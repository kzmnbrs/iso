// Code generated from iso/gen. DO NOT EDIT.

package iso

import (
	"database/sql/driver"
	"errors"
	"regexp"
)

type (
	Country     uint8
	NullCountry struct {
		Valid   bool
		Country Country
	}
)

const (
	_ Country = iota
	CountryAD
	CountryAE
	CountryAF
	CountryAG
	CountryAI
	CountryAL
	CountryAM
	CountryAO
	CountryAQ
	CountryAR
	CountryAS
	CountryAT
	CountryAU
	CountryAW
	CountryAX
	CountryAZ
	CountryBA
	CountryBB
	CountryBD
	CountryBE
	CountryBF
	CountryBG
	CountryBH
	CountryBI
	CountryBJ
	CountryBL
	CountryBM
	CountryBN
	CountryBO
	CountryBQ
	CountryBR
	CountryBS
	CountryBT
	CountryBV
	CountryBW
	CountryBY
	CountryBZ
	CountryCA
	CountryCC
	CountryCD
	CountryCF
	CountryCG
	CountryCH
	CountryCI
	CountryCK
	CountryCL
	CountryCM
	CountryCN
	CountryCO
	CountryCR
	CountryCU
	CountryCV
	CountryCW
	CountryCX
	CountryCY
	CountryCZ
	CountryDE
	CountryDJ
	CountryDK
	CountryDM
	CountryDO
	CountryDZ
	CountryEC
	CountryEE
	CountryEG
	CountryEH
	CountryER
	CountryES
	CountryET
	CountryFI
	CountryFJ
	CountryFK
	CountryFM
	CountryFO
	CountryFR
	CountryGA
	CountryGB
	CountryGD
	CountryGE
	CountryGF
	CountryGG
	CountryGH
	CountryGI
	CountryGL
	CountryGM
	CountryGN
	CountryGP
	CountryGQ
	CountryGR
	CountryGS
	CountryGT
	CountryGU
	CountryGW
	CountryGY
	CountryHK
	CountryHM
	CountryHN
	CountryHR
	CountryHT
	CountryHU
	CountryID
	CountryIE
	CountryIL
	CountryIM
	CountryIN
	CountryIO
	CountryIQ
	CountryIR
	CountryIS
	CountryIT
	CountryJE
	CountryJM
	CountryJO
	CountryJP
	CountryKE
	CountryKG
	CountryKH
	CountryKI
	CountryKM
	CountryKN
	CountryKP
	CountryKR
	CountryKW
	CountryKY
	CountryKZ
	CountryLA
	CountryLB
	CountryLC
	CountryLI
	CountryLK
	CountryLR
	CountryLS
	CountryLT
	CountryLU
	CountryLV
	CountryLY
	CountryMA
	CountryMC
	CountryMD
	CountryME
	CountryMF
	CountryMG
	CountryMH
	CountryMK
	CountryML
	CountryMM
	CountryMN
	CountryMO
	CountryMP
	CountryMQ
	CountryMR
	CountryMS
	CountryMT
	CountryMU
	CountryMV
	CountryMW
	CountryMX
	CountryMY
	CountryMZ
	CountryNA
	CountryNC
	CountryNE
	CountryNF
	CountryNG
	CountryNI
	CountryNL
	CountryNO
	CountryNP
	CountryNR
	CountryNU
	CountryNZ
	CountryOM
	CountryPA
	CountryPE
	CountryPF
	CountryPG
	CountryPH
	CountryPK
	CountryPL
	CountryPM
	CountryPN
	CountryPR
	CountryPS
	CountryPT
	CountryPW
	CountryPY
	CountryQA
	CountryRE
	CountryRO
	CountryRS
	CountryRU
	CountryRW
	CountrySA
	CountrySB
	CountrySC
	CountrySD
	CountrySE
	CountrySG
	CountrySH
	CountrySI
	CountrySJ
	CountrySK
	CountrySL
	CountrySM
	CountrySN
	CountrySO
	CountrySR
	CountrySS
	CountryST
	CountrySV
	CountrySX
	CountrySY
	CountrySZ
	CountryTC
	CountryTD
	CountryTF
	CountryTG
	CountryTH
	CountryTJ
	CountryTK
	CountryTL
	CountryTM
	CountryTN
	CountryTO
	CountryTR
	CountryTT
	CountryTV
	CountryTW
	CountryTZ
	CountryUA
	CountryUG
	CountryUM
	CountryUS
	CountryUY
	CountryUZ
	CountryVA
	CountryVC
	CountryVE
	CountryVG
	CountryVI
	CountryVN
	CountryVU
	CountryWF
	CountryWS
	CountryYE
	CountryYT
	CountryZA
	CountryZM
	CountryZW
)

const (
	CountryAND = CountryAD
	CountryARE = CountryAE
	CountryAFG = CountryAF
	CountryATG = CountryAG
	CountryAIA = CountryAI
	CountryALB = CountryAL
	CountryARM = CountryAM
	CountryAGO = CountryAO
	CountryATA = CountryAQ
	CountryARG = CountryAR
	CountryASM = CountryAS
	CountryAUT = CountryAT
	CountryAUS = CountryAU
	CountryABW = CountryAW
	CountryALA = CountryAX
	CountryAZE = CountryAZ
	CountryBIH = CountryBA
	CountryBRB = CountryBB
	CountryBGD = CountryBD
	CountryBEL = CountryBE
	CountryBFA = CountryBF
	CountryBGR = CountryBG
	CountryBHR = CountryBH
	CountryBDI = CountryBI
	CountryBEN = CountryBJ
	CountryBLM = CountryBL
	CountryBMU = CountryBM
	CountryBRN = CountryBN
	CountryBOL = CountryBO
	CountryBES = CountryBQ
	CountryBRA = CountryBR
	CountryBHS = CountryBS
	CountryBTN = CountryBT
	CountryBVT = CountryBV
	CountryBWA = CountryBW
	CountryBLR = CountryBY
	CountryBLZ = CountryBZ
	CountryCAN = CountryCA
	CountryCCK = CountryCC
	CountryCOD = CountryCD
	CountryCAF = CountryCF
	CountryCOG = CountryCG
	CountryCHE = CountryCH
	CountryCIV = CountryCI
	CountryCOK = CountryCK
	CountryCHL = CountryCL
	CountryCMR = CountryCM
	CountryCHN = CountryCN
	CountryCOL = CountryCO
	CountryCRI = CountryCR
	CountryCUB = CountryCU
	CountryCPV = CountryCV
	CountryCUW = CountryCW
	CountryCXR = CountryCX
	CountryCYP = CountryCY
	CountryCZE = CountryCZ
	CountryDEU = CountryDE
	CountryDJI = CountryDJ
	CountryDNK = CountryDK
	CountryDMA = CountryDM
	CountryDOM = CountryDO
	CountryDZA = CountryDZ
	CountryECU = CountryEC
	CountryEST = CountryEE
	CountryEGY = CountryEG
	CountryESH = CountryEH
	CountryERI = CountryER
	CountryESP = CountryES
	CountryETH = CountryET
	CountryFIN = CountryFI
	CountryFJI = CountryFJ
	CountryFLK = CountryFK
	CountryFSM = CountryFM
	CountryFRO = CountryFO
	CountryFRA = CountryFR
	CountryGAB = CountryGA
	CountryGBR = CountryGB
	CountryGRD = CountryGD
	CountryGEO = CountryGE
	CountryGUF = CountryGF
	CountryGGY = CountryGG
	CountryGHA = CountryGH
	CountryGIB = CountryGI
	CountryGRL = CountryGL
	CountryGMB = CountryGM
	CountryGIN = CountryGN
	CountryGLP = CountryGP
	CountryGNQ = CountryGQ
	CountryGRC = CountryGR
	CountrySGS = CountryGS
	CountryGTM = CountryGT
	CountryGUM = CountryGU
	CountryGNB = CountryGW
	CountryGUY = CountryGY
	CountryHKG = CountryHK
	CountryHMD = CountryHM
	CountryHND = CountryHN
	CountryHRV = CountryHR
	CountryHTI = CountryHT
	CountryHUN = CountryHU
	CountryIDN = CountryID
	CountryIRL = CountryIE
	CountryISR = CountryIL
	CountryIMN = CountryIM
	CountryIND = CountryIN
	CountryIOT = CountryIO
	CountryIRQ = CountryIQ
	CountryIRN = CountryIR
	CountryISL = CountryIS
	CountryITA = CountryIT
	CountryJEY = CountryJE
	CountryJAM = CountryJM
	CountryJOR = CountryJO
	CountryJPN = CountryJP
	CountryKEN = CountryKE
	CountryKGZ = CountryKG
	CountryKHM = CountryKH
	CountryKIR = CountryKI
	CountryCOM = CountryKM
	CountryKNA = CountryKN
	CountryPRK = CountryKP
	CountryKOR = CountryKR
	CountryKWT = CountryKW
	CountryCYM = CountryKY
	CountryKAZ = CountryKZ
	CountryLAO = CountryLA
	CountryLBN = CountryLB
	CountryLCA = CountryLC
	CountryLIE = CountryLI
	CountryLKA = CountryLK
	CountryLBR = CountryLR
	CountryLSO = CountryLS
	CountryLTU = CountryLT
	CountryLUX = CountryLU
	CountryLVA = CountryLV
	CountryLBY = CountryLY
	CountryMAR = CountryMA
	CountryMCO = CountryMC
	CountryMDA = CountryMD
	CountryMNE = CountryME
	CountryMAF = CountryMF
	CountryMDG = CountryMG
	CountryMHL = CountryMH
	CountryMKD = CountryMK
	CountryMLI = CountryML
	CountryMMR = CountryMM
	CountryMNG = CountryMN
	CountryMAC = CountryMO
	CountryMNP = CountryMP
	CountryMTQ = CountryMQ
	CountryMRT = CountryMR
	CountryMSR = CountryMS
	CountryMLT = CountryMT
	CountryMUS = CountryMU
	CountryMDV = CountryMV
	CountryMWI = CountryMW
	CountryMEX = CountryMX
	CountryMYS = CountryMY
	CountryMOZ = CountryMZ
	CountryNAM = CountryNA
	CountryNCL = CountryNC
	CountryNER = CountryNE
	CountryNFK = CountryNF
	CountryNGA = CountryNG
	CountryNIC = CountryNI
	CountryNLD = CountryNL
	CountryNOR = CountryNO
	CountryNPL = CountryNP
	CountryNRU = CountryNR
	CountryNIU = CountryNU
	CountryNZL = CountryNZ
	CountryOMN = CountryOM
	CountryPAN = CountryPA
	CountryPER = CountryPE
	CountryPYF = CountryPF
	CountryPNG = CountryPG
	CountryPHL = CountryPH
	CountryPAK = CountryPK
	CountryPOL = CountryPL
	CountrySPM = CountryPM
	CountryPCN = CountryPN
	CountryPRI = CountryPR
	CountryPSE = CountryPS
	CountryPRT = CountryPT
	CountryPLW = CountryPW
	CountryPRY = CountryPY
	CountryQAT = CountryQA
	CountryREU = CountryRE
	CountryROU = CountryRO
	CountrySRB = CountryRS
	CountryRUS = CountryRU
	CountryRWA = CountryRW
	CountrySAU = CountrySA
	CountrySLB = CountrySB
	CountrySYC = CountrySC
	CountrySDN = CountrySD
	CountrySWE = CountrySE
	CountrySGP = CountrySG
	CountrySHN = CountrySH
	CountrySVN = CountrySI
	CountrySJM = CountrySJ
	CountrySVK = CountrySK
	CountrySLE = CountrySL
	CountrySMR = CountrySM
	CountrySEN = CountrySN
	CountrySOM = CountrySO
	CountrySUR = CountrySR
	CountrySSD = CountrySS
	CountrySTP = CountryST
	CountrySLV = CountrySV
	CountrySXM = CountrySX
	CountrySYR = CountrySY
	CountrySWZ = CountrySZ
	CountryTCA = CountryTC
	CountryTCD = CountryTD
	CountryATF = CountryTF
	CountryTGO = CountryTG
	CountryTHA = CountryTH
	CountryTJK = CountryTJ
	CountryTKL = CountryTK
	CountryTLS = CountryTL
	CountryTKM = CountryTM
	CountryTUN = CountryTN
	CountryTON = CountryTO
	CountryTUR = CountryTR
	CountryTTO = CountryTT
	CountryTUV = CountryTV
	CountryTWN = CountryTW
	CountryTZA = CountryTZ
	CountryUKR = CountryUA
	CountryUGA = CountryUG
	CountryUMI = CountryUM
	CountryUSA = CountryUS
	CountryURY = CountryUY
	CountryUZB = CountryUZ
	CountryVAT = CountryVA
	CountryVCT = CountryVC
	CountryVEN = CountryVE
	CountryVGB = CountryVG
	CountryVIR = CountryVI
	CountryVNM = CountryVN
	CountryVUT = CountryVU
	CountryWLF = CountryWF
	CountryWSM = CountryWS
	CountryYEM = CountryYE
	CountryMYT = CountryYT
	CountryZAF = CountryZA
	CountryZMB = CountryZM
	CountryZWE = CountryZW
)

var ErrBadCountry = errors.New("bad country")

