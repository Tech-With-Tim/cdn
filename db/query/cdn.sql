-- name: CreateAsset :one
WITH fid AS (
    INSERT into files (id, mimetype, name, data)
        VALUES (create_snowflake(), $1, $2, $3)
        RETURNING id
)
INSERT INTO assets (id, name, url_path, file_id, creator_id)
    VALUES (create_snowflake(), $4, $5, (SELECT id from fid), $6)
    RETURNING id;

-- name: GetFile :one
SELECT data, mimetype
    FROM files f
    WHERE f.id = (
        SELECT file_id
            FROM assets a
            WHERE a.url_path = $1
    );

-- name: GetAssetDetailsByUrl :one
SELECT id, name, creator_id
    FROM assets
    WHERE url_path = $1;

-- name: GetAssetDetailsById :one
SELECT url_path, name, creator_id
    FROM assets
    WHERE id = $1;

-- name: DeleteAsset :exec
DELETE
    FROM files
    WHERE id = (SELECT file_id FROM assets WHERE url_path = $1 AND creator_id = $2);

-- name: ListAssetByCreator :many
SELECT *
FROM assets
WHERE creator_id = $1
ORDER BY id
LIMIT $2 -- PageSize
    OFFSET $3; -- ((Pagenumber - 1) * PageSize)
