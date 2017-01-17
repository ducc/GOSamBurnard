$(document).ready(function() {
    var $toggle = $('#nav-toggle');
    var $menu = $('#nav-menu');
    var $navHeading = $('#nav-heading');

    $toggle.click(function() {
        $(this).toggleClass('is-active');
        $menu.toggleClass('is-active');
        $navHeading.toggleClass('is-hidden');
    });

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
        $(".admin-button").each(function() {
            $(this).css("background-color", "");
        });
        $(this).css("background-color", "#f5f5f5");
    });

    $(".slider-edit-button").click(function() {
        var editSectionId = $(this).attr("data-toggle");
        $(".slider-edit-section").each(function() {
            $(this).css("display", "none");
        });
        $(editSectionId).css("display", "");
    });

    $(".edit-section-close").click(function() {
        $(this).parent().css("display", "none");
    });

    $(".portfolio-edit-button").click(function() {
        var editSectionId = $(this).attr("data-toggle");
        $(".portfolio-edit-section").each(function() {
            $(this).css("display", "none");
        });
        $(editSectionId).css("display", "");
    });

    $(".notification-hide").click(function() {
        $(this).parent().remove();
    });

    var hash = $(location).attr("hash");
    if (hash.startsWith("#admin-")) {
        var elementId = "#" + hash.substring(7, hash.length) + "-section";

        var activeElement = $(".active-section");
        activeElement.removeClass("active-section");
        activeElement.css("display", "none");

        $(elementId).addClass("active-section");
        $(elementId).css("display", "");
    }
});