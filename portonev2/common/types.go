package common

import (
	"encoding/json"
)

// Currency 통화 단위
type Currency string

const (
	CurrencyKRW Currency = "KRW" // 대한민국 원화
	CurrencyUSD Currency = "USD" // 미국 달러
	CurrencyJPY Currency = "JPY" // 일본 엔화
	CurrencyAED Currency = "AED" // UAE Dirham
	CurrencyAFN Currency = "AFN" // Afghani
	CurrencyALL Currency = "ALL" // Lek
	CurrencyAMD Currency = "AMD" // Armenian Dram
	CurrencyANG Currency = "ANG" // Netherlands Antillean Guilder
	CurrencyAOA Currency = "AOA" // Kwanza
	CurrencyARS Currency = "ARS" // Argentine Peso
	CurrencyAUD Currency = "AUD" // Australian Dollar
	CurrencyAWG Currency = "AWG" // Aruban Florin
	CurrencyAZN Currency = "AZN" // Azerbaijan Manat
	CurrencyBAM Currency = "BAM" // Convertible Mark
	CurrencyBBD Currency = "BBD" // Barbados Dollar
	CurrencyBDT Currency = "BDT" // Taka
	CurrencyBGN Currency = "BGN" // Bulgarian Lev
	CurrencyBHD Currency = "BHD" // Bahraini Dinar
	CurrencyBIF Currency = "BIF" // Burundi Franc
	CurrencyBMD Currency = "BMD" // Bermudian Dollar
	CurrencyBND Currency = "BND" // Brunei Dollar
	CurrencyBOB Currency = "BOB" // Boliviano
	CurrencyBOV Currency = "BOV" // Mvdol
	CurrencyBRL Currency = "BRL" // Brazilian Real
	CurrencyBSD Currency = "BSD" // Bahamian Dollar
	CurrencyBTN Currency = "BTN" // Ngultrum
	CurrencyBWP Currency = "BWP" // Pula
	CurrencyBYN Currency = "BYN" // Belarusian Ruble
	CurrencyBZD Currency = "BZD" // Belize Dollar
	CurrencyCAD Currency = "CAD" // Canadian Dollar
	CurrencyCDF Currency = "CDF" // Congolese Franc
	CurrencyCHE Currency = "CHE" // WIR Euro
	CurrencyCHF Currency = "CHF" // Swiss Franc
	CurrencyCHW Currency = "CHW" // WIR Franc
	CurrencyCLF Currency = "CLF" // Unidad de Fomento
	CurrencyCLP Currency = "CLP" // Chilean Peso
	CurrencyCNY Currency = "CNY" // Yuan Renminbi
	CurrencyCOP Currency = "COP" // Colombian Peso
	CurrencyCOU Currency = "COU" // Unidad de Valor Real
	CurrencyCRC Currency = "CRC" // Costa Rican Colon
	CurrencyCUC Currency = "CUC" // Peso Convertible
	CurrencyCUP Currency = "CUP" // Cuban Peso
	CurrencyCVE Currency = "CVE" // Cabo Verde Escudo
	CurrencyCZK Currency = "CZK" // Czech Koruna
	CurrencyDJF Currency = "DJF" // Djibouti Franc
	CurrencyDKK Currency = "DKK" // Danish Krone
	CurrencyDOP Currency = "DOP" // Dominican Peso
	CurrencyDZD Currency = "DZD" // Algerian Dinar
	CurrencyEGP Currency = "EGP" // Egyptian Pound
	CurrencyERN Currency = "ERN" // Nakfa
	CurrencyETB Currency = "ETB" // Ethiopian Birr
	CurrencyEUR Currency = "EUR" // Euro
	CurrencyFJD Currency = "FJD" // Fiji Dollar
	CurrencyFKP Currency = "FKP" // Falkland Islands Pound
	CurrencyGBP Currency = "GBP" // Pound Sterling
	CurrencyGEL Currency = "GEL" // Lari
	CurrencyGHS Currency = "GHS" // Ghana Cedi
	CurrencyGIP Currency = "GIP" // Gibraltar Pound
	CurrencyGMD Currency = "GMD" // Dalasi
	CurrencyGNF Currency = "GNF" // Guinean Franc
	CurrencyGTQ Currency = "GTQ" // Quetzal
	CurrencyGYD Currency = "GYD" // Guyana Dollar
	CurrencyHKD Currency = "HKD" // Hong Kong Dollar
	CurrencyHNL Currency = "HNL" // Lempira
	CurrencyHRK Currency = "HRK" // Kuna (Replaced by EUR)
	CurrencyHTG Currency = "HTG" // Gourde
	CurrencyHUF Currency = "HUF" // Forint
	CurrencyIDR Currency = "IDR" // Rupiah
	CurrencyILS Currency = "ILS" // New Israeli Sheqel
	CurrencyINR Currency = "INR" // Indian Rupee
	CurrencyIQD Currency = "IQD" // Iraqi Dinar
	CurrencyIRR Currency = "IRR" // Iranian Rial
	CurrencyISK Currency = "ISK" // Iceland Krona
	CurrencyJMD Currency = "JMD" // Jamaican Dollar
	CurrencyJOD Currency = "JOD" // Jordanian Dinar
	CurrencyKES Currency = "KES" // Kenyan Shilling
	CurrencyKGS Currency = "KGS" // Som
	CurrencyKHR Currency = "KHR" // Riel
	CurrencyKMF Currency = "KMF" // Comorian Franc
	CurrencyKPW Currency = "KPW" // North Korean Won
	CurrencyKWD Currency = "KWD" // Kuwaiti Dinar
	CurrencyKYD Currency = "KYD" // Cayman Islands Dollar
	CurrencyKZT Currency = "KZT" // Tenge
	CurrencyLAK Currency = "LAK" // Lao Kip
	CurrencyLBP Currency = "LBP" // Lebanese Pound
	CurrencyLKR Currency = "LKR" // Sri Lanka Rupee
	CurrencyLRD Currency = "LRD" // Liberian Dollar
	CurrencyLSL Currency = "LSL" // Loti
	CurrencyLYD Currency = "LYD" // Libyan Dinar
	CurrencyMAD Currency = "MAD" // Moroccan Dirham
	CurrencyMDL Currency = "MDL" // Moldovan Leu
	CurrencyMGA Currency = "MGA" // Malagasy Ariary
	CurrencyMKD Currency = "MKD" // Denar
	CurrencyMMK Currency = "MMK" // Kyat
	CurrencyMNT Currency = "MNT" // Tugrik
	CurrencyMOP Currency = "MOP" // Pataca
	CurrencyMRU Currency = "MRU" // Ouguiya
	CurrencyMUR Currency = "MUR" // Mauritius Rupee
	CurrencyMVR Currency = "MVR" // Rufiyaa
	CurrencyMWK Currency = "MWK" // Malawi Kwacha
	CurrencyMXN Currency = "MXN" // Mexican Peso
	CurrencyMXV Currency = "MXV" // Mexican Unidad de Inversion (UDI)
	CurrencyMYR Currency = "MYR" // Malaysian Ringgit
	CurrencyMZN Currency = "MZN" // Mozambique Metical
	CurrencyNAD Currency = "NAD" // Namibia Dollar
	CurrencyNGN Currency = "NGN" // Naira
	CurrencyNIO Currency = "NIO" // Cordoba Oro
	CurrencyNOK Currency = "NOK" // Norwegian Krone
	CurrencyNPR Currency = "NPR" // Nepalese Rupee
	CurrencyNZD Currency = "NZD" // New Zealand Dollar
	CurrencyOMR Currency = "OMR" // Rial Omani
	CurrencyPAB Currency = "PAB" // Balboa
	CurrencyPEN Currency = "PEN" // Sol
	CurrencyPGK Currency = "PGK" // Kina
	CurrencyPHP Currency = "PHP" // Philippine Peso
	CurrencyPKR Currency = "PKR" // Pakistan Rupee
	CurrencyPLN Currency = "PLN" // Zloty
	CurrencyPYG Currency = "PYG" // Guarani
	CurrencyQAR Currency = "QAR" // Qatari Rial
	CurrencyRON Currency = "RON" // Romanian Leu
	CurrencyRSD Currency = "RSD" // Serbian Dinar
	CurrencyRUB Currency = "RUB" // Russian Ruble
	CurrencyRWF Currency = "RWF" // Rwanda Franc
	CurrencySAR Currency = "SAR" // Saudi Riyal
	CurrencySBD Currency = "SBD" // Solomon Islands Dollar
	CurrencySCR Currency = "SCR" // Seychelles Rupee
	CurrencySDG Currency = "SDG" // Sudanese Pound
	CurrencySEK Currency = "SEK" // Swedish Krona
	CurrencySGD Currency = "SGD" // Singapore Dollar
	CurrencySHP Currency = "SHP" // Saint Helena Pound
	CurrencySLE Currency = "SLE" // Leone
	CurrencySLL Currency = "SLL" // Leone
	CurrencySOS Currency = "SOS" // Somali Shilling
	CurrencySRD Currency = "SRD" // Surinam Dollar
	CurrencySSP Currency = "SSP" // South Sudanese Pound
	CurrencySTN Currency = "STN" // Dobra
	CurrencySVC Currency = "SVC" // El Salvador Colon
	CurrencySYP Currency = "SYP" // Syrian Pound
	CurrencySZL Currency = "SZL" // Lilangeni
	CurrencyTHB Currency = "THB" // Baht
	CurrencyTJS Currency = "TJS" // Somoni
	CurrencyTMT Currency = "TMT" // Turkmenistan New Manat
	CurrencyTND Currency = "TND" // Tunisian Dinar
	CurrencyTOP Currency = "TOP" // Pa'anga
	CurrencyTRY Currency = "TRY" // Turkish Lira
	CurrencyTTD Currency = "TTD" // Trinidad and Tobago Dollar
	CurrencyTWD Currency = "TWD" // New Taiwan Dollar
	CurrencyTZS Currency = "TZS" // Tanzanian Shilling
	CurrencyUAH Currency = "UAH" // Hryvnia
	CurrencyUGX Currency = "UGX" // Uganda Shilling
	CurrencyUSN Currency = "USN" // US Dollar (Next day)
	CurrencyUYI Currency = "UYI" // Uruguay Peso en Unidades Indexadas (UI)
	CurrencyUYU Currency = "UYU" // Peso Uruguayo
	CurrencyUYW Currency = "UYW" // Unidad Previsional
	CurrencyUZS Currency = "UZS" // Uzbekistan Sum
	CurrencyVED Currency = "VED" // Bolívar Soberano
	CurrencyVES Currency = "VES" // Bolívar Soberano
	CurrencyVND Currency = "VND" // Dong
	CurrencyVUV Currency = "VUV" // Vatu
	CurrencyWST Currency = "WST" // Tala
	CurrencyXAF Currency = "XAF" // CFA Franc BEAC
	CurrencyXAG Currency = "XAG" // Silver
	CurrencyXAU Currency = "XAU" // Gold
	CurrencyXBA Currency = "XBA" // Bond Markets Unit European Composite Unit (EURCO)
	CurrencyXBB Currency = "XBB" // Bond Markets Unit European Monetary Unit (E.M.U.-6)
	CurrencyXBC Currency = "XBC" // Bond Markets Unit European Unit of Account 9 (E.U.A.-9)
	CurrencyXBD Currency = "XBD" // Bond Markets Unit European Unit of Account 17 (E.U.A.-17)
	CurrencyXCD Currency = "XCD" // East Caribbean Dollar
	CurrencyXDR Currency = "XDR" // SDR (Special Drawing Right)
	CurrencyXOF Currency = "XOF" // CFA Franc BCEAO
	CurrencyXPD Currency = "XPD" // Palladium
	CurrencyXPF Currency = "XPF" // CFP Franc
	CurrencyXPT Currency = "XPT" // Platinum
	CurrencyXSU Currency = "XSU" // Sucre
	CurrencyXTS Currency = "XTS" // Codes specifically reserved for testing purposes
	CurrencyXUA Currency = "XUA" // ADB Unit of Account
	CurrencyXXX Currency = "XXX" // The codes assigned for transactions where no currency is involved
	CurrencyYER Currency = "YER" // Yemeni Rial
	CurrencyZAR Currency = "ZAR" // Rand
	CurrencyZMW Currency = "ZMW" // Zambian Kwacha
	CurrencyZWL Currency = "ZWL" // Zimbabwe Dollar
)

