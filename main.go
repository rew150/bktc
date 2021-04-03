package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"math/big"
	"net/http"
	"strconv"
)

var addr string

var accountValue map[string]big.Int

const formatURL = "http://api-ropsten.etherscan.io/api?module=account&action=tokentx&address=%s&startblock=0&endblock=999999999&sort=asc&apikey=K7ST5DC6VP2Z5ZVWWD1IB3JDB5AHIEV274"

func init() {
	flag.StringVar(&addr, "a", "0", "address of the account you want to trace")
	flag.Parse()
	accountValue = make(map[string]big.Int)
}

func main() {
	out := recur(addr, 0)
	for _, v := range out {
		fmt.Println(v.Hash, "\t", v.From, "\t", v.To, "\t", v.Value)
	}
	fmt.Println(len(out))
	for k, v := range accountValue {
		fmt.Println(k, "\t", v.String())
	}
}

func recur(addr string, startTimestamp uint64) []OutEntry {
	var res TxnResponse
	var out []OutEntry
	err := json.Unmarshal(getTxn(addr), &res)
	for err != nil {
		err = json.Unmarshal(getTxn(addr), &res)
	}
	for _, v := range res.Result {
		timestamp, _ := strconv.ParseUint(v.Timestamp, 10, 64)
		if v.Tokensymbol == "BKTC" && timestamp > startTimestamp {
			out = append(out, OutEntry{
				From:  addr,
				To:    v.To,
				Value: v.Value,
				Hash:  v.Hash,
			})
			recurred := recur(addr, timestamp)
			out = append(out, recurred...)
		}
		if v.Tokensymbol == "BKTC" {
			value, _ := new(big.Int).SetString(v.Value, 10)
			from := accountValue[addr]
			from.Sub(&from, value)
			accountValue[addr] = from
			to := accountValue[v.To]
			to.Add(&to, value)
			accountValue[v.To] = to
		}
	}
	return out
}

func getTxn(addr string) []byte {
	resp, err := http.Get(fmt.Sprintf(formatURL, addr))
	for err != nil {
		resp.Body.Close()
		resp, err = http.Get(fmt.Sprintf(formatURL, addr))
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("parse error")
	}
	return body
}
