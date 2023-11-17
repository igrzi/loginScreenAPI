# Login screen API

> This is an API for creating and checking users in a MySQL database.
## API Documentation

<div align="center">
  <h3>POST | User register</h3>
</div>

**Receives data from a JSON**

If the user is on the database, returns a 409 status code and a string to indicate user is already on the database.  
If the user isn't on the database, returns a 200 status code, encrypts the password and saves on the database

```http
  POST /submitUser
```

| Parameters   | Type       | Description                                |
| :---------- | :--------- | :---------------------------------- |
| `"Email"` | `string` | **Required**. New user email |
| `"Password"` | `string` | **Required**. New user password |

<div align="center">
  <h3>GET | User checking</h3>
</div>

Returns a 200 status code if the user is on the database.  
Returns a 404 status code if the user isn't on the database.

```http
  GET /checkUser?email=&password=
```

| Parameters   | Type       | Description                                |
| :---------- | :--------- | :------------------------------------------ |
| `email`      | `string` | **Required**. User email |
| `password`      | `string` | **Required**. User password |
