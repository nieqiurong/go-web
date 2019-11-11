package model

type Page struct {
	Page int `form:"page" binding:"gt=0"`
	Size int `form:"size" binding:"gt=0"`
}

func (p Page) GetOffset() int {
	return (p.Page - 1) * p.Size
}
