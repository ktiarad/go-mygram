package services

import (
	"go-mygram/models"
	"go-mygram/params"
	"go-mygram/repositories"
	"net/http"
)

type SocialMediaServices struct {
	SocialMediaRepo repositories.SocialMediaRepo
}

func NewSocialMediaService(SocialMediaRepo repositories.SocialMediaRepo) *SocialMediaServices {
	return &SocialMediaServices{
		SocialMediaRepo: SocialMediaRepo,
	}
}

func (s *SocialMediaServices) CreateSocialMedia(req *params.SocialMediaCreate) *params.Response {
	var socialmedia = &models.SocialMedia{
		Name:           req.Name,
		SocialMediaUrl: req.SocialMediaUrl,
		UserID:         req.UserID,
	}

	id, err := s.SocialMediaRepo.CreateSocialMedia(socialmedia)

	if err != nil {
		return &params.Response{
			Status:         http.StatusInternalServerError,
			Error:          "INTERNAL SERVER ERROR, when create social media",
			AdditionalInfo: err.Error(),
		}
	}

	createdData, err := s.SocialMediaRepo.GetSocialMediaById(id)

	if err != nil {
		return &params.Response{
			Status:         http.StatusInternalServerError,
			Error:          "INTERNAL SERVER ERROR, when retrieve created social media",
			AdditionalInfo: err.Error(),
		}
	}
	createdSocialMedia := models.SocialMedia{
		ID:             id,
		Name:           createdData.Name,
		SocialMediaUrl: createdData.SocialMediaUrl,
		UserID:         createdData.UserID,
		// CreatedAt:      createdData.CreatedAt,
		// TODO : add CreatedAt untuk Payload
	}

	return &params.Response{
		Status:  http.StatusCreated,
		Payload: createdSocialMedia,
	}
}

func (s *SocialMediaServices) GetAllSocialMedias(id int) *params.Response {
	socialmedias, err := s.SocialMediaRepo.GetAllSocialMedias(id)

	if err != nil {
		return &params.Response{
			Status:         http.StatusInternalServerError,
			Error:          "INTERNAL SERVER ERROR, when get all social medias",
			AdditionalInfo: err.Error(),
		}
	}

	return &params.Response{
		Status:  http.StatusOK,
		Payload: socialmedias, // TODO : cek konten berisi :id, name, social_media_url, user_id, created_at, updated_at, User{id, username, profile_image_url}
	}
}

func (s *SocialMediaServices) UpdateSocialMedia(req *params.SocialMediaCreate, id int) *params.Response {
	var socialmedia = &models.SocialMedia{
		Name:           req.Name,
		SocialMediaUrl: req.SocialMediaUrl,
	}

	err := s.SocialMediaRepo.UpdateSocialMedia(socialmedia, id)

	if err != nil {
		return &params.Response{
			Status:         http.StatusInternalServerError,
			Error:          "INTERNAL SERVER ERROR, when update social media",
			AdditionalInfo: err.Error(),
		}
	}

	updatedData, err := s.SocialMediaRepo.GetSocialMediaById(id)

	if err != nil {
		return &params.Response{
			Status:         http.StatusInternalServerError,
			Error:          "INTERNAL SERVER ERROR, when retrieve updated social media",
			AdditionalInfo: err.Error(),
		}
	}

	return &params.Response{
		Status:  http.StatusOK,
		Payload: updatedData, // TODO : cek payload terdiri dari : id, name, social_media_url, user_id, updated_at
	}
}

func (s *SocialMediaServices) DeleteSocialMedia(id int) *params.Response {
	err := s.SocialMediaRepo.DeleteSocialMedia(id)
	if err != nil {
		return &params.Response{
			Status:         http.StatusInternalServerError,
			Error:          "INTERNAL SERVER ERROR, when delete social media",
			AdditionalInfo: err.Error(),
		}
	}

	return &params.Response{
		Status:  http.StatusOK,
		Message: "Your social media has been successfully deleted",
	}

}