var countryToAlpha2 = map[Country]string{
	CountryAD: "AD",
	CountryAE: "AE",
	CountryAF: "AF",
	CountryAG: "AG",
	CountryAI: "AI",
	CountryAL: "AL",
	CountryAM: "AM",
	CountryAO: "AO",
	CountryAQ: "AQ",
	CountryAR: "AR",
	CountryAS: "AS",
	CountryAT: "AT",
	CountryAU: "AU",
	CountryAW: "AW",
	CountryAX: "AX",
	CountryAZ: "AZ",
	CountryBA: "BA",
	CountryBB: "BB",
	CountryBD: "BD",
	CountryBE: "BE",
	CountryBF: "BF",
	CountryBG: "BG",
	CountryBH: "BH",
	CountryBI: "BI",
	CountryBJ: "BJ",
	CountryBL: "BL",
	CountryBM: "BM",
	CountryBN: "BN",
	CountryBO: "BO",
	CountryBQ: "BQ",
	CountryBR: "BR",
	CountryBS: "BS",
	CountryBT: "BT",
	CountryBV: "BV",
	CountryBW: "BW",
	CountryBY: "BY",
	CountryBZ: "BZ",
	CountryCA: "CA",
	CountryCC: "CC",
	CountryCD: "CD",
	CountryCF: "CF",
	CountryCG: "CG",
	CountryCH: "CH",
	CountryCI: "CI",
	CountryCK: "CK",
	CountryCL: "CL",
	CountryCM: "CM",
	CountryCN: "CN",
	CountryCO: "CO",
	CountryCR: "CR",
	CountryCU: "CU",
	CountryCV: "CV",
	CountryCW: "CW",
	CountryCX: "CX",
	CountryCY: "CY",
	CountryCZ: "CZ",
	CountryDE: "DE",
	CountryDJ: "DJ",
	CountryDK: "DK",
	CountryDM: "DM",
	CountryDO: "DO",
	CountryDZ: "DZ",
	CountryEC: "EC",
	CountryEE: "EE",
	CountryEG: "EG",
	CountryEH: "EH",
	CountryER: "ER",
	CountryES: "ES",
	CountryET: "ET",
	CountryFI: "FI",
	CountryFJ: "FJ",
	CountryFK: "FK",
	CountryFM: "FM",
	CountryFO: "FO",
	CountryFR: "FR",
	CountryGA: "GA",
	CountryGB: "GB",
	CountryGD: "GD",
	CountryGE: "GE",
	CountryGF: "GF",
	CountryGG: "GG",
	CountryGH: "GH",
	CountryGI: "GI",
	CountryGL: "GL",
	CountryGM: "GM",
	CountryGN: "GN",
	CountryGP: "GP",
	CountryGQ: "GQ",
	CountryGR: "GR",
	CountryGS: "GS",
	CountryGT: "GT",
	CountryGU: "GU",
	CountryGW: "GW",
	CountryGY: "GY",
	CountryHK: "HK",
	CountryHM: "HM",
	CountryHN: "HN",
	CountryHR: "HR",
	CountryHT: "HT",
	CountryHU: "HU",
	CountryID: "ID",
	CountryIE: "IE",
	CountryIL: "IL",
	CountryIM: "IM",
	CountryIN: "IN",
	CountryIO: "IO",
	CountryIQ: "IQ",
	CountryIR: "IR",
	CountryIS: "IS",
	CountryIT: "IT",
	CountryJE: "JE",
	CountryJM: "JM",
	CountryJO: "JO",
	CountryJP: "JP",
	CountryKE: "KE",
	CountryKG: "KG",
	CountryKH: "KH",
	CountryKI: "KI",
	CountryKM: "KM",
	CountryKN: "KN",
	CountryKP: "KP",
	CountryKR: "KR",
	CountryKW: "KW",
	CountryKY: "KY",
	CountryKZ: "KZ",
	CountryLA: "LA",
	CountryLB: "LB",
	CountryLC: "LC",
	CountryLI: "LI",
	CountryLK: "LK",
	CountryLR: "LR",
	CountryLS: "LS",
	CountryLT: "LT",
	CountryLU: "LU",
	CountryLV: "LV",
	CountryLY: "LY",
	CountryMA: "MA",
	CountryMC: "MC",
	CountryMD: "MD",
	CountryME: "ME",
	CountryMF: "MF",
	CountryMG: "MG",
	CountryMH: "MH",
	CountryMK: "MK",
	CountryML: "ML",
	CountryMM: "MM",
	CountryMN: "MN",
	CountryMO: "MO",
	CountryMP: "MP",
	CountryMQ: "MQ",
	CountryMR: "MR",
	CountryMS: "MS",
	CountryMT: "MT",
	CountryMU: "MU",
	CountryMV: "MV",
	CountryMW: "MW",
	CountryMX: "MX",
	CountryMY: "MY",
	CountryMZ: "MZ",
	CountryNA: "NA",
	CountryNC: "NC",
	CountryNE: "NE",
	CountryNF: "NF",
	CountryNG: "NG",
	CountryNI: "NI",
	CountryNL: "NL",
	CountryNO: "NO",
	CountryNP: "NP",
	CountryNR: "NR",
	CountryNU: "NU",
	CountryNZ: "NZ",
	CountryOM: "OM",
	CountryPA: "PA",
	CountryPE: "PE",
	CountryPF: "PF",
	CountryPG: "PG",
	CountryPH: "PH",
	CountryPK: "PK",
	CountryPL: "PL",
	CountryPM: "PM",
	CountryPN: "PN",
	CountryPR: "PR",
	CountryPS: "PS",
	CountryPT: "PT",
	CountryPW: "PW",
	CountryPY: "PY",
	CountryQA: "QA",
	CountryRE: "RE",
	CountryRO: "RO",
	CountryRS: "RS",
	CountryRU: "RU",
	CountryRW: "RW",
	CountrySA: "SA",
	CountrySB: "SB",
	CountrySC: "SC",
	CountrySD: "SD",
	CountrySE: "SE",
	CountrySG: "SG",
	CountrySH: "SH",
	CountrySI: "SI",
	CountrySJ: "SJ",
	CountrySK: "SK",
	CountrySL: "SL",
	CountrySM: "SM",
	CountrySN: "SN",
	CountrySO: "SO",
	CountrySR: "SR",
	CountrySS: "SS",
	CountryST: "ST",
	CountrySV: "SV",
	CountrySX: "SX",
	CountrySY: "SY",
	CountrySZ: "SZ",
	CountryTC: "TC",
	CountryTD: "TD",
	CountryTF: "TF",
	CountryTG: "TG",
	CountryTH: "TH",
	CountryTJ: "TJ",
	CountryTK: "TK",
	CountryTL: "TL",
	CountryTM: "TM",
	CountryTN: "TN",
	CountryTO: "TO",
	CountryTR: "TR",
	CountryTT: "TT",
	CountryTV: "TV",
	CountryTW: "TW",
	CountryTZ: "TZ",
	CountryUA: "UA",
	CountryUG: "UG",
	CountryUM: "UM",
	CountryUS: "US",
	CountryUY: "UY",
	CountryUZ: "UZ",
	CountryVA: "VA",
	CountryVC: "VC",
	CountryVE: "VE",
	CountryVG: "VG",
	CountryVI: "VI",
	CountryVN: "VN",
	CountryVU: "VU",
	CountryWF: "WF",
	CountryWS: "WS",
	CountryYE: "YE",
	CountryYT: "YT",
	CountryZA: "ZA",
	CountryZM: "ZM",
	CountryZW: "ZW",
}

var countryToAlpha3 = map[Country]string{
	CountryAND: "AND",
	CountryARE: "ARE",
	CountryAFG: "AFG",
	CountryATG: "ATG",
	CountryAIA: "AIA",
	CountryALB: "ALB",
	CountryARM: "ARM",
	CountryAGO: "AGO",
	CountryATA: "ATA",
	CountryARG: "ARG",
	CountryASM: "ASM",
	CountryAUT: "AUT",
	CountryAUS: "AUS",
	CountryABW: "ABW",
	CountryALA: "ALA",
	CountryAZE: "AZE",
	CountryBIH: "BIH",
	CountryBRB: "BRB",
	CountryBGD: "BGD",
	CountryBEL: "BEL",
	CountryBFA: "BFA",
	CountryBGR: "BGR",
	CountryBHR: "BHR",
	CountryBDI: "BDI",
	CountryBEN: "BEN",
	CountryBLM: "BLM",
	CountryBMU: "BMU",
	CountryBRN: "BRN",
	CountryBOL: "BOL",
	CountryBES: "BES",
	CountryBRA: "BRA",
	CountryBHS: "BHS",
	CountryBTN: "BTN",
	CountryBVT: "BVT",
	CountryBWA: "BWA",
	CountryBLR: "BLR",
	CountryBLZ: "BLZ",
	CountryCAN: "CAN",
	CountryCCK: "CCK",
	CountryCOD: "COD",
	CountryCAF: "CAF",
	CountryCOG: "COG",
	CountryCHE: "CHE",
	CountryCIV: "CIV",
	CountryCOK: "COK",
	CountryCHL: "CHL",
	CountryCMR: "CMR",
	CountryCHN: "CHN",
	CountryCOL: "COL",
	CountryCRI: "CRI",
	CountryCUB: "CUB",
	CountryCPV: "CPV",
	CountryCUW: "CUW",
	CountryCXR: "CXR",
	CountryCYP: "CYP",
	CountryCZE: "CZE",
	CountryDEU: "DEU",
	CountryDJI: "DJI",
	CountryDNK: "DNK",
	CountryDMA: "DMA",
	CountryDOM: "DOM",
	CountryDZA: "DZA",
	CountryECU: "ECU",
	CountryEST: "EST",
	CountryEGY: "EGY",
	CountryESH: "ESH",
	CountryERI: "ERI",
	CountryESP: "ESP",
	CountryETH: "ETH",
	CountryFIN: "FIN",
	CountryFJI: "FJI",
	CountryFLK: "FLK",
	CountryFSM: "FSM",
	CountryFRO: "FRO",
	CountryFRA: "FRA",
	CountryGAB: "GAB",
	CountryGBR: "GBR",
	CountryGRD: "GRD",
	CountryGEO: "GEO",
	CountryGUF: "GUF",
	CountryGGY: "GGY",
	CountryGHA: "GHA",
	CountryGIB: "GIB",
	CountryGRL: "GRL",
	CountryGMB: "GMB",
	CountryGIN: "GIN",
	CountryGLP: "GLP",
	CountryGNQ: "GNQ",
	CountryGRC: "GRC",
	CountrySGS: "SGS",
	CountryGTM: "GTM",
	CountryGUM: "GUM",
	CountryGNB: "GNB",
	CountryGUY: "GUY",
	CountryHKG: "HKG",
	CountryHMD: "HMD",
	CountryHND: "HND",
	CountryHRV: "HRV",
	CountryHTI: "HTI",
	CountryHUN: "HUN",
	CountryIDN: "IDN",
	CountryIRL: "IRL",
	CountryISR: "ISR",
	CountryIMN: "IMN",
	CountryIND: "IND",
	CountryIOT: "IOT",
	CountryIRQ: "IRQ",
	CountryIRN: "IRN",
	CountryISL: "ISL",
	CountryITA: "ITA",
	CountryJEY: "JEY",
	CountryJAM: "JAM",
	CountryJOR: "JOR",
	CountryJPN: "JPN",
	CountryKEN: "KEN",
	CountryKGZ: "KGZ",
	CountryKHM: "KHM",
	CountryKIR: "KIR",
	CountryCOM: "COM",
	CountryKNA: "KNA",
	CountryPRK: "PRK",
	CountryKOR: "KOR",
	CountryKWT: "KWT",
	CountryCYM: "CYM",
	CountryKAZ: "KAZ",
	CountryLAO: "LAO",
	CountryLBN: "LBN",
	CountryLCA: "LCA",
	CountryLIE: "LIE",
	CountryLKA: "LKA",
	CountryLBR: "LBR",
	CountryLSO: "LSO",
	CountryLTU: "LTU",
	CountryLUX: "LUX",
	CountryLVA: "LVA",
	CountryLBY: "LBY",
	CountryMAR: "MAR",
	CountryMCO: "MCO",
	CountryMDA: "MDA",
	CountryMNE: "MNE",
	CountryMAF: "MAF",
	CountryMDG: "MDG",
	CountryMHL: "MHL",
	CountryMKD: "MKD",
	CountryMLI: "MLI",
	CountryMMR: "MMR",
	CountryMNG: "MNG",
	CountryMAC: "MAC",
	CountryMNP: "MNP",
	CountryMTQ: "MTQ",
	CountryMRT: "MRT",
	CountryMSR: "MSR",
	CountryMLT: "MLT",
	CountryMUS: "MUS",
	CountryMDV: "MDV",
	CountryMWI: "MWI",
	CountryMEX: "MEX",
	CountryMYS: "MYS",
	CountryMOZ: "MOZ",
	CountryNAM: "NAM",
	CountryNCL: "NCL",
	CountryNER: "NER",
	CountryNFK: "NFK",
	CountryNGA: "NGA",
	CountryNIC: "NIC",
	CountryNLD: "NLD",
	CountryNOR: "NOR",
	CountryNPL: "NPL",
	CountryNRU: "NRU",
	CountryNIU: "NIU",
	CountryNZL: "NZL",
	CountryOMN: "OMN",
	CountryPAN: "PAN",
	CountryPER: "PER",
	CountryPYF: "PYF",
	CountryPNG: "PNG",
	CountryPHL: "PHL",
	CountryPAK: "PAK",
	CountryPOL: "POL",
	CountrySPM: "SPM",
	CountryPCN: "PCN",
	CountryPRI: "PRI",
	CountryPSE: "PSE",
	CountryPRT: "PRT",
	CountryPLW: "PLW",
	CountryPRY: "PRY",
	CountryQAT: "QAT",
	CountryREU: "REU",
	CountryROU: "ROU",
	CountrySRB: "SRB",
	CountryRUS: "RUS",
	CountryRWA: "RWA",
	CountrySAU: "SAU",
	CountrySLB: "SLB",
	CountrySYC: "SYC",
	CountrySDN: "SDN",
	CountrySWE: "SWE",
	CountrySGP: "SGP",
	CountrySHN: "SHN",
	CountrySVN: "SVN",
	CountrySJM: "SJM",
	CountrySVK: "SVK",
	CountrySLE: "SLE",
	CountrySMR: "SMR",
	CountrySEN: "SEN",
	CountrySOM: "SOM",
	CountrySUR: "SUR",
	CountrySSD: "SSD",
	CountrySTP: "STP",
	CountrySLV: "SLV",
	CountrySXM: "SXM",
	CountrySYR: "SYR",
	CountrySWZ: "SWZ",
	CountryTCA: "TCA",
	CountryTCD: "TCD",
	CountryATF: "ATF",
	CountryTGO: "TGO",
	CountryTHA: "THA",
	CountryTJK: "TJK",
	CountryTKL: "TKL",
	CountryTLS: "TLS",
	CountryTKM: "TKM",
	CountryTUN: "TUN",
	CountryTON: "TON",
	CountryTUR: "TUR",
	CountryTTO: "TTO",
	CountryTUV: "TUV",
	CountryTWN: "TWN",
	CountryTZA: "TZA",
	CountryUKR: "UKR",
	CountryUGA: "UGA",
	CountryUMI: "UMI",
	CountryUSA: "USA",
	CountryURY: "URY",
	CountryUZB: "UZB",
	CountryVAT: "VAT",
	CountryVCT: "VCT",
	CountryVEN: "VEN",
	CountryVGB: "VGB",
	CountryVIR: "VIR",
	CountryVNM: "VNM",
	CountryVUT: "VUT",
	CountryWLF: "WLF",
	CountryWSM: "WSM",
	CountryYEM: "YEM",
	CountryMYT: "MYT",
	CountryZAF: "ZAF",
	CountryZMB: "ZMB",
	CountryZWE: "ZWE",
}

