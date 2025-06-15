# ----------------------------------------
# Check Health
# ----------------------------------------
GET http://localhost:8000/health

HTTP 200 
[Asserts]
header "X-XSRF-TOKEN" not exists

# ----------------------------------------
# Register with data without CSRF
# ----------------------------------------
POST http://localhost:8000/api/auth/register
{
  "email": "{{newUuid}}@test.com",
  "name": "test name",
  "password": "password"
}

HTTP 419
[Asserts]
header "X-XSRF-TOKEN" not exists

# ----------------------------------------
# Create Session
# ----------------------------------------
GET http://localhost:8000/sanctum/csrf-cookie

HTTP 204
[Captures]
temp-app-session: cookie "app_session"
temp-csrf-header: header "X-XSRF-TOKEN"

# ----------------------------------------
# Register without data with CSRF
# ----------------------------------------
POST http://localhost:8000/api/auth/register
X-XSRF-TOKEN: {{temp-csrf-header}}
[Cookies]
app_session: {{temp-app-session}}

HTTP 400
[Asserts]
header "X-XSRF-TOKEN" exists

# ----------------------------------------
# Register with data with csrf
# ----------------------------------------
POST http://localhost:8000/api/auth/register
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
# Register with data with old csrf
# ----------------------------------------
POST http://localhost:8000/api/auth/register
X-XSRF-TOKEN: {{temp-csrf-header}}
[Cookies]
app_session: {{temp-app-session}}
{
  "email": "{{newUuid}}@test.com",
  "name": "test name",
  "password": "password"
}

HTTP 419
[Asserts]
header "X-XSRF-TOKEN" not exists
cookie "app_session" exists
cookie "app_session[Value]" == "" 

# ----------------------------------------
# Login without data and without CSRF
# ----------------------------------------
POST http://localhost:8000/api/auth/login

HTTP 419

# ----------------------------------------
# Create Session
# ----------------------------------------
GET http://localhost:8000/sanctum/csrf-cookie

HTTP 204
[Captures]
temp-app-session: cookie "app_session"
temp-csrf-header: header "X-XSRF-TOKEN"

# ----------------------------------------
# Login with incorrect data
# ----------------------------------------
POST http://localhost:8000/api/auth/login
X-XSRF-TOKEN: {{temp-csrf-header}}
[Cookies]
app_session: {{temp-app-session}}
{
  "email": "{{email}}",
  "password": "incorrent-password"
}

HTTP 401
[Asserts]
header "X-XSRF-TOKEN" exists

# ----------------------------------------
# Login with correct data
# ----------------------------------------
POST http://localhost:8000/api/auth/login
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
# Get user without session 
# ----------------------------------------
GET http://localhost:8000/api/user

HTTP 401
[Asserts]
header "X-XSRF-TOKEN" not exists

# ----------------------------------------
# Get user with session (cookie)
# ----------------------------------------
GET http://localhost:8000/api/user
[Cookies]
app_session: "{{app-session}}"

HTTP 200
[Asserts]
header "X-XSRF-TOKEN" exists
header "X-XSRF-TOKEN" == {{csrf-header}}
jsonpath "$.data.name" == {{name}}
jsonpath "$.data.email" == {{email}}

# ----------------------------------------
# Login with current active session (cookie)
# ----------------------------------------
POST http://localhost:8000/api/auth/login
X-XSRF-TOKEN: {{csrf-header}}
[Cookies]
app_session: "{{app-session}}"
{
  "email": "{{email}}",
  "password": "password"
}

HTTP 403
[Asserts]
header "X-XSRF-TOKEN" exists
header "X-XSRF-TOKEN" == {{csrf-header}}
header "X-XSRF-TOKEN" != {{temp-csrf-header}}

# ----------------------------------------
# Login with current active session (cookie) with old CSRF token
# ----------------------------------------
POST http://localhost:8000/api/auth/login
X-XSRF-TOKEN: {{temp-csrf-header}}
[Cookies]
app_session: "{{app-session}}"
{
  "email": "{{email}}",
  "password": "password"
}

HTTP 419

# ----------------------------------------
# Register with current active session (cookie)
# ----------------------------------------
POST http://localhost:8000/api/auth/register
X-XSRF-TOKEN: {{csrf-header}}
[Cookies]
app_session: "{{app-session}}"
{
  "email": "{{newUuid}}@test.com",
  "name": "test name",
  "password": "password"
}

HTTP 403
[Asserts]
header "X-XSRF-TOKEN" exists
header "X-XSRF-TOKEN" == "{{csrf-header}}"
header "X-XSRF-TOKEN" != {{temp-csrf-header}}

# ----------------------------------------
# Register with current active session (cookie) with old CSRF token
# ----------------------------------------
POST http://localhost:8000/api/auth/register
X-XSRF-TOKEN: {{temp-csrf-header}}
[Cookies]
app_session: "{{app-session}}"
{
  "email": "{{newUuid}}@test.com",
  "name": "test name",
  "password": "password"
}

HTTP 419

# ----------------------------------------
# Logout without session and without CSRF Token
# ----------------------------------------
DELETE http://localhost:8000/api/auth/logout

HTTP 419

# ----------------------------------------
# Logout without session
# ----------------------------------------
DELETE http://localhost:8000/api/auth/logout
X-XSRF-TOKEN: {{csrf-header}}

HTTP 419

# ----------------------------------------
# Logout
# ----------------------------------------
DELETE http://localhost:8000/api/auth/logout
X-XSRF-TOKEN: {{csrf-header}}
[Cookies]
app_session: {{app-session}}

HTTP 204

# ----------------------------------------
# Get user with session (cookie)
# ----------------------------------------
GET http://localhost:8000/api/user
[Cookies]
app_session: "{{app-session}}"
[Options]
very-verbose: true

HTTP 401
