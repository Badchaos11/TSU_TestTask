package repository

import (
	"context"
	"time"

	"github.com/Badchaos11/TSU_TestTask/model"
	"github.com/Badchaos11/TSU_TestTask/repository/cache"
	"github.com/fatih/structs"
	"github.com/jackc/pgx/v5"
)

type IRepository interface {
	CreateUser(ctx context.Context, u model.User) (int64, error)
	ChangeUser(ctx context.Context, u model.User) (bool, error)
	DeleteUser(ctx context.Context, userId int64) (bool, error)
	GetUserByID(ctx context.Context, userId int) (*model.User, error)
	GetUsersFiltered(ctx context.Context, filter model.UserFilter) ([]model.User, error)
	ClearCache(ctx context.Context) error
}

type Repo struct {
	PGXRepo *pgx.Conn
	KVRepo  cache.ICacheRepository
	timeout time.Duration
}

func NewRepository(ctx context.Context, dsn string, cacheUrl string) (IRepository, error) {
	conn, err := pgx.Connect(ctx, dsn)
	if err != nil {
		return nil, err
	}

	cacheConn, err := cache.NewCacheClient(ctx, cacheUrl)
	if err != nil {
		return nil, err
	}

	return &Repo{
		PGXRepo: conn,
		KVRepo:  cacheConn,
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

func (r *Repo) ClearCache(ctx context.Context) error {
	return r.KVRepo.ClearCache(ctx)
}
