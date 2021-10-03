$(function(e) {
    let content = $('.article-content').text()
        //console.log("content is ", content)
    window.maineditor.txt.html(content)
    window.maineditor.txt.append('请继续编辑...')
    let title = $('.article-title').text()
    $('#title-edit').val(title)
    let subtitle = $('.article-subtitle').text()
    $('.sub-tittle input').val(subtitle)
})

$('.update-btn').on('click', function(e) {
    loadinst.loading()
    let id = $('.article-id').text()
    console.log('article id is ', id)
    let data = {}
    data['title'] = $('#title-edit').val()
    data['subtitle'] = $('.sub-tittle input').val()
    data['id'] = id
    data['cat'] = $('.cat-text').text()
    data['subcat'] = $('.subcat-text').text()
    data['lastedit'] = Math.round(new Date().valueOf() / 1000)
    data['author'] = $('.author input').val()
    data['content'] = window.maineditor.txt.html()
    console.log('update data is ', data)
    let data_json = JSON.stringify(data)
    console.log('data_json is ', data_json)

    $.ajax({
        type: "POST",
        url: "/admin/updatearticle",
        contentType: "application/json",
        data: data_json, //参数列表
        dataType: "json",
        success: function(result) {
            loadinst.unloading()
                //请求正确之后的操作
            console.log('post success , result is ', result)
            if (result.code != 0) {
                $('.msg-tip').text(result.msg).fadeIn(1000, function(e) {
                    $('.msg-tip').fadeOut(1000)
                })
                return
            }

            window.location.href = "/admin"
        },
        error: function(XMLHttpRequest, textStatus, errorThrown) {
            //请求失败之后的操作
            loadinst.unloading()
            console.log('post failed')
                // 状态码
            console.log(XMLHttpRequest.status);
            // 状态
            console.log(XMLHttpRequest.readyState);
            // 错误信息   
            console.log(textStatus);
        }
    })
})