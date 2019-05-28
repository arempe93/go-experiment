package models

type Audit struct {
	ID     uint   `gorm:"primary_key"`
	Action string `gorm:"size:64"`
	Data   string `gorm:"type:text"`
}
