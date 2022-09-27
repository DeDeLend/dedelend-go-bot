package ddl_contract

import (
	"context"
	"ddl/liq/contracts"
	"ddl/liq/database"
	"log"
	"math/big"

	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Contract struct {
	Client     *ethclient.Client
	DDL        *contracts.DDL
	ctx        context.Context
	privateKey string
}

func CheckUniqueOptions(arr []database.Options, value database.Options) bool {
	for _, opt := range arr {
		if opt.ID == value.ID {
			return false
		}
	}
	return true
}

func InitContract(contractAddress string, web3ProviderUrl string, privateKey string) *Contract {
	client, err := ethclient.Dial(web3ProviderUrl)
	if err != nil {
		panic(err)
	}
	ddl, err := contracts.NewDDL(common.HexToAddress(contractAddress), client)
	if err != nil {
		panic(err)
	}
	return &Contract{
		client,
		ddl,
		context.Background(),
		privateKey,
	}
}

func (cont *Contract) GetCurrentBlock() uint64 {
	header, err := cont.Client.HeaderByNumber(cont.ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	return header.Number.Uint64()
}

func (cont *Contract) GetBorrowEvent(fromBlock uint64) []database.Options {
	iterator, err := cont.DDL.FilterBorrow(&bind.FilterOpts{
		Start: fromBlock,
		// End:   &toBlock,
	}, nil, nil)
	if err != nil {
		panic(err)
	}
	var res []database.Options
	for iterator.Next() {
		event := iterator.Event
		option := database.Options{ID: event.OptionID.Uint64(), State: false}
		if CheckUniqueOptions(res, option) {
			log.Println(option)
			res = append(res, database.Options{ID: event.OptionID.Uint64(), State: true})
		}
	}
	return res
}

func (cont *Contract) GetUnlockEvent(fromBlock uint64) []database.Options {
	iterator, err := cont.DDL.FilterUnlock(&bind.FilterOpts{
		Start: fromBlock,
		// End:   &toBlock,
	}, nil, nil)
	if err != nil {
		panic(err)
	}
	var res []database.Options
	for iterator.Next() {
		event := iterator.Event
		option := database.Options{ID: event.OptionID.Uint64(), State: false}
		if CheckUniqueOptions(res, option) {
			res = append(res, database.Options{ID: event.OptionID.Uint64(), State: false})
		}
	}
	return res
}

func (cont *Contract) GetLiquidateEvent(fromBlock uint64) []database.Options {
	iterator, err := cont.DDL.FilterLiquidate(&bind.FilterOpts{
		Start: fromBlock,
		// End:   &toBlock,
	}, nil, nil)
	if err != nil {
		panic(err)
	}
	var res []database.Options
	for iterator.Next() {
		event := iterator.Event
		res = append(res, database.Options{ID: event.OptionID.Uint64(), State: false})
	}
	return res
}

func (cont *Contract) GetForcedExerciseEvent(fromBlock uint64) []database.Options {
	iterator, err := cont.DDL.FilterForcedExercise(&bind.FilterOpts{
		Start: fromBlock,
		// End:   &toBlock,
	}, nil, nil)
	if err != nil {
		panic(err)
	}
	var res []database.Options
	for iterator.Next() {
		event := iterator.Event
		res = append(res, database.Options{ID: event.OptionID.Uint64(), State: false})
	}
	return res
}

func (cont *Contract) GetExerciseByPriorLiqPriceEvent(fromBlock uint64) []database.Options {
	iterator, err := cont.DDL.FilterExerciseByPriorLiqPrice(&bind.FilterOpts{
		Start: fromBlock,
		// End:   &toBlock,
	}, nil, nil)
	if err != nil {
		panic(err)
	}
	var res []database.Options
	for iterator.Next() {
		event := iterator.Event
		res = append(res, database.Options{ID: event.OptionID.Uint64(), State: false})
	}
	return res
}

func (cont *Contract) GetCollateralInfo(id uint64) struct {
	Owner    common.Address
	Strategy contracts.DDLoptionInfo
} {
	collateral, err := cont.DDL.CollateralInfo(nil, big.NewInt(int64(id)))
	if err != nil {
		panic(err)
	}
	return collateral
}

func (cont *Contract) GetExpirationByID(id uint64) uint64 {
	res := cont.GetCollateralInfo(id)
	return res.Strategy.Expiration.Uint64()
}

func (cont *Contract) IsOptionExpired(id uint64) bool {
	now := time.Now().Unix()
	experation := cont.GetExpirationByID(id)
	return now > int64(experation)
}

func (cont *Contract) LoanState(id uint64) bool {
	state, err := cont.DDL.LoanState(nil, big.NewInt(int64(id)))
	if err != nil {
		panic(err)
	}
	return state
}

func (cont *Contract) LoanStateByPriorLiqPrice(id uint64) bool {
	state, err := cont.DDL.LoanStateByPriorLiqPrice(nil, big.NewInt(int64(id)))
	if err != nil {
		panic(err)
	}
	return state
}

func (cont *Contract) Liquidate(id uint64) {
	privateKey, _ := crypto.HexToECDSA(cont.privateKey)
	tropt, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(int64(42161)))
	if err != nil {
		panic(err)
	}
	tropt.GasPrice = big.NewInt(2e9)
	log.Println(cont.DDL.Liquidate(tropt, big.NewInt(int64(id))))
}

func (cont *Contract) ForcedExercise(id uint64) {
	privateKey, _ := crypto.HexToECDSA(cont.privateKey)
	tropt, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(int64(42161)))
	if err != nil {
		panic(err)
	}
	tropt.GasPrice = big.NewInt(2e9)
	log.Println(cont.DDL.ForcedExercise(tropt, big.NewInt(int64(id))))
}

func (cont *Contract) ExerciseByPriorLiqPrice(id uint64) {
	privateKey, _ := crypto.HexToECDSA(cont.privateKey)
	tropt, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(int64(42161)))
	if err != nil {
		panic(err)
	}
	tropt.GasPrice = big.NewInt(2e9)
	log.Println(cont.DDL.ExerciseByPriorLiqPrice(tropt, big.NewInt(int64(id))))
}
