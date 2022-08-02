package crypto

//import (
//	"encoding/json"
//	"fmt"
//	"io"
//	"net/http"
//	"os"
//
//	"everest/utils"
//
//	log "github.com/sirupsen/logrus"
//)
//
//type Balance struct {
//	Status  string
//	Message string
//	Result  string
//}
//type Transactions struct {
//	Message string
//	Result  []TxnRes
//	Status  string
//}
//
//type TxnRes struct {
//	BlockHash         string `json:"blockHash"`
//	BlockNumber       string `json:"blockNumber"`
//	Confirmations     string `json:"confirmations"`
//	ContractAddress   string `json:"ContractAddress"`
//	CumulativeGasUsed string `json:"CumulativeGasUsed"`
//	From              string `json:"from"`
//	Gas               string `json:"gas"`
//	GasPrice          string `json:"gasPrice"`
//	GasUsed           string `json:"gasUsed"`
//	Hash              string `json:"hash"`
//	Input             string `json:"input"`
//	IsError           string `json:"isError"`
//	Nonce             string `json:"nonce"`
//	TimeStamp         string `json:"timeStamp"`
//	To                string `json:"to"`
//	TransactionIndex  string `json:"transactionIndex"`
//	Txreceipt_status  string `json:"txreceipt_status"`
//	Value             string `json:"value"`
//}
//
//type Coin struct {
//	balance      interface{}
//	transactions []TxnRes
//}
//
//type Btc struct {
//	Address             string `json:"address"`
//	Total_received      int    `json:"total_received"`
//	Total_sent          int    `json:"total_sent"`
//	Balance             int    `json:"balance"`
//	Unconfirmed_balance int    `json:"unconfirmed_balance"`
//	Final_balance       int    `json:"final_balance"`
//	N_tx                int    `json:"n_tx"`
//	Unconfirmed_n_tx    int    `json:"unconfirmed_n_tx"`
//	Final_n_tx          int    `json:"final_n_tx"`
//}
//
//func GetBalance(address, coin string) (interface{}, error) {
//	BSC_API_KEY := os.Getenv("BSC_API_KEY")
//	ETH_API_KEY := os.Getenv("ETH_API_KEY")
//	var url string
//	if coin == "ETH" {
//		url = fmt.Sprintf("https://api.etherscan.io/api?module=account&action=balance&address=%s&apikey=%s", address, ETH_API_KEY)
//	}
//	if coin == "BNB" {
//		url = fmt.Sprintf("https://api.bscscan.com/api?module=account&action=balance&address=%s&apikey=%s", address, BSC_API_KEY)
//		log.Info(url)
//
//	}
//	log.Info(url)
//
//	resp, err := http.Get(url)
//
//	if err != nil {
//		fmt.Println(err)
//		return nil, err
//	}
//	defer resp.Body.Close()
//	body, _ := io.ReadAll(resp.Body)
//
//	var result Balance
//	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to the go struct pointer
//		fmt.Println("Can not unmarshal JSON")
//	}
//	fmt.Println(utils.PrettyPrint(result))
//	return result.Result, nil
//}
//
//func GetTxn(address, coin string) ([]TxnRes, error) {
//	BSC_API_KEY := os.Getenv("BSC_API_KEY")
//	ETH_API_KEY := os.Getenv("ETH_API_KEY")
//
//	var url string
//	if coin == "ETH" {
//		url = fmt.Sprintf("https://api.etherscan.io/api?module=account&action=txlist&address=%s&startblock=0&endblock=99999999&page=1&offset=30&sort=asc&apikey=%s", address, ETH_API_KEY)
//
//	}
//	if coin == "BNB" {
//		url = fmt.Sprintf("https://api.bscscan.com/api?module=account&action=txlist&address=%s&startblock=0&endblock=99999999&page=1&offset=30&sort=asc&apikey=%s", address, BSC_API_KEY)
//
//	}
//	log.Info(url)
//
//	resp, err := http.Get(url)
//
//	if err != nil {
//		fmt.Println(err)
//		return nil, err
//	}
//	defer resp.Body.Close()
//	body, _ := io.ReadAll(resp.Body)
//
//	var result Transactions
//	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to the go struct pointer
//		fmt.Println("Can not unmarshal JSON")
//	}
//	fmt.Println(utils.PrettyPrint(result))
//	return result.Result, nil
//}
//
//func FetchBNB(address string) *Coin {
//
//	bal, _ := GetBalance(address, "BNB")
//	txn, _ := GetTxn(address, "BNB")
//
//	res := &Coin{
//		balance:      bal,
//		transactions: txn,
//	}
//	return res
//}
//
//func FetchETH(address string) *Coin {
//	bal, _ := GetBalance(address, "ETH")
//	txn, _ := GetTxn(address, "ETH")
//
//	res := &Coin{
//		balance:      bal,
//		transactions: txn,
//	}
//	return res
//}
//
//func FetchBTC(btcaddress string) interface{} {
//	resp, err := http.Get(fmt.Sprintf("https://api.blockcypher.com/v1/btc/main/addrs/%s/balance", btcaddress))
//
//	if err != nil {
//		fmt.Println(err)
//	}
//	defer resp.Body.Close()
//	body, _ := io.ReadAll(resp.Body)
//
//	var result Btc
//	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to the go struct pointer
//		fmt.Println("Can not unmarshal JSON")
//	}
//	fmt.Println(utils.PrettyPrint(result))
//	return result
//
//}
