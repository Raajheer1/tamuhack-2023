curl -X POST --compressed "https://api.duffel.com/air/offer_requests?return_offers=false&supplier_timeout=10000" \
  -H "Accept-Encoding: gzip" \
  -H "Accept: application/json" \
  -H "Content-Type: application/json" \
  -H "Duffel-Version: v1" \
  -H "Authorization: Bearer duffel_test_D_iatNCJUYaoBy6Q3n5JR4LVHkGIDVyT9wJcTrhPvyt" \
  -d '{
  "data": {
    "slices": [
      {
        "origin": "LHR",
        "destination": "JFK",
        "departure_time": {
          "to": "17:00",
          "from": "09:45"
        },
        "departure_date": "2023-04-24",
        "arrival_time": {
          "to": "17:00",
          "from": "09:45"
        }
      }
    ],
    "private_fares": {
      "QF": [
        {
          "corporate_code": "FLX53",
          "tracking_reference": "ABN:2345678"
        }
      ],
      "UA": [
        {
          "corporate_code": "1234"
        }
      ]
    },
    "passengers": [
      {
        "family_name": "Earhart",
        "given_name": "Amelia",
        "loyalty_programme_accounts": [
          {
            "account_number": "12901014",
            "airline_iata_code": "BA"
          }
        ],
        "type": "adult"
      },
      {
        "age": 14
      },
      {
        "fare_type": "student"
      },
      {
        "age": 5,
        "fare_type": "contract_bulk_child"
      }
    ],
    "max_connections": 0,
    "cabin_class": "economy"
  }
}'