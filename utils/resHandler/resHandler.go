package resHandler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrResponse struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
}

func SendResponse(c *gin.Context, code int, err error, msg interface{}) {
	if code == http.StatusOK {
		c.JSON(code, msg)
		return
	}

	// TODO: logger
	fmt.Println(err.Error())
	errRes := ErrResponse{
		Status: code,
		Msg:    err.Error(),
	}
	c.JSON(code, errRes)
}

func SendRedirect(c *gin.Context, url string) {
	c.Redirect(http.StatusFound, url)
}
