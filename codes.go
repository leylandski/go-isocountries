package isocountries

import (
	"unicode/utf8"
)

/*
CountryCode is a named string type to hold ISO 3166-1 alpha-3 country code constants.
*/
type CountryCode string

const (
	Aruba                                  CountryCode = "ABW"
	Afghanistan                            CountryCode = "AFG"
	Angola                                 CountryCode = "AGO"
	Anguilla                               CountryCode = "AIA"
	AlandIslands                           CountryCode = "ALA"
	Albania                                CountryCode = "ALB"
	Andorra                                CountryCode = "AND"
	UnitedArabEmirates                     CountryCode = "ARE"
	Argentina                              CountryCode = "ARG"
	Armenia                                CountryCode = "ARM"
	AmericanSamoa                          CountryCode = "ASM"
	Antarctica                             CountryCode = "ATA"
	FrenchSouthernTerritories              CountryCode = "ATF"
	AntiguaAndBarbuda                      CountryCode = "ATG"
	Australia                              CountryCode = "AUS"
	Austria                                CountryCode = "AUT"
	Azerbaijan                             CountryCode = "AZE"
	Burundi                                CountryCode = "BDI"
	Belgium                                CountryCode = "BEL"
	Benin                                  CountryCode = "BEN"
	BonaireSintEustatiusAndSaba            CountryCode = "BES"
	BurkinaFaso                            CountryCode = "BFA"
	Bangladesh                             CountryCode = "BGD"
	Bulgaria                               CountryCode = "BGR"
	Bahrain                                CountryCode = "BHR"
	Bahamas                                CountryCode = "BHS"
	BosniaAndHerzegovina                   CountryCode = "BIH"
	SaintBarthelemy                        CountryCode = "BLM"
	Belarus                                CountryCode = "BLR"
	Belize                                 CountryCode = "BLZ"
	Bermuda                                CountryCode = "BMU"
	Bolivia                                CountryCode = "BOL"
	Brazil                                 CountryCode = "BRA"
	Barbados                               CountryCode = "BRB"
	BruneiDarussalam                       CountryCode = "BRN"
	Bhutan                                 CountryCode = "BTN"
	BouvetIsland                           CountryCode = "BVT"
	Botswana                               CountryCode = "BWA"
	CentralAfricanRepublic                 CountryCode = "CAF"
	Canada                                 CountryCode = "CAN"
	CocosIslands                           CountryCode = "CCK"
	Switzerland                            CountryCode = "CHE"
	Chile                                  CountryCode = "CHL"
	China                                  CountryCode = "CHN"
	CotedIvoire                            CountryCode = "CIV"
	Cameroon                               CountryCode = "CMR"
	CongoDR                                CountryCode = "COD"
	Congo                                  CountryCode = "COG"
	CookIslands                            CountryCode = "COK"
	Colombia                               CountryCode = "COL"
	Comoros                                CountryCode = "COM"
	CaboVerde                              CountryCode = "CPV"
	CostaRica                              CountryCode = "CRI"
	Cuba                                   CountryCode = "CUB"
	Curacao                                CountryCode = "CUW"
	ChristmasIsland                        CountryCode = "CXR"
	CaymanIslands                          CountryCode = "CYM"
	Cyprus                                 CountryCode = "CYP"
	CzechRepublic                          CountryCode = "CZE"
	Germany                                CountryCode = "DEU"
	Djibouti                               CountryCode = "DJI"
	Dominica                               CountryCode = "DMA"
	Denmark                                CountryCode = "DNK"
	DominicanRepublic                      CountryCode = "DOM"
	Algeria                                CountryCode = "DZA"
	Ecuador                                CountryCode = "ECU"
	Egypt                                  CountryCode = "EGY"
	Eritrea                                CountryCode = "ERI"
	WesternSahara                          CountryCode = "ESH"
	Spain                                  CountryCode = "ESP"
	Estonia                                CountryCode = "EST"
	Ethiopia                               CountryCode = "ETH"
	Finland                                CountryCode = "FIN"
	Fiji                                   CountryCode = "FJI"
	FalklandIslands                        CountryCode = "FLK"
	France                                 CountryCode = "FRA"
	FaroeIslands                           CountryCode = "FRO"
	FederatedStatesOfMicronesia            CountryCode = "FSM"
	Gabon                                  CountryCode = "GAB"
	UnitedKingdom                          CountryCode = "GBR"
	Georgia                                CountryCode = "GEO"
	Guernsey                               CountryCode = "GGY"
	Ghana                                  CountryCode = "GHA"
	Gibraltar                              CountryCode = "GIB"
	Guinea                                 CountryCode = "GIN"
	Guadeloupe                             CountryCode = "GLP"
	Gambia                                 CountryCode = "GMB"
	GuineaBissau                           CountryCode = "GNB"
	EquatorialGuinea                       CountryCode = "GNQ"
	Greece                                 CountryCode = "GRC"
	Grenada                                CountryCode = "GRD"
	Greenland                              CountryCode = "GRL"
	Guatemala                              CountryCode = "GTM"
	FrenchGuiana                           CountryCode = "GUF"
	Guam                                   CountryCode = "GUM"
	Guyana                                 CountryCode = "GUY"
	HongKong                               CountryCode = "HKG"
	HeardIslandsAndMcDonaldIslands         CountryCode = "HMD"
	Honduras                               CountryCode = "HND"
	Croatia                                CountryCode = "HRV"
	Haiti                                  CountryCode = "HTI"
	Hungary                                CountryCode = "HUN"
	Indonesia                              CountryCode = "IDN"
	IsleOfMan                              CountryCode = "IMN"
	India                                  CountryCode = "IND"
	BritishIndianOceanTerritory            CountryCode = "IOT"
	Ireland                                CountryCode = "IRL"
	Iran                                   CountryCode = "IRN"
	Iraq                                   CountryCode = "IRQ"
	Iceland                                CountryCode = "ISL"
	Israel                                 CountryCode = "ISR"
	Italy                                  CountryCode = "ITA"
	Jamaica                                CountryCode = "JAM"
	Jersey                                 CountryCode = "JEY"
	Jordan                                 CountryCode = "JOR"
	Japan                                  CountryCode = "JPN"
	Kazakhstan                             CountryCode = "KAZ"
	Kenya                                  CountryCode = "KEN"
	Kyrgyzstan                             CountryCode = "KGZ"
	Cambodia                               CountryCode = "KHM"
	Kiribati                               CountryCode = "KIR"
	SaintKittsAndNevis                     CountryCode = "KNA"
	KoreaRepublic                          CountryCode = "KOR"
	Kuwait                                 CountryCode = "KWT"
	Laos                                   CountryCode = "LAO"
	Lebanon                                CountryCode = "LBN"
	Liberia                                CountryCode = "LBR"
	Libya                                  CountryCode = "LBY"
	SaintLucia                             CountryCode = "LCA"
	Liechtenstein                          CountryCode = "LIE"
	SriLanka                               CountryCode = "LKA"
	Lesotho                                CountryCode = "LSO"
	Lithuania                              CountryCode = "LTU"
	Luxembourg                             CountryCode = "LUX"
	Latvia                                 CountryCode = "LVA"
	Macao                                  CountryCode = "MAC"
	SaintMartinFr                          CountryCode = "MAF"
	Morocco                                CountryCode = "MAR"
	Monaco                                 CountryCode = "MCO"
	Moldova                                CountryCode = "MDA"
	Madagascar                             CountryCode = "MDG"
	Maldives                               CountryCode = "MDV"
	Mexico                                 CountryCode = "MAX"
	MarshallIslands                        CountryCode = "MHL"
	MacedoniaFYR                           CountryCode = "MKD"
	Mali                                   CountryCode = "MLI"
	Malta                                  CountryCode = "MLT"
	Myanmar                                CountryCode = "MMR"
	Montenegro                             CountryCode = "MNE"
	Mongolia                               CountryCode = "MNG"
	NorthernMarianaIslands                 CountryCode = "MNP"
	Mozambique                             CountryCode = "MOZ"
	Mauritania                             CountryCode = "MRT"
	Montserrat                             CountryCode = "MSR"
	Martinique                             CountryCode = "MTQ"
	Mauritius                              CountryCode = "MUS"
	Malawi                                 CountryCode = "MWI"
	Malaysia                               CountryCode = "MYS"
	Mayotte                                CountryCode = "MYT"
	Namibia                                CountryCode = "NAM"
	NewCaledonia                           CountryCode = "NCL"
	Niger                                  CountryCode = "NER"
	NorfolkIsland                          CountryCode = "NFK"
	Nigeria                                CountryCode = "NGA"
	Nicaragua                              CountryCode = "NIC"
	Niue                                   CountryCode = "NIU"
	Netherlands                            CountryCode = "NLD"
	Norway                                 CountryCode = "NOR"
	Nepal                                  CountryCode = "NPL"
	Nauru                                  CountryCode = "NRU"
	NewZealand                             CountryCode = "NZL"
	Oman                                   CountryCode = "OMN"
	Pakistan                               CountryCode = "PAK"
	Panama                                 CountryCode = "PAN"
	Pitcairn                               CountryCode = "PCN"
	Peru                                   CountryCode = "PER"
	Philippines                            CountryCode = "PHL"
	Palau                                  CountryCode = "PLW"
	PapuaNewGuinea                         CountryCode = "PNG"
	Poland                                 CountryCode = "POL"
	PuertoRico                             CountryCode = "PRI"
	KoreaDPR                               CountryCode = "PRK"
	Portugal                               CountryCode = "PRT"
	Paraguay                               CountryCode = "PRY"
	Palestine                              CountryCode = "PSE"
	FrenchPolynesia                        CountryCode = "PYF"
	Qatar                                  CountryCode = "QAT"
	Reunion                                CountryCode = "REU"
	Romania                                CountryCode = "ROU"
	RussianFederation                      CountryCode = "RUS"
	Rwanda                                 CountryCode = "RWA"
	SaudiArabia                            CountryCode = "SAU"
	Sudan                                  CountryCode = "SDN"
	Senegal                                CountryCode = "SEN"
	Singapore                              CountryCode = "SGP"
	SouthGeorgiaAndTheSouthSandwichIslands CountryCode = "SGS"
	SaintHelenaAscensionAndTristanDaCunha  CountryCode = "SHN"
	SvalbardAndJanMayen                    CountryCode = "SJM"
	SolomonIslands                         CountryCode = "SLB"
	SierraLeone                            CountryCode = "SLE"
	ElSalvador                             CountryCode = "SLV"
	SanMarino                              CountryCode = "SMR"
	Somalia                                CountryCode = "SOM"
	SaintPierreAndMiquelon                 CountryCode = "SPM"
	Serbia                                 CountryCode = "SRB"
	SouthSudan                             CountryCode = "SSD"
	SaoTomeAndPrincipe                     CountryCode = "STP"
	Suriname                               CountryCode = "SUR"
	Slovakia                               CountryCode = "SVK"
	Slovenia                               CountryCode = "SVN"
	Sweden                                 CountryCode = "SWE"
	Swaziland                              CountryCode = "SWZ"
	SintMaartenNl                          CountryCode = "SXM"
	Seychelles                             CountryCode = "SYC"
	SyrianArabRepublic                     CountryCode = "SYR"
	TurksAndCaicosIslands                  CountryCode = "TCA"
	Chad                                   CountryCode = "TCD"
	Togo                                   CountryCode = "TGO"
	Thailand                               CountryCode = "THA"
	Tajikstan                              CountryCode = "TJK"
	Tokelau                                CountryCode = "TKL"
	Turkmenistan                           CountryCode = "TKM"
	TimorLeste                             CountryCode = "TLS"
	Tonga                                  CountryCode = "TON"
	TrinidadAndTobago                      CountryCode = "TTO"
	Tunisia                                CountryCode = "TUN"
	Turkey                                 CountryCode = "TUR"
	Tuvalu                                 CountryCode = "TUV"
	Taiwan                                 CountryCode = "TWN"
	Tanzania                               CountryCode = "TZA"
	Uganda                                 CountryCode = "UGA"
	Ukraine                                CountryCode = "UKR"
	UnitedStatesMinorOutlyingIslands       CountryCode = "UMI"
	Uruguay                                CountryCode = "URY"
	UnitedStates                           CountryCode = "USA"
	Uzbekistan                             CountryCode = "UZB"
	Vatican                                CountryCode = "VAT"
	SaintVincentAndTheGrenadines           CountryCode = "VCT"
	Venezuela                              CountryCode = "VEN"
	BritishVirginIslands                   CountryCode = "VGB"
	VirginIslandsUS                        CountryCode = "VIR"
	Vietnam                                CountryCode = "VNM"
	Vanuatu                                CountryCode = "VUT"
	WallisAndFutuna                        CountryCode = "WLF"
	Samoa                                  CountryCode = "WSM"
	Yemen                                  CountryCode = "YEM"
	SouthAfrica                            CountryCode = "ZAF"
	Zambia                                 CountryCode = "ZMB"
	Zimbabwe                               CountryCode = "ZWE"
)

