## Edit book info
Deescrription : This API will edit the book info

### HTTP Request
`PUT/book/bookid`

### URL Parameters
/book/bookid

### Query Parameters
N/A


### Request Headers
```
Content-Type: application/x-www-form-urlencoded
```

### Request Body
| Parameter | Format | Description                                |
|-----------|--------|--------------------------------------------|
| bookName     | Book name |
| BookId   | String | ID of bbook that is to be updated       |


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
    "Message": Book info updated sucessfully"
}
```

### Bad Request Response when Password validation failed
```
{
    "Message": "Invalid password. Must be atleast 8 characters long with at least 1 capital letter, 1 small letter, 1 digit and 1 symbol"
}
```

### Bad Request Response when invalid info entered
```
{
    "Message": "Invalid info entered, try again"
}
```

### Forbidden Response when Book not present 
```
{
    "Message": "Unable to verify token. Please contact your administrator"
}
```