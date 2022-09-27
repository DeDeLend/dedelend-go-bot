package database

import (
	"context"
	"database/sql"
	"log"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
)

type Options struct {
	ID     uint64 `bun:",allowzero,unique"`
	State  bool
	Symbol string
}

type LastUpdate struct {
	BlockNumber uint64
}

type DataBase struct {
	db  *bun.DB
	ctx context.Context
}

func NewDataBase(dsn string) *DataBase {
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	db := bun.NewDB(sqldb, pgdialect.New())
	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
		bundebug.FromEnv("BUNDEBUG"),
	))
	return &DataBase{
		db,
		context.Background(),
	}
}

func (data *DataBase) CreateTable(event any) {
	res, err := data.db.NewCreateTable().Model(event).Exec(data.ctx)
	log.Println(res)
	if err != nil {
		panic(err)
	}
}

func (data *DataBase) DropTable(event any) {
	res, err := data.db.NewDropTable().Model(event).Exec(data.ctx)
	log.Println(res)
	if err != nil {
		panic(err)
	}
}

func (data *DataBase) InsertTable(event any) {
	res, err := data.db.NewInsert().Model(event).Exec(data.ctx)
	log.Println(res)
	if err != nil {
		panic(err)
	}
}

func (data *DataBase) UpdateTable(event any) {
	res, err := data.db.NewUpdate().
		Model(event).
		Where("true").
		Exec(data.ctx)
	log.Println(res)
	if err != nil {
		panic(err)
	}
}

func (data *DataBase) UpdateTableByID(event any, id uint64) {
	res, err := data.db.NewUpdate().
		Model(event).
		Where("id = ?", id).
		Exec(data.ctx)
	log.Println(res)
	if err != nil {
		panic(err)
	}
}

func (data *DataBase) Select(event any) {
	err := data.db.NewSelect().
		Model(event).
		Scan(data.ctx)
	log.Println(event)
	if err != nil {
		panic(err)
	}
}

func (data *DataBase) SelectActiveOptions(symbol string) []Options {
	res := new([]Options)
	err := data.db.NewSelect().
		Table("options").
		Where("options.State = ?", true).
		Where("options.Symbol = ?", symbol).
		Scan(data.ctx, res)
	if err != nil {
		panic(err)
	}
	return *res
}

func (data *DataBase) SelectLastUpdate(event any) {
	var res uint64
	err := data.db.NewSelect().
		ColumnExpr("last_update.block_number").
		Table("last_update").
		Scan(data.ctx, &res)
	log.Println(event)
	if err != nil {
		panic(err)
	}
}
