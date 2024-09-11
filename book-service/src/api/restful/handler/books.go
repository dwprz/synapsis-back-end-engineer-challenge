package handler

import (
	"book-service/src/common/helper"
	"book-service/src/interface/service"
	"book-service/src/model/dto"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Book struct {
	bookService service.Book
}

func NewBook(bs service.Book) *Book {
	return &Book{
		bookService: bs,
	}
}

func (h *Book) Add(c *fiber.Ctx) error {
	req := new(dto.AddBookReq)
	if err := c.BodyParser(req); err != nil {
		return err
	}

	if err := h.bookService.Add(c.Context(), req); err != nil {
		return err
	}

	return c.Status(201).JSON(fiber.Map{"data": "successfully added the book"})
}

func (h *Book) Get(c *fiber.Ctx) error {
	req, err := helper.ParseGetBookQueryParams(c)
	if err != nil {
		return err
	}

	res, err := h.bookService.FindMany(c.Context(), req)
	if err != nil {
		return err
	}

	return c.Status(200).JSON(fiber.Map{"data": res.Data, "paging": res.Paging})
}

func (h *Book) GetRecommendations(c *fiber.Ctx) error {
	res, err := h.bookService.FindManyPopularBook(c.Context())
	if err != nil {
		return err
	}

	return c.Status(200).JSON(fiber.Map{"data": res})
}

func (h *Book) Update(c *fiber.Ctx) error {
	bookId, err := strconv.Atoi(c.Params("bookId"))
	if err != nil {
		return err
	}

	req := new(dto.UpdateBookReq)
	if err := c.BodyParser(req); err != nil {
		return err
	}

	req.BookId = bookId

	res, err := h.bookService.Update(c.Context(), req)
	if err != nil {
		return err
	}

	return c.Status(200).JSON(fiber.Map{"data": res})
}

func (h *Book) Delete(c *fiber.Ctx) error {
	bookId, err := strconv.Atoi(c.Params("bookId"))
	if err != nil {
		return err
	}

	err = h.bookService.Delete(c.Context(), bookId)
	if err != nil {
		return err
	}

	return c.Status(200).JSON(fiber.Map{"data": "successfully deleted the book"})
}
