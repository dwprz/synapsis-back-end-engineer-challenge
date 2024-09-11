package handler

import (
	"book-category-service/src/interface/service"
	"book-category-service/src/model/dto"
	"book-category-service/src/model/entity"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type BookCategory struct {
	bookCategoryService service.BookCategory
}

func NewBookCategory(bs service.BookCategory) *BookCategory {
	return &BookCategory{
		bookCategoryService: bs,
	}
}

func (h *BookCategory) Create(c *fiber.Ctx) error {
	req := new(dto.CreateBookCategoryReq)
	if err := c.BodyParser(&req); err != nil {
		return err
	}

	if err := h.bookCategoryService.Create(c.Context(), req); err != nil {
		return err
	}

	return c.Status(201).JSON(fiber.Map{"data": "successfully created the category book"})
}

func (h *BookCategory) GetByCategory(c *fiber.Ctx) error {
	category := strings.ToUpper(c.Params("category"))

	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil {
		return err
	}

	res, err := h.bookCategoryService.FindManyByCategory(c.Context(), &dto.FindManyByCategoryReq{
		Category: entity.Category(category),
		Page:     page,
	})

	if err != nil {
		return err
	}

	return c.Status(200).JSON(fiber.Map{"data": res.Data, "paging": res.Paging})
}

func (h *BookCategory) DeleteByCategory(c *fiber.Ctx) error {
	category := c.Params("category")

	if err := h.bookCategoryService.Delete(c.Context(), &dto.DeleteBookCategoryReq{Category: (*entity.Category)(&category)}); err != nil {
		return err
	}

	return c.Status(200).JSON(fiber.Map{"data": "successfully deleted the category by category"})
}

func (h *BookCategory) DeleteByBookId(c *fiber.Ctx) error {
	bookId, err := strconv.Atoi(c.Params("bookId"))
	if err != nil {
		return err
	}

	if err := h.bookCategoryService.Delete(c.Context(), &dto.DeleteBookCategoryReq{BookId: &bookId}); err != nil {
		return err
	}

	return c.Status(200).JSON(fiber.Map{"data": "successfully deleted the category by book id"})
}

func (h *BookCategory) DeleteFromBookCategory(c *fiber.Ctx) error {
	req := new(dto.DeleteBookCategoryReq)
	if err := c.BodyParser(req); err != nil {
		return err
	}

	if err := h.bookCategoryService.Delete(c.Context(), req); err != nil {
		return err
	}

	return c.Status(200).JSON(fiber.Map{"data": "successfully deleted book from the category"})
}
