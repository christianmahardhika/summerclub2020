// models/todo.go

package models

type Twitter struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Username string `json:"username"`
	Post     string `json:"post"`
}
