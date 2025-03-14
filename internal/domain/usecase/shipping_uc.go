package usecase

import (
	"github.com/GratiaManullang03/ekspedisi-go/internal/domain/model"
	"github.com/GratiaManullang03/ekspedisi-go/internal/domain/repository"
)

type ShippingUseCase struct {
	repo *repository.ShippingRepository
}

func NewShippingUseCase(repo *repository.ShippingRepository) *ShippingUseCase {
	return &ShippingUseCase{repo: repo}
}

func (uc *ShippingUseCase) GetShipping(role string, nik string, costCenter string) ([]model.ShippingResponse, error) {
	return uc.repo.SelectShipping(role, nik, costCenter)
}

func (uc *ShippingUseCase) GetShippingByID(trID int) (model.ShippingResponse, error) {
	return uc.repo.SelectByID(trID)
}