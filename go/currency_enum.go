// Code generated from iso/gen. DO NOT EDIT.

package iso

import (
	"database/sql/driver"
	"errors"
)

type (
	Currency     uint8
	NullCurrency struct {
		Valid    bool
		Currency Currency
	}
)

const (
	_ Currency = iota
	CurrencyEUR
	CurrencyAED
	CurrencyAFN
	CurrencyXCD
	CurrencyALL
	CurrencyAMD
	CurrencyAOA
	CurrencyUSD
	CurrencyARS
	CurrencyAUD
	CurrencyAWG
	CurrencyAZN
	CurrencyBAM
	CurrencyBBD
	CurrencyBDT
	CurrencyXOF
	CurrencyBGN
	CurrencyBHD
	CurrencyBIF
	CurrencyBMD
	CurrencyBND
	CurrencyBOB
	CurrencyBRL
	CurrencyBSD
	CurrencyBTN
	CurrencyNOK
	CurrencyBWP
	CurrencyBYN
	CurrencyBZD
	CurrencyCAD
	CurrencyCDF
	CurrencyXAF
	CurrencyCHF
	CurrencyNZD
	CurrencyCLP
	CurrencyCNY
	CurrencyCOP
	CurrencyCRC
	CurrencyCUP
	CurrencyCVE
	CurrencyANG
	CurrencyCZK
	CurrencyDJF
	CurrencyDKK
	CurrencyDOP
	CurrencyDZD
	CurrencyEGP
	CurrencyMAD
	CurrencyETB
	CurrencyFJD
	CurrencyFKP
	CurrencyGBP
	CurrencyGEL
	CurrencyGHS
	CurrencyGIP
	CurrencyGMD
	CurrencyGNF
	CurrencyGTQ
	CurrencyGYD
	CurrencyHKD
	CurrencyHNL
	CurrencyHTG
	CurrencyHUF
	CurrencyIDR
	CurrencyILS
	CurrencyINR
	CurrencyIQD
	CurrencyIRR
	CurrencyISK
	CurrencyJMD
	CurrencyJOD
	CurrencyJPY
	CurrencyKES
	CurrencyKGS
	CurrencyKHR
	CurrencyKMF
	CurrencyKPW
	CurrencyKRW
	CurrencyKWD
	CurrencyKYD
	CurrencyKZT
	CurrencyLAK
	CurrencyLBP
	CurrencyLKR
	CurrencyLRD
	CurrencyLSL
	CurrencyLYD
	CurrencyMDL
	CurrencyMGA
	CurrencyMKD
	CurrencyMMK
	CurrencyMNT
	CurrencyMOP
	CurrencyMRU
	CurrencyMUR
	CurrencyMVR
	CurrencyMWK
	CurrencyMXN
	CurrencyMYR
	CurrencyMZN
	CurrencyNAD
	CurrencyXPF
	CurrencyNGN
	CurrencyNIO
	CurrencyNPR
	CurrencyOMR
	CurrencyPAB
	CurrencyPEN
	CurrencyPGK
	CurrencyPHP
	CurrencyPKR
	CurrencyPLN
	CurrencyPYG
	CurrencyQAR
	CurrencyRON
	CurrencyRSD
	CurrencyRUB
	CurrencyRWF
	CurrencySAR
	CurrencySBD
	CurrencySCR
	CurrencySDG
	CurrencySEK
	CurrencySGD
	CurrencySHP
	CurrencySLL
	CurrencySOS
	CurrencySRD
	CurrencySSP
	CurrencySTD
	CurrencySYP
	CurrencySZL
	CurrencyTHB
	CurrencyTJS
	CurrencyTMT
	CurrencyTND
	CurrencyTOP
	CurrencyTRY
	CurrencyTTD
	CurrencyTWD
	CurrencyTZS
	CurrencyUAH
	CurrencyUGX
	CurrencyUYU
	CurrencyUZS
	CurrencyVES
	CurrencyVND
	CurrencyVUV
	CurrencyWST
	CurrencyYER
	CurrencyZAR
	CurrencyZMW
)

var ErrBadCurrency = errors.New("bad currency")

