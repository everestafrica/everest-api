package crypto

import (
	"encoding/json"
	"fmt"
	"github.com/everestafrica/everest-api/internal/commons/log"
	"github.com/everestafrica/everest-api/internal/commons/utils"
	"github.com/everestafrica/everest-api/internal/config"
	"io"
	"net/http"
)

type Balance struct {
	Status  string
	Message string
	Result  string
}
type Transactions struct {
	Message string
	Result  []TxnRes
	Status  string
}

type TxnRes struct {
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

type Coin struct {
	balance      interface{}
	transactions []TxnRes
}

type Btc struct {
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

func GetBalance(address, coin string) (interface{}, error) {
	BscApiKey := config.GetConf().BscApiKey
	EthApiKey := config.GetConf().EthApiKey
	var url string
	if coin == "ETH" {
		url = fmt.Sprintf("https://api.etherscan.io/api?module=account&action=balance&address=%s&apikey=%s", address, EthApiKey)
	}
	if coin == "BNB" {
		url = fmt.Sprintf("https://api.bscscan.com/api?module=account&action=balance&address=%s&apikey=%s", address, BscApiKey)
		log.Info(url)

	}
	log.Info(url)

	resp, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	var result Balance
	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to the go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}
	fmt.Println(util.PrettyPrint(result))
	return result.Result, nil
}

func GetTxn(address, coin string) ([]TxnRes, error) {
	BscApiKey := config.GetConf().BscApiKey
	EthApiKey := config.GetConf().EthApiKey

	var url string
	if coin == "ETH" {
		url = fmt.Sprintf("https://api.etherscan.io/api?module=account&action=txlist&address=%s&startblock=0&endblock=99999999&page=1&offset=30&sort=asc&apikey=%s", address, EthApiKey)

	}
	if coin == "BNB" {
		url = fmt.Sprintf("https://api.bscscan.com/api?module=account&action=txlist&address=%s&startblock=0&endblock=99999999&page=1&offset=30&sort=asc&apikey=%s", address, BscApiKey)

	}
	log.Info(url)

	resp, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	var result Transactions
	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to the go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}
	fmt.Println(util.PrettyPrint(result))
	return result.Result, nil
}

func FetchBNB(address string) *Coin {

	bal, _ := GetBalance(address, "BNB")
	txn, _ := GetTxn(address, "BNB")

	res := &Coin{
		balance:      bal,
		transactions: txn,
	}
	return res
}

func FetchETH(address string) *Coin {
	bal, _ := GetBalance(address, "ETH")
	txn, _ := GetTxn(address, "ETH")

	res := &Coin{
		balance:      bal,
		transactions: txn,
	}
	return res
}

func FetchBTC(address string) (*BtcTransaction, error) {
	resp, err := http.Get(fmt.Sprintf("https://blockchain.info/rawaddr/%s", address))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	var result BtcTransaction
	if err = json.Unmarshal(body, &result); err != nil { // Parse []byte to the go struct pointer
		return nil, err
	}
	fmt.Println(util.PrettyPrint(result))
	return &result, nil
}

type BtcTransaction struct {
	Hash160       string `json:"hash160"`
	Address       string `json:"address"`
	NTx           int    `json:"n_tx"`
	NUnredeemed   int    `json:"n_unredeemed"`
	TotalReceived int    `json:"total_received"`
	TotalSent     int    `json:"total_sent"`
	FinalBalance  int    `json:"final_balance"`
	Txs           []struct {
		Hash        string `json:"hash"`
		Ver         int    `json:"ver"`
		VinSz       int    `json:"vin_sz"`
		VoutSz      int    `json:"vout_sz"`
		Size        int    `json:"size"`
		Weight      int    `json:"weight"`
		Fee         int    `json:"fee"`
		RelayedBy   string `json:"relayed_by"`
		LockTime    int    `json:"lock_time"`
		TxIndex     int64  `json:"tx_index"`
		DoubleSpend bool   `json:"double_spend"`
		Time        int    `json:"time"`
		BlockIndex  int    `json:"block_index"`
		BlockHeight int    `json:"block_height"`
		Inputs      []struct {
			Sequence int64  `json:"sequence"`
			Witness  string `json:"witness"`
			Script   string `json:"script"`
			Index    int    `json:"index"`
			PrevOut  struct {
				TxIndex           int64  `json:"tx_index"`
				Value             int    `json:"value"`
				N                 int    `json:"n"`
				Type              int    `json:"type"`
				Spent             bool   `json:"spent"`
				Script            string `json:"script"`
				SpendingOutpoints []struct {
					TxIndex int64 `json:"tx_index"`
					N       int   `json:"n"`
				} `json:"spending_outpoints"`
				Addr string `json:"addr"`
			} `json:"prev_out"`
		} `json:"inputs"`
		Out []struct {
			Type              int  `json:"type"`
			Spent             bool `json:"spent"`
			Value             int  `json:"value"`
			SpendingOutpoints []struct {
				TxIndex int64 `json:"tx_index"`
				N       int   `json:"n"`
			} `json:"spending_outpoints"`
			N       int    `json:"n"`
			TxIndex int64  `json:"tx_index"`
			Script  string `json:"script"`
			Addr    string `json:"addr"`
		} `json:"out"`
		Result  int `json:"result"`
		Balance int `json:"balance"`
	} `json:"txs"`
}

//fmt.Sprintf("https://api.blockcypher.com/v1/btc/main/addrs/%s/balance", address)
//spent:true --debit/credit
//hash: ""
//value:92300
//time:1657802535

//multiplier = 0.00000001
