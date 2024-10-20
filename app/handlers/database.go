package handlers

import (
	"fmt"

	"gorm.io/gorm"

	"url_shortener/models"
)

type handler struct {
	DB *gorm.DB
}

func New(db *gorm.DB) handler {
	return handler{db}
}

func (h *handler) AddURLPair(original, short string) {
	if result := h.DB.Create(&models.StorageURL{
		OriginalURL: original,
		ShortURL:    short,
	}); result.Error != nil {
		fmt.Println(result.Error)
	}
}

func (h *handler) GetLongByShort(short string) (string, error) {
	var storage models.StorageURL

	if result := h.DB.Find(
		&storage,
		models.StorageURL{ShortURL: short},
	); result.Error != nil {
		return "", result.Error
	}
	return storage.OriginalURL, nil
}

func (h *handler) GetShortByLong(original string) (string, error) {
	var storage models.StorageURL

	if result := h.DB.Find(
		&storage,
		models.StorageURL{OriginalURL: original},
	); result.Error != nil {
		return "", result.Error
	}
	return storage.ShortURL, nil
}
