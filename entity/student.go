package entity

import (
	"go-web/model"
	"log"
)

type Student struct {
	Id   int    `xorm:"pk autoincr" json:"id"`
	Name string `xorm:"varchar(50) name" json:"name"`
	Sex  int    `xorm:"varchar(50) sex" json:"sex"`
}

func (Student) TableName() string {
	return "t_student"
}

func Save(name string, sex int) (err error) {
	student := Student{Name: name, Sex: sex}
	_, e := db.Insert(&student)
	if e != nil {
		log.Println("save student fail ", e)
		return e
	}
	return nil
}

func Delete(id int) (err error) {
	_, e := db.Where("id = ?", id).Delete(&Student{})
	if e != nil {
		log.Println("delete student fail ", e)
		return e
	}
	return nil
}

func Page(page model.Page) (student []*Student, err error) {
	var students []*Student
	e := db.Limit(page.Size, page.GetOffset()).Find(&students)
	if e != nil {
		log.Println("select student fail ", e)
		return nil, e
	}
	return students, nil
}

func Update(name string, id int) (err error) {
	student := Student{Name: name, Id: id}
	_, e := db.Id(id).Update(&student)
	if e != nil {
		log.Println("update student fail ", e)
		return e
	}
	return nil
}
