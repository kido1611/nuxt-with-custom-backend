# ----------------------------------------
# Create Session
# ----------------------------------------
GET http://localhost:8000/sanctum/csrf-cookie
Origin: http://localhost:3000/

HTTP 204
[Captures]
temp-app-session: cookie "app_session"
temp-csrf-header: header "X-XSRF-TOKEN"

# ----------------------------------------
# Register with data with csrf
# ----------------------------------------
POST http://localhost:8000/api/auth/register
Origin: http://localhost:3000/
X-XSRF-TOKEN: {{temp-csrf-header}}
[Cookies]
app_session: {{temp-app-session}}
{
  "email": "{{newUuid}}@test.com",
  "name": "test name",
  "password": "password"
}

HTTP 201
[Asserts]
jsonpath "$.data.email" exists
header "X-XSRF-TOKEN" exists
cookie "app_session" not exists
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
temp-app-session: cookie "app_session"
temp-csrf-header: header "X-XSRF-TOKEN"

# ----------------------------------------
# Login with correct data
# ----------------------------------------
POST http://localhost:8000/api/auth/login
Origin: http://localhost:3000/
X-XSRF-TOKEN: {{temp-csrf-header}}
[Cookies]
app_session: {{temp-app-session}}
{
  "email": "{{email}}",
  "password": "password"
}

HTTP 200
[Asserts]
header "X-XSRF-TOKEN" exists
cookie "app_session" exists
jsonpath "$.data.email" == {{email}}
header "X-XSRF-TOKEN" != {{temp-csrf-header}}
[Captures]
app-session: cookie "app_session"
csrf-header: header "X-XSRF-TOKEN"

# ----------------------------------------
# Get user notes
# ----------------------------------------
GET http://localhost:8000/api/notes
Origin: http://localhost:3000/
X-XSRF-TOKEN: {{csrf-header}}
[Cookies]
app_session: {{app-session}}

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
[Cookies]
app_session: {{app-session}}
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
[Cookies]
app_session: {{app-session}}

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
[Cookies]
app_session: {{app-session}}

HTTP 204

# ----------------------------------------
# Get user notes
# ----------------------------------------
GET http://localhost:8000/api/notes
Origin: http://localhost:3000/
X-XSRF-TOKEN: {{csrf-header}}
[Cookies]
app_session: {{app-session}}

HTTP 200
[Asserts]
jsonpath "$.data" exists
jsonpath "$.data" count == 0
