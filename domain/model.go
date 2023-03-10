// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package domain

type Node interface {
	IsNode()
	GetID() string
}

type Task struct {
	ID          string `json:"id"`
	TaskID      string `json:"taskId"`
	Body        string `json:"body"`
	IsCompleted bool   `json:"isCompleted"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

func (Task) IsNode()            {}
func (this Task) GetID() string { return this.ID }

type User struct {
	ID        string `json:"id"`
	UserID    string `json:"userId"`
	Username  string `json:"username"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

func (User) IsNode()            {}
func (this User) GetID() string { return this.ID }

type UserCreateInput struct {
	Name string `json:"name"`
}
