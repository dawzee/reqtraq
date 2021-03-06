// Daedalean-specific configuration file, defining a bunch of constants that are company-specific. Create your own and
// mark it with your build tag, then remove the !ddln tag below.
// +build ddln !ddln

package config

type RequirementLevel int

// Requirement levels according to DO-178C (do not change!)
const (
	SYSTEM RequirementLevel = iota
	HIGH
	LOW
)

// Document types:
// ORD - Overall (aka System) Requirement Document
// SRD - Software Requirements Data
// SDD - Software Design Description
// HRD - Hardware Requirements Data
// HDD - Hardware Design Description

// Requirement types:
// SYS - System/overall requirements (defined in ORD documents)
// SWH - Sofware high-level requirements (defined in SRD documents)
// SWL - Software low-level requirements (defined in SDD documents)
// HWH - Hardware high-level requirements (defined in HRD documents)
// HWL - Hardware low-level requirements (defined in HDD documents)

// Map from requirement type to requirement level.
var ReqTypeToReqLevel = map[string]RequirementLevel{
	"SYS": SYSTEM,
	"SWH": HIGH,
	"HWH": HIGH,
	"SWL": LOW,
	"HWL": LOW,
}

// Map from document type to requirement type.
var DocTypeToReqType = map[string]string{
	"ORD": "SYS",
	"SRD": "SWH",
	"HRD": "HWH",
	"SDD": "SWL",
	"HDD": "HWL",
}

// Map from requirement type to document type.
var ReqTypeToDocType = map[string]string{
	"SYS": "ORD",
	"SWH": "SRD",
	"SWL": "SDD",
	"HWH": "HRD",
	"HWL": "HDD",
}

// Map from document type to document ID.
// https://a.daedalean.ai/organisation-of-documentation
var DocTypeToDocId = map[string]string{
	"H":       "0",
	"DS":      "1",
	"CLSRS":   "5",
	"RS":      "6",
	"SDS":     "7",
	"CS":      "8",
	"AS":      "9",
	"HCS":     "10",
	"HDS":     "11",
	"HVVS":    "12",
	"HARS":    "13",
	"CLCMP":   "21",
	"CLPAP":   "23",
	"CLTQP":   "25",
	"CMP":     "26",
	"CLSCMP":  "27",
	"QAP":     "28",
	"CLSQAP":  "29",
	"SDP":     "30",
	"CLHPAP":  "33",
	"TPPR":    "50",
	"TPSQAR":  "51",
	"ORD":     "100",
	"ICD":     "101",
	"CP":      "102",
	"DP":      "103",
	"DD":      "104",
	"VAP":     "105",
	"VEP":     "106",
	"CI":      "107",
	"FHA":     "108",
	"SFHA":    "109",
	"PSSA":    "110",
	"SSA":     "111",
	"CCA":     "112",
	"SPP":     "113",
	"VAD":     "114",
	"VED":     "115",
	"ECM":     "116",
	"EPA":     "117",
	"CSCR":    "118",
	"PSAC":    "134",
	"TQP":     "135",
	"SVP":     "136",
	"SRD":     "137",
	"SDD":     "138",
	"SLECI":   "141",
	"SCI":     "142",
	"SAS":     "143",
	"STD":     "144",
	"SQARI":   "145",
	"SPVP":    "146",
	"SVVDVP":  "147",
	"STCP":    "148",
	"SPVR":    "149",
	"SVVDVR":  "150",
	"STR":     "151",
	"SRA":     "152",
	"PHAC":    "167",
	"HDP":     "168",
	"HVVP":    "169",
	"ECMP":    "170",
	"HECI":    "173",
	"HCI":     "174",
	"HRD":     "175",
	"HDD":     "176",
	"HTD":     "177",
	"HRAP":    "178",
	"HRAR":    "179",
	"HTP":     "180",
	"HTR":     "181",
	"HATC":    "182",
	"HAS":     "183",
	"FFPA":    "184",
	"TPORD":   "200",
	"TPICD":   "201",
	"TPCP":    "202",
	"TPDP":    "203",
	"TPDD":    "204",
	"TPVAP":   "205",
	"TPVEP":   "206",
	"TPCI":    "207",
	"TPFHA":   "208",
	"TPSFHA":  "209",
	"TPPSSA":  "210",
	"TPSSA":   "211",
	"TPCCA":   "212",
	"TPSPP":   "213",
	"TPVAD":   "214",
	"TPVED":   "215",
	"TPECM":   "216",
	"TPEPA":   "217",
	"TPCSCR":  "218",
	"TPPSAC":  "234",
	"TPSDP":   "235",
	"TPSVP":   "236",
	"TPSRD":   "237",
	"TPSDD":   "238",
	"TPSVCP":  "239",
	"TPSVR":   "240",
	"TPSLECI": "241",
	"TPSCI":   "242",
	"TPSAS":   "243",
	"TPSTD":   "244",
	"TPPHAC":  "267",
	"TPHDP":   "268",
	"TPHVEP":  "269",
	"TPECMP":  "270",
	"TPHVAP":  "271",
	"TPHCMP":  "272",
	"TPHECI":  "273",
	"TPHCI":   "274",
	"TPHRD":   "275",
	"TPHDD":   "276",
	"TPHTD":   "277",
	"TPHRAP":  "278",
	"TPHRAR":  "279",
	"TPHTP":   "280",
	"TPHTR":   "281",
	"TPHATC":  "282",
	"TPHACS":  "283",
	"TPFFPA":  "284",
	"CLORD":   "300",
	"CLICD":   "301",
	"CLCP":    "302",
	"CLDP":    "303",
	"CLDD":    "304",
	"CLVAP":   "305",
	"CLVEP":   "306",
	"CLCI":    "307",
	"CLFHA":   "308",
	"CLSFHA":  "309",
	"CLPSSA":  "310",
	"CLSSA":   "311",
	"CLCCA":   "312",
	"CLSPP":   "313",
	"CLVAD":   "314",
	"CLVED":   "315",
	"CLECM":   "316",
	"CLEPA":   "317",
	"CLCSCR":  "318",
	"CLPSAC":  "334",
	"CLSDP":   "335",
	"CLSVP":   "336",
	"CLSRD":   "337",
	"CLSDD":   "338",
	"CLSVCP":  "339",
	"CLSVR":   "340",
	"CLSLECI": "341",
	"CLSCI":   "342",
	"CLSAS":   "343",
	"CLSTD":   "344",
	"CLSOI1":  "345",
	"CLSOI2":  "346",
	"CLSOI3":  "347",
	"CLSOI4":  "348",
	"CLSPR":   "349",
	"CLRA":    "350",
	"CLSCR":   "351",
	"CLPHAC":  "367",
	"CLHDP":   "368",
	"CLHVEP":  "369",
	"CLECMP":  "370",
	"CLHVAP":  "371",
	"CLHCMP":  "372",
	"CLHECI":  "373",
	"CLHCI":   "374",
	"CLHRD":   "375",
	"CLHDD":   "376",
	"CLHTD":   "377",
	"CLHRAP":  "378",
	"CLHRAR":  "379",
	"CLHTP":   "380",
	"CLHTR":   "381",
	"CLHATC":  "382",
	"CLHACS":  "383",
	"CLFFPA":  "384",
}
