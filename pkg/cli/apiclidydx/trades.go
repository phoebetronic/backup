package apiclidydx

import "time"

//
//    {
//        "trades": [{}, {}, {}]
//    }
//
type Response struct {
	Trades []Trade `json:"trades"`
}

//
//    {
//        "side": "SELL",
//        "size": "1.8",
//        "price": "1071.3",
//        "createdAt": "2022-06-30T23:59:24.375Z",
//        "liquidation": false
//    }
//
type Trade struct {
	Side        string    `json:"side"`
	Size        string    `json:"size"`
	Price       string    `json:"price"`
	CreatedAt   time.Time `json:"createdAt"`
	Liquidation bool      `json:"liquidation"`
}
