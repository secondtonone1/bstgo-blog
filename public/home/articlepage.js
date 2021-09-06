//相关推荐文章点击
$('.article-recommend>a').on('click', function(e) {
    let article_id = $(this).attr('article-id')
    console.log("recommend clicked, article id is ", article_id)
    if (article_id && article_id != "") {
        window.location.href = "/articlepage?id=" + article_id
    }
    demo.loading()
})