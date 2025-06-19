# ----------------------------------------
# Does not contain origign, does not contain CORS
# ----------------------------------------
GET http://localhost:8000/health
Referer: http://localhost:3000

HTTP 200
[Asserts]
header "Access-Control-Allow-Origin" not exists
header "Access-Control-Allow-Credentials" not exists
header "Access-Control-Allow-Methods" not exists
header "Access-Control-Allow-Headers" not exists
header "Vary" exists
header "Vary" == "Origin"


# ----------------------------------------
# Correct origin, contain CORS
# ----------------------------------------
GET http://localhost:8000/health
Origin: http://localhost:3000

HTTP 200
[Asserts]
header "Access-Control-Allow-Origin" exists
header "Access-Control-Allow-Credentials" exists
header "Access-Control-Allow-Methods" not exists
header "Access-Control-Allow-Headers" not exists
header "Vary" exists
header "Vary" == "Origin"


# ----------------------------------------
# Preflight
# ----------------------------------------
OPTIONS http://localhost:8000/api/notes
Access-Control-Request-Method: POST
Origin: http://localhost:3000

HTTP 204
[Asserts]
header "Access-Control-Allow-Origin" exists
header "Access-Control-Allow-Credentials" exists
header "Access-Control-Allow-Methods" exists
header "Access-Control-Allow-Headers" exists
header "Vary" exists
header "Vary" == "Origin"


# ----------------------------------------
# Incorrect origin, Does not contain CORS
# ----------------------------------------
GET http://localhost:8000/health
Origin: http://localhost:5000

HTTP 403
[Asserts]
header "Access-Control-Allow-Origin" not exists
header "Access-Control-Allow-Credentials" not exists
header "Access-Control-Allow-Methods" not exists
header "Access-Control-Allow-Headers" not exists
header "Vary" exists
header "Vary" == "Origin"


# ----------------------------------------
# Preflight, incorrect origin
# ----------------------------------------
OPTIONS http://localhost:8000/api/notes
Access-Control-Request-Method: POST
Origin: http://localhost:5000

HTTP 403
[Asserts]
header "Access-Control-Allow-Origin" not exists
header "Access-Control-Allow-Credentials" not exists
header "Access-Control-Allow-Methods" not exists
header "Access-Control-Allow-Headers" not exists
header "Vary" exists
header "Vary" == "Origin"
