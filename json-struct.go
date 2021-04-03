package main

type TxnResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Result  []struct {
		Blocknumber       string `json:"blockNumber"`
		Timestamp         string `json:"timeStamp"`
		Hash              string `json:"hash"`
		Nonce             string `json:"nonce"`
		Blockhash         string `json:"blockHash"`
		From              string `json:"from"`
		Contractaddress   string `json:"contractAddress"`
		To                string `json:"to"`
		Value             string `json:"value"`
		Tokenname         string `json:"tokenName"`
		Tokensymbol       string `json:"tokenSymbol"`
		Tokendecimal      string `json:"tokenDecimal"`
		Transactionindex  string `json:"transactionIndex"`
		Gas               string `json:"gas"`
		Gasprice          string `json:"gasPrice"`
		Gasused           string `json:"gasUsed"`
		Cumulativegasused string `json:"cumulativeGasUsed"`
		Input             string `json:"input"`
		Confirmations     string `json:"confirmations"`
	} `json:"result"`
}
