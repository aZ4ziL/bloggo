{{ define "title" }}New Article{{ end }}

{{ define "content" }}
<!-- <div class="col-xl-6 col-md-6 mb-3">
    <div class="card shadow h-100">
        <img src="{{ .Logo }}" alt="" class="card-img-top">
        <div class="card-body">
            <h5 class="card-title">{{ .Title }}</h5>
            <p class="card-text">
                {{ .Desc }}
            </p>
        </div>
        <div class="card-footer d-flex align-items-center justify-content-between">
            <a href="generic.html" class="btn btn-primary">Read More</a>
            <div>
                <i class="fas fa-eye"></i>
                &nbsp;
                <i class="fas fa-comments"></i>
            </div>
        </div>
    </div>
</div> -->

<!-- Content -->
<div class="container pt-4">
    <div class="row">

        <!-- Right -->
        <div class="col-xl-8 col-md-12 mb-4">

            <div class="row" id="contents">
                <div class="d-flex align-items-center justify-content-between">
                    <h1 class="fw-bold">Newest Posts</h1>
                    {{ if .user }}
                    <button type="button" class="btn btn-primary" data-bs-toggle="modal"
                        data-bs-target="#addPostModal">Create New Post</button>
                    {{ end }}
                </div>
                <!-- Modal -->
                <!-- Modal -->
                <!-- Modal -->
                {{ if .user }}
                <div class="modal fade" id="addPostModal" data-bs-backdrop="static" data-bs-keyboard="false"
                    tabindex="-1" aria-labelledby="addPostModalLabel" aria-hidden="true">
                    <div class="modal-dialog">
                        <div class="modal-content">
                            <div class="modal-header">
                                <h1 class="modal-title fs-5" id="addPostModalLabel">Add Post</h1>
                                <button type="button" class="btn-close" data-bs-dismiss="modal"
                                    aria-label="Close"></button>
                            </div>
                            <div class="modal-body">
                                <form action="/api/v1/article" method="post" enctype="multipart/form-data" id="addForm">
                                    <input type="hidden" name="author_id" id="author_id" value="{{ .user.id }}">
                                    <div class="form-group mb-2">
                                        <label for="title">Title</label>
                                        <input type="text" name="title" id="title" class="form-control" required>
                                    </div>
                                    <div class="form-group mb-2">
                                        <label for="slug">Slug</label>
                                        <input type="text" name="slug" id="slug" class="form-control" required>
                                    </div>
                                    <div class="form-group mb-2">
                                        <label for="category">Category</label>
                                        <select name="category_id" id="category_id" class="form-control form-select"
                                            required>
                                            <option selected>Please select category</option>
                                            {{ range .categories }}
                                            <option value="{{ .ID }}">{{ .Name }}</option>
                                            {{ end }}
                                        </select>
                                    </div>
                                    <div class="form-group mb-2">
                                        <label for="logo">Logo</label>
                                        <input type="file" name="logo" id="logo" class="form-control">
                                    </div>
                                    <div class="form-group mb-2">
                                        <label for="desc">Description</label>
                                        <input type="text" name="desc" id="desc" class="form-control" max="255"
                                            required>
                                    </div>
                                    <div class="form-group mb-2">
                                        <label for="content">Content</label>
                                        <textarea name="content" id="content" rows="10" class="form-control"></textarea>
                                        <div class="form-text">Please use Markdown only!</div>
                                    </div>
                                    <div class="form-group mb-2">
                                        <label for="status">Status</label>
                                        <select name="status" id="status" class="form-select">
                                            <option value="DRAFTED" selected>Drafted</option>
                                            <option value="PUBLISHED">Published</option>
                                        </select>
                                    </div>
                                    <div class="d-grid gap-2">
                                        <button type="submit" class="btn btn-primary">Post</button>
                                        <button type="reset" class="btn btn-danger">Cancel</button>
                                        <button type="button" class="btn btn-secondary"
                                            data-bs-dismiss="modal">Close</button>
                                    </div>
                                </form>
                            </div>
                        </div>
                    </div>
                </div>
                <!-- Modal -->
                {{ end }}
            </div>

        </div>
        <!-- Right -->

        <!-- Left -->
        <div class="col-xl-4 col-md-12 mb-4">
            <h4>Popular Post</h4>

            <ul class="list-group list-group-flush">
                <li class="list-group-item d-flex align-items-center">
                    <a href="#">Lorem ipsum dolor, sit amet consectetur adipisicing elit. Magni, asperiores?</a>
                </li>
                <li class="list-group-item d-flex align-items-center">
                    <a href="#">Lorem ipsum dolor, sit amet consectetur adipisicing elit. Magni, asperiores?</a>
                </li>
                <li class="list-group-item d-flex align-items-center">
                    <a href="#">Lorem ipsum dolor, sit amet consectetur adipisicing elit. Magni, asperiores?</a>
                </li>
            </ul>
        </div>
        <!-- Left -->

    </div>
</div>
<!-- Content -->
{{ end }}

{{ define "script" }}
<script src="//cdn.jsdelivr.net/npm/sweetalert2@11"></script>
<script>
    $.ajax({
        url: "/api/v1/article",
        type: "GET",
        dataType: "JSON",
        success: function (response) {
            for (var i = 0; i < response.length; i++) {
                var col = document.createElement("div");
                col.classList.add("col-xl-6", "col-md-6", "mb-3");

                var card = document.createElement("div");
                card.classList.add("card", "shadow", "h-100");

                var img = document.createElement("img");
                img.classList.add("card-img-top");
                img.setAttribute("src", response[i].logo);

                var body = document.createElement("div");
                body.classList.add("card-body");
                var h5 = document.createElement("h5");
                h5.classList.add("card-title");
                h5.innerText = response[i].title;
                var p = document.createElement("p");
                p.classList.add("card-text");
                p.innerText = response[i].desc;


                body.appendChild(h5);
                body.appendChild(p)

                // Footer
                var footer = document.createElement("div");
                footer.classList.add("card-footer", "d-flex", "align-items-center", "justify-content-between");
                var btn = document.createElement("a");
                btn.classList.add("btn", "btn-primary");
                btn.setAttribute("href", "/detail/" + response[i].slug);
                btn.innerText = "Read More";
                footer.appendChild(btn);

                // Append
                card.appendChild(img);
                card.appendChild(body);
                card.appendChild(footer);

                col.appendChild(card);

                $("#contents").append(col);
            }
        }
    })
</script>

<script>
    $("#addForm").submit(function (event) {
        var formData = new FormData($(this)[0]);

        $.ajax({
            url: "/api/v1/article",
            type: "post",
            dataType: "json",
            data: formData,
            cache: false,
            async: false,
            contentType: false,
            processData: false,
            success: function (response) {
                $("#addPostModal").modal("hide");
                Swal.fire({
                    icon: response.status,
                    text: response.message,
                }).then(result => {
                    if (result.isConfirmed) {
                        window.location.reload();
                    }
                })
            }
        });

        event.preventDefault();
    })
</script>
{{ end }}