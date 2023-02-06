package model

import "time"

type User struct {
	Id         int64     `json:"id,omitempty"`
	Name       string    `json:"name"`
	Surname    string    `json:"surname"`
	Patronymic string    `json:"patronymic,omitempty"`
	Sex        string    `json:"sex"`
	Status     string    `json:"status"`
	BirthDate  time.Time `json:"birthdate,omitempty"`
	Created    time.Time `json:"created"`
}

type UserFilter struct {
	Sex      string `json:"sex,omitempty"`
	Status   string `json:"status,omitempty"`
	FullName string `json:"fullname,omitempty"`
	SortDesc bool   `json:"sort_desc,omitempty"`
	Limit    int64  `json:"limit,omitempty"`
	Offset   int64  `json:"offset,omitempty"`
}
