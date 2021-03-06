// Code generated by sqlc. DO NOT EDIT.
// source: cdn.sql

package db

import (
	"context"
)

const createAsset = `-- name: CreateAsset :one
WITH fid AS (
    INSERT into files (id, mimetype, name, data)
        VALUES (create_snowflake(), $1, $2, $3)
        RETURNING id
)
INSERT INTO assets (id, name, url_path, file_id, creator_id)
    VALUES (create_snowflake(), $4, $5, (SELECT id from fid), $6)
    RETURNING id
`

type CreateAssetParams struct {
	Mimetype  string `json:"mimetype"`
	Name      string `json:"name"`
	Data      []byte `json:"data"`
	Name_2    string `json:"name2"`
	UrlPath   string `json:"urlPath"`
	CreatorID int64  `json:"creatorID"`
}

func (q *Queries) CreateAsset(ctx context.Context, arg CreateAssetParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, createAsset,
		arg.Mimetype,
		arg.Name,
		arg.Data,
		arg.Name_2,
		arg.UrlPath,
		arg.CreatorID,
	)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const createUser = `-- name: CreateUser :exec


INSERT INTO users (id, username, discriminator) VALUES ($1, $2, $3)
`

type CreateUserParams struct {
	ID            int64  `json:"id"`
	Username      string `json:"username"`
	Discriminator string `json:"discriminator"`
}

// ((Pagenumber - 1) * PageSize)
func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) error {
	_, err := q.db.ExecContext(ctx, createUser, arg.ID, arg.Username, arg.Discriminator)
	return err
}

const deleteAsset = `-- name: DeleteAsset :exec
DELETE
    FROM files
    WHERE id = (SELECT file_id FROM assets WHERE url_path = $1 AND creator_id = $2)
`

type DeleteAssetParams struct {
	UrlPath   string `json:"urlPath"`
	CreatorID int64  `json:"creatorID"`
}

func (q *Queries) DeleteAsset(ctx context.Context, arg DeleteAssetParams) error {
	_, err := q.db.ExecContext(ctx, deleteAsset, arg.UrlPath, arg.CreatorID)
	return err
}

const getAssetDetailsById = `-- name: GetAssetDetailsById :one
SELECT url_path, name, creator_id
    FROM assets
    WHERE id = $1
`

type GetAssetDetailsByIdRow struct {
	UrlPath   string `json:"urlPath"`
	Name      string `json:"name"`
	CreatorID int64  `json:"creatorID"`
}

func (q *Queries) GetAssetDetailsById(ctx context.Context, id int64) (GetAssetDetailsByIdRow, error) {
	row := q.db.QueryRowContext(ctx, getAssetDetailsById, id)
	var i GetAssetDetailsByIdRow
	err := row.Scan(&i.UrlPath, &i.Name, &i.CreatorID)
	return i, err
}

const getAssetDetailsByUrl = `-- name: GetAssetDetailsByUrl :one
SELECT id, name, creator_id
    FROM assets
    WHERE url_path = $1
`

type GetAssetDetailsByUrlRow struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	CreatorID int64  `json:"creatorID"`
}

func (q *Queries) GetAssetDetailsByUrl(ctx context.Context, urlPath string) (GetAssetDetailsByUrlRow, error) {
	row := q.db.QueryRowContext(ctx, getAssetDetailsByUrl, urlPath)
	var i GetAssetDetailsByUrlRow
	err := row.Scan(&i.ID, &i.Name, &i.CreatorID)
	return i, err
}

const getFile = `-- name: GetFile :one
SELECT data, mimetype
    FROM files f
    WHERE f.id = (
        SELECT file_id
            FROM assets a
            WHERE a.url_path = $1
    )
`

type GetFileRow struct {
	Data     []byte `json:"data"`
	Mimetype string `json:"mimetype"`
}

func (q *Queries) GetFile(ctx context.Context, urlPath string) (GetFileRow, error) {
	row := q.db.QueryRowContext(ctx, getFile, urlPath)
	var i GetFileRow
	err := row.Scan(&i.Data, &i.Mimetype)
	return i, err
}

const listAssetByCreator = `-- name: ListAssetByCreator :many
SELECT id, name, url_path, file_id, creator_id
FROM assets
WHERE creator_id = $1
ORDER BY id
LIMIT $2 -- PageSize
    OFFSET $3
`

type ListAssetByCreatorParams struct {
	CreatorID int64 `json:"creatorID"`
	Limit     int32 `json:"limit"`
	Offset    int32 `json:"offset"`
}

func (q *Queries) ListAssetByCreator(ctx context.Context, arg ListAssetByCreatorParams) ([]Assets, error) {
	rows, err := q.db.QueryContext(ctx, listAssetByCreator, arg.CreatorID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Assets
	for rows.Next() {
		var i Assets
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.UrlPath,
			&i.FileID,
			&i.CreatorID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
