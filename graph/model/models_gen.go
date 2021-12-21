// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type OrderCreate struct {
	ID                string  `json:"id"`
	Strategy          string  `json:"strategy"`
	Symbol            string  `json:"symbol"`
	TradeType         string  `json:"trade_type"`
	Csize             int     `json:"Csize"`
	Signal            string  `json:"Signal"`
	Entry             float64 `json:"Entry"`
	Target            float64 `json:"Target"`
	Sl                float64 `json:"SL"`
	StopLoss          float64 `json:"StopLoss"`
	DeepStopLoss      float64 `json:"DeepStopLoss"`
	DelayedStopLoss   int     `json:"DelayedStopLoss"`
	StallDetectPeriod int     `json:"StallDetectPeriod"`
	TrailTarget       bool    `json:"TrailTarget"`
	PositionRevarsal  bool    `json:"PositionRevarsal"`
}

type OrderInfo struct {
	ID                string   `json:"id"`
	Strategy          string   `json:"strategy"`
	Symbol            string   `json:"symbol"`
	TradeType         string   `json:"trade_type"`
	Csize             int      `json:"Csize"`
	Signal            string   `json:"Signal"`
	Entry             float64  `json:"Entry"`
	Target            float64  `json:"Target"`
	Sl                float64  `json:"SL"`
	StopLoss          float64  `json:"StopLoss"`
	DeepStopLoss      float64  `json:"DeepStopLoss"`
	DelayedStopLoss   int      `json:"DelayedStopLoss"`
	StallDetectPeriod int      `json:"StallDetectPeriod"`
	TrailTarget       bool     `json:"TrailTarget"`
	PositionRevarsal  bool     `json:"PositionRevarsal"`
	Orderinfo         string   `json:"orderinfo"`
	Timestamp         string   `json:"timestamp"`
	Status            string   `json:"status"`
	EntryTime         *string  `json:"EntryTime"`
	Exit              *float64 `json:"Exit"`
	ExitTime          *string  `json:"ExitTime"`
	Reason            *string  `json:"Reason"`
	Result            *float64 `json:"Result"`
	ResultPerc        *float64 `json:"ResultPerc"`
	SMax              *float64 `json:"SMax"`
	SMaxD             *float64 `json:"SMaxD"`
	SMaxTime          *string  `json:"SMaxTime"`
	SMin              *float64 `json:"SMin"`
	SMinD             *float64 `json:"SMinD"`
	SMinTime          *string  `json:"SMinTime"`
	ExitCriteria      *string  `json:"ExitCriteria"`
}

type ServerStatus struct {
	Kite bool `json:"kite"`
	Db   bool `json:"db"`
	Env  bool `json:"env"`
}