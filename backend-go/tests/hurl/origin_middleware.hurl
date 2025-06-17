# ----------------------------------------
# Access without origin
# ----------------------------------------
GET http://localhost:8000/health

HTTP 403

# ----------------------------------------
# Access with origin
# ----------------------------------------
GET http://localhost:8000/health
Origin: http://localhost:3000/

HTTP 200

# ----------------------------------------
# Access with incorrect origin
# ----------------------------------------
GET http://localhost:8000/health
Origin: http://localhost:5000

HTTP 403

# ----------------------------------------
# Fallback to referer if origin not exist
# ----------------------------------------
GET http://localhost:8000/health
Referer: http://localhost:3000

HTTP 200
