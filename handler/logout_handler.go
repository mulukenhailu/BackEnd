package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ph *profileHandler) Logout(c *gin.Context) {
	// metadata, _ := ph.tk.ExtractedMetaData(c.Request)

	// if metadata != nil {
	// 	deleteErr := ph.rd.DeleteTokens(metadata)
	// 	if deleteErr != nil {
	// 		c.JSON(http.StatusBadRequest, deleteErr.Error())
	// 		return
	// 	}
	// }
	c.JSON(http.StatusOK, "Successfully log out!")
}
