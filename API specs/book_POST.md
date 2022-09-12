## Add Book API
Deescrription : This API will be used to add new books in library

### HTTP Request
`POST/book`

### URL Parameters
/user/book/create

### Query Parameters
N/A


### Request Headers
```
Content-Type: application/x-www-form-urlencoded
```

### Request Body
| Parameter | Format | Description                                |
|-----------|--------|--------------------------------------------|
| Name     | String | Name of Book |
| Id   | String | Unique ID of book       |
| Publisher     | String | Author of Book |
| No of copies     | String | count  of Book copies |
| Status     | String | Status of Book |


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
    "Message": Added Book Successfully "
```

### Bad Request Response when book addition failed
```
{
    "Message": "Book addition failed, please try again"
}
```

### Forbidden Response when role doesn't match
```
{
    "Message": "Access restricted"
}
```