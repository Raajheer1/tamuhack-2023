package duffel

import "time"

type Offer struct {
	Destination      string    `json:"destination,omitempty"`
	Origin           string    `json:"origin,omitempty"`
	DepartureDate    string    `json:"departure,omitempty"`
	UpdatedAt        time.Time `json:"updated_at,omitempty"`
	TotalEmissionsKg string    `json:"total_emissions_kg,omitempty"`
	TotalCurrency    string    `json:"total_currency,omitempty"`
	TotalAmount      string    `json:"total_amount,omitempty"`
	TaxCurrency      string    `json:"tax_currency,omitempty"`
	TaxAmount        string    `json:"tax_amount,omitempty"`
	Slices           []struct {
		Segments []struct {
			Passengers []struct {
				PassengerID             string `json:"passenger_id,omitempty"`
				FareBasisCode           string `json:"fare_basis_code,omitempty"`
				CabinClassMarketingName string `json:"cabin_class_marketing_name,omitempty"`
				CabinClass              string `json:"cabin_class,omitempty"`
				Baggages                []struct {
					Type     string `json:"type,omitempty"`
					Quantity int    `json:"quantity,omitempty"`
				} `json:"baggages,omitempty"`
			} `json:"passengers,omitempty"`
			OriginTerminal string `json:"origin_terminal,omitempty"`
			Origin         struct {
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
					Type            string `json:"type,omitempty"`
					Name            string `json:"name,omitempty"`
					ID              string `json:"id,omitempty"`
					IataCountryCode string `json:"iata_country_code,omitempty"`
					IataCode        string `json:"iata_code,omitempty"`
					IataCityCode    string `json:"iata_city_code,omitempty"`
				} `json:"city,omitempty"`
			} `json:"origin,omitempty"`
			OperatingCarrierFlightNumber string `json:"operating_carrier_flight_number,omitempty"`
			OperatingCarrier             struct {
				Name          string `json:"name,omitempty"`
				LogoSymbolURL string `json:"logo_symbol_url,omitempty"`
				ID            string `json:"id,omitempty"`
				IataCode      string `json:"iata_code,omitempty"`
			} `json:"operating_carrier,omitempty"`
			MarketingCarrierFlightNumber string `json:"marketing_carrier_flight_number,omitempty"`
			MarketingCarrier             struct {
				Name          string `json:"name,omitempty"`
				LogoSymbolURL string `json:"logo_symbol_url,omitempty"`
				ID            string `json:"id,omitempty"`
				IataCode      string `json:"iata_code,omitempty"`
			} `json:"marketing_carrier,omitempty"`
			ID                  string `json:"id,omitempty"`
			Duration            string `json:"duration,omitempty"`
			Distance            string `json:"distance,omitempty"`
			DestinationTerminal string `json:"destination_terminal,omitempty"`
			Destination         struct {
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
					Type            string `json:"type,omitempty"`
					Name            string `json:"name,omitempty"`
					ID              string `json:"id,omitempty"`
					IataCountryCode string `json:"iata_country_code,omitempty"`
					IataCode        string `json:"iata_code,omitempty"`
					IataCityCode    string `json:"iata_city_code,omitempty"`
				} `json:"city,omitempty"`
			} `json:"destination,omitempty"`
			DepartingAt string `json:"departing_at,omitempty"`
			ArrivingAt  string `json:"arriving_at,omitempty"`
			Aircraft    struct {
				Name     string `json:"name,omitempty"`
				ID       string `json:"id,omitempty"`
				IataCode string `json:"iata_code,omitempty"`
			} `json:"aircraft,omitempty"`
		} `json:"segments,omitempty"`
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
				Type            string `json:"type,omitempty"`
				Name            string `json:"name,omitempty"`
				ID              string `json:"id,omitempty"`
				IataCountryCode string `json:"iata_country_code,omitempty"`
				IataCode        string `json:"iata_code,omitempty"`
				IataCityCode    string `json:"iata_city_code,omitempty"`
			} `json:"city,omitempty"`
		} `json:"origin,omitempty"`
		SliceID         string `json:"id,omitempty"`
		FareBrandName   string `json:"fare_brand_name,omitempty"`
		Duration        string `json:"duration,omitempty"`
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
				Type            string `json:"type,omitempty"`
				Name            string `json:"name,omitempty"`
				ID              string `json:"id,omitempty"`
				IataCountryCode string `json:"iata_country_code,omitempty"`
				IataCode        string `json:"iata_code,omitempty"`
				IataCityCode    string `json:"iata_city_code,omitempty"`
			} `json:"city,omitempty"`
		} `json:"destination,omitempty"`
		Conditions struct {
			ChangeBeforeDeparture struct {
			} `json:"change_before_departure,omitempty"`
		} `json:"conditions,omitempty"`
	} `json:"slices,omitempty"`
	PaymentRequirements struct {
		PriceGuaranteeExpiresAt time.Time `json:"price_guarantee_expires_at,omitempty"`
		PaymentRequiredBy       time.Time `json:"payment_required_by,omitempty"`
	} `json:"payment_requirements,omitempty"`
	Passengers []struct {
		Type string `json:"type,omitempty"`
		ID   string `json:"id,omitempty"`
	} `json:"passengers,omitempty"`
	Owner struct {
		Name          string `json:"name,omitempty"`
		LogoSymbolURL string `json:"logo_symbol_url,omitempty"`
		OwnerID       string `json:"id,omitempty"`
		IataCode      string `json:"iata_code,omitempty"`
	} `json:"owner,omitempty" gorm:"embedded"`
	ID         string    `json:"id,omitempty"`
	ExpiresAt  time.Time `json:"expires_at,omitempty"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
	Conditions struct {
		RefundBeforeDeparture struct {
			PenaltyCurrency string `json:"penalty_currency,omitempty"`
			PenaltyAmount   string `json:"penalty_amount,omitempty"`
			Allowed         bool   `json:"allowed,omitempty"`
		} `json:"refund_before_departure,omitempty"`
		ChangeBeforeDeparture struct {
		} `json:"change_before_departure,omitempty"`
	} `json:"conditions,omitempty"`
	BaseCurrency string `json:"base_currency,omitempty"`
	BaseAmount   string `json:"base_amount,omitempty"`
}
