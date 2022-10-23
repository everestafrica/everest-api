package crypto

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/everestafrica/everest-api/internal/commons/types"
	"github.com/everestafrica/everest-api/internal/config"
)

type Transaction struct {
	WalletAddress string                `json:"wallet_address"`
	Hash          string                `json:"hash"`
	Fees          string                `json:"fees"`
	Value         float64               `json:"value"`
	Date          time.Time             `json:"date"`
	Type          types.TransactionType `json:"type"`
}
type Balance struct {
	WalletAddress string  `json:"wallet_address"`
	Value         float64 `json:"value"`
}

type EthTransaction struct {
	Message string
	Result  []EthTxnRes
	Status  string
}
type EthTxnRes struct {
	BlockHash         string `json:"blockHash"`
	BlockNumber       string `json:"blockNumber"`
	Confirmations     string `json:"confirmations"`
	ContractAddress   string `json:"ContractAddress"`
	CumulativeGasUsed string `json:"CumulativeGasUsed"`
	From              string `json:"from"`
	Gas               string `json:"gas"`
	GasPrice          string `json:"gasPrice"`
	GasUsed           string `json:"gasUsed"`
	Hash              string `json:"hash"`
	Input             string `json:"input"`
	IsError           string `json:"isError"`
	Nonce             string `json:"nonce"`
	TimeStamp         string `json:"timeStamp"`
	To                string `json:"to"`
	TransactionIndex  string `json:"transactionIndex"`
	TxreceiptStatus   string `json:"txreceipt_status"`
	Value             string `json:"value"`
}
type EthBalance struct {
	Status  string
	Message string
	Result  string
}

type BtcTxn struct {
	Data struct {
		Page  int `json:"page"`
		Limit int `json:"limit"`
		Pages int `json:"pages"`
		List  []struct {
			Segwit            bool        `json:"segwit"`
			Rbf               bool        `json:"rbf"`
			TxID              string      `json:"txId"`
			Version           int         `json:"version"`
			Size              int         `json:"size"`
			VSize             int         `json:"vSize"`
			BSize             int         `json:"bSize"`
			LockTime          int         `json:"lockTime"`
			Confirmations     int         `json:"confirmations"`
			BlockTime         int         `json:"blockTime"`
			BlockIndex        int         `json:"blockIndex"`
			Coinbase          bool        `json:"coinbase"`
			Fee               int         `json:"fee"`
			Data              interface{} `json:"data"`
			Amount            int         `json:"amount"`
			Weight            int         `json:"weight"`
			BlockHeight       int         `json:"blockHeight"`
			Timestamp         int         `json:"timestamp"`
			InputsAmount      int         `json:"inputsAmount"`
			InputAddressCount int         `json:"inputAddressCount"`
			OutAddressCount   int         `json:"outAddressCount"`
			InputsCount       int         `json:"inputsCount"`
			OutsCount         int         `json:"outsCount"`
			OutputsAmount     int         `json:"outputsAmount"`
			AddressReceived   int         `json:"addressReceived"`
			AddressOuts       int         `json:"addressOuts"`
			AddressSent       int         `json:"addressSent"`
			AddressInputs     int         `json:"addressInputs"`
		} `json:"list"`
	} `json:"data"`
	Time float64 `json:"time"`
}
type BtcBal struct {
	Address            string `json:"address"`
	TotalReceived      int    `json:"total_received"`
	TotalSent          int    `json:"total_sent"`
	Balance            int    `json:"balance"`
	UnconfirmedBalance int    `json:"unconfirmed_balance"`
	FinalBalance       int    `json:"final_balance"`
	NTx                int    `json:"n_tx"`
	UnconfirmedNTx     int    `json:"unconfirmed_n_tx"`
	FinalNTx           int    `json:"final_n_tx"`
}

type SolTransaction struct {
	BlockTime         int      `json:"blockTime"`
	Slot              int      `json:"slot"`
	TxHash            string   `json:"txHash"`
	Fee               int      `json:"fee"`
	Status            string   `json:"status"`
	Lamport           int      `json:"lamport"`
	Signer            []string `json:"signer"`
	ParsedInstruction []struct {
		ProgramID string `json:"programId"`
		Program   string `json:"program"`
		Type      string `json:"type"`
	} `json:"parsedInstruction"`
	IncludeSPLTransfer bool `json:"includeSPLTransfer,omitempty"`
}
type SolBal struct {
	Lamports     int    `json:"lamports"`
	OwnerProgram string `json:"ownerProgram"`
	Type         string `json:"type"`
	RentEpoch    int    `json:"rentEpoch"`
	Executable   bool   `json:"executable"`
	Account      string `json:"account"`
}

const SolPerLamport = 0.000000001
const BtcPerSat = 0.00000001
const EthPerWei = 0.000000000000000001

