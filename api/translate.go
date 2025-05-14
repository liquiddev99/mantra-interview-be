package api

import (
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/liquiddev99/mantra-interview-be/util"
)

func (server *Server) upload(ctx *gin.Context) {
	toLang := ctx.DefaultQuery("toLang", "English")
	fontFamily := ctx.DefaultQuery("fontFamily", "WildWords")

	file, err := ctx.FormFile("file")
	if err != nil {
		fmt.Println("error", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "No file received"})
		return
	}

	src, err := file.Open()
	mlServerAddress := fmt.Sprintf(
		server.config.MlServerAddress+"/translate/?fromLang=%s&toLang=%s&translationModel=%s&fontFamily=%s&deeplAPIKey=%s",
		"",
		url.QueryEscape(toLang),
		"Gemini2.0",
		fontFamily,
		"",
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, responseError(err))
		return
	}
	defer src.Close()

	fileBytes, err := io.ReadAll(src)
	imageData, _, err := util.SendImageDataToMl(mlServerAddress, fileBytes, file.Filename)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, responseError(err))
		return
	}

	ctx.Data(http.StatusOK, "image/jpeg", imageData)
}