var currencyToCode = map[Currency]string{
	CurrencyEUR: "EUR",
	CurrencyAED: "AED",
	CurrencyAFN: "AFN",
	CurrencyXCD: "XCD",
	CurrencyALL: "ALL",
	CurrencyAMD: "AMD",
	CurrencyAOA: "AOA",
	CurrencyUSD: "USD",
	CurrencyARS: "ARS",
	CurrencyAUD: "AUD",
	CurrencyAWG: "AWG",
	CurrencyAZN: "AZN",
	CurrencyBAM: "BAM",
	CurrencyBBD: "BBD",
	CurrencyBDT: "BDT",
	CurrencyXOF: "XOF",
	CurrencyBGN: "BGN",
	CurrencyBHD: "BHD",
	CurrencyBIF: "BIF",
	CurrencyBMD: "BMD",
	CurrencyBND: "BND",
	CurrencyBOB: "BOB",
	CurrencyBRL: "BRL",
	CurrencyBSD: "BSD",
	CurrencyBTN: "BTN",
	CurrencyNOK: "NOK",
	CurrencyBWP: "BWP",
	CurrencyBYN: "BYN",
	CurrencyBZD: "BZD",
	CurrencyCAD: "CAD",
	CurrencyCDF: "CDF",
	CurrencyXAF: "XAF",
	CurrencyCHF: "CHF",
	CurrencyNZD: "NZD",
	CurrencyCLP: "CLP",
	CurrencyCNY: "CNY",
	CurrencyCOP: "COP",
	CurrencyCRC: "CRC",
	CurrencyCUP: "CUP",
	CurrencyCVE: "CVE",
	CurrencyANG: "ANG",
	CurrencyCZK: "CZK",
	CurrencyDJF: "DJF",
	CurrencyDKK: "DKK",
	CurrencyDOP: "DOP",
	CurrencyDZD: "DZD",
	CurrencyEGP: "EGP",
	CurrencyMAD: "MAD",
	CurrencyETB: "ETB",
	CurrencyFJD: "FJD",
	CurrencyFKP: "FKP",
	CurrencyGBP: "GBP",
	CurrencyGEL: "GEL",
	CurrencyGHS: "GHS",
	CurrencyGIP: "GIP",
	CurrencyGMD: "GMD",
	CurrencyGNF: "GNF",
	CurrencyGTQ: "GTQ",
	CurrencyGYD: "GYD",
	CurrencyHKD: "HKD",
	CurrencyHNL: "HNL",
	CurrencyHTG: "HTG",
	CurrencyHUF: "HUF",
	CurrencyIDR: "IDR",
	CurrencyILS: "ILS",
	CurrencyINR: "INR",
	CurrencyIQD: "IQD",
	CurrencyIRR: "IRR",
	CurrencyISK: "ISK",
	CurrencyJMD: "JMD",
	CurrencyJOD: "JOD",
	CurrencyJPY: "JPY",
	CurrencyKES: "KES",
	CurrencyKGS: "KGS",
	CurrencyKHR: "KHR",
	CurrencyKMF: "KMF",
	CurrencyKPW: "KPW",
	CurrencyKRW: "KRW",
	CurrencyKWD: "KWD",
	CurrencyKYD: "KYD",
	CurrencyKZT: "KZT",
	CurrencyLAK: "LAK",
	CurrencyLBP: "LBP",
	CurrencyLKR: "LKR",
	CurrencyLRD: "LRD",
	CurrencyLSL: "LSL",
	CurrencyLYD: "LYD",
	CurrencyMDL: "MDL",
	CurrencyMGA: "MGA",
	CurrencyMKD: "MKD",
	CurrencyMMK: "MMK",
	CurrencyMNT: "MNT",
	CurrencyMOP: "MOP",
	CurrencyMRU: "MRU",
	CurrencyMUR: "MUR",
	CurrencyMVR: "MVR",
	CurrencyMWK: "MWK",
	CurrencyMXN: "MXN",
	CurrencyMYR: "MYR",
	CurrencyMZN: "MZN",
	CurrencyNAD: "NAD",
	CurrencyXPF: "XPF",
	CurrencyNGN: "NGN",
	CurrencyNIO: "NIO",
	CurrencyNPR: "NPR",
	CurrencyOMR: "OMR",
	CurrencyPAB: "PAB",
	CurrencyPEN: "PEN",
	CurrencyPGK: "PGK",
	CurrencyPHP: "PHP",
	CurrencyPKR: "PKR",
	CurrencyPLN: "PLN",
	CurrencyPYG: "PYG",
	CurrencyQAR: "QAR",
	CurrencyRON: "RON",
	CurrencyRSD: "RSD",
	CurrencyRUB: "RUB",
	CurrencyRWF: "RWF",
	CurrencySAR: "SAR",
	CurrencySBD: "SBD",
	CurrencySCR: "SCR",
	CurrencySDG: "SDG",
	CurrencySEK: "SEK",
	CurrencySGD: "SGD",
	CurrencySHP: "SHP",
	CurrencySLL: "SLL",
	CurrencySOS: "SOS",
	CurrencySRD: "SRD",
	CurrencySSP: "SSP",
	CurrencySTD: "STD",
	CurrencySYP: "SYP",
	CurrencySZL: "SZL",
	CurrencyTHB: "THB",
	CurrencyTJS: "TJS",
	CurrencyTMT: "TMT",
	CurrencyTND: "TND",
	CurrencyTOP: "TOP",
	CurrencyTRY: "TRY",
	CurrencyTTD: "TTD",
	CurrencyTWD: "TWD",
	CurrencyTZS: "TZS",
	CurrencyUAH: "UAH",
	CurrencyUGX: "UGX",
	CurrencyUYU: "UYU",
	CurrencyUZS: "UZS",
	CurrencyVES: "VES",
	CurrencyVND: "VND",
	CurrencyVUV: "VUV",
	CurrencyWST: "WST",
	CurrencyYER: "YER",
	CurrencyZAR: "ZAR",
	CurrencyZMW: "ZMW",
}

