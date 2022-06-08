package repositories

import (
	"go-mygram/models"

	"gorm.io/gorm"
)

type SocialMediaRepo interface {
	CreateSocialMedia(request *models.SocialMedia) (int, error)
	GetSocialMediaById(id int) (*models.SocialMedia, error)
	GetAllSocialMedias(id int) (*[]models.SocialMedia, error)
	UpdateSocialMedia(request *models.SocialMedia, id int) error
	DeleteSocialMedia(id int) error
}

type socialMediaRepo struct {
	db *gorm.DB
}

func NewMediaSocialRepo(db *gorm.DB) SocialMediaRepo {
	return &socialMediaRepo{
		db: db,
	}
}

func (s *socialMediaRepo) CreateSocialMedia(request *models.SocialMedia) (int, error) {
	result := s.db.Create(&request)

	err := result.Error
	id := request.ID

	return id, err
}

func (s *socialMediaRepo) GetSocialMediaById(id int) (*models.SocialMedia, error) {
	var socialmedia models.SocialMedia

	result := s.db.First(&socialmedia, "id=?", id)
	err := result.Error

	return &socialmedia, err
}

func (s *socialMediaRepo) GetAllSocialMedias(id int) (*[]models.SocialMedia, error) {
	var socialmedias []models.SocialMedia

	result := s.db.Model(&socialmedias).Where("user_id=?", id).Find(&socialmedias)
	err := result.Error

	return &socialmedias, err
}

func (s *socialMediaRepo) UpdateSocialMedia(request *models.SocialMedia, id int) error {
	var socialmedia models.SocialMedia

	result := s.db.Model(&socialmedia).Where("id=?", id).Updates(models.SocialMedia{
		Name:           request.Name,
		SocialMediaUrl: request.SocialMediaUrl,
	})
	err := result.Error

	return err
}

func (s *socialMediaRepo) DeleteSocialMedia(id int) error {
	var socialmedia models.SocialMedia

	result := s.db.Model(&socialmedia).Where("id=?", id).Delete(&socialmedia)
	err := result.Error

	return err
}
