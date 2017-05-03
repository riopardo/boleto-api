package api

import (
	"errors"

	"bitbucket.org/mundipagg/boletoapi/config"
	"bitbucket.org/mundipagg/boletoapi/log"
	"bitbucket.org/mundipagg/boletoapi/models"
	gin "gopkg.in/gin-gonic/gin.v1"
)

// ReturnHeaders 'seta' os headers padrões de resposta
func ReturnHeaders() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		c.Next()
	}
}

func executionController() gin.HandlerFunc {
	return func(c *gin.Context) {
		if config.IsRunning() {
			c.AbortWithError(500, errors.New("A aplicação está sendo finalizada"))
			return
		}
	}
}

//ParseBoleto trata a entrada de boleto em todos os requests
func ParseBoleto() gin.HandlerFunc {
	return func(c *gin.Context) {
		boleto := models.BoletoRequest{}
		errBind := c.BindJSON(&boleto)
		if errBind != nil {
			e := models.NewFormatError(errBind.Error())
			checkError(c, e, log.CreateLog())
			return
		}
		c.Set("boleto", boleto)
		c.Next()
	}
}
