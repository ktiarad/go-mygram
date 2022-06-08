package services

import (
	"go-mygram/models"
	"go-mygram/params"
	"go-mygram/repositories"
	"net/http"
)

type CommentServices struct {
	CommentRepo repositories.CommentRepo
}

func NewCommentService(CommentRepo repositories.CommentRepo) *CommentServices {
	return &CommentServices{
		CommentRepo: CommentRepo,
	}
}

func (c *CommentServices) CreateComment(req *params.CommentCreate) *params.Response {
	var comment = &models.Comment{
		Message: req.Message,
		PhotoID: req.PhotoID,
		UserID:  req.UserID,
	}

	id, err := c.CommentRepo.CreateComment(comment)
	if err != nil {
		return &params.Response{
			Status:         http.StatusInternalServerError,
			Error:          "INTERNAL SERVER ERROR, when create comment",
			AdditionalInfo: err.Error(),
		}
	}

	dataCreated, err := c.CommentRepo.GetCommentById(id)
	if err != nil {
		return &params.Response{
			Status:         http.StatusInternalServerError,
			Error:          "INTERNAL SERVER ERROR, when get param comment id",
			AdditionalInfo: err.Error(),
		}
	}

	return &params.Response{
		Status:  http.StatusCreated,
		Payload: dataCreated,
	}

}

func (c *CommentServices) GetAllComments(id int) *params.Response {
	comments, err := c.CommentRepo.GetAllComments(id)

	if err != nil {
		return &params.Response{
			Status:         http.StatusInternalServerError,
			Error:          "INTERNAL SERVER ERROR, when get all comments",
			AdditionalInfo: err.Error(),
		}
	}

	return &params.Response{
		Status:  http.StatusOK,
		Payload: comments, // TODO : payload berisi data User dan Photo, dilakukan join
	}
}

func (c *CommentServices) UpdateComment(req *params.CommentUpdate, id int) *params.Response {

	var comment = &models.Comment{
		Message: req.Message,
	}

	err := c.CommentRepo.UpdateComment(comment, id)

	if err != nil {
		return &params.Response{
			Status:         http.StatusInternalServerError,
			Error:          "INTERNAL SERVER ERROR, when update comment",
			AdditionalInfo: err.Error(),
		}
	}

	updatedData, err := c.CommentRepo.GetCommentById(id)

	if err != nil {
		return &params.Response{
			Status:         http.StatusInternalServerError,
			Error:          "INTERNAL SERVER ERROR, when get updated comment",
			AdditionalInfo: err.Error(),
		}
	}

	return &params.Response{
		Status:  http.StatusOK,
		Payload: &updatedData, // TODO : updatedData terdiri dari : id, title, caption, photo_url, user_id, updated_at (yang merupakan models.Photo)
	}
}

func (c *CommentServices) DeleteComment(id int) *params.Response {
	err := c.CommentRepo.DeleteComment(id)

	if err != nil {
		return &params.Response{
			Status:         http.StatusInternalServerError,
			Error:          "INTERNAL SERVER ERROR, when delete comment",
			AdditionalInfo: err.Error(),
		}
	}

	return &params.Response{
		Status:  http.StatusOK,
		Message: "Your comment has been successfully deleted",
	}
}
