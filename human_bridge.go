package ergo

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/text/language"
)

type HumanBridge struct {
	engine *gin.Engine
	ergo   Ergo
}

func (human *HumanBridge) CreateServer(address string) {
	human.engine.GET("/what-is", func(c *gin.Context) {
		errorID := c.Query("id")
		lang := c.Query("lang")
		l, err := language.Parse(lang)
		if err != nil {
			l = language.English
		}

		forHuman, err := human.ergo.ConsultErrorAsHuman([]byte(errorID), l)
		if err != nil {

		}

		fmt.Println(forHuman)
	})
}
