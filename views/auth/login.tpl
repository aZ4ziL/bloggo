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
        <div class="row mt-5 justify-content-center">
            <div class="col-lg-6">
                <h1 class="my-4">Login Form</h1>
                    <div id="alert"></div>
                    <div class="form-group mb-2">
                        <label for="username">Username</label>
                        <input type="text" name="username" id="username" class="form-control" required>
                    </div>
                    <div class="form-group mb-2">
                        <label for="password">Password</label>
                        <input type="password" name="password" id="password" class="form-control" required>
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
            var username = $("#username").val();
            var password = $("#password").val();

            $.ajax({
                url: "/api/v1/auth/user/login",
                type: "post",
                dataType: "json",
                data: {
                    "username": username,
                    "password": password,
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
                    } else {
                        window.location.href = "/";
                    }
                }
            })
        })
    </script>
</body>
</html>