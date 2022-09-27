package main

import (
	"ddl/liq/database"
	"encoding/json"
	"io/ioutil"
)

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
	db := database.NewDataBase(conf.PsqlUrl)

	// opt := new(database.Options)
	// lu := new(database.LastUpdate)
	// db.CreateTable(opt)
	// db.CreateTable(lu)
	db.UpdateTable(&database.LastUpdate{BlockNumber: conf.BlockCreate})
}
