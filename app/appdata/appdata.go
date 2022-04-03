package appdata

import (
	"time"
)

// Global variables
var (
	ChNseTicks chan TickData
	ChStkTick  chan TickData
)

type TickData struct {
	Timestamp       time.Time
	LastTradedPrice float64
	Symbol          string
	LastPrice       float64
	Buy_Demand      uint32
	Sell_Demand     uint32
	TradesTillNow   uint32
	OpenInterest    uint32
}

type Percentage struct {
	Target       float64 // "target": 1,
	SL           float64 // "sl": 1,
	DeepSL       float64 // "deepsl": 1
	Limit_budget float64 // "limit_budget": 50%,
}

type TargetControls struct {
	Trail_target_en         bool      // 	"trail_target_en": true,
	Position_reversal_en    bool      // 	"position_reversal_en": true,
	Delayed_stoploss_min    time.Time // 	"delayed_stoploss_min": "00:30:00",
	Stall_detect_period_min time.Time // 	"stall_detect_period_min": "00:30:00"
}

type Kite_Setting struct {
	Products     string
	Varieties    string
	OrderType    string
	Validities   string
	PositionType string
}

type Trade_setting struct {
	OrderRoute            string
	OptionLevel           int
	OptionExpiryWeek      int
	FuturesExpiryMonth    int
	SkipExipryWeekFutures bool
	LimitAmount           float64
}

type ControlParams struct {
	Percentages     Percentage
	Target_Controls TargetControls
	KiteSettings    Kite_Setting
	TradeSettings   Trade_setting
}

type Strategies struct {
	Strategy     string    // 0
	Enabled      bool      // 1
	Engine       string    // 2
	Trigger_time time.Time // 3
	Trigger_days string    // 4
	Cdl_size     int       // 6
	Instruments  string    // 7
	Controls     string
	CtrlParam    ControlParams
}

type TradeSignal struct {
	Id             uint16    // 1
	Date           time.Time // 2
	Instr          string    // 3
	Strategy       string    // 4
	Dir            string    // 5
	Entry          float64   // 6
	Entry_time     time.Time // 7
	Target         float64   // 8
	Stoploss       float64   // 9
	Trade_id       uint64    // 10
	Exit_val       float64   // 11
	Exit_time      time.Time // 12
	Exit_reason    string    // 13
	Swing_min      float64   // 14
	Swing_max      float64   // 15
	Swing_min_time time.Time // 16
	Swing_max_time time.Time // 17
}

// Env variables required
var UserSettings = []string{
	"APP_LIVE_TRADING_MODE",
	"ZERODHA_USER_ID",
	"ZERODHA_PASSWORD",
	"ZERODHA_API_KEY",
	"ZERODHA_PIN",
	"ZERODHA_API_SECRET",
	"ZERODHA_TOTP_SECRET_KEY",
	"ZERODHA_REQ_TOKEN_URL",
	"TIMESCALEDB_ADDRESS",
	"TIMESCALEDB_USERNAME",
	"TIMESCALEDB_PASSWORD",
	"TIMESCALEDB_PORT"}