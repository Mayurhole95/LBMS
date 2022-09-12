## Book Status
Deescrription : This API will display the status of Book (available or unavailable-issued already)

### HTTP Request
`GET/book/{bookid}`

### URL Parameters
/book/{bookid}

### Query Parameters
N/A


### Request Headers
```
Content-Type: application/x-www-form-urlencoded
```

### Request Body
| Parameter | Format | Description                                |
|-----------|--------|--------------------------------------------|
| BookName     | String | Book Name |
| Book ID   | String | Book id     |


### Sample cURL request
```

```

### Status codes and errors
| Value | Description           |
|-------|-----------------------|
| 200   | OK                    |
| 400   | Bad Request           |
| 403   | Forbidden             |
| 410   | Gone                  |
| 500   | Internal Server Error |

### Response Headers
N/A

### Success Response Body
```
{
    "Message": Status: ""
}
```

### Bad Request Response when Password validation failed
```
{
    "Message": "Invalid password. Must be atleast 8 characters long with at least 1 capital letter, 1 small letter, 1 digit and 1 symbol"
}
```

### Forbidden Response when Book not present in library
```
{
    "Message": "Unable to verify token. Please contact your administrator"
}
```