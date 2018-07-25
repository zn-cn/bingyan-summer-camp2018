$(document).ready(function () {
    $('.who').click(function () {
        value = $("input:radio[name='who']:checked").val();//注意作用域的问题，每一个回调函数一个作用域？
        if (value === "1") {
            $("#div1").remove();
            $("h1").after("<div id='div1'></div>");
            $("div").append("id:<input id='id' type='text' name='id'><br>");
            $("div").append("password:<input id='password' type='password' name='password'><br>");
            $("div").append("rewrite_password:<input id='rpassword' type='password' name='password'><br>");
            $("div").append("email:<input type='text' name='email'><br>");
            $("div").append("name:<input type='text' name='name'><br>");
            $("div").append("addr:<input type='text' name='addr'><br>");
            $("div").append("phone:<input type='text' name='phone'><br>");
            $('#rpassword').blur(function () {//注意异步加载的问题
                $("#message").remove();
                p1 = $('#password').val();
                p2 = $('#rpassword').val();
                if (p1 !== p2) {
                    $("#rpassword").after("<p id='message'>two password is not equal</p>");
                }
            });
        }
        else if (value === "2") {
            $("#div1").remove();
            $("h1").after("<div id='div1'></div>");
            $("div").append("id:<input id='id' type='text' name='id'><br>");
            $("div").append("password:<input id='password' type='password' name='password'><br>");
            $("div").append("rewrite_password:<input id='rpassword' type='password' name='password'><br>");
            $("div").append("email:<input type='text' name='email'><br>");
            $("div").append("addr:<input type='text' name='addr'><br>");
            $('#rpassword').blur(function () {//注意异步加载的问题
                $("#message").remove();
                p1 = $('#password').val();
                p2 = $('#rpassword').val();
                if (p1 !== p2) {
                    $("#rpassword").after("<p id='message'>two password is not equal</p>");
                }
            });
        }
        else if (value === "3") {
            $("#div1").remove();
            $("h1").after("<div id='div1'></div>");
            $("div").append("id:<input id='id' type='text' name='id'><br>");
            $("div").append("password:<input id='password' type='password' name='password'><br>");
            $("div").append("rewrite_password:<input id='rpassword' type='password' name='password'><br>");
            $("div").append("email:<input type='text' name='email'><br>");
            $('#rpassword').blur(function () {//注意异步加载的问题
                $("#message").remove();
                p1 = $('#password').val();
                p2 = $('#rpassword').val();
                if (p1 !== p2) {
                    $("#rpassword").after("<p id='message'>two password is not equal</p>");
                }
            });
        }
    });
});