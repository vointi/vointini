package locales

import (
	"embed"
	"fmt"
	"github.com/vointi/vointini/pkg/locales"
	"golang.org/x/text/message/catalog"
)

//go:embed */messages.json
var localeFiles embed.FS

var Translations catalog.Catalog

// init loads translations from embedded file(s)
func init() {
	var err error
	Translations, err = locales.LoadLang(localeFiles)
	if err != nil {
		panic(fmt.Sprintf(`could not load translations: %v`, err))
	}
}