var codeToCurrency = map[string]Currency{
	"EUR": CurrencyEUR,
	"AED": CurrencyAED,
	"AFN": CurrencyAFN,
	"XCD": CurrencyXCD,
	"ALL": CurrencyALL,
	"AMD": CurrencyAMD,
	"AOA": CurrencyAOA,
	"USD": CurrencyUSD,
	"ARS": CurrencyARS,
	"AUD": CurrencyAUD,
	"AWG": CurrencyAWG,
	"AZN": CurrencyAZN,
	"BAM": CurrencyBAM,
	"BBD": CurrencyBBD,
	"BDT": CurrencyBDT,
	"XOF": CurrencyXOF,
	"BGN": CurrencyBGN,
	"BHD": CurrencyBHD,
	"BIF": CurrencyBIF,
	"BMD": CurrencyBMD,
	"BND": CurrencyBND,
	"BOB": CurrencyBOB,
	"BRL": CurrencyBRL,
	"BSD": CurrencyBSD,
	"BTN": CurrencyBTN,
	"NOK": CurrencyNOK,
	"BWP": CurrencyBWP,
	"BYN": CurrencyBYN,
	"BZD": CurrencyBZD,
	"CAD": CurrencyCAD,
	"CDF": CurrencyCDF,
	"XAF": CurrencyXAF,
	"CHF": CurrencyCHF,
	"NZD": CurrencyNZD,
	"CLP": CurrencyCLP,
	"CNY": CurrencyCNY,
	"COP": CurrencyCOP,
	"CRC": CurrencyCRC,
	"CUP": CurrencyCUP,
	"CVE": CurrencyCVE,
	"ANG": CurrencyANG,
	"CZK": CurrencyCZK,
	"DJF": CurrencyDJF,
	"DKK": CurrencyDKK,
	"DOP": CurrencyDOP,
	"DZD": CurrencyDZD,
	"EGP": CurrencyEGP,
	"MAD": CurrencyMAD,
	"ETB": CurrencyETB,
	"FJD": CurrencyFJD,
	"FKP": CurrencyFKP,
	"GBP": CurrencyGBP,
	"GEL": CurrencyGEL,
	"GHS": CurrencyGHS,
	"GIP": CurrencyGIP,
	"GMD": CurrencyGMD,
	"GNF": CurrencyGNF,
	"GTQ": CurrencyGTQ,
	"GYD": CurrencyGYD,
	"HKD": CurrencyHKD,
	"HNL": CurrencyHNL,
	"HTG": CurrencyHTG,
	"HUF": CurrencyHUF,
	"IDR": CurrencyIDR,
	"ILS": CurrencyILS,
	"INR": CurrencyINR,
	"IQD": CurrencyIQD,
	"IRR": CurrencyIRR,
	"ISK": CurrencyISK,
	"JMD": CurrencyJMD,
	"JOD": CurrencyJOD,
	"JPY": CurrencyJPY,
	"KES": CurrencyKES,
	"KGS": CurrencyKGS,
	"KHR": CurrencyKHR,
	"KMF": CurrencyKMF,
	"KPW": CurrencyKPW,
	"KRW": CurrencyKRW,
	"KWD": CurrencyKWD,
	"KYD": CurrencyKYD,
	"KZT": CurrencyKZT,
	"LAK": CurrencyLAK,
	"LBP": CurrencyLBP,
	"LKR": CurrencyLKR,
	"LRD": CurrencyLRD,
	"LSL": CurrencyLSL,
	"LYD": CurrencyLYD,
	"MDL": CurrencyMDL,
	"MGA": CurrencyMGA,
	"MKD": CurrencyMKD,
	"MMK": CurrencyMMK,
	"MNT": CurrencyMNT,
	"MOP": CurrencyMOP,
	"MRU": CurrencyMRU,
	"MUR": CurrencyMUR,
	"MVR": CurrencyMVR,
	"MWK": CurrencyMWK,
	"MXN": CurrencyMXN,
	"MYR": CurrencyMYR,
	"MZN": CurrencyMZN,
	"NAD": CurrencyNAD,
	"XPF": CurrencyXPF,
	"NGN": CurrencyNGN,
	"NIO": CurrencyNIO,
	"NPR": CurrencyNPR,
	"OMR": CurrencyOMR,
	"PAB": CurrencyPAB,
	"PEN": CurrencyPEN,
	"PGK": CurrencyPGK,
	"PHP": CurrencyPHP,
	"PKR": CurrencyPKR,
	"PLN": CurrencyPLN,
	"PYG": CurrencyPYG,
	"QAR": CurrencyQAR,
	"RON": CurrencyRON,
	"RSD": CurrencyRSD,
	"RUB": CurrencyRUB,
	"RWF": CurrencyRWF,
	"SAR": CurrencySAR,
	"SBD": CurrencySBD,
	"SCR": CurrencySCR,
	"SDG": CurrencySDG,
	"SEK": CurrencySEK,
	"SGD": CurrencySGD,
	"SHP": CurrencySHP,
	"SLL": CurrencySLL,
	"SOS": CurrencySOS,
	"SRD": CurrencySRD,
	"SSP": CurrencySSP,
	"STD": CurrencySTD,
	"SYP": CurrencySYP,
	"SZL": CurrencySZL,
	"THB": CurrencyTHB,
	"TJS": CurrencyTJS,
	"TMT": CurrencyTMT,
	"TND": CurrencyTND,
	"TOP": CurrencyTOP,
	"TRY": CurrencyTRY,
	"TTD": CurrencyTTD,
	"TWD": CurrencyTWD,
	"TZS": CurrencyTZS,
	"UAH": CurrencyUAH,
	"UGX": CurrencyUGX,
	"UYU": CurrencyUYU,
	"UZS": CurrencyUZS,
	"VES": CurrencyVES,
	"VND": CurrencyVND,
	"VUV": CurrencyVUV,
	"WST": CurrencyWST,
	"YER": CurrencyYER,
	"ZAR": CurrencyZAR,
	"ZMW": CurrencyZMW,
}

