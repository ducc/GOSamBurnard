$(document).ready(function() {
    $(".modal-toggler").click(function() {
        var modalId = $(this).attr("data-toggle");
        $(modalId).addClass("is-active");
    });

    $(".modal-button").click(function() {
        var target = $(this).data('target');
        $('html').addClass('is-clipped');
        $(target).addClass('is-active');
    });

    $(".modal-background, .modal-close").click(function() {
        $('html').removeClass('is-clipped');
        $(this).parent().removeClass('is-active');
    });

    $(".admin-button").click(function() {
        var section = $(this).attr("data-toggle");
        var activeElement = $(".active-section");
        activeElement.removeClass("active-section");
        activeElement.css("display", "none");
        $(section).addClass("active-section");
        $(section).css("display", "");
    });
});