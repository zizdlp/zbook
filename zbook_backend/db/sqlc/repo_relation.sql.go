// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: repo_relation.sql

package db

import (
	"context"
	"time"
)

const createRepoRelation = `-- name: CreateRepoRelation :exec
INSERT INTO repo_relations (
  user_id,
  repo_id,
  relation_type
) VALUES ($1,$2,$3)
`

type CreateRepoRelationParams struct {
	UserID       int64  `json:"user_id"`
	RepoID       int64  `json:"repo_id"`
	RelationType string `json:"relation_type"`
}

func (q *Queries) CreateRepoRelation(ctx context.Context, arg CreateRepoRelationParams) error {
	_, err := q.db.Exec(ctx, createRepoRelation, arg.UserID, arg.RepoID, arg.RelationType)
	return err
}

const deleteRepoRelation = `-- name: DeleteRepoRelation :exec
DELETE FROM repo_relations
WHERE user_id=$1 and repo_id=$2 and relation_type = $3
`

type DeleteRepoRelationParams struct {
	UserID       int64  `json:"user_id"`
	RepoID       int64  `json:"repo_id"`
	RelationType string `json:"relation_type"`
}

func (q *Queries) DeleteRepoRelation(ctx context.Context, arg DeleteRepoRelationParams) error {
	_, err := q.db.Exec(ctx, deleteRepoRelation, arg.UserID, arg.RepoID, arg.RelationType)
	return err
}

const getRepoRelation = `-- name: GetRepoRelation :one
SELECT relation_id, relation_type, user_id, repo_id, created_at
FROM repo_relations
WHERE user_id = $1 and repo_id=$2 and relation_type = $3
`

type GetRepoRelationParams struct {
	UserID       int64  `json:"user_id"`
	RepoID       int64  `json:"repo_id"`
	RelationType string `json:"relation_type"`
}

func (q *Queries) GetRepoRelation(ctx context.Context, arg GetRepoRelationParams) (RepoRelation, error) {
	row := q.db.QueryRow(ctx, getRepoRelation, arg.UserID, arg.RepoID, arg.RelationType)
	var i RepoRelation
	err := row.Scan(
		&i.RelationID,
		&i.RelationType,
		&i.UserID,
		&i.RepoID,
		&i.CreatedAt,
	)
	return i, err
}

const getSelectedUserByRepoCount = `-- name: GetSelectedUserByRepoCount :one
SELECT COUNT(*)
FROM repos as r
LEFT JOIN repo_relations as rr ON rr.repo_id=r.repo_id
JOIN users as u ON u.user_id = rr.user_id
WHERE r.repo_id=$1 AND rr.relation_type = 'visi'
`

func (q *Queries) GetSelectedUserByRepoCount(ctx context.Context, repoID int64) (int64, error) {
	row := q.db.QueryRow(ctx, getSelectedUserByRepoCount, repoID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const listSelectedUserByRepo = `-- name: ListSelectedUserByRepo :many
SELECT u.user_id, u.username, u.email, u.hashed_password, u.blocked, u.verified, u.motto, u.user_role, u.onboarding, u.created_at, u.updated_at, u.unread_count, u.unread_count_updated_at, u.fts_username
FROM repos as r
LEFT JOIN repo_relations as rr ON rr.repo_id=r.repo_id
JOIN users as u ON u.user_id = rr.user_id
WHERE r.repo_id=$3 AND rr.relation_type = 'visi'
ORDER BY rr.created_at DESC
LIMIT $1
OFFSET $2
`

type ListSelectedUserByRepoParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
	RepoID int64 `json:"repo_id"`
}

func (q *Queries) ListSelectedUserByRepo(ctx context.Context, arg ListSelectedUserByRepoParams) ([]User, error) {
	rows, err := q.db.Query(ctx, listSelectedUserByRepo, arg.Limit, arg.Offset, arg.RepoID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []User{}
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.UserID,
			&i.Username,
			&i.Email,
			&i.HashedPassword,
			&i.Blocked,
			&i.Verified,
			&i.Motto,
			&i.UserRole,
			&i.Onboarding,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.UnreadCount,
			&i.UnreadCountUpdatedAt,
			&i.FtsUsername,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const querySelectedUserByRepo = `-- name: QuerySelectedUserByRepo :many
SELECT
   u.user_id, u.username, u.email, u.hashed_password, u.blocked, u.verified, u.motto, u.user_role, u.onboarding, u.created_at, u.updated_at, u.unread_count, u.unread_count_updated_at, u.fts_username,
   CASE WHEN MAX(rr.user_id) IS NOT NULL THEN true ELSE false END AS is_visible
FROM 
  users as u 
LEFT JOIN 
    repo_relations rr ON rr.user_id = u.user_id AND rr.repo_id=$3
WHERE u.username=$4
GROUP BY u.user_id,rr.created_at
ORDER BY rr.created_at DESC
LIMIT $1
OFFSET $2
`

type QuerySelectedUserByRepoParams struct {
	Limit    int32  `json:"limit"`
	Offset   int32  `json:"offset"`
	RepoID   int64  `json:"repo_id"`
	Username string `json:"username"`
}

type QuerySelectedUserByRepoRow struct {
	UserID               int64     `json:"user_id"`
	Username             string    `json:"username"`
	Email                string    `json:"email"`
	HashedPassword       string    `json:"hashed_password"`
	Blocked              bool      `json:"blocked"`
	Verified             bool      `json:"verified"`
	Motto                string    `json:"motto"`
	UserRole             string    `json:"user_role"`
	Onboarding           bool      `json:"onboarding"`
	CreatedAt            time.Time `json:"created_at"`
	UpdatedAt            time.Time `json:"updated_at"`
	UnreadCount          int32     `json:"unread_count"`
	UnreadCountUpdatedAt time.Time `json:"unread_count_updated_at"`
	FtsUsername          string    `json:"fts_username"`
	IsVisible            bool      `json:"is_visible"`
}

func (q *Queries) QuerySelectedUserByRepo(ctx context.Context, arg QuerySelectedUserByRepoParams) ([]QuerySelectedUserByRepoRow, error) {
	rows, err := q.db.Query(ctx, querySelectedUserByRepo,
		arg.Limit,
		arg.Offset,
		arg.RepoID,
		arg.Username,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []QuerySelectedUserByRepoRow{}
	for rows.Next() {
		var i QuerySelectedUserByRepoRow
		if err := rows.Scan(
			&i.UserID,
			&i.Username,
			&i.Email,
			&i.HashedPassword,
			&i.Blocked,
			&i.Verified,
			&i.Motto,
			&i.UserRole,
			&i.Onboarding,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.UnreadCount,
			&i.UnreadCountUpdatedAt,
			&i.FtsUsername,
			&i.IsVisible,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