func CurrencyFromCode(v string) Currency {
	return codeToCurrency[v]
}

func AppendCurrencies(dst []Currency) []Currency {
	dst = append(dst, CurrencyEUR)
	dst = append(dst, CurrencyAED)
	dst = append(dst, CurrencyAFN)
	dst = append(dst, CurrencyXCD)
	dst = append(dst, CurrencyALL)
	dst = append(dst, CurrencyAMD)
	dst = append(dst, CurrencyAOA)
	dst = append(dst, CurrencyUSD)
	dst = append(dst, CurrencyARS)
	dst = append(dst, CurrencyAUD)
	dst = append(dst, CurrencyAWG)
	dst = append(dst, CurrencyAZN)
	dst = append(dst, CurrencyBAM)
	dst = append(dst, CurrencyBBD)
	dst = append(dst, CurrencyBDT)
	dst = append(dst, CurrencyXOF)
	dst = append(dst, CurrencyBGN)
	dst = append(dst, CurrencyBHD)
	dst = append(dst, CurrencyBIF)
	dst = append(dst, CurrencyBMD)
	dst = append(dst, CurrencyBND)
	dst = append(dst, CurrencyBOB)
	dst = append(dst, CurrencyBRL)
	dst = append(dst, CurrencyBSD)
	dst = append(dst, CurrencyBTN)
	dst = append(dst, CurrencyNOK)
	dst = append(dst, CurrencyBWP)
	dst = append(dst, CurrencyBYN)
	dst = append(dst, CurrencyBZD)
	dst = append(dst, CurrencyCAD)
	dst = append(dst, CurrencyCDF)
	dst = append(dst, CurrencyXAF)
	dst = append(dst, CurrencyCHF)
	dst = append(dst, CurrencyNZD)
	dst = append(dst, CurrencyCLP)
	dst = append(dst, CurrencyCNY)
	dst = append(dst, CurrencyCOP)
	dst = append(dst, CurrencyCRC)
	dst = append(dst, CurrencyCUP)
	dst = append(dst, CurrencyCVE)
	dst = append(dst, CurrencyANG)
	dst = append(dst, CurrencyCZK)
	dst = append(dst, CurrencyDJF)
	dst = append(dst, CurrencyDKK)
	dst = append(dst, CurrencyDOP)
	dst = append(dst, CurrencyDZD)
	dst = append(dst, CurrencyEGP)
	dst = append(dst, CurrencyMAD)
	dst = append(dst, CurrencyETB)
	dst = append(dst, CurrencyFJD)
	dst = append(dst, CurrencyFKP)
	dst = append(dst, CurrencyGBP)
	dst = append(dst, CurrencyGEL)
	dst = append(dst, CurrencyGHS)
	dst = append(dst, CurrencyGIP)
	dst = append(dst, CurrencyGMD)
	dst = append(dst, CurrencyGNF)
	dst = append(dst, CurrencyGTQ)
	dst = append(dst, CurrencyGYD)
	dst = append(dst, CurrencyHKD)
	dst = append(dst, CurrencyHNL)
	dst = append(dst, CurrencyHTG)
	dst = append(dst, CurrencyHUF)
	dst = append(dst, CurrencyIDR)
	dst = append(dst, CurrencyILS)
	dst = append(dst, CurrencyINR)
	dst = append(dst, CurrencyIQD)
	dst = append(dst, CurrencyIRR)
	dst = append(dst, CurrencyISK)
	dst = append(dst, CurrencyJMD)
	dst = append(dst, CurrencyJOD)
	dst = append(dst, CurrencyJPY)
	dst = append(dst, CurrencyKES)
	dst = append(dst, CurrencyKGS)
	dst = append(dst, CurrencyKHR)
	dst = append(dst, CurrencyKMF)
	dst = append(dst, CurrencyKPW)
	dst = append(dst, CurrencyKRW)
	dst = append(dst, CurrencyKWD)
	dst = append(dst, CurrencyKYD)
	dst = append(dst, CurrencyKZT)
	dst = append(dst, CurrencyLAK)
	dst = append(dst, CurrencyLBP)
	dst = append(dst, CurrencyLKR)
	dst = append(dst, CurrencyLRD)
	dst = append(dst, CurrencyLSL)
	dst = append(dst, CurrencyLYD)
	dst = append(dst, CurrencyMDL)
	dst = append(dst, CurrencyMGA)
	dst = append(dst, CurrencyMKD)
	dst = append(dst, CurrencyMMK)
	dst = append(dst, CurrencyMNT)
	dst = append(dst, CurrencyMOP)
	dst = append(dst, CurrencyMRU)
	dst = append(dst, CurrencyMUR)
	dst = append(dst, CurrencyMVR)
	dst = append(dst, CurrencyMWK)
	dst = append(dst, CurrencyMXN)
	dst = append(dst, CurrencyMYR)
	dst = append(dst, CurrencyMZN)
	dst = append(dst, CurrencyNAD)
	dst = append(dst, CurrencyXPF)
	dst = append(dst, CurrencyNGN)
	dst = append(dst, CurrencyNIO)
	dst = append(dst, CurrencyNPR)
	dst = append(dst, CurrencyOMR)
	dst = append(dst, CurrencyPAB)
	dst = append(dst, CurrencyPEN)
	dst = append(dst, CurrencyPGK)
	dst = append(dst, CurrencyPHP)
	dst = append(dst, CurrencyPKR)
	dst = append(dst, CurrencyPLN)
	dst = append(dst, CurrencyPYG)
	dst = append(dst, CurrencyQAR)
	dst = append(dst, CurrencyRON)
	dst = append(dst, CurrencyRSD)
	dst = append(dst, CurrencyRUB)
	dst = append(dst, CurrencyRWF)
	dst = append(dst, CurrencySAR)
	dst = append(dst, CurrencySBD)
	dst = append(dst, CurrencySCR)
	dst = append(dst, CurrencySDG)
	dst = append(dst, CurrencySEK)
	dst = append(dst, CurrencySGD)
	dst = append(dst, CurrencySHP)
	dst = append(dst, CurrencySLL)
	dst = append(dst, CurrencySOS)
	dst = append(dst, CurrencySRD)
	dst = append(dst, CurrencySSP)
	dst = append(dst, CurrencySTD)
	dst = append(dst, CurrencySYP)
	dst = append(dst, CurrencySZL)
	dst = append(dst, CurrencyTHB)
	dst = append(dst, CurrencyTJS)
	dst = append(dst, CurrencyTMT)
	dst = append(dst, CurrencyTND)
	dst = append(dst, CurrencyTOP)
	dst = append(dst, CurrencyTRY)
	dst = append(dst, CurrencyTTD)
	dst = append(dst, CurrencyTWD)
	dst = append(dst, CurrencyTZS)
	dst = append(dst, CurrencyUAH)
	dst = append(dst, CurrencyUGX)
	dst = append(dst, CurrencyUYU)
	dst = append(dst, CurrencyUZS)
	dst = append(dst, CurrencyVES)
	dst = append(dst, CurrencyVND)
	dst = append(dst, CurrencyVUV)
	dst = append(dst, CurrencyWST)
	dst = append(dst, CurrencyYER)
	dst = append(dst, CurrencyZAR)
	dst = append(dst, CurrencyZMW)
	return dst
}

