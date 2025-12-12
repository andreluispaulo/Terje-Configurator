package db

import (
	"time"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type FileHistory struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	FilePath  string    `gorm:"index" json:"filePath"`
	CreatedAt time.Time `json:"createdAt"`
	Content   string    `json:"-"` // Don't send content in list view
}

var DB *gorm.DB

func InitDB(dbPath string) error {
	var err error
	DB, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		return err
	}

	return DB.AutoMigrate(&FileHistory{})
}

func SaveSnapshot(path string, content string) error {
	history := FileHistory{
		FilePath:  path,
		Content:   content,
		CreatedAt: time.Now(),
	}
	return DB.Create(&history).Error
}

func GetHistory(path string) ([]FileHistory, error) {
	var history []FileHistory
	// Return last 20 entries
	err := DB.Where("file_path = ?", path).Order("created_at desc").Limit(20).Find(&history).Error
	return history, err
}

func GetSnapshot(id uint) (*FileHistory, error) {
	var history FileHistory
	err := DB.First(&history, id).Error
	return &history, err
}
