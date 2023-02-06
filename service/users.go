package service

import "net/http"

func (s *service) CreateNewUser(w http.ResponseWriter, r *http.Request)

func (s *service) CreateUsersFromExcell(w http.ResponseWriter, r *http.Request)

func (s *service) DeleteUser(w http.ResponseWriter, r *http.Request)

func (s *service) ChangeUser(w http.ResponseWriter, r *http.Request)

func (s *service) GetUserByID(w http.ResponseWriter, r *http.Request)

func (s *service) GetFilteredUsers(w http.ResponseWriter, r *http.Request)
