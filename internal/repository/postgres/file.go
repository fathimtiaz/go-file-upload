package postgres

import (
	"context"
	"database/sql"
	"go-file-upload/internal/domain"
	"go-file-upload/internal/repository"
)

func (db *sqlDB) SaveFile(ctx context.Context, file *domain.File) error {
	return db.ddb.QueryRowContext(ctx, `
		INSERT INTO file_ (name_, chunk_size, num_of_chunks, created_at)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`,
		file.Name,
		file.ChunkSize,
		file.NumOfChunks,
		file.CreatedAt,
	).Scan(&file.Id)
}

func (db *sqlDB) UpdateFile(ctx context.Context, file domain.File) (err error) {
	_, err = db.ddb.ExecContext(ctx, `
		UPDATE file_ SET name_ = $1, chunk_size = $2, num_of_chunks = $3, updated_at = $4 WHERE id = $5
	`, file.Name, file.ChunkSize, file.NumOfChunks, file.UpdatedAt, file.Id)

	return
}

func (db *sqlDB) SaveFileChunk(ctx context.Context, file *domain.FileChunk) (err error) {
	_, err = db.ddb.ExecContext(ctx, `
			INSERT INTO file_chunk_ (file_id, size, index_, data_, created_at)
			VALUES ($1, $2, $3, $4, $5)`,
		file.FileId, file.Size, file.Index, file.Data, file.CreatedAt,
	)

	return
}

func (db *sqlDB) GetFile(ctx context.Context, query repository.FileQuery) (file domain.File, err error) {
	var fileCreatedAt, fileUpdatedAt sql.NullTime

	if err = db.ddb.QueryRowContext(ctx, `
		SELECT id, name_, chunk_size, num_of_chunks, created_at, updated_at
		FROM file_
		WHERE id = $1
	`, query.Id).Scan(
		&file.Id,
		&file.Name,
		&file.ChunkSize,
		&file.NumOfChunks,
		&fileCreatedAt,
		&fileUpdatedAt,
	); err != nil {
		return
	}

	file.CreatedAt = fileCreatedAt.Time
	file.UpdatedAt = fileUpdatedAt.Time

	return
}

func (db *sqlDB) GetFileChunk(ctx context.Context, query repository.FileChunkQuery) (chunk domain.FileChunk, err error) {
	var createdAt, updatedAt sql.NullTime

	if err = db.ddb.QueryRowContext(ctx, `
		SELECT id, file_id, size, index_, data_, created_at, updated_at
		FROM file_chunk_
		WHERE id = $1
	`, query.Id).Scan(
		&chunk.Id,
		&chunk.FileId,
		&chunk.Size,
		&chunk.Index,
		&chunk.Data,
		&createdAt,
		&updatedAt,
	); err != nil {
		return
	}

	chunk.CreatedAt = createdAt.Time
	chunk.UpdatedAt = updatedAt.Time

	return
}

func (db *sqlDB) GetFileChunkIds(ctx context.Context, query repository.FileQuery) (ids []int, err error) {
	rows, err := db.ddb.QueryContext(ctx, `
		SELECT id
		FROM file_chunk_
		WHERE file_id = $1
	`, query.Id)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var id int

		if err = rows.Scan(&id); err != nil {
			return
		}

		ids = append(ids, id)
	}

	return
}
