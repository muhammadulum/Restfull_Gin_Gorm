package usecase

import (
    "gin_restfull/internal/domain"
    "gin_restfull/internal/model"
)

type CustomerUseCase struct {
    repo domain.CustomerRepository
}

func NewCustomerUseCase(repo domain.CustomerRepository) *CustomerUseCase {
    return &CustomerUseCase{repo: repo}
}

func (uc *CustomerUseCase) Create(req model.CustomerRequest) error {
    customer := domain.Customer{
        Name:    req.Name,
        Email:   req.Email,
        Phone:   req.Phone,
        Address: req.Address,
    }
    return uc.repo.Create(&customer)
}

func (uc *CustomerUseCase) GetAll() ([]domain.Customer, error) {
    return uc.repo.GetAll()
}

func (uc *CustomerUseCase) GetByID(id uint) (*domain.Customer, error) {
    return uc.repo.GetByID(id)
}

func (uc *CustomerUseCase) Update(id uint, req model.CustomerRequest) error {
    existing, err := uc.repo.GetByID(id)
    if err != nil {
        return err
    }

    existing.Name = req.Name
    existing.Email = req.Email
    existing.Phone = req.Phone
    existing.Address = req.Address

    return uc.repo.Update(existing)
}

func (uc *CustomerUseCase) Delete(id uint) error {
    return uc.repo.Delete(id)
}
