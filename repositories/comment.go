package repositories

import (
	"go-mygram/models"

	"gorm.io/gorm"
)

type CommentRepo interface {
	CreateComment(request *models.Comment) (int, error)
	GetAllComments(id int) (*[]models.Comment, error)
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

func (c *commentRepo) GetAllComments(id int) (*[]models.Comment, error) {
	var comments []models.Comment

	result := c.db.Where("user_id=?", id).Find(&comments)
	err := result.Error

	return &comments, err
}

func (c *commentRepo) GetCommentById(id int) (*models.Comment, error) {
	var comment models.Comment

	result := c.db.First(&comment, "id=?", id)
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
