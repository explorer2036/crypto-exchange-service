### Clone and build the application
```
$ make
```

### Run with binance real network
```
$ TEST_NETWORK=no ./crypto-exchange-service serve

$ TEST_NETWORK=no API_KEY=xxxxx SECRET_KEY=xxxxx ./crypto-exchange-service serve
```

### Testing

**Query order by id and symbol**
```
curl -d '{"id":2750377045, "marketId":"LTCUSDT"}' -H "Content-Type: application/json" -X GET "http://localhost:9000/order"
```

**Create new order**
```
curl -d '{"exchange":"binance","marketId": "LTCUSDT","side":"buy","quantity":0.15,"price":125}' -H "Content-Type: application/json" -X POST "http://localhost:9000/order"
```
