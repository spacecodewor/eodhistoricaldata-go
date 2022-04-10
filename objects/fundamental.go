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

// EconomicEvent ...
type EconomicEvent struct {
	Type             string   `json:"type"`
	Comparison       *string  `json:"comparison"`
	Period           *string  `json:"period"`
	Country          string   `json:"country"`
	Date             string   `json:"date"`
	Actual           *float64 `json:"actual"`
	Previous         *float64 `json:"previous"`
	Estimate         *float64 `json:"estimate"`
	Change           *float64 `json:"change"`
	ChangePercentage *float64 `json:"change_percentage"`
}

// InsiderTransaction ...
type InsiderTransaction struct {
	Code                        string      `json:"code"`
	Exchange                    string      `json:"exchange"`
	Date                        string      `json:"date"`
	ReportDate                  string      `json:"reportDate"`
	OwnerCik                    interface{} `json:"ownerCik"`
	OwnerName                   string      `json:"ownerName"`
	OwnerRelationship           interface{} `json:"ownerRelationship"`
	OwnerTitle                  string      `json:"ownerTitle"`
	TransactionDate             string      `json:"transactionDate"`
	TransactionCode             string      `json:"transactionCode"`
	TransactionAmount           int         `json:"transactionAmount"`
	TransactionPrice            float64     `json:"transactionPrice"`
	TransactionAcquiredDisposed string      `json:"transactionAcquiredDisposed"`
	PostTransactionAmount       int         `json:"postTransactionAmount"`
	Link                        string      `json:"link"`
}
