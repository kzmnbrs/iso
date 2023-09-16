// Code generated by qtc from "country.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line templates/country.qtpl:1
package templates

//line templates/country.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line templates/country.qtpl:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line templates/country.qtpl:2
type CountryOptions struct {
	Entries []struct {
		Alpha2       string
		Alpha3       string
		CurrencyCode string
		Name         string
		LongName     string
		PhonePrefix  string
		PostalCodeRe string
	}
}

//line templates/country.qtpl:15
func StreamCountry(qw422016 *qt422016.Writer, opts *CountryOptions) {
//line templates/country.qtpl:15
	qw422016.N().S(`
// Code generated from iso/gen. DO NOT EDIT.

package iso

import (
    "database/sql/driver"
    "errors"
    "regexp"
)

type (
    Country uint8
    NullCountry struct {
        Valid bool
        Country Country
    }
)

const (
    _ Country = iota
`)
//line templates/country.qtpl:36
	for _, ent := range opts.Entries {
//line templates/country.qtpl:36
		qw422016.N().S(`    Country`)
//line templates/country.qtpl:37
		qw422016.N().S(ent.Alpha2)
//line templates/country.qtpl:37
		qw422016.N().S(`
`)
//line templates/country.qtpl:38
	}
//line templates/country.qtpl:38
	qw422016.N().S(`)

const (
`)
//line templates/country.qtpl:42
	for _, ent := range opts.Entries {
//line templates/country.qtpl:42
		qw422016.N().S(`    Country`)
//line templates/country.qtpl:43
		qw422016.N().S(ent.Alpha3)
//line templates/country.qtpl:43
		qw422016.N().S(` = Country`)
//line templates/country.qtpl:43
		qw422016.N().S(ent.Alpha2)
//line templates/country.qtpl:43
		qw422016.N().S(`
`)
//line templates/country.qtpl:44
	}
//line templates/country.qtpl:44
	qw422016.N().S(`)

var ErrBadCountry = errors.New("bad country")

var countryToAlpha2 = map[Country]string{
`)
//line templates/country.qtpl:50
	for _, ent := range opts.Entries {
//line templates/country.qtpl:50
		qw422016.N().S(`    Country`)
//line templates/country.qtpl:51
		qw422016.N().S(ent.Alpha2)
//line templates/country.qtpl:51
		qw422016.N().S(`: "`)
//line templates/country.qtpl:51
		qw422016.N().S(ent.Alpha2)
//line templates/country.qtpl:51
		qw422016.N().S(`",
`)
//line templates/country.qtpl:52
	}
//line templates/country.qtpl:52
	qw422016.N().S(`}

var countryToAlpha3 = map[Country]string{
`)
//line templates/country.qtpl:56
	for _, ent := range opts.Entries {
//line templates/country.qtpl:56
		qw422016.N().S(`    Country`)
//line templates/country.qtpl:57
		qw422016.N().S(ent.Alpha3)
//line templates/country.qtpl:57
		qw422016.N().S(`: "`)
//line templates/country.qtpl:57
		qw422016.N().S(ent.Alpha3)
//line templates/country.qtpl:57
		qw422016.N().S(`",
`)
//line templates/country.qtpl:58
	}
//line templates/country.qtpl:58
	qw422016.N().S(`}

var countryToCurrency = map[Country]Currency{
`)
//line templates/country.qtpl:62
	for _, ent := range opts.Entries {
//line templates/country.qtpl:62
		qw422016.N().S(`    Country`)
//line templates/country.qtpl:63
		qw422016.N().S(ent.Alpha2)
//line templates/country.qtpl:63
		qw422016.N().S(`: Currency`)
//line templates/country.qtpl:63
		qw422016.N().S(ent.CurrencyCode)
//line templates/country.qtpl:63
		qw422016.N().S(`,
`)
//line templates/country.qtpl:64
	}
//line templates/country.qtpl:64
	qw422016.N().S(`}

var alpha2ToCountry = map[string]Country{
`)
//line templates/country.qtpl:68
	for _, ent := range opts.Entries {
//line templates/country.qtpl:68
		qw422016.N().S(`   "`)
//line templates/country.qtpl:69
		qw422016.N().S(ent.Alpha2)
//line templates/country.qtpl:69
		qw422016.N().S(`": Country`)
//line templates/country.qtpl:69
		qw422016.N().S(ent.Alpha2)
//line templates/country.qtpl:69
		qw422016.N().S(`,
`)
//line templates/country.qtpl:70
	}
//line templates/country.qtpl:70
	qw422016.N().S(`}

var alpha3ToCountry = map[string]Country{
`)
//line templates/country.qtpl:74
	for _, ent := range opts.Entries {
//line templates/country.qtpl:74
		qw422016.N().S(`   "`)
//line templates/country.qtpl:75
		qw422016.N().S(ent.Alpha3)
//line templates/country.qtpl:75
		qw422016.N().S(`": Country`)
//line templates/country.qtpl:75
		qw422016.N().S(ent.Alpha3)
//line templates/country.qtpl:75
		qw422016.N().S(`,
`)
//line templates/country.qtpl:76
	}
//line templates/country.qtpl:76
	qw422016.N().S(`}

var countryToName = map[Country]string{
`)
//line templates/country.qtpl:80
	for _, ent := range opts.Entries {
//line templates/country.qtpl:80
		qw422016.N().S(`   Country`)
//line templates/country.qtpl:81
		qw422016.N().S(ent.Alpha2)
//line templates/country.qtpl:81
		qw422016.N().S(`: "`)
//line templates/country.qtpl:81
		qw422016.N().S(ent.Name)
//line templates/country.qtpl:81
		qw422016.N().S(`",
`)
//line templates/country.qtpl:82
	}
//line templates/country.qtpl:82
	qw422016.N().S(`}

var countryToLongName = map[Country]string{
`)
//line templates/country.qtpl:86
	for _, ent := range opts.Entries {
//line templates/country.qtpl:86
		qw422016.N().S(`   Country`)
//line templates/country.qtpl:87
		qw422016.N().S(ent.Alpha2)
//line templates/country.qtpl:87
		qw422016.N().S(`: "`)
//line templates/country.qtpl:87
		qw422016.N().S(ent.LongName)
//line templates/country.qtpl:87
		qw422016.N().S(`",
`)
//line templates/country.qtpl:88
	}
//line templates/country.qtpl:88
	qw422016.N().S(`}

var countryToPhonePrefix = map[Country]string{
`)
//line templates/country.qtpl:92
	for _, ent := range opts.Entries {
//line templates/country.qtpl:92
		qw422016.N().S(`   Country`)
//line templates/country.qtpl:93
		qw422016.N().S(ent.Alpha2)
//line templates/country.qtpl:93
		qw422016.N().S(`: "`)
//line templates/country.qtpl:93
		qw422016.N().S(ent.PhonePrefix)
//line templates/country.qtpl:93
		qw422016.N().S(`",
`)
//line templates/country.qtpl:94
	}
//line templates/country.qtpl:94
	qw422016.N().S(`}

var countryToPostalCodeRe map[Country]*regexp.Regexp

func init() {
    countryToPostalCodeRe = make(map[Country]*regexp.Regexp, len(alpha2ToCountry))
`)
//line templates/country.qtpl:101
	for _, ent := range opts.Entries {
//line templates/country.qtpl:101
		qw422016.N().S(`    `)
//line templates/country.qtpl:102
		if len(ent.PostalCodeRe) > 0 {
//line templates/country.qtpl:102
			qw422016.N().S(`
    {
        re := regexp.MustCompile(`)
//line templates/country.qtpl:102
			qw422016.N().S("`")
//line templates/country.qtpl:104
			qw422016.N().S(ent.PostalCodeRe)
//line templates/country.qtpl:104
			qw422016.N().S(``)
//line templates/country.qtpl:104
			qw422016.N().S("`")
//line templates/country.qtpl:104
			qw422016.N().S(`)
        countryToPostalCodeRe[Country`)
//line templates/country.qtpl:105
			qw422016.N().S(ent.Alpha2)
//line templates/country.qtpl:105
			qw422016.N().S(`] = re
    }
    `)
//line templates/country.qtpl:107
		}
//line templates/country.qtpl:107
		qw422016.N().S(`
`)
//line templates/country.qtpl:108
	}
//line templates/country.qtpl:108
	qw422016.N().S(`}

func CountryFromCode(v string) Country {
    x, found := alpha2ToCountry[v]
    if !found {
        return alpha3ToCountry[v]
    }
    return x
}

func CountryFromAlpha3(v string) Country {
    return alpha3ToCountry[v]
}

func AppendCountries(dst []Country) []Country {
`)
//line templates/country.qtpl:124
	for _, ent := range opts.Entries {
//line templates/country.qtpl:124
		qw422016.N().S(`    dst = append(dst, Country`)
//line templates/country.qtpl:125
		qw422016.N().S(ent.Alpha2)
//line templates/country.qtpl:125
		qw422016.N().S(`)
`)
//line templates/country.qtpl:126
	}
//line templates/country.qtpl:126
	qw422016.N().S(`    return dst
}

func (v Country) String() string {
	return countryToAlpha3[v]
}

func (v Country) Alpha2() string {
    return countryToAlpha2[v]
}

func (v Country) Name() string {
    return countryToName[v]
}

func (v Country) LongName() string {
    return countryToLongName[v]
}

func (v Country) PhonePrefix() string {
    return countryToPhonePrefix[v]
}

func (v Country) PostalCodeRe() *regexp.Regexp {
    return countryToPostalCodeRe[v]
}

func (v Country) Currency() Currency {
    return countryToCurrency[v]
}

func (v Country) IsValid() bool {
    return v.String() != ""
}

func (v Country) MarshalText() ([]byte, error) {
	code := v.String()
	if len(code) == 0 {
		return nil, ErrBadCountry
	}
	return s2b(code), nil
}

func (v *Country) UnmarshalText(bb []byte) error {
	v1, found := alpha2ToCountry[b2s(bb)]
	if !found {
	    v1, found = alpha3ToCountry[b2s(bb)]
	}
    if !found {
        return ErrBadCountry
    }
	*v = v1
	return nil
}

func (v *Country) UnmarshalJSON(bb []byte) error {
	if len(bb) < 2 {
		return ErrBadCountry
	}
	return v.UnmarshalText(unquote(bb))
}

func (v Country) Value() (driver.Value, error) {
	return v.String(), nil
}

func (v *Country) Scan(v1 interface{}) error {
	switch x := v1.(type) {
	case string:
		return v.UnmarshalText(s2b(x))
	case []byte:
		return v.UnmarshalText(x)
	}
	return ErrBadCountry
}

func (nv *NullCountry) UnmarshalJSON(bb []byte) error {
    if isJsonNull(bb) {
        nv.Valid = false
        return nil
    }

    v := Country(0)
    if err := v.UnmarshalJSON(bb); err != nil {
        nv.Valid = false
        return err
    }

    nv.Country = v
    nv.Valid = true
    return nil
}

func (nv *NullCountry) Scan(v1 interface{}) error {
    if v1 == nil {
        nv.Valid = false
        return nil
    }

    v := Country(0)
    if err := v.Scan(v1); err != nil {
        nv.Valid = false
        return err
    }

    nv.Country = v
    nv.Valid = true
    return nil
}
`)
//line templates/country.qtpl:236
}

