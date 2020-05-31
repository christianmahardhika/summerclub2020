// models/todo.go

package models

type Todo struct {
  ID     uint   `json:"id" gorm:"primary_key"`
  Title  string `json:"title"`
  Note  string `json:"note"`
}