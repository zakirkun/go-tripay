package examples

import (
	"fmt"

	"github.com/zakirkun/go-tripay/client"
	"github.com/zakirkun/go-tripay/utils"
)

func Example_merchant_transactions() {
	c := client.Client{
		MerchantCode: "T14302",
		ApiKey:       "your_api_key_here",
		PrivateKey:   "your_private_key_here",
		Mode:         utils.MODE_DEVELOPMENT,
	}
	response, err := c.MerchantTransactions()
	if err != nil {
		panic(err)
	}

	fmt.Printf("response: %v\n", response)
	/*
		response:
		&{true Success [{DEV-T14302154789W4DPV INV345675  BCAVA BCA Virtual Account John Doe johndoe@gmail.com <nil> <nil> https://thisisreturnurl.com/redirect 50000 5500 0 5500 44500 0 <nil> https://tripay.co.id/checkout/DEV-T14302154789W4DPV [{Produk1 nama produk 1 50000 1 50000}] UNPAID <nil> 1715263201 1715349367 <nil>} {DEV-T14302154790HPMWS INV345675  BCAVA BCA Virtual Account John Doe johndoe@gmail.com <nil> <nil> https://thisisreturnurl.com/redirect 50000 5500 0 5500 44500 0 <nil> https://tripay.co.id/checkout/DEV-T14302154790HPMWS [{Produk1 nama produk 1 50000 1 50000}] UNPAID <nil> 1715263347 1715349514 <nil>} {DEV-T14302154794FZ2ZT INV345675  BCAVA BCA Virtual Account John Doe johndoe@gmail.com <nil> <nil> https://thisisreturnurl.com/redirect 50000 5500 0 5500 44500 0 <nil> https://tripay.co.id/checkout/DEV-T14302154794FZ2ZT [{Produk1 nama produk 1 50000 1 50000}] UNPAID <nil> 1715265393 1715351557 <nil>} {DEV-T14302154796GC64X INV345675  BCAVA BCA Virtual Account John Doe johndoe@gmail.com <nil> <nil> https://thisisreturnurl.com/redirect 50000 5500 0 5500 44500 0 <nil> https://tripay.co.id/checkout/DEV-T14302154796GC64X [{Produk1 nama produk 1 50000 1 50000}] UNPAID <nil> 1715266465 1715352865 <nil>} {DEV-T14302154833B5XYD INV345675  BCAVA BCA Virtual Account John Doe johndoe@gmail.com <nil> <nil> https://thisisreturnurl.com/redirect 50000 5500 0 5500 44500 0 <nil> https://tripay.co.id/checkout/DEV-T14302154833B5XYD [{Produk1 nama produk 1 50000 1 50000}] UNPAID <nil> 1715333695 1715420094 <nil>} {DEV-T14302154834TY1ML INV345675  QRIS_SHOPEEPAY QRIS Custom by ShopeePay Farda Ayu Nurfatika fardaayu@gmail.com <nil> <nil> https://thisisreturnurl.com/redirect 50000 1100 0 1100 48900 0 <nil> https://tripay.co.id/checkout/DEV-T14302154834TY1ML [{Produk1 nama produk 1 50000 1 50000}] UNPAID <nil> 1715339528 1715343068 <nil>}] {asc {1 6} 1 <nil> <nil> 1 50 6}}
	*/
}

func Example_merchant_transactions_with_param() {
	c := client.Client{
		MerchantCode: "T14302",
		ApiKey:       "your_api_key_here",
		PrivateKey:   "your_private_key_here",
		Mode:         utils.MODE_DEVELOPMENT,
	}
	merchanTransactionParam := client.MerchantTransactionsParam{
		Page:    1,
		PerPage: 10,
		Sort:    "asc", // asc or desc
	}

	response, err := c.MerchantTransactions(merchanTransactionParam)
	if err != nil {
		panic(err)
	}

	fmt.Printf("response: %v\n", response)
	/*
		response:
		 &{true Success [{DEV-T14302154789W4DPV INV345675  BCAVA BCA Virtual Account John Doe johndoe@gmail.com <nil> <nil> https://thisisreturnurl.com/redirect 50000 5500 0 5500 44500 0 <nil> https://tripay.co.id/checkout/DEV-T14302154789W4DPV [{Produk1 nama produk 1 50000 1 50000}] UNPAID <nil> 1715263201 1715349367 <nil>} {DEV-T14302154790HPMWS INV345675  BCAVA BCA Virtual Account John Doe johndoe@gmail.com <nil> <nil> https://thisisreturnurl.com/redirect 50000 5500 0 5500 44500 0 <nil> https://tripay.co.id/checkout/DEV-T14302154790HPMWS [{Produk1 nama produk 1 50000 1 50000}] UNPAID <nil> 1715263347 1715349514 <nil>} {DEV-T14302154794FZ2ZT INV345675  BCAVA BCA Virtual Account John Doe johndoe@gmail.com <nil> <nil> https://thisisreturnurl.com/redirect 50000 5500 0 5500 44500 0 <nil> https://tripay.co.id/checkout/DEV-T14302154794FZ2ZT [{Produk1 nama produk 1 50000 1 50000}] UNPAID <nil> 1715265393 1715351557 <nil>} {DEV-T14302154796GC64X INV345675  BCAVA BCA Virtual Account John Doe johndoe@gmail.com <nil> <nil> https://thisisreturnurl.com/redirect 50000 5500 0 5500 44500 0 <nil> https://tripay.co.id/checkout/DEV-T14302154796GC64X [{Produk1 nama produk 1 50000 1 50000}] UNPAID <nil> 1715266465 1715352865 <nil>} {DEV-T14302154833B5XYD INV345675  BCAVA BCA Virtual Account John Doe johndoe@gmail.com <nil> <nil> https://thisisreturnurl.com/redirect 50000 5500 0 5500 44500 0 <nil> https://tripay.co.id/checkout/DEV-T14302154833B5XYD [{Produk1 nama produk 1 50000 1 50000}] UNPAID <nil> 1715333695 1715420094 <nil>} {DEV-T14302154834TY1ML INV345675  QRIS_SHOPEEPAY QRIS Custom by ShopeePay Farda Ayu Nurfatika fardaayu@gmail.com <nil> <nil> https://thisisreturnurl.com/redirect 50000 1100 0 1100 48900 0 <nil> https://tripay.co.id/checkout/DEV-T14302154834TY1ML [{Produk1 nama produk 1 50000 1 50000}] UNPAID <nil> 1715339528 1715343068 <nil>}] {asc {1 6} 1 <nil> <nil> 1 10 6}}
	*/
}
