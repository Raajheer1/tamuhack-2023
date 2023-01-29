package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

type SearchInput struct {
	DepartureDay      string `json:"departure_day" binding:"required"`
	ReturnDay         string `json:"return_day"`
	FlightDestination string `json:"flight_destination" binding:"required"`
	FlightOrigin      string `json:"flight_origin" binding:"required"`
}

func Search(c *gin.Context) {
	fmt.Println("Searching flights...")

	var input SearchInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	flights, err := DuffelQuery(input.FlightOrigin, input.FlightDestination, input.DepartureDay, input.ReturnDay)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, flights)

}

type Payload struct {
	Data Data `json:"data"`
}
type Slices struct {
	Origin        string `json:"origin"`
	Destination   string `json:"destination"`
	DepartureDate string `json:"departure_date"`
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
type DuffelResponse struct {
	Data struct {
		Slices []struct {
			OriginType string `json:"origin_type,omitempty"`
			Origin     struct {
				Type            string  `json:"type,omitempty"`
				TimeZone        string  `json:"time_zone,omitempty"`
				Name            string  `json:"name,omitempty"`
				Longitude       float64 `json:"longitude,omitempty"`
				Latitude        float64 `json:"latitude,omitempty"`
				ID              string  `json:"id,omitempty"`
				IcaoCode        string  `json:"icao_code,omitempty"`
				IataCountryCode string  `json:"iata_country_code,omitempty"`
				IataCode        string  `json:"iata_code,omitempty"`
				IataCityCode    string  `json:"iata_city_code,omitempty"`
				CityName        string  `json:"city_name,omitempty"`
				City            struct {
					Type            string      `json:"type,omitempty"`
					TimeZone        interface{} `json:"time_zone,omitempty"`
					Name            string      `json:"name,omitempty"`
					Longitude       interface{} `json:"longitude,omitempty"`
					Latitude        interface{} `json:"latitude,omitempty"`
					ID              string      `json:"id,omitempty"`
					IcaoCode        interface{} `json:"icao_code,omitempty"`
					IataCountryCode string      `json:"iata_country_code,omitempty"`
					IataCode        string      `json:"iata_code,omitempty"`
					IataCityCode    string      `json:"iata_city_code,omitempty"`
					CityName        interface{} `json:"city_name,omitempty"`
				} `json:"city,omitempty"`
				Airports interface{} `json:"airports,omitempty"`
			} `json:"origin,omitempty"`
			DestinationType string `json:"destination_type,omitempty"`
			Destination     struct {
				Type            string  `json:"type,omitempty"`
				TimeZone        string  `json:"time_zone,omitempty"`
				Name            string  `json:"name,omitempty"`
				Longitude       float64 `json:"longitude,omitempty"`
				Latitude        float64 `json:"latitude,omitempty"`
				ID              string  `json:"id,omitempty"`
				IcaoCode        string  `json:"icao_code,omitempty"`
				IataCountryCode string  `json:"iata_country_code,omitempty"`
				IataCode        string  `json:"iata_code,omitempty"`
				IataCityCode    string  `json:"iata_city_code,omitempty"`
				CityName        string  `json:"city_name,omitempty"`
				City            struct {
					Type            string      `json:"type,omitempty"`
					TimeZone        interface{} `json:"time_zone,omitempty"`
					Name            string      `json:"name,omitempty"`
					Longitude       interface{} `json:"longitude,omitempty"`
					Latitude        interface{} `json:"latitude,omitempty"`
					ID              string      `json:"id,omitempty"`
					IcaoCode        interface{} `json:"icao_code,omitempty"`
					IataCountryCode string      `json:"iata_country_code,omitempty"`
					IataCode        string      `json:"iata_code,omitempty"`
					IataCityCode    string      `json:"iata_city_code,omitempty"`
					CityName        interface{} `json:"city_name,omitempty"`
				} `json:"city,omitempty"`
				Airports interface{} `json:"airports,omitempty"`
			} `json:"destination,omitempty"`
			DepartureDate string    `json:"departure_date,omitempty"`
			CreatedAt     time.Time `json:"created_at,omitempty"`
		} `json:"slices,omitempty"`
		Passengers []struct {
			Type                     string `json:"type,omitempty"`
			LoyaltyProgrammeAccounts []struct {
				AirlineIataCode string `json:"airline_iata_code,omitempty"`
				AccountNumber   string `json:"account_number,omitempty"`
			} `json:"loyalty_programme_accounts,omitempty"`
			ID         string      `json:"id,omitempty"`
			GivenName  string      `json:"given_name,omitempty"`
			FareType   interface{} `json:"fare_type,omitempty"`
			FamilyName string      `json:"family_name,omitempty"`
			Age        interface{} `json:"age,omitempty"`
		} `json:"passengers,omitempty"`
		LiveMode   bool      `json:"live_mode,omitempty"`
		ID         string    `json:"id,omitempty"`
		CreatedAt  time.Time `json:"created_at,omitempty"`
		CabinClass string    `json:"cabin_class,omitempty"`
	} `json:"data,omitempty"`
}

func DuffelQuery(origin string, destination string, departureDate string, returnDate string) (*DuffelResponse, error) {
	data := Payload{
		Data: Data{
			Slices: []Slices{
				{
					Origin:        origin,
					Destination:   destination,
					DepartureDate: departureDate,
				},
				{
					Origin:        origin,
					Destination:   destination,
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
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", "https://api.duffel.com/air/offer_requests?return_offers=false", body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Duffel-Version", "v1")
	req.Header.Set("Authorization", os.ExpandEnv("Bearer $DUFFEL_API_TOKEN"))

	resp, err := http.DefaultClient.Do(req)
	fmt.Println(resp.Status)
	if err != nil {
		fmt.Println("err creating http req")
		return nil, err
	}
	defer resp.Body.Close()

	duffelResp := DuffelResponse{}
	err = json.NewDecoder(resp.Body).Decode(&duffelResp)
	if err != nil {
		fmt.Println("err decoding response body")
		return nil, err
	}

	if err != nil {
		fmt.Println("err decoding response body")
		return nil, err
	}

	return &duffelResp, nil
}
