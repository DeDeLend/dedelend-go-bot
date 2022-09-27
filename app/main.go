package main

import (
	"ddl/liq/database"
	"ddl/liq/ddl_contract"
	"encoding/json"
	"io/ioutil"
	"log"
	"strconv"
	"time"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

type SymbolContract struct {
	Symbol   string
	Contract *ddl_contract.Contract
}

type Config struct {
	DdlAddressEth        string
	DdlAddressBtc        string
	PsqlUrl              string
	BlockCreate          uint64
	Web3LocalProviderUrl string
	Web3ProviderUrl      string
	PrivateKey           string
	TelegramToken        string
	ChatID               int64
}

func main() {
	data, err := ioutil.ReadFile("./config.json")
	if err != nil {
		panic(err)
	}
	var conf Config
	err = json.Unmarshal(data, &conf)
	if err != nil {
		panic(err)
	}
	arrContrats := []SymbolContract{
		SymbolContract{
			Symbol:   "eth",
			Contract: ddl_contract.InitContract(conf.DdlAddressEth, conf.Web3ProviderUrl, conf.PrivateKey),
		},
		SymbolContract{
			Symbol:   "btc",
			Contract: ddl_contract.InitContract(conf.DdlAddressBtc, conf.Web3ProviderUrl, conf.PrivateKey),
		},
	}
	db := database.NewDataBase(conf.PsqlUrl)
	bot, err := tgbotapi.NewBotAPI(conf.TelegramToken)
	if err != nil {
		panic(err)
	}
	bot.Debug = true
	for {
		start := new(database.LastUpdate)
		db.Select(start)
		fromBlock := start.BlockNumber
		for _, cont := range arrContrats {
			arrBorrow := cont.Contract.GetBorrowEvent(fromBlock)
			log.Println("arrBorrow", arrBorrow)
			for _, v := range arrBorrow {
				db.InsertTable(&database.Options{ID: v.ID, State: v.State, Symbol: cont.Symbol})
			}
			arrCloseEvents := append(
				cont.Contract.GetUnlockEvent(fromBlock),
				append(
					cont.Contract.GetLiquidateEvent(fromBlock),
					append(
						cont.Contract.GetForcedExerciseEvent(fromBlock),
						cont.Contract.GetExerciseByPriorLiqPriceEvent(fromBlock)...,
					)...,
				)...)
			for _, v := range arrCloseEvents {
				db.UpdateTableByID(&database.Options{ID: v.ID, State: v.State, Symbol: cont.Symbol}, v.ID)
			}
		}
		db.UpdateTable(&database.LastUpdate{BlockNumber: arrContrats[0].Contract.GetCurrentBlock()})
		for _, cont := range arrContrats {
			log.Println("Check ", cont.Symbol)
			msg := tgbotapi.NewMessage(conf.ChatID, "Check "+cont.Symbol)
			bot.Send(msg)
			arrActive := db.SelectActiveOptions(cont.Symbol)
			log.Println(arrActive)
			for _, option := range arrActive {
				msg := tgbotapi.NewMessage(conf.ChatID, "Active option: "+strconv.FormatUint(option.ID, 10))
				bot.Send(msg)
				if cont.Contract.LoanState(option.ID) {
					cont.Contract.Liquidate(option.ID)
					log.Println("Liquidate id: " + strconv.FormatUint(option.ID, 10))
					msg := tgbotapi.NewMessage(conf.ChatID, "Liquidate id: "+strconv.FormatUint(option.ID, 10))
					bot.Send(msg)
				} else if cont.Contract.IsOptionExpired(option.ID) {
					cont.Contract.ForcedExercise(option.ID)
					log.Println("ForcedExercise id: " + strconv.FormatUint(option.ID, 10))
					msg := tgbotapi.NewMessage(conf.ChatID, "ForcedExercise id: "+strconv.FormatUint(option.ID, 10))
					bot.Send(msg)
				} else if cont.Contract.LoanStateByPriorLiqPrice(option.ID) {
					cont.Contract.ExerciseByPriorLiqPrice(option.ID)
					log.Println("ExerciseByPriorLiqPrice id: " + strconv.FormatUint(option.ID, 10))
					msg := tgbotapi.NewMessage(conf.ChatID, "ExerciseByPriorLiqPrice id: "+strconv.FormatUint(option.ID, 10))
					bot.Send(msg)
				}
			}
		}
		msg := tgbotapi.NewMessage(conf.ChatID, "wait 60 seconds...")
		bot.Send(msg)
		log.Println("wait 60 seconds...")
		time.Sleep(60 * time.Second)
	}

}
