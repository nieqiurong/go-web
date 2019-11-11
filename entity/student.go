package entity

import (
	"go-web/model"
)

type Student struct {
	Id   int `gorm:"primary_key;AUTO_INCREMENT"`
	Name string
	Sex  int
}

func Save(name string, sex int) {
	student := Student{Name: name, Sex: sex}
	db.Create(&student)
}

func Delete(id int) {
	db.Where("id = ?", id).Delete(&Student{})
}

func Page(page model.Page) (student []*Student, err error) {
	var students []*Student
	e := db.Offset(page.GetOffset()).Limit(page.Size).Find(&students).Error
	return students, e
}
