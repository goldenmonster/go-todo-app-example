package models

import (
	"example/todo-app/database"
	"fmt"

	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Title string `gorm:"size:255;not null;unique" json:"name" binding:"required"`  // `key: "value" key:"value"` uri: 
	Hours int `gorm:"not null" json:"hours" binding:"required"`
	Done bool `gorm:"not null" json:"done,omitempty" binding:"required"`
}

func GetAllTodos(todos *[]Todo) (err error) {
	if err = database.Database.Find(todos).Error; err != nil {
		return err
	}

	return nil
}

func CreateATodo(todo *Todo) (err error) {
	if err = database.Database.Create(todo).Error; err != nil {
		return err
	}

	return nil
}

func GetATodo(todo *Todo, id string) (err error) {
	if err := database.Database.Where("id = ?", id).First(todo).Error; err != nil {
		return err
	}
	return nil
}

func UpdateATodo(todo *Todo, id string) (err error) {
	fmt.Println(todo)
	database.Database.Save(todo)
	return nil
}

func DeleteATodo(todo *Todo, id string) (err error) {
	database.Database.Where("id = ?", id).Delete(todo)
	return nil

}


func (todo *Todo) Save() (*Todo, error) {
	err := database.Database.Create(&todo).Error
	if err != nil {
		return &Todo{}, err
	}

	return todo, nil
}


func (todo *Todo) BeforeSave(tx *gorm.DB) {}