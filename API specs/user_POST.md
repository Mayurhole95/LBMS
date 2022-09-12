## Creare new user
Deescrription : This API will create a new user

### HTTP Request
`POST /user

### URL Parameters
/user/create

### Query Parameters
N/A


### Request Headers
```
Content-Type: application/x-www-form-urlencoded
```

### Request Body
| Parameter | Format | Description                                |
|-----------|--------|--------------------------------------------|
| Email     | String | Email Id of user requesting password reset |
| Name   | String | Name of user      |
| Password   | String | password of user      |
| Gender   | String | Gender of user      |
| Role   | String | role of user (admin,superadmin,enduser)     |
| Number   | longint |Conact Number of user      |


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
    "Message": User Created Successfully "
}
```

### Bad Request Response when wrong info entered
```
{
    "Message": "Invalid info entered.Please check again"
}
```

### Bad Request Response when user already exists
```
{
    "Message": "Your username has already been set. Try logging in."
}


### Forbidden Response when field is empty
```
{
    "Message": Please enter valid credentials"
}
```
