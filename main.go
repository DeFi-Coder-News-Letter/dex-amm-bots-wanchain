package main

import (
	"fmt"
	"os"

	"github.com/shopspring/decimal"
	"github.com/wanchain/dex-amm-bots/algorithm"
	"github.com/wanchain/dex-amm-bots/client"
	"strconv"
)

func main() {

	fmt.Println("start")

	botType := os.Getenv("BOT_TYPE")

	switch botType {
	case "CONST_PRODUCT":
		startConstProductBot()
	}

}

/*
Env checklist:
 - BOT_PRIVATE_KEY
 - BOT_QUOTE_TOKEN
 - BOT_BASE_URL
 - BOT_MIN_PRICE
 - BOT_MAX_PRICE
 - BOT_PRICE_GAP
 - BOT_EXPAND_INVENTORY
 - BOT_WEB3_URL
 - BOT_OPERATOR_ID
*/
func startConstProductBot() {
	privateKey := os.Getenv("BOT_PRIVATE_KEY")

	makerClient := client.NewHydroClient(
		privateKey,
		os.Getenv("BOT_BASE_TOKEN"),
		os.Getenv("BOT_QUOTE_TOKEN"),
		os.Getenv("BOT_BASE_URL"),
	)

	// makerClient.CancelAllPendingOrders()
	// return

	minPrice, _ := decimal.NewFromString(os.Getenv("BOT_MIN_PRICE"))
	maxPrice, _ := decimal.NewFromString(os.Getenv("BOT_MAX_PRICE"))
	priceGap, _ := decimal.NewFromString(os.Getenv("BOT_PRICE_GAP"))
	expandInventory, _ := decimal.NewFromString(os.Getenv("BOT_EXPAND_INVENTORY"))
	web3Url := os.Getenv("BOT_WEB3_URL")
	operatorID, _ := strconv.Atoi(os.Getenv("BOT_OPERATOR_ID"))

	bot := algorithm.NewConstProductBot(
		makerClient,
		minPrice,
		maxPrice,
		priceGap,
		expandInventory,
		web3Url,
		operatorID,
	)

	bot.Run()
}
