## cURL

```
curl -X POST "localhost:8080/login" \
  -u "admin:secret"

// { "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2
VybmFtZSI6InVzZXJuYW1lIiwiZXhwIjoxNzA3MTM5MDU5fQ.re0Mb00kWBT1_CrCRp7BkqPNs8PxDBclekT-m6uAw30" }

curl -H "Authorization: Bearer token
" -X GET "localhost:8080/resource"
```