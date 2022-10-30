<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Login User</title>
    <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.2.2/dist/css/bootstrap.min.css" rel="stylesheet"
        integrity="sha384-Zenh87qX5JnK2Jl0vWa8Ck2rdkQ2Bzep5IDxbcnCeuOxjzrPF/et3URy9Bv1WTRi" crossorigin="anonymous">
    <link rel="preconnect" href="https://fonts.googleapis.com">
    <link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
    <link
        href="https://fonts.googleapis.com/css2?family=Josefin+Sans:ital,wght@0,300;0,400;0,500;0,600;0,700;1,300;1,400;1,500;1,600;1,700&display=swap"
        rel="stylesheet">
    <script src="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.2.0/js/all.min.js"
        integrity="sha512-naukR7I+Nk6gp7p5TMA4ycgfxaZBJ7MO5iC3Fp6ySQyKFHOGfpkSZkYVWV5R7u7cfAicxanwYQ5D1e17EfJcMA=="
        crossorigin="anonymous" referrerpolicy="no-referrer"></script>
</head>
<body>
    <div class="container">
        <div class="row mt-2 justify-content-center">
            <div class="col-lg-6">
                <h1 class="my-4">Login Form</h1>
                    <div id="alert"></div>
                    <div class="form-group mb-2">
                        <label for="first_name">First name*</label>
                        <input type="text" name="first_name" id="first_name" class="form-control" required>
                    </div>
                    <div class="form-group mb-2">
                        <label for="last_name">Last name*</label>
                        <input type="text" name="last_name" id="last_name" class="form-control" required>
                    </div>
                    <div class="form-group mb-2">
                        <label for="username">Username*</label>
                        <input type="text" name="username" id="username" class="form-control" required>
                    </div>
                    <div class="form-group mb-2">
                        <label for="email">Email*</label>
                        <input type="email" name="email" id="email" class="form-control" required>
                    </div>
                    <div class="form-group mb-2">
                        <label for="password1">Password*</label>
                        <input type="password" name="password1" id="password1" class="form-control" required>
                    </div>
                    <div class="form-group mb-2">
                        <label for="password2">Password Confirm*</label>
                        <input type="password" name="password2" id="password2" class="form-control" required>
                    </div>
                    <button id="loginHandler" type="submit" class="btn btn-primary">Login</button>
            </div>
        </div>
    </div>

    <script src="https://code.jquery.com/jquery-3.6.1.min.js"
        integrity="sha256-o88AwQnZB+VDvE9tvIXrMQaPlFFSUTR+nldQm1LuPXQ=" crossorigin="anonymous"></script>
    <script src="https://cdn.jsdelivr.net/npm/bootstrap@5.2.2/dist/js/bootstrap.bundle.min.js"
        integrity="sha384-OERcA2EqjJCMA+/3y+gxIOqMEjwtxJY7qPCqsdltbNJuaOe923+mo//f6V8Qbsw3"
        crossorigin="anonymous"></script>

    <script>
        $("#loginHandler").on("click", function() {
            var firstName = $("#first_name").val();
            var lastName = $("#last_name").val();
            var username = $("#username").val();
            var email = $("#email").val();
            var password1 = $("#password1").val();
            var password2 = $("#password2").val();

            if (password1 != password2) {
                var div = document.createElement("div")
                div.classList.add("alert")
                div.classList.add("alert-danger")
                div.classList.add("alert-dismissible")
                div.setAttribute("role", "alert")
                div.innerText = "Please enter password and password confirmation.";

                var btnClose = document.createElement("button");
                btnClose.classList.add("btn-close");
                btnClose.setAttribute("data-bs-dismiss", "alert");

                div.appendChild(btnClose)

                $("#alert").append(div);
                return
            }

            $.ajax({
                url: "/api/v1/auth/user/register",
                type: "post",
                dataType: "json",
                data: {
                    "username": username,
                    "password": password1,
                    "first_name": firstName,
                    "last_name": lastName,
                    "email": email,
                },
                success: function(response) {
                    if (response.status == "error") {
                        var div = document.createElement("div")
                        div.classList.add("alert")
                        div.classList.add("alert-danger")
                        div.classList.add("alert-dismissible")
                        div.setAttribute("role", "alert")
                        div.innerText = response.message;

                        var btnClose = document.createElement("button");
                        btnClose.classList.add("btn-close");
                        btnClose.setAttribute("data-bs-dismiss", "alert");

                        div.appendChild(btnClose)

                        $("#alert").append(div);
                    } else if (response.status == "success") {
                        var div = document.createElement("div")
                        div.classList.add("alert")
                        div.classList.add("alert-success")
                        div.classList.add("alert-dismissible")
                        div.setAttribute("role", "alert")
                        div.innerHTML = "Successfully to register. Please login now at <a href=\"/auth/login\">here</a>";

                        var btnClose = document.createElement("button");
                        btnClose.classList.add("btn-close");
                        btnClose.setAttribute("data-bs-dismiss", "alert");

                        div.appendChild(btnClose)

                        $("#alert").append(div);
                    }
                }
            })
        })
    </script>
</body>
</html>