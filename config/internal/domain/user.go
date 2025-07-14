package domain

type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Name     string `json:"name"`
	Email    string `gorm:"uniqueIndex" json:"email"`
	Password string `json:"-"`
	Role     string `json:"role"`
}

type UserRepository interface {
	FindByEmail(email string) (*User, error)
	Create(user *User) error
}