package idpay

// Client type
type Client struct {
	APIKEY     string
	host       string
	isTestMode bool
}

type requestStatus struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type payment struct {
	TrackId      int32  `json:"track_id,string"`
	Amount       int64  `json:"amount,string"`
	Date         int64  `json:"date,string"`
	CardNo       string `json:"card_no"`
	HashedCardNo string `json:"hashed_card_no"`
}

type verify struct {
	Date int64 `json:"date"`
}

type CreatePaymentReq struct {
	OrderId  string `json:"order_id"`
	Amount   int64  `json:"amount"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Mail     string `json:"mail"`
	Desc     string `json:"desc"`
	Callback string `json:"callback"`
}

type CreatePaymentRes struct {
	ReqStatus    requestStatus
	ErrorCode    int    `json:"error_code"`
	ErrorMessage string `json:"error_message"`
	Id           string `json:"id"`
	Link         string `json:"link"`
}

type VerifyReq struct {
	Id      string `json:"id"`
	OrderId string `json:"order_id"`
}

type VerifyRes struct {
	ReqStatus    requestStatus
	ErrorCode    int    `json:"error_code"`
	ErrorMessage string `json:"error_message"`
	Status       int    `json:"status"`
	TrackId      int32  `json:"track_id,string"`
	Id           string `json:"id"`
	OrderId      string `json:"order_id"`
	Amount       int64  `json:"amount,string"`
	Date         int64  `json:"date,string"`
	Payment      payment
	Verify       verify
}
