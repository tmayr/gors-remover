# cors-remover-go

Simple reverse proxy to bypass any CORS blockage while developing

## usage

run main.go and make requests to your `:8080` passing `url` with your url as a query param

```
go run *.go
curl https://localhost:8080/?url=https://google.cl
```

## todo

- reorganize app
- dynamic port
- https (lego)
