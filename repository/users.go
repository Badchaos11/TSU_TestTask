package repository

import (
	"context"

	"github.com/Badchaos11/TSU_TestTask/model"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
)

func (r *PGXRepo) CreateUser(ctx context.Context, u model.User) (int64, error) {
	const query = `INSERT INTO users (name, surname, patronymic, sex, status, birth, created) VALUES ($1, $2, $3, $4, $5, $6, now()) RETURNING id;`
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()
	var id int64

	err := r.Conn.QueryRow(ctx, query, u.Name, u.Surname, u.Patronymic, u.Sex, u.Status, u.BirthDate).Scan(&id)
	if err != nil {
		logrus.Errorf("Error creating user: %v", err)
		return 0, err
	}

	return id, nil
}

func (r *PGXRepo) ChangeUser(ctx context.Context, u model.User) (bool, error) {
	ub := squirrel.Update("users").PlaceholderFormat(squirrel.Dollar).Where("id", u.Id)
	fvMap := makeFieldValMap(u)
	for k, v := range fvMap {
		if v != "" {
			ub = ub.Set(k, v)
		}
	}
	sql, args, _ := ub.ToSql()
	_, err := r.Conn.Exec(ctx, sql, args...)
	if err != nil {
		if err == pgx.ErrNoRows {
			logrus.Error("There is no user with this user ID %d", u.Id)
			return false, nil
		}
		logrus.Errorf("Error updating user %v", err)
		return false, err
	}

	return true, nil
}

func (r *PGXRepo) DeleteUser(ctx context.Context, userId int64) (bool, error) {
	const query = `DELETE FROM users WHERE id=$1`
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	_, err := r.Conn.Exec(ctx, query, userId)
	if err != nil {
		if err == pgx.ErrNoRows {
			logrus.Errorf("There is no user with id %d", userId)
			return false, nil
		}
		logrus.Errorf("Error deleting user: %v", err)
		return false, err
	}

	return true, nil
}

func (r *PGXRepo) GetUserByID(ctx context.Context, userId int64) (*model.User, error) {
	const query = `SELECT * FROM users WHERE id=$1`
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()
	var u model.User

	err := r.Conn.QueryRow(ctx, query, userId).Scan(&u.Id, &u.Name, &u.Surname, &u.Patronymic, &u.Sex, &u.Status, &u.BirthDate, &u.Created)
	if err != nil {
		if err == pgx.ErrNoRows {
			logrus.Errorf("There is no user with user ID %d", userId)
			return nil, nil
		}
		logrus.Errorf("Error getting user %v", err)
		return nil, err
	}

	return &u, nil
}

func (r *PGXRepo) GetUserFilter(ctx context.Context, filter model.UserFilter) ([]model.User, error)
