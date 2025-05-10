package handlers

import (
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nil0j/jirafeitor/config"
	"github.com/nil0j/jirafeitor/models/responses"
	"github.com/nil0j/jirafeitor/repository"
)

func VideoPage(c *gin.Context) {
	c.HTML(http.StatusOK, "video", gin.H{})
}

func GetVideo(c *gin.Context) {
	folderPath := config.Data.Filesystem + c.Param("id")
	files, err := os.ReadDir(folderPath)
	if err != nil {
		responses.HandleError(c, err)
		return
	}

	if len(files) == 0 {
		responses.HandleError(c, err)
		return
	}

	var targetFileName string
	if filepath.Ext(files[0].Name()) == ".png" {
		targetFileName = files[1].Name()
	} else {
		targetFileName = files[0].Name()
	}

	filePath := filepath.Join(folderPath, targetFileName)
	c.File(filePath)
}

func GetVideoThumbnail(c *gin.Context) {
	folderPath := config.Data.Filesystem + c.Param("id")
	files, err := os.ReadDir(folderPath)
	if err != nil {
		responses.HandleError(c, err)
		return
	}

	if len(files) == 0 {
		responses.HandleError(c, err)
		return
	}

	filePath := filepath.Join(folderPath, "thumbnail.png")
	c.File(filePath)
}

func GetVideoInfo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		responses.HandleError(c, err)
		return
	}

	videoInfo, err := repository.GetVideoInfo(id)
	if err != nil {
		responses.HandleError(c, err)
		return
	}

	responses.HandleVideoInfo(c, videoInfo)
}

func GetRecentVideos(c *gin.Context) {
	var limit int
	{
		limitInput := c.Query("limit")

		var err error
		limit, err = strconv.Atoi(limitInput)
		if err != nil {
			limit = 10 // default limit
		}
	}

	repository.GetRecentVideos(limit)
}
