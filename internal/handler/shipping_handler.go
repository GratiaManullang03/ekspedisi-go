package handler

import (
	"log"
	// "net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/GratiaManullang03/ekspedisi-go/internal/domain/usecase"
	"github.com/GratiaManullang03/ekspedisi-go/pkg/utils"
)

type ShippingHandler struct {
	useCase *usecase.ShippingUseCase
}

func NewShippingHandler(useCase *usecase.ShippingUseCase) *ShippingHandler {
	return &ShippingHandler{useCase: useCase}
}

// GetShipping handler untuk mendapatkan daftar shipping
func (h *ShippingHandler) GetShipping(c *gin.Context) {
	// Ambil data user dari context (diset oleh middleware)
	nik, _ := c.Get("nik")
	costCenter, _ := c.Get("costCenter")
	levels, _ := c.Get("levels")

	// Konversi ke tipe yang benar
	nikStr := nik.(string)
	costCenterStr := costCenter.(string)
	userLevels := levels.([]int)

	// Tentukan role berdasarkan level
	var role string
	if contains(userLevels, 0) {
		role = "SUPER_ADMIN"
	} else if contains(userLevels, 100) {
		role = "MANAGER"
	} else {
		role = "USER"
	}

	log.Printf("Role Identified: %s", role)

	// Panggil use case
	shipping, err := h.useCase.GetShipping(role, nikStr, costCenterStr)
	if err != nil {
		log.Printf("Error in get_shipping: %v", err)
		utils.ErrorResponse(c, err.Error())
		return
	}

	log.Printf("Retrieved Shipping Data: %d records", len(shipping))
	utils.SuccessResponse(c, "success", shipping)
}

// GetShippingByID handler untuk mendapatkan shipping berdasarkan ID
func (h *ShippingHandler) GetShippingByID(c *gin.Context) {
	// Ambil ID dari URL
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.ErrorResponse(c, "Invalid ID format")
		return
	}

	// Panggil use case
	shipping, err := h.useCase.GetShippingByID(id)
	if err != nil {
		utils.ErrorResponse(c, err.Error())
		return
	}

	utils.SuccessResponse(c, "success", []interface{}{shipping})
}

// Utils function untuk memeriksa apakah slice berisi nilai tertentu
func contains(slice []int, val int) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}