func GetBalance(address string, symbol types.CryptoSymbol) (*Balance, error) {

	BscApiKey := config.GetConf().BscApiKey
	EthApiKey := config.GetConf().EthApiKey

	var url string
	var result Balance

	switch symbol {
	case types.ETH:
		url = fmt.Sprintf("https://api.etherscan.io/api?module=account&action=balance&address=%s&apikey=%s", address, EthApiKey)
		var response EthBalance

		v, err := Get(url, response)
		if err != nil {
			return nil, err
		}
		fmt.Println("v: ", v)
		res := v.(EthBalance)
		val, _ := strconv.Atoi(res.Result)
		bal := Balance{
			WalletAddress: address,
			Value:         float64(val) * EthPerWei,
		}
		result = bal
	case types.BSC:
		url = fmt.Sprintf("https://api.bscscan.com/api?module=account&action=balance&address=%s&apikey=%s", address, BscApiKey)
		var response EthBalance

		v, err := Get(url, response)
		if err != nil {
			return nil, err
		}
		res := v.(EthBalance)
		val, _ := strconv.Atoi(res.Result)
		bal := Balance{
			WalletAddress: address,
			Value:         float64(val) * EthPerWei,
		}
		result = bal
	case types.SOL:
		url = fmt.Sprintf("https://public-api.solscan.io/account/%s", address)
		var response SolBal

		v, err := Get(url, response)
		if err != nil {
			return nil, err
		}
		res := v.(SolBal)
		bal := Balance{
			WalletAddress: address,
			Value:         float64(res.Lamports) * SolPerLamport,
		}
		result = bal
	case types.BTC:
		url = fmt.Sprintf("https://api.bitaps.com/btc/v1/blockchain/address/state/%s", address)
		var response BtcBal

		v, err := Get(url, response)
		if err != nil {
			return nil, err
		}
		res := v.(BtcBal)
		bal := Balance{
			WalletAddress: address,
			Value:         float64(res.FinalBalance) * BtcPerSat,
		}
		result = bal
	}

	return &result, nil
}

func GetTransaction(address string, symbol types.CryptoSymbol) (*[]Transaction, error) {
	BscApiKey := config.GetConf().BscApiKey
	EthApiKey := config.GetConf().EthApiKey

	var url string
	var result []Transaction

	switch symbol {
	case types.ETH:
		url = fmt.Sprintf("https://api.etherscan.io/api?module=account&action=txlist&address=%s&startblock=0&endblock=99999999&page=1&offset=30&sort=asc&apikey=%s", address, EthApiKey)
		var response EthTransaction

		v, err := Get(url, response)
		if err != nil {
			return nil, err
		}
		res := v.(EthTransaction)
		for _, data := range res.Result {
			txnType := types.Debit
			if data.To == address {
				txnType = types.Credit
			}
			t, _ := strconv.Atoi(data.TimeStamp)
			v, _ := strconv.Atoi(data.Value)
			txnTime := time.Unix(int64(t), 0)
			txn := Transaction{
				WalletAddress: address,
				Hash:          data.Hash,
				Fees:          data.GasPrice,
				Value:         float64(v),
				Date:          txnTime,
				Type:          txnType,
			}
			result = append(result, txn)
		}
	case types.BSC:
		url = fmt.Sprintf("https://api.bscscan.com/api?module=account&action=txlist&address=%s&startblock=0&endblock=99999999&page=1&offset=30&sort=asc&apikey=%s", address, BscApiKey)
		var response EthTransaction

		v, err := Get(url, response)
		if err != nil {
			return nil, err
		}
		transaction := v.(EthTransaction)
		for _, data := range transaction.Result {
			txnType := types.Debit
			if data.To == address {
				txnType = types.Credit
			}
			t, _ := strconv.Atoi(data.TimeStamp)
			v, _ := strconv.Atoi(data.Value)
			txnTime := time.Unix(int64(t), 0)
			txn := Transaction{
				WalletAddress: address,
				Hash:          data.Hash,
				Fees:          data.GasPrice,
				Value:         float64(v),
				Date:          txnTime,
				Type:          txnType,
			}
			result = append(result, txn)
		}
	case types.SOL:
		url = fmt.Sprintf("https://public-api.solscan.io/account/transactions?account=%s", address)
		var response []SolTransaction

		v, err := Get(url, response)
		if err != nil {
			return nil, err
		}
		transactions := v.([]SolTransaction)
		for _, data := range transactions {
			txnType := types.Debit
			if data.Signer[0] == address {
				txnType = types.Credit
			}
			txn := Transaction{
				WalletAddress: address,
				Hash:          data.TxHash,
				Fees:          string(rune(data.Fee)),
				Value:         float64(data.Lamport) * SolPerLamport,
				Date:          time.Unix(int64(data.BlockTime), 0),
				Type:          txnType,
			}
			result = append(result, txn)
		}
	case types.BTC:
		url = fmt.Sprintf("https://api.bitaps.com/btc/v1/blockchain/address/transactions/%s", address)
		var response BtcTxn

		v, err := Get(url, response)
		if err != nil {
			return nil, err
		}
		res := v.(BtcTxn)
		for _, data := range res.Data.List {
			txnType := types.Debit
			if data.Amount > 0 {
				txnType = types.Credit
			}
			txn := Transaction{
				WalletAddress: address,
				Hash:          data.TxID,
				Fees:          string(rune(data.Fee)),
				Value:         float64(data.Amount) * BtcPerSat,
				Date:          time.Unix(int64(data.BlockTime), 0),
				Type:          txnType,
			}
			result = append(result, txn)
		}
	}
	return &result, nil
}

func Get(url string, response interface{}) (interface{}, error) {
	var result interface{}
	resp, err := http.Get(url)

	if err != nil {
		//logger := log.WithField("error in Mono GET request", err)
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			return
		}
	}(resp.Body)
	if resp.StatusCode >= http.StatusBadRequest {
		return nil, errors.New(resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &result)

	if err != nil {
		return nil, err
	}
	err = mapstructure.Decode(result, &response)

	if err != nil {
		return nil, err
	}

	return response, nil
}
