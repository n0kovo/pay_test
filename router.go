package main

import (
	"github.com/gocraft/web"

	"github.com/n0kovo/pay_test/cryptocurrencies"
	"github.com/n0kovo/pay_test/cryptocurrencies/bitcoin"
	"github.com/n0kovo/pay_test/cryptocurrencies/bitcoincash"
	"github.com/n0kovo/pay_test/cryptocurrencies/ethereum"

	"github.com/n0kovo/pay_test/exchange"
)

func ConfigureRouter(router *web.Router) *web.Router {

	bitcoinRouter := router.Subrouter(Context{}, "/bitcoin")
	bitcoin.ConfigureRouter(bitcoinRouter)

	bitcoinCashRouter := router.Subrouter(Context{}, "/bitcoin_cash")
	bitcoincash.ConfigureRouter(bitcoinCashRouter)

	ethereumRouter := router.Subrouter(Context{}, "/ethereum")
	ethereum.ConfigureRouter(ethereumRouter)

	// Exchange
	exchangeRouter := router.Subrouter(Context{}, "/exchange")
	exchange.ConfigureRouter(exchangeRouter)

	// Currency
	router.Get("/currency/:base_currency", cryptocurrencies.ViewShowCurrency)
	return router
}
