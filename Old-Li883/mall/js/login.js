$(document).ready(function () {
    $("#submit").click(function () {
        value = $("input:radio[name='who']:checked").val();
        id = $("#id").val();
        password = $('#password').val();
        jdata = { "id": id, "password": password };
        $.ajax({
            url: '/api/login?status=' + value,
            data: JSON.stringify(jdata),
            type: 'post',
            dataType: 'json',
            headers: {
                Accept: "application/json",
                "Content-Type": "application/json"
            },
            processData: false,
            cache: false
        }).done(function (data, status) {
            if (data['status'] == 200) {
                alert("degnluchegngong");
            }
            else if (data['status'] == 403) {
                $("p").remove("#note");
                $("p").append("<p id='note'>you have login</p>");
            }
            else if (data['status'] == 400) {
                $("p").remove("#note");
                $("p").append('<p id="note">you have not been identified</p>');
            }
            else if (data['status'] == 401) {
                $("p").remove("#note");
                $("p").append('<p id="note">password wrong</p>');
            }
            else if (data['status'] == 402) {
                $("p").remove("#note");
                $("p").append('<p id="note">id wrong</p>');
            }
        });
    });
});