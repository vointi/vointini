package locales

import (
	"embed"
	"fmt"
	"github.com/vointini/vointini/pkg/locales"
	"golang.org/x/text/message/catalog"
)

//go:embed */messages.json
var localeFiles embed.FS

var Translations catalog.Catalog

// init loads translations
func init() {
	var err error
	Translations, err = locales.LoadLang(localeFiles)
	if err != nil {
		panic(fmt.Sprintf(`could not load translations: %v`, err))
	}
}
