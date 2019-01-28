package ergo

import (
	"github.com/dgraph-io/badger"
	"github.com/gin-gonic/gin"
	"github.com/oklog/ulid"
	"golang.org/x/text/language"
	"net/http"
)

type ServerClient struct {
	ergo   Service
	engine *gin.Engine
}

type ServerClientOptions struct {
	// TODO: Add custom options
}

func (client *ServerClient) RegisterRoutes(opts ...ServerClientOptions) {
	client.engine.GET("/get-message", func(c *gin.Context) {
		id := c.Query("id")
		lang := c.Query("lang")

		errorID, err := ulid.Parse(id)
		switch err {
		case nil:
			break
		case ulid.ErrInvalidCharacters:
			c.String(http.StatusOK, invalidErrorCodeMessage)
			return
		default:
			c.String(http.StatusOK, invalidErrorCodeMessage)
			return
		}

		ergoError, err := client.ergo.GetErrorByID(errorID)
		switch err {
		case nil:
			break
		case badger.ErrKeyNotFound:
			c.String(http.StatusOK, errorCodeNotExistMessage)
			return
		default:
			c.String(http.StatusOK, err.Error())
			return
		}

		ergoErrorLang := DefaultLanguage

		if lang != "" {
			tag, err := language.Parse(lang)
			if err == nil {
				ergoErrorLang = tag
			}
		}

		msg, err := client.ergo.GetErrorMessageByLanguage(errorID, ergoErrorLang, true)
		if err != nil {
			c.String(http.StatusOK, "[ERGO ERROR] "+err.Error())
			return
		}

		if msg == "" {
			msg = ergoError.Message
			if msg == "" {
				msg = "x_x, ergo not found your error message | " + ergoError.Error
			}
		}

		c.String(http.StatusOK, msg)

	})
}
