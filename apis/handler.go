package apis

import (
	"fmt"
	"log"
	"net/http"

	. "gin-gonic/models"

	"github.com/gin-gonic/gin"
)

func IndexApi(c *gin.Context) {
	c.String(http.StatusOK, "It works")
}

func AddPersonApi(c *gin.Context) {
	first_name := c.Request.FormValue("first_name")
	last_name := c.Request.FormValue("last_name")

	p := Person{FirstName: first_name, LastName: last_name}

	ra, err := p.AddPerson()
	if err != nil {
		log.Fatalln(err)
	}
	msg := fmt.Sprintf("insert successful %d", ra)
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}
