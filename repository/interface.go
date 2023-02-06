package repository

import (
	"context"
	"time"

	"github.com/Badchaos11/TSU_TestTask/model"
	"github.com/fatih/structs"
	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
)

type IRepository interface {
	CreateUser(ctx context.Context, u model.User) (int64, error)
	ChangeUser(ctx context.Context, u model.User) (bool, error)
	DeleteUser(ctx context.Context, userId int64) (bool, error)
	GetUserByID(ctx context.Context, userId int64) (*model.User, error)
	GetUserFilter(ctx context.Context, filter model.UserFilter) ([]model.User, error)
}

type PGXRepo struct {
	Conn    *pgx.Conn
	timeout time.Duration
}

func NewRepository(ctx context.Context, dsn string) (IRepository, error) {
	conn, err := pgx.Connect(ctx, dsn)
	if err != nil {
		logrus.Errorf("Error connecting to database: %v", err)
		return nil, err
	}

	return &PGXRepo{
		Conn:    conn,
		timeout: time.Second * 5,
	}, nil
}

func makeFieldValMap(u model.User) map[string]string {
	fields := structs.Fields(u)
	res := make(map[string]string, 6)

	for _, field := range fields {
		f := field.Tag("json")
		res[f] = field.Value().(string)
	}

	return res
}
