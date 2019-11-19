package entity

import (
	"go-web/model"
	"log"
)

type Student struct {
	Id   int    `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Name string `json:"name"`
	Sex  int    `json:"sex"`
}

func Save(name string, sex int) (err error) {
	student := Student{Name: name, Sex: sex}
	e := db.Create(&student).Error
	if e != nil {
		log.Println("save student fail ", e)
		return e
	}
	return nil
}

func Delete(id int) (err error) {
	e := db.Where("id = ?", id).Delete(&Student{}).Error
	if e != nil {
		log.Println("delete student fail ", e)
		return e
	}
	return nil
}

func Page(page model.Page) (student []*Student, err error) {
	var students []*Student
	e := db.Offset(page.GetOffset()).Limit(page.Size).Find(&students).Error
	if e != nil {
		log.Println("select student fail ", e)
		return nil, e
	}
	return students, nil
}

func Update(name string, id int) (err error) {
	student := Student{Name: name, Id: id}
	e := db.Model(&student).Update(&student).Error
	if e != nil {
		log.Println("update student fail ", e)
		return e
	}
	return nil
}
