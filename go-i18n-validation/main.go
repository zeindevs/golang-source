package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var tagPrefixMap = map[string]string{
	"required": "Required",
	"email":    "InvalidEmail",
	"min":      "ShouldMin",
	"max":      "ShouldMax",
	"len":      "ShouldLen",
	"eq":       "ShouldEq",
	"gt":       "ShouldGt",
	"gte":      "ShouldGte",
	"lt":       "ShouldLt",
	"lte":      "ShouldLte",
}

var bundle *i18n.Bundle
var localizer *i18n.Localizer

func i18nTr(msgID string, field string, params ...string) string {
	fmt.Printf("%s %s %+v\n", msgID, field, params)
	locaization, _ := localizer.Localize(&i18n.LocalizeConfig{
		MessageID: msgID,
		TemplateData: map[string]string{
			"Field": field,
			"Max":   params[0],
		},
	})
	return locaization
}

func composeMsgID(e validator.FieldError) string {
	if prefix, ok := tagPrefixMap[e.Tag()]; ok {
		return fmt.Sprintf("%s", prefix)
	}
	return ""
}

func translateError(err error) map[string]string {
	errs := make(map[string]string)
	validationErrors, ok := err.(validator.ValidationErrors)
	if ok {
		for _, e := range validationErrors {
			errs[e.Field()] = i18nTr(composeMsgID(e), e.Field(), e.Param())
		}
	} else {
		errs["Message"] = err.Error()
	}
	return errs
}

type Form struct {
	Username string `json:"username" binding:"required,max=20"`
	Email    string `json:"email" binding:"required,email,max=100"`
}

func handler(c *gin.Context) {
	form := &Form{
		// Email: "test@example.com",
	}
	if err := c.ShouldBindJSON(form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": translateError(err),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Form received successfullt",
	})
}

func main() {
	bundle := i18n.NewBundle(language.English)
	bundle.LoadMessageFile("resources/en.json")
	localizer = i18n.NewLocalizer(bundle, language.English.String(), language.Indonesian.String())

	r := gin.Default()

	r.POST("/form", handler)

	r.Run(":9001")
}
