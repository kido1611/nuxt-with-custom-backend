# ----------------------------------------
# Check Health
# ----------------------------------------
GET http://localhost:8000/health
Origin: http://localhost:3000/

HTTP 200 
[Asserts]
cookie "X-XSRF-TOKEN" not exists

# ----------------------------------------
# Register with data without CSRF
# ----------------------------------------
POST http://localhost:8000/api/auth/register
Origin: http://localhost:3000/
{
  "email": "{{newUuid}}@test.com",
  "name": "test name",
  "password": "password"
}

HTTP 419
[Asserts]
cookie "X-XSRF-TOKEN" not exists

# ----------------------------------------
# Create Session
# ----------------------------------------
GET http://localhost:8000/sanctum/csrf-cookie
Origin: http://localhost:3000/

HTTP 204
[Asserts]
cookie "app_session" exists
cookie "XSRF-TOKEN" exists
[Captures]
temp-csrf-header: cookie "XSRF-TOKEN"

# ----------------------------------------
# Register without data with CSRF
# ----------------------------------------
POST http://localhost:8000/api/auth/register
Origin: http://localhost:3000/
X-XSRF-TOKEN: {{temp-csrf-header}}

HTTP 400
[Asserts]
cookie "XSRF-TOKEN" exists

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
cookie "XSRF-TOKEN" == ""
cookie "app_session" exists
cookie "app_session" == ""
[Captures]
name: jsonpath "$['data']['name']"
email: jsonpath "$['data']['email']"

# ----------------------------------------
# Register with data with old csrf
# ----------------------------------------
POST http://localhost:8000/api/auth/register
Origin: http://localhost:3000/
X-XSRF-TOKEN: {{temp-csrf-header}}
{
  "email": "{{newUuid}}@test.com",
  "name": "test name",
  "password": "password"
}

HTTP 419
[Asserts]
cookie "XSRF-TOKEN" not exists
cookie "app_session" not exists

# ----------------------------------------
# Get user with session (cookie)
# ----------------------------------------
GET http://localhost:8000/api/user
Origin: http://localhost:3000/

HTTP 401

# ----------------------------------------
# Login without data and without CSRF
# ----------------------------------------
POST http://localhost:8000/api/auth/login
Origin: http://localhost:3000/

HTTP 419

# ----------------------------------------
# Create Session
# ----------------------------------------
GET http://localhost:8000/sanctum/csrf-cookie
Origin: http://localhost:3000/

HTTP 204
[Asserts]
cookie "app_session" exists
cookie "XSRF-TOKEN" exists
[Captures]
temp-csrf-header: cookie "XSRF-TOKEN"

# ----------------------------------------
# Login with incorrect data
# ----------------------------------------
POST http://localhost:8000/api/auth/login
Origin: http://localhost:3000/
X-XSRF-TOKEN: {{temp-csrf-header}}
{
  "email": "{{email}}",
  "password": "incorrent-password"
}

HTTP 401
[Asserts]
cookie "XSRF-TOKEN" exists

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
cookie "XSRF-TOKEN" != {{temp-csrf-header}}
jsonpath "$.data.email" == {{email}}
[Captures]
csrf-header: cookie "XSRF-TOKEN"

# ----------------------------------------
# Get user without session 
# ----------------------------------------
GET http://localhost:8000/api/user
Origin: http://localhost:3000/

HTTP 200
[Asserts]
cookie "XSRF-TOKEN" exists
cookie "XSRF-TOKEN" == {{csrf-header}}
cookie "app_session" not exists
jsonpath "$.data.name" == {{name}}
jsonpath "$.data.email" == {{email}}

# ----------------------------------------
# Login with current active session (cookie)
# ----------------------------------------
POST http://localhost:8000/api/auth/login
Origin: http://localhost:3000/
X-XSRF-TOKEN: {{csrf-header}}
{
  "email": "{{email}}",
  "password": "password"
}

HTTP 403
[Asserts]
cookie "XSRF-TOKEN" exists
cookie "XSRF-TOKEN" == {{csrf-header}}
cookie "XSRF-TOKEN" != {{temp-csrf-header}}
cookie "app_session" not exists

# ----------------------------------------
# Login with current active session (cookie) with old CSRF token
# ----------------------------------------
POST http://localhost:8000/api/auth/login
Origin: http://localhost:3000/
X-XSRF-TOKEN: {{temp-csrf-header}}
{
  "email": "{{email}}",
  "password": "password"
}

HTTP 419

# ----------------------------------------
# Register with current active session (cookie)
# ----------------------------------------
POST http://localhost:8000/api/auth/register
Origin: http://localhost:3000/
X-XSRF-TOKEN: {{csrf-header}}
{
  "email": "{{newUuid}}@test.com",
  "name": "test name",
  "password": "password"
}

HTTP 403
[Asserts]
cookie "XSRF-TOKEN" exists
cookie "XSRF-TOKEN" == "{{csrf-header}}"
cookie "XSRF-TOKEN" != {{temp-csrf-header}}
cookie "app_session" not exists

# ----------------------------------------
# Register with current active session (cookie) with old CSRF token
# ----------------------------------------
POST http://localhost:8000/api/auth/register
Origin: http://localhost:3000/
X-XSRF-TOKEN: {{temp-csrf-header}}
{
  "email": "{{newUuid}}@test.com",
  "name": "test name",
  "password": "password"
}

HTTP 419

# ----------------------------------------
# Logout without CSRF Token
# ----------------------------------------
DELETE http://localhost:8000/api/auth/logout
Origin: http://localhost:3000/

HTTP 419

# ----------------------------------------
# Logout
# ----------------------------------------
DELETE http://localhost:8000/api/auth/logout
Origin: http://localhost:3000/
X-XSRF-TOKEN: {{csrf-header}}

HTTP 204
[Asserts]
cookie "app_session" exists
cookie "app_session[Value]" == ""
cookie "XSRF-TOKEN" exists
cookie "XSRF-TOKEN[Value]" == ""

# ----------------------------------------
# Get user with session (cookie)
# ----------------------------------------
GET http://localhost:8000/api/user
Origin: http://localhost:3000/

HTTP 401
