$(document).ready(function() {

    var d = new Date();
    var n = d.getHours();
    if (n > 18 || n < 6) {
        document.body.className = "night";
    }else {
        document.body.className = "day";
    }

    $('#english').load("terms.txt").html();
    $('#arabic').load("arabic.html").html();

    $('fieldset').addClass("ui-widget ui-widget-content ui-corner-all");
    $('input').addClass("ui-widget ui-widget-content ui-corner-all");
    //$("#login").hide();
    //$("#connectForm").submit(function(e) {
    //    e.preventDefault();
    //    $("#connect").hide();
    //    $("#login").show();
    //    $("#terms").hide();
    //});

    var tsantsA = Math.random().toString(36).substr(2, 5)
    $('#tsantsa').value = tsantsA

    $("form[name='lForm']").validate({

        rules: {

            phone: {
                required: true,
                digits: true,
                minlength: 8,
                maxlength: 8
            }
        },

        messages: {

            phone: "please enter a valid phone number of 8 digits"

        },

        submitHandler: function(form) {
            form.submit();
        }
    });

    $("form[name='vForm']").validate({

        rules: {

            pincode: {
                required: true,
                digits: true,
                minlength: 4,
                maxlength: 4
            }
        },

        messages: {

            pincode: "please enter a valid pincode number of 4 digits"

        },

        submitHandler: function(form) {
            form.submit();
        }
    });


    var method;
    var noop = function () {};
    var methods = [
        'assert', 'clear', 'count', 'debug', 'dir', 'dirxml', 'error',
        'exception', 'group', 'groupCollapsed', 'groupEnd', 'info', 'log',
        'markTimeline', 'profile', 'profileEnd', 'table', 'time', 'timeEnd',
        'timeline', 'timelineEnd', 'timeStamp', 'trace', 'warn'
    ];
    var length = methods.length;
    var console = (window.console = window.console || {});

    while (length--) {
        method = methods[length];

        // Only stub undefined methods.
        if (!console[method]) {
            console[method] = noop;
        }
    }
    
});