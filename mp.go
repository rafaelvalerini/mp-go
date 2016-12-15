package mp

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const (
	mercadopagoPayment           = "https://api.mercadopago.com/v1/payments?"
	mercadopagoAccessTokenParam  = "access_token="
	mercadopagoHeaderContent     = "Content-Type"
	mercadopagoContentJson       = "application/json"
	mercadopagoAccept            = "accept"
	mercadopagoPaymentsMethod    = "POST"
	mercadopagoPaymentsMethodGet = "GET"
)

func Pay(mp MPRequestPayments, accessToken string) (MPResponsePayments, error) {

	var response MPResponsePayments

	client := &http.Client{}

	b := new(bytes.Buffer)

	json.NewEncoder(b).Encode(mp)

	req, err := http.NewRequest(mercadopagoPaymentsMethod, mercadopagoPayment+mercadopagoAccessTokenParam+accessToken, b)

	if err != nil {

		return response, err

	}

	req.Header.Add(mercadopagoHeaderContent, mercadopagoContentJson)

	req.Header.Add(mercadopagoAccept, mercadopagoContentJson)

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

	req, err := http.NewRequest(mercadopagoPaymentsMethodGet, mercadopagoPayment+"/"+idPay+mercadopagoAccessTokenParam+accessToken, nil)

	if err != nil {

		return response, err

	}

	req.Header.Add(mercadopagoHeaderContent, mercadopagoContentJson)

	req.Header.Add(mercadopagoAccept, mercadopagoContentJson)

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
