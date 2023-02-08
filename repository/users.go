package repository

import (
	"context"

	"github.com/Badchaos11/TSU_TestTask/model"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
)

func (r *Repo) CreateUser(ctx context.Context, u model.User) (int64, error) {
	const query = `INSERT INTO users (name, surname, patronymic, sex, status, birth, created) VALUES ($1, $2, $3, $4, $5, $6, now()) RETURNING id;`
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()
	var id int64

	err := r.PGXRepo.QueryRow(ctx, query, u.Name, u.Surname, u.Patronymic, u.Sex, u.Status, u.BirthDate).Scan(&id)
	if err != nil {
		logrus.Errorf("Error creating user: %v", err)
		return 0, err
	}

	return id, nil
}

func (r *Repo) ChangeUser(ctx context.Context, u model.User) (bool, error) {
	ub := squirrel.Update("users").PlaceholderFormat(squirrel.Dollar).Where("id", u.Id)
	fvMap := makeFieldValMap(u)
	for k, v := range fvMap {
		if v != "" {
			ub = ub.Set(k, v)
		}
	}
	sql, args, _ := ub.ToSql()
	_, err := r.PGXRepo.Exec(ctx, sql, args...)
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

func (r *Repo) DeleteUser(ctx context.Context, userId int64) (bool, error) {
	const query = `DELETE FROM users WHERE id=$1`
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	_, err := r.PGXRepo.Exec(ctx, query, userId)
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

func (r *Repo) GetUserByID(ctx context.Context, userId int) (*model.User, error) {
	const query = `SELECT * FROM users WHERE id=$1`
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()
	var u model.User

	err := r.PGXRepo.QueryRow(ctx, query, userId).Scan(&u.Id, &u.Name, &u.Surname, &u.Patronymic, &u.Sex, &u.Status, &u.BirthDate, &u.Created)
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

func (r *Repo) GetUsersFiltered(ctx context.Context, filter model.UserFilter) ([]model.User, error) {
	sq := squirrel.Select("*").From("users").PlaceholderFormat(squirrel.Dollar)
	if filter.Limit > 0 {
		sq = sq.Limit(filter.Limit)
	}
	if filter.Offset > 0 {
		sq = sq.Offset(filter.Offset)
	}

	if filter.Sex != "" {
		sq = sq.Where("sex", filter.Sex)
	}
	if filter.Status != "" {
		sq = sq.Where("status", filter.Status)
	}

	if filter.OrderBy == "sex" || filter.OrderBy == "status" {
		if filter.Desc != nil && *filter.Desc {
			sq = sq.OrderBy(filter.OrderBy, "desc")
		}
		sq = sq.OrderBy(filter.OrderBy)
	}

	if filter.ByName != nil && *filter.ByName {

	}

	sql, args, _ := sq.ToSql()
	var res []model.User
	rows, err := r.PGXRepo.Query(ctx, sql, args...)
	if err != nil {
		if err == pgx.ErrNoRows {
			logrus.Error("No userf for this filter")
			return nil, nil
		}
		logrus.Errorf("error select users error %v", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var row model.User
		err := rows.Scan(&row.Id, &row.Name, &row.Surname, &row.Patronymic, &row.Sex, &row.Status, &row.BirthDate, &row.Created)
		if err != nil {
			logrus.Errorf("error scannig row %v", err)
			return nil, err
		}
		res = append(res, row)
	}

	return res, nil
}
