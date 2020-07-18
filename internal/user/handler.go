package user

import (
	"github.com/AlexRipoll/enchante_technical_interview/pkg/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type handler struct {
	service Service
}

type Handler interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
	Search(c *gin.Context)
	Delete(c *gin.Context)
	FindAll(c *gin.Context)
}

func NewHandler(service Service) Handler {
	return &handler{service}
}

func (h *handler) Register(c *gin.Context) {
	var account Account
	if err := c.ShouldBindJSON(&account); err != nil {
		apiErr := errors.NewBadRequestError("invalid json body")
		c.JSON(apiErr.Status, apiErr)
		return
	}

	if serviceErr := h.service.Register(account.Username, account.Email, account.Password); serviceErr != nil {
		c.JSON(serviceErr.Status, serviceErr)
		return
	}
	c.JSON(http.StatusCreated, map[string]string{"status": "created"})
}

func (h *handler) Login(c *gin.Context) {
	// TODO implementation
}

func (h *handler) Search(c *gin.Context) {
	// TODO implementation
}

func (h *handler) Delete(c *gin.Context) {
	// TODO implementation
}

func (h *handler) FindAll(c *gin.Context) {
	// TODO implementation
}
