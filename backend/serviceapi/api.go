package serviceapi

import (
	"context"
	"github.com/vointi/vointini/backend/filestorage"
	"github.com/vointi/vointini/backend/serviceapi/locales"
	"github.com/vointi/vointini/backend/storage/storageimpl"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"io"
	"log"
	"time"
)

// Service interacts with backend storage to save and fetch data transparently
type Service struct {
	storage     storageimpl.Storage // Backend storage for permanent data
	tr          *message.Printer    // Translations
	log         *log.Logger
	fileStorage *filestorage.FileStorage
}

// New returns *Service which uses given backend storage
// Note: Service handles all possible caching
func New(storage storageimpl.Storage, fStorage *filestorage.FileStorage, defaultLanguage language.Tag, logStream io.Writer) (svc *Service) {
	svc = &Service{
		storage:     storage,
		fileStorage: fStorage,
		tr: message.NewPrinter(defaultLanguage,
			message.Catalog(locales.Translations),
		),
		log: log.New(logStream, `service`, log.LstdFlags),
	}

	err := svc.Initialize(context.TODO())
	if err != nil {
		svc.log.Printf(`could not initialize: %v`, err)
		return nil
	}

	// Start generating re-occurring tasks
	svc.generateReoccurringTasks(time.Minute * 1)

	return svc
}
