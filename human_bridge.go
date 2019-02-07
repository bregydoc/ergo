package ergo

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/oklog/ulid"
	"golang.org/x/text/language"
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
		code := c.Query("code")
		lang := c.QueryArray("lang")

		var langs []language.Tag
		for _, l := range lang {
			tag, err := language.Parse(l)
			if err != nil {
				tag = language.English
			}
			langs = append(langs, tag)
		}

		// Search by ID
		id, err := ulid.Parse(errorID)
		if err != nil {
			c.JSON(http.StatusOK, ergoIsNotWorkingForHumans)
			return
		}
		// If the id is correct
		if err == nil {
			forHuman, err := human.ergo.ConsultErrorAsHumanByID(id[:], langs...)
			if err != nil {
				c.JSON(http.StatusOK, ergoIsNotWorkingForHumans)
				return
			}
			c.JSON(http.StatusOK, forHuman)
			return
		}

		// Search by code
		cc, err := strconv.Atoi(code)
		if err != nil {
			c.JSON(http.StatusOK, ergoIsNotWorkingForHumans)
			return
		}

		forHuman, err := human.ergo.ConsultErrorAsHumanByCode(uint64(cc), langs...)
		if err != nil {
			c.JSON(http.StatusOK, ergoIsNotWorkingForHumans)
			return
		}
		c.JSON(http.StatusOK, forHuman)
		return

	})

	err := human.engine.Run(address)

	return err
}