/*
IsValidIsoCC returns a boolean denoting whether the given string appears in the list of ISO country code constants.
*/
func IsValidIsoCC(code CountryCode) bool {

	if utf8.RuneCountInString(string(code)) != 3 {
		return false
	}

	if code == Aruba ||
		code == Afghanistan ||
		code == Angola ||
		code == Anguilla ||
		code == AlandIslands ||
		code == Albania ||
		code == Andorra ||
		code == UnitedArabEmirates ||
		code == Argentina ||
		code == Armenia ||
		code == AmericanSamoa ||
		code == Antarctica ||
		code == FrenchSouthernTerritories ||
		code == AntiguaAndBarbuda ||
		code == Australia ||
		code == Austria ||
		code == Azerbaijan ||
		code == Burundi ||
		code == Belgium ||
		code == Benin ||
		code == BonaireSintEustatiusAndSaba ||
		code == BurkinaFaso ||
		code == Bangladesh ||
		code == Bulgaria ||
		code == Bahrain ||
		code == Bahamas ||
		code == BosniaAndHerzegovina ||
		code == SaintBarthelemy ||
		code == Belarus ||
		code == Belize ||
		code == Bermuda ||
		code == Bolivia ||
		code == Brazil ||
		code == Barbados ||
		code == BruneiDarussalam ||
		code == Bhutan ||
		code == BouvetIsland ||
		code == Botswana ||
		code == CentralAfricanRepublic ||
		code == Canada ||
		code == CocosIslands ||
		code == Switzerland ||
		code == Chile ||
		code == China ||
		code == CotedIvoire ||
		code == Cameroon ||
		code == CongoDR ||
		code == Congo ||
		code == CookIslands ||
		code == Colombia ||
		code == Comoros ||
		code == CaboVerde ||
		code == CostaRica ||
		code == Cuba ||
		code == Curacao ||
		code == ChristmasIsland ||
		code == CaymanIslands ||
		code == Cyprus ||
		code == CzechRepublic ||
		code == Germany ||
		code == Djibouti ||
		code == Dominica ||
		code == Denmark ||
		code == DominicanRepublic ||
		code == Algeria ||
		code == Ecuador ||
		code == Egypt ||
		code == Eritrea ||
		code == WesternSahara ||
		code == Spain ||
		code == Estonia ||
		code == Ethiopia ||
		code == Finland ||
		code == Fiji ||
		code == FalklandIslands ||
		code == France ||
		code == FaroeIslands ||
		code == FederatedStatesOfMicronesia ||
		code == Gabon ||
		code == UnitedKingdom ||
		code == Georgia ||
		code == Guernsey ||
		code == Ghana ||
		code == Gibraltar ||
		code == Guinea ||
		code == Guadeloupe ||
		code == Gambia ||
		code == GuineaBissau ||
		code == EquatorialGuinea ||
		code == Greece ||
		code == Grenada ||
		code == Greenland ||
		code == Guatemala ||
		code == FrenchGuiana ||
		code == Guam ||
		code == Guyana ||
		code == HongKong ||
		code == HeardIslandsAndMcDonaldIslands ||
		code == Honduras ||
		code == Croatia ||
		code == Haiti ||
		code == Hungary ||
		code == Indonesia ||
		code == IsleOfMan ||
		code == India ||
		code == BritishIndianOceanTerritory ||
		code == Ireland ||
		code == Iran ||
		code == Iraq ||
		code == Iceland ||
		code == Israel ||
		code == Italy ||
		code == Jamaica ||
		code == Jersey ||
		code == Jordan ||
		code == Japan ||
		code == Kazakhstan ||
		code == Kenya ||
		code == Kyrgyzstan ||
		code == Cambodia ||
		code == Kiribati ||
		code == SaintKittsAndNevis ||
		code == KoreaRepublic ||
		code == Kuwait ||
		code == Laos ||
		code == Lebanon ||
		code == Liberia ||
		code == Libya ||
		code == SaintLucia ||
		code == Liechtenstein ||
		code == SriLanka ||
		code == Lesotho ||
		code == Lithuania ||
		code == Luxembourg ||
		code == Latvia ||
		code == Macao ||
		code == SaintMartinFr ||
		code == Morocco ||
		code == Monaco ||
		code == Moldova ||
		code == Madagascar ||
		code == Maldives ||
		code == Mexico ||
		code == MarshallIslands ||
		code == MacedoniaFYR ||
		code == Mali ||
		code == Malta ||
		code == Myanmar ||
		code == Montenegro ||
		code == Mongolia ||
		code == NorthernMarianaIslands ||
		code == Mozambique ||
		code == Mauritania ||
		code == Montserrat ||
		code == Martinique ||
		code == Mauritius ||
		code == Malawi ||
		code == Malaysia ||
		code == Mayotte ||
		code == Namibia ||
		code == NewCaledonia ||
		code == Niger ||
		code == NorfolkIsland ||
		code == Nigeria ||
		code == Nicaragua ||
		code == Niue ||
		code == Netherlands ||
		code == Norway ||
		code == Nepal ||
		code == Nauru ||
		code == NewZealand ||
		code == Oman ||
		code == Pakistan ||
		code == Panama ||
		code == Pitcairn ||
		code == Peru ||
		code == Philippines ||
		code == Palau ||
		code == PapuaNewGuinea ||
		code == Poland ||
		code == PuertoRico ||
		code == KoreaDPR ||
		code == Portugal ||
		code == Paraguay ||
		code == Palestine ||
		code == FrenchPolynesia ||
		code == Qatar ||
		code == Reunion ||
		code == Romania ||
		code == RussianFederation ||
		code == Rwanda ||
		code == SaudiArabia ||
		code == Sudan ||
		code == Senegal ||
		code == Singapore ||
		code == SouthGeorgiaAndTheSouthSandwichIslands ||
		code == SaintHelenaAscensionAndTristanDaCunha ||
		code == SvalbardAndJanMayen ||
		code == SolomonIslands ||
		code == SierraLeone ||
		code == ElSalvador ||
		code == SanMarino ||
		code == Somalia ||
		code == SaintPierreAndMiquelon ||
		code == Serbia ||
		code == SouthSudan ||
		code == SaoTomeAndPrincipe ||
		code == Suriname ||
		code == Slovakia ||
		code == Slovenia ||
		code == Sweden ||
		code == Swaziland ||
		code == SintMaartenNl ||
		code == Seychelles ||
		code == SyrianArabRepublic ||
		code == TurksAndCaicosIslands ||
		code == Chad ||
		code == Togo ||
		code == Thailand ||
		code == Tajikstan ||
		code == Tokelau ||
		code == Turkmenistan ||
		code == TimorLeste ||
		code == Tonga ||
		code == TrinidadAndTobago ||
		code == Tunisia ||
		code == Turkey ||
		code == Tuvalu ||
		code == Taiwan ||
		code == Tanzania ||
		code == Uganda ||
		code == Ukraine ||
		code == UnitedStatesMinorOutlyingIslands ||
		code == Uruguay ||
		code == UnitedStates ||
		code == Uzbekistan ||
		code == Vatican ||
		code == SaintVincentAndTheGrenadines ||
		code == Venezuela ||
		code == BritishVirginIslands ||
		code == VirginIslandsUS ||
		code == Vietnam ||
		code == Vanuatu ||
		code == WallisAndFutuna ||
		code == Samoa ||
		code == Yemen ||
		code == SouthAfrica ||
		code == Zambia ||
		code == Zimbabwe {
		return true
	}

	return false

}
