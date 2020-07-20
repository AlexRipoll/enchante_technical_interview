package cart

import (
	"github.com/AlexRipoll/enchante_technical_interview/pkg/errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type handler struct {
	service Service
}

type Handler interface {
	Purchase(c *gin.Context)
}

func NewHandler(service Service) Handler {
	return &handler{service}
}

func (h *handler) Purchase(c *gin.Context) {
	userId := strings.TrimSpace(c.Param("id"))
	var order Order
	if err := c.ShouldBindJSON(&order); err != nil {
		apiErr := errors.NewBadRequestError("invalid json body")
		c.JSON(apiErr.Status, apiErr)
		return
	}

	if err := h.service.Purchase(userId, order.Items); err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusCreated, map[string]string{"status": "created"})
}