var countryToCurrency = map[Country]Currency{
	CountryAD: CurrencyEUR,
	CountryAE: CurrencyAED,
	CountryAF: CurrencyAFN,
	CountryAG: CurrencyXCD,
	CountryAI: CurrencyXCD,
	CountryAL: CurrencyALL,
	CountryAM: CurrencyAMD,
	CountryAO: CurrencyAOA,
	CountryAQ: CurrencyUSD,
	CountryAR: CurrencyARS,
	CountryAS: CurrencyUSD,
	CountryAT: CurrencyEUR,
	CountryAU: CurrencyAUD,
	CountryAW: CurrencyAWG,
	CountryAX: CurrencyEUR,
	CountryAZ: CurrencyAZN,
	CountryBA: CurrencyBAM,
	CountryBB: CurrencyBBD,
	CountryBD: CurrencyBDT,
	CountryBE: CurrencyEUR,
	CountryBF: CurrencyXOF,
	CountryBG: CurrencyBGN,
	CountryBH: CurrencyBHD,
	CountryBI: CurrencyBIF,
	CountryBJ: CurrencyXOF,
	CountryBL: CurrencyEUR,
	CountryBM: CurrencyBMD,
	CountryBN: CurrencyBND,
	CountryBO: CurrencyBOB,
	CountryBQ: CurrencyUSD,
	CountryBR: CurrencyBRL,
	CountryBS: CurrencyBSD,
	CountryBT: CurrencyBTN,
	CountryBV: CurrencyNOK,
	CountryBW: CurrencyBWP,
	CountryBY: CurrencyBYN,
	CountryBZ: CurrencyBZD,
	CountryCA: CurrencyCAD,
	CountryCC: CurrencyAUD,
	CountryCD: CurrencyCDF,
	CountryCF: CurrencyXAF,
	CountryCG: CurrencyXAF,
	CountryCH: CurrencyCHF,
	CountryCI: CurrencyXOF,
	CountryCK: CurrencyNZD,
	CountryCL: CurrencyCLP,
	CountryCM: CurrencyXAF,
	CountryCN: CurrencyCNY,
	CountryCO: CurrencyCOP,
	CountryCR: CurrencyCRC,
	CountryCU: CurrencyCUP,
	CountryCV: CurrencyCVE,
	CountryCW: CurrencyANG,
	CountryCX: CurrencyAUD,
	CountryCY: CurrencyEUR,
	CountryCZ: CurrencyCZK,
	CountryDE: CurrencyEUR,
	CountryDJ: CurrencyDJF,
	CountryDK: CurrencyDKK,
	CountryDM: CurrencyXCD,
	CountryDO: CurrencyDOP,
	CountryDZ: CurrencyDZD,
	CountryEC: CurrencyUSD,
	CountryEE: CurrencyEUR,
	CountryEG: CurrencyEGP,
	CountryEH: CurrencyMAD,
	CountryER: CurrencyETB,
	CountryES: CurrencyEUR,
	CountryET: CurrencyETB,
	CountryFI: CurrencyEUR,
	CountryFJ: CurrencyFJD,
	CountryFK: CurrencyFKP,
	CountryFM: CurrencyUSD,
	CountryFO: CurrencyDKK,
	CountryFR: CurrencyEUR,
	CountryGA: CurrencyXAF,
	CountryGB: CurrencyGBP,
	CountryGD: CurrencyXCD,
	CountryGE: CurrencyGEL,
	CountryGF: CurrencyEUR,
	CountryGG: CurrencyGBP,
	CountryGH: CurrencyGHS,
	CountryGI: CurrencyGIP,
	CountryGL: CurrencyDKK,
	CountryGM: CurrencyGMD,
	CountryGN: CurrencyGNF,
	CountryGP: CurrencyEUR,
	CountryGQ: CurrencyXAF,
	CountryGR: CurrencyEUR,
	CountryGS: CurrencyGBP,
	CountryGT: CurrencyGTQ,
	CountryGU: CurrencyUSD,
	CountryGW: CurrencyXOF,
	CountryGY: CurrencyGYD,
	CountryHK: CurrencyHKD,
	CountryHM: CurrencyAUD,
	CountryHN: CurrencyHNL,
	CountryHR: CurrencyEUR,
	CountryHT: CurrencyHTG,
	CountryHU: CurrencyHUF,
	CountryID: CurrencyIDR,
	CountryIE: CurrencyEUR,
	CountryIL: CurrencyILS,
	CountryIM: CurrencyGBP,
	CountryIN: CurrencyINR,
	CountryIO: CurrencyUSD,
	CountryIQ: CurrencyIQD,
	CountryIR: CurrencyIRR,
	CountryIS: CurrencyISK,
	CountryIT: CurrencyEUR,
	CountryJE: CurrencyGBP,
	CountryJM: CurrencyJMD,
	CountryJO: CurrencyJOD,
	CountryJP: CurrencyJPY,
	CountryKE: CurrencyKES,
	CountryKG: CurrencyKGS,
	CountryKH: CurrencyKHR,
	CountryKI: CurrencyAUD,
	CountryKM: CurrencyKMF,
	CountryKN: CurrencyXCD,
	CountryKP: CurrencyKPW,
	CountryKR: CurrencyKRW,
	CountryKW: CurrencyKWD,
	CountryKY: CurrencyKYD,
	CountryKZ: CurrencyKZT,
	CountryLA: CurrencyLAK,
	CountryLB: CurrencyLBP,
	CountryLC: CurrencyXCD,
	CountryLI: CurrencyCHF,
	CountryLK: CurrencyLKR,
	CountryLR: CurrencyLRD,
	CountryLS: CurrencyLSL,
	CountryLT: CurrencyEUR,
	CountryLU: CurrencyEUR,
	CountryLV: CurrencyEUR,
	CountryLY: CurrencyLYD,
	CountryMA: CurrencyMAD,
	CountryMC: CurrencyEUR,
	CountryMD: CurrencyMDL,
	CountryME: CurrencyEUR,
	CountryMF: CurrencyEUR,
	CountryMG: CurrencyMGA,
	CountryMH: CurrencyUSD,
	CountryMK: CurrencyMKD,
	CountryML: CurrencyXOF,
	CountryMM: CurrencyMMK,
	CountryMN: CurrencyMNT,
	CountryMO: CurrencyMOP,
	CountryMP: CurrencyUSD,
	CountryMQ: CurrencyEUR,
	CountryMR: CurrencyMRU,
	CountryMS: CurrencyXCD,
	CountryMT: CurrencyEUR,
	CountryMU: CurrencyMUR,
	CountryMV: CurrencyMVR,
	CountryMW: CurrencyMWK,
	CountryMX: CurrencyMXN,
	CountryMY: CurrencyMYR,
	CountryMZ: CurrencyMZN,
	CountryNA: CurrencyNAD,
	CountryNC: CurrencyXPF,
	CountryNE: CurrencyXOF,
	CountryNF: CurrencyAUD,
	CountryNG: CurrencyNGN,
	CountryNI: CurrencyNIO,
	CountryNL: CurrencyEUR,
	CountryNO: CurrencyNOK,
	CountryNP: CurrencyNPR,
	CountryNR: CurrencyAUD,
	CountryNU: CurrencyNZD,
	CountryNZ: CurrencyNZD,
	CountryOM: CurrencyOMR,
	CountryPA: CurrencyPAB,
	CountryPE: CurrencyPEN,
	CountryPF: CurrencyXPF,
	CountryPG: CurrencyPGK,
	CountryPH: CurrencyPHP,
	CountryPK: CurrencyPKR,
	CountryPL: CurrencyPLN,
	CountryPM: CurrencyEUR,
	CountryPN: CurrencyNZD,
	CountryPR: CurrencyUSD,
	CountryPS: CurrencyILS,
	CountryPT: CurrencyEUR,
	CountryPW: CurrencyUSD,
	CountryPY: CurrencyPYG,
	CountryQA: CurrencyQAR,
	CountryRE: CurrencyEUR,
	CountryRO: CurrencyRON,
	CountryRS: CurrencyRSD,
	CountryRU: CurrencyRUB,
	CountryRW: CurrencyRWF,
	CountrySA: CurrencySAR,
	CountrySB: CurrencySBD,
	CountrySC: CurrencySCR,
	CountrySD: CurrencySDG,
	CountrySE: CurrencySEK,
	CountrySG: CurrencySGD,
	CountrySH: CurrencySHP,
	CountrySI: CurrencyEUR,
	CountrySJ: CurrencyNOK,
	CountrySK: CurrencyEUR,
	CountrySL: CurrencySLL,
	CountrySM: CurrencyEUR,
	CountrySN: CurrencyXOF,
	CountrySO: CurrencySOS,
	CountrySR: CurrencySRD,
	CountrySS: CurrencySSP,
	CountryST: CurrencySTD,
	CountrySV: CurrencyUSD,
	CountrySX: CurrencyANG,
	CountrySY: CurrencySYP,
	CountrySZ: CurrencySZL,
	CountryTC: CurrencyUSD,
	CountryTD: CurrencyXAF,
	CountryTF: CurrencyEUR,
	CountryTG: CurrencyXOF,
	CountryTH: CurrencyTHB,
	CountryTJ: CurrencyTJS,
	CountryTK: CurrencyNZD,
	CountryTL: CurrencyIDR,
	CountryTM: CurrencyTMT,
	CountryTN: CurrencyTND,
	CountryTO: CurrencyTOP,
	CountryTR: CurrencyTRY,
	CountryTT: CurrencyTTD,
	CountryTV: CurrencyAUD,
	CountryTW: CurrencyTWD,
	CountryTZ: CurrencyTZS,
	CountryUA: CurrencyUAH,
	CountryUG: CurrencyUGX,
	CountryUM: CurrencyUSD,
	CountryUS: CurrencyUSD,
	CountryUY: CurrencyUYU,
	CountryUZ: CurrencyUZS,
	CountryVA: CurrencyEUR,
	CountryVC: CurrencyXCD,
	CountryVE: CurrencyVES,
	CountryVG: CurrencyUSD,
	CountryVI: CurrencyUSD,
	CountryVN: CurrencyVND,
	CountryVU: CurrencyVUV,
	CountryWF: CurrencyXPF,
	CountryWS: CurrencyWST,
	CountryYE: CurrencyYER,
	CountryYT: CurrencyEUR,
	CountryZA: CurrencyZAR,
	CountryZM: CurrencyZMW,
	CountryZW: CurrencyUSD,
}

var alpha2ToCountry = map[string]Country{
	"AD": CountryAD,
	"AE": CountryAE,
	"AF": CountryAF,
	"AG": CountryAG,
	"AI": CountryAI,
	"AL": CountryAL,
	"AM": CountryAM,
	"AO": CountryAO,
	"AQ": CountryAQ,
	"AR": CountryAR,
	"AS": CountryAS,
	"AT": CountryAT,
	"AU": CountryAU,
	"AW": CountryAW,
	"AX": CountryAX,
	"AZ": CountryAZ,
	"BA": CountryBA,
	"BB": CountryBB,
	"BD": CountryBD,
	"BE": CountryBE,
	"BF": CountryBF,
	"BG": CountryBG,
	"BH": CountryBH,
	"BI": CountryBI,
	"BJ": CountryBJ,
	"BL": CountryBL,
	"BM": CountryBM,
	"BN": CountryBN,
	"BO": CountryBO,
	"BQ": CountryBQ,
	"BR": CountryBR,
	"BS": CountryBS,
	"BT": CountryBT,
	"BV": CountryBV,
	"BW": CountryBW,
	"BY": CountryBY,
	"BZ": CountryBZ,
	"CA": CountryCA,
	"CC": CountryCC,
	"CD": CountryCD,
	"CF": CountryCF,
	"CG": CountryCG,
	"CH": CountryCH,
	"CI": CountryCI,
	"CK": CountryCK,
	"CL": CountryCL,
	"CM": CountryCM,
	"CN": CountryCN,
	"CO": CountryCO,
	"CR": CountryCR,
	"CU": CountryCU,
	"CV": CountryCV,
	"CW": CountryCW,
	"CX": CountryCX,
	"CY": CountryCY,
	"CZ": CountryCZ,
	"DE": CountryDE,
	"DJ": CountryDJ,
	"DK": CountryDK,
	"DM": CountryDM,
	"DO": CountryDO,
	"DZ": CountryDZ,
	"EC": CountryEC,
	"EE": CountryEE,
	"EG": CountryEG,
	"EH": CountryEH,
	"ER": CountryER,
	"ES": CountryES,
	"ET": CountryET,
	"FI": CountryFI,
	"FJ": CountryFJ,
	"FK": CountryFK,
	"FM": CountryFM,
	"FO": CountryFO,
	"FR": CountryFR,
	"GA": CountryGA,
	"GB": CountryGB,
	"GD": CountryGD,
	"GE": CountryGE,
	"GF": CountryGF,
	"GG": CountryGG,
	"GH": CountryGH,
	"GI": CountryGI,
	"GL": CountryGL,
	"GM": CountryGM,
	"GN": CountryGN,
	"GP": CountryGP,
	"GQ": CountryGQ,
	"GR": CountryGR,
	"GS": CountryGS,
	"GT": CountryGT,
	"GU": CountryGU,
	"GW": CountryGW,
	"GY": CountryGY,
	"HK": CountryHK,
	"HM": CountryHM,
	"HN": CountryHN,
	"HR": CountryHR,
	"HT": CountryHT,
	"HU": CountryHU,
	"ID": CountryID,
	"IE": CountryIE,
	"IL": CountryIL,
	"IM": CountryIM,
	"IN": CountryIN,
	"IO": CountryIO,
	"IQ": CountryIQ,
	"IR": CountryIR,
	"IS": CountryIS,
	"IT": CountryIT,
	"JE": CountryJE,
	"JM": CountryJM,
	"JO": CountryJO,
	"JP": CountryJP,
	"KE": CountryKE,
	"KG": CountryKG,
	"KH": CountryKH,
	"KI": CountryKI,
	"KM": CountryKM,
	"KN": CountryKN,
	"KP": CountryKP,
	"KR": CountryKR,
	"KW": CountryKW,
	"KY": CountryKY,
	"KZ": CountryKZ,
	"LA": CountryLA,
	"LB": CountryLB,
	"LC": CountryLC,
	"LI": CountryLI,
	"LK": CountryLK,
	"LR": CountryLR,
	"LS": CountryLS,
	"LT": CountryLT,
	"LU": CountryLU,
	"LV": CountryLV,
	"LY": CountryLY,
	"MA": CountryMA,
	"MC": CountryMC,
	"MD": CountryMD,
	"ME": CountryME,
	"MF": CountryMF,
	"MG": CountryMG,
	"MH": CountryMH,
	"MK": CountryMK,
	"ML": CountryML,
	"MM": CountryMM,
	"MN": CountryMN,
	"MO": CountryMO,
	"MP": CountryMP,
	"MQ": CountryMQ,
	"MR": CountryMR,
	"MS": CountryMS,
	"MT": CountryMT,
	"MU": CountryMU,
	"MV": CountryMV,
	"MW": CountryMW,
	"MX": CountryMX,
	"MY": CountryMY,
	"MZ": CountryMZ,
	"NA": CountryNA,
	"NC": CountryNC,
	"NE": CountryNE,
	"NF": CountryNF,
	"NG": CountryNG,
	"NI": CountryNI,
	"NL": CountryNL,
	"NO": CountryNO,
	"NP": CountryNP,
	"NR": CountryNR,
	"NU": CountryNU,
	"NZ": CountryNZ,
	"OM": CountryOM,
	"PA": CountryPA,
	"PE": CountryPE,
	"PF": CountryPF,
	"PG": CountryPG,
	"PH": CountryPH,
	"PK": CountryPK,
	"PL": CountryPL,
	"PM": CountryPM,
	"PN": CountryPN,
	"PR": CountryPR,
	"PS": CountryPS,
	"PT": CountryPT,
	"PW": CountryPW,
	"PY": CountryPY,
	"QA": CountryQA,
	"RE": CountryRE,
	"RO": CountryRO,
	"RS": CountryRS,
	"RU": CountryRU,
	"RW": CountryRW,
	"SA": CountrySA,
	"SB": CountrySB,
	"SC": CountrySC,
	"SD": CountrySD,
	"SE": CountrySE,
	"SG": CountrySG,
	"SH": CountrySH,
	"SI": CountrySI,
	"SJ": CountrySJ,
	"SK": CountrySK,
	"SL": CountrySL,
	"SM": CountrySM,
	"SN": CountrySN,
	"SO": CountrySO,
	"SR": CountrySR,
	"SS": CountrySS,
	"ST": CountryST,
	"SV": CountrySV,
	"SX": CountrySX,
	"SY": CountrySY,
	"SZ": CountrySZ,
	"TC": CountryTC,
	"TD": CountryTD,
	"TF": CountryTF,
	"TG": CountryTG,
	"TH": CountryTH,
	"TJ": CountryTJ,
	"TK": CountryTK,
	"TL": CountryTL,
	"TM": CountryTM,
	"TN": CountryTN,
	"TO": CountryTO,
	"TR": CountryTR,
	"TT": CountryTT,
	"TV": CountryTV,
	"TW": CountryTW,
	"TZ": CountryTZ,
	"UA": CountryUA,
	"UG": CountryUG,
	"UM": CountryUM,
	"US": CountryUS,
	"UY": CountryUY,
	"UZ": CountryUZ,
	"VA": CountryVA,
	"VC": CountryVC,
	"VE": CountryVE,
	"VG": CountryVG,
	"VI": CountryVI,
	"VN": CountryVN,
	"VU": CountryVU,
	"WF": CountryWF,
	"WS": CountryWS,
	"YE": CountryYE,
	"YT": CountryYT,
	"ZA": CountryZA,
	"ZM": CountryZM,
	"ZW": CountryZW,
}

