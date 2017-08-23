package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/syedomair/weather/app/common"
)

func GetAction(c *gin.Context) {
	/*
		dbInterface, _ := c.Get("DB")
		db := dbInterface.(models.Datastore)
		bks, _ := db.AllBooks()
		fmt.Println(bks[0].Id)

		c.JSON(http.StatusOK, common.SuccessResponse(bks[0].Id))
	*/
	c.JSON(http.StatusOK, common.SuccessResponse("success"))
	return
}
