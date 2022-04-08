package objects

// Profile ...
type Profile struct {
	General    ProfileGeneral              `json:"General"`
	Components map[string]ProfileComponent `json:"Components"`
}

// ProfileGeneral ...
type ProfileGeneral struct {
	Code           string `json:"Code"`
	Type           string `json:"Type"`
	Name           string `json:"Name"`
	Exchange       string `json:"Exchange"`
	CurrencyCode   string `json:"CurrencyCode"`
	CurrencyName   string `json:"CurrencyName"`
	CurrencySymbol string `json:"CurrencySymbol"`
	CountryName    string `json:"CountryName"`
	CountryISO     string `json:"CountryISO"`
}

// ProfileComponent ...
type ProfileComponent struct {
	Code     string `json:"Code"`
	Exchange string `json:"Exchange"`
	Name     string `json:"Name"`
	Sector   string `json:"Sector"`
	Industry string `json:"Industry"`
}
