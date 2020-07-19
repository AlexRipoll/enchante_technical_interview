package product

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
	Search(c *gin.Context)
	Add(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	SearchAll(c *gin.Context)
}

func NewHandler(service Service) Handler {
	return &handler{service}
}

func (h *handler) Search(c *gin.Context) {
	h.checkAccessRights(c)
	id := strings.TrimSpace(c.Param("product_id"))
	if err := uuidv4.NewService().Validate(id); err != nil {
		apiErr := errors.NewBadRequestError(err.Error())
		c.JSON(apiErr.Status, apiErr)
		return
	}

	p, serviceErr := h.service.Find(id)
	if serviceErr != nil {
		c.JSON(serviceErr.Status, serviceErr)
		return
	}
	c.JSON(http.StatusOK, p)
}

func (h *handler) Add(c *gin.Context) {
	h.checkAccessRights(c)
	sellerId := strings.TrimSpace(c.Param("id"))
	var p Product
	if err := c.ShouldBindJSON(&p); err != nil {
		apiErr := errors.NewBadRequestError("invalid json body")
		c.JSON(apiErr.Status, apiErr)
		return
	}
	if err := h.service.Add(p.Name, p.Price, sellerId); err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusCreated, map[string]string{"status": "created"})
}

func (h *handler) Update(c *gin.Context) {
	h.checkAccessRights(c)
	id := strings.TrimSpace(c.Param("product_id"))
	if err := uuidv4.NewService().Validate(id); err != nil {
		apiErr := errors.NewBadRequestError(err.Error())
		c.JSON(apiErr.Status, apiErr)
		return
	}
	sellerId := strings.TrimSpace(c.Param("id"))
	if err := uuidv4.NewService().Validate(id); err != nil {
		apiErr := errors.NewBadRequestError(err.Error())
		c.JSON(apiErr.Status, apiErr)
		return
	}

	var p Product
	if err := c.ShouldBindJSON(&p); err != nil {
		apiErr := errors.NewBadRequestError("invalid json body")
		c.JSON(apiErr.Status, apiErr)
		return
	}

	if err := h.service.Update(id, p.Name, p.Price, sellerId); err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"status": "updated"})
}

func (h *handler) Delete(c *gin.Context) {
	h.checkAccessRights(c)
	id := strings.TrimSpace(c.Param("product_id"))
	if err := uuidv4.NewService().Validate(id); err != nil {
		apiErr := errors.NewBadRequestError(err.Error())
		c.JSON(apiErr.Status, apiErr)
		return
	}

	if err := h.service.Delete(id); err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

func (h *handler) SearchAll(c *gin.Context) {
	h.checkAccessRights(c)
	products, err := h.service.FindAll()
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, products)
}

func (h *handler) checkAccessRights(c *gin.Context) {
	id, ok := c.Get("id")
	if !ok {
		apiErr := errors.NewUnauthorizedError("missing claim")
		c.JSON(apiErr.Status, apiErr)
		c.Abort()
		return
	}
	sellerId := strings.TrimSpace(c.Param("id"))
	if id != sellerId {
		apiErr := errors.NewForbiddenAccessError("forbidden access")
		c.JSON(apiErr.Status, apiErr)
		c.Abort()
		return
	}

	role, ok := c.Get("role")
	if !ok {
		apiErr := errors.NewUnauthorizedError("missing role claim")
		c.JSON(apiErr.Status, apiErr)
		c.Abort()
		return
	}
	if role != "seller" {
		apiErr := errors.NewForbiddenAccessError("forbidden access")
		c.JSON(apiErr.Status, apiErr)
		c.Abort()
		return
	}
}

