package service

import "gin-rest-api/model"

type Auth struct {
	Username string
	Password string
}

func (a *Auth) Check() (bool, error) {
	return model.CheckAuth(a.Username, a.Password)
}
