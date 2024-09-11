package service

import (
	"context"
	"user-service/src/common/errors"
	"user-service/src/common/helper"
	"user-service/src/common/log"
	v "user-service/src/common/validator"
	"user-service/src/interface/cache"
	"user-service/src/interface/repository"
	"user-service/src/interface/service"
	"user-service/src/model/dto"
	"user-service/src/model/entity"

	"github.com/jinzhu/copier"
	"github.com/matoous/go-nanoid/v2"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type AuthImpl struct {
	userRepo  repository.User
	userCache cache.User
}

func NewAuth(ur repository.User, uc cache.User) service.Auth {
	return &AuthImpl{
		userRepo:  ur,
		userCache: uc,
	}
}

func (s *AuthImpl) Register(ctx context.Context, data *dto.RegisterReq) error {
	if err := v.Validate.Struct(data); err != nil {
		return err
	}

	user := s.userCache.FindByEmail(ctx, data.Email)
	if user == nil {
		res, err := s.userRepo.FindByFields(ctx, &entity.User{Email: data.Email})
		if err != nil {
			return err
		}

		user = res
	}

	if user != nil {
		return &errors.Response{HttpCode: 409, Message: "user already exists"}
	}

	userId, err := gonanoid.New()
	if err != nil {
		return err
	}

	data.UserId = userId

	encryptPwd, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	data.Password = string(encryptPwd)

	err = s.userRepo.Create(ctx, data)
	return err
}

func (s *AuthImpl) Login(ctx context.Context, data *dto.LoginReq) (*dto.LoginRes, error) {
	if err := v.Validate.Struct(data); err != nil {
		return nil, err
	}

	user := s.userCache.FindByEmail(ctx, data.Email)
	if user == nil {
		res, err := s.userRepo.FindByFields(ctx, &entity.User{Email: data.Email})
		if err != nil {
			return nil, err
		}

		user = res
	}

	if user == nil {
		return nil, &errors.Response{HttpCode: 404, Message: "there are no users that match this email"}
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password)); err != nil {
		return nil, &errors.Response{HttpCode: 401, Message: "password is invalid"}
	}

	accessToken, err := helper.GenerateAccessToken(user.UserId, user.Email, user.Role)
	if err != nil {
		return nil, err
	}

	refreshToken, err := helper.GenerateRefreshToken()
	if err != nil {
		return nil, err
	}

	go func() {
		if _, err := s.userRepo.UpdateByUserId(ctx, &dto.UpdateUserReq{UserId: user.UserId, RefreshToken: refreshToken}); err != nil {
			log.Logger.WithFields(logrus.Fields{"location": "service.AuthImpl/Login", "section": "userRepo.UpdateByUserId"}).Error(err)
		}
	}()

	userSanitized := new(entity.SanitizedUser)
	if err := copier.Copy(userSanitized, user); err != nil {
		return nil, err
	}

	return &dto.LoginRes{
		Data: userSanitized,
		Tokens: &entity.Tokens{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	}, nil
}

func (s *AuthImpl) RefreshToken(ctx context.Context, refreshToken string) (*entity.Tokens, error) {
	if err := v.Validate.Struct(&entity.RefreshToken{Token: refreshToken}); err != nil {
		return nil, err
	}

	user, err := s.userRepo.FindByFields(ctx, &entity.User{RefreshToken: &refreshToken})

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, &errors.Response{HttpCode: 404, Message: "there are no users that match this refresh token"}
	}

	accessToken, err := helper.GenerateAccessToken(user.UserId, user.Email, user.Role)
	if err != nil {
		return nil, err
	}

	return &entity.Tokens{
		AccessToken: accessToken,
	}, nil
}

func (h *AuthImpl) SetNullRefreshToken(ctx context.Context, refreshToken string) error {
	if err := v.Validate.Struct(&entity.RefreshToken{Token: refreshToken}); err != nil {
		return err
	}

	err := h.userRepo.SetNullRefreshToken(ctx, refreshToken)
	return err
}
