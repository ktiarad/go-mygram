package controllers

import (
	"go-mygram/params"
	"go-mygram/services"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type CommentController struct {
	commentService *services.CommentServices
}

func NewCommentController(commentServices *services.CommentServices) *CommentController {
	return &CommentController{
		commentService: commentServices,
	}
}

func (c *CommentController) CreateComment(ctx *gin.Context) {
	var comment params.CommentCreate

	err := ctx.BindJSON(&comment)

	if err != nil {
		response := params.Response{
			Status:         http.StatusBadRequest,
			Error:          "BAD REQUEST, when create comment",
			AdditionalInfo: err.Error(),
		}

		params.WriteJsonResponse(ctx.Writer, &response)
	}

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := int(userData["userID"].(float64))

	comment.UserID = userID

	response := c.commentService.CreateComment(&comment)

	params.WriteJsonResponse(ctx.Writer, response)
}

func (c *CommentController) GetAllComments(ctx *gin.Context) {
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := int(userData["userID"].(float64))

	response := c.commentService.GetAllComments(userID)

	params.WriteJsonResponse(ctx.Writer, response)
}

func (c *CommentController) UpdateComment(ctx *gin.Context) {
	var comment params.CommentUpdate

	err := ctx.BindJSON(&comment)

	if err != nil {
		response := params.Response{
			Status:         http.StatusBadRequest,
			Error:          "BAD REQUEST, when update comment",
			AdditionalInfo: err.Error(),
		}

		params.WriteJsonResponse(ctx.Writer, &response)
	}

	commentId := ctx.Param("commentID")
	id, err := strconv.Atoi(commentId)
	if err != nil {
		response := params.Response{
			Status:         http.StatusBadRequest,
			Error:          "BAD REQUEST, when get param comment id",
			AdditionalInfo: err.Error(),
		}
		params.WriteJsonResponse(ctx.Writer, &response)
	}

	response := c.commentService.UpdateComment(&comment, id)
	params.WriteJsonResponse(ctx.Writer, response)

}

func (c *CommentController) DeleteComment(ctx *gin.Context) {
	commentId := ctx.Param("commentID")
	id, err := strconv.Atoi(commentId)

	if err != nil {
		response := &params.Response{
			Status:         http.StatusBadRequest,
			Message:        "BAD REQUEST, when get param comment id",
			AdditionalInfo: err.Error(),
		}

		params.WriteJsonResponse(ctx.Writer, response)
	}

	response := c.commentService.DeleteComment(id)
	params.WriteJsonResponse(ctx.Writer, response)
}
