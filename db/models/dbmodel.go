package models

import (
	"time"
)

type DbModel struct {
	Id         int       `gorm:"primaryKey"`
	Identifier string    `gorm:"default:newid()"`
	CreatedOn  time.Time `gorm:"default:getdate()"`
}

/*
func (m *DbModel) AfterFind(tx *gorm.DB) (err error) {
	m.FixUUID()
	return
}

func (m *DbModel) AfterCreate(tx *gorm.DB) (err error) {
	m.FixUUID()
	return
}

// Hack to fix SQL Server guids
func (m *DbModel) FixUUID() {
	uuid := &m.Identifier
	uuid[0], uuid[1], uuid[2], uuid[3] = uuid[3], uuid[2], uuid[1], uuid[0]
	uuid[4], uuid[5] = uuid[5], uuid[4]
	uuid[6], uuid[7] = uuid[7], uuid[6]
}
*/
