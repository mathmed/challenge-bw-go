package pagseguro

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"os"
	"strings"

	dtos "github.com/mathmed/challenge-bw-go/src/UseCases/CreatePayment/Dtos"
)

type holder struct {
	Name string `json:"name"`
} 
type card struct {
	Number string `json:"number"`
	ExpMonth string `json:"exp_month"`
	ExpYear string `json:"exp_year"`
	SecurityCode string `json:"security_code"`
	Holder holder `json:"holder"`
}
type paymentMethod struct {
	Type string `json:"type"`
	Installments int `json:"installments"`
	Capture bool `json:"capture"`
	Card card `json:"card"`
}
type amount struct {
	Value float64 `json:"value"`
	Currency string `json:"currency"`
}
type pagSeguroRequestData struct {
	ReferenceID string `json:"reference_id"`
	Description string `json:"description"`
	Amount amount `json:"amount"`
	PaymentMethod paymentMethod `json:"payment_method"`
}
type pagSeguroResponseData struct {
	ID string `json:"id"`
	Status string `json:"status"`
	PaymentResponse struct {
		Code string `json:"code"`
	} `json:"payment_response"`
}

// CreatePaymentPagSeguro .
func CreatePaymentPagSeguro(paymentData dtos.CreatePayment) string {
	
	pagseguroPayload, err:= json.Marshal(createPayloadToSend(paymentData))

	if err != nil {
		fmt.Println(err)
	}

	req, err := http.NewRequest("POST", os.Getenv("PAGSEGURO_URL") + "charges", bytes.NewReader(pagseguroPayload))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("x-api-version", "application/json")
	req.Header.Add("Authorization", os.Getenv("PAGSEGURO_TOKEN"))

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Println(err)
	}

	var pagseguroResponse pagSeguroResponseData

	err = json.NewDecoder(res.Body).Decode(&pagseguroResponse)
	res.Body.Close()

	if err != nil {
		fmt.Println(err)
	}

	if pagseguroResponse.Status != "PAID" || 
		pagseguroResponse.PaymentResponse.Code != "20000" {
		fmt.Println("error")
	}

	return pagseguroResponse.ID
}


func createPayloadToSend(paymentData dtos.CreatePayment) pagSeguroRequestData {

	payload := pagSeguroRequestData{
		"Reference ID",
		"Description", 
		amount{math.Round(paymentData.Card.Amount) * 100, "BRL"},
		paymentMethod {
			"CREDIT_CARD",
			paymentData.Card.Parcels,
			true,
			card {
				paymentData.Card.Number,
				strings.Split(paymentData.Card.Expiration, "/")[0],
				"20" + strings.Split(paymentData.Card.Expiration, "/")[1],
				paymentData.Card.Cvv,
				holder {
					paymentData.Card.Owner,
				},
			},
		},
	}
	return payload
}