// Country 국가 코드 (ISO 3166-1 alpha-2)
type Country string

const (
	CountryAD Country = "AD" // Andorra
	CountryAE Country = "AE" // United Arab Emirates
	CountryAF Country = "AF" // Afghanistan
	CountryAG Country = "AG" // Antigua and Barbuda
	CountryAI Country = "AI" // Anguilla
	CountryAL Country = "AL" // Albania
	CountryAM Country = "AM" // Armenia
	CountryAO Country = "AO" // Angola
	CountryAQ Country = "AQ" // Antarctica
	CountryAR Country = "AR" // Argentina
	CountryAS Country = "AS" // American Samoa
	CountryAT Country = "AT" // Austria
	CountryAU Country = "AU" // Australia
	CountryAW Country = "AW" // Aruba
	CountryAX Country = "AX" // Åland Islands
	CountryAZ Country = "AZ" // Azerbaijan
	CountryBA Country = "BA" // Bosnia and Herzegovina
	CountryBB Country = "BB" // Barbados
	CountryBD Country = "BD" // Bangladesh
	CountryBE Country = "BE" // Belgium
	CountryBF Country = "BF" // Burkina Faso
	CountryBG Country = "BG" // Bulgaria
	CountryBH Country = "BH" // Bahrain
	CountryBI Country = "BI" // Burundi
	CountryBJ Country = "BJ" // Benin
	CountryBL Country = "BL" // Saint Barthélemy
	CountryBM Country = "BM" // Bermuda
	CountryBN Country = "BN" // Brunei Darussalam
	CountryBO Country = "BO" // Bolivia
	CountryBQ Country = "BQ" // Bonaire, Sint Eustatius and Saba
	CountryBR Country = "BR" // Brazil
	CountryBS Country = "BS" // Bahamas
	CountryBT Country = "BT" // Bhutan
	CountryBV Country = "BV" // Bouvet Island
	CountryBW Country = "BW" // Botswana
	CountryBY Country = "BY" // Belarus
	CountryBZ Country = "BZ" // Belize
	CountryCA Country = "CA" // Canada
	CountryCC Country = "CC" // Cocos (Keeling) Islands
	CountryCD Country = "CD" // Congo (Democratic Republic)
	CountryCF Country = "CF" // Central African Republic
	CountryCG Country = "CG" // Congo
	CountryCH Country = "CH" // Switzerland
	CountryCI Country = "CI" // Côte d'Ivoire
	CountryCK Country = "CK" // Cook Islands
	CountryCL Country = "CL" // Chile
	CountryCM Country = "CM" // Cameroon
	CountryCN Country = "CN" // China
	CountryCO Country = "CO" // Colombia
	CountryCR Country = "CR" // Costa Rica
	CountryCU Country = "CU" // Cuba
	CountryCV Country = "CV" // Cabo Verde
	CountryCW Country = "CW" // Curaçao
	CountryCX Country = "CX" // Christmas Island
	CountryCY Country = "CY" // Cyprus
	CountryCZ Country = "CZ" // Czechia
	CountryDE Country = "DE" // Germany
	CountryDJ Country = "DJ" // Djibouti
	CountryDK Country = "DK" // Denmark
	CountryDM Country = "DM" // Dominica
	CountryDO Country = "DO" // Dominican Republic
	CountryDZ Country = "DZ" // Algeria
	CountryEC Country = "EC" // Ecuador
	CountryEE Country = "EE" // Estonia
	CountryEG Country = "EG" // Egypt
	CountryEH Country = "EH" // Western Sahara
	CountryER Country = "ER" // Eritrea
	CountryES Country = "ES" // Spain
	CountryET Country = "ET" // Ethiopia
	CountryFI Country = "FI" // Finland
	CountryFJ Country = "FJ" // Fiji
	CountryFK Country = "FK" // Falkland Islands
	CountryFM Country = "FM" // Micronesia
	CountryFO Country = "FO" // Faroe Islands
	CountryFR Country = "FR" // France
	CountryGA Country = "GA" // Gabon
	CountryGB Country = "GB" // United Kingdom
	CountryGD Country = "GD" // Grenada
	CountryGE Country = "GE" // Georgia
	CountryGF Country = "GF" // French Guiana
	CountryGG Country = "GG" // Guernsey
	CountryGH Country = "GH" // Ghana
	CountryGI Country = "GI" // Gibraltar
	CountryGL Country = "GL" // Greenland
	CountryGM Country = "GM" // Gambia
	CountryGN Country = "GN" // Guinea
	CountryGP Country = "GP" // Guadeloupe
	CountryGQ Country = "GQ" // Equatorial Guinea
	CountryGR Country = "GR" // Greece
	CountryGS Country = "GS" // South Georgia and the South Sandwich Islands
	CountryGT Country = "GT" // Guatemala
	CountryGU Country = "GU" // Guam
	CountryGW Country = "GW" // Guinea-Bissau
	CountryGY Country = "GY" // Guyana
	CountryHK Country = "HK" // Hong Kong
	CountryHM Country = "HM" // Heard Island and McDonald Islands
	CountryHN Country = "HN" // Honduras
	CountryHR Country = "HR" // Croatia
	CountryHT Country = "HT" // Haiti
	CountryHU Country = "HU" // Hungary
	CountryID Country = "ID" // Indonesia
	CountryIE Country = "IE" // Ireland
	CountryIL Country = "IL" // Israel
	CountryIM Country = "IM" // Isle of Man
	CountryIN Country = "IN" // India
	CountryIO Country = "IO" // British Indian Ocean Territory
	CountryIQ Country = "IQ" // Iraq
	CountryIR Country = "IR" // Iran
	CountryIS Country = "IS" // Iceland
	CountryIT Country = "IT" // Italy
	CountryJE Country = "JE" // Jersey
	CountryJM Country = "JM" // Jamaica
	CountryJO Country = "JO" // Jordan
	CountryJP Country = "JP" // Japan
	CountryKE Country = "KE" // Kenya
	CountryKG Country = "KG" // Kyrgyzstan
	CountryKH Country = "KH" // Cambodia
	CountryKI Country = "KI" // Kiribati
	CountryKM Country = "KM" // Comoros
	CountryKN Country = "KN" // Saint Kitts and Nevis
	CountryKP Country = "KP" // North Korea
	CountryKR Country = "KR" // South Korea
	CountryKW Country = "KW" // Kuwait
	CountryKY Country = "KY" // Cayman Islands
	CountryKZ Country = "KZ" // Kazakhstan
	CountryLA Country = "LA" // Lao
	CountryLB Country = "LB" // Lebanon
	CountryLC Country = "LC" // Saint Lucia
	CountryLI Country = "LI" // Liechtenstein
	CountryLK Country = "LK" // Sri Lanka
	CountryLR Country = "LR" // Liberia
	CountryLS Country = "LS" // Lesotho
	CountryLT Country = "LT" // Lithuania
	CountryLU Country = "LU" // Luxembourg
	CountryLV Country = "LV" // Latvia
	CountryLY Country = "LY" // Libya
	CountryMA Country = "MA" // Morocco
	CountryMC Country = "MC" // Monaco
	CountryMD Country = "MD" // Moldova
	CountryME Country = "ME" // Montenegro
	CountryMF Country = "MF" // Saint Martin (French part)
	CountryMG Country = "MG" // Madagascar
	CountryMH Country = "MH" // Marshall Islands
	CountryMK Country = "MK" // North Macedonia
	CountryML Country = "ML" // Mali
	CountryMM Country = "MM" // Myanmar
	CountryMN Country = "MN" // Mongolia
	CountryMO Country = "MO" // Macao
	CountryMP Country = "MP" // Northern Mariana Islands
	CountryMQ Country = "MQ" // Martinique
	CountryMR Country = "MR" // Mauritania
	CountryMS Country = "MS" // Montserrat
	CountryMT Country = "MT" // Malta
	CountryMU Country = "MU" // Mauritius
	CountryMV Country = "MV" // Maldives
	CountryMW Country = "MW" // Malawi
	CountryMX Country = "MX" // Mexico
	CountryMY Country = "MY" // Malaysia
	CountryMZ Country = "MZ" // Mozambique
	CountryNA Country = "NA" // Namibia
	CountryNC Country = "NC" // New Caledonia
	CountryNE Country = "NE" // Niger
	CountryNF Country = "NF" // Norfolk Island
	CountryNG Country = "NG" // Nigeria
	CountryNI Country = "NI" // Nicaragua
	CountryNL Country = "NL" // Netherlands
	CountryNO Country = "NO" // Norway
	CountryNP Country = "NP" // Nepal
	CountryNR Country = "NR" // Nauru
	CountryNU Country = "NU" // Niue
	CountryNZ Country = "NZ" // New Zealand
	CountryOM Country = "OM" // Oman
	CountryPA Country = "PA" // Panama
	CountryPE Country = "PE" // Peru
	CountryPF Country = "PF" // French Polynesia
	CountryPG Country = "PG" // Papua New Guinea
	CountryPH Country = "PH" // Philippines
	CountryPK Country = "PK" // Pakistan
	CountryPL Country = "PL" // Poland
	CountryPM Country = "PM" // Saint Pierre and Miquelon
	CountryPN Country = "PN" // Pitcairn
	CountryPR Country = "PR" // Puerto Rico
	CountryPS Country = "PS" // Palestine
	CountryPT Country = "PT" // Portugal
	CountryPW Country = "PW" // Palau
	CountryPY Country = "PY" // Paraguay
	CountryQA Country = "QA" // Qatar
	CountryRE Country = "RE" // Réunion
	CountryRO Country = "RO" // Romania
	CountryRS Country = "RS" // Serbia
	CountryRU Country = "RU" // Russia
	CountryRW Country = "RW" // Rwanda
	CountrySA Country = "SA" // Saudi Arabia
	CountrySB Country = "SB" // Solomon Islands
	CountrySC Country = "SC" // Seychelles
	CountrySD Country = "SD" // Sudan
	CountrySE Country = "SE" // Sweden
	CountrySG Country = "SG" // Singapore
	CountrySH Country = "SH" // Saint Helena
	CountrySI Country = "SI" // Slovenia
	CountrySJ Country = "SJ" // Svalbard and Jan Mayen
	CountrySK Country = "SK" // Slovakia
	CountrySL Country = "SL" // Sierra Leone
	CountrySM Country = "SM" // San Marino
	CountrySN Country = "SN" // Senegal
	CountrySO Country = "SO" // Somalia
	CountrySR Country = "SR" // Suriname
	CountrySS Country = "SS" // South Sudan
	CountryST Country = "ST" // Sao Tome and Principe
	CountrySV Country = "SV" // El Salvador
	CountrySX Country = "SX" // Sint Maarten (Dutch part)
	CountrySY Country = "SY" // Syria
	CountrySZ Country = "SZ" // Eswatini
	CountryTC Country = "TC" // Turks and Caicos Islands
	CountryTD Country = "TD" // Chad
	CountryTF Country = "TF" // French Southern Territories
	CountryTG Country = "TG" // Togo
	CountryTH Country = "TH" // Thailand
	CountryTJ Country = "TJ" // Tajikistan
	CountryTK Country = "TK" // Tokelau
	CountryTL Country = "TL" // Timor-Leste
	CountryTM Country = "TM" // Turkmenistan
	CountryTN Country = "TN" // Tunisia
	CountryTO Country = "TO" // Tonga
	CountryTR Country = "TR" // Türkiye
	CountryTT Country = "TT" // Trinidad and Tobago
	CountryTV Country = "TV" // Tuvalu
	CountryTW Country = "TW" // Taiwan
	CountryTZ Country = "TZ" // Tanzania
	CountryUA Country = "UA" // Ukraine
	CountryUG Country = "UG" // Uganda
	CountryUM Country = "UM" // United States Minor Outlying Islands
	CountryUS Country = "US" // United States of America
	CountryUY Country = "UY" // Uruguay
	CountryUZ Country = "UZ" // Uzbekistan
	CountryVA Country = "VA" // Holy See
	CountryVC Country = "VC" // Saint Vincent and the Grenadines
	CountryVE Country = "VE" // Venezuela
	CountryVG Country = "VG" // Virgin Islands (British)
	CountryVI Country = "VI" // Virgin Islands (U.S.)
	CountryVN Country = "VN" // Viet Nam
	CountryVU Country = "VU" // Vanuatu
	CountryWF Country = "WF" // Wallis and Futuna
	CountryWS Country = "WS" // Samoa
	CountryYE Country = "YE" // Yemen
	CountryYT Country = "YT" // Mayotte
	CountryZA Country = "ZA" // South Africa
	CountryZM Country = "ZM" // Zambia
	CountryZW Country = "ZW" // Zimbabwe
)

