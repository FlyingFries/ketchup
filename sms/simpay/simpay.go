package simpay

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

var InvalidNumerReturned = errors.New("Invalid number retured")
var BadCode = errors.New("bad code")
var InvalidNumber = errors.New("invalid number")

type Auth struct {
	Key    string `json:"key"`
	Secret string `json:"secret"`
}

type Error struct {
	Code  string      `json:"error_code"`
	Name  string      `json:"error_name"`
	Value interface{} `json:"error_value"`
}

func (err Error) Error() string {
	return fmt.Sprintf("SimpayError %s %s %v", err.Code, err.Name, err.Value)
}

type checkParams struct {
	Auth      Auth   `json:"auth"`
	ServiceId string `json:"service_id"`
	Number    string `json:"number"`
	Code      string `json:"code"`
}

func Check(auth Auth, service_id, number, code string) (price int, err error) {
	params := &struct {
		Params checkParams `json:"params"`
	}{Params: checkParams{Auth: auth, ServiceId: service_id, Number: number, Code: code}}

	var buf = &bytes.Buffer{}
	err = json.NewEncoder(buf).Encode(params)
	if err != nil {
		return
	}

	price, err = NumberToPrice(number)
	if err != nil {
		return
	}

	resp, err := http.Post("https://simpay.pl/api/1/status", "application/json", buf)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return 0, fmt.Errorf("%d != 200", resp.StatusCode)
	}

	json_data := struct {
		Error []Error `json:"error"`
		//Params interface{}
		Respond struct {
			Status string `json:"status"`
		} `json:"respond"`
	}{}
	err = json.NewDecoder(resp.Body).Decode(&json_data)
	if err != nil {
		return 0, err
	}
	if len(json_data.Error) > 0 {
		return 0, json_data.Error[0]
	}
	if json_data.Respond.Status != "OK" {
		return 0, BadCode
	}

	return
}

// konwertuje numer na cene netto sms w groszach
func NumberToPrice(number string) (price int, err error) {
	switch {
	case number[0:1] == "7":
		price, err = strconv.Atoi(number[1:2])
		if err != nil {
			return 0, err
		}
		if price == 0 {
			price = 50
		} else {
			price *= 100
		}
	case number[0:1] == "9":
		price, err = strconv.Atoi(number[1:3])
		if err != nil {
			return 0, err
		}
		if price < 10 {
			return 0, errors.New("invalid price")
		}
		price *= 100
	default:
		return 0, InvalidNumber
	}
	return
}
