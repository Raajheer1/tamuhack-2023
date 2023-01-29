curl -X POST --compressed "https://api.duffel.com/air/offer_requests?return_offers=false&supplier_timeout=10000" \
  -H "Accept-Encoding: gzip" \
  -H "Accept: application/json" \
  -H "Content-Type: application/json" \
  -H "Duffel-Version: v1" \
  -H "Authorization: Bearer duffel_test_D_iatNCJUYaoBy6Q3n5JR4LVHkGIDVyT9wJcTrhPvyt" \
  -d '{
  "data": {
    "slices" : [
      {
        "origin": "NYC",
        "destination": "ATL",
        "departure_date": "2023-06-21"
      },
      {
        "origin": "ATL",
        "destination": "NYC",
        "departure_date": "2023-07-21"
      }
    ],
    "passengers": [{ "type": "adult" }, { "type": "adult" }, { "age": 1 }],
    "cabin_class": "business"
  }
}'