package util

import (
	"context"
	"user-service/src/common/log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

type UserTest struct {
	pool *pgxpool.Pool
}

func NewUserTest(p *pgxpool.Pool) *UserTest {
	return &UserTest{
		pool: p,
	}
}

func (u *UserTest) Delete() {
	ctx := context.Background()

	conn, err := u.pool.Acquire(ctx)
	if err != nil {
		log.Logger.WithFields(logrus.Fields{"location": "util.UserTest/Delete", "section": "pool.Acquire"}).Error(err)
		return
	}

	defer conn.Release()

	if _, err = conn.Exec(ctx, `DELETE FROM users;`); err != nil {
		log.Logger.WithFields(logrus.Fields{"location": "util.UserTest/Delete", "section": "conn.Exec"})
	}
}

func (u *UserTest) Create() {
	ctx := context.Background()

	conn, err := u.pool.Acquire(ctx)
	if err != nil {
		log.Logger.WithFields(logrus.Fields{"location": "util.UserTest/Create", "section": "pool.Acquire"}).Error(err)
		return
	}

	defer conn.Release()

	query := `
	INSERT INTO
		users(user_id, email, full_name, role, password, refresh_token)
	VALUES
		('ynA1nZIULkXLrfy0fvz5t', 'johndoe123@gmail.com', 'John Doe', 'ADMIN', 'rahasia', 'example-refresh-token');
	`

	if _, err = conn.Exec(ctx, query); err != nil {
		log.Logger.WithFields(logrus.Fields{"location": "util.UserTest/CreateMany", "section": "conn.Exec"})
	}
}
