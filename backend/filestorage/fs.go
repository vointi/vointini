package filestorage

import (
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	"github.com/spf13/afero"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"io"
	"os"
	"strings"
)

type FileStorage struct {
	vfs afero.Fs
	tr  *message.Printer
}

func New(basepath string, defaultLanguage language.Tag) *FileStorage {
	if basepath == `` {
		panic(`empty basepath`)
	}

	return &FileStorage{
		vfs: afero.NewBasePathFs(afero.NewOsFs(), basepath),
		tr:  message.NewPrinter(defaultLanguage),
	}
}

// upload saves a stream to a targetDir with its contents hash as a filename
func (fs FileStorage) upload(ctx context.Context, extension string, r io.ReadCloser, targetDir string) (newname string, err error) {
	if strings.HasPrefix(extension, `.`) {
		return ``, fmt.Errorf(`extension contains dot ('.') prefix`)
	}

	tmpf, err := afero.TempFile(fs.vfs, `/`, `upload-**********.`+extension)
	if err != nil {
		return newname, err
	}
	defer tmpf.Close()

	newname = tmpf.Name()

	h := sha256.New()

	buffer := make([]byte, 8192)

	for {
		readBytes, err := r.Read(buffer)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			return newname, err
		}

		_, err = h.Write(buffer[:readBytes])
		if err != nil {
			return newname, err
		}

		_, err = tmpf.Write(buffer[:readBytes])
		if err != nil {
			return newname, err
		}

	}

	err = tmpf.Sync()
	if err != nil {
		return newname, err
	}

	newname = fmt.Sprintf(`%x.%s`, h.Sum(nil), extension)
	renameTo := fmt.Sprintf(`%s/%s`, targetDir, newname)

	// Check if file exists already
	_, err = fs.vfs.Stat(renameTo)
	if err != nil {
		if errors.Is(err, os.ErrExist) {
			// File exists already, nothing to do here
			return newname, nil
		}

		if !errors.Is(err, os.ErrNotExist) {
			return newname, fmt.Errorf(`couldn't stat %q: %w`, renameTo, err)
		}
	}

	// Create directory
	err = fs.vfs.MkdirAll(targetDir, os.ModeDir)
	if err != nil {
		if !errors.Is(err, os.ErrExist) {
			return newname, fmt.Errorf(`can't create directory %q: %w`, targetDir, err)
		}
	}

	// Move file
	err = fs.vfs.Rename(tmpf.Name(), renameTo)
	if err != nil {
		return newname, fmt.Errorf(fs.tr.Sprintf(`can't move file %q to %q`, tmpf.Name(), renameTo)+`: %w`, err)
	}

	return newname, nil
}

func (fs FileStorage) AddResolutionFile(ctx context.Context, id int, extension string, r io.ReadCloser) (string, error) {
	fname, err := fs.upload(ctx, extension, r, fmt.Sprintf(`/resolutions/%d/`, id))
	if err != nil {
		return ``, err
	}

	return fname, nil
}
