package repositories

import (
	"go-mygram/models"

	"gorm.io/gorm"
)

type CommentRepo interface {
	CreateComment(request *models.Comment) (int, error)
	GetAllComments() (*[]models.Comment, error)
	GetCommentById(id int) (*models.Comment, error)
	UpdateComment(request *models.Comment, id int) error
	DeleteComment(id int) error
}

type commentRepo struct {
	db *gorm.DB
}

func NewCommentRepo(db *gorm.DB) CommentRepo {
	return &commentRepo{
		db: db,
	}
}

func (c *commentRepo) CreateComment(request *models.Comment) (int, error) {
	result := c.db.Create(request)
	err := result.Error
	id := request.ID

	return id, err
}

func (c *commentRepo) GetAllComments() (*[]models.Comment, error) {
	var comment []models.Comment

	result := c.db.Preload("Comments").Find(&comment)
	err := result.Error

	return &comment, err
}

func (c *commentRepo) GetCommentById(id int) (*models.Comment, error) {
	var comment models.Comment

	result := c.db.First("id=?", id, &comment)
	err := result.Error

	return &comment, err
}

func (c *commentRepo) UpdateComment(request *models.Comment, id int) error {
	var comment models.Comment

	result := c.db.Model(&comment).Where("id=?", id).Updates(models.Comment{
		Message: request.Message,
	})

	err := result.Error

	return err
}

func (c *commentRepo) DeleteComment(id int) error {
	var comment models.Comment

	result := c.db.Model(&comment).Where("id=?", id).Delete(&comment)

	err := result.Error

	return err
}
