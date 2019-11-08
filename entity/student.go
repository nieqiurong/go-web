package entity

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
