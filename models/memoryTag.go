package models

type MemoryTag struct {
	MemoryID 	uint64 `gorm:"not null"`
	TagID 		uint64 `gorm:"not null"`
}
