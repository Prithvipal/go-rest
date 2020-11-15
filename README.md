# go-rest

## APIs

### Create a new TODO

**API:** localhost:8080/api/v1/todo<br/>
**Method:** POST<br/>
**Payload**:
```
{
    "value": "Todo Number one"
}
```

**Response:**
```
{
    "msg": "TODO created successfully"
}
```

### List TODOs
**API:** localhost:8080/api/v1/todo<br/>
**Method:** GET<br/>
**Description:** When page and limit is not passed then this API will take page as 0 and limit as 10.<br/>
**Response:**
```
{
    "records": [
        {
            "id": "1",
            "value": "Todo Number one"
        },
        {
            "id": "2",
            "value": "Todo Number one"
        },
        {
            "id": "3",
            "value": "Todo Number one"
        },
        {
            "id": "4",
            "value": "Todo Number one"
        },
        {
            "id": "5",
            "value": "Todo Number one"
        },
        {
            "id": "6",
            "value": "Todo Number one"
        }
    ],
    "page": 0,
    "limit": 10,
    "count": 6,
    "total_count": 6,
    "total_pages": 0
}
```


### Get TODO by ID

### Delete TODO by ID

### Update TODO by ID
