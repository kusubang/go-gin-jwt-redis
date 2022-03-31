# Go + Gin + JWT + Redis

Go언어로 인증서버 구현

## reference
- https://covenant.tistory.com/203

### VHandling Go configuration with Viper
- https://blog.logrocket.com/handling-go-configuration-viper/

```
curl -d "key1=value1&key2=value2" \
-H "Content-Type: application/x-www-form-urlencoded" \
-X POST http://localhost:8000/data
```

## Login 
### Request
```
curl -d '{"username":"username", "password":"password"}' \
-H "Content-Type: application/json" \
-X POST http://localhost:8080/login/
```
### Response
```
{
    "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NfdXVpZCI6IjRiYmZhYzZmLWEyOTgtNGVjYi04MDQ0LTNjNmM0MjQ0YTdiNiIsImF1dGhvcml6ZWQiOnRydWUsImV4cCI6MTY0ODY5MDI4NCwidXNlcl9pZCI6MX0.ltvGFAQJqG3bih65aEgX_krsrcVaBeaKS10Gzqy18nY",
    "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDkyOTQxODQsInJlZnJlc2hfdXVpZCI6IjNhZWRmNzIyLTYxMjMtNDRkOC04MmQzLWE2Zjk0MGYxZDM5MiIsInVzZXJfaWQiOjF9.NGrfOlEw-adU0scVcpxqLCxtnhojHkKPimzUSLVn7Oo"
}
```

## Logout
### Request
```
curl -d "key1=value1&key2=value2" \
-H "Authorization: a eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NfdXVpZCI6ImVkOTFhYjU4LTM0Y2MtNGVjOS05ZGM2LTBiNzg1MTNjZGVlZiIsImF1dGhvcml6ZWQiOnRydWUsImV4cCI6MTY0ODYxNDQ4NiwidXNlcl9pZCI6MX0._TMubyGpJUJpnrp5nlYgJki8ipAHIix0cW6tEUy_CcQ" \
-X POST http://localhost:8080/logout/
```
### Response
```
"Token is expired"
```

```
"unauthorized"
```