func (v Currency) String() string {
	return currencyToCode[v]
}

func (v Currency) IsValid() bool {
	return v.String() != ""
}

func (v Currency) MarshalText() ([]byte, error) {
	code := v.String()
	if len(code) == 0 {
		return nil, ErrBadCurrency
	}
	return s2b(code), nil
}

func (v *Currency) UnmarshalText(bb []byte) error {
	v1, found := codeToCurrency[b2s(bb)]
	if !found {
		return ErrBadCurrency
	}
	*v = v1
	return nil
}

func (v *Currency) UnmarshalJSON(bb []byte) error {
	if len(bb) < 2 {
		return ErrBadCurrency
	}
	return v.UnmarshalText(unquote(bb))
}

func (v Currency) Value() (driver.Value, error) {
	return v.String(), nil
}

func (v *Currency) Scan(v1 interface{}) error {
	switch x := v1.(type) {
	case string:
		return v.UnmarshalText(s2b(x))
	case []byte:
		return v.UnmarshalText(x)
	}
	return ErrBadCurrency
}

func (nv *NullCurrency) UnmarshalJSON(bb []byte) error {
	if isJsonNull(bb) {
		nv.Valid = false
		return nil
	}

	v := Currency(0)
	if err := v.UnmarshalJSON(bb); err != nil {
		nv.Valid = false
		return err
	}

	nv.Currency = v
	nv.Valid = true
	return nil
}

func (nv *NullCurrency) Scan(v1 interface{}) error {
	if v1 == nil {
		nv.Valid = false
		return nil
	}

	v := Currency(0)
	if err := v.Scan(v1); err != nil {
		nv.Valid = false
		return err
	}

	nv.Currency = v
	nv.Valid = true
	return nil
}
