package handler

import (
	"user-service/src/interface/service"
	"user-service/src/model/dto"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type User struct {
	userService service.User
}

func NewUser(us service.User) *User {
	return &User{
		userService: us,
	}
}

func (h *User) GetByCurrent(c *fiber.Ctx) error {
	userData := c.Locals("user_data").(jwt.MapClaims)
	email := userData["email"].(string)

	res, err := h.userService.FindByEmail(c.Context(), email)
	if err != nil {
		return err
	}

	return c.Status(200).JSON(fiber.Map{"data": res})
}

func (h *User) Update(c *fiber.Ctx) error {
	userData := c.Locals("user_data").(jwt.MapClaims)
	userId := userData["user_id"].(string)

	req := new(dto.UpdateUserReq)
	if err := c.BodyParser(req); err != nil {
		return err
	}

	req.UserId = userId

	res, err := h.userService.Update(c.Context(), req)
	if err != nil {
		return err
	}

	return c.Status(200).JSON(fiber.Map{"data": res})
}

func (h *User) Delete(c *fiber.Ctx) error {
	userData := c.Locals("user_data").(jwt.MapClaims)
	userId := userData["user_id"].(string)

	if err := h.userService.Delete(c.Context(), userId); err != nil {
		return err
	}

	return c.Status(200).JSON(fiber.Map{"data": "successfully deleted the user"})
}
