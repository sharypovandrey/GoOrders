# Docker Compose: Go Orders REST API with MySQL database

In `docker-compose.yaml` replace `GOOGLE_MAPS_KEY` key with real Google Maps API Key.  

Then run `docker-compose up --build` and open `http://localhost:8080`

## CURL commands to check an API
PLACE ORDER: 

`
curl --location --request POST "http://mytrainer.tech/orders" \
  --header "Content-Type: application/json" \
  --data "{
    \"origin\": [\"22.338622\", \"114.167996\"],
    \"destination\": [\"22.294748\", \"114.172452\"]
}"
`

TAKE ORDER: 


`
curl --location --request PATCH "http://mytrainer.tech/orders/1" \
  --header "Content-Type: application/json" \
  --data "{
	\"status\": \"TAKEN\"
}"
`


ORDERS LIST: 

`
curl --location --request GET "http://mytrainer.tech/orders?page=1&limit=3" \
  --header "Content-Type: application/json"
`