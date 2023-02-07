package service

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/Badchaos11/TSU_TestTask/model"
	jsoniter "github.com/json-iterator/go"
	"github.com/sirupsen/logrus"
)

func (s *service) CreateNewUser(w http.ResponseWriter, r *http.Request) {
	logrus.Info("Got CreateNewUser Request. Starting process.")
	body, err := io.ReadAll(r.Body)
	if err != nil {
		logrus.Errorf("Error reading request body error %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Не удалось создать подьзователя из-за внутренней ошибки."))
	}
	var req model.User
	err = jsoniter.Unmarshal(body, &req)
	if err != nil {
		logrus.Errorf("Error unmarshalling request body error %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Не удалось создать пользователя из-за внутренней ошибки"))
	}
	ctx := context.Background()
	id, err := s.Repo.CreateUser(ctx, req)
	if err != nil {
		logrus.Errorf("Error creqting user error %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Не удалось создать пользователя"))
	}

	logrus.Info("User succesfully created")
	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte(fmt.Sprintf("Пользователь успешно создан. ID пользователя %d", id)))
}

func (s *service) CreateUsersFromExcell(w http.ResponseWriter, r *http.Request)

func (s *service) DeleteUser(w http.ResponseWriter, r *http.Request)

func (s *service) ChangeUser(w http.ResponseWriter, r *http.Request)

func (s *service) GetUserByID(w http.ResponseWriter, r *http.Request)

func (s *service) GetFilteredUsers(w http.ResponseWriter, r *http.Request)
