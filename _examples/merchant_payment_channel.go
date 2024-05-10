package examples

import (
	"fmt"

	"github.com/zakirkun/go-tripay/client"
	"github.com/zakirkun/go-tripay/utils"
)

func Example_merchant_payment_channel() {
	c := client.Client{
		MerchantCode: "T14302",
		ApiKey:       "your_api_key_here",
		PrivateKey:   "your_private_key_here",
		Mode:         utils.MODE_DEVELOPMENT,
	}

	response, err := c.MerchantPay()
	if err != nil {
		panic(err)
	}
	fmt.Printf("response: %v\n", response)
	/*
		response:
		&{true Success [{Virtual Account MYBVA Maybank Virtual Account direct {4250 0} {0 0} {4250 0} 0 0 https://assets.tripay.co.id/upload/payment-icon/ZT91lrOEad1582929126.png true} {Virtual Account PERMATAVA Permata Virtual Account direct {4250 0} {0 0} {4250 0} 0 0 https://assets.tripay.co.id/upload/payment-icon/szezRhAALB1583408731.png true} {Virtual Account BNIVA BNI Virtual Account direct {4250 0} {0 0} {4250 0} 0 0 https://assets.tripay.co.id/upload/payment-icon/n22Qsh8jMa1583433577.png true} {Virtual Account BRIVA BRI Virtual Account direct {4250 0} {0 0} {4250 0} 0 0 https://assets.tripay.co.id/upload/payment-icon/8WQ3APST5s1579461828.png true} {Virtual Account MANDIRIVA Mandiri Virtual Account direct {4250 0} {0 0} {4250 0} 0 0 https://assets.tripay.co.id/upload/payment-icon/T9Z012UE331583531536.png true} {Virtual Account BCAVA BCA Virtual Account direct {5500 0} {0 0} {5500 0} 0 0 https://assets.tripay.co.id/upload/payment-icon/ytBKvaleGy1605201833.png true} {Virtual Account MUAMALATVA Muamalat Virtual Account direct {4250 0} {0 0} {4250 0} 0 0 https://assets.tripay.co.id/upload/payment-icon/GGwwcgdYaG1611929720.png true} {Virtual Account CIMBVA CIMB Niaga Virtual Account direct {4250 0} {0 0} {4250 0} 0 0 https://assets.tripay.co.id/upload/payment-icon/WtEJwfuphn1614003973.png true} {Virtual Account BSIVA BSI Virtual Account direct {4250 0} {0 0} {4250 0} 0 0 https://assets.tripay.co.id/upload/payment-icon/tEclz5Assb1643375216.png true} {Virtual Account OCBCVA OCBC NISP Virtual Account direct {4250 0} {0 0} {4250 0} 0 0 https://assets.tripay.co.id/upload/payment-icon/ysiSToLvKl1644244798.png true} {Virtual Account DANAMONVA Danamon Virtual Account direct {4250 0} {0 0} {4250 0} 0 0 https://assets.tripay.co.id/upload/payment-icon/F3pGzDOLUz1644245546.png true} {Virtual Account OTHERBANKVA Other Bank Virtual Account direct {4250 0} {0 0} {4250 0} 0 0 https://assets.tripay.co.id/upload/payment-icon/qQYo61sIDa1702995837.png true} {Convenience Store ALFAMART Alfamart direct {3500 0} {0 0} {3500 0} 0 0 https://assets.tripay.co.id/upload/payment-icon/jiGZMKp2RD1583433506.png true} {Convenience Store INDOMARET Indomaret direct {3500 0} {0 0} {3500 0} 0 0 https://assets.tripay.co.id/upload/payment-icon/zNzuO5AuLw1583513974.png true} {Convenience Store ALFAMIDI Alfamidi direct {3500 0} {0 0} {3500 0} 0 0 https://assets.tripay.co.id/upload/payment-icon/aQTdaUC2GO1593660384.png true} {E-Wallet OVO OVO redirect {0 3} {0 0} {0 0} 1000 0 https://assets.tripay.co.id/upload/payment-icon/fH6Y7wDT171586199243.png true} {E-Wallet QRIS QRIS by ShopeePay direct {750 0} {0 0} {750 0} 0 0 https://assets.tripay.co.id/upload/payment-icon/BpE4BPVyIw1605597490.png true} {E-Wallet QRISC QRIS (Customizable) direct {750 0} {0 0} {750 0} 0 0 https://assets.tripay.co.id/upload/payment-icon/m9FtFwaBCg1623157494.png true} {E-Wallet QRIS2 QRIS direct {750 0} {0 0} {750 0} 0 0 https://assets.tripay.co.id/upload/payment-icon/8ewGzP6SWe1649667701.png true} {E-Wallet DANA DANA redirect {0 3} {0 0} {0 0} 1000 0 https://assets.tripay.co.id/upload/payment-icon/sj3UHLu8Tu1655719621.png true} {E-Wallet SHOPEEPAY ShopeePay redirect {0 3} {0 0} {0 0} 1000 0 https://assets.tripay.co.id/upload/payment-icon/d204uajhlS1655719774.png true} {E-Wallet QRIS_SHOPEEPAY QRIS Custom by ShopeePay direct {750 0} {0 0} {750 0} 0 0 https://assets.tripay.co.id/upload/payment-icon/DM8sBd1i9y1681718593.png true}]}
	*/
}
