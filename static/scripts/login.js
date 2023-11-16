document.addEventListener("DOMContentLoaded", function () {
    // Get the username and password input elements
    const emailInput = document.getElementById("email");
    const passwordInput = document.getElementById("password");

    // Get the submit button
    const submitButton = document.getElementById("button");

    // Add an event listener to the submit button
    submitButton.addEventListener("click", function (event) {
        // Prevent the default form submission
        event.preventDefault();

        // Get the values from the input fields
        const emailValue = emailInput.value;
        const passwordValue = passwordInput.value;

        // Check if both email and password are provided
        if (emailValue !== "" && passwordValue !== "") {
            const apiUrl = `http://localhost:3000/checkUser?email=${emailValue}&password=${passwordValue}`;

            // Make an HTTP GET request to the server
            fetch(apiUrl)
            .then(response => {
                if (response.status === 200) {
                    // Redirect to another HTML page
                    window.location.href = "/static/index.html";
                } else {
                    // Handle other response statuses here
                    return response.json();
                }
            })
            .then(result => {
                // Handle the server response here
                console.log(result);
            })
            .catch(error => {
                console.error('Error:', error);
            });
        }
    });
});
