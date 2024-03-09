#!/bin/bash

# Set the URL for the POST request
url="http://127.0.0.1:8080/api/v1/vote"

# Set the JSON payload as a string
payload='{
  "vote_id": 1,
  "vote_options": [1, 2]
}'

credentials=sjxiang
# Set the cookie as a string
cookie='credentials=sjxiang; path=/;'

# Send the POST request with the JSON payload and cookie using curl
curl -s \
  -X POST \
  -H "Content-Type: application/json" \
  -H "Cookie: $cookie" \
  -d "$payload" \
  "$url"

