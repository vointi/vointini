package locales

// See: https://pkg.go.dev/golang.org/x/text/message#hdr-Translation_Pipeline

import (
	"embed"
	"encoding/json"
	"fmt"
	"golang.org/x/text/language"
	"golang.org/x/text/message/catalog"
	"path"
)

type Msg struct {
	Id                string `json:"id"`
	Message           string `json:"message"`
	Translation       string `json:"translation"`
	TranslatorComment string `json:"translatorComment"`
}

type Locale struct {
	Language string `json:"language"`
	Messages []Msg  `json:"messages"`
}

func LoadLang(src embed.FS) (cat catalog.Catalog, err error) {
	languages := []string{
		`en-US`, `fi-FI`,
	}

	tr := catalog.NewBuilder(
		catalog.Fallback(language.English),
	)

	for _, name := range languages {
		b, err := src.ReadFile(path.Join(name, `messages.json`))
		if err != nil {
			return nil, err
		}

		var translations Locale
		err = json.Unmarshal(b, &translations)
		if err != nil {
			return nil, err
		}
		b = nil

		lang := language.English
		switch translations.Language {
		case `en-US`:
		case `fi-FI`:
			lang = language.Finnish
		default:
			return nil, fmt.Errorf(`invalid language: %q`, name)
		}

		for _, i := range translations.Messages {
			err = tr.SetString(lang, i.Id, i.Translation)
			if err != nil {
				return nil, fmt.Errorf(`could not set string: %v`, err)
			}
		}
	}

	return tr, nil
}
