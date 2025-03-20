// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: friendship.sql

package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const addFriend = `-- name: AddFriend :one
insert into friend_lists(
    user_id,
    friend_id,
    status
) values ( $1,$2,$3) returning id, user_id, friend_id, status, created_at, updated_at
`

type AddFriendParams struct {
	UserID   uuid.UUID
	FriendID uuid.UUID
	Status   sql.NullString
}

func (q *Queries) AddFriend(ctx context.Context, arg AddFriendParams) (FriendList, error) {
	row := q.db.QueryRow(ctx, addFriend, arg.UserID, arg.FriendID, arg.Status)
	var i FriendList
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.FriendID,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getFiendByUserName = `-- name: GetFiendByUserName :one
select id, user_name, email, password, created_at, updated_at from users where email = $1
`

func (q *Queries) GetFiendByUserName(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRow(ctx, getFiendByUserName, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.UserName,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getFriendList = `-- name: GetFriendList :many

select users.id,user_name,email,password from users join friend_lists
on users.id=friend_lists.user_id
`

type GetFriendListRow struct {
	ID       uuid.UUID
	UserName string
	Email    string
	Password string
}

func (q *Queries) GetFriendList(ctx context.Context) ([]GetFriendListRow, error) {
	rows, err := q.db.Query(ctx, getFriendList)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetFriendListRow
	for rows.Next() {
		var i GetFriendListRow
		if err := rows.Scan(
			&i.ID,
			&i.UserName,
			&i.Email,
			&i.Password,
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

const updateFriendStatus = `-- name: UpdateFriendStatus :one
update friend_lists 
set status=$3 where 
user_id=$1 and friend_id=$2 
returning id, user_id, friend_id, status, created_at, updated_at
`

type UpdateFriendStatusParams struct {
	UserID   uuid.UUID
	FriendID uuid.UUID
	Status   sql.NullString
}

func (q *Queries) UpdateFriendStatus(ctx context.Context, arg UpdateFriendStatusParams) (FriendList, error) {
	row := q.db.QueryRow(ctx, updateFriendStatus, arg.UserID, arg.FriendID, arg.Status)
	var i FriendList
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.FriendID,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
