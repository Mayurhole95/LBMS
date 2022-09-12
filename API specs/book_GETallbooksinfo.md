##Display all books
Deescrription : This API will display list of all books 

### HTTP Request
`GET/book

### URL Parameters
/book/info

### Query Parameters
N/A


### Request Headers
```
Content-Type: application/x-www-form-urlencoded
```

### Request Body
| Parameter | Format | Description                                |
|-----------|--------|--------------------------------------------|
| Book     | String | book name |
| BookID    | String | book ID |



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
    "Message": Book Info ""
}
```

### Bad Request Response when Password validation failed
```
{
    "Message": "Invalid password. Must be atleast 8 characters long with at least 1 capital letter, 1 small letter, 1 digit and 1 symbol"
}
```

### Forbidden Response when Email not valid
```
{
    "Message": "Unable to verify email. Please contact your administrator"
}
```
