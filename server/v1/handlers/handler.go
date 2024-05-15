package handlers

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	connector "github.com/bartoszpiechnik25/static-files-server/aws-connector"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	s3Connection connector.Agent
}

func NewHandler() (*Handler, error) {
	agent, err := connector.NewS3PersistentAgent(context.TODO())
	if err != nil {
		return nil, err
	}
	return &Handler{
		s3Connection: agent,
	}, nil
}

func (h *Handler) OptionsHandler(ctx *gin.Context) {
	ctx.Header("Access-Controll-Allow-Headers", "Content-Type, Authorization")
	ctx.Header("Access-Controll-Allow-Methods", "POST, GET, PUT, OPTIONS")
	ctx.Header("Accept", "application/json,multipart/form-data")

	ctx.AbortWithStatus(http.StatusNoContent)
}

func (h *Handler) CreateAsset(ctx *gin.Context) {
	if contains := strings.Contains(ctx.GetHeader("Content-Type"), "multipart/form-data"); !contains {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"code":    strconv.Itoa(http.StatusBadRequest),
				"message": "Missing header 'multipart/form-data'",
			},
		)
		return
	}

	facilityId := ctx.Param("facility-id")
	file, err := ctx.FormFile("file")

	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{
				"message": "could not form file to upload",
			},
		)
		return
	}

	fileContent, err := file.Open()
	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{
				"message": "could not open file for upload",
			},
		)
		return
	}
	dirs, filename := []string{facilityId}, file.Filename
	filePath := connector.CreateFilepath(dirs, filename)

	if exist, _ := h.s3Connection.FileExists(ctx, filePath); exist {
		ctx.AbortWithStatusJSON(
			http.StatusForbidden,
			gin.H{
				"message": fmt.Sprintf("file: %s already exist", filePath),
			},
		)
		return
	}

	err = h.s3Connection.UploadFile(ctx, dirs, filename, fileContent)

	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"message": err.Error(),
			},
		)
		return
	}

	ctx.JSON(
		http.StatusCreated,
		gin.H{
			"message": "asset uploaded successfully",
			"url":     fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", connector.BUCKET_NAME, connector.BUCKET_REGION, filePath),
		},
	)
}
