package models

type StorageURL struct {
	Id          int    `json:"id"           gorm:"primaryKey"`
	OriginalURL string `json:"original_url"`
	ShortURL    string `json:"short_url"`
}
