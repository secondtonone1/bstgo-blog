function loginAdmin() {
    console.log('login admin submit')
    var myDate = new Date(); //当前时间
    let email = $('#inputEmail3').val()
    let pwd = $('#inputPassword3').val()
    let localTime = myDate.toLocaleTimeString()
    console.log('localTime is ', localTime)
        //sha1加密
    let pwdsecret = sha1(pwd + localTime)
    console.log('pwd secret is ', pwdsecret)
    $.ajax({
        type: "POST",
        url: "/admin/loginsub",
        contentType: "application/json",
        data: JSON.stringify({
            'email': email,
            'pwd': pwdsecret,
            'salt': localTime,
        }), //参数列表
        dataType: "json",
        success: function(result) {
            //请求正确之后的操作
            console.log('post success , result is ', result)
            console.log('post success , result code is ', result.code)

            if (result.code == 1006) {
                $('.alert-danger').text("登陆失败次数过多，一分钟后再尝试")
                $('.alert-danger').stop().fadeIn(1000, function() {
                    $(this).fadeOut(2000)
                })
                return
            }

            if (result.code != 0) {
                $('.alert-danger').text("邮箱或者密码不正确")
                $('.alert-danger').stop().fadeIn(1000, function() {
                    $(this).fadeOut(2000)
                })
                return
            }

            window.location.href = "/admin/"
            return
        },
        error: function(XMLHttpRequest, textStatus, errorThrown) {
            //请求失败之后的操作
            console.log('post failed')
                // 状态码
            console.log(XMLHttpRequest.status);
            // 状态
            console.log(XMLHttpRequest.readyState);
            // 错误信息   
            console.log(textStatus);
        }
    })
    return false
}