package helper

import (
	"book-service/src/model/dto"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func ParseGetBookQueryParams(c *fiber.Ctx) (*dto.GetBookReq, error) {
	req := new(dto.GetBookReq)

	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil {
		return nil, err
	}

	req.Page = page

	if bookIdStr := c.Query("bookId"); bookIdStr != "" {
		bookId, err := strconv.Atoi(bookIdStr)
		if err != nil {
			return nil, err
		}

		req.BookId = bookId
	}

	req.Title = c.Query("title")
	req.Author = c.Query("author")

	if yearStr := c.Query("publishedYear"); yearStr != "" {
		year, err := strconv.Atoi(yearStr)
		if err != nil {
			return nil, err
		}

		req.PublishedYear = year
	}

	if stockStr := c.Query("stock"); stockStr != "" {
		stock, err := strconv.Atoi(stockStr)
		if err != nil {
			return nil, err
		}

		req.Stock = &stock
	}

	req.Location = c.Query("location")

	return req, nil
}
