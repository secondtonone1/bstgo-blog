/*点击标题栏 */
$('.category>li>a').on('click', function() {
        $(this).parent().addClass("active").siblings().removeClass("active")
    })
    // 点击文章索引边栏的一级目录
$('.sidebar>li>a').on('click', function() {
    //隐藏所有二级折叠框
    $('.sidebar>li>div.collapse').collapse('hide')
        // 移除其他二级标题active属性
    $(".sub-sidebar > li").removeClass("active")
        //将点击的一级标题设置active类，其他一级标题移除active
    $(this).parent().addClass("active").siblings().removeClass("active")
        // 将其他一级目录图标切换为初始化右箭头
        //console.log($(this).parent().siblings().children('a').children('span'))
    $(this).parent().siblings().children('a').children('span').removeClass('glyphicon glyphicon-menu-down').removeClass('glyphicon glyphicon-menu-right').addClass('glyphicon glyphicon-menu-right')

    //将自己的一级目录图标切换
    if ($(this).children('span').attr('class').trim() == 'glyphicon glyphicon-menu-right') {
        $(this).children('span').removeClass('glyphicon glyphicon-menu-right')
        $(this).children('span').addClass('glyphicon glyphicon-menu-down')
    } else {
        $(this).children('span').removeClass('glyphicon glyphicon-menu-down')
        $(this).children('span').addClass('glyphicon glyphicon-menu-right')
    }
})

//点击文章索引边栏的二级目录
$('.sub-sidebar>li>a').on('click', function() {
    // 移除一级标题active类
    $(".sidebar > li").removeClass("active").siblings().removeClass("active")
        //设置二级标题active选中效果
    $(this).parent().addClass("active").siblings().removeClass("active")
})

//小屏幕下点击列表按钮
$('.article-index-btn').on('click', function() {
    $('.ancestor').addClass('show')
    $('body,html').addClass('fix-html');
})

//蒙板点击后侧边栏弹回
$('.mask').on('click', function() {
    $('body,html').removeClass('fix-html');
    $('.ancestor').removeClass('show')
})

//小屏幕侧边栏点击一级标题
$('.xs-article-index-wrapper>ul>li>a').on('click', function() {
    //折叠其他一级标题
    $(this).parent().siblings().children().collapse('hide')
        //将点击的一级标题设置active类，其他一级标题移除active
    $(this).parent().addClass("active").siblings().removeClass("active")
        //设置二级标题取消active
    $('.xs-article-index-wrapper>ul>li>div a').parent().removeClass('active')
})

//小屏幕侧边栏点击二级标题
$('.xs-article-index-wrapper>ul>li>div a').on('click', function() {
    //移除一级标题选中状态
    $('.xs-article-index-wrapper>ul>li>a').parent().removeClass('active')
        //移除其他二级标题active
    $('.xs-article-index-wrapper>ul>li>div a').parent().removeClass('active')
        //设置点击的二级标题active
    $(this).parent().addClass('active')
})

//按钮点击事件
$('.comment-commit-btn').on('click', function() {
    if (window.editor.txt.html().trim() == "") {
        return
    }
    var comment_data = {
        "comment_name": "江边皓月",
        "comment_head": "https://profile.csdnimg.cn/2/D/2/3_akiss33",
        "comment_date": ' ' + formatDate(new Date().getTime()),
        "comment_content": window.editor.txt.html(),
    }
    var html = template('comment-li-tpl', comment_data)
    $(".comment-list-ul").prepend(html)
    window.editor.txt.clear()
    return
})

//点击评论数
$('#comment-num').on('click', function() {
    console.log("comment-num clicked")
    $('html , body').animate({ scrollTop: $('#comment-label-id').offset().top - 200 }, 300);
    window.editor.txt.html("")
})

//点击回复数开启回复功能
$('.comment-list-ul').on('click', '.reply-num', function(e) {
    $(this).parent().parent().siblings('.reply-text').stop().slideToggle()
    $(this).parent().parent().siblings('.reply-text').find('textarea').focus()
})

// $('.comment-list-ul').on('blur', '.reply-text textarea', function(e) {
//     $(this).parent().parent().parent(".reply-text").stop().slideUp();
// })

$(function() {
    const E = window.wangEditor
    const editor = new E('.comment-editor')
        //    设置编辑区域高度为 500 px
    editor.config.height = 100
        //配置菜单
    editor.config.menus = [
        'head',
        'bold',
        // 'fontSize',
        // 'fontName',
        'italic',
        //'underline',
        //'strikeThrough',
        // 'indent',
        //'lineHeight',
        //'foreColor',
        //'backColor',
        'link',
        'list',
        //'todo',
        // 'justify',
        'quote',
        'emoticon',
        //'image',
        //'video',
        //'table',
        'code',
        'splitLine',
        'undo',
        'redo',
    ]
    hljs.initHighlightingOnLoad(); // 初始化
    hljs.initLineNumbersOnLoad();
    editor.highlight = hljs
    editor.config.languageTab = '    '
    editor.config.pasteIgnoreImg = false
    editor.config.uploadImgShowBase64 = true
    editor.create()
    window.editor = editor
})