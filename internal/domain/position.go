package domain

type Position struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	TitlePosition     string `json:"name"`
	Description    string `gorm:"uniqueIndex" json:"email"`
	IsActive      bool   `gorm:"default:true" json:"is_active"`
}
