package handler

import (
	"errors"
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

func(h *MerchantUserHandler) Login(c *gin.Context) {

	var req model.LoginRequest

	err := c.ShouldBindJSON(&req)
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	merchantUser, err := h.service.Login(req.Email, req.Password)

	if err != nil {
		if errors.Is(err, service.ErrInvalidCredentials) {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid email or password",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H {
		"data": merchantUser,
	})

}