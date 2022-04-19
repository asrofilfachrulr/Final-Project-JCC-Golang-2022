package sql

type UserCredential struct {
	CredentialID uint   `gorm:"primaryKey"`
	UserID       uint   `gorm:"not null"`
	Password     string `gorm:"type:varchar;not null"`
	BearerToken  string
}
