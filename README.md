# multisignature-service
 
run the following to start server. I think that dockerfile is later needed.

```
cd back-end

go run main.go
```

# API mock data
1. POST /save-tx

```
{
    "tx_id" : "124",
    "tx_body" : "this is a transaction body"
}
```

2. POST /save-sign

```
{
    "tx_id": "124",
    "address": "cosmos123",
    "sign_body": "this is a sign body"
}
```
