package main

import (
	"encoding/json"
	"fmt"
	"github.com/vointini/vointini/pkg/locales"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
)

func updateTranslationsInDir(basepath string) {
	de, err := os.ReadDir(basepath)
	if err != nil {
		panic(err)
	}

	for _, dir := range de {
		if !dir.IsDir() {
			continue
		}

		err = updateLocales(path.Join(basepath, dir.Name()))
		if err != nil {
			log.Fatalf(`error: %v`, err)
		}
	}
}

func updateLocales(fpath string) error {
	log.Printf(`Reading %q`, fpath)

	base, err := loadLocale(path.Join(fpath, `..`, `en-US`, `out.gotext.json`))
	if err != nil {
		return err
	}

	translated, err := loadLocale(path.Join(fpath, `messages.json`))
	if err != nil {
		return err
	}

	var idOrder []string
	mergeHelper := make(map[string]locales.Msg)

	var baseIds []string

	for _, msg := range base.Messages {
		idOrder = append(idOrder, msg.Id)
		baseIds = append(baseIds, msg.Id)
		mergeHelper[msg.Id] = msg
	}

	var translatedIds []string

	for _, msg := range translated.Messages {
		translatedIds = append(translatedIds, msg.Id)

		if _, ok := mergeHelper[msg.Id]; ok {
			mergeHelper[msg.Id] = msg
		}
	}

	merged := locales.Locale{
		Language: translated.Language,
	}

	for _, v := range idOrder {
		merged.Messages = append(merged.Messages, mergeHelper[v])
	}

	b, err := json.MarshalIndent(merged, ``, "\t")
	if err != nil {
		return err
	}

	mf, err := ioutil.TempFile(`/tmp`, `merged-*.json`)
	if err != nil {
		return err
	}
	_, err = mf.Write(b)
	if err != nil {
		return err
	}
	err = mf.Close()
	if err != nil {
		return err
	}

	log.Printf(`Wrote %q`, mf.Name())

	err = copyFile(mf.Name(), path.Join(fpath, `messages.json`))
	if err != nil {
		return err
	}

	log.Printf(`Copied %q to %q`, mf.Name(), path.Join(fpath, `messages.json`))

	return nil
}

func loadLocale(fpath string) (l locales.Locale, err error) {
	b, err := os.ReadFile(fpath)
	if err != nil {
		return locales.Locale{}, err
	}

	err = json.Unmarshal(b, &l)
	if err != nil {
		return locales.Locale{}, err
	}

	return l, nil
}

func copyFile(src string, dst string) (err error) {
	fStat, err := os.Stat(src)
	if err != nil {
		return err
	}

	if !fStat.Mode().IsRegular() {
		return fmt.Errorf("%q is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destination.Close()

	_, err = io.Copy(destination, source)

	return err
}
