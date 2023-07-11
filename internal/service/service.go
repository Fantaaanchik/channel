package service

import (
	"chans/models"
	"github.com/gin-gonic/gin"
)

type Services struct {
	Repository RepoInterface
}

func (s Services) GetCardFromDB(c *gin.Context) ([]models.Card, error) {
	return s.Repository.GetCardFromDB(c)
}

func NewService(repo RepoInterface) *Services {
	return &Services{Repository: repo}
}

type RepoInterface interface {
	GetCardFromDB(c *gin.Context) ([]models.Card, error)
}
