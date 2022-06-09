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

type CreatePaymentReq struct {
	OrderId  string `json:"order_id"`
	Amount   int    `json:"amount"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Mail     string `json:"mail"`
	Desc     string `json:"desc"`
	Callback string `json:"callback"`
}

type CreatePaymentRes struct {
	ReqStatus    requestStatus
	Id           string `json:"id"`
	Link         string `json:"link"`
	ErrorCode    int    `json:"error_code"`
	ErrorMessage string `json:"error_message"`
}