// CardBrand 카드 브랜드
type CardBrand string

const (
	CardBrandLOCAL    CardBrand = "LOCAL"
	CardBrandMASTER   CardBrand = "MASTER"
	CardBrandUNIONPAY CardBrand = "UNIONPAY"
	CardBrandVISA     CardBrand = "VISA"
	CardBrandJCB      CardBrand = "JCB"
	CardBrandAMEX     CardBrand = "AMEX"
	CardBrandDINERS   CardBrand = "DINERS"
)

// CardType 카드 유형
type CardType string

const (
	CardTypeCREDIT CardType = "CREDIT" // 신용카드
	CardTypeDEBIT  CardType = "DEBIT"  // 체크카드
	CardTypeGIFT   CardType = "GIFT"   // 기프트카드
)

// CardOwnerType 카드 소유주 유형
type CardOwnerType string

const (
	CardOwnerTypePERSONAL  CardOwnerType = "PERSONAL"  // 개인
	CardOwnerTypeCORPORATE CardOwnerType = "CORPORATE" // 법인
)

// Gender 성별
type Gender string

const (
	GenderMALE   Gender = "MALE"   // 남성
	GenderFEMALE Gender = "FEMALE" // 여성
	GenderOTHER  Gender = "OTHER"  // 그 외 성별
)

