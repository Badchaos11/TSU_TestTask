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
	ByName   bool   `json:"by_name,omitempty"`
	FullName string `json:"fullname,omitempty"`
	OrderBy  string `json:"order_by,omitempty"`
	Desc     bool   `json:"sort_desc,omitempty"`
	Limit    uint64 `json:"limit,omitempty"`
	Offset   uint64 `json:"offset,omitempty"`
}

type ResponseMessage struct {
	Message string `json:"message"`
}