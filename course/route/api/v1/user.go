package v1

import (
	"course/models"
	"course/pkg/e"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
)

const (
	NAME     = "name"
	STATE    = "state"
	CREATED  = "created_by"
	MODIFIED = "modified_by"
	ID       = "id"
)

func GetUser(c *gin.Context) {
	userName := com.StrTo(c.Param("userName")).String()
	password := com.StrTo(c.Param("password")).String()
	user := models.GetUser(userName, password)
	if user {
		code := e.SUCCESS

		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
			"data": "data",
		})
	}

}
