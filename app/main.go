package main

import (
	"ddl/liq/database"
	"ddl/liq/ddl_contract"
	"ddl/liq/recoverer"
	"ddl/liq/telegram"
	"encoding/json"
	"io/ioutil"
	"log"
	"strconv"
	"time"
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
	botTelegram := telegram.NewTelegram(conf.TelegramToken, conf.ChatID)
	rec := recoverer.Recoverer{
		PrintError: func(x string) { botTelegram.SendMessage(x) },
		SleepTime:  10,
	}
	for {
		start := new(database.LastUpdate)
		db.Select(start)
		fromBlock := start.BlockNumber
		for _, cont := range arrContrats {
			var arrBorrow []database.Options
			rec.RecoverPanicFunction(func() { arrBorrow = cont.Contract.GetBorrowEvent(fromBlock) })
			for _, v := range arrBorrow {
				rec.RecoverPanicFunction(func() {
					db.InsertTable(&database.Options{ID: v.ID, State: v.State, Symbol: cont.Symbol})
				})
			}
			var UnlockEvents []database.Options
			var LiquidateEvents []database.Options
			var ForcedExerciseEvents []database.Options
			var ExerciseByPriorLiqPriceEvents []database.Options

			rec.RecoverPanicFunction(func() { UnlockEvents = cont.Contract.GetUnlockEvent(fromBlock) })
			rec.RecoverPanicFunction(func() { LiquidateEvents = cont.Contract.GetLiquidateEvent(fromBlock) })
			rec.RecoverPanicFunction(func() { ForcedExerciseEvents = cont.Contract.GetForcedExerciseEvent(fromBlock) })
			rec.RecoverPanicFunction(func() { ExerciseByPriorLiqPriceEvents = cont.Contract.GetExerciseByPriorLiqPriceEvent(fromBlock) })

			arrCloseEvents := append(
				UnlockEvents,
				append(
					LiquidateEvents,
					append(
						ForcedExerciseEvents,
						ExerciseByPriorLiqPriceEvents...,
					)...,
				)...)
			for _, v := range arrCloseEvents {
				rec.RecoverPanicFunction(func() {
					db.UpdateTableByID(&database.Options{ID: v.ID, State: v.State, Symbol: cont.Symbol}, v.ID)
				})
			}
		}
		rec.RecoverPanicFunction(func() {
			db.UpdateTable(&database.LastUpdate{BlockNumber: arrContrats[0].Contract.GetCurrentBlock()})
		})
		for _, cont := range arrContrats {
			botTelegram.SendMessage("Check " + cont.Symbol)
			var arrActive []database.Options
			rec.RecoverPanicFunction(func() { arrActive = db.SelectActiveOptions(cont.Symbol) })
			log.Println(arrActive)
			for _, option := range arrActive {
				botTelegram.SendMessage("Active option: " + strconv.FormatUint(option.ID, 10))
				var LoanState bool
				var IsOptionExpired bool
				var LoanStateByPriorLiqPrice bool
				rec.RecoverPanicFunction(func() { LoanState = cont.Contract.LoanState(option.ID) })
				rec.RecoverPanicFunction(func() { IsOptionExpired = cont.Contract.IsOptionExpired(option.ID) })
				rec.RecoverPanicFunction(func() { LoanStateByPriorLiqPrice = cont.Contract.LoanStateByPriorLiqPrice(option.ID) })
				if LoanState {
					rec.RecoverPanicFunction(func() { cont.Contract.Liquidate(option.ID) })
					botTelegram.SendMessage("Liquidate id: " + strconv.FormatUint(option.ID, 10))
				} else if IsOptionExpired {
					rec.RecoverPanicFunction(func() { cont.Contract.ForcedExercise(option.ID) })
					botTelegram.SendMessage("ForcedExercise id: " + strconv.FormatUint(option.ID, 10))
				} else if LoanStateByPriorLiqPrice {
					rec.RecoverPanicFunction(func() { cont.Contract.ExerciseByPriorLiqPrice(option.ID) })
					botTelegram.SendMessage("ExerciseByPriorLiqPrice id: " + strconv.FormatUint(option.ID, 10))
				}
			}
		}
		botTelegram.SendMessage("wait 60 seconds...")
		time.Sleep(60 * time.Second)
	}

}