// SortOrder 정렬 순서
type SortOrder string

const (
	SortOrderDESC SortOrder = "DESC"
	SortOrderASC  SortOrder = "ASC"
)

// PortOneVersion 포트원 버전
type PortOneVersion string

const (
	PortOneVersionV1 PortOneVersion = "V1"
	PortOneVersionV2 PortOneVersion = "V2"
)

// SelectedChannelType 채널 타입
type SelectedChannelType string

const (
	SelectedChannelTypeLIVE SelectedChannelType = "LIVE"
	SelectedChannelTypeTEST SelectedChannelType = "TEST"
)

// PaymentClientType 결제 클라이언트 타입
type PaymentClientType string

const (
	PaymentClientTypeSDK_MOBILE PaymentClientType = "SDK_MOBILE"
	PaymentClientTypeSDK_PC     PaymentClientType = "SDK_PC"
	PaymentClientTypeAPI        PaymentClientType = "API"
)

// PaymentProductType 상품 유형
type PaymentProductType string

const (
	PaymentProductTypePHYSICAL PaymentProductType = "PHYSICAL"
	PaymentProductTypeDIGITAL  PaymentProductType = "DIGITAL"
)

// CashReceiptType 현금영수증 유형
type CashReceiptType string

const (
	CashReceiptTypePERSONAL   CashReceiptType = "PERSONAL"   // 소득공제용
	CashReceiptTypeCORPORATE  CashReceiptType = "CORPORATE"  // 지출증빙용
	CashReceiptTypeANONYMOUS  CashReceiptType = "ANONYMOUS"  // 자진발급
)

