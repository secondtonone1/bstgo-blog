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

                let index_find = result.indexOf('Sign in')
                    //找到Sign in 说明是登录页面
                if (index_find != -1) {
                    window.location.href = "/admin/login"
                    return
                }

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
$('.sub-sidebar').on('click', '.sub-li>a', function() {
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
    console.log('二级标题点完了')
})

//点击文章索引边栏的三级目录
$('.sub-sidebar').on('click', '.mini-li>span', function(event) {
    console.log('点击了三级标题')
        // 移除一级标题active类
    $(".sidebar>ul>li").removeClass("active")
        // 移除二级标题active类
    $('.sub-sidebar>li').removeClass('active')
        //设置自己三级选中效果，清除其他三级选中效果
    $('.mini-li').removeClass('active')
    $(this).parent().addClass("active")

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

                let index_find = result.indexOf('Sign in')
                    //找到Sign in 说明是登录页面
                if (index_find != -1) {
                    window.location.href = "/admin/login"
                    return
                }

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

//点击分类下拉菜单
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

//点击排序
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
                let index_find = result.indexOf('Sign in')
                    //找到Sign in 说明是登录页面
                if (index_find != -1) {
                    window.location.href = "/admin/login"
                    return
                }
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


//点击保存
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
                let index_find = result.indexOf('Sign in')
                    //找到Sign in 说明是登录页面
                if (index_find != -1) {
                    window.location.href = "/admin/login"
                    return
                }

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

//发布文章
$('#article-content').on('click', '.article-pub-btn', function() {

    let sub_cat = $('li.mini-li.active').children('span').text()
    let cat = $('li.mini-li.active').parents('.sub-li').children('a').text()
    console.log('cat is ', cat)
    console.log('sub_cat is ', sub_cat)

    window.location.href = '/admin/articledit?cat=' + cat + '&subcat=' + sub_cat

})


// 创建分类
$('#myModalDialog .sure').on('click', function() {
    let ctg_val = $(this).parent().siblings('.model-input').children('input').val()
    console.log(ctg_val);
    $('#myModalDialog').modal('hide')
    $.ajax({
        type: "POST",
        url: "/admin/createctg",
        contentType: "application/json",
        data: JSON.stringify({ "category": ctg_val }), //参数列表
        dataType: "html",
        success: function(result) {
            //请求正确之后的操作
            console.log('post success , result is ', result)

            let index_find = result.indexOf('Sign in')
                //找到Sign in 说明是登录页面
            if (index_find != -1) {
                window.location.href = "/admin/login"
                return
            }

            let res_find = result.indexOf('res-success')
            if (res_find == -1) {
                return
            }

            $('.sub-sidebar').append(result)
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

//点击确定按钮，提交子分类
$('#subCtgDialog  .sure').on('click', function() {

    if (window.cur_subctg == undefined || window.cur_subctg == null) {
        return
    }

    let ctg_val = $(this).parent().siblings('.model-input').children('input').val()
    console.log(ctg_val);
    console.log(window.cur_subctg.parent().parent().attr('id'))
    $('#subCtgDialog').modal('hide')
    $.ajax({
        type: "POST",
        url: "/admin/createsubctg",
        contentType: "application/json",
        data: JSON.stringify({ "subcategory": ctg_val, "parentid": window.cur_subctg.parent().parent().attr('id') }), //参数列表
        dataType: "html",
        success: function(result) {
            //请求正确之后的操作
            console.log('post success , result is ', result)
            console.log(window.cur_subctg)

            let index_find = result.indexOf('Sign in')
                //找到Sign in 说明是登录页面
            if (index_find != -1) {
                window.location.href = "/admin/login"
                return
            }

            //页面结果信息不成功则不添加
            let res_find = result.indexOf('res-success')
            if (res_find == -1) {
                return
            }

            window.cur_subctg.parent().append(result)

            var items = window.cur_subctg[0].parentNode.getElementsByClassName('mini-li');
            //console.log('items are ', items);
            [].forEach.call(Array.prototype.slice.call(items, items.length - 1), actions);
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

//创建子分类，将当前节点保存起来
$('.sub-sidebar').on('click', '.sub-ctg', function() {
    window.cur_subctg = $(this)
    console.log(window.cur_subctg)
})