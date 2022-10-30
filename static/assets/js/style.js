// Scroll navbar dark
$(window).scroll(function () {
    // console.log($(window).scrollTop());
    if ($(window).scrollTop() <= 100) {
        if ($("#navbar").hasClass("bg-dark")) {
            $("#navbar").removeClass("bg-dark");
            $("#navbar").addClass("bg-light");
            $("#navbar").removeClass("navbar-dark");
        }

        if ($("#toTop").hasClass("show")) {
            $("#toTop").removeClass("show");
            $("#toTop").addClass("hide");
        }
    }

    if ($(window).scrollTop() >= 100) {
        if ($("#navbar").hasClass("bg-light")) {
            $("#navbar").removeClass("bg-light");
            $("#navbar").addClass("bg-dark");
            $("#navbar").addClass("navbar-dark");
        }

        if ($("#toTop").hasClass("hide")) {
            $("#toTop").removeClass("hide");
            $("#toTop").addClass("show");
        }
    }
})

// Scroll To Top
$("#toTop").on("click", function () {
    $("html, body").animate({
        scrollTop: 0,
    })
})