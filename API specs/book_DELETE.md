##Delete book API 
Deescrription : This API will delete the book

### HTTP Request
`DELETE/bookid`

### URL Parameters
//books/delete

### Query Parameters
N/A


### Request Headers
```
Content-Type: application/x-www-form-urlencoded
```

### Request Body
| Parameter | Format | Description                                |
|-----------|--------|--------------------------------------------|



### Sample cURL request
```
'localhost:8004/books/94ba9ccc-ed96-4110-a957-29d0e0ee492a' 
 
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
    "Message": Book Deleted Successfuly"
}
```

### Bad Request Response when Book deletion failed
```
{
    "Message": "Invalid operation, please try again"
}
```

### Bad Request Response when book doesn't exist
```
{
    "Message": "Book does not exist"
}
```

### Forbidden Response when auth failed
```
{
    "Message": "Access Denied"
}
```