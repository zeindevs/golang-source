package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

func TestUseMessage(t *testing.T) {
	messageEn := i18n.Message{
		ID:    "hello",
		Other: "Hello!",
	}
	messageId := i18n.Message{
		ID:    "hello",
		Other: "Hallo!",
	}

	bundle := i18n.NewBundle(language.English)
	bundle.AddMessages(language.English, &messageEn)
	bundle.AddMessages(language.Indonesian, &messageId)

	localizer := i18n.NewLocalizer(bundle, language.Indonesian.String(), language.English.String())
	localizeConfig := i18n.LocalizeConfig{
		MessageID: "hello",
	}
	localization, _ := localizer.Localize(&localizeConfig)

	fmt.Println(localization)
}

func TestUseDefaultMessage(t *testing.T) {
	defaultMessageEn := i18n.Message{
		ID:    "welcome",
		Other: "Welcome to my app!",
	}
	bundle := i18n.NewBundle(language.English)
	localizer := i18n.NewLocalizer(bundle, language.Indonesian.String(), language.English.String())

	localizeConfigWithDefault := i18n.LocalizeConfig{
		MessageID:      "welcome",
		DefaultMessage: &defaultMessageEn,
	}
	localizationReturningDefault, _ := localizer.Localize(&localizeConfigWithDefault)

	fmt.Println(localizationReturningDefault)
}

func TestUsingJsonFile(t *testing.T) {
	bundle := i18n.NewBundle(language.English)

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

func TestUseHttpRequest(t *testing.T) {
	bundle := i18n.NewBundle(language.English)

	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	bundle.LoadMessageFile("resources/en.json")
	bundle.LoadMessageFile("resources/id.json")

	handler := func(w http.ResponseWriter, r *http.Request) {
		lang := r.FormValue("lang")
		accept := r.Header.Get("Accept-Language")
		localizer := i18n.NewLocalizer(bundle, lang, accept)

		valToLocalize := r.URL.Query().Get("msg")
		name := r.URL.Query().Get("name")

		localizeConfig := i18n.LocalizeConfig{
			MessageID: valToLocalize,
			TemplateData: map[string]string{
				"Name": name,
			},
		}
		locatization, _ := localizer.Localize(&localizeConfig)
		fmt.Fprintln(w, locatization)
	}

	req := httptest.NewRequest(http.MethodGet, "/?lang=en&msg=greeting&name=Zen", nil)
	w := httptest.NewRecorder()
	handler(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, _ := io.ReadAll(res.Body)

	fmt.Println(string(data))
}
