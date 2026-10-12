package handler

import (
	"net/http"
	"payflow/internal/model"
	"payflow/internal/service"

	"github.com/gin-gonic/gin"
)

type MerchantUserHandler struct {
	service *service.MerchantUserService
}

func NewMerchantUserHandler(merchantUserService *service.MerchantUserService) *MerchantUserHandler {
	return &MerchantUserHandler{
		service: merchantUserService,
	}
}

func (h *MerchantUserHandler) Register(c *gin.Context) {
	var req model.RegisterMerchantRequest

	err := c.ShouldBindJSON(&req)

	merchantUser, err := h.service.Register(req)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": merchantUser,
	})

}
