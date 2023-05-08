// Code generated by sqlc. DO NOT EDIT.

package search_queries

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.addAuthorStmt, err = db.PrepareContext(ctx, addAuthor); err != nil {
		return nil, fmt.Errorf("error preparing query AddAuthor: %w", err)
	}
	if q.addPostStmt, err = db.PrepareContext(ctx, addPost); err != nil {
		return nil, fmt.Errorf("error preparing query AddPost: %w", err)
	}
	if q.getAuthorStmt, err = db.PrepareContext(ctx, getAuthor); err != nil {
		return nil, fmt.Errorf("error preparing query GetAuthor: %w", err)
	}
	if q.getAuthorStatsStmt, err = db.PrepareContext(ctx, getAuthorStats); err != nil {
		return nil, fmt.Errorf("error preparing query GetAuthorStats: %w", err)
	}
	if q.getAuthorsByHandleStmt, err = db.PrepareContext(ctx, getAuthorsByHandle); err != nil {
		return nil, fmt.Errorf("error preparing query GetAuthorsByHandle: %w", err)
	}
	if q.getOldestPresentParentStmt, err = db.PrepareContext(ctx, getOldestPresentParent); err != nil {
		return nil, fmt.Errorf("error preparing query GetOldestPresentParent: %w", err)
	}
	if q.getPostStmt, err = db.PrepareContext(ctx, getPost); err != nil {
		return nil, fmt.Errorf("error preparing query GetPost: %w", err)
	}
	if q.getThreadViewStmt, err = db.PrepareContext(ctx, getThreadView); err != nil {
		return nil, fmt.Errorf("error preparing query GetThreadView: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.addAuthorStmt != nil {
		if cerr := q.addAuthorStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing addAuthorStmt: %w", cerr)
		}
	}
	if q.addPostStmt != nil {
		if cerr := q.addPostStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing addPostStmt: %w", cerr)
		}
	}
	if q.getAuthorStmt != nil {
		if cerr := q.getAuthorStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getAuthorStmt: %w", cerr)
		}
	}
	if q.getAuthorStatsStmt != nil {
		if cerr := q.getAuthorStatsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getAuthorStatsStmt: %w", cerr)
		}
	}
	if q.getAuthorsByHandleStmt != nil {
		if cerr := q.getAuthorsByHandleStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getAuthorsByHandleStmt: %w", cerr)
		}
	}
	if q.getOldestPresentParentStmt != nil {
		if cerr := q.getOldestPresentParentStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getOldestPresentParentStmt: %w", cerr)
		}
	}
	if q.getPostStmt != nil {
		if cerr := q.getPostStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getPostStmt: %w", cerr)
		}
	}
	if q.getThreadViewStmt != nil {
		if cerr := q.getThreadViewStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getThreadViewStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                         DBTX
	tx                         *sql.Tx
	addAuthorStmt              *sql.Stmt
	addPostStmt                *sql.Stmt
	getAuthorStmt              *sql.Stmt
	getAuthorStatsStmt         *sql.Stmt
	getAuthorsByHandleStmt     *sql.Stmt
	getOldestPresentParentStmt *sql.Stmt
	getPostStmt                *sql.Stmt
	getThreadViewStmt          *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                         tx,
		tx:                         tx,
		addAuthorStmt:              q.addAuthorStmt,
		addPostStmt:                q.addPostStmt,
		getAuthorStmt:              q.getAuthorStmt,
		getAuthorStatsStmt:         q.getAuthorStatsStmt,
		getAuthorsByHandleStmt:     q.getAuthorsByHandleStmt,
		getOldestPresentParentStmt: q.getOldestPresentParentStmt,
		getPostStmt:                q.getPostStmt,
		getThreadViewStmt:          q.getThreadViewStmt,
	}
}