// CashReceiptInputType 현금영수증 입력 유형
type CashReceiptInputType string

const (
	CashReceiptInputTypePERSONAL  CashReceiptInputType = "PERSONAL"
	CashReceiptInputTypeCORPORATE CashReceiptInputType = "CORPORATE"
	CashReceiptInputTypeNO_RECEIPT CashReceiptInputType = "NO_RECEIPT"
)

// PaymentMethodType 결제 수단 유형
type PaymentMethodType string

const (
	PaymentMethodTypeCARD               PaymentMethodType = "CARD"
	PaymentMethodTypeTRANSFER           PaymentMethodType = "TRANSFER"
	PaymentMethodTypeVIRTUAL_ACCOUNT    PaymentMethodType = "VIRTUAL_ACCOUNT"
	PaymentMethodTypeGIFT_CERTIFICATE   PaymentMethodType = "GIFT_CERTIFICATE"
	PaymentMethodTypeMOBILE             PaymentMethodType = "MOBILE"
	PaymentMethodTypeEASY_PAY           PaymentMethodType = "EASY_PAY"
)

// Bank 은행
type Bank string

const (
	BankBANK_OF_KOREA         Bank = "BANK_OF_KOREA"         // 한국은행
	BankKDB                   Bank = "KDB"                   // 산업은행
	BankIBK                   Bank = "IBK"                   // 기업은행
	BankKOOKMIN               Bank = "KOOKMIN"               // 국민은행
	BankSUHYUP                Bank = "SUHYUP"                // 수협은행
	BankKEXIM                 Bank = "KEXIM"                 // 수출입은행
	BankNONGHYUP              Bank = "NONGHYUP"              // NH농협은행
	BankLOCAL_NONGHYUP        Bank = "LOCAL_NONGHYUP"        // 지역농축협
	BankWOORI                 Bank = "WOORI"                 // 우리은행
	BankSTANDARD_CHARTERED    Bank = "STANDARD_CHARTERED"    // SC제일은행
	BankCITI                  Bank = "CITI"                  // 한국씨티은행
	BankSUHYUP_FEDERATION     Bank = "SUHYUP_FEDERATION"     // 수협중앙회
	BankDAEGU                 Bank = "DAEGU"                 // 아이엠뱅크
	BankBUSAN                 Bank = "BUSAN"                 // 부산은행
	BankKWANGJU               Bank = "KWANGJU"               // 광주은행
	BankJEJU                  Bank = "JEJU"                  // 제주은행
	BankJEONBUK               Bank = "JEONBUK"               // 전북은행
	BankKYONGNAM              Bank = "KYONGNAM"              // 경남은행
	BankKFCC                  Bank = "KFCC"                  // 새마을금고
	BankSHINHYUP              Bank = "SHINHYUP"              // 신협
	BankSAVINGS_BANK          Bank = "SAVINGS_BANK"          // 저축은행
	BankMORGAN_STANLEY        Bank = "MORGAN_STANLEY"        // 모간스탠리은행
	BankHSBC                  Bank = "HSBC"                  // HSBC은행
	BankDEUTSCHE              Bank = "DEUTSCHE"              // 도이치은행
	BankJPMC                  Bank = "JPMC"                  // 제이피모간체이스은행
	BankMIZUHO                Bank = "MIZUHO"                // 미즈호은행
	BankMUFG                  Bank = "MUFG"                  // 엠유에프지은행
	BankBANK_OF_AMERICA       Bank = "BANK_OF_AMERICA"       // BOA은행
	BankBNP_PARIBAS           Bank = "BNP_PARIBAS"           // 비엔피파리바은행
	BankICBC                  Bank = "ICBC"                  // 중국공상은행
	BankBANK_OF_CHINA         Bank = "BANK_OF_CHINA"         // 중국은행
	BankNFCF                  Bank = "NFCF"                  // 산림조합중앙회
	BankUOB                   Bank = "UOB"                   // 대화은행
	BankBOCOM                 Bank = "BOCOM"                 // 교통은행
	BankCCB                   Bank = "CCB"                   // 중국건설은행
	BankPOST                  Bank = "POST"                  // 우체국
	BankKODIT                 Bank = "KODIT"                 // 신용보증기금
	BankKIBO                  Bank = "KIBO"                  // 기술보증기금
	BankHANA                  Bank = "HANA"                  // 하나은행
	BankSHINHAN               Bank = "SHINHAN"               // 신한은행
	BankK_BANK                Bank = "K_BANK"                // 케이뱅크
	BankKAKAO                 Bank = "KAKAO"                 // 카카오뱅크
	BankTOSS                  Bank = "TOSS"                  // 토스뱅크
	BankMISC_FOREIGN          Bank = "MISC_FOREIGN"          // 기타 외국계은행
	BankSGI                   Bank = "SGI"                   // 서울보증보험
	BankKCIS                  Bank = "KCIS"                  // 한국신용정보원
	BankYUANTA_SECURITIES     Bank = "YUANTA_SECURITIES"     // 유안타증권
	BankKB_SECURITIES         Bank = "KB_SECURITIES"         // KB증권
	BankSANGSANGIN_SECURITIES Bank = "SANGSANGIN_SECURITIES" // 상상인증권
	BankHANYANG_SECURITIES    Bank = "HANYANG_SECURITIES"    // 한양증권
	BankLEADING_SECURITIES    Bank = "LEADING_SECURITIES"    // 리딩투자증권
	BankBNK_SECURITIES        Bank = "BNK_SECURITIES"        // BNK투자증권
	BankIBK_SECURITIES        Bank = "IBK_SECURITIES"        // IBK투자증권
	BankDAOL_SECURITIES       Bank = "DAOL_SECURITIES"       // 다올투자증권
	BankMIRAE_ASSET_SECURITIES Bank = "MIRAE_ASSET_SECURITIES" // 미래에셋증권
	BankSAMSUNG_SECURITIES    Bank = "SAMSUNG_SECURITIES"    // 삼성증권
	BankKOREA_SECURITIES      Bank = "KOREA_SECURITIES"      // 한국투자증권
	BankNH_SECURITIES         Bank = "NH_SECURITIES"         // NH투자증권
	BankKYOBO_SECURITIES      Bank = "KYOBO_SECURITIES"      // 교보증권
	BankHI_SECURITIES         Bank = "HI_SECURITIES"         // 하이투자증권
	BankHYUNDAI_MOTOR_SECURITIES Bank = "HYUNDAI_MOTOR_SECURITIES" // 현대차증권
	BankKIWOOM_SECURITIES     Bank = "KIWOOM_SECURITIES"     // 키움증권
	BankEBEST_SECURITIES      Bank = "EBEST_SECURITIES"      // LS증권
	BankSK_SECURITIES         Bank = "SK_SECURITIES"         // SK증권
	BankDAISHIN_SECURITIES    Bank = "DAISHIN_SECURITIES"    // 대신증권
	BankHANHWA_SECURITIES     Bank = "HANHWA_SECURITIES"     // 한화투자증권
	BankHANA_SECURITIES       Bank = "HANA_SECURITIES"       // 하나증권
	BankTOSS_SECURITIES       Bank = "TOSS_SECURITIES"       // 토스증권
	BankSHINHAN_SECURITIES    Bank = "SHINHAN_SECURITIES"    // 신한투자증권
	BankDB_SECURITIES         Bank = "DB_SECURITIES"         // DB금융투자
	BankEUGENE_SECURITIES     Bank = "EUGENE_SECURITIES"     // 유진투자증권
	BankMERITZ_SECURITIES     Bank = "MERITZ_SECURITIES"     // 메리츠증권
	BankKAKAO_PAY_SECURITIES  Bank = "KAKAO_PAY_SECURITIES"  // 카카오페이증권
	BankBOOKOOK_SECURITIES    Bank = "BOOKOOK_SECURITIES"    // 부국증권
	BankSHINYOUNG_SECURITIES  Bank = "SHINYOUNG_SECURITIES"  // 신영증권
	BankCAPE_SECURITIES       Bank = "CAPE_SECURITIES"       // 케이프투자증권
	BankKOREA_SECURITIES_FINANCE Bank = "KOREA_SECURITIES_FINANCE" // 한국증권금융
	BankKOREA_FOSS_SECURITIES Bank = "KOREA_FOSS_SECURITIES" // 한국포스증권
	BankWOORI_INVESTMENT_BANK Bank = "WOORI_INVESTMENT_BANK" // 우리종합금융
)

