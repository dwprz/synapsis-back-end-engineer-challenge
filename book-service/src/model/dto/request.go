package dto

type AddBookReq struct {
	Title         string  `json:"title" validate:"required,max=100"`
	Author        string  `json:"author" validate:"required,max=100"`
	ISBN          string  `json:"isbn" validate:"required,max=50"`
	Synopsis      *string `json:"synopsis" validate:"omitempty,max=500"`
	PublishedYear int     `json:"published_year" validate:"required"`
	Stock         int     `json:"stock" validate:"required"`
	Location      string  `json:"location" validate:"required,max=100"`
}

type GetBookReq struct {
	Page          int    `json:"page" validate:"required,max=100"`
	BookId        int    `json:"book_id" validate:"omitempty"`
	Title         string `json:"title" validate:"omitempty,max=100"`
	Author        string `json:"author" validate:"omitempty,max=100"`
	ISBN          string `json:"isbn" validate:"omitempty,max=50"`
	PublishedYear int    `json:"published_year" validate:"omitempty"`
	Stock         *int   `json:"stock" validate:"omitempty"`
	Location      string `json:"location" validate:"omitempty,max=100"`
}

type UpdateBookReq struct {
	BookId        int     `json:"book_id" validate:"required"`
	Title         string  `json:"title" validate:"omitempty,max=100"`
	Author        string  `json:"author" validate:"omitempty,max=100"`
	ISBN          string  `json:"isbn" validate:"omitempty,max=50"`
	Synopsis      *string `json:"synopsis" validate:"omitempty,max=500"`
	PublishedYear int     `json:"published_year" validate:"omitempty"`
	Stock         *int    `json:"stock" validate:"omitempty"`
	Location      string  `json:"location" validate:"omitempty,max=100"`
}

type FindManyByIdsReq struct {
	BookIds []uint32 `json:"book_ids" validate:"max=20,dive,required"`
}
