package serviceapi

import (
	"github.com/vointi/vointini/backend/serviceapi/locales"
	"github.com/vointi/vointini/backend/storage/storageimpl"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"time"
)

// Service interacts with backend storage to save and fetch data transparently
type Service struct {
	storage storageimpl.Storage // Backend storage for permanent data
	tr      *message.Printer    // Translations
}

// New returns *Service which uses given backend storage
// Note: Service handles all possible caching
func New(storage storageimpl.Storage, defaultLanguage language.Tag) (svc *Service) {
	svc = &Service{
		storage: storage,
		tr: message.NewPrinter(defaultLanguage,
			message.Catalog(locales.Translations),
		),
	}

	// Start generating re-occurring tasks
	svc.generateReoccurringTasks(time.Minute * 1)

	return svc
}
