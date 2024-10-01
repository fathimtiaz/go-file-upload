package repository

import (
	"context"
	"go-file-upload/internal/domain"
)

type FileRepo interface {
	SaveFile(context.Context, *domain.File) error
	UpdateFile(context.Context, domain.File) error
	GetFile(ctx context.Context, query FileQuery) (domain.File, error)

	SaveFileChunk(ctx context.Context, file *domain.FileChunk) error
	GetFileChunk(context.Context, FileChunkQuery) (domain.FileChunk, error)
	GetFileChunkIds(context.Context, FileQuery) (ids []int, err error)

	WrapTx(context.Context, func(FileRepo) error) error
}
