  curl -X GET --compressed "https://api.duffel.com/air/offers?offer_request_id=$1&sort=total_amount"\
    -H "Accept-Encoding: gzip"\
    -H "Accept: application/json"\
    -H "Duffel-Version: v1"\
    -H "Authorization: Bearer duffel_test_D_iatNCJUYaoBy6Q3n5JR4LVHkGIDVyT9wJcTrhPvyt"