package ethereum

import (
	"github.com/n0kovo/pay_test/settings"
)

var (
	PAYMENT_GATE_SETTINGS = settings.LoadPaymentGateSettings()
	WEI_IN_ETH            = float64(1e18)
)
