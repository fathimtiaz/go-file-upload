package domain

import (
	"mime/multipart"
	"time"
)

type File struct {
	Id          int
	Name        string
	NumOfChunks int
	ChunkSize   int

	Timestamp
}

func NewFile(data multipart.File, name string) File {
	return File{
		Name:      name,
		Timestamp: Timestamp{CreatedAt: time.Now()},
	}
}

type FileChunk struct {
	Id     int
	FileId int
	Data   []byte
	Size   int
	Index  int

	Timestamp
}

func NewChunk(fileId, index, size int) FileChunk {
	return FileChunk{
		FileId:    fileId,
		Index:     index,
		Size:      size,
		Data:      make([]byte, size),
		Timestamp: Timestamp{CreatedAt: time.Now()},
	}
}
