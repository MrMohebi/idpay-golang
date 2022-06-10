package main

import "github.com/MrMohebi/idpay-golang"

func main() {
	idpayC := idpay.NewClient("2bb467dd-aaaa-aaaa-aaaa-aaaaafc90add", true)

	createRes := idpayC.CreatePayment("123", 5000, "https://google.com", "", "", "", "")
	println(createRes.Link)

	verifyRes := idpayC.Verify("3adf8e642413ed0d5d17b9e4901f9065", "123")
	println(verifyRes.TrackId)
}
