$(document).ready(function() {
    $(".modal-toggler").click(function() {
        var modalId = $(".modal-toggler").attr("data-toggle");
        $(modalId).addClass("is-active");
    });

    $('.modal-button').click(function() {
        var target = $(this).data('target');
        $('html').addClass('is-clipped');
        $(target).addClass('is-active');
    });

    $('.modal-background, .modal-close').click(function() {
        $('html').removeClass('is-clipped');
        $(this).parent().removeClass('is-active');
    });
});