// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package dao

import ()

type Task struct {
	TaskID      int64
	UserID      int64
	Body        string
	IsCompleted int64
	CreatedAt   string
	UpdatedAt   string
}

type User struct {
	UserID    int64
	Username  string
	CreatedAt string
	UpdatedAt string
}
