package service

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

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
		return
	}
	var req model.User
	err = jsoniter.Unmarshal(body, &req)
	if err != nil {
		logrus.Errorf("Error unmarshalling request body error %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Не удалось создать пользователя из-за внутренней ошибки"))
		return
	}
	ctx := context.Background()
	id, err := s.repo.CreateUser(ctx, req)
	if err != nil {
		logrus.Errorf("Error creqting user error %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Не удалось создать пользователя"))
		return
	}

	logrus.Info("User succesfully created")
	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte(fmt.Sprintf("Пользователь успешно создан. ID пользователя %d", id)))
}

func (s *service) CreateUsersFromExcell(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(0)
	if err != nil {
		logrus.Errorf("Error parsing multipart form: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Не удалось прочитать файл"))
		return
	}
	h := r.MultipartForm.File["user"][0]
	fName := h.Filename
	f, err := h.Open()
	if err != nil {
		logrus.Errorf("error opening multipart form: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Не удалось открыть файл"))
		return
	}
	tmpfile, _ := os.Create("./userfiles/" + h.Filename)
	io.Copy(tmpfile, f)
	f.Close()

	user, err := s.GetUserFromFile(fName)
	if err != nil {
		logrus.Errorf("Error getting user frmo file %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Не удалось прочитать файл"))
		return
	}
	ctx := context.Background()
	id, err := s.repo.CreateUser(ctx, *user)
	if err != nil {
		logrus.Errorf("Error creating user %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Не удалось создать пользователя"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Пользователь успешно создан, id %d", id)))

}

func (s *service) DeleteUser(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Не удалось прочитать тело запроса"))
		return
	}
	var req model.DeleteUserRequest
	if err := jsoniter.Unmarshal(body, &req); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Не удалось сериалтзовать тело"))
		return
	}
	ctx := context.Background()
	succes, err := s.repo.DeleteUser(ctx, req.UserID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Не удалось удалить пользователя"))
		return
	}
	if !succes {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Пользователя с таким id не существует"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Пользователь успешно удален"))
}

func (s *service) ChangeUser(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Не удалось прочитать тело запроса"))
		return
	}
	var req model.ChangeUserRequest

	if err := jsoniter.Unmarshal(body, &req); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Не удалось сериализовать тело запроса"))
		return
	}
	ctx := context.Background()
	exists, err := s.repo.ChangeUser(ctx, model.User{Id: req.Id, Name: req.Name, Surname: req.Surname,
		Patronymic: req.Patronymic, Sex: req.Sex, Status: req.Status})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Не удалось обновить данные пользователя"))
		return
	}
	if !exists {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Пользователя с таким id не существует"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Данные пользователя успешно изменены"))
}

func (s *service) GetUserByID(w http.ResponseWriter, r *http.Request) {
	userIdStr := r.URL.Query().Get("user_id")
	if userIdStr == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Введен пустой id пользователя"))
		return
	}
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Введен некорректный id пользователя, можно использовать только цифры"))
		return
	}
	ctx := context.Background()
	user, err := s.repo.GetUserByID(ctx, userId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Не удалось найти пользователя"))
		return
	}
	if user == nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Пользователь с таким id не существует"))
		return
	}

	strBody, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Ошибка сериализации пользователя"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(strBody))
}

func (s *service) GetFilteredUsers(w http.ResponseWriter, r *http.Request) {
	filter := s.GetUserFilter(r)

	ctx := context.Background()
	users, err := s.repo.GetUsersFiltered(ctx, filter)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Ошибка поиска пользователей по заданному фильтру"))
		return
	}

	if users == nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Для данного фильтра пользователи не найдены"))
		return
	}

	strBody, err := json.Marshal(users)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Ошибка сериализации резальтата"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(strBody))
}
