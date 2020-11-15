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
**Description:** When page and limit is not passed then this API will take page = 0 and limit = 10 as default.<br/>
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
            "value": "Todo Number two"
        },
        {
            "id": "3",
            "value": "Todo Number three"
        },
        {
            "id": "4",
            "value": "Todo Number four"
        },
        {
            "id": "5",
            "value": "Todo Number five"
        },
        {
            "id": "6",
            "value": "Todo Number six"
        }
    ],
    "page": 0,
    "limit": 10,
    "count": 6,
    "total_count": 6,
    "total_pages": 0
}
```

**API:** localhost:8080/api/v1/todo?page=0<br/>
**Method:** GET<br/>
**Description:** API page starts with 0. If there is only one page then we need to pass 0 as first page. 
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
        },
        {
            "id": "7",
            "value": "Todo Number 7"
        },
        {
            "id": "8",
            "value": "Todo Number 8"
        },
        {
            "id": "9",
            "value": "Todo Number 9"
        },
        {
            "id": "10",
            "value": "Todo Number 10"
        }
    ],
    "page": 0,
    "limit": 10,
    "count": 10,
    "total_count": 12,
    "total_pages": 1
}
```

### Get TODO by ID

### Delete TODO by ID

### Update TODO by ID