var alpha3ToCountry = map[string]Country{
	"AND": CountryAND,
	"ARE": CountryARE,
	"AFG": CountryAFG,
	"ATG": CountryATG,
	"AIA": CountryAIA,
	"ALB": CountryALB,
	"ARM": CountryARM,
	"AGO": CountryAGO,
	"ATA": CountryATA,
	"ARG": CountryARG,
	"ASM": CountryASM,
	"AUT": CountryAUT,
	"AUS": CountryAUS,
	"ABW": CountryABW,
	"ALA": CountryALA,
	"AZE": CountryAZE,
	"BIH": CountryBIH,
	"BRB": CountryBRB,
	"BGD": CountryBGD,
	"BEL": CountryBEL,
	"BFA": CountryBFA,
	"BGR": CountryBGR,
	"BHR": CountryBHR,
	"BDI": CountryBDI,
	"BEN": CountryBEN,
	"BLM": CountryBLM,
	"BMU": CountryBMU,
	"BRN": CountryBRN,
	"BOL": CountryBOL,
	"BES": CountryBES,
	"BRA": CountryBRA,
	"BHS": CountryBHS,
	"BTN": CountryBTN,
	"BVT": CountryBVT,
	"BWA": CountryBWA,
	"BLR": CountryBLR,
	"BLZ": CountryBLZ,
	"CAN": CountryCAN,
	"CCK": CountryCCK,
	"COD": CountryCOD,
	"CAF": CountryCAF,
	"COG": CountryCOG,
	"CHE": CountryCHE,
	"CIV": CountryCIV,
	"COK": CountryCOK,
	"CHL": CountryCHL,
	"CMR": CountryCMR,
	"CHN": CountryCHN,
	"COL": CountryCOL,
	"CRI": CountryCRI,
	"CUB": CountryCUB,
	"CPV": CountryCPV,
	"CUW": CountryCUW,
	"CXR": CountryCXR,
	"CYP": CountryCYP,
	"CZE": CountryCZE,
	"DEU": CountryDEU,
	"DJI": CountryDJI,
	"DNK": CountryDNK,
	"DMA": CountryDMA,
	"DOM": CountryDOM,
	"DZA": CountryDZA,
	"ECU": CountryECU,
	"EST": CountryEST,
	"EGY": CountryEGY,
	"ESH": CountryESH,
	"ERI": CountryERI,
	"ESP": CountryESP,
	"ETH": CountryETH,
	"FIN": CountryFIN,
	"FJI": CountryFJI,
	"FLK": CountryFLK,
	"FSM": CountryFSM,
	"FRO": CountryFRO,
	"FRA": CountryFRA,
	"GAB": CountryGAB,
	"GBR": CountryGBR,
	"GRD": CountryGRD,
	"GEO": CountryGEO,
	"GUF": CountryGUF,
	"GGY": CountryGGY,
	"GHA": CountryGHA,
	"GIB": CountryGIB,
	"GRL": CountryGRL,
	"GMB": CountryGMB,
	"GIN": CountryGIN,
	"GLP": CountryGLP,
	"GNQ": CountryGNQ,
	"GRC": CountryGRC,
	"SGS": CountrySGS,
	"GTM": CountryGTM,
	"GUM": CountryGUM,
	"GNB": CountryGNB,
	"GUY": CountryGUY,
	"HKG": CountryHKG,
	"HMD": CountryHMD,
	"HND": CountryHND,
	"HRV": CountryHRV,
	"HTI": CountryHTI,
	"HUN": CountryHUN,
	"IDN": CountryIDN,
	"IRL": CountryIRL,
	"ISR": CountryISR,
	"IMN": CountryIMN,
	"IND": CountryIND,
	"IOT": CountryIOT,
	"IRQ": CountryIRQ,
	"IRN": CountryIRN,
	"ISL": CountryISL,
	"ITA": CountryITA,
	"JEY": CountryJEY,
	"JAM": CountryJAM,
	"JOR": CountryJOR,
	"JPN": CountryJPN,
	"KEN": CountryKEN,
	"KGZ": CountryKGZ,
	"KHM": CountryKHM,
	"KIR": CountryKIR,
	"COM": CountryCOM,
	"KNA": CountryKNA,
	"PRK": CountryPRK,
	"KOR": CountryKOR,
	"KWT": CountryKWT,
	"CYM": CountryCYM,
	"KAZ": CountryKAZ,
	"LAO": CountryLAO,
	"LBN": CountryLBN,
	"LCA": CountryLCA,
	"LIE": CountryLIE,
	"LKA": CountryLKA,
	"LBR": CountryLBR,
	"LSO": CountryLSO,
	"LTU": CountryLTU,
	"LUX": CountryLUX,
	"LVA": CountryLVA,
	"LBY": CountryLBY,
	"MAR": CountryMAR,
	"MCO": CountryMCO,
	"MDA": CountryMDA,
	"MNE": CountryMNE,
	"MAF": CountryMAF,
	"MDG": CountryMDG,
	"MHL": CountryMHL,
	"MKD": CountryMKD,
	"MLI": CountryMLI,
	"MMR": CountryMMR,
	"MNG": CountryMNG,
	"MAC": CountryMAC,
	"MNP": CountryMNP,
	"MTQ": CountryMTQ,
	"MRT": CountryMRT,
	"MSR": CountryMSR,
	"MLT": CountryMLT,
	"MUS": CountryMUS,
	"MDV": CountryMDV,
	"MWI": CountryMWI,
	"MEX": CountryMEX,
	"MYS": CountryMYS,
	"MOZ": CountryMOZ,
	"NAM": CountryNAM,
	"NCL": CountryNCL,
	"NER": CountryNER,
	"NFK": CountryNFK,
	"NGA": CountryNGA,
	"NIC": CountryNIC,
	"NLD": CountryNLD,
	"NOR": CountryNOR,
	"NPL": CountryNPL,
	"NRU": CountryNRU,
	"NIU": CountryNIU,
	"NZL": CountryNZL,
	"OMN": CountryOMN,
	"PAN": CountryPAN,
	"PER": CountryPER,
	"PYF": CountryPYF,
	"PNG": CountryPNG,
	"PHL": CountryPHL,
	"PAK": CountryPAK,
	"POL": CountryPOL,
	"SPM": CountrySPM,
	"PCN": CountryPCN,
	"PRI": CountryPRI,
	"PSE": CountryPSE,
	"PRT": CountryPRT,
	"PLW": CountryPLW,
	"PRY": CountryPRY,
	"QAT": CountryQAT,
	"REU": CountryREU,
	"ROU": CountryROU,
	"SRB": CountrySRB,
	"RUS": CountryRUS,
	"RWA": CountryRWA,
	"SAU": CountrySAU,
	"SLB": CountrySLB,
	"SYC": CountrySYC,
	"SDN": CountrySDN,
	"SWE": CountrySWE,
	"SGP": CountrySGP,
	"SHN": CountrySHN,
	"SVN": CountrySVN,
	"SJM": CountrySJM,
	"SVK": CountrySVK,
	"SLE": CountrySLE,
	"SMR": CountrySMR,
	"SEN": CountrySEN,
	"SOM": CountrySOM,
	"SUR": CountrySUR,
	"SSD": CountrySSD,
	"STP": CountrySTP,
	"SLV": CountrySLV,
	"SXM": CountrySXM,
	"SYR": CountrySYR,
	"SWZ": CountrySWZ,
	"TCA": CountryTCA,
	"TCD": CountryTCD,
	"ATF": CountryATF,
	"TGO": CountryTGO,
	"THA": CountryTHA,
	"TJK": CountryTJK,
	"TKL": CountryTKL,
	"TLS": CountryTLS,
	"TKM": CountryTKM,
	"TUN": CountryTUN,
	"TON": CountryTON,
	"TUR": CountryTUR,
	"TTO": CountryTTO,
	"TUV": CountryTUV,
	"TWN": CountryTWN,
	"TZA": CountryTZA,
	"UKR": CountryUKR,
	"UGA": CountryUGA,
	"UMI": CountryUMI,
	"USA": CountryUSA,
	"URY": CountryURY,
	"UZB": CountryUZB,
	"VAT": CountryVAT,
	"VCT": CountryVCT,
	"VEN": CountryVEN,
	"VGB": CountryVGB,
	"VIR": CountryVIR,
	"VNM": CountryVNM,
	"VUT": CountryVUT,
	"WLF": CountryWLF,
	"WSM": CountryWSM,
	"YEM": CountryYEM,
	"MYT": CountryMYT,
	"ZAF": CountryZAF,
	"ZMB": CountryZMB,
	"ZWE": CountryZWE,
}

