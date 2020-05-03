package api

import (
	"GoDemo/Project/Mall/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Main(c *gin.Context) {
	var banners, _ = model.GetTopBanner(5)
	c.JSON(http.StatusOK, banners)
}
