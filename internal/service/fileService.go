package service

import (
	"context"
	"errors"
	"go-file-upload/config"
	"go-file-upload/internal/domain"
	"go-file-upload/internal/repository"
	"io"
	"mime/multipart"
	"sync"
	"time"
)

type FileSvc struct {
	cfg      *config.Config
	fileRepo repository.FileRepo
}

func NewFileSvc(cfg *config.Config, fileRepo repository.FileRepo) *FileSvc {
	return &FileSvc{cfg, fileRepo}
}

func (s *FileSvc) SaveFile(ctx context.Context, reqFile multipart.File, reqFileHeader *multipart.FileHeader) (file domain.File, err error) {
	if reqFileHeader.Size > s.cfg.FIle.MaxSize.Int64() {
		return file, errors.New("file too big")
	}

	file = domain.NewFile(reqFile, reqFileHeader.Filename)

	if err = s.fileRepo.WrapTx(ctx, func(repo repository.FileRepo) (err error) {
		if err = repo.SaveFile(ctx, &file); err != nil {
			return
		}

		var bytesRead int
		var index int
		var chunkSize int = s.cfg.FIle.ChunkSize.Int()
		var buffer = make([]byte, chunkSize)
		var wg = &sync.WaitGroup{}

		for {
			bytesRead, err = reqFile.Read(buffer)
			if err != nil && err != io.EOF {
				return
			}
			err = nil
			if bytesRead == 0 {
				break
			}

			var chunk = domain.NewChunk(file.Id, index, chunkSize)
			copy(chunk.Data, buffer[:bytesRead])
			index++

			wg.Add(1)
			go func() {
				defer wg.Done()
				if err = repo.SaveFileChunk(ctx, &chunk); err != nil {
					return
				}
			}()
		}

		wg.Wait()

		file.ChunkSize = chunkSize
		file.NumOfChunks = index + 1
		file.UpdatedAt = time.Now()

		return
	}); err != nil {
		return
	}

	if err = s.fileRepo.UpdateFile(ctx, file); err != nil {
		return
	}

	return
}

func (s *FileSvc) GetFileInfo(ctx context.Context, query repository.FileQuery) (result domain.File, err error) {
	return s.fileRepo.GetFile(ctx, query)
}

func (s *FileSvc) GetFileData(ctx context.Context, query repository.FileQuery) (fileName string, data []byte, err error) {
	var file domain.File
	var ids []int

	if file, err = s.fileRepo.GetFile(ctx, query); err != nil {
		return
	}

	if ids, err = s.fileRepo.GetFileChunkIds(ctx, query); err != nil {
		return
	}

	var byteData = make([]byte, file.ChunkSize*file.NumOfChunks)
	var wg = &sync.WaitGroup{}

	wg.Add(len(ids))
	for i := range ids {
		go func(id int) {
			defer wg.Done()

			var chunk domain.FileChunk

			if chunk, err = s.fileRepo.GetFileChunk(ctx, repository.FileChunkQuery{Id: id}); err != nil {
				return
			}

			copy(byteData[chunk.Index*chunk.Size:(chunk.Index+1)*chunk.Size], chunk.Data)
		}(ids[i])
	}

	wg.Wait()

	data = byteData
	fileName = file.Name

	return
}
