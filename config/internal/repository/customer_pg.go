package repository

import (
    "gin_restfull/internal/domain"
    "gorm.io/gorm"
)

type CustomerPG struct {
    db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) domain.CustomerRepository {
    return &CustomerPG{db}
}

func (r *CustomerPG) Create(c *domain.Customer) error {
    return r.db.Create(c).Error
}

func (r *CustomerPG) GetAll() ([]domain.Customer, error) {
    var customers []domain.Customer
    err := r.db.Find(&customers).Error
    return customers, err
}

func (r *CustomerPG) GetByID(id uint) (*domain.Customer, error) {
    var customer domain.Customer
    err := r.db.First(&customer, id).Error
    return &customer, err
}

func (r *CustomerPG) Update(c *domain.Customer) error {
    return r.db.Save(c).Error
}

func (r *CustomerPG) Delete(id uint) error {
    return r.db.Delete(&domain.Customer{}, id).Error
}
