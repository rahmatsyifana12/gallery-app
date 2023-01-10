package models

type MemoryTags struct {
	MemoryID 	uint64 `gorm:"not null"`
	TagsID 		uint64 `gorm:"not null"`
}
