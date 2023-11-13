# Functions

### UserCreate 
1. Receives a Email and Password in JSON format
2. Checks if the user isn't already on the database
    * If the user isn't on the database, returns a 200 (Success), encrypts the password using SHA256 and saves it on the database;
    * If the user is on the database, returns a 409 (Conflict) and a string to indicate that the user is already on the database;

### UserCheck
1. Receives a Email and Password as query parameters in the URL
2. Checks if the user is on the database, return 404 if the user isn't found, and returns 200 if the user is found.