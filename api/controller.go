package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/nanoohlaing1997/online-ordering-items/env"
	"github.com/nanoohlaing1997/online-ordering-items/models"
)

var (
	environ  = env.GetEnviroment()
	validate = validator.New()
)

type Controller struct {
	dbm *models.DatabaseManger
}

type AuthResponse struct {
	UserID       uint64 `json:"user_id"`
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

func NewControllerManager(dbm *models.DatabaseManger) *Controller {
	return &Controller{
		dbm: dbm,
	}
}
