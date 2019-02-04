package ergo

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/text/language"
	"net/http"
)

type HumanBridge struct {
	engine *gin.Engine
	ergo   Ergo
}

func NewHumanBridge(engine *gin.Engine, ergo *Ergo) *HumanBridge {
	return &HumanBridge{
		engine: engine,
		ergo:   *ergo,
	}
}

func (human *HumanBridge) LaunchServerForHumans(address string) error {
	human.engine.GET("/what-is", func(c *gin.Context) {
		errorID := c.Query("id")
		lang := c.QueryArray("lang")

		var langs []language.Tag
		for _, l := range lang {
			tag, err := language.Parse(l)
			if err != nil {
				tag = language.English
			}
			langs = append(langs, tag)
		}

		forHuman, err := human.ergo.ConsultErrorAsHuman([]byte(errorID), langs...)
		if err != nil {
			c.JSON(http.StatusOK, ergoIsNotWorkingForHumans)
		}

		c.JSON(http.StatusOK, forHuman)

	})

	err := human.engine.Run(address)

	return err
}
