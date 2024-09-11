package service

import (
	"context"
	"user-service/src/common/errors"
	v "user-service/src/common/validator"
	"user-service/src/interface/cache"
	"user-service/src/interface/repository"
	"user-service/src/interface/service"
	"user-service/src/model/dto"
	"user-service/src/model/entity"

	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
)

type UserImpl struct {
	userRepo  repository.User
	userCache cache.User
}

func NewUser(ur repository.User, uc cache.User) service.User {
	return &UserImpl{
		userRepo:  ur,
		userCache: uc,
	}
}

func (s *UserImpl) FindByEmail(ctx context.Context, email string) (*entity.SanitizedUser, error) {
	if err := v.Validate.Var(email, `required,email,min=10,max=100`); err != nil {
		return nil, err
	}

	user := s.userCache.FindByEmail(ctx, email)
	if user == nil {
		res, err := s.userRepo.FindByFields(ctx, &entity.User{Email: email})
		if err != nil {
			return nil, err
		}

		user = res
	}
	userSanitized := new(entity.SanitizedUser)
	if err := copier.Copy(userSanitized, user); err != nil {
		return nil, err
	}

	return userSanitized, nil
}

func (s *UserImpl) Update(ctx context.Context, data *dto.UpdateUserReq) (*entity.SanitizedUser, error) {
	if err := v.Validate.Struct(data); err != nil {
		return nil, err
	}

	if data.Email != "" {
		user := s.userCache.FindByEmail(ctx, data.Email)
		if user == nil {
			res, err := s.userRepo.FindByFields(ctx, &entity.User{Email: data.Email})
			if err != nil {
				return nil, err
			}

			user = res
		}

		if user != nil {
			return nil, &errors.Response{HttpCode: 409, Message: "email already exists"}
		}
	}

	if data.Password != "" {
		encryptPwd, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
	
		data.Password = string(encryptPwd)
	}

	user, err := s.userRepo.UpdateByUserId(ctx, data)
	if err != nil {
		return nil, err
	}

	userSanitized := new(entity.SanitizedUser)
	if err := copier.Copy(userSanitized, user); err != nil {
		return nil, err
	}

	return userSanitized, nil
}

func (s *UserImpl) Delete(ctx context.Context, userId string) error {
	if err := v.Validate.Var(userId, `required,min=21,max=21`); err != nil {
		return err
	}

	err := s.userRepo.Delete(ctx, userId)	
	return err
}
