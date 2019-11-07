package entity

type Student struct {
	Id   int `gorm:"primary_key"`
	Name string
	Sex  int
}
