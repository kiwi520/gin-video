package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"gin-video/models"
)

//列表数据
func Users(c *gin.Context) {
	var user models.User
	user.Username = c.Request.FormValue("username")
	user.Password = c.Request.FormValue("password")
	result, err := user.Users()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "抱歉未找到相关信息",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data":   result,
	})
}

//添加数据
func Add(c *gin.Context) {
	var user models.User


	err := c.BindJSON(&user)
	if err != nil {
		//log.Info(err)
		c.JSON(200, gin.H{"errcode": 400, "description": "Post Data Err"})
		return
	} else {
		id, err := user.Insert()

		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    -1,
				"message": "添加失败",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code":  1,
			"message": "添加成功",
			"data":    id,
		})
	}


}

//修改数据
func Update(c *gin.Context) {
	var user models.User
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	err= c.BindJSON(&user)
	if err != nil {
		//log.Info(err)
		c.JSON(200, gin.H{"errcode": 400, "description": "Post Data Err"})
		return
	} else {
		result, err := user.Update(id)
		if err != nil || result.ID == 0 {
			c.JSON(http.StatusOK, gin.H{
				"code":    -1,
				"message": "修改失败",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code":  1,
			"message": "修改成功",
		})
	}

}

//删除数据
func Destroy(c *gin.Context) {
	var user models.User
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	result, err := user.Destroy(id)
	if err != nil || result.ID == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "删除失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":  1,
		"message": "删除成功",
	})
}