package handler

import (
	"chans/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	Engine   *gin.Engine
	services ServInterface
}

type ServInterface interface {
	GetCardFromDB(c *gin.Context) ([]models.Card, error)
}

func NewH(services ServInterface, engine *gin.Engine) *Handler {
	return &Handler{services: services, Engine: engine}
}

func (h Handler) AllRoutes() {

	h.Engine.GET("/get_cards", h.GetCards)

}

func (h Handler) GetCards(c *gin.Context) {

	cards, err := h.services.GetCardFromDB(c)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{"cards": cards})
}