// PgProvider PG사 결제 모듈
type PgProvider string

const (
	PgProviderHTML5_INICIS     PgProvider = "HTML5_INICIS"
	PgProviderPAYPAL           PgProvider = "PAYPAL"
	PgProviderPAYPAL_V2        PgProvider = "PAYPAL_V2"
	PgProviderINICIS           PgProvider = "INICIS"
	PgProviderDANAL            PgProvider = "DANAL"
	PgProviderNICE             PgProvider = "NICE"
	PgProviderDANAL_TPAY       PgProvider = "DANAL_TPAY"
	PgProviderJTNET            PgProvider = "JTNET"
	PgProviderUPLUS            PgProvider = "UPLUS"
	PgProviderNAVERPAY         PgProvider = "NAVERPAY"
	PgProviderKAKAO            PgProvider = "KAKAO"
	PgProviderSETTLE           PgProvider = "SETTLE"
	PgProviderKCP              PgProvider = "KCP"
	PgProviderMOBILIANS        PgProvider = "MOBILIANS"
	PgProviderKAKAOPAY         PgProvider = "KAKAOPAY"
	PgProviderNAVERCO          PgProvider = "NAVERCO"
	PgProviderSYRUP            PgProvider = "SYRUP"
	PgProviderKICC             PgProvider = "KICC"
	PgProviderEXIMBAY          PgProvider = "EXIMBAY"
	PgProviderSMILEPAY         PgProvider = "SMILEPAY"
	PgProviderPAYCO            PgProvider = "PAYCO"
	PgProviderKCP_BILLING      PgProvider = "KCP_BILLING"
	PgProviderALIPAY           PgProvider = "ALIPAY"
	PgProviderPAYPLE           PgProvider = "PAYPLE"
	PgProviderCHAI             PgProvider = "CHAI"
	PgProviderBLUEWALNUT       PgProvider = "BLUEWALNUT"
	PgProviderSMARTRO          PgProvider = "SMARTRO"
	PgProviderSMARTRO_V2       PgProvider = "SMARTRO_V2"
	PgProviderPAYMENTWALL      PgProvider = "PAYMENTWALL"
	PgProviderTOSSPAYMENTS     PgProvider = "TOSSPAYMENTS"
	PgProviderKCP_QUICK        PgProvider = "KCP_QUICK"
	PgProviderDAOU             PgProvider = "DAOU"
	PgProviderGALAXIA          PgProvider = "GALAXIA"
	PgProviderTOSSPAY          PgProvider = "TOSSPAY"
	PgProviderKCP_DIRECT       PgProvider = "KCP_DIRECT"
	PgProviderSETTLE_ACC       PgProvider = "SETTLE_ACC"
	PgProviderSETTLE_FIRM      PgProvider = "SETTLE_FIRM"
	PgProviderINICIS_UNIFIED   PgProvider = "INICIS_UNIFIED"
	PgProviderKSNET            PgProvider = "KSNET"
	PgProviderPINPAY           PgProvider = "PINPAY"
	PgProviderNICE_V2          PgProvider = "NICE_V2"
	PgProviderTOSS_BRANDPAY    PgProvider = "TOSS_BRANDPAY"
	PgProviderWELCOME          PgProvider = "WELCOME"
	PgProviderTOSSPAY_V2       PgProvider = "TOSSPAY_V2"
	PgProviderINICIS_V2        PgProvider = "INICIS_V2"
	PgProviderKPN              PgProvider = "KPN"
	PgProviderKCP_V2           PgProvider = "KCP_V2"
	PgProviderHYPHEN           PgProvider = "HYPHEN"
	PgProviderEXIMBAY_V2       PgProvider = "EXIMBAY_V2"
	PgProviderINICIS_JP        PgProvider = "INICIS_JP"
	PgProviderPAYLETTER_GLOBAL PgProvider = "PAYLETTER_GLOBAL"
)

