package models

type MemoryTag struct {
	MemoryID 	uint64 `gorm:"not null"`
	TagsID 		uint64 `gorm:"not null"`
}
