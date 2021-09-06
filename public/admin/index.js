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

$('#rt-index').on('click', function(e) {
    e.preventDefault()
    let year = $('.year-text').text()
    let month = $('.month-text').text()
    let cat = $('.category-text').text()
    let keywords = $('.keysearch-div>input').val()
    let condition = {}
    condition.year = year
    condition.month = month
    condition.cat = cat
    condition.keywords = keywords
    condition.page = 1
    let condJson = JSON.stringify(condition)
    console.log("condJson is ", condJson)

    $('#article-content').parent().fadeOut(100)
    $.ajax({
        type: "POST",
        url: "/admin/index",
        contentType: "application/json",
        data: condJson, //参数列表
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
            let total = $('#page-total').text()
            let cur = $('#page-cur').text()
            resetPage(cur, total)
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
    let catedata = $(this).text()
    console.log('post data is ', JSON.stringify({ 'category': catedata }))
    event.preventDefault(); //使a自带的方法失效，即无法调整到href中的URL()
    $('#article-content').parent().fadeOut(100, function() {
        $.ajax({
            type: "POST",
            url: "/admin/category",
            contentType: "application/json",
            data: JSON.stringify({ 'category': catedata }), //参数列表
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

    console.log("text_data is ", $(this).children('a').text())
    if ($(this).children('a').text() != "不限") {
        $('.month-div>button').attr('disabled', false)
    } else {
        $('.month-div>button').attr('disabled', true)
    }
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

//搜索文章
$('#article-content').on('click', '.searchbtn-div>button', function() {
    console.log('clicked article totoal search btn')
    let year = $('.year-text').text()
    let month = $('.month-text').text()
    let cat = $('.category-text').text()
    let keywords = $('.keysearch-div>input').val()
    let condition = {}
    condition.year = year
    condition.month = month
    condition.cat = cat
    condition.keywords = keywords
    condition.page = 1
    let condJson = JSON.stringify(condition)
    console.log("condJson is ", condJson)

    $('.blog-footer').fadeOut(700)
    $('.article-page').fadeOut(700)

    $('.article-div').fadeOut(700, function() {
        $.ajax({
            type: "POST",
            url: "/admin/articlesearch",
            contentType: "application/json",
            data: condJson, //参数列表
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
                $('.article-div').html(result)
                let total = $('#page-total').text()
                let cur = $('#page-cur').text()
                console.log('cur is ', cur)
                console.log('total is ', total)
                resetPage(cur, total)
                $('.article-div').fadeIn(700)
                $('.blog-footer').fadeIn(700)
                $('.article-page').fadeIn(700)
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

    })
})

//点击排序
$('#article-content').on('click', '.sort-edit-btn', function() {
    let categorytxt = $('.mini-li.active>span').text()
    $(this).parent().parent().parent().fadeOut(100, function() {
        $.ajax({
            type: "POST",
            url: "/admin/sort",
            contentType: "application/json",
            data: JSON.stringify({ 'category': categorytxt }), //参数列表
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
    let categorytxt = $('.mini-li.active>span').text()
    let subCat = {}
    subCat.sortlist = []
    subCat.subcat = categorytxt
    $('.sort-article-list').each(function(index, element) {
        console.log(index)
        console.log(element)
        console.log($(element).attr('articleid'))
        console.log($(element).children('span').text())
        let article = {
            'articleid': $(element).attr('articleid'),
            'title': $(element).children('span').text(),
            'index': index - 0 + 1
        }
        subCat.sortlist.push(article)
    })

    let jsonData = JSON.stringify(subCat)
    console.log('jsondata is ', jsonData)
    $('#article-content').parent().fadeOut(100, function() {
        $.ajax({
            type: "POST",
            url: "/admin/sortsave",
            contentType: "application/json",
            data: jsonData, //参数列表
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

    let sub_cat = $('li.mini-li.active').attr('subcatid')
    let cat = $('li.mini-li.active').parents('.sub-li').children('div').attr('id')
    console.log('cat is ', cat)
    console.log('sub_cat is ', sub_cat)

    window.location.href = '/admin/articledit?cat=' + cat + '&subcat=' + sub_cat

})

//点击创建分类标签
$('.cat-ctg').on('click', function() {
    console.log('sibling sub-li length is ', $(this).siblings('.sub-li').length)
    window.catnum = $(this).siblings('.sub-li').length
    console.log(window.catnum)
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
        data: JSON.stringify({ "category": ctg_val, "index": window.catnum }), //参数列表
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
        data: JSON.stringify({ "subcategory": ctg_val, "parentid": window.cur_subctg.parent().parent().attr('id'), "index": window.subcatnum }), //参数列表
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
    console.log('mini-li length is ', $(this).siblings('.mini-li').length)
    window.subcatnum = $(this).siblings('.mini-li').length
})

$(function() {
    $('.month-div>button').attr('disabled', true)
    let total = $('#page-total').text()
    let cur = $('#page-cur').text()
    resetPage(cur, total)
})

//点击文章删除span
$('#article-content').on('click', '.del-span', function() {
    let tips = $(this).parent().siblings('.article-title').text()
    $('#deldialog .del-article-span').text(tips)
    $('#deldialog').modal('show')

})

//点击删除模态框提交
$('#deldialog  .sure').on('click', function() {
    let tips = $('#deldialog .del-article-span').text()
    let delSuccess = 0;
    let inputdata = $("#deldialog input").val()
    if (inputdata != tips) {
        console.log('input data is ', inputdata)
        console.log('tips is ', tips)
        console.log('del data failed, input error!')
        $('#deldialog').modal('hide')
        $('.error-tips').text('del data failed, input error!')
        $('.error-tips').fadeIn(1000, function() {
            $('.error-tips').fadeOut(2000)
        })
        return
    }
    //发送删除给后端
    $.ajax({
        type: "POST",
        url: "/admin/delarticle",
        contentType: "application/json",
        data: JSON.stringify({ "title": tips }), //参数列表
        dataType: "json",
        success: function(result) {
            //请求正确之后的操作
            console.log('post success , result is ', result)
            delSuccess = 1
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

    $('#deldialog').modal('hide')

    if ($('#rt-index').hasClass('active')) {

        let year = $('.year-text').text()
        let month = $('.month-text').text()
        let cat = $('.category-text').text()
        let keywords = $('.keysearch-div>input').val()
        let condition = {}
        condition.year = year
        condition.month = month
        condition.cat = cat
        condition.keywords = keywords
        condition.page = 1
        let condJson = JSON.stringify(condition)
        console.log("condJson is ", condJson)

        $('#article-content').parent().fadeOut(50, function() {
            $.ajax({
                type: "POST",
                url: "/admin/index",
                contentType: "application/json",
                data: condJson, //参数列表
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
                    console.log($('#article-content').parent())
                    let total = $('#page-total').text()
                    let cur = $('#page-cur').text()
                    resetPage(cur, total)
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
    } else if ($('.mini-li.active').length > 0) {
        let catedata = $('.mini-li.active>span').text()
        console.log('post data is ', JSON.stringify({ 'category': catedata }))

        $('#article-content').parent().fadeOut(50, function() {
            $.ajax({
                type: "POST",
                url: "/admin/category",
                contentType: "application/json",
                data: JSON.stringify({ 'category': catedata }), //参数列表
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
    } else {
        window.location.href = "/admin"
    }

})

function pageli(num) {
    return '<li class="page-li" id=\"page-' + num + '\"><a href="#">' + num + '</a></li>'
}

function pageliActive(num) {
    return '<li class="active page-li" id=\"page-' + num + '\"><a href="#">' + num + '</a></li>'
}

//翻页插件设置页码
function resetPage(cur, total) {
    $('.article-page .pagination').html('')
    console.log('cur page is ', cur)
    console.log('total page is ', total)
    if (cur <= 0 || total <= 0) {
        return
    }
    if ($('.article-list').length <= 0) {
        return
    }
    window.curpage = cur - 0
    window.totalpage = total - 0
    let prev = '<li class="page-prev"> <a href="#" aria-label="Previous">'
    prev += '<span aria-hidden="true">&laquo;</span></a></li>'

    let next = '<li class="page-next"><a href="#" aria-label="Next">'
    next += '<span aria-hidden="true">&raquo;</span></a></li>'

    let ellipsis_prev = '<li id="ellipsis-prev"><span >...</span></li>'
    let ellipsis_next = '<li id="ellipsis-next"><span >...</span></li>'

    $('.article-page .pagination').append(pageliActive(cur))

    //当前页为第一页
    if (cur <= 1) {}

    //当前页为第二页

    if (cur == 2) {
        $(pageli(cur - 1)).insertBefore("#page-" + cur)
        $(prev).insertBefore("#page-" + (cur - 1))
    }

    //当前页为第三页
    if (cur == 3) {
        $(pageli(cur - 1)).insertBefore("#page-" + cur)
        $(pageli(cur - 2)).insertBefore("#page-" + (cur - 1))
        $(prev).insertBefore("#page-" + (cur - 2))
    }

    //当前页码大于3
    if (cur > 3) {
        $(pageli(cur - 1)).insertBefore("#page-" + cur)
        $(ellipsis_prev).insertBefore("#page-" + (cur - 1))
        $(pageli(1)).insertBefore("#ellipsis-prev")
        $(prev).insertBefore("#page-1")
    }

    if (cur >= total) {

    }

    if (cur == total - 1) {
        $(pageli(cur - 0 + 1)).insertAfter("#page-" + cur); //在id为test的元素后插入<p>测试</p>
        $(next).insertAfter("#page-" + (cur - 0 + 1))
    }

    if (cur == total - 2) {
        $(pageli(cur - 0 + 1)).insertAfter("#page-" + cur); //在id为test的元素后插入<p>测试</p>
        $(pageli(cur - 0 + 2)).insertAfter("#page-" + (cur - 0 + 1))
        $(next).insertAfter("#page-" + (cur - 0 + 2))
    }

    if (cur < total - 2) {
        $(pageli(cur - 0 + 1)).insertAfter("#page-" + cur); //在id为test的元素后插入<p>测试</p>
        $(ellipsis_next).insertAfter("#page-" + (cur - 0 + 1))
        $(pageli(total)).insertAfter("#ellipsis-next")
        $(next).insertAfter("#page-" + total)
    }

}

//点击页码
$('#article-content').on('click', '.page-li', function(e) {
    e.preventDefault()
    page = $(this).children('a').text()
    console.log("click page is ", page)
    let year = $('.year-text').text()
    let month = $('.month-text').text()
    let cat = $('.category-text').text()
    let keywords = $('.keysearch-div>input').val()
    let condition = {}
    condition.year = year
    condition.month = month
    condition.cat = cat
    condition.keywords = keywords
    condition.page = page - 0
    let condJson = JSON.stringify(condition)
    console.log("condJson is ", condJson)

    //请求page页数据给后端发送ajax请求
    $('#article-content').parent().fadeOut(100, function() {
        $.ajax({
            type: "POST",
            url: $('.urlinfo').text(),
            contentType: "application/json",
            data: condJson, //参数列表
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
                let total = $('#page-total').text()
                let cur = $('#page-cur').text()
                resetPage(cur, total)
                $('#article-content').parent().fadeIn(500)
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

//点击next翻页操作
$('#article-content').on('click', '.page-next', function(e) {
    e.preventDefault()
    let next_page = window.curpage - 0 + 1
        //请求page页数据给后端发送ajax请求
    let total_page = window.totalpage - 0
    console.log("next page is ", next_page)
    console.log("total_page page is ", total_page)
    if (next_page > total_page) {
        return
    }

    let year = $('.year-text').text()
    let month = $('.month-text').text()
    let cat = $('.category-text').text()
    let keywords = $('.keysearch-div>input').val()
    let condition = {}
    condition.year = year
    condition.month = month
    condition.cat = cat
    condition.keywords = keywords
    condition.page = next_page - 0
    let condJson = JSON.stringify(condition)
    console.log("condJson is ", condJson)

    //发送请求next_page页数据
    //请求page页数据给后端发送ajax请求
    $('#article-content').parent().fadeOut(100, function() {
        $.ajax({
            type: "POST",
            url: $('.urlinfo').text(),
            contentType: "application/json",
            data: condJson, //参数列表
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
                let total = $('#page-total').text()
                let cur = $('#page-cur').text()
                resetPage(cur, total)
                $('#article-content').parent().fadeIn(500)
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

//点击prev翻页操作
$('#article-content').on('click', '.page-prev', function(e) {
    e.preventDefault()
    let prev_page = window.curpage - 1
        //请求page页数据给后端发送ajax请求
    console.log("prev_page is ", prev_page)
    if (prev_page < 1) {
        return
    }

    let year = $('.year-text').text()
    let month = $('.month-text').text()
    let cat = $('.category-text').text()
    let keywords = $('.keysearch-div>input').val()
    let condition = {}
    condition.year = year
    condition.month = month
    condition.cat = cat
    condition.keywords = keywords
    condition.page = prev_page - 0
    let condJson = JSON.stringify(condition)
    console.log("condJson is ", condJson)


    //发送请求prev_page页数据
    //请求page页数据给后端发送ajax请求
    $('#article-content').parent().fadeOut(100, function() {
        $.ajax({
            type: "POST",
            url: $('.urlinfo').text(),
            contentType: "application/json",
            data: condJson, //参数列表
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
                let total = $('#page-total').text()
                let cur = $('#page-cur').text()
                resetPage(cur, total)
                $('#article-content').parent().fadeIn(500)
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

//编辑文章
$('#article-content').on('click', '.edit-span', function() {
    let id = $(this).parents(".article-ele").attr("id")
    window.location.href = "/admin/articlemodify?id=" + id
})

//点击回收站
$('#draftbox').on('click', function() {

})