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

	return &params.Response{
		Status:  http.StatusCreated,
		Message: "CREATE PHOTO SUCCESS",
		Payload: id, // TODO : payload berupa id, title, caption, photo_url, user_id, created_at
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
		Payload: photos, // TODO : cek payload terdiri dari []models.Photo : id, title, caption, photo_url, user_id, created_at, updated_at, User{email, username}

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

	updatedData, err := p.PhotoRepo.GetPhotoById(id)

	var updatedPhoto = models.Photo{
		ID:        updatedData.ID,
		Title:     updatedData.Title,
		Caption:   updatedData.Caption,
		PhotoUrl:  updatedData.PhotoUrl,
		UserID:    updatedData.UserID,
		UpdatedAt: updatedData.UpdatedAt,
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
		Payload: updatedPhoto,
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
