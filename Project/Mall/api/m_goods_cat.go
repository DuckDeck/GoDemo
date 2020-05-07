package api

import (
	"GoDemo/Project/Mall/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//理论上后端的接口要和前端分开
func getGoodCats(c *gin.Context) {
	var cat, yes = c.Params.Get("cat")
	if !yes {
		var res = model.Result{Code: 0, Count: 0, Msg: "success", Data: nil}
		c.JSON(http.StatusOK, res)
		return
	}
	var catID, err = strconv.Atoi(cat)
	if err != nil {
		var res = model.Result{Code: 0, Count: 0, Msg: "success", Data: nil}
		c.JSON(http.StatusOK, res)
		return
	}

	var arr = []model.GoodsCat{model.GoodsCat{ID: 0, PID: 0, Sort: 0, Name: "热门推荐", FullPath: "热门推荐", Level: 1}}
	var tmp, _ = model.GetCats(catID)

	arr = append(arr, tmp...)
	var res = model.Result{Code: 0, Count: 0, Msg: "success", Data: cat}
	c.JSON(http.StatusOK, res)
}
