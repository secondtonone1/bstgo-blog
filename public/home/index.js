/*点击标题栏 */
$('.category>li>a').on('click', function() {
        $(this).parent().addClass("active").siblings().removeClass("active")
    })
    // 点击文章索引边栏的一级目录
$('.sidebar>li>a').on('click', function() {
    //隐藏所有二级折叠框
    $('.sidebar>li>div.collapse').collapse('hide')
        // 移除其他二级标题active属性
    $(".sub-sidebar > li").removeClass("active").siblings().removeClass("active")
        //将点击的一级标题设置active类，其他一级标题移除active
    $(this).parent().addClass("active").siblings().removeClass("active")
})

//点击文章索引边栏的二级目录
$('.sub-sidebar>li>a').on('click', function() {
    // 移除一级标题active类
    $(".sidebar > li").removeClass("active").siblings().removeClass("active")
        //设置二级标题active选中效果
    $(this).parent().addClass("active").siblings().removeClass("active")
})

//小屏幕下点击列表按钮
$('.navbar-toggle').on('click', function() {
    // 滚动到页面顶部
    $(window).scrollTop(0);
    //设置其他标题栏的弹出框折叠
    let collapseid = $(this).siblings('.navbar-toggle').attr('data-target')
    console.log(collapseid)
    $(collapseid).collapse('hide')
})