// EasyPayProvider 간편결제사
type EasyPayProvider string

const (
	EasyPayProviderPAYCO       EasyPayProvider = "PAYCO"
	EasyPayProviderSAMSUNGPAY  EasyPayProvider = "SAMSUNGPAY"
	EasyPayProviderSSGPAY      EasyPayProvider = "SSGPAY"
	EasyPayProviderKAKAOPAY    EasyPayProvider = "KAKAOPAY"
	EasyPayProviderNAVERPAY    EasyPayProvider = "NAVERPAY"
	EasyPayProviderCHAIPAY     EasyPayProvider = "CHAIPAY"
	EasyPayProviderLPAY        EasyPayProvider = "LPAY"
	EasyPayProviderKBANK       EasyPayProvider = "KBANK"
	EasyPayProviderTOSSPAY     EasyPayProvider = "TOSSPAY"
	EasyPayProviderAPPLEPAY    EasyPayProvider = "APPLEPAY"
	EasyPayProviderPINPAY      EasyPayProvider = "PINPAY"
	EasyPayProviderSKPAY       EasyPayProvider = "SKPAY"
	EasyPayProviderHYPHEN      EasyPayProvider = "HYPHEN"
	EasyPayProviderKCP_QUICKPAY EasyPayProvider = "KCP_QUICKPAY"
	EasyPayProviderWOORI       EasyPayProvider = "WOORI"
)

// PgCompany PG사
type PgCompany string

const (
	PgCompanyINICIS       PgCompany = "INICIS"
	PgCompanyNICE         PgCompany = "NICE"
	PgCompanyKCP          PgCompany = "KCP"
	PgCompanyDANAL        PgCompany = "DANAL"
	PgCompanyTOSSPAYMENTS PgCompany = "TOSSPAYMENTS"
	PgCompanyKICC         PgCompany = "KICC"
	PgCompanySMARTRO      PgCompany = "SMARTRO"
	PgCompanyDAOU         PgCompany = "DAOU"
	PgCompanyBLUEWALNUT   PgCompany = "BLUEWALNUT"
	PgCompanySETTLE       PgCompany = "SETTLE"
	PgCompanyGALAXIA      PgCompany = "GALAXIA"
	PgCompanyPAYPAL       PgCompany = "PAYPAL"
	PgCompanyEXIMBAY      PgCompany = "EXIMBAY"
	PgCompanyKSNET        PgCompany = "KSNET"
	PgCompanyWELCOME      PgCompany = "WELCOME"
	PgCompanyKPN          PgCompany = "KPN"
	PgCompanyHYPHEN       PgCompany = "HYPHEN"
	PgCompanyMOBILIANS    PgCompany = "MOBILIANS"
	PgCompanyJTNET        PgCompany = "JTNET"
	PgCompanyPAYLETTER    PgCompany = "PAYLETTER"
)

// Locale 로케일
type Locale string

const (
	LocaleKO_KR Locale = "KO_KR"
	LocaleEN_US Locale = "EN_US"
)

// Card 카드 상세 정보
type Card struct {
	// 발행사 코드
	Publisher *string `json:"publisher,omitempty"`
	// 발급사 코드
	Issuer *string `json:"issuer,omitempty"`
	// 카드 브랜드
	Brand *CardBrand `json:"brand,omitempty"`
	// 카드 유형
	Type *CardType `json:"type,omitempty"`
	// 카드 소유주 유형
	OwnerType *CardOwnerType `json:"ownerType,omitempty"`
	// 카드 번호 앞 6자리 또는 8자리의 BIN (Bank Identification Number)
	Bin *string `json:"bin,omitempty"`
	// 카드 상품명
	Name *string `json:"name,omitempty"`
	// 마스킹된 카드 번호
	Number *string `json:"number,omitempty"`
}

// CardCredential 카드 인증 정보
type CardCredential struct {
	// 카드 번호
	Number string `json:"number"`
	// 유효 기간 (MMYY)
	ExpiryMonth string `json:"expiryMonth"`
	ExpiryYear  string `json:"expiryYear"`
	// 생년월일 (YYMMDD) 또는 사업자등록번호 (10자리)
	BirthOrBusinessRegistrationNumber *string `json:"birthOrBusinessRegistrationNumber,omitempty"`
	// 카드 비밀번호 앞 두자리
	PasswordTwoDigits *string `json:"passwordTwoDigits,omitempty"`
}

// AddressType 주소 타입
type AddressType string

const (
	AddressTypeONE_LINE  AddressType = "ONE_LINE"
	AddressTypeSEPARATED AddressType = "SEPARATED"
)

// Address 주소 (discriminated union)
type Address struct {
	Type AddressType `json:"type"`
	// 주소 (한 줄)
	OneLine string `json:"oneLine"`
	// 상세 주소 1 (SEPARATED 타입에서만 사용)
	AddressLine1 *string `json:"addressLine1,omitempty"`
	// 상세 주소 2 (SEPARATED 타입에서만 사용)
	AddressLine2 *string `json:"addressLine2,omitempty"`
	// 시/군/구 (SEPARATED 타입에서만 사용)
	City *string `json:"city,omitempty"`
	// 주/도/시 (SEPARATED 타입에서만 사용)
	Province *string `json:"province,omitempty"`
	// 국가 (SEPARATED 타입에서만 사용)
	Country *Country `json:"country,omitempty"`
}

// IsOneLine 한 줄 형식 주소인지 확인
func (a *Address) IsOneLine() bool {
	return a.Type == AddressTypeONE_LINE
}

// IsSeparated 분리 형식 주소인지 확인
func (a *Address) IsSeparated() bool {
	return a.Type == AddressTypeSEPARATED
}

// SeparatedAddressInput 분리 형식 주소 입력
type SeparatedAddressInput struct {
	// 상세 주소 1
	AddressLine1 string `json:"addressLine1"`
	// 상세 주소 2
	AddressLine2 string `json:"addressLine2"`
	// 시/군/구
	City *string `json:"city,omitempty"`
	// 주/도/시
	Province *string `json:"province,omitempty"`
	// 국가
	Country *Country `json:"country,omitempty"`
}

