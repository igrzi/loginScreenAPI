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

        // Print the values to the console
        if (emailValue != "" && passwordValue != ""){
            console.log("Email:", emailValue);
            console.log("Password:", passwordValue);
        }
    });
});
