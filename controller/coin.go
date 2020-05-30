package controllers

import (
	"coinproject/models"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetStatus(c *gin.Context) {
	c.Request.ParseForm()
	var longurl string
	longurl = c.Request.FormValue("longurl")
	if longurl != "" {
		urlid := models.Hash(longurl)
		var s = strconv.FormatUint(uint64(urlid), 10)
		shorturl := "http://" + os.Getenv("DBHOST") + ":" + os.Getenv("PORT") + "/status/" + s
		models.InsertUrl(longurl, shorturl, urlid)
	}
	c.HTML(
		http.StatusOK,
		"home.html",
		gin.H{
			"title": "Tiny-Url",
		},
	)
}

func OriginalPage(c *gin.Context) {
	urlID := c.Param("id")
	originalUrl := models.GetOriginalUrl(urlID)
	c.Redirect(301, originalUrl)
}

func GetUrlDetail(c *gin.Context) {
	ID := c.Param("id")
	singleurldetails := models.SingleUrlDetail(ID)
	c.JSON(http.StatusOK, singleurldetails)
}

func GetUrls(c *gin.Context) {
	allurls := models.AllUrlDetails()
	c.JSON(http.StatusOK, allurls)
}
