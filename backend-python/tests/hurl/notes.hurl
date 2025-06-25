# ----------------------------------------
# Create Session
# ----------------------------------------
GET http://localhost:8000/sanctum/csrf-cookie
Origin: http://localhost:3000/

HTTP 204
[Captures]
temp-csrf-header: cookie "XSRF-TOKEN"

# ----------------------------------------
# Register with data with csrf
# ----------------------------------------
POST http://localhost:8000/api/auth/register
Origin: http://localhost:3000/
X-XSRF-TOKEN: {{temp-csrf-header}}
{
  "email": "{{newUuid}}@test.com",
  "name": "test name",
  "password": "password"
}

HTTP 201
[Asserts]
jsonpath "$.data.email" exists
cookie "XSRF-TOKEN" exists
cookie "XSRF-TOKEN" == "\"\""
cookie "app_session" exists
cookie "app_session" == "\"\""
[Captures]
name: jsonpath "$['data']['name']"
email: jsonpath "$['data']['email']"

# ----------------------------------------
# Create Session
# ----------------------------------------
GET http://localhost:8000/sanctum/csrf-cookie
Origin: http://localhost:3000/

HTTP 204
[Captures]
temp-csrf-header: cookie "XSRF-TOKEN"

# ----------------------------------------
# Login with correct data
# ----------------------------------------
POST http://localhost:8000/api/auth/login
Origin: http://localhost:3000/
X-XSRF-TOKEN: {{temp-csrf-header}}
{
  "email": "{{email}}",
  "password": "password"
}

HTTP 200
[Asserts]
cookie "XSRF-TOKEN" exists
cookie "app_session" exists
jsonpath "$.data.email" == {{email}}
cookie "XSRF-TOKEN" != {{temp-csrf-header}}
[Captures]
csrf-header: cookie "XSRF-TOKEN"

# ----------------------------------------
# Get user notes
# ----------------------------------------
GET http://localhost:8000/api/notes
Origin: http://localhost:3000/
X-XSRF-TOKEN: {{csrf-header}}

HTTP 200
[Asserts]
jsonpath "$.data" exists
jsonpath "$.data" count == 0

# ----------------------------------------
# Create note
# ----------------------------------------
POST http://localhost:8000/api/notes
Origin: http://localhost:3000/
X-XSRF-TOKEN: {{csrf-header}}
{
  "title": "Title {{newUuid}}",
  "description": "Description {{newUuid}}"
}

HTTP 201
[Asserts]
jsonpath "$.data" exists
[Captures]
noteId : jsonpath "$.data.id"

# ----------------------------------------
# Get user notes
# ----------------------------------------
GET http://localhost:8000/api/notes
Origin: http://localhost:3000/
X-XSRF-TOKEN: {{csrf-header}}

HTTP 200
[Asserts]
jsonpath "$.data" exists
jsonpath "$.data" count == 1

# ----------------------------------------
# Delete user notes
# ----------------------------------------
DELETE http://localhost:8000/api/notes/{{noteId}}
Origin: http://localhost:3000/
X-XSRF-TOKEN: {{csrf-header}}

HTTP 204

# ----------------------------------------
# Get user notes
# ----------------------------------------
GET http://localhost:8000/api/notes
Origin: http://localhost:3000/
X-XSRF-TOKEN: {{csrf-header}}

HTTP 200
[Asserts]
jsonpath "$.data" exists
jsonpath "$.data" count == 0
