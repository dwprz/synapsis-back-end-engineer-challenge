package entity

import "time"

type Book struct {
	BookId        int       `json:"book_id"`
	Title         string    `json:"title"`
	Author        string    `json:"author"`
	ISBN          string    `json:"isbn"`
	Synopsis      *string   `json:"synopsis"`
	PublishedYear int       `json:"published_year"`
	Stock         int       `json:"stock"`
	Location      string    `json:"location"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type BookId struct {
	Id int `json:"book_id" validate:"required"`
}
