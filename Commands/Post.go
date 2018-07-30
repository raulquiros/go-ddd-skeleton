package Commands

import (
	"io/ioutil"
	"io"
	"net/http"
	"github.com/gin-gonic/gin"
	"go-ddd-structure/Application/Service"
)

func Post(c *gin.Context) {

	body, err := ioutil.ReadAll(io.LimitReader(c.Request.Body, 1048576))
	if err != nil {
		ReturnErr(c.Writer, err)
	}
	if err := c.Request.Body.Close(); err != nil {
		ReturnErr(c.Writer, err)
	}

	var container DependencyContainer
	_ = GetContainer(&container)

	valid, err := Service.Post()

	if err != nil {
		ReturnErr(c.Writer, err)
	}

	response := Response{valid, "Ok"}
	MakeResponse(c.Writer, response, http.StatusOK)
}

