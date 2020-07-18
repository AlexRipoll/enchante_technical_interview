package user

import (
	"github.com/AlexRipoll/enchante_technical_interview/pkg/errors"
	"github.com/AlexRipoll/enchante_technical_interview/pkg/uuidv4"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type handler struct {
	service Service
}

type Handler interface {
	Register(c *gin.Context)
	RegisterUser(c *gin.Context)
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
	var account Account
	if err := c.ShouldBindJSON(&account); err != nil {
		apiErr := errors.NewBadRequestError("invalid json body")
		c.JSON(apiErr.Status, apiErr)
		return
	}

	jwt, err := h.service.Login(account.Email, account.Password)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusCreated, map[string]string{"token": jwt})
}

func (h *handler) RegisterUser(c *gin.Context) {
	h.checkAccessRights(c)
	var a Account
	if err := c.ShouldBindJSON(&a); err != nil {
		apiErr := errors.NewBadRequestError("invalid json body")
		c.JSON(apiErr.Status, apiErr.Message)
		return
	}
	if serviceErr := h.service.RegisterUser(a.Username, a.Email, a.Password, a.Role); serviceErr != nil {
		c.JSON(serviceErr.Status, serviceErr)
		return
	}
	c.JSON(http.StatusCreated, map[string]string{"status": "created"})
}

func (h *handler) Search(c *gin.Context) {
	id := strings.TrimSpace(c.Param("id"))
	if err := uuidv4.NewService().Validate(id); err != nil {
		apiErr := errors.NewBadRequestError(err.Error())
		c.JSON(apiErr.Status, apiErr)
		return
	}

	u, serviceErr := h.service.Search(id)
	if serviceErr != nil {
		c.JSON(serviceErr.Status, serviceErr)
		return
	}
	c.JSON(http.StatusOK, u)
}

func (h *handler) Delete(c *gin.Context) {
	id := strings.TrimSpace(c.Param("id"))
	if err := uuidv4.NewService().Validate(id); err != nil {
		apiErr := errors.NewBadRequestError(err.Error())
		c.JSON(apiErr.Status, apiErr)
		return
	}

	h.checkAccessRights(c)
	if serviceErr := h.service.Delete(id); serviceErr != nil {
		c.JSON(serviceErr.Status, serviceErr)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

func (h *handler) FindAll(c *gin.Context) {
	// TODO implementation
}


func (h *handler) checkAccessRights(c *gin.Context) {
	id, ok := c.Get("id")
	if !ok {
		apiErr := errors.NewUnauthorizedError("missing claim")
		c.JSON(apiErr.Status, apiErr)
		return
	}
	account, err := h.service.Search(id.(string))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	if account.Role != "admin" {
		apiErr := errors.NewForbiddenAccessError("forbidden access")
		c.JSON(apiErr.Status, apiErr)
		return
	}
}