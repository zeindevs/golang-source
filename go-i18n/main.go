package main

import (
	"encoding/json"
	"fmt"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var localizer *i18n.Localizer
var bundle *i18n.Bundle

func main() {
	bundle = i18n.NewBundle(language.English)

	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	bundle.LoadMessageFile("resources/en.json")
	bundle.LoadMessageFile("resources/id.json")

	localizer := i18n.NewLocalizer(bundle, language.Indonesian.String(), language.English.String())

	localizeConfigWelcome := i18n.LocalizeConfig{
		MessageID: "greeting",
		TemplateData: map[string]string{
			"Name": "Zen",
		},
	}

	localizationUsingJson, _ := localizer.Localize(&localizeConfigWelcome)

	fmt.Println(localizationUsingJson)
}
