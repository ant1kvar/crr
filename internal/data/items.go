// Package data contains static data for drums
package data

// Country represents a country with name and code
type Country struct {
	Name string
	Code string
}

// Countries is the list of countries for selection
var Countries = []Country{
	{"The United States Of America", "US"},
	{"Germany", "DE"},
	{"The Russian Federation", "RU"},
	{"France", "FR"},
	{"Australia", "AU"},
	{"China", "CN"},
	{"The United Kingdom Of Great Britain And Northern Ireland", "GB"},
	{"Greece", "GR"},
	{"Mexico", "MX"},
	{"Italy", "IT"},
	{"Canada", "CA"},
	{"Brazil", "BR"},
	{"Spain", "ES"},
	{"The Netherlands", "NL"},
	{"India", "IN"},
	{"Poland", "PL"},
	{"The United Arab Emirates", "AE"},
	{"Argentina", "AR"},
	{"The Philippines", "PH"},
	{"Switzerland", "CH"},
	{"Romania", "RO"},
	{"Colombia", "CO"},
	{"Türkiye", "TR"},
	{"Belgium", "BE"},
	{"Indonesia", "ID"},
	{"Uganda", "UG"},
	{"Chile", "CL"},
	{"Hungary", "HU"},
	{"Serbia", "RS"},
	{"Austria", "AT"},
	{"Ukraine", "UA"},
	{"Portugal", "PT"},
	{"Croatia", "HR"},
	{"Czechia", "CZ"},
	{"Bulgaria", "BG"},
	{"Denmark", "DK"},
	{"Ireland", "IE"},
	{"Sweden", "SE"},
	{"New Zealand", "NZ"},
	{"Peru", "PE"},
	{"South Africa", "ZA"},
	{"Taiwan, Republic Of China", "TW"},
	{"Japan", "JP"},
	{"Slovakia", "SK"},
	{"Ecuador", "EC"},
	{"Uruguay", "UY"},
	{"Bolivarian Republic Of Venezuela", "VE"},
	{"Bosnia And Herzegovina", "BA"},
	{"Finland", "FI"},
	{"Israel", "IL"},
	{"Afghanistan", "AF"},
	{"Slovenia", "SI"},
	{"Saudi Arabia", "SA"},
	{"Estonia", "EE"},
	{"Norway", "NO"},
	{"Latvia", "LV"},
	{"Lithuania", "LT"},
	{"The Dominican Republic", "DO"},
	{"The Republic Of Korea", "KR"},
	{"Belarus", "BY"},
	{"Thailand", "TH"},
	{"Tunisia", "TN"},
	{"Malaysia", "MY"},
	{"Pakistan", "PK"},
	{"Hong Kong", "HK"},
	{"Guatemala", "GT"},
	{"Paraguay", "PY"},
	{"Sri Lanka", "LK"},
	{"Kenya", "KE"},
	{"Bolivia", "BO"},
	{"Lebanon", "LB"},
	{"El Salvador", "SV"},
	{"Singapore", "SG"},
	{"Republic Of North Macedonia", "MK"},
	{"Nigeria", "NG"},
	{"Honduras", "HN"},
	{"Cyprus", "CY"},
	{"Egypt", "EG"},
	{"Morocco", "MA"},
	{"Jamaica", "JM"},
	{"Costa Rica", "CR"},
	{"Montenegro", "ME"},
	{"The Republic Of Moldova", "MD"},
	{"Azerbaijan", "AZ"},
	{"Kazakhstan", "KZ"},
	{"Puerto Rico", "PR"},
	{"Vietnam", "VN"},
	{"Albania", "AL"},
	{"Ethiopia", "ET"},
	{"Haiti", "HT"},
	{"Macao", "MO"},
	{"Senegal", "SN"},
	{"Luxembourg", "LU"},
	{"Islamic Republic Of Iran", "IR"},
	{"Nepal", "NP"},
	{"Panama", "PA"},
	{"Georgia", "GE"},
	{"Iraq", "IQ"},
	{"Iceland", "IS"},
	{"Jordan", "JO"},
}

// CountryNames returns list of country names for UI
func CountryNames() []string {
	names := make([]string, len(Countries))
	for i, c := range Countries {
		names[i] = c.Name
	}
	return names
}

// CountryCodeByName returns country code by name
func CountryCodeByName(name string) string {
	for _, c := range Countries {
		if c.Name == name {
			return c.Code
		}
	}
	return ""
}

// Genre is the list of music genres
var Genre = []string{
	"Electronic",
	"Pop",
	"Rock",
	"Hip-Hop",
	"Dance",
	"House",
	"Techno",
	"Trance",
	"Lo-Fi",
	"Chillout",
	"Ambient",
	"Jazz",
	"Blues",
	"Classical",
	"Indie",
	"Metal",
	"Reggae",
	"R&B",
	"Soul",
	"Funk",
	"Darkwave",
	"Synthwave",
	"Industrial",
	"Post-punk",
	"Drum",
}
