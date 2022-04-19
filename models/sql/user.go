package sql

type User struct {
	ID             uint   `gorm:"primaryKey"`
	Fullname       string `gorm:"type:varchar"`
	Username       string `gorm:"type:varchar(20);not null;unique"`
	Email          string `gorm:"type:varchar(50);not null;unique"`
	PhoneNumber    string `gorm:"type:varchar(20);not null;unique"`
	AddressLine    string
	UserCredential *UserCredential `gorm:"constraint:OnDelete:CASCADE;"`
}
