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
            // Construct the API URLc
            const apiUrl = "http://localhost:3000/submitUser";

            // Create an object with the email and password
            const data = {
                Email: emailValue,
                Password: passwordValue
            };

            // Make a POST request to the API
            fetch(apiUrl, {
                method: "POST",
                headers: {
                    "Content-Type": "application/json"
                },
                body: JSON.stringify(data)
            })
            .then(response => {
                // Check if the response status is 200
                if (response.status === 200) {
                    // Redirect to another HTML page
                    window.location.href = "/static/pages/login.html";
                } else {
                    // Handle other response statuses here
                    return response.json();
                }
            })
            .then(result => {
                // Handle the API response here
                console.log("API Response:", result);
            })
            .catch(error => {
                console.error("Error:", error);
            });
        }
    });
});