// Customer 고객 정보
type Customer struct {
	// 고객 아이디 (고객사가 지정한 고객의 고유 식별자)
	ID *string `json:"id,omitempty"`
	// 이름
	Name *string `json:"name,omitempty"`
	// 출생 연도
	BirthYear *string `json:"birthYear,omitempty"`
	// 출생 월
	BirthMonth *string `json:"birthMonth,omitempty"`
	// 출생 일자
	BirthDay *string `json:"birthDay,omitempty"`
	// 성별
	Gender *Gender `json:"gender,omitempty"`
	// 이메일
	Email *string `json:"email,omitempty"`
	// 전화번호
	PhoneNumber *string `json:"phoneNumber,omitempty"`
	// 주소
	Address *Address `json:"address,omitempty"`
	// 우편번호
	Zipcode *string `json:"zipcode,omitempty"`
}

// CustomerNameInput 고객 이름 입력
type CustomerNameInput struct {
	// 전체 이름
	Full *string `json:"full,omitempty"`
	// 이름 (분리된 형식)
	Separated *CustomerSeparatedName `json:"separated,omitempty"`
}

// CustomerSeparatedName 분리된 고객 이름
type CustomerSeparatedName struct {
	// 이름
	First *string `json:"first,omitempty"`
	// 성
	Last *string `json:"last,omitempty"`
}

// CustomerInput 고객 정보 입력
type CustomerInput struct {
	// 고객 아이디
	ID *string `json:"id,omitempty"`
	// 이름
	Name *CustomerNameInput `json:"name,omitempty"`
	// 출생 연도
	BirthYear *string `json:"birthYear,omitempty"`
	// 출생 월
	BirthMonth *string `json:"birthMonth,omitempty"`
	// 출생 일자
	BirthDay *string `json:"birthDay,omitempty"`
	// 성별
	Gender *Gender `json:"gender,omitempty"`
	// 이메일
	Email *string `json:"email,omitempty"`
	// 전화번호
	PhoneNumber *string `json:"phoneNumber,omitempty"`
	// 주소
	Address *SeparatedAddressInput `json:"address,omitempty"`
	// 우편번호
	Zipcode *string `json:"zipcode,omitempty"`
}

// ChannelGroupSummary 채널 그룹 요약 정보
type ChannelGroupSummary struct {
	// 채널 그룹 아이디
	ID string `json:"id"`
	// 채널 그룹 이름
	Name string `json:"name"`
	// 테스트 채널 그룹 여부
	IsForTest bool `json:"isForTest"`
}

// SelectedChannel 선택된 채널
type SelectedChannel struct {
	// 채널 타입
	Type SelectedChannelType `json:"type"`
	// 채널 아이디
	ID *string `json:"id,omitempty"`
	// 채널 키
	Key *string `json:"key,omitempty"`
	// 채널 이름
	Name *string `json:"name,omitempty"`
	// PG사
	PgProvider PgProvider `json:"pgProvider"`
	// PG사 상점 아이디
	PgMerchantId string `json:"pgMerchantId"`
}

// PageInfo 페이지 정보
type PageInfo struct {
	// 현재 페이지 번호
	Number int `json:"number"`
	// 페이지당 항목 수
	Size int `json:"size"`
	// 전체 항목 수
	TotalCount int `json:"totalCount"`
}

// PageInput 페이지 입력
type PageInput struct {
	// 페이지 번호 (0부터 시작)
	Number *int `json:"number,omitempty"`
	// 페이지당 항목 수
	Size *int `json:"size,omitempty"`
}

// DateTimeRange 시간 범위
type DateTimeRange struct {
	// 시작 시간 (RFC 3339 형식)
	From string `json:"from"`
	// 종료 시간 (RFC 3339 형식)
	Until string `json:"until"`
}

// PaymentAmountInput 결제 금액 입력
type PaymentAmountInput struct {
	// 총 결제 금액
	Total int64 `json:"total"`
	// 면세 금액
	TaxFree *int64 `json:"taxFree,omitempty"`
	// 부가세
	Vat *int64 `json:"vat,omitempty"`
}

// PaymentProduct 결제 상품
type PaymentProduct struct {
	// 상품 아이디
	ID string `json:"id"`
	// 상품 이름
	Name string `json:"name"`
	// 상품 코드
	Code *string `json:"code,omitempty"`
	// 상품 단가
	Amount int64 `json:"amount"`
	// 수량
	Quantity int `json:"quantity"`
	// 상품 태그
	Tag *string `json:"tag,omitempty"`
}

// CashReceiptInput 현금영수증 입력
type CashReceiptInput struct {
	// 현금영수증 유형
	Type CashReceiptInputType `json:"type"`
	// 개인: 휴대폰 번호 또는 현금영수증 카드 번호
	// 법인: 사업자 등록 번호
	CustomerIdentityNumber *string `json:"customerIdentityNumber,omitempty"`
}

// BillingKeyPaymentInput 빌링키 결제 입력
type BillingKeyPaymentInput struct {
	// 빌링키
	BillingKey string `json:"billingKey"`
	// 주문 이름
	OrderName string `json:"orderName"`
	// 결제 금액
	Amount PaymentAmountInput `json:"amount"`
	// 통화
	Currency Currency `json:"currency"`
	// 할부 개월 수
	InstallMonth *int `json:"installMonth,omitempty"`
	// 무이자 할부 사용 여부
	UseFreeInterestFromMerchant *bool `json:"useFreeInterestFromMerchant,omitempty"`
	// 카드 포인트 사용 여부
	UseCardPoint *bool `json:"useCardPoint,omitempty"`
	// 현금영수증 입력
	CashReceipt *CashReceiptInput `json:"cashReceipt,omitempty"`
	// 국가
	Country *Country `json:"country,omitempty"`
	// 결제 알림 URL
	NoticeUrls []string `json:"noticeUrls,omitempty"`
	// 상품 목록
	Products []PaymentProduct `json:"products,omitempty"`
	// 상품 수
	ProductCount *int `json:"productCount,omitempty"`
	// 상품 유형
	ProductType *PaymentProductType `json:"productType,omitempty"`
	// 배송지 주소
	ShippingAddress *SeparatedAddressInput `json:"shippingAddress,omitempty"`
	// 프로모션 아이디
	PromotionId *string `json:"promotionId,omitempty"`
	// 결제 유저 아이피
	Bypass *json.RawMessage `json:"bypass,omitempty"`
}

// BankInfoName 은행 명칭
type BankInfoName struct {
	// 한국어 명칭
	Ko string `json:"ko"`
}

// BankInfo 은행 정보
type BankInfo struct {
	// 은행
	Bank Bank `json:"bank"`
	// 언어별 명칭
	Name BankInfoName `json:"name"`
}

// GetBankInfosResponse 은행 정보 조회 응답
type GetBankInfosResponse struct {
	// 조회된 은행 정보 리스트
	Items []BankInfo `json:"items"`
}
