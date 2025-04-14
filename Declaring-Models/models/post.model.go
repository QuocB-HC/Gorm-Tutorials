package models

type Post struct {
	ID               uint
	Title            string
	Content          string
	AuditInfo `gorm:"embedded;embeddedPrefix:audit_"`
}