package duffel

import "time"

type OfferRequestResponse struct {
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
