package domain

type Customer struct {
    ID      uint   `gorm:"primaryKey" json:"id"`
    Name    string `json:"name"`
    Email   string `gorm:"unique" json:"email"`
    Phone   string `json:"phone"`
    Address string `json:"address"`
}

type CustomerRepository interface {
    Create(customer *Customer) error
    GetAll() ([]Customer, error)
    GetByID(id uint) (*Customer, error)
    Update(customer *Customer) error
    Delete(id uint) error
}
