package objects

// Quote ...
type Quote struct {
	Code          string      `json:"code"`
	Timestamp     int64       `json:"timestamp"`
	Gmtoffset     int64       `json:"gmtoffset"`
	Open          float64     `json:"open"`
	High          float64     `json:"high"`
	Low           float64     `json:"low"`
	Close         float64     `json:"close"`
	Volume        interface{} `json:"volume"` // int64 || "NA"
	PreviousClose float64     `json:"previousClose"`
	Change        float64     `json:"change"`
	ChangeP       float64     `json:"change_p"`
}

// Candle ...
type Candle struct {
	Date          string  `json:"date"`
	Open          float64 `json:"open"`
	High          float64 `json:"high"`
	Low           float64 `json:"low"`
	Close         float64 `json:"close"`
	AdjustedClose float64 `json:"adjusted_close"`
	Volume        float64 `json:"volume"`
}
