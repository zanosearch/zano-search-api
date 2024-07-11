package zano

import (
	"bytes"
	"fmt"
	"github.com/bytedance/sonic"
	"io"
	"log"
	"net/http"
)

type MarketplaceOffers struct {
	Ap             string `json:"ap"`
	At             string `json:"at"`
	B              string `json:"b"`
	Cat            string `json:"cat"`
	Cnt            string `json:"cnt"`
	Com            string `json:"com"`
	Do             string `json:"do"`
	Et             int    `json:"et"`
	Fee            int64  `json:"fee"`
	IndexInTx      int    `json:"index_in_tx"`
	Lci            string `json:"lci"`
	Lco            string `json:"lco"`
	Ot             int    `json:"ot"`
	P              string `json:"p"`
	Pt             string `json:"pt"`
	Security       string `json:"security"`
	T              string `json:"t"`
	Timestamp      int    `json:"timestamp"`
	TxHash         string `json:"tx_hash"`
	TxOriginalHash string `json:"tx_original_hash"`
	Url            string `json:"url"`
}

type MarketplaceResponse struct {
	Id      int    `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		Offers      []MarketplaceOffers `json:"offers"`
		Status      string              `json:"status"`
		TotalOffers int                 `json:"total_offers"`
	} `json:"result"`
}

func GetOffers(daemonUrl string, limit int) (MarketplaceResponse, error) {
	jsonBody := fmt.Sprintf(`{
	  "id": 0,
	  "jsonrpc": "2.0",
	  "method": "marketplace_global_get_offers_ex",
	  "params": {
		"filter": {
		  "amount_low_limit": 0,
		  "amount_up_limit": 0,
		  "bonus": false,
		  "category": "",
		  "keyword": "",
		  "limit": %d,
		  "location_city": "",
		  "location_country": "",
		  "offer_type_mask": 0,
		  "offset": 0,
		  "order_by": 0,
		  "primary": "",
		  "rate_low_limit": "",
		  "rate_up_limit": "",
		  "reverse": false,
		  "target": "",
		  "timestamp_start": 0,
		  "timestamp_stop": 0
		}
	  }
	}`, limit)

	request, err := http.NewRequest("POST", daemonUrl, bytes.NewBuffer([]byte(jsonBody)))
	if err != nil {
		return MarketplaceResponse{}, err
	}
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}
	res, err := client.Do(request)
	if err != nil {
		return MarketplaceResponse{}, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println(err)
		}
	}(res.Body)

	body, _ := io.ReadAll(res.Body)
	data := MarketplaceResponse{}

	_ = sonic.Unmarshal(body, &data)

	return data, nil
}
