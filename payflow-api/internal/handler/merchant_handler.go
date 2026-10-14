package handler

import (
	"errors"
	"net/http"
	"uuid"

	"payflow/internal/model"
	"payflow/internal/repository"
	"payflow/internal/service"

	"github.com/gin-gonic/gin"
)

type MerchantHandler struct {
	service *service.MerchantService
}

func NewMerchantHandler(merchantService *service.MerchantService) *MerchantHandler {
	return &MerchantHandler{
		service: merchantService,
	}
}

func (h *MerchantHandler) GetAll(c *gin.Context) {
	merchants, err := h.service.GetAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": merchants,
	})
}

func (h *MerchantHandler) Create(c *gin.Context) {

	var req model.CreateMerchantRequest

	err := c.ShouldBindJSON(&req)

	merchant, err := h.service.Create(req)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": merchant,
	})

}

func (h *MerchantHandler) GetByID(c *gin.Context) {
	getId := c.Param("id")

	id, err := uuid.Parse(getId)
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid UUID format",
		})
		return
	}

	merchantIDValue, exists := c.Get("merchant_id")
	if !exists {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Authentication required",
		})
		return
	}

	merchantIDStr, ok := merchantIDValue.(string)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid authentication context",
		})
		return
	}

	merchantID, err := uuid.Parse(merchantIDStr)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid merchant identity",
		})
		return
	}

	if id != merchantID {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "You cannot access another merchant",
		})
		return
	}

	merchant, err := h.service.GetByID(id)

	if err != nil {
		if errors.Is(err, repository.ErrMerchantNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": merchant,
	})

}

func (h *MerchantHandler) Update(c *gin.Context) {
	id := c.Param("id")

	update_id, err := uuid.Parse(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid merchant ID",
		})
		return
	}

	merchantIDValue, exists := c.Get("merchant_id")
	if !exists {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H {
			"error": "Authentication required",
		})
		return
	}

	merchantIDStr, ok := merchantIDValue.(string)

	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H {
			"error": "Invalid authentication context",
		})
		return
	}

	merchantID, err := uuid.Parse(merchantIDStr)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H {
			"error": "Invalid merchant identity",
		})
		return
	}

	if update_id != merchantID {
		c.JSON(http.StatusForbidden, gin.H {
			"error": "You cannnot update another merchant",
		})
		return
	}

	var req model.CreateMerchantRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	merchant, err := h.service.Update(update_id, merchantID, req)
	if err != nil {
		if errors.Is(err, service.ErrForbidden) {
			c.JSON(http.StatusForbidden, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": merchant,
	})

}

func (h *MerchantHandler) Delete(c *gin.Context) {
	id := c.Param("id")

	deleteId, err := uuid.Parse(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid merchant ID",
		})
		return
	}

	err = h.service.Delete(deleteId)

	if err != nil {
		if errors.Is(err, repository.ErrMerchantNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Merchant deleted successfully",
	})

}

func (h *MerchantHandler) Activate(c *gin.Context) {
	id := c.Param("id")

	activateId, err := uuid.Parse(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid merchant ID",
		})
		return
	}

	merchantIDValue, exists := c.Get("merchant_id")
	if !exists {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Authentication required",
		})
		return
	}

	merchantIDStr, ok := merchantIDValue.(string)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid authentication context",
		})
		return
	}

	merchantID, err := uuid.Parse(merchantIDStr)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid merchant identity",
		})
		return
	}

	if activateId != merchantID {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "You cannot activate another merchant",
		})
		return
	}

	err = h.service.Activate(activateId)

	if err != nil {
		if errors.Is(err, repository.ErrMerchantNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": "merchant status activated successfully!",
	})

}


func (h *MerchantHandler) Me(c *gin.Context) {
	merchantID, exists := c.Get("merchant_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H {
			"error": "Unauthorized",
		})
		return
	}

	merchantIDStr, ok := merchantID.(string)
	if !ok || merchantIDStr == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid merchant identity",
		})
		return
	}

	id, err := uuid.Parse(merchantIDStr)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H {
			"error": "Invalid merchant identity",
		})
		return
	}

	merchant, err := h.service.GetByID(id)

	if err != nil {
		if errors.Is(err, repository.ErrMerchantNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Merchant not found",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H {
		"data": merchant,
	})
}