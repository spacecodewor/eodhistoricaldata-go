package objects

// ExchangeDetail ...
type ExchangeDetail struct {
	Name             string      `json:"name"`
	Code             string      `json:"code"`
	Operatingmic     string      `json:"operatingMIC"`
	Country          string      `json:"country"`
	Currency         string      `json:"currency"`
	Timezone         string      `json:"timezone"`
	Exchangeholidays interface{} `json:"exchangeHolidays"`
	IsOpen           bool        `json:"isOpen"`
	Tradinghours     struct {
		Open        string `json:"open"`
		Close       string `json:"close"`
		Openutc     string `json:"openUTC"`
		Closeutc    string `json:"closeUTC"`
		Workingdays string `json:"workingDays"`
	} `json:"tradingHours"`
	Activetickers             int64 `json:"activeTickers"`
	Previousdayupdatedtickers int64 `json:"previousDayUpdatedTickers"`
	Updatedtickers            int64 `json:"updatedTickers"`
}

// Exchange ...
type Exchange struct {
	Name         string `json:"name"`
	Code         string `json:"code"`
	Operatingmic string `json:"operatingMIC"`
	Country      string `json:"country"`
	Currency     string `json:"currency"`
}

// Symbol ...
type Symbol struct {
	Code     string `json:"code"`
	Name     string `json:"name"`
	Country  string `json:"country"`
	Exchange string `json:"exchange"`
	Currency string `json:"currency"`
	Type     string `json:"type"`
}
