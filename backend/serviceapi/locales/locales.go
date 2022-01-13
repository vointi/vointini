package locales

import (
	"embed"
	"encoding/json"
	"fmt"
	"golang.org/x/text/language"
	"golang.org/x/text/message/catalog"
	"path"
)

// See: https://pkg.go.dev/golang.org/x/text/message#hdr-Translation_Pipeline

//go:embed */*.json
var localeFiles embed.FS

var Translations = catalog.NewBuilder(
	catalog.Fallback(language.English),
)

// init loads translations
func init() {
	de, err := localeFiles.ReadDir(path.Join(`.`))
	if err != nil {
		panic(err)
	}

	for _, e := range de {
		if !e.IsDir() {
			continue
		}

		err = addLang(e.Name())
		if err != nil {
			panic(err)
		}
	}

}

type message struct {
	Id                string `json:"id"`
	Message           string `json:"message"`
	Translation       string `json:"translation"`
	TranslatorComment string `json:"translatorComment"`
	Fuzzy             bool   `json:"fuzzy"`
}

type base struct {
	Language string    `json:"language"`
	Messages []message `json:"messages"`
}

func addLang(name string) error {
	b, err := localeFiles.ReadFile(path.Join(name, `out.gotext.json`))
	if err != nil {
		return err
	}

	var translations base
	err = json.Unmarshal(b, &translations)
	if err != nil {
		return err
	}
	b = nil

	lang := language.English
	switch translations.Language {
	case `en-US`:
	case `fi-FI`:
		lang = language.Finnish
	default:
		return fmt.Errorf(`invalid language: %q`, name)
	}

	for _, i := range translations.Messages {
		err = Translations.SetString(lang, i.Id, i.Translation)
		if err != nil {
			return err
		}
	}

	return nil
}
