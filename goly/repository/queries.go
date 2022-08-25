package repository

import "gorm.io/gorm"

type Repository interface {
	GetGolies() ([]Goly, error)
	GetGoly(id uint64) (Goly, error)
	GetGolyByURL(url string) (Goly, error)
	CreateGoly(goly *Goly) error
	UpdateGoly(goly *Goly) error
	DeleteGoly(goly *Goly) error
}

type repo struct {
	DB *gorm.DB
}

func (p *repo) GetGolies() ([]Goly, error) {
	golies := []Goly{}
	tx := p.DB.Find(&golies)
	return golies, tx.Error
}

func (p *repo) GetGoly(id uint64) (Goly, error) {
	goly := Goly{}
	tx := p.DB.Where("id = ?", id).First(&goly)
	return goly, tx.Error
}

func (p *repo) GetGolyByURL(url string) (Goly, error) {
	goly := Goly{}
	tx := p.DB.Where("goly = ?", url).First(&goly)
	return goly, tx.Error
}

func (p *repo) CreateGoly(goly *Goly) error {
	return p.DB.Create(goly).Error
}

func (p *repo) UpdateGoly(goly *Goly) error {
	return p.DB.Save(&goly).Error
}

func (p *repo) DeleteGoly(goly *Goly) error {
	return p.DB.Unscoped().Delete(&goly).Error
}

func NewRepository(db *gorm.DB) Repository {
	return &repo{
		DB: db,
	}
}
