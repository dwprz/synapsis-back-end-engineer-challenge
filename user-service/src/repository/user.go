package repository

import (
	"context"
	"errors"
	errcstm "user-service/src/common/errors"
	"user-service/src/common/helper"
	"user-service/src/interface/cache"
	"user-service/src/interface/repository"
	"user-service/src/model/dto"
	"user-service/src/model/entity"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserImpl struct {
	pool      *pgxpool.Pool
	userCache cache.User
}

func NewUser(p *pgxpool.Pool, uc cache.User) repository.User {
	return &UserImpl{
		pool:      p,
		userCache: uc,
	}
}

func (r *UserImpl) Create(ctx context.Context, data *dto.RegisterReq) error {
	conn, err := r.pool.Acquire(ctx)
	if err != nil {
		return err
	}

	defer conn.Release()

	query := `
	INSERT INTO 
		users(user_id, email, full_name, password) 
	VALUES 
		($1, $2, $3, $4) 
	RETURNING *;`

	row := conn.QueryRow(ctx, query, data.UserId, data.Email, data.FullName, data.Password)

	var user entity.User
	err = row.Scan(&user.UserId, &user.Email, &user.FullName, &user.Whatsapp, &user.Role, &user.Password,
		&user.RefreshToken, &user.CreatedAt, &user.UpdatedAt)

	if err == nil {
		go r.userCache.Cache(context.Background(), &user)
	}

	return err
}

func (r *UserImpl) FindByFields(ctx context.Context, data *entity.User) (*entity.User, error) {
	conn, err := r.pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}

	defer conn.Release()

	query, args, err := helper.BuildFindByFieldsQuery(data)
	if err != nil {
		return nil, err
	}

	row := conn.QueryRow(ctx, query, args...)

	var user entity.User
	err = row.Scan(&user.UserId, &user.Email, &user.FullName, &user.Whatsapp, &user.Role, &user.Password,
		&user.RefreshToken, &user.CreatedAt, &user.UpdatedAt)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}

	if err == nil {
		go r.userCache.Cache(context.Background(), &user)
	}

	return &user, err
}

func (r *UserImpl) UpdateByUserId(ctx context.Context, data *dto.UpdateUserReq) (*entity.User, error) {
	conn, err := r.pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}

	defer conn.Release()

	query, args, err := helper.BuildUpdateUserQuery(data)
	if err != nil {
		return nil, err
	}

	row := conn.QueryRow(ctx, query, args...)

	var user entity.User
	err = row.Scan(&user.UserId, &user.Email, &user.FullName, &user.Whatsapp, &user.Role, &user.Password,
		&user.RefreshToken, &user.CreatedAt, &user.UpdatedAt)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, &errcstm.Response{HttpCode: 404, Message: "user not found"}
	}

	if err == nil {
		go r.userCache.Cache(context.Background(), &user)
	}

	return &user, err
}

func (r *UserImpl) SetNullRefreshToken(ctx context.Context, refreshToken string) error {
	conn, err := r.pool.Acquire(ctx)
	if err != nil {
		return err
	}

	defer conn.Release()

	query := `
	UPDATE 
		users SET refresh_token = NULL, updated_at = now() 
	WHERE 
		refresh_token = $1 
	RETURNING 
		user_id, email, full_name, whatsapp, role, password, refresh_token, created_at, updated_at;
	`
	row := conn.QueryRow(ctx, query, refreshToken)

	var user entity.User
	err = row.Scan(&user.UserId, &user.Email, &user.FullName, &user.Whatsapp, &user.Role, &user.Password,
		&user.RefreshToken, &user.CreatedAt, &user.UpdatedAt)

	if errors.Is(err, pgx.ErrNoRows) {
		return &errcstm.Response{HttpCode: 404, Message: "user not found"}
	}

	if err == nil {
		go r.userCache.Cache(context.Background(), &user)
	}

	return err
}

func (r *UserImpl) Delete(ctx context.Context, userId string) error {
	conn, err := r.pool.Acquire(ctx)
	if err != nil {
		return err
	}

	defer conn.Release()

	query := `DELETE FROM users WHERE user_id = $1 RETURNING email;`

	var user entity.User
	err = conn.QueryRow(ctx, query, userId).Scan(&user.Email)

	if errors.Is(err, pgx.ErrNoRows) {
		return &errcstm.Response{HttpCode: 404, Message: "user not found"}
	}

	if err == nil {
		go r.userCache.DeleteByEmail(context.Background(), user.Email)
	}

	return err
}
