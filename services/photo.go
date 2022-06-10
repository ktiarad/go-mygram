package services

import (
	"go-mygram/models"
	"go-mygram/params"
	"go-mygram/repositories"
	"net/http"
)

type PhotoServices struct {
	PhotoRepo repositories.PhotoRepo
}

func NewPhotoService(PhotoRepo repositories.PhotoRepo) *PhotoServices {
	return &PhotoServices{
		PhotoRepo: PhotoRepo,
	}
}

func (p *PhotoServices) CreatePhoto(req *params.PhotoCreate) *params.Response {
	var photo = &models.Photo{
		Title:    req.Title,
		Caption:  req.Caption,
		PhotoUrl: req.PhotoUrl,
		UserID:   req.UserID,
	}

	id, err := p.PhotoRepo.CreatePhoto(photo)

	if err != nil {
		return &params.Response{
			Status:         http.StatusInternalServerError,
			Error:          "INTERNAL SERVER ERROR, when create photo",
			AdditionalInfo: err.Error(),
		}
	}

	photoDb, err := p.PhotoRepo.GetPhotoById(id)

	if err != nil {
		return &params.Response{
			Status:         http.StatusInternalServerError,
			Error:          "INTERNAL SERVER ERROR, when get photo by id",
			AdditionalInfo: err.Error(),
		}
	}

	payload := map[string]interface{}{
		"id":         id,
		"title":      photoDb.Title,
		"caption":    photoDb.Caption,
		"photo_url":  photoDb.PhotoUrl,
		"user_id":    photoDb.UserID,
		"created_at": photoDb.CreatedAt,
	}

	return &params.Response{
		Status:  http.StatusCreated,
		Message: "CREATE PHOTO SUCCESS",
		Payload: payload,
	}
}

func (p *PhotoServices) GetAllPhotos(id int) *params.Response {
	photos, err := p.PhotoRepo.GetAllPhotos(id)

	if err != nil {
		return &params.Response{
			Status:         http.StatusInternalServerError,
			Error:          "INTERNAL SERVER ERROR, when get all photos",
			AdditionalInfo: err.Error(),
		}
	}

	return &params.Response{
		Status:  http.StatusOK,
		Payload: photos,
	}

}

func (p *PhotoServices) UpdatePhoto(req *params.PhotoCreate, id int) *params.Response {
	var photo = &models.Photo{
		Title:    req.Title,
		Caption:  req.Caption,
		PhotoUrl: req.PhotoUrl,
	}

	err := p.PhotoRepo.UpdatePhoto(photo, id)

	if err != nil {
		return &params.Response{
			Status:         http.StatusInternalServerError,
			Error:          "INTERNAL SERVER ERROR, when update photo",
			AdditionalInfo: err.Error(),
		}
	}

	photoDb, err := p.PhotoRepo.GetPhotoById(id)

	payload := map[string]interface{}{
		"id":         id,
		"title":      photoDb.Title,
		"caption":    photoDb.Caption,
		"photo_url":  photoDb.PhotoUrl,
		"user_id":    photoDb.UserID,
		"updated_at": photoDb.UpdatedAt,
	}
	if err != nil {
		return &params.Response{
			Status:         http.StatusInternalServerError,
			Error:          "INTERNAL SERVER ERROR, when retrieve updated photo",
			AdditionalInfo: err.Error(),
		}
	}

	return &params.Response{
		Status:  http.StatusOK,
		Payload: payload,
	}

}

func (p *PhotoServices) DeletePhoto(id int) *params.Response {
	err := p.PhotoRepo.DeletePhoto(id)
	if err != nil {
		return &params.Response{
			Status:         http.StatusInternalServerError,
			Error:          "INTERNAL SERVER ERROR, when delete photo",
			AdditionalInfo: err.Error(),
		}
	}

	return &params.Response{
		Status:  http.StatusOK,
		Message: "Your photo has been successfully deleted",
	}

}
