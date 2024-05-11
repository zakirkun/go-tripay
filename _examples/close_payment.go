package examples

import (
	"fmt"

	"github.com/zakirkun/go-tripay/client"
	"github.com/zakirkun/go-tripay/utils"
)

func Example_close_payment_req_transaction() {
	c := client.Client{
		MerchantCode: "T14302",
		ApiKey:       "your_api_key_here",
		PrivateKey:   "your_private_key_here",
		Mode:         utils.MODE_DEVELOPMENT,
	}

	signStr := utils.Signature{
		Amount:       50000,
		PrivateKey:   "your_private_key_here",
		MerchantCode: "T14302",
		MerchanReff:  "INV345675",
	}

	c.SetSignature(signStr)

	bodyReq := client.ClosePaymentBodyRequest{
		Method:        utils.CHANNEL_QRIS_SHOPEEPAY,
		MerchantRef:   "INV345675",
		Amount:        50000,
		CustomerName:  "Farda Ayu Nurfatika",
		CustomerEmail: "fardaayu@gmail.com",
		CustomerPhone: "6285111990223",
		ReturnURL:     "https://thisisreturnurl.com/redirect",
		ExpiredTime:   client.SetTripayExpiredTime(24), // 24 Hour
		Signature:     c.GetSignature(),
		OrderItems: []client.OrderItemClosePaymentRequest{
			{
				SKU:        "Produk1",
				Name:       "nama produk 1",
				Price:      50000,
				Quantity:   1,
				ProductURL: "https://producturl.com",
				ImageURL:   "https://imageurl.com",
			},
		},
	}

	response, err := c.ClosePaymentRequestTransaction(bodyReq)
	if err != nil {
		panic(err)
	}

	fmt.Printf("response: %v\n", response)
	/*
		response :
		&{true  {DEV-T14302154834TY1ML INV345675 static QRIS_SHOPEEPAY QRIS Custom by ShopeePay Farda Ayu Nurfatika fardaayu@gmail.com   https://thisisreturnurl.com/redirect 50000 1100 0 1100 48900  <nil> https://tripay.co.id/checkout/DEV-T14302154834TY1ML UNPAID 1715343068 [{Produk1 nama produk 1 50000 1 50000 https://producturl.com https://imageurl.com}] [{Pembayaran via QRIS [Masuk ke aplikasi dompet digital Anda yang telah mendukung QRIS Pindai/Scan QR Code yang tersedia Akan muncul detail transaksi. Pastikan data transaksi sudah sesuai Selesaikan proses pembayaran Anda Transaksi selesai. Simpan bukti pembayaran Anda]} {Pembayaran via QRIS (Mobile) [Download QR Code pada invoice Masuk ke aplikasi dompet digital Anda yang telah mendukung QRIS Upload QR Code yang telah di download tadi Akan muncul detail transaksi. Pastikan data transaksi sudah sesuai Selesaikan proses pembayaran Anda Transaksi selesai. Simpan bukti pembayaran Anda]}] SANDBOX MODE https://tripay.co.id/qr/DEV-T14302154834TY1ML}}
	*/
}

func Example_close_payment_get_detail_transaction() {
	c := client.Client{
		MerchantCode: "T14302",
		ApiKey:       "your_api_key_here",
		PrivateKey:   "your_private_key_here",
		Mode:         utils.MODE_DEVELOPMENT,
	}

	signStr := utils.Signature{
		Amount:       50000,
		PrivateKey:   "your_private_key_here",
		MerchantCode: "T14302",
		MerchanReff:  "INV345675",
	}

	c.SetSignature(signStr)

	referenceId := "DEV-T14302154834TY1ML"
	response, err := c.ClosePaymentTransactionGetDetail(referenceId)
	if err != nil {
		panic(err)
	}
	fmt.Printf("response: %v\n", response)
	/*
		response :
		 &{true Transaction found {DEV-T14302154834TY1ML INV345675 static QRIS_SHOPEEPAY QRIS Custom by ShopeePay Farda Ayu Nurfatika fardaayu@gmail.com   https://thisisreturnurl.com/redirect 50000 1100 0 1100 48900  <nil> https://tripay.co.id/checkout/DEV-T14302154834TY1ML UNPAID 1715343068 [{Produk1 nama produk 1 50000 1 50000 https://producturl.com https://imageurl.com}] [{Pembayaran via QRIS [Masuk ke aplikasi dompet digital Anda yang telah mendukung QRIS Pindai/Scan QR Code yang tersedia Akan muncul detail transaksi. Pastikan data transaksi sudah sesuai Selesaikan proses pembayaran Anda Transaksi selesai. Simpan bukti pembayaran Anda]} {Pembayaran via QRIS (Mobile) [Download QR Code pada invoice Masuk ke aplikasi dompet digital Anda yang telah mendukung QRIS Upload QR Code yang telah di download tadi Akan muncul detail transaksi. Pastikan data transaksi sudah sesuai Selesaikan proses pembayaran Anda Transaksi selesai. Simpan bukti pembayaran Anda]}] SANDBOX MODE https://tripay.co.id/qr/DEV-T14302154834TY1ML}}
	*/
}
