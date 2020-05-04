package api

import (
	"GoDemo/Project/Mall/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Main(c *gin.Context) {
	var banners, _ = model.GetTopBanner(5)
	var data = make(map[string]interface{}, 4)
	data["banner"] = banners
	var res = model.Result{Code: 0, Count: 0, Msg: "success", Data: data}
	c.JSON(http.StatusOK, res)
}
