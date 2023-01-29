package duffel

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Payload struct {
	Data Data `json:"data"`
}
type Slices struct {
	Origin        string `json:"origin"`
	Destination   string `json:"destination"`
	DepartureDate string `json:"departure_date"`
}
type Airport struct {
	IataCode string `json:"iata_code"`
}
type Passengers struct {
	Type string `json:"type,omitempty"`
	Age  int    `json:"age,omitempty"`
}
type Data struct {
	Slices     []Slices     `json:"slices"`
	Passengers []Passengers `json:"passengers"`
	CabinClass string       `json:"cabin_class"`
}

func Search(origin string, destination string, departureDate string, returnDate string) (*OfferList, error) {
	fmt.Println(origin, destination, departureDate, returnDate)
	var data Payload
	if departureDate == returnDate {
		data = Payload{
			Data: Data{
				Slices: []Slices{
					{
						Origin:        origin,
						Destination:   destination,
						DepartureDate: departureDate,
					},
				},
				Passengers: []Passengers{
					{
						Type: "adult",
					},
				},
			},
		}
	} else {
		data = Payload{
			Data: Data{
				Slices: []Slices{
					{
						Origin:        origin,
						Destination:   destination,
						DepartureDate: departureDate,
					},
					{
						Origin:        destination,
						Destination:   origin,
						DepartureDate: returnDate,
					},
				},
				Passengers: []Passengers{
					{
						Type: "adult",
					},
				},
			},
		}
	}
	payloadJson, err := json.Marshal(data)
	if err != nil {
		fmt.Println("err marshalling payload")
		return nil, err
	}
	body := bytes.NewReader(payloadJson)

	req, err := http.NewRequest("POST", "https://api.duffel.com/air/offer_requests?return_offers=false", body)
	if err != nil {
		fmt.Println("err creating http req")
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Duffel-Version", "v1")
	req.Header.Set("Authorization", os.ExpandEnv("Bearer $DUFFEL_API_TOKEN"))

	resp, err := http.DefaultClient.Do(req)
	fmt.Println(resp.Status, "Request offer ID")
	if err != nil {
		fmt.Println("err completing http req")
		return nil, err
	}
	defer resp.Body.Close()

	offerRequestResp := OfferRequestResponse{}
	err = json.NewDecoder(resp.Body).Decode(&offerRequestResp)
	fmt.Println(offerRequestResp)
	if err != nil {
		fmt.Println("err decoding response body")
		return nil, err
	}

	fmt.Println(offerRequestResp.Data.ID)
	searchURL := "https://api.duffel.com/air/offers?offer_request_id=" + offerRequestResp.Data.ID + "&sort=total_amount&limit=5"
	fmt.Println(searchURL)
	searchRequest, err := http.NewRequest("GET", searchURL, nil)
	searchRequest.Header.Set("Accept", "application/json")
	searchRequest.Header.Set("Content-Type", "application/json")
	searchRequest.Header.Set("Duffel-Version", "v1")
	searchRequest.Header.Set("Authorization", os.ExpandEnv("Bearer $DUFFEL_API_TOKEN"))

	searchResp, err := http.DefaultClient.Do(searchRequest)
	fmt.Println(searchResp.Status, "Search results")
	if err != nil {
		fmt.Println("err requesting search results")
		return nil, err
	}

	defer searchResp.Body.Close()

	offerList := OfferList{}
	err = json.NewDecoder(searchResp.Body).Decode(&offerList)
	fmt.Println(offerList)
	return &offerList, nil
}