var countryToName = map[Country]string{
	CountryAD: "Andorra",
	CountryAE: "United Arab Emirates",
	CountryAF: "Afghanistan",
	CountryAG: "Antigua and Barbuda",
	CountryAI: "Anguilla",
	CountryAL: "Albania",
	CountryAM: "Armenia",
	CountryAO: "Angola",
	CountryAQ: "Antarctica",
	CountryAR: "Argentina",
	CountryAS: "American Samoa",
	CountryAT: "Austria",
	CountryAU: "Australia",
	CountryAW: "Aruba",
	CountryAX: "Åland Islands",
	CountryAZ: "Azerbaijan",
	CountryBA: "Bosnia and Herzegovina",
	CountryBB: "Barbados",
	CountryBD: "Bangladesh",
	CountryBE: "Belgium",
	CountryBF: "Burkina Faso",
	CountryBG: "Bulgaria",
	CountryBH: "Bahrain",
	CountryBI: "Burundi",
	CountryBJ: "Benin",
	CountryBL: "Saint Barthélemy",
	CountryBM: "Bermuda",
	CountryBN: "Brunei Darussalam",
	CountryBO: "Bolivia (Plurinational State of)",
	CountryBQ: "Bonaire, Sint Eustatius and Saba",
	CountryBR: "Brazil",
	CountryBS: "Bahamas",
	CountryBT: "Bhutan",
	CountryBV: "Bouvet Island",
	CountryBW: "Botswana",
	CountryBY: "Belarus",
	CountryBZ: "Belize",
	CountryCA: "Canada",
	CountryCC: "Cocos (Keeling) Islands",
	CountryCD: "Congo (Democratic Republic of the)",
	CountryCF: "Central African Republic",
	CountryCG: "Congo",
	CountryCH: "Switzerland",
	CountryCI: "Côte d'Ivoire",
	CountryCK: "Cook Islands",
	CountryCL: "Chile",
	CountryCM: "Cameroon",
	CountryCN: "China",
	CountryCO: "Colombia",
	CountryCR: "Costa Rica",
	CountryCU: "Cuba",
	CountryCV: "Cabo Verde",
	CountryCW: "Curaçao",
	CountryCX: "Christmas Island",
	CountryCY: "Cyprus",
	CountryCZ: "Czechia",
	CountryDE: "Germany",
	CountryDJ: "Djibouti",
	CountryDK: "Denmark",
	CountryDM: "Dominica",
	CountryDO: "Dominican Republic",
	CountryDZ: "Algeria",
	CountryEC: "Ecuador",
	CountryEE: "Estonia",
	CountryEG: "Egypt",
	CountryEH: "Western Sahara",
	CountryER: "Eritrea",
	CountryES: "Spain",
	CountryET: "Ethiopia",
	CountryFI: "Finland",
	CountryFJ: "Fiji",
	CountryFK: "Falkland Islands (Malvinas)",
	CountryFM: "Micronesia (Federated States of)",
	CountryFO: "Faroe Islands",
	CountryFR: "France",
	CountryGA: "Gabon",
	CountryGB: "United Kingdom of Great Britain and Northern Ireland",
	CountryGD: "Grenada",
	CountryGE: "Georgia",
	CountryGF: "French Guiana",
	CountryGG: "Guernsey",
	CountryGH: "Ghana",
	CountryGI: "Gibraltar",
	CountryGL: "Greenland",
	CountryGM: "Gambia",
	CountryGN: "Guinea",
	CountryGP: "Guadeloupe",
	CountryGQ: "Equatorial Guinea",
	CountryGR: "Greece",
	CountryGS: "South Georgia and the South Sandwich Islands",
	CountryGT: "Guatemala",
	CountryGU: "Guam",
	CountryGW: "Guinea-Bissau",
	CountryGY: "Guyana",
	CountryHK: "Hong Kong",
	CountryHM: "Heard Island and McDonald Islands",
	CountryHN: "Honduras",
	CountryHR: "Croatia",
	CountryHT: "Haiti",
	CountryHU: "Hungary",
	CountryID: "Indonesia",
	CountryIE: "Ireland",
	CountryIL: "Israel",
	CountryIM: "Isle of Man",
	CountryIN: "India",
	CountryIO: "British Indian Ocean Territory",
	CountryIQ: "Iraq",
	CountryIR: "Iran (Islamic Republic of)",
	CountryIS: "Iceland",
	CountryIT: "Italy",
	CountryJE: "Jersey",
	CountryJM: "Jamaica",
	CountryJO: "Jordan",
	CountryJP: "Japan",
	CountryKE: "Kenya",
	CountryKG: "Kyrgyzstan",
	CountryKH: "Cambodia",
	CountryKI: "Kiribati",
	CountryKM: "Comoros",
	CountryKN: "Saint Kitts and Nevis",
	CountryKP: "Korea (Democratic People's Republic of)",
	CountryKR: "Korea (Republic of)",
	CountryKW: "Kuwait",
	CountryKY: "Cayman Islands",
	CountryKZ: "Kazakhstan",
	CountryLA: "Lao People's Democratic Republic",
	CountryLB: "Lebanon",
	CountryLC: "Saint Lucia",
	CountryLI: "Liechtenstein",
	CountryLK: "Sri Lanka",
	CountryLR: "Liberia",
	CountryLS: "Lesotho",
	CountryLT: "Lithuania",
	CountryLU: "Luxembourg",
	CountryLV: "Latvia",
	CountryLY: "Libya",
	CountryMA: "Morocco",
	CountryMC: "Monaco",
	CountryMD: "Moldova (Republic of)",
	CountryME: "Montenegro",
	CountryMF: "Saint Martin (French part)",
	CountryMG: "Madagascar",
	CountryMH: "Marshall Islands",
	CountryMK: "North Macedonia",
	CountryML: "Mali",
	CountryMM: "Myanmar",
	CountryMN: "Mongolia",
	CountryMO: "Macao",
	CountryMP: "Northern Mariana Islands",
	CountryMQ: "Martinique",
	CountryMR: "Mauritania",
	CountryMS: "Montserrat",
	CountryMT: "Malta",
	CountryMU: "Mauritius",
	CountryMV: "Maldives",
	CountryMW: "Malawi",
	CountryMX: "Mexico",
	CountryMY: "Malaysia",
	CountryMZ: "Mozambique",
	CountryNA: "Namibia",
	CountryNC: "New Caledonia",
	CountryNE: "Niger",
	CountryNF: "Norfolk Island",
	CountryNG: "Nigeria",
	CountryNI: "Nicaragua",
	CountryNL: "Netherlands",
	CountryNO: "Norway",
	CountryNP: "Nepal",
	CountryNR: "Nauru",
	CountryNU: "Niue",
	CountryNZ: "New Zealand",
	CountryOM: "Oman",
	CountryPA: "Panama",
	CountryPE: "Peru",
	CountryPF: "French Polynesia",
	CountryPG: "Papua New Guinea",
	CountryPH: "Philippines",
	CountryPK: "Pakistan",
	CountryPL: "Poland",
	CountryPM: "Saint Pierre and Miquelon",
	CountryPN: "Pitcairn",
	CountryPR: "Puerto Rico",
	CountryPS: "Palestine, State of",
	CountryPT: "Portugal",
	CountryPW: "Palau",
	CountryPY: "Paraguay",
	CountryQA: "Qatar",
	CountryRE: "Réunion",
	CountryRO: "Romania",
	CountryRS: "Serbia",
	CountryRU: "Russian Federation",
	CountryRW: "Rwanda",
	CountrySA: "Saudi Arabia",
	CountrySB: "Solomon Islands",
	CountrySC: "Seychelles",
	CountrySD: "Sudan",
	CountrySE: "Sweden",
	CountrySG: "Singapore",
	CountrySH: "Saint Helena, Ascension and Tristan da Cunha",
	CountrySI: "Slovenia",
	CountrySJ: "Svalbard and Jan Mayen",
	CountrySK: "Slovakia",
	CountrySL: "Sierra Leone",
	CountrySM: "San Marino",
	CountrySN: "Senegal",
	CountrySO: "Somalia",
	CountrySR: "Suriname",
	CountrySS: "South Sudan",
	CountryST: "Sao Tome and Principe",
	CountrySV: "El Salvador",
	CountrySX: "Sint Maarten (Dutch part)",
	CountrySY: "Syrian Arab Republic",
	CountrySZ: "Eswatini",
	CountryTC: "Turks and Caicos Islands",
	CountryTD: "Chad",
	CountryTF: "French Southern Territories",
	CountryTG: "Togo",
	CountryTH: "Thailand",
	CountryTJ: "Tajikistan",
	CountryTK: "Tokelau",
	CountryTL: "Timor-Leste",
	CountryTM: "Turkmenistan",
	CountryTN: "Tunisia",
	CountryTO: "Tonga",
	CountryTR: "Türkiye",
	CountryTT: "Trinidad and Tobago",
	CountryTV: "Tuvalu",
	CountryTW: "Taiwan, Province of China",
	CountryTZ: "Tanzania, United Republic of",
	CountryUA: "Ukraine",
	CountryUG: "Uganda",
	CountryUM: "United States Minor Outlying Islands",
	CountryUS: "United States of America",
	CountryUY: "Uruguay",
	CountryUZ: "Uzbekistan",
	CountryVA: "Holy See",
	CountryVC: "Saint Vincent and the Grenadines",
	CountryVE: "Venezuela (Bolivarian Republic of)",
	CountryVG: "Virgin Islands (British)",
	CountryVI: "Virgin Islands (U.S.)",
	CountryVN: "Viet Nam",
	CountryVU: "Vanuatu",
	CountryWF: "Wallis and Futuna",
	CountryWS: "Samoa",
	CountryYE: "Yemen",
	CountryYT: "Mayotte",
	CountryZA: "South Africa",
	CountryZM: "Zambia",
	CountryZW: "Zimbabwe",
}

var countryToLongName = map[Country]string{
	CountryAD: "The Principality of Andorra",
	CountryAE: "The United Arab Emirates",
	CountryAF: "The Islamic Republic of Afghanistan",
	CountryAG: "Antigua and Barbuda",
	CountryAI: "Anguilla",
	CountryAL: "The Republic of Albania",
	CountryAM: "The Republic of Armenia",
	CountryAO: "The Republic of Angola",
	CountryAQ: "Antarctica",
	CountryAR: "The Argentine Republic",
	CountryAS: "The Territory of American Samoa",
	CountryAT: "The Republic of Austria",
	CountryAU: "The Commonwealth of Australia",
	CountryAW: "Aruba",
	CountryAX: "Åland",
	CountryAZ: "The Republic of Azerbaijan",
	CountryBA: "Bosnia and Herzegovina",
	CountryBB: "Barbados",
	CountryBD: "The People's Republic of Bangladesh",
	CountryBE: "The Kingdom of Belgium",
	CountryBF: "Burkina Faso",
	CountryBG: "The Republic of Bulgaria",
	CountryBH: "The Kingdom of Bahrain",
	CountryBI: "The Republic of Burundi",
	CountryBJ: "The Republic of Benin",
	CountryBL: "The Collectivity of Saint-Barthélemy",
	CountryBM: "Bermuda",
	CountryBN: "The Nation of Brunei, the Abode of Peace",
	CountryBO: "The Plurinational State of Bolivia",
	CountryBQ: "Bonaire, Sint Eustatius and Saba",
	CountryBR: "The Federative Republic of Brazil",
	CountryBS: "The Commonwealth of The Bahamas",
	CountryBT: "The Kingdom of Bhutan",
	CountryBV: "Bouvet Island",
	CountryBW: "The Republic of Botswana",
	CountryBY: "The Republic of Belarus",
	CountryBZ: "Belize",
	CountryCA: "Canada",
	CountryCC: "The Territory of Cocos (Keeling) Islands",
	CountryCD: "The Democratic Republic of the Congo",
	CountryCF: "The Central African Republic",
	CountryCG: "The Republic of the Congo",
	CountryCH: "The Swiss Confederation",
	CountryCI: "The Republic of Côte d'Ivoire",
	CountryCK: "The Cook Islands",
	CountryCL: "The Republic of Chile",
	CountryCM: "The Republic of Cameroon",
	CountryCN: "The People's Republic of China",
	CountryCO: "The Republic of Colombia",
	CountryCR: "The Republic of Costa Rica",
	CountryCU: "The Republic of Cuba",
	CountryCV: "The Republic of Cabo Verde",
	CountryCW: "The Country of Curaçao",
	CountryCX: "The Territory of Christmas Island",
	CountryCY: "The Republic of Cyprus",
	CountryCZ: "The Czech Republic",
	CountryDE: "The Federal Republic of Germany",
	CountryDJ: "The Republic of Djibouti",
	CountryDK: "The Kingdom of Denmark",
	CountryDM: "The Commonwealth of Dominica",
	CountryDO: "The Dominican Republic",
	CountryDZ: "The People's Democratic Republic of Algeria",
	CountryEC: "The Republic of Ecuador",
	CountryEE: "The Republic of Estonia",
	CountryEG: "The Arab Republic of Egypt",
	CountryEH: "The Sahrawi Arab Democratic Republic",
	CountryER: "The State of Eritrea",
	CountryES: "The Kingdom of Spain",
	CountryET: "The Federal Democratic Republic of Ethiopia",
	CountryFI: "The Republic of Finland",
	CountryFJ: "The Republic of Fiji",
	CountryFK: "The Falkland Islands",
	CountryFM: "The Federated States of Micronesia",
	CountryFO: "The Faroe Islands",
	CountryFR: "The French Republic",
	CountryGA: "The Gabonese Republic",
	CountryGB: "The United Kingdom of Great Britain and Northern Ireland",
	CountryGD: "Grenada",
	CountryGE: "Georgia",
	CountryGF: "Guyane",
	CountryGG: "The Bailiwick of Guernsey",
	CountryGH: "The Republic of Ghana",
	CountryGI: "Gibraltar",
	CountryGL: "Kalaallit Nunaat",
	CountryGM: "The Republic of The Gambia",
	CountryGN: "The Republic of Guinea",
	CountryGP: "Guadeloupe",
	CountryGQ: "The Republic of Equatorial Guinea",
	CountryGR: "The Hellenic Republic",
	CountryGS: "South Georgia and the South Sandwich Islands",
	CountryGT: "The Republic of Guatemala",
	CountryGU: "The Territory of Guam",
	CountryGW: "The Republic of Guinea-Bissau",
	CountryGY: "The Co-operative Republic of Guyana",
	CountryHK: "The Hong Kong Special Administrative Region of China",
	CountryHM: "The Territory of Heard Island and McDonald Islands",
	CountryHN: "The Republic of Honduras",
	CountryHR: "The Republic of Croatia",
	CountryHT: "The Republic of Haiti",
	CountryHU: "Hungary",
	CountryID: "The Republic of Indonesia",
	CountryIE: "Ireland",
	CountryIL: "The State of Israel",
	CountryIM: "The Isle of Man",
	CountryIN: "The Republic of India",
	CountryIO: "The British Indian Ocean Territory",
	CountryIQ: "The Republic of Iraq",
	CountryIR: "The Islamic Republic of Iran",
	CountryIS: "Iceland",
	CountryIT: "The Italian Republic",
	CountryJE: "The Bailiwick of Jersey",
	CountryJM: "Jamaica",
	CountryJO: "The Hashemite Kingdom of Jordan",
	CountryJP: "Japan",
	CountryKE: "The Republic of Kenya",
	CountryKG: "The Kyrgyz Republic",
	CountryKH: "The Kingdom of Cambodia",
	CountryKI: "The Republic of Kiribati",
	CountryKM: "The Union of the Comoros",
	CountryKN: "Saint Kitts and Nevis",
	CountryKP: "The Democratic People's Republic of Korea",
	CountryKR: "The Republic of Korea",
	CountryKW: "The State of Kuwait",
	CountryKY: "The Cayman Islands",
	CountryKZ: "The Republic of Kazakhstan",
	CountryLA: "The Lao People's Democratic Republic",
	CountryLB: "The Lebanese Republic",
	CountryLC: "Saint Lucia",
	CountryLI: "The Principality of Liechtenstein",
	CountryLK: "The Democratic Socialist Republic of Sri Lanka",
	CountryLR: "The Republic of Liberia",
	CountryLS: "The Kingdom of Lesotho",
	CountryLT: "The Republic of Lithuania",
	CountryLU: "The Grand Duchy of Luxembourg",
	CountryLV: "The Republic of Latvia",
	CountryLY: "The State of Libya",
	CountryMA: "The Kingdom of Morocco",
	CountryMC: "The Principality of Monaco",
	CountryMD: "The Republic of Moldova",
	CountryME: "Montenegro",
	CountryMF: "The Collectivity of Saint-Martin",
	CountryMG: "The Republic of Madagascar",
	CountryMH: "The Republic of the Marshall Islands",
	CountryMK: "The Republic of North Macedonia",
	CountryML: "The Republic of Mali",
	CountryMM: "The Republic of the Union of Myanmar",
	CountryMN: "Mongolia",
	CountryMO: "The Macao Special Administrative Region of China",
	CountryMP: "The Commonwealth of the Northern Mariana Islands",
	CountryMQ: "Martinique",
	CountryMR: "The Islamic Republic of Mauritania",
	CountryMS: "Montserrat",
	CountryMT: "The Republic of Malta",
	CountryMU: "The Republic of Mauritius",
	CountryMV: "The Republic of Maldives",
	CountryMW: "The Republic of Malawi",
	CountryMX: "The United Mexican States",
	CountryMY: "Malaysia",
	CountryMZ: "The Republic of Mozambique",
	CountryNA: "The Republic of Namibia",
	CountryNC: "New Caledonia",
	CountryNE: "The Republic of the Niger",
	CountryNF: "The Territory of Norfolk Island",
	CountryNG: "The Federal Republic of Nigeria",
	CountryNI: "The Republic of Nicaragua",
	CountryNL: "The Kingdom of the Netherlands",
	CountryNO: "The Kingdom of Norway",
	CountryNP: "The Federal Democratic Republic of Nepal",
	CountryNR: "The Republic of Nauru",
	CountryNU: "Niue",
	CountryNZ: "New Zealand",
	CountryOM: "The Sultanate of Oman",
	CountryPA: "The Republic of Panamá",
	CountryPE: "The Republic of Perú",
	CountryPF: "French Polynesia",
	CountryPG: "The Independent State of Papua New Guinea",
	CountryPH: "The Republic of the Philippines",
	CountryPK: "The Islamic Republic of Pakistan",
	CountryPL: "The Republic of Poland",
	CountryPM: "The Overseas Collectivity of Saint-Pierre and Miquelon",
	CountryPN: "The Pitcairn, Henderson, Ducie and Oeno Islands",
	CountryPR: "The Commonwealth of Puerto Rico",
	CountryPS: "The State of Palestine",
	CountryPT: "The Portuguese Republic",
	CountryPW: "The Republic of Palau",
	CountryPY: "The Republic of Paraguay",
	CountryQA: "The State of Qatar",
	CountryRE: "Réunion",
	CountryRO: "Romania",
	CountryRS: "The Republic of Serbia",
	CountryRU: "The Russian Federation",
	CountryRW: "The Republic of Rwanda",
	CountrySA: "The Kingdom of Saudi Arabia",
	CountrySB: "The Solomon Islands",
	CountrySC: "The Republic of Seychelles",
	CountrySD: "The Republic of the Sudan",
	CountrySE: "The Kingdom of Sweden",
	CountrySG: "The Republic of Singapore",
	CountrySH: "Saint Helena, Ascension and Tristan da Cunha",
	CountrySI: "The Republic of Slovenia",
	CountrySJ: "Svalbard and Jan Mayen",
	CountrySK: "The Slovak Republic",
	CountrySL: "The Republic of Sierra Leone",
	CountrySM: "The Republic of San Marino",
	CountrySN: "The Republic of Senegal",
	CountrySO: "The Federal Republic of Somalia",
	CountrySR: "The Republic of Suriname",
	CountrySS: "The Republic of South Sudan",
	CountryST: "The Democratic Republic of São Tomé and Príncipe",
	CountrySV: "The Republic of El Salvador",
	CountrySX: "Sint Maarten",
	CountrySY: "The Syrian Arab Republic",
	CountrySZ: "The Kingdom of Eswatini",
	CountryTC: "The Turks and Caicos Islands",
	CountryTD: "The Republic of Chad",
	CountryTF: "The French Southern and Antarctic Lands",
	CountryTG: "The Togolese Republic",
	CountryTH: "The Kingdom of Thailand",
	CountryTJ: "The Republic of Tajikistan",
	CountryTK: "Tokelau",
	CountryTL: "The Democratic Republic of Timor-Leste",
	CountryTM: "Turkmenistan",
	CountryTN: "The Republic of Tunisia",
	CountryTO: "The Kingdom of Tonga",
	CountryTR: "The Republic of Türkiye",
	CountryTT: "The Republic of Trinidad and Tobago",
	CountryTV: "Tuvalu",
	CountryTW: "Taiwan, Province of China",
	CountryTZ: "The United Republic of Tanzania",
	CountryUA: "Ukraine",
	CountryUG: "The Republic of Uganda",
	CountryUM: "United States Minor Outlying Islands",
	CountryUS: "The United States of America",
	CountryUY: "The Oriental Republic of Uruguay",
	CountryUZ: "The Republic of Uzbekistan",
	CountryVA: "The Holy See",
	CountryVC: "Saint Vincent and the Grenadines",
	CountryVE: "The Bolivarian Republic of Venezuela",
	CountryVG: "The Virgin Islands",
	CountryVI: "The Virgin Islands of the United States",
	CountryVN: "The Socialist Republic of Viet Nam",
	CountryVU: "The Republic of Vanuatu",
	CountryWF: "The Territory of the Wallis and Futuna Islands",
	CountryWS: "The Independent State of Samoa",
	CountryYE: "The Republic of Yemen",
	CountryYT: "The Department of Mayotte",
	CountryZA: "The Republic of South Africa",
	CountryZM: "The Republic of Zambia",
	CountryZW: "The Republic of Zimbabwe",
}

