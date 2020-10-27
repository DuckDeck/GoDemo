package api

import (
	"GoDemo/Project/Mall/model"
	"GoDemo/Project/Mall/tool"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Main(c *gin.Context) {
	var banners, _ = model.GetTopBanner(5)
	var data = make(map[string]interface{}, 4)
	var recommend, _ = model.GetRecommend(4)
	var current = make(map[string]string, 2)
	current["title"] = "本周流行"
	current["img"] = "/static/img/recommend_bg.794b662.jpg"
	data["banner"] = banners
	data["recommend"] = recommend
	data["currency"] = current
	var res = model.Result{Code: 0, Count: 0, Msg: "success", Data: data}
	c.JSON(http.StatusOK, res)
}

func List(c *gin.Context) {
	var res = tool.CheckNum(c, "type")
	if res.Code != 0 {
		c.JSON(http.StatusOK, res)
		return
	}
	res = tool.CheckPage(c)
	if res.Code != 0 {
		c.JSON(http.StatusOK, res)
		return
	}
	// var type = c.GetInt("type")
	// var pageIndex = c.GetInt("pageindex")
	// var pageSize = c.GetInt("pagesize")

}
