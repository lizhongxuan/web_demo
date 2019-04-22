/*
@Time : 2019-04-22 16:37 
@Author : zhongxuanli
@File : test_handler
@Software: GoLand
*/

package test

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FileHandler(c *gin.Context) {
	fname := c.Query("_w_fname")
	fmt.Println("fname:",fname)
	c.JSON(http.StatusOK, fname)
}