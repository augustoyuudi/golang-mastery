## cURL

```
curl -X POST "localhost:8080/login" \
  -u "admin:secret"

// { "token": "3e91398a70275ee96a16196a5df02f7b00b81d0b" }

curl -H "Authorization: Bearer token
" -X GET "localhost:8080/resource"
```