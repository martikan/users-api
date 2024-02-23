package common

type Pageable interface {
	GetSize() int
	GetOffset() int
}

type DefaultPageable struct {
	Page int `form:"page" binding:"min=1" default:"1"`
	Size int `form:"size" binding:"min=5,max=50" default:"20"`
}

func (d DefaultPageable) GetSize() int {
	return d.Size
}

func (d DefaultPageable) GetOffset() int {
	return (d.Page - 1) * d.Size
}
