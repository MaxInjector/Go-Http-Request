package api

import (
	"cryptomasters/data"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const apiUrl = "https://cex.io/api/ticker/%s/USD"

func GetRate(currency string) (*data.Rate, error) {
	UpCurrency := strings.ToUpper(currency)
	res, err := http.Get(fmt.Sprintf(apiUrl, UpCurrency))

	if err != nil {
		return nil, err
	}
	var response CEXResponse
	if res.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(res.Body)

		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(bodyBytes, &response)
		if err != nil {
			return nil, err
		}

	} else {
		return nil, fmt.Errorf("status code recieved: %v", res.StatusCode)
	}
	rate := data.Rate{Currency: currency, Price: response.Bid}
	return &rate, nil
}
