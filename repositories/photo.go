package repositories

import (
	"go-mygram/models"

	"gorm.io/gorm"
)

type PhotoRepo interface {
	CreatePhoto(request *models.Photo) (int, error)
	GetAllPhotos() (*[]models.Photo, error)
	UpdatePhoto(request *models.Photo, id int) error
	GetPhotoById(id int) (*models.Photo, error)
	DeletePhoto(id int) error
}

type photoRepo struct {
	db *gorm.DB
}

func NewPhotoRepo(db *gorm.DB) PhotoRepo {
	return &photoRepo{
		db: db,
	}
}

func (p *photoRepo) CreatePhoto(request *models.Photo) (int, error) {
	result := p.db.Create(request)
	err := result.Error
	id := request.ID

	return id, err
}

func (p *photoRepo) GetAllPhotos() (*[]models.Photo, error) {
	var photos []models.Photo

	// TODO : get all photo bersama dengan get User
	result := p.db.Preload("Photos").Find(&photos)
	err := result.Error

	return &photos, err
}

func (p *photoRepo) GetPhotoById(id int) (*models.Photo, error) {
	var photo models.Photo

	result := p.db.First(&photo, "id=?", id)

	err := result.Error

	return &photo, err
}

func (p *photoRepo) UpdatePhoto(request *models.Photo, id int) error {
	var photo []models.Photo

	result := p.db.Model(&photo).Where("id=?", id).Updates(models.Photo{
		Title:    request.Title,
		Caption:  request.Caption,
		PhotoUrl: request.PhotoUrl,
	})

	err := result.Error

	return err
}

func (p *photoRepo) DeletePhoto(id int) error {
	var photo models.Photo

	result := p.db.Model(&photo).Where("id=?", id).Delete(&photo)
	err := result.Error

	return err
}