//line templates/country.qtpl:236
func WriteCountry(qq422016 qtio422016.Writer, opts *CountryOptions) {
//line templates/country.qtpl:236
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/country.qtpl:236
	StreamCountry(qw422016, opts)
//line templates/country.qtpl:236
	qt422016.ReleaseWriter(qw422016)
//line templates/country.qtpl:236
}

//line templates/country.qtpl:236
func Country(opts *CountryOptions) string {
//line templates/country.qtpl:236
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/country.qtpl:236
	WriteCountry(qb422016, opts)
//line templates/country.qtpl:236
	qs422016 := string(qb422016.B)
//line templates/country.qtpl:236
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/country.qtpl:236
	return qs422016
//line templates/country.qtpl:236
}

//line templates/country.qtpl:238
func StreamCountrySQL(qw422016 *qt422016.Writer, opts *CountryOptions) {
//line templates/country.qtpl:238
	qw422016.N().S(`CREATE TYPE iso_country AS ENUM (
`)
//line templates/country.qtpl:240
	for i, ent := range opts.Entries {
//line templates/country.qtpl:240
		qw422016.N().S(`    '`)
//line templates/country.qtpl:241
		qw422016.N().S(ent.Alpha2)
//line templates/country.qtpl:241
		qw422016.N().S(`'`)
//line templates/country.qtpl:241
		if i != len(opts.Entries)-1 {
//line templates/country.qtpl:241
			qw422016.N().S(`,`)
//line templates/country.qtpl:241
		}
//line templates/country.qtpl:241
		qw422016.N().S(`
`)
//line templates/country.qtpl:242
	}
//line templates/country.qtpl:242
	qw422016.N().S(`)
`)
//line templates/country.qtpl:244
}

