package mp

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const (
	MERCADOPAGO_PAYMENT             = "https://api.mercadopago.com/v1/payments?"
	MERCADOPAGO_ACCESS_TOKEN_PARAM  = "access_token="
	MERCADOPAGO_HEADER_CONTENT      = "Content-Type"
	MERCADOPAGO_CONTENT_JSON        = "application/json"
	MERCADOPAGO_ACCEPT              = "accept"
	MERCADOPAGO_PAYMENTS_METHOD     = "POST"
	MERCADOPAGO_PAYMENTS_METHOD_GET = "GET"
)

func Pay(mp MPRequestPayments, accessToken string) (MPResponsePayments, error) {

	var response MPResponsePayments

	client := &http.Client{}

	b := new(bytes.Buffer)

	json.NewEncoder(b).Encode(mp)

	req, err := http.NewRequest(MERCADOPAGO_PAYMENTS_METHOD, MERCADOPAGO_PAYMENT+MERCADOPAGO_ACCESS_TOKEN_PARAM+accessToken, b)

	if err != nil {

		return response, err

	}

	req.Header.Add(MERCADOPAGO_HEADER_CONTENT, MERCADOPAGO_CONTENT_JSON)

	req.Header.Add(MERCADOPAGO_ACCEPT, MERCADOPAGO_CONTENT_JSON)

	resp, err := client.Do(req)

	defer resp.Body.Close()

	if err != nil {
		return response, err
	}

	htmlData, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return response, err
	}

	c := []byte(string(htmlData))

	err = json.Unmarshal(c, &response)

	if err != nil {

		return response, err

	}

	return response, nil

}

func GetPay(idPay string, accessToken string) (MPResponsePayments, error) {

	var response MPResponsePayments

	client := &http.Client{}

	req, err := http.NewRequest(MERCADOPAGO_PAYMENTS_METHOD_GET, MERCADOPAGO_PAYMENT + "/" + idPay + MERCADOPAGO_ACCESS_TOKEN_PARAM + accessToken, nil)

	if err != nil {

		return response, err

	}

	req.Header.Add(MERCADOPAGO_HEADER_CONTENT, MERCADOPAGO_CONTENT_JSON)

	req.Header.Add(MERCADOPAGO_ACCEPT, MERCADOPAGO_CONTENT_JSON)

	resp, err := client.Do(req)

	defer resp.Body.Close()

	if err != nil {
		return response, err
	}

	htmlData, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return response, err
	}

	c := []byte(string(htmlData))

	err = json.Unmarshal(c, &response)

	if err != nil {

		return response, err

	}

	return response, nil

}
