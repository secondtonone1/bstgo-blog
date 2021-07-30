// 点击侧边栏一级标题

$(".sidebar > ul > li > a").on('click', function() {
    console.log($(this).attr('href'))
        //隐藏所有二级折叠框
    $('.sidebar li>div.collapse').collapse('hide')

    let collapes_id = $(this).attr('href')
        // 设置右箭头为向下箭头
        //将自己的一级目录图标切换
    if ($(this).children('span').attr('class') != undefined) {
        if ($(this).children('span').attr('class').trim().indexOf('glyphicon-menu-right') != -1) {
            $(this).children('span').removeClass('glyphicon glyphicon-menu-right')
            $(this).children('span').addClass('glyphicon glyphicon-menu-down')
        } else {
            $(this).children('span').removeClass('glyphicon glyphicon-menu-down')
            $(this).children('span').addClass('glyphicon glyphicon-menu-right')
        }
    }


    //设置点击的一级标题为选中状态，其他一级标题取消选中
    $(this).parent('li').addClass('active').siblings('li').removeClass('active')

    // 移除其他二级标题active属性
    $(".sub-sidebar > li").removeClass("active")
        //设置二级标题箭头恢复为向右
    $(".sub-sidebar > li >a>span").removeClass('glyphicon glyphicon-menu-down').addClass('glyphicon glyphicon-menu-right')
        // 将其他一级目录图标切换为初始化右箭头
        //console.log($(this).parent().siblings().children('a').children('span'))
    $(this).parent().siblings('li').children('a').children('span').removeClass('glyphicon glyphicon-menu-down').removeClass('glyphicon glyphicon-menu-right').addClass('glyphicon glyphicon-menu-right')


})

$('#rt-index').on('click', function() {
    $('#article-content').parent().fadeOut(100, function() {
        $.ajax({
            type: "POST",
            url: "/admin/index",
            contentType: "application/json",
            data: JSON.stringify({ category: $(this).text() }), //参数列表
            dataType: "html",
            success: function(result) {
                //请求正确之后的操作
                //console.log('post success , result is ', result)
                $('#article-content').html(result)
                console.log($('#article-content').parent())
                $('#article-content').parent().fadeIn(700)
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
        });
    })
})


//点击文章索引边栏的二级目录
$('.sub-sidebar>li>a').on('click', function() {
    console.log('二级标题点击了')
        // 移除一级标题active类
    $(".sidebar>ul>li").removeClass("active")
        // 移除三级标题active类
    $('.mini-sidebar>li').removeClass('active')
        //设置二级标题active选中效果
    $(this).parent().addClass("active").siblings().removeClass("active")

    // 设置右箭头为向下箭头
    //将自己的一级目录图标切换
    if ($(this).children('span').attr('class') != undefined) {
        if ($(this).children('span').attr('class').trim().indexOf('glyphicon-menu-right') != -1) {
            console.log('切换为下箭头')
            $(this).children('span').removeClass('glyphicon glyphicon-menu-right')
            $(this).children('span').addClass('glyphicon glyphicon-menu-down')
        } else {
            console.log('切换为右箭头')
            $(this).children('span').removeClass('glyphicon glyphicon-menu-down')
            $(this).children('span').addClass('glyphicon glyphicon-menu-right')
        }
    }

})

//点击文章索引边栏的三级目录
$('.sub-sidebar').on('click', '.mini-li', function(event) {
    console.log('点击了三级标题')
        // 移除一级标题active类
    $(".sidebar>ul>li").removeClass("active")
        // 移除二级标题active类
    $('.sub-sidebar>li').removeClass('active')
        //设置自己三级选中效果，清除其他三级选中效果
    $(this).parent().addClass("active").siblings().removeClass('active')

    event.preventDefault(); //使a自带的方法失效，即无法调整到href中的URL()
    $('#article-content').parent().fadeOut(100, function() {
        $.ajax({
            type: "POST",
            url: "/admin/category",
            contentType: "application/json",
            data: JSON.stringify({ category: $(this).text() }), //参数列表
            dataType: "html",
            success: function(result) {
                //请求正确之后的操作
                // console.log('post success , result is ', result)
                $('#article-content').html(result)
                $('#article-content').parent().fadeIn(700)
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
        });
    })

})



//点击年份下拉菜单
$('.year-div button').on('click', function() {
    //console.log("year button clicked")
    //设置下拉列表选中状态
    let text_data = $(this).children('.year-text').text()
        //清空之前的选中状态
    $(this).siblings('ul').children('li').removeClass('active')
        //设置下拉列表选中状态

    $(this).siblings('ul').children('li').each(function() {
        if ($(this).children('a').text() == text_data) {
            $(this).addClass('active')
        }
    })
})

//点击年份列表更新按钮信息
$('.year-div ul li').on('click', function() {
    $(this).addClass('active').siblings().removeClass('active')
    $(this).parent().siblings('button').children('.year-text').text($(this).children('a').text())
})

//点击月份下拉菜单
$('.month-div button').on('click', function() {
    //console.log("year button clicked")
    //设置下拉列表选中状态
    let text_data = $(this).children('.month-text').text()
        //清空之前的选中状态
    $(this).siblings('ul').children('li').removeClass('active')
        //设置下拉列表选中状态

    $(this).siblings('ul').children('li').each(function() {
        if ($(this).children('a').text() == text_data) {
            $(this).addClass('active')
        }
    })
})

//点击月份列表更新按钮信息
$('.month-div ul li').on('click', function() {
    $(this).addClass('active').siblings().removeClass('active')
    $(this).parent().siblings('button').children('.month-text').text($(this).children('a').text())
})

//点击月份下拉菜单
$('.category-div button').on('click', function() {
    //console.log("year button clicked")
    //设置下拉列表选中状态
    let text_data = $(this).children('.category-text').text()
        //清空之前的选中状态
    $(this).siblings('ul').children('li').removeClass('active')
        //设置下拉列表选中状态

    $(this).siblings('ul').children('li').each(function() {
        if ($(this).children('a').text() == text_data) {
            $(this).addClass('active')
        }
    })
})

//点击月份列表更新按钮信息
$('.category-div ul li').on('click', function() {
    $(this).addClass('active').siblings().removeClass('active')
    $(this).parent().siblings('button').children('.category-text').text($(this).children('a').text())
})

$('#article-content').on('click', '.sort-edit-btn', function() {

    $(this).parent().parent().parent().fadeOut(100, function() {
        $.ajax({
            type: "POST",
            url: "/admin/sort",
            contentType: "application/json",
            data: JSON.stringify({ category: $(this).text() }), //参数列表
            dataType: "html",
            success: function(result) {
                //请求正确之后的操作
                console.log('post success , result is ', result)
                $('#article-content').html(result)
                var article_items = document.getElementsByClassName('sort-article-list');
                [].forEach.call(article_items, actions_article);
                $('#article-content').parent().fadeIn(700)

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
        });
    })

})

$('#article-content').on('click', '.sort-save-btn', function() {
    $('#article-content').parent().fadeOut(100, function() {
        $.ajax({
            type: "POST",
            url: "/admin/sortsave",
            contentType: "application/json",
            data: JSON.stringify({ category: $(this).text() }), //参数列表
            dataType: "html",
            success: function(result) {
                //请求正确之后的操作
                console.log('post success , result is ', result)
                $('#article-content').html(result)
                $('#article-content').parent().fadeIn(1000)
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
        });
    })
})