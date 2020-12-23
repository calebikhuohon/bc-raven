package circle

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"log"
	"net"
	"net/http"
)

type CardDetails struct {
	Number string `json:"number"`
	Cvv    string `json:"cvv"`
}

type Billing struct {
	Name       string `json:"name"`
	City       string `json:"city"`
	Country    string `json:"country"`
	Line1      string `json:"line1"`
	PostalCode string `json:"postalCode"`
	ExpMonth   int32  `json:"expMonth"`
	ExpYear    int32  `json:"expYear"`
}

type Metadata struct {
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	SessionId   string `json:"sessionId"`
	IpAddress   string `json:"ipAddress"`
}

type Card struct {
	IdempotencyKey string      `json:"idempotencyKey"`
	KeyId          string      `json:"keyId"`
	CardDetails    CardDetails `json:"encryptedData"`
	BillingDetails Billing     `json:"billingDetails"`
	Metadata       Metadata    `json:"metadata"`
}

type PaymentAmount struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}

type Source struct {
	Id   string `json:"id"`
	Type string `json:"type"`
}

type PaymentData struct {
	Cvv string `json:"cvv"`
}

type Payment struct {
	IdempotencyKey string        `json:"idempotencyKey"`
	KeyId          string        `json:"keyId"`
	Metadata       Metadata      `json:"metadata"`
	Amount         PaymentAmount `json:"amount"`
	Verification   string        `json:"verification"`
	Source         Source        `json:"source"`
	Description    string        `json:"description"`
	PaymentData    PaymentData   `json:"encryptedData"`
}

func CreateCard(card *Card) (map[string]interface{}, error) {
	res, err := CreateCardResponse(card)
	if err != nil {
		return nil, err
	}

	result := map[string]interface{}{}
	if err := json.Unmarshal([]byte(res), &result); err != nil {
		return nil, err
	}

	return result, nil
}

func GetIPAddress() (string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}

	for _, i := range ifaces {
		//if i.Flags&net.FlagUp == 0 {
		//	continue
		//}
		//if i.Flags&net.FlagLoopback != 0 {
		//	continue
		//}

		addrs, err := i.Addrs()
		if err != nil {
			return "", err
		}

		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}

			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue
			}
			return ip.String(), nil
		}
	}
	return "", errors.New("no network connection")
}

func MakeCardPayment(p *Payment) (map[string]interface{}, error) {
	response, err := MakePaymentResponse(p)
	if err != nil {
		return nil, err
	}
	result := map[string]interface{}{}
	if err := json.Unmarshal([]byte(response), &result); err != nil {
		return nil, err
	}
	return result, nil
}

func CheckCardPaymentStatus(status string) bool {
	var isCardChargeSuccessful bool
	if status == "confirmed" {
		isCardChargeSuccessful = true
	}
	//TODO update transaction in db

	return isCardChargeSuccessful
}

func StartSubscriber() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/subscribe", Subscribe)

	log.Fatal(http.ListenAndServe(":55005", router))
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func Subscribe(w http.ResponseWriter, r *http.Request) {
	var res map[string]interface{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&res); err != nil {
		respondWithError(w, http.StatusBadRequest, "invalid request payload")
		return
	}
	defer r.Body.Close()

	status := res["Message"].(map[string]interface{})["Status"].(string)
	isSuccess := CheckCardPaymentStatus(status)
	if !isSuccess {
		respondWithError(w, http.StatusBadRequest, "Payment is not yet confirmed")
		return
	}
	respondWithJSON(w, http.StatusOK, res)
}

func CheckPaymentSettlementStatus() {

}
