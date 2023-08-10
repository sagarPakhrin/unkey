// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: api_find.sql

package database

import (
	"context"
)

const findApi = `-- name: FindApi :one
SELECT
    id, name, workspace_id, ip_whitelist, auth_type, key_auth_id
FROM
    apis
WHERE
    id = ?
`

func (q *Queries) FindApi(ctx context.Context, id string) (Api, error) {
	row := q.db.QueryRowContext(ctx, findApi, id)
	var i Api
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.WorkspaceID,
		&i.IpWhitelist,
		&i.AuthType,
		&i.KeyAuthID,
	)
	return i, err
}
