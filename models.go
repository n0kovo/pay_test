package main

import (
	"github.com/n0kovo/pay_test/cryptocurrencies/bitcoin"
	"github.com/n0kovo/pay_test/cryptocurrencies/bitcoincash"
	"github.com/n0kovo/pay_test/cryptocurrencies/ethereum"
	"github.com/n0kovo/pay_test/db"
)

func SyncDatabase() {
	db.Database.AutoMigrate(
		&bitcoin.BitcoinWallet{},
		&bitcoin.BitcoinWalletBalance{},

		&bitcoincash.BitcoinCashWallet{},
		&bitcoincash.BitcoinCashWalletBalance{},

		&ethereum.EthereumWallet{},
		&ethereum.EthereumWalletBalance{},
	)

	ethereum.SetupViews()
	bitcoin.SetupViews()
	bitcoincash.SetupViews()
}
