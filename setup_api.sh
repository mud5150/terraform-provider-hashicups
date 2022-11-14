export HASHICUPS_TOKEN=$(curl -X POST localhost:19090/signup -d '{"username":"education", "password":"test123"}' | jq .token -r)
export HASHICUPS_USERNAME=education
export HASHICUPS_PASSWORD=test123
curl -X POST -H "Authorization: ${HASHICUPS_TOKEN}" localhost:19090/orders -d '[{"coffee": { "id":1 }, "quantity":4}, {"coffee": { "id":3 }, "quantity":3}]'
curl -X GET -H "Authorization: ${HASHICUPS_TOKEN}" localhost:19090/orders/1