//line templates/country.qtpl:244
func WriteCountrySQL(qq422016 qtio422016.Writer, opts *CountryOptions) {
//line templates/country.qtpl:244
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/country.qtpl:244
	StreamCountrySQL(qw422016, opts)
//line templates/country.qtpl:244
	qt422016.ReleaseWriter(qw422016)
//line templates/country.qtpl:244
}

//line templates/country.qtpl:244
func CountrySQL(opts *CountryOptions) string {
//line templates/country.qtpl:244
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/country.qtpl:244
	WriteCountrySQL(qb422016, opts)
//line templates/country.qtpl:244
	qs422016 := string(qb422016.B)
//line templates/country.qtpl:244
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/country.qtpl:244
	return qs422016
//line templates/country.qtpl:244
}

//line templates/country.qtpl:246
func StreamCountry3SQL(qw422016 *qt422016.Writer, opts *CountryOptions) {
//line templates/country.qtpl:246
	qw422016.N().S(`CREATE TYPE iso_country3 AS ENUM (
`)
//line templates/country.qtpl:248
	for i, ent := range opts.Entries {
//line templates/country.qtpl:248
		qw422016.N().S(`    '`)
//line templates/country.qtpl:249
		qw422016.N().S(ent.Alpha3)
//line templates/country.qtpl:249
		qw422016.N().S(`'`)
//line templates/country.qtpl:249
		if i != len(opts.Entries)-1 {
//line templates/country.qtpl:249
			qw422016.N().S(`,`)
//line templates/country.qtpl:249
		}
//line templates/country.qtpl:249
		qw422016.N().S(`
`)
//line templates/country.qtpl:250
	}
//line templates/country.qtpl:250
	qw422016.N().S(`)
`)
//line templates/country.qtpl:252
}

//line templates/country.qtpl:252
func WriteCountry3SQL(qq422016 qtio422016.Writer, opts *CountryOptions) {
//line templates/country.qtpl:252
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/country.qtpl:252
	StreamCountry3SQL(qw422016, opts)
//line templates/country.qtpl:252
	qt422016.ReleaseWriter(qw422016)
//line templates/country.qtpl:252
}

//line templates/country.qtpl:252
func Country3SQL(opts *CountryOptions) string {
//line templates/country.qtpl:252
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/country.qtpl:252
	WriteCountry3SQL(qb422016, opts)
//line templates/country.qtpl:252
	qs422016 := string(qb422016.B)
//line templates/country.qtpl:252
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/country.qtpl:252
	return qs422016
//line templates/country.qtpl:252
}
