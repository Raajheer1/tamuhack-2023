package duffel

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type SingleOfferResponse struct {
	Data Offer `json:"data"`
}

func GetSingleOfferById(offer_id string) (Offer, error) {
	req, err := http.NewRequest("GET", "https://api.duffel.com/air/offers/{id}?return_available_services=false", nil)
	if err != nil {
		return Offer{}, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Duffel-Version", "v1")
	req.Header.Set("Authorization", os.ExpandEnv("Bearer $DUFFEL_API_TOKEN"))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return Offer{}, err
	}
	defer resp.Body.Close()
	fmt.Println(resp.Status, "Single offer request")

	singleOfferResp := SingleOfferResponse{}
	err = json.NewDecoder(resp.Body).Decode(&singleOfferResp)
	if err != nil {
		return Offer{}, err
	}

	return singleOfferResp.Data, nil

}
