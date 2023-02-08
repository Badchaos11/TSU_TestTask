package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Badchaos11/TSU_TestTask/model"
	"github.com/sirupsen/logrus"
	"github.com/xuri/excelize/v2"
)

func (s *service) WriteResponse(w http.ResponseWriter, code int, msg string) {
	w.WriteHeader(code)
	type response struct {
		Message string `json:"message"`
	}
	res := response{Message: msg}
	body, _ := json.Marshal(&res)
	w.Write(body)
}

func (s *service) GetDataFromCache(key string) ([]model.User, error) {
	// Типа получили данные из кеша
	// И обработали ошибку
	var data string
	var out []model.User

	if err := json.Unmarshal([]byte(data), &out); err != nil {
		logrus.Errorf("error unmarshalling data from cache: %v", err)
		return nil, err
	}

	return out, nil
}

func (s *service) GetUserFilter(r *http.Request) model.UserFilter {
	var limit, offset int64
	var byName, desc *bool

	sex := r.URL.Query().Get("sex")
	status := r.URL.Query().Get("status")
	fullName := r.URL.Query().Get("full_name")
	orderBy := r.URL.Query().Get("order_by")

	byNameStr := r.URL.Query().Get("by_name")
	descStr := r.URL.Query().Get("desc")
	limitStr := r.URL.Query().Get("limit")
	offsetStr := r.URL.Query().Get("offset")

	limit, err := strconv.ParseInt(limitStr, 10, 64)
	if err != nil {
		limit = 0
	}
	offset, err = strconv.ParseInt(offsetStr, 10, 64)
	if err != nil {
		offset = 0
	}

	*byName, err = strconv.ParseBool(byNameStr)
	if err != nil {
		byName = nil
	}

	*desc, err = strconv.ParseBool(descStr)
	if err != nil {
		desc = nil
	}

	return model.UserFilter{
		Sex:      sex,
		Status:   status,
		FullName: fullName,
		OrderBy:  orderBy,
		Limit:    uint64(limit),
		Offset:   uint64(offset),
		Desc:     desc,
		ByName:   byName,
	}
}

func (s *service) GetUserFromFile(fileName string) (*model.User, error) {
	f, err := excelize.OpenFile(fmt.Sprintf("./userfiles/%s.xlsx", fileName))
	if err != nil {
		return nil, err
	}

	cellName, _ := f.GetCellValue("Sheet1", "A1")
	cellSurname, _ := f.GetCellValue("Sheet1", "B1")
	cellPatr, _ := f.GetCellValue("Sheet1", "C1")
	cellSex, _ := f.GetCellValue("Sheet1", "D1")
	// cellMail, _ := f.GetCellValue("Sheet1", "E1")

	return &model.User{
		Name:       cellName,
		Surname:    cellSurname,
		Patronymic: cellPatr,
		Sex:        cellSex,
	}, nil
}