var countryToPhonePrefix = map[Country]string{
	CountryAD: "376",
	CountryAE: "971",
	CountryAF: "93",
	CountryAG: "1",
	CountryAI: "1",
	CountryAL: "355",
	CountryAM: "374",
	CountryAO: "244",
	CountryAQ: "672",
	CountryAR: "54",
	CountryAS: "1",
	CountryAT: "43",
	CountryAU: "61",
	CountryAW: "297",
	CountryAX: "358",
	CountryAZ: "994",
	CountryBA: "387",
	CountryBB: "1",
	CountryBD: "880",
	CountryBE: "32",
	CountryBF: "226",
	CountryBG: "359",
	CountryBH: "973",
	CountryBI: "257",
	CountryBJ: "229",
	CountryBL: "590",
	CountryBM: "1",
	CountryBN: "673",
	CountryBO: "591",
	CountryBQ: "599",
	CountryBR: "55",
	CountryBS: "1",
	CountryBT: "975",
	CountryBV: "47",
	CountryBW: "267",
	CountryBY: "375",
	CountryBZ: "501",
	CountryCA: "1",
	CountryCC: "61",
	CountryCD: "243",
	CountryCF: "236",
	CountryCG: "242",
	CountryCH: "41",
	CountryCI: "225",
	CountryCK: "682",
	CountryCL: "56",
	CountryCM: "237",
	CountryCN: "86",
	CountryCO: "57",
	CountryCR: "506",
	CountryCU: "53",
	CountryCV: "238",
	CountryCW: "599",
	CountryCX: "61",
	CountryCY: "357",
	CountryCZ: "420",
	CountryDE: "49",
	CountryDJ: "253",
	CountryDK: "45",
	CountryDM: "1",
	CountryDO: "1",
	CountryDZ: "213",
	CountryEC: "593",
	CountryEE: "372",
	CountryEG: "20",
	CountryEH: "212",
	CountryER: "291",
	CountryES: "34",
	CountryET: "251",
	CountryFI: "358",
	CountryFJ: "679",
	CountryFK: "500",
	CountryFM: "691",
	CountryFO: "298",
	CountryFR: "33",
	CountryGA: "241",
	CountryGB: "44",
	CountryGD: "1",
	CountryGE: "995",
	CountryGF: "594",
	CountryGG: "44",
	CountryGH: "233",
	CountryGI: "350",
	CountryGL: "299",
	CountryGM: "220",
	CountryGN: "224",
	CountryGP: "590",
	CountryGQ: "240",
	CountryGR: "30",
	CountryGS: "500",
	CountryGT: "502",
	CountryGU: "1",
	CountryGW: "245",
	CountryGY: "592",
	CountryHK: "852",
	CountryHM: "672",
	CountryHN: "504",
	CountryHR: "385",
	CountryHT: "509",
	CountryHU: "36",
	CountryID: "62",
	CountryIE: "353",
	CountryIL: "972",
	CountryIM: "44",
	CountryIN: "91",
	CountryIO: "246",
	CountryIQ: "964",
	CountryIR: "98",
	CountryIS: "354",
	CountryIT: "39",
	CountryJE: "44",
	CountryJM: "1",
	CountryJO: "962",
	CountryJP: "81",
	CountryKE: "254",
	CountryKG: "996",
	CountryKH: "855",
	CountryKI: "686",
	CountryKM: "269",
	CountryKN: "1",
	CountryKP: "850",
	CountryKR: "82",
	CountryKW: "965",
	CountryKY: "1",
	CountryKZ: "7",
	CountryLA: "856",
	CountryLB: "961",
	CountryLC: "1",
	CountryLI: "423",
	CountryLK: "94",
	CountryLR: "231",
	CountryLS: "266",
	CountryLT: "370",
	CountryLU: "352",
	CountryLV: "371",
	CountryLY: "218",
	CountryMA: "212",
	CountryMC: "377",
	CountryMD: "373",
	CountryME: "382",
	CountryMF: "590",
	CountryMG: "261",
	CountryMH: "692",
	CountryMK: "389",
	CountryML: "223",
	CountryMM: "95",
	CountryMN: "976",
	CountryMO: "853",
	CountryMP: "1",
	CountryMQ: "596",
	CountryMR: "222",
	CountryMS: "1",
	CountryMT: "356",
	CountryMU: "230",
	CountryMV: "960",
	CountryMW: "265",
	CountryMX: "52",
	CountryMY: "60",
	CountryMZ: "258",
	CountryNA: "264",
	CountryNC: "687",
	CountryNE: "227",
	CountryNF: "672",
	CountryNG: "234",
	CountryNI: "505",
	CountryNL: "31",
	CountryNO: "47",
	CountryNP: "977",
	CountryNR: "674",
	CountryNU: "683",
	CountryNZ: "64",
	CountryOM: "968",
	CountryPA: "507",
	CountryPE: "51",
	CountryPF: "689",
	CountryPG: "675",
	CountryPH: "63",
	CountryPK: "92",
	CountryPL: "48",
	CountryPM: "508",
	CountryPN: "64",
	CountryPR: "1",
	CountryPS: "970",
	CountryPT: "351",
	CountryPW: "680",
	CountryPY: "595",
	CountryQA: "974",
	CountryRE: "262",
	CountryRO: "40",
	CountryRS: "381",
	CountryRU: "7",
	CountryRW: "250",
	CountrySA: "966",
	CountrySB: "677",
	CountrySC: "248",
	CountrySD: "249",
	CountrySE: "46",
	CountrySG: "65",
	CountrySH: "290",
	CountrySI: "386",
	CountrySJ: "47",
	CountrySK: "421",
	CountrySL: "232",
	CountrySM: "378",
	CountrySN: "221",
	CountrySO: "252",
	CountrySR: "597",
	CountrySS: "211",
	CountryST: "239",
	CountrySV: "503",
	CountrySX: "1",
	CountrySY: "963",
	CountrySZ: "268",
	CountryTC: "1",
	CountryTD: "235",
	CountryTF: "262",
	CountryTG: "228",
	CountryTH: "66",
	CountryTJ: "992",
	CountryTK: "690",
	CountryTL: "670",
	CountryTM: "993",
	CountryTN: "216",
	CountryTO: "676",
	CountryTR: "90",
	CountryTT: "1",
	CountryTV: "688",
	CountryTW: "886",
	CountryTZ: "255",
	CountryUA: "380",
	CountryUG: "256",
	CountryUM: "1",
	CountryUS: "1",
	CountryUY: "598",
	CountryUZ: "998",
	CountryVA: "39",
	CountryVC: "1",
	CountryVE: "58",
	CountryVG: "1",
	CountryVI: "1",
	CountryVN: "84",
	CountryVU: "678",
	CountryWF: "681",
	CountryWS: "685",
	CountryYE: "967",
	CountryYT: "262",
	CountryZA: "27",
	CountryZM: "260",
	CountryZW: "263",
}

var countryToPostalCodeRe map[Country]*regexp.Regexp

