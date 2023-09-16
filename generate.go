package main

import (
	"bytes"
	"fmt"
	"go/format"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/kzmnbrs/iso/templates"
	"gopkg.in/yaml.v3"
)

//go:generate go get -u github.com/valyala/quicktemplate/qtc
//go:generate qtc -dir=templates

type Options struct {
	Currency templates.CurrencyOptions
	Country  templates.CountryOptions
}

func main() {
	opts, err := collect()
	if err != nil {
		log.Fatal(err)
	}

	err = writeAndFormat("currency", func(w io.Writer) {
		templates.WriteCurrency(w, &opts.Currency)
	})
	if err != nil {
		log.Fatal(err)
	}
	err = writeSQL("currency", func(w io.Writer) {
		templates.WriteCurrencySQL(w, &opts.Currency)
	})
	if err != nil {
		log.Fatal(err)
	}

	err = writeAndFormat("country", func(w io.Writer) {
		templates.WriteCountry(w, &opts.Country)
	})
	if err != nil {
		log.Fatal(err)
	}
	err = writeSQL("country", func(w io.Writer) {
		templates.WriteCountrySQL(w, &opts.Country)
	})
	if err != nil {
		log.Fatal(err)
	}
	err = writeSQL("country3", func(w io.Writer) {
		templates.WriteCountry3SQL(w, &opts.Country)
	})
	if err != nil {
		log.Fatal(err)
	}
}

func collect() (*Options, error) {
	dir, err := os.MkdirTemp("", "ruby_iso_countries*")
	defer func() { _ = os.RemoveAll(dir) }()

	_, err = git.PlainClone(dir, false, &git.CloneOptions{
		URL:          "https://github.com/countries/countries",
		Progress:     os.Stdout,
		SingleBranch: true,
	})
	if err != nil {
		return nil, fmt.Errorf(")) %w", err)
	}

	dataDir := filepath.Join(dir, "lib/countries/data/countries")

	opts := &Options{}
	err = filepath.WalkDir(dataDir, func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}

		var data = map[string]struct {
			Alpha2       string `yaml:"alpha2"`
			Alpha3       string `yaml:"alpha3"`
			CurrencyCode string `yaml:"currency_code"`
			PhonePrefix  string `yaml:"country_code"`
			Name         string `yaml:"iso_short_name"`
			LongName     string `yaml:"iso_long_name"`
			PostalCodeRe string `yaml:"postal_code_format"`
		}{}

		bb, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		if err := yaml.Unmarshal(bb, &data); err != nil {
			return err
		}

		if len(data) != 1 {
			return strconv.ErrSyntax
		}
		for key := range data {
			opts.Currency.Codes = appendOnce(opts.Currency.Codes, data[key].CurrencyCode)
			opts.Country.Entries = appendOnce(opts.Country.Entries, struct {
				Alpha2       string
				Alpha3       string
				CurrencyCode string
				Name         string
				LongName     string
				PhonePrefix  string
				PostalCodeRe string
			}{data[key].Alpha2, data[key].Alpha3, data[key].CurrencyCode,
				data[key].Name, data[key].LongName, data[key].PhonePrefix,
				data[key].PostalCodeRe})
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return opts, nil
}

func writeAndFormat(name string, render func(io.Writer)) error {
	buf := bytes.Buffer{}
	render(&buf)

	fname := filepath.Join("go", strings.ToLower(name)+"_enum.go")
	fd, err := os.OpenFile(
		fname,
		os.O_WRONLY|os.O_CREATE|os.O_TRUNC,
		0644,
	)
	if err != nil {
		return err
	}
	defer func() { _ = fd.Close() }()

	bs, err := format.Source(buf.Bytes())
	if err != nil {
		return err
	}
	_, err = fd.Write(bs)
	return err
}

func writeSQL(name string, render func(io.Writer)) error {
	buf := bytes.Buffer{}
	render(&buf)

	fname := "sql/" + strings.ToLower(name) + ".sql"
	fd, err := os.OpenFile(
		fname,
		os.O_WRONLY|os.O_CREATE|os.O_TRUNC,
		0644,
	)
	if err != nil {
		return err
	}
	defer func() { _ = fd.Close() }()

	_, err = fd.Write(buf.Bytes())
	return err
}

func appendOnce[T comparable](dst []T, x T) []T {
	for i := range dst {
		if dst[i] == x {
			return dst
		}
	}
	return append(dst, x)
}
