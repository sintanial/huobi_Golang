package algoorder

type GetSpecificOrderResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    *struct {
		AccountId       int    `json:"accountId"`
		Source          string `json:"source"`
		ClientOrderId   string `json:"clientOrderId"`
		OrderId         int    `json:"orderId"`
		Symbol          string `json:"symbol"`
		OrderPrice      string `json:"orderPrice"`
		OrderSize       string `json:"orderSize"`
		OrderValue      string `json:"orderValue"`
		OrderSide       string `json:"orderSide"`
		TimeInForce     string `json:"timeInForce"`
		OrderType       string `json:"orderType"`
		StopPrice       string `json:"stopPrice"`
		TrailingRate    string `json:"trailingRate"`
		OrderOrigTime   int64  `json:"orderOrigTime"`
		LastActTime     int64  `json:"lastActTime"`
		OrderCreateTime int64  `json:"orderCreateTime"`
		OrderStatus     string `json:"orderStatus"`
		ErrCode         string `json:"errCode"`
		ErrMessage      string `json:"errMessage"`
	}
}
