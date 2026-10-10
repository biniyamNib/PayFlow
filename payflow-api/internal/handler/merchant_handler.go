package handler

import (
	"net/http"
	"uuid"

	"payflow/internal/model"
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
	}

	merchant, err := h.service.GetByID(id)

	if err != nil {
		if err.Error() == "merchant not found" {
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

	var req model.CreateMerchantRequest

	err = c.ShouldBindJSON(&req)

	if err != nil {
		if err.Error() == "merchant not found" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	merchant, err := h.service.Update(update_id, req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
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
		if err.Error() == "merchant not found" {
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

	err = h.service.Activate(activateId)

	if err != nil {
		if err.Error() == "merchant not found" {
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

	c.JSON(http.StatusOK, gin.H {
		"data": "merchant status activated successfully!",
	})

}