func init() {
	countryToPostalCodeRe = make(map[Country]*regexp.Regexp, len(alpha2ToCountry))

	{
		re := regexp.MustCompile(`AD[1-7]0\d`)
		countryToPostalCodeRe[CountryAD] = re
	}

	{
		re := regexp.MustCompile(`\d{4}`)
		countryToPostalCodeRe[CountryAF] = re
	}

	{
		re := regexp.MustCompile(`(?:AI-)?2640`)
		countryToPostalCodeRe[CountryAI] = re
	}

	{
		re := regexp.MustCompile(`\d{4}`)
		countryToPostalCodeRe[CountryAL] = re
	}

	{
		re := regexp.MustCompile(`(?:37)?\d{4}`)
		countryToPostalCodeRe[CountryAM] = re
	}

	{
		re := regexp.MustCompile(`((?:[A-HJ-NP-Z])?\d{4})([A-Z]{3})?`)
		countryToPostalCodeRe[CountryAR] = re
	}

	{
		re := regexp.MustCompile(`(96799)(?:[ \-](\d{4}))?`)
		countryToPostalCodeRe[CountryAS] = re
	}

	{
		re := regexp.MustCompile(`\d{4}`)
		countryToPostalCodeRe[CountryAT] = re
	}

	{
		re := regexp.MustCompile(`\d{4}`)
		countryToPostalCodeRe[CountryAU] = re
	}

	{
		re := regexp.MustCompile(`22\d{3}`)
		countryToPostalCodeRe[CountryAX] = re
	}

	{
		re := regexp.MustCompile(`\d{4}`)
		countryToPostalCodeRe[CountryAZ] = re
	}

	{
		re := regexp.MustCompile(`\d{5}`)
		countryToPostalCodeRe[CountryBA] = re
	}

	{
		re := regexp.MustCompile(`BB\d{5}`)
		countryToPostalCodeRe[CountryBB] = re
	}

	{
		re := regexp.MustCompile(`\d{4}`)
		countryToPostalCodeRe[CountryBD] = re
	}

	{
		re := regexp.MustCompile(`\d{4}`)
		countryToPostalCodeRe[CountryBE] = re
	}

	{
		re := regexp.MustCompile(`\d{4}`)
		countryToPostalCodeRe[CountryBG] = re
	}

	{
		re := regexp.MustCompile(`(?:\d|1[0-2])\d{2}`)
		countryToPostalCodeRe[CountryBH] = re
	}

	{
		re := regexp.MustCompile(`9[78][01]\d{2}`)
		countryToPostalCodeRe[CountryBL] = re
	}

	{
		re := regexp.MustCompile(`[A-Z]{2} ?[A-Z0-9]{2}`)
		countryToPostalCodeRe[CountryBM] = re
	}

	{
		re := regexp.MustCompile(`[A-Z]{2} ?\d{4}`)
		countryToPostalCodeRe[CountryBN] = re
	}

	{
		re := regexp.MustCompile(`\d{5}-?\d{3}`)
		countryToPostalCodeRe[CountryBR] = re
	}

	{
		re := regexp.MustCompile(`\d{5}`)
		countryToPostalCodeRe[CountryBT] = re
	}

	{
		re := regexp.MustCompile(`\d{6}`)
		countryToPostalCodeRe[CountryBY] = re
	}

	{
		re := regexp.MustCompile(`[ABCEGHJKLMNPRSTVXY]\d[ABCEGHJ-NPRSTV-Z] ?\d[ABCEGHJ-NPRSTV-Z]\d`)
		countryToPostalCodeRe[CountryCA] = re
	}

	{
		re := regexp.MustCompile(`6799`)
		countryToPostalCodeRe[CountryCC] = re
	}

	{
		re := regexp.MustCompile(`\d{4}`)
		countryToPostalCodeRe[CountryCH] = re
	}

	{
		re := regexp.MustCompile(`\d{7}`)
		countryToPostalCodeRe[CountryCL] = re
	}

	{
		re := regexp.MustCompile(`\d{6}`)
		countryToPostalCodeRe[CountryCN] = re
	}

	{
		re := regexp.MustCompile(`\d{6}`)
		countryToPostalCodeRe[CountryCO] = re
	}

	{
		re := regexp.MustCompile(`\d{4,5}|\d{3}-\d{4}`)
		countryToPostalCodeRe[CountryCR] = re
	}

	{
		re := regexp.MustCompile(`\d{5}`)
		countryToPostalCodeRe[CountryCU] = re
	}

	{
		re := regexp.MustCompile(`\d{4}`)
		countryToPostalCodeRe[CountryCV] = re
	}

	{
		re := regexp.MustCompile(`6798`)
		countryToPostalCodeRe[CountryCX] = re
	}

	{
		re := regexp.MustCompile(`\d{4}`)
		countryToPostalCodeRe[CountryCY] = re
	}

	{
		re := regexp.MustCompile(`\d{3} ?\d{2}`)
		countryToPostalCodeRe[CountryCZ] = re
	}

	{
		re := regexp.MustCompile(`\d{5}`)
		countryToPostalCodeRe[CountryDE] = re
	}

	{
		re := regexp.MustCompile(`\d{4}`)
		countryToPostalCodeRe[CountryDK] = re
	}

	{
		re := regexp.MustCompile(`\d{5}`)
		countryToPostalCodeRe[CountryDO] = re
	}

	{
		re := regexp.MustCompile(`\d{5}`)
		countryToPostalCodeRe[CountryDZ] = re
	}

	{
		re := regexp.MustCompile(`\d{6}`)
		countryToPostalCodeRe[CountryEC] = re
	}

	{
		re := regexp.MustCompile(`\d{5}`)
		countryToPostalCodeRe[CountryEE] = re
	}

	{
		re := regexp.MustCompile(`\d{5}`)
		countryToPostalCodeRe[CountryEG] = re
	}

	{
		re := regexp.MustCompile(`\d{5}`)
		countryToPostalCodeRe[CountryEH] = re
	}

	{
		re := regexp.MustCompile(`\d{5}`)
		countryToPostalCodeRe[CountryES] = re
	}

	{
		re := regexp.MustCompile(`\d{4}`)
		countryToPostalCodeRe[CountryET] = re
	}

	{
		re := regexp.MustCompile(`\d{5}`)
		countryToPostalCodeRe[CountryFI] = re
	}

	{
		re := regexp.MustCompile(`FIQQ 1ZZ`)
		countryToPostalCodeRe[CountryFK] = re
	}

	{
		re := regexp.MustCompile(`(9694[1-4])(?:[ \-](\d{4}))?`)
		countryToPostalCodeRe[CountryFM] = re
	}

	{
		re := regexp.MustCompile(`\d{3}`)
		countryToPostalCodeRe[CountryFO] = re
	}

	{
		re := regexp.MustCompile(`\d{2} ?\d{3}`)
		countryToPostalCodeRe[CountryFR] = re
	}

	{
		re := regexp.MustCompile(`GIR ?0AA|(?:(?:AB|AL|B|BA|BB|BD|BF|BH|BL|BN|BR|BS|BT|BX|CA|CB|CF|CH|CM|CO|CR|CT|CV|CW|DA|DD|DE|DG|DH|DL|DN|DT|DY|E|EC|EH|EN|EX|FK|FY|G|GL|GY|GU|HA|HD|HG|HP|HR|HS|HU|HX|IG|IM|IP|IV|JE|KA|KT|KW|KY|L|LA|LD|LE|LL|LN|LS|LU|M|ME|MK|ML|N|NE|NG|NN|NP|NR|NW|OL|OX|PA|PE|PH|PL|PO|PR|RG|RH|RM|S|SA|SE|SG|SK|SL|SM|SN|SO|SP|SR|SS|ST|SW|SY|TA|TD|TF|TN|TQ|TR|TS|TW|UB|W|WA|WC|WD|WF|WN|WR|WS|WV|YO|ZE)(?:\d[\dA-Z]? ?\d[ABD-HJLN-UW-Z]{2}))|BFPO ?\d{1,4}`)
		countryToPostalCodeRe[CountryGB] = re
	}

	{
		re := regexp.MustCompile(`\d{4}`)
		countryToPostalCodeRe[CountryGE] = re
	}

	{
		re := regexp.MustCompile(`9[78]3\d{2}`)
		countryToPostalCodeRe[CountryGF] = re
	}

	{
		re := regexp.MustCompile(`GY\d[\dA-Z]? ?\d[ABD-HJLN-UW-Z]{2}`)
		countryToPostalCodeRe[CountryGG] = re
	}

	{
		re := regexp.MustCompile(`GX11 1AA`)
		countryToPostalCodeRe[CountryGI] = re
	}

	{
		re := regexp.MustCompile(`39\d{2}`)
		countryToPostalCodeRe[CountryGL] = re
	}

	{
		re := regexp.MustCompile(`\d{3}`)
		countryToPostalCodeRe[CountryGN] = re
	}

	{
		re := regexp.MustCompile(`9[78][01]\d{2}`)
		countryToPostalCodeRe[CountryGP] = re
	}

	{
		re := regexp.MustCompile(`\d{3} ?\d{2}`)
		countryToPostalCodeRe[CountryGR] = re
	}

	{
		re := regexp.MustCompile(`SIQQ 1ZZ`)
		countryToPostalCodeRe[CountryGS] = re
	}

	{
		re := regexp.MustCompile(`\d{5}`)
		countryToPostalCodeRe[CountryGT] = re
	}

	{
		re := regexp.MustCompile(`(969(?:[12]\d|3[12]))(?:[ \-](\d{4}))?`)
		countryToPostalCodeRe[CountryGU] = re
	}

	{
		re := regexp.MustCompile(`\d{4}`)
		countryToPostalCodeRe[CountryGW] = re
	}

	{
		re := regexp.MustCompile(`\d{4}`)
		countryToPostalCodeRe[CountryHM] = re
	}

	{
		re := regexp.MustCompile(`\d{5}`)
		countryToPostalCodeRe[CountryHN] = re
	}

	{
		re := regexp.MustCompile(`\d{5}`)
		countryToPostalCodeRe[CountryHR] = re
	}

	{
		re := regexp.MustCompile(`\d{4}`)
		countryToPostalCodeRe[CountryHT] = re
	}

	{
		re := regexp.MustCompile(`\d{4}`)
		countryToPostalCodeRe[CountryHU] = re
	}

	{
		re := regexp.MustCompile(`\d{5}`)
		countryToPostalCodeRe[CountryID] = re
	}

	{
		re := regexp.MustCompile(`[\dA-Z]{3} ?[\dA-Z]{4}`)
		countryToPostalCodeRe[CountryIE] = re
	}

	{
		re := regexp.MustCompile(`\d{5}(?:\d{2})?`)
		countryToPostalCodeRe[CountryIL] = re
	}

	{
		re := regexp.MustCompile(`IM\d[\dA-Z]? ?\d[ABD-HJLN-UW-Z]{2}`)
		countryToPostalCodeRe[CountryIM] = re
	}

	{
		re := regexp.MustCompile(`\d{6}`)
		countryToPostalCodeRe[CountryIN] = re
	}

	{
		re := regexp.MustCompile(`BBND 1ZZ`)
		countryToPostalCodeRe[CountryIO] = re
	}

	{
		re := regexp.MustCompile(`\d{5}`)
		countryToPostalCodeRe[CountryIQ] = re
	}

	{
		re := regexp.MustCompile(`\d{5}-?\d{5}`)
		countryToPostalCodeRe[CountryIR] = re
	}

	{
		re := regexp.MustCompile(`\d{3}`)
		countryToPostalCodeRe[CountryIS] = re
	}

	{
		re := regexp.MustCompile(`\d{5}`)
		countryToPostalCodeRe[CountryIT] = re
	}

	{
		re := regexp.MustCompile(`JE\d[\dA-Z]? ?\d[ABD-HJLN-UW-Z]{2}`)
		countryToPostalCodeRe[CountryJE] = re
	}

	{
		re := regexp.MustCompile(`\d{5}`)
		countryToPostalCodeRe[CountryJO] = re
	}

	{
		re := regexp.MustCompile(`\d{3}-?\d{4}`)
		countryToPostalCodeRe[CountryJP] = re
	}

	{
		re := regexp.MustCompile(`\d{5}`)
		countryToPostalCodeRe[CountryKE] = re
	}

	{
		re := regexp.MustCompile(`\d{6}`)
		countryToPostalCodeRe[CountryKG] = re
	}

	{
		re := regexp.MustCompile(`\d{5,6}`)
		countryToPostalCodeRe[CountryKH] = re
	}

	{
		re := regexp.MustCompile(`\d{5}`)
		countryToPostalCodeRe[CountryKR] = re
	}

	{
		re := regexp.MustCompile(`\d{5}`)
		countryToPostalCodeRe[CountryKW] = re
	}

	{
		re := regexp.MustCompile(`KY\d-\d{4}`)
		countryToPostalCodeRe[CountryKY] = re
	}

	{
		re := regexp.MustCompile(`\d{6}`)
		countryToPostalCodeRe[CountryKZ] = re
	}

	{
		re := regexp.MustCompile(`\d{5}`)
		countryToPostalCodeRe[CountryLA] = re
	}

	{
		re := regexp.MustCompile(`(?:\d{4})(?: ?(?:\d{4}))?`)
		countryToPostalCodeRe[CountryLB] = re
	}

	{
		re := regexp.MustCompile(`948[5-9]|949[0-8]`)
		countryToPostalCodeRe[CountryLI] = re
	}

	{
		re := regexp.MustCompile(`\d{5}`)
		countryToPostalCodeRe[CountryLK] = re
	}

	{
		re := regexp.MustCompile(`\d{4}`)
		countryToPostalCodeRe[CountryLR] = re
	}

	{
		re := regexp.MustCompile(`\d{3}`)
		countryToPostalCodeRe[CountryLS] = re
	}

	{
		re := regexp.MustCompile(`\d{5}`)
		countryToPostalCodeRe[CountryLT] = re
	}

	{
		re := regexp.MustCompile(`\d{4}`)
		countryToPostalCodeRe[CountryLU] = re
	}

	{
		re := regexp.MustCompile(`LV-\d{4}`)
		countryToPostalCodeRe[CountryLV] = re
	}

	{
		re := regexp.MustCompile(`\d{5}`)
		countryToPostalCodeRe[CountryMA] = re
	}

	{
		re := regexp.MustCompile(`980\d{2}`)
		countryToPostalCodeRe[CountryMC] = re
	}

	{
		re := regexp.MustCompile(`\d{4}`)
		countryToPostalCodeRe[CountryMD] = re
	}

	{
		re := regexp.MustCompile(`8\d{4}`)
		countryToPostalCodeRe[CountryME] = re
	}

	{
		re := regexp.MustCompile(`9[78][01]\d{2}`)
		countryToPostalCodeRe[CountryMF] = re
	}

	{
		re := regexp.MustCompile(`\d{3}`)
		countryToPostalCodeRe[CountryMG] = re
	}

	{
		re := regexp.MustCompile(`(969[67]\d)(?:[ \-](\d{4}))?`)
		countryToPostalCodeRe[CountryMH] = re
	}

	{
		re := regexp.MustCompile(`\d{4}`)
		countryToPostalCodeRe[CountryMK] = re
	}

	{
		re := regexp.MustCompile(`\d{5}`)
		countryToPostalCodeRe[CountryMM] = re
	}

	{
		re := regexp.MustCompile(`\d{5}`)
		countryToPostalCodeRe[CountryMN] = re
	}

	{
		re := regexp.MustCompile(`(9695[012])(?:[ \-](\d{4}))?`)
		countryToPostalCodeRe[CountryMP] = re
	}

	{
		re := regexp.MustCompile(`9[78]2\d{2}`)
		countryToPostalCodeRe[CountryMQ] = re
	}

	{
		re := regexp.MustCompile(`[A-Z]{3} ?\d{2,4}`)
		countryToPostalCodeRe[CountryMT] = re
	}

	{
		re := regexp.MustCompile(`\d{3}(?:\d{2}|[A-Z]{2}\d{3})`)
		countryToPostalCodeRe[CountryMU] = re
	}

	{
		re := regexp.MustCompile(`\d{5}`)
		countryToPostalCodeRe[CountryMV] = re
	}

	{
		re := regexp.MustCompile(`\d{5}`)
		countryToPostalCodeRe[CountryMX] = re
	}

	{
		re := regexp.MustCompile(`\d{5}`)
		countryToPostalCodeRe[CountryMY] = re
	}

	{
		re := regexp.MustCompile(`\d{4}`)
		countryToPostalCodeRe[CountryMZ] = re
	}

	{
		re := regexp.MustCompile(`\d{5}`)
		countryToPostalCodeRe[CountryNA] = re
	}

	{
		re := regexp.MustCompile(`988\d{2}`)
		countryToPostalCodeRe[CountryNC] = re
	}

	{
		re := regexp.MustCompile(`\d{4}`)
		countryToPostalCodeRe[CountryNE] = re
	}

	{
		re := regexp.MustCompile(`2899`)
		countryToPostalCodeRe[CountryNF] = re
	}

	{
		re := regexp.MustCompile(`\d{6}`)
		countryToPostalCodeRe[CountryNG] = re
	}

	{
		re := regexp.MustCompile(`\d{5}`)
		countryToPostalCodeRe[CountryNI] = re
	}

	{
		re := regexp.MustCompile(`\d{4} ?[A-Z]{2}`)
		countryToPostalCodeRe[CountryNL] = re
	}

	{
		re := regexp.MustCompile(`\d{4}`)
		countryToPostalCodeRe[CountryNO] = re
	}

	{
		re := regexp.MustCompile(`\d{5}`)
		countryToPostalCodeRe[CountryNP] = re
	}

	{
		re := regexp.MustCompile(`\d{4}`)
		countryToPostalCodeRe[CountryNZ] = re
	}

	{
		re := regexp.MustCompile(`(?:PC )?\d{3}`)
		countryToPostalCodeRe[CountryOM] = re
	}

	{
		re := regexp.MustCompile(`(?:LIMA \d{1,2}|CALLAO 0?\d)|[0-2]\d{4}`)
		countryToPostalCodeRe[CountryPE] = re
	}

	{
		re := regexp.MustCompile(`987\d{2}`)
		countryToPostalCodeRe[CountryPF] = re
	}

	{
		re := regexp.MustCompile(`\d{3}`)
		countryToPostalCodeRe[CountryPG] = re
	}

	{
		re := regexp.MustCompile(`\d{4}`)
		countryToPostalCodeRe[CountryPH] = re
	}

	{
		re := regexp.MustCompile(`\d{5}`)
		countryToPostalCodeRe[CountryPK] = re
	}

	{
		re := regexp.MustCompile(`\d{2}-\d{3}`)
		countryToPostalCodeRe[CountryPL] = re
	}

	{
		re := regexp.MustCompile(`9[78]5\d{2}`)
		countryToPostalCodeRe[CountryPM] = re
	}

	{
		re := regexp.MustCompile(`PCRN 1ZZ`)
		countryToPostalCodeRe[CountryPN] = re
	}

	{
		re := regexp.MustCompile(`(00[679]\d{2})(?:[ \-](\d{4}))?`)
		countryToPostalCodeRe[CountryPR] = re
	}

	{
		re := regexp.MustCompile(`\d{4}-\d{3}`)
		countryToPostalCodeRe[CountryPT] = re
	}

	{
		re := regexp.MustCompile(`(969(?:39|40))(?:[ \-](\d{4}))?`)
		countryToPostalCodeRe[CountryPW] = re
	}

	{
		re := regexp.MustCompile(`\d{4}`)
		countryToPostalCodeRe[CountryPY] = re
	}

	{
		re := regexp.MustCompile(`9[78]4\d{2}`)
		countryToPostalCodeRe[CountryRE] = re
	}

	{
		re := regexp.MustCompile(`\d{6}`)
		countryToPostalCodeRe[CountryRO] = re
	}

	{
		re := regexp.MustCompile(`\d{5,6}`)
		countryToPostalCodeRe[CountryRS] = re
	}

	{
		re := regexp.MustCompile(`\d{6}`)
		countryToPostalCodeRe[CountryRU] = re
	}

	{
		re := regexp.MustCompile(`\d{5}`)
		countryToPostalCodeRe[CountrySA] = re
	}

	{
		re := regexp.MustCompile(`\d{5}`)
		countryToPostalCodeRe[CountrySD] = re
	}

	{
		re := regexp.MustCompile(`\d{3} ?\d{2}`)
		countryToPostalCodeRe[CountrySE] = re
	}

	{
		re := regexp.MustCompile(`\d{6}`)
		countryToPostalCodeRe[CountrySG] = re
	}

	{
		re := regexp.MustCompile(`(?:ASCN|STHL) 1ZZ`)
		countryToPostalCodeRe[CountrySH] = re
	}

	{
		re := regexp.MustCompile(`\d{4}`)
		countryToPostalCodeRe[CountrySI] = re
	}

	{
		re := regexp.MustCompile(`\d{4}`)
		countryToPostalCodeRe[CountrySJ] = re
	}

	{
		re := regexp.MustCompile(`\d{3} ?\d{2}`)
		countryToPostalCodeRe[CountrySK] = re
	}

	{
		re := regexp.MustCompile(`4789\d`)
		countryToPostalCodeRe[CountrySM] = re
	}

	{
		re := regexp.MustCompile(`\d{5}`)
		countryToPostalCodeRe[CountrySN] = re
	}

	{
		re := regexp.MustCompile(`[A-Z]{2} ?\d{5}`)
		countryToPostalCodeRe[CountrySO] = re
	}

	{
		re := regexp.MustCompile(`CP [1-3][1-7][0-2]\d`)
		countryToPostalCodeRe[CountrySV] = re
	}

	{
		re := regexp.MustCompile(`[HLMS]\d{3}`)
		countryToPostalCodeRe[CountrySZ] = re
	}

	{
		re := regexp.MustCompile(`TKCA 1ZZ`)
		countryToPostalCodeRe[CountryTC] = re
	}

	{
		re := regexp.MustCompile(`\d{5}`)
		countryToPostalCodeRe[CountryTH] = re
	}

	{
		re := regexp.MustCompile(`\d{6}`)
		countryToPostalCodeRe[CountryTJ] = re
	}

	{
		re := regexp.MustCompile(`\d{6}`)
		countryToPostalCodeRe[CountryTM] = re
	}

	{
		re := regexp.MustCompile(`\d{4}`)
		countryToPostalCodeRe[CountryTN] = re
	}

	{
		re := regexp.MustCompile(`\d{5}`)
		countryToPostalCodeRe[CountryTR] = re
	}

	{
		re := regexp.MustCompile(`\d{3}(?:\d{2,3})?`)
		countryToPostalCodeRe[CountryTW] = re
	}

	{
		re := regexp.MustCompile(`\d{4,5}`)
		countryToPostalCodeRe[CountryTZ] = re
	}

	{
		re := regexp.MustCompile(`\d{5}`)
		countryToPostalCodeRe[CountryUA] = re
	}

	{
		re := regexp.MustCompile(`96898`)
		countryToPostalCodeRe[CountryUM] = re
	}

	{
		re := regexp.MustCompile(`(\d{5})(?:[ \-](\d{4}))?`)
		countryToPostalCodeRe[CountryUS] = re
	}

	{
		re := regexp.MustCompile(`\d{5}`)
		countryToPostalCodeRe[CountryUY] = re
	}

	{
		re := regexp.MustCompile(`\d{6}`)
		countryToPostalCodeRe[CountryUZ] = re
	}

	{
		re := regexp.MustCompile(`00120`)
		countryToPostalCodeRe[CountryVA] = re
	}

	{
		re := regexp.MustCompile(`VC\d{4}`)
		countryToPostalCodeRe[CountryVC] = re
	}

	{
		re := regexp.MustCompile(`\d{4}`)
		countryToPostalCodeRe[CountryVE] = re
	}

	{
		re := regexp.MustCompile(`VG\d{4}`)
		countryToPostalCodeRe[CountryVG] = re
	}

	{
		re := regexp.MustCompile(`(008(?:(?:[0-4]\d)|(?:5[01])))(?:[ \-](\d{4}))?`)
		countryToPostalCodeRe[CountryVI] = re
	}

	{
		re := regexp.MustCompile(`\d{5}\d?`)
		countryToPostalCodeRe[CountryVN] = re
	}

	{
		re := regexp.MustCompile(`986\d{2}`)
		countryToPostalCodeRe[CountryWF] = re
	}

	{
		re := regexp.MustCompile(`976\d{2}`)
		countryToPostalCodeRe[CountryYT] = re
	}

	{
		re := regexp.MustCompile(`\d{4}`)
		countryToPostalCodeRe[CountryZA] = re
	}

	{
		re := regexp.MustCompile(`\d{5}`)
		countryToPostalCodeRe[CountryZM] = re
	}

}

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
	dst = append(dst, CountryAD)
	dst = append(dst, CountryAE)
	dst = append(dst, CountryAF)
	dst = append(dst, CountryAG)
	dst = append(dst, CountryAI)
	dst = append(dst, CountryAL)
	dst = append(dst, CountryAM)
	dst = append(dst, CountryAO)
	dst = append(dst, CountryAQ)
	dst = append(dst, CountryAR)
	dst = append(dst, CountryAS)
	dst = append(dst, CountryAT)
	dst = append(dst, CountryAU)
	dst = append(dst, CountryAW)
	dst = append(dst, CountryAX)
	dst = append(dst, CountryAZ)
	dst = append(dst, CountryBA)
	dst = append(dst, CountryBB)
	dst = append(dst, CountryBD)
	dst = append(dst, CountryBE)
	dst = append(dst, CountryBF)
	dst = append(dst, CountryBG)
	dst = append(dst, CountryBH)
	dst = append(dst, CountryBI)
	dst = append(dst, CountryBJ)
	dst = append(dst, CountryBL)
	dst = append(dst, CountryBM)
	dst = append(dst, CountryBN)
	dst = append(dst, CountryBO)
	dst = append(dst, CountryBQ)
	dst = append(dst, CountryBR)
	dst = append(dst, CountryBS)
	dst = append(dst, CountryBT)
	dst = append(dst, CountryBV)
	dst = append(dst, CountryBW)
	dst = append(dst, CountryBY)
	dst = append(dst, CountryBZ)
	dst = append(dst, CountryCA)
	dst = append(dst, CountryCC)
	dst = append(dst, CountryCD)
	dst = append(dst, CountryCF)
	dst = append(dst, CountryCG)
	dst = append(dst, CountryCH)
	dst = append(dst, CountryCI)
	dst = append(dst, CountryCK)
	dst = append(dst, CountryCL)
	dst = append(dst, CountryCM)
	dst = append(dst, CountryCN)
	dst = append(dst, CountryCO)
	dst = append(dst, CountryCR)
	dst = append(dst, CountryCU)
	dst = append(dst, CountryCV)
	dst = append(dst, CountryCW)
	dst = append(dst, CountryCX)
	dst = append(dst, CountryCY)
	dst = append(dst, CountryCZ)
	dst = append(dst, CountryDE)
	dst = append(dst, CountryDJ)
	dst = append(dst, CountryDK)
	dst = append(dst, CountryDM)
	dst = append(dst, CountryDO)
	dst = append(dst, CountryDZ)
	dst = append(dst, CountryEC)
	dst = append(dst, CountryEE)
	dst = append(dst, CountryEG)
	dst = append(dst, CountryEH)
	dst = append(dst, CountryER)
	dst = append(dst, CountryES)
	dst = append(dst, CountryET)
	dst = append(dst, CountryFI)
	dst = append(dst, CountryFJ)
	dst = append(dst, CountryFK)
	dst = append(dst, CountryFM)
	dst = append(dst, CountryFO)
	dst = append(dst, CountryFR)
	dst = append(dst, CountryGA)
	dst = append(dst, CountryGB)
	dst = append(dst, CountryGD)
	dst = append(dst, CountryGE)
	dst = append(dst, CountryGF)
	dst = append(dst, CountryGG)
	dst = append(dst, CountryGH)
	dst = append(dst, CountryGI)
	dst = append(dst, CountryGL)
	dst = append(dst, CountryGM)
	dst = append(dst, CountryGN)
	dst = append(dst, CountryGP)
	dst = append(dst, CountryGQ)
	dst = append(dst, CountryGR)
	dst = append(dst, CountryGS)
	dst = append(dst, CountryGT)
	dst = append(dst, CountryGU)
	dst = append(dst, CountryGW)
	dst = append(dst, CountryGY)
	dst = append(dst, CountryHK)
	dst = append(dst, CountryHM)
	dst = append(dst, CountryHN)
	dst = append(dst, CountryHR)
	dst = append(dst, CountryHT)
	dst = append(dst, CountryHU)
	dst = append(dst, CountryID)
	dst = append(dst, CountryIE)
	dst = append(dst, CountryIL)
	dst = append(dst, CountryIM)
	dst = append(dst, CountryIN)
	dst = append(dst, CountryIO)
	dst = append(dst, CountryIQ)
	dst = append(dst, CountryIR)
	dst = append(dst, CountryIS)
	dst = append(dst, CountryIT)
	dst = append(dst, CountryJE)
	dst = append(dst, CountryJM)
	dst = append(dst, CountryJO)
	dst = append(dst, CountryJP)
	dst = append(dst, CountryKE)
	dst = append(dst, CountryKG)
	dst = append(dst, CountryKH)
	dst = append(dst, CountryKI)
	dst = append(dst, CountryKM)
	dst = append(dst, CountryKN)
	dst = append(dst, CountryKP)
	dst = append(dst, CountryKR)
	dst = append(dst, CountryKW)
	dst = append(dst, CountryKY)
	dst = append(dst, CountryKZ)
	dst = append(dst, CountryLA)
	dst = append(dst, CountryLB)
	dst = append(dst, CountryLC)
	dst = append(dst, CountryLI)
	dst = append(dst, CountryLK)
	dst = append(dst, CountryLR)
	dst = append(dst, CountryLS)
	dst = append(dst, CountryLT)
	dst = append(dst, CountryLU)
	dst = append(dst, CountryLV)
	dst = append(dst, CountryLY)
	dst = append(dst, CountryMA)
	dst = append(dst, CountryMC)
	dst = append(dst, CountryMD)
	dst = append(dst, CountryME)
	dst = append(dst, CountryMF)
	dst = append(dst, CountryMG)
	dst = append(dst, CountryMH)
	dst = append(dst, CountryMK)
	dst = append(dst, CountryML)
	dst = append(dst, CountryMM)
	dst = append(dst, CountryMN)
	dst = append(dst, CountryMO)
	dst = append(dst, CountryMP)
	dst = append(dst, CountryMQ)
	dst = append(dst, CountryMR)
	dst = append(dst, CountryMS)
	dst = append(dst, CountryMT)
	dst = append(dst, CountryMU)
	dst = append(dst, CountryMV)
	dst = append(dst, CountryMW)
	dst = append(dst, CountryMX)
	dst = append(dst, CountryMY)
	dst = append(dst, CountryMZ)
	dst = append(dst, CountryNA)
	dst = append(dst, CountryNC)
	dst = append(dst, CountryNE)
	dst = append(dst, CountryNF)
	dst = append(dst, CountryNG)
	dst = append(dst, CountryNI)
	dst = append(dst, CountryNL)
	dst = append(dst, CountryNO)
	dst = append(dst, CountryNP)
	dst = append(dst, CountryNR)
	dst = append(dst, CountryNU)
	dst = append(dst, CountryNZ)
	dst = append(dst, CountryOM)
	dst = append(dst, CountryPA)
	dst = append(dst, CountryPE)
	dst = append(dst, CountryPF)
	dst = append(dst, CountryPG)
	dst = append(dst, CountryPH)
	dst = append(dst, CountryPK)
	dst = append(dst, CountryPL)
	dst = append(dst, CountryPM)
	dst = append(dst, CountryPN)
	dst = append(dst, CountryPR)
	dst = append(dst, CountryPS)
	dst = append(dst, CountryPT)
	dst = append(dst, CountryPW)
	dst = append(dst, CountryPY)
	dst = append(dst, CountryQA)
	dst = append(dst, CountryRE)
	dst = append(dst, CountryRO)
	dst = append(dst, CountryRS)
	dst = append(dst, CountryRU)
	dst = append(dst, CountryRW)
	dst = append(dst, CountrySA)
	dst = append(dst, CountrySB)
	dst = append(dst, CountrySC)
	dst = append(dst, CountrySD)
	dst = append(dst, CountrySE)
	dst = append(dst, CountrySG)
	dst = append(dst, CountrySH)
	dst = append(dst, CountrySI)
	dst = append(dst, CountrySJ)
	dst = append(dst, CountrySK)
	dst = append(dst, CountrySL)
	dst = append(dst, CountrySM)
	dst = append(dst, CountrySN)
	dst = append(dst, CountrySO)
	dst = append(dst, CountrySR)
	dst = append(dst, CountrySS)
	dst = append(dst, CountryST)
	dst = append(dst, CountrySV)
	dst = append(dst, CountrySX)
	dst = append(dst, CountrySY)
	dst = append(dst, CountrySZ)
	dst = append(dst, CountryTC)
	dst = append(dst, CountryTD)
	dst = append(dst, CountryTF)
	dst = append(dst, CountryTG)
	dst = append(dst, CountryTH)
	dst = append(dst, CountryTJ)
	dst = append(dst, CountryTK)
	dst = append(dst, CountryTL)
	dst = append(dst, CountryTM)
	dst = append(dst, CountryTN)
	dst = append(dst, CountryTO)
	dst = append(dst, CountryTR)
	dst = append(dst, CountryTT)
	dst = append(dst, CountryTV)
	dst = append(dst, CountryTW)
	dst = append(dst, CountryTZ)
	dst = append(dst, CountryUA)
	dst = append(dst, CountryUG)
	dst = append(dst, CountryUM)
	dst = append(dst, CountryUS)
	dst = append(dst, CountryUY)
	dst = append(dst, CountryUZ)
	dst = append(dst, CountryVA)
	dst = append(dst, CountryVC)
	dst = append(dst, CountryVE)
	dst = append(dst, CountryVG)
	dst = append(dst, CountryVI)
	dst = append(dst, CountryVN)
	dst = append(dst, CountryVU)
	dst = append(dst, CountryWF)
	dst = append(dst, CountryWS)
	dst = append(dst, CountryYE)
	dst = append(dst, CountryYT)
	dst = append(dst, CountryZA)
	dst = append(dst, CountryZM)
	dst = append(dst, CountryZW)
	return dst
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
