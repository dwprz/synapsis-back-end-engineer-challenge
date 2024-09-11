package handler

import (
	"user-service/src/common/errors"
	"user-service/src/common/helper"
	"user-service/src/interface/service"
	"user-service/src/model/dto"
	"time"

	"github.com/gofiber/fiber/v2"
)

type Auth struct {
	authService service.Auth
}

func NewAuth(as service.Auth) *Auth {
	return &Auth{
		authService: as,
	}
}

func (h *Auth) Register(c *fiber.Ctx) error {
	request := new(dto.RegisterReq)

	if err := c.BodyParser(request); err != nil {
		return err
	}

	err := h.authService.Register(c.Context(), request)
	if err != nil {
		return err
	}

	return c.Status(201).JSON(fiber.Map{"data": "successfully register"})
}

func (h *Auth) Login(c *fiber.Ctx) error {
	req := new(dto.LoginReq)
	if err := c.BodyParser(req); err != nil {
		return err
	}

	res, err := h.authService.Login(c.Context(), req)
	if err != nil {
		return err
	}

	c.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    res.Tokens.AccessToken,
		Path:     "/",
		HTTPOnly: true,
		Expires:  time.Now().Add(1 * time.Hour),
	})

	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    res.Tokens.RefreshToken,
		Path:     "/",
		HTTPOnly: true,
		Expires:  time.Now().Add(30 * 24 * time.Hour),
	})

	return c.Status(200).JSON(fiber.Map{"data": res.Data})
}

func (h *Auth) RefreshToken(c *fiber.Ctx) error {
	refreshToken := c.Cookies("refresh_token")

	if _, err := helper.VerifyJwt(refreshToken); err != nil {
		return &errors.Response{HttpCode: 401, Message: "refresh token is required"}
	}

	res, err := h.authService.RefreshToken(c.Context(), refreshToken)
	if err != nil {
		return err
	}

	c.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    res.AccessToken,
		Path:     "/",
		HTTPOnly: true,
		Expires:  time.Now().Add(1 * time.Hour),
	})

	return c.Status(201).JSON(fiber.Map{"data": "successfully refreshed the token"})
}

func (h *Auth) Logout(c *fiber.Ctx) error {
	refreshToken := c.Cookies("refresh_token")

	err := h.authService.SetNullRefreshToken(c.Context(), refreshToken)
	if err != nil {
		return err
	}

	// clear cookie
	c.Cookie(helper.ClearCookie("refresh_token", "/"))
	c.Cookie(helper.ClearCookie("access_token", "/"))

	return c.Status(200).JSON(fiber.Map{"data": "successfully logged out"})
}
