package dto

import "book-category-service/src/model/entity"

type CreateBookCategoryReq struct {
	Category entity.Category `json:"category" validate:"required,bookcategory"`
	BookIds  []int               `json:"book_ids" validate:"dive,required"`
}

type DeleteBookCategoryReq struct {
	Category *entity.Category `json:"category" validate:"omitempty,bookcategory"`
	BookId   *int                 `json:"book_id" validate:"omitempty"`
}

type FindManyByCategoryReq struct {
	Category entity.Category `json:"category" validate:"required,bookcategory"`
	Page     int                 `json:"page" validate:"required,max=100"`
}
