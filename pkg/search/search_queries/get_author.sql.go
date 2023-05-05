// Code generated by sqlc. DO NOT EDIT.
// source: get_author.sql

package search_queries

import (
	"context"
)

const getAuthor = `-- name: GetAuthor :one
SELECT did, handle
FROM authors
WHERE did = $1
`

func (q *Queries) GetAuthor(ctx context.Context, did string) (Author, error) {
	row := q.queryRow(ctx, q.getAuthorStmt, getAuthor, did)
	var i Author
	err := row.Scan(&i.Did, &i.Handle)
	return i, err
}