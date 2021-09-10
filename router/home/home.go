package home

import (
	"bstgo-blog/model"
	mongocli "bstgo-blog/mongo"
	"html/template"
	"log"
	"math"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/segmentio/ksuid"
)

func GetArticleDeails(c *gin.Context) {
	req := &model.ArtdetailsReq{}
	detailR := &model.ArticleDetailsR{}
	detailR.Msg = model.MSG_SUCCESS
	defer func() {
		c.HTML(http.StatusOK, "home/articledetails.html", detailR)
	}()
	err := c.BindJSON(req)
	if err != nil {
		log.Println("json unpack failed, err is ", err)
		detailR.Msg = model.MSG_JSON_UNPACK
		return
	}

	//获取总页数
	count, err := mongocli.ArticleTotalCount()
	if err != nil {
		detailR.Msg = "get article total count failed"
		log.Println("get article total count failed, err is ", err)
		return
	}

	//获取当前页
	var page_f float64 = float64(count) / 5
	detailR.TotalPage = int(math.Ceil(page_f))
	if req.Page <= 1 {
		detailR.CurPage = 1
	} else if req.Page >= detailR.TotalPage {
		detailR.CurPage = detailR.TotalPage
	} else {
		detailR.CurPage = req.Page
	}

	detailR.NextPage = detailR.CurPage + 1
	detailR.PrevPage = detailR.CurPage - 1
	if detailR.NextPage >= detailR.TotalPage {
		detailR.NextPage = detailR.TotalPage
	}

	if detailR.PrevPage <= 0 {
		detailR.PrevPage = 1
	}

	// 获取当前页文章列表
	details, err := mongocli.GetArticleDetailsByPage(detailR.CurPage)
	if err != nil {
		detailR.Msg = "get article details by page failed"
		log.Println("get article total count failed, err is ", err)
		return
	}

	for _, detail := range details {
		articleR := &model.HomeArticleR{}
		lasttm := time.Unix(detail.LastEdit, 0)
		articleR.LastEdit = lasttm.Format("2006-01-02")
		articleR.Cat = detail.Cat
		articleR.Title = detail.Title
		articleR.Id = detail.ArticleInfo.Id
		content := TrimEmptyTag(detail.Content)
		//log.Println("content TrimEmptyTag is ", content)
		index := strings.Index(content, "/p>")
		if index == -1 {
			log.Println("not found match /p>")
			continue
		}
		fIndex := index + 3
		if fIndex >= len(content) {
			fIndex = len(content)
		}
		articleR.Content = template.HTML(content[:fIndex])
		detailR.IndexArticleList = append(detailR.IndexArticleList, articleR)
	}

}

func Home(c *gin.Context) {
	//c.String(http.StatusOK, "Hello World")
	homeIndex := &model.HomeIndexR{}
	val, b := c.Get("visitnum")
	if !b {
		log.Println("get visit num from midware failed")
		c.HTML(http.StatusOK, "home/errorpage.html", "get visitnum failed, after 2 seconds return to home")
		return
	} else {
		homeIndex.VisitNum = val.(int64)
	}

	//nav 标题栏cat 信息
	menus, err := mongocli.GetMenuListByParent("")
	if err != nil {
		c.HTML(http.StatusOK, "home/errorpage.html", "get menu list by parent failed")
		return
	}

	for _, menu := range menus {
		navCat := &model.NavCatR{}
		navCat.CatId = menu.CatId
		navCat.Index = menu.Index
		navCat.Name = menu.Name
		homeIndex.NavCatList = append(homeIndex.NavCatList, navCat)
	}

	hotarticles, err := mongocli.HotArticles()
	if err != nil {
		c.HTML(http.StatusOK, "home/errorpage.html", "get hot articles failed")
		return
	}

	for _, hot := range hotarticles {
		homeR := &model.HomeArticleR{}
		homeR.Id = hot.Id
		homeR.Title = hot.Title
		homeR.LoveNum = hot.LoveNum
		homeR.ScanNum = hot.ScanNum
		homeIndex.HotList = append(homeIndex.HotList, homeR)
	}

	newcomments, err := mongocli.GetNewComments()
	if err != nil {
		c.HTML(http.StatusOK, "home/errorpage.html", "get hot new comments failed")
		return
	}

	for _, newcomment := range newcomments {
		commentR := &model.CommentR{}
		commentR.Content = template.HTML(trimHtml(newcomment.Content))
		commentR.UserName = newcomment.UserName
		commentR.ArtId = newcomment.ArtId
		//log.Println("commentR.ArtId is ", newcomment.ArtId)
		info, err := mongocli.GetArticleInfo(commentR.ArtId)
		if err != nil {
			log.Println("get article ", commentR.ArtId, " info failed, err is ", err)
			continue
		}
		commentR.ArtTitle = info.Title

		homeIndex.CommentList = append(homeIndex.CommentList, commentR)
	}

	articles, err := mongocli.GetArticleDetailsByPage(1)
	for _, article := range articles {
		articleR := &model.HomeArticleR{}
		lasttm := time.Unix(article.LastEdit, 0)
		articleR.LastEdit = lasttm.Format("2006-01-02")
		articleR.Cat = article.Cat
		articleR.Title = article.Title
		articleR.Id = article.ArticleInfo.Id
		content := TrimEmptyTag(article.Content)
		//log.Println("content TrimEmptyTag is ", content)
		index := strings.Index(content, "/p>")
		if index == -1 {
			log.Println("not found match /p>")
			continue
		}
		fIndex := index + 3
		if fIndex >= len(content) {
			fIndex = len(content)
		}
		articleR.Content = template.HTML(content[:fIndex])
		homeIndex.IndexArticleList = append(homeIndex.IndexArticleList, articleR)
	}

	//获取总页数
	count, err := mongocli.ArticleTotalCount()
	if err != nil {
		c.HTML(http.StatusOK, "home/errorpage.html", "get article total count failed")
		log.Println("get article total count failed, err is ", err)
		return
	}

	//获取当前页
	var page_f float64 = float64(count) / 5
	homeIndex.TotalPage = int(math.Ceil(page_f))
	homeIndex.CurPage = 1

	homeIndex.NextPage = 2

	if homeIndex.NextPage >= homeIndex.TotalPage {
		homeIndex.NextPage = homeIndex.TotalPage
	}

	c.HTML(http.StatusOK, "home/index.html", homeIndex)
}

// 去除空标签

func TrimEmptyTag(src string) string {
	re, _ := regexp.Compile(`<p>[\s]+?</p>`)
	res := re.ReplaceAllString(src, "")
	re2, _ := regexp.Compile(`<h[1-9]>[\s]+?</h[1-9]>`)
	res = re2.ReplaceAllString(res, "")
	re3, _ := regexp.Compile(`<p></p>`)
	res = re3.ReplaceAllString(res, "")
	re4, _ := regexp.Compile(`<h[1-9]></h[1-9]>`)
	res = re4.ReplaceAllString(res, "")
	return res
}

func Category(c *gin.Context) {
	catid := c.Query("catid")
	log.Println("id is ", catid)
	if catid == "" {
		c.HTML(http.StatusOK, "home/errorpage.html", "invalid page request , id is null, after 2 seconds return to home")
		return
	}

	cateIndex := &model.HomeCategoryR{}
	cateIndex.NavCatList = []*model.NavCatR{}
	cateIndex.LeftCatList = []*model.LeftCatR{}
	cateIndex.ActiveId = catid
	cateIndex.CommentList = []*model.CommentR{}
	cateIndex.HotList = []*model.HomeArticleR{}
	cateIndex.Comments = []*model.CommentR{}
	//获取访问量
	val, b := c.Get("visitnum")
	if !b {
		log.Println("get visit num from midware failed")
		c.HTML(http.StatusOK, "home/errorpage.html", "get visitnum failed, after 2 seconds return to home")
		return
	} else {
		cateIndex.VisitNum = val.(int64)
	}

	//nav 标题栏cat 信息
	menus, err := mongocli.GetMenuListByParent("")
	if err != nil {
		c.HTML(http.StatusOK, "home/errorpage.html", "get menu list by parent failed")
		return
	}

	for _, menu := range menus {
		navCat := &model.NavCatR{}
		navCat.CatId = menu.CatId
		navCat.Index = menu.Index
		navCat.Name = menu.Name
		cateIndex.NavCatList = append(cateIndex.NavCatList, navCat)
	}

	//获取侧边栏菜单信息
	menulv1, err := mongocli.GetMenuById(catid)
	if err != nil {
		log.Println("get menu by id ", catid, " failed, err is ", err)
		c.HTML(http.StatusOK, "home/errorpage.html", "get menu by id failed, after 2 seconds return to home")
		return
	}

	cateIndex.CategoryName = menulv1.Name

	menulv2s, err := mongocli.GetMenuListByParent(catid)
	if err != nil {
		log.Println("get menu by parent ", menulv1.Name, " failed, err is ", err)
		c.HTML(http.StatusOK, "home/errorpage.html", "get menu by parent failed, after 2 seconds return to home")
		return
	}

	if len(menulv2s) == 0 {
		log.Println("menu ", menulv1.Name, " has no child menu ")
		//渲染Category分类主页
		c.HTML(http.StatusOK, "home/categary.html", cateIndex)
		return
	}

	//用map管理二级菜单
	menulv2map := make(map[string]*model.LeftCatR)
	for _, menulv2 := range menulv2s {
		leftCatR := &model.LeftCatR{}
		leftCatR.CatId = menulv2.CatId
		leftCatR.Name = menulv2.Name
		cateIndex.LeftCatList = append(cateIndex.LeftCatList, leftCatR)
		menulv2map[menulv2.Name] = leftCatR
	}

	catartinfos, err := mongocli.CatArtInfos(menulv1.Name)
	if err != nil {
		log.Println("get art info failed by cat , err is ", err)
		//渲染Category分类主页
		c.HTML(http.StatusOK, "home/errorpage.html", "get art info failed by cat, after 2 seconds return to home")
		return
	}

	for _, catartinfo := range catartinfos {
		leftCatR, ok := menulv2map[catartinfo.Subcat]
		if !ok {
			log.Println("key ", catartinfo.Subcat, " is not in menulv2 map")
			continue
		}
		infoR := &model.ArtInfoR{}
		infoR.ArtId = catartinfo.Id
		infoR.ArtSubTitle = catartinfo.Subtitle
		leftCatR.SubArticle = append(leftCatR.SubArticle, infoR)
	}

	//中间文章详情
	article, err := mongocli.GetFirstArtByCat(menulv1.Name, menulv2s[0].Name)

	if err != nil {
		log.Println("get first article by cat failed, err is ", err)
		c.HTML(http.StatusOK, "home/errorpage.html", "get article failed, after 2 seconds return to home")
		return
	}

	cateIndex.Author = article.Author
	cateIndex.Cat = article.Cat
	cateIndex.Content = template.HTML(article.Content)
	createtm := time.Unix(article.CreateAt, 0)
	cateIndex.CreateAt = createtm.Format("2006-01-02 15:04:05")

	lasttm := time.Unix(article.LastEdit, 0)
	cateIndex.LastEdit = lasttm.Format("2006-01-02 15:04:05")
	cateIndex.Id = article.ArticleInfo.Id
	cateIndex.Index = article.Index
	cateIndex.LoveNum = article.LoveNum
	cateIndex.ScanNum = article.ScanNum
	cateIndex.Subcat = article.Subcat
	cateIndex.Subtitle = article.Subtitle
	cateIndex.Title = article.Title

	//获取评论信息
	comments, err := mongocli.GetCommentByParent(article.ArticleInfo.Id)
	if err != nil {
		c.HTML(http.StatusOK, "home/errorpage.html", "get comments failed")
		return
	}
	cateIndex.CommentNum = len(comments)
	for _, comment := range comments {
		tm := time.Unix(int64(comment.Time), 0)
		timestr := tm.Format("2006-01-02 15:04:05")

		commentR := &model.CommentR{}
		commentR.Time = timestr
		commentR.Content = template.HTML(comment.Content)
		commentR.Id = comment.Id
		commentR.LoveNum = comment.LoveNum
		commentR.Parent = comment.Parent
		commentR.Replys = []*model.ReplyR{}
		commentR.UserName = comment.UserName
		commentR.HeadIcon = comment.HeadIcon

		replys, err := mongocli.GetCommentByParent(comment.Id)
		if err != nil {
			log.Println("get reply by comment id ", comment.Id, "failed, error is ", err)
			continue
		}

		commentR.ReplyNum = len(replys)

		for _, reply := range replys {
			replyR := &model.ReplyR{}
			replyR.Content = template.HTML(reply.Content)
			replyR.Id = reply.Id
			replyR.LoveNum = reply.LoveNum
			replyR.Parent = reply.Parent
			rtm := time.Unix(int64(reply.Time), 0)
			timestr := rtm.Format("2006-01-02 15:04:05")
			replyR.Time = timestr
			replyR.UserName = reply.UserName
			replyR.HeadIcon = reply.HeadIcon
			commentR.Replys = append(commentR.Replys, replyR)
		}

		cateIndex.Comments = append(cateIndex.Comments, commentR)
	}

	hotarticles, err := mongocli.HotArticles()
	if err != nil {
		c.HTML(http.StatusOK, "home/errorpage.html", "get hot articles failed")
		return
	}

	for _, hot := range hotarticles {
		homeR := &model.HomeArticleR{}
		homeR.Id = hot.Id
		homeR.Title = hot.Title
		homeR.LoveNum = hot.LoveNum
		homeR.ScanNum = hot.ScanNum
		cateIndex.HotList = append(cateIndex.HotList, homeR)
	}

	newcomments, err := mongocli.GetNewComments()
	if err != nil {
		c.HTML(http.StatusOK, "home/errorpage.html", "get hot new comments failed")
		return
	}

	for _, newcomment := range newcomments {
		commentR := &model.CommentR{}
		commentR.Content = template.HTML(trimHtml(newcomment.Content))
		commentR.UserName = newcomment.UserName
		commentR.ArtId = newcomment.ArtId
		//log.Println("commentR.ArtId is ", newcomment.ArtId)
		info, err := mongocli.GetArticleInfo(commentR.ArtId)
		if err != nil {
			log.Println("get article ", commentR.ArtId, " info failed, err is ", err)
			continue
		}
		commentR.ArtTitle = info.Title

		cateIndex.CommentList = append(cateIndex.CommentList, commentR)
	}

	//渲染Category分类主页
	c.HTML(http.StatusOK, "home/categary.html", cateIndex)
}

func trimHtml(src string) string {
	//将HTML标签全转换成小写
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllStringFunc(src, strings.ToLower)
	//去除STYLE
	re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	src = re.ReplaceAllString(src, "")
	//去除SCRIPT
	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	src = re.ReplaceAllString(src, "")
	//去除所有尖括号内的HTML代码，并换成换行符
	re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllString(src, "\n")
	//去除连续的换行符
	re, _ = regexp.Compile("\\s{2,}")
	src = re.ReplaceAllString(src, "\n")
	return strings.TrimSpace(src)
}

//请求文章信息
func ArticlePage(c *gin.Context) {

	id := c.Query("id")
	log.Println("id is ", id)
	if id == "" {
		c.HTML(http.StatusOK, "home/errorpage.html", "invalid page request , id is null, after 2 seconds return to home")
		return
	}

	num, bres := c.Get("visitnum")
	if bres == false {
		log.Println("get visitnum failed , err is ")
		c.HTML(http.StatusOK, "home/errorpage.html", "get visitnum failed, after 2 seconds return to home")
		return
	}

	//内容区文章信息和内容
	err := mongocli.AddArticleScan(id)
	if err != nil {
		log.Println("add article scan num failed , error is ", err)
	}

	article, err := mongocli.GetArticleId(id)

	if err != nil {
		c.HTML(http.StatusOK, "home/errorpage.html", "get article failed, after 2 seconds return to home")
		return
	}

	//左侧相关推荐
	recommends, err := mongocli.RelRecommend(article.Cat)
	if err != nil {
		c.HTML(http.StatusOK, "home/errorpage.html", "get recommends failed")
		return
	}

	articleR := &model.ArticlePageR{}
	articleR.Author = article.Author
	articleR.Cat = article.Cat
	articleR.Content = template.HTML(article.Content)
	createtm := time.Unix(article.CreateAt, 0)
	articleR.CreateAt = createtm.Format("2006-01-02 15:04:05")

	lasttm := time.Unix(article.LastEdit, 0)
	articleR.LastEdit = lasttm.Format("2006-01-02 15:04:05")
	articleR.Id = article.ArticleInfo.Id
	articleR.Index = article.Index
	articleR.LoveNum = article.LoveNum
	articleR.ScanNum = article.ScanNum
	articleR.Subcat = article.Subcat
	articleR.Subtitle = article.Subtitle
	articleR.Title = article.Title
	articleR.Comments = []*model.CommentR{}
	articleR.VisitNum = num.(int64)
	articleR.RecommendList = []*model.HomeArticleR{}
	articleR.CommentList = []*model.CommentR{}
	articleR.HotList = []*model.HomeArticleR{}
	articleR.NavCatList = []*model.NavCatR{}

	//获取评论信息
	comments, err := mongocli.GetCommentByParent(article.ArticleInfo.Id)
	if err != nil {
		c.HTML(http.StatusOK, "home/errorpage.html", "get comments failed")
		return
	}
	articleR.CommentNum = len(comments)
	for _, comment := range comments {
		tm := time.Unix(int64(comment.Time), 0)
		timestr := tm.Format("2006-01-02 15:04:05")

		commentR := &model.CommentR{}
		commentR.Time = timestr
		commentR.Content = template.HTML(comment.Content)
		commentR.Id = comment.Id
		commentR.LoveNum = comment.LoveNum
		commentR.Parent = comment.Parent
		commentR.Replys = []*model.ReplyR{}
		commentR.UserName = comment.UserName
		commentR.HeadIcon = comment.HeadIcon

		replys, err := mongocli.GetCommentByParent(comment.Id)
		if err != nil {
			log.Println("get reply by comment id ", comment.Id, "failed, error is ", err)
			continue
		}

		commentR.ReplyNum = len(replys)

		for _, reply := range replys {
			replyR := &model.ReplyR{}
			replyR.Content = template.HTML(reply.Content)
			replyR.Id = reply.Id
			replyR.LoveNum = reply.LoveNum
			replyR.Parent = reply.Parent
			rtm := time.Unix(int64(reply.Time), 0)
			timestr := rtm.Format("2006-01-02 15:04:05")
			replyR.Time = timestr
			replyR.UserName = reply.UserName
			replyR.HeadIcon = reply.HeadIcon
			commentR.Replys = append(commentR.Replys, replyR)
		}

		articleR.Comments = append(articleR.Comments, commentR)
	}

	for _, recommend := range recommends {
		homeR := &model.HomeArticleR{}
		homeR.Id = recommend.Id
		homeR.Title = recommend.Title
		homeR.LoveNum = recommend.LoveNum
		homeR.ScanNum = recommend.ScanNum
		articleR.RecommendList = append(articleR.RecommendList, homeR)
	}

	hotarticles, err := mongocli.HotArticles()
	if err != nil {
		c.HTML(http.StatusOK, "home/errorpage.html", "get hot articles failed")
		return
	}

	for _, hot := range hotarticles {
		homeR := &model.HomeArticleR{}
		homeR.Id = hot.Id
		homeR.Title = hot.Title
		homeR.LoveNum = hot.LoveNum
		homeR.ScanNum = hot.ScanNum
		articleR.HotList = append(articleR.HotList, homeR)
	}

	newcomments, err := mongocli.GetNewComments()
	if err != nil {
		c.HTML(http.StatusOK, "home/errorpage.html", "get hot new comments failed")
		return
	}

	for _, newcomment := range newcomments {
		commentR := &model.CommentR{}
		commentR.Content = template.HTML(trimHtml(newcomment.Content))
		commentR.UserName = newcomment.UserName
		commentR.ArtId = newcomment.ArtId
		//log.Println("commentR.ArtId is ", newcomment.ArtId)
		info, err := mongocli.GetArticleInfo(commentR.ArtId)
		if err != nil {
			log.Println("get article ", commentR.ArtId, " info failed, err is ", err)
			continue
		}
		commentR.ArtTitle = info.Title

		articleR.CommentList = append(articleR.CommentList, commentR)
	}

	menus, err := mongocli.GetMenuListByParent("")
	if err != nil {
		c.HTML(http.StatusOK, "home/errorpage.html", "get menu list by parent failed")
		return
	}

	for _, menu := range menus {
		navCat := &model.NavCatR{}
		navCat.CatId = menu.CatId
		navCat.Index = menu.Index
		navCat.Name = menu.Name
		articleR.NavCatList = append(articleR.NavCatList, navCat)
	}

	c.HTML(http.StatusOK, "home/articlepage.html", articleR)
}

//文章增加点赞数
func AddLoveNum(c *gin.Context) {

	rsp := &model.AddLoveNumRsp{}
	rsp.Code = model.SUCCESS_NO
	rsp.Msg = model.MSG_SUCCESS
	defer func() {
		c.JSON(http.StatusOK, rsp)
	}()
	req := &model.AddLoveNumReq{}
	err := c.BindJSON(req)
	if err != nil {
		rsp.Code = model.ERR_JSON_UNPACK
		rsp.Msg = model.MSG_JSON_UNPACK
		log.Println(model.MSG_JSON_UNPACK)
		return
	}

	err = mongocli.UpdateArticleLoveNum(req)
	if err != nil {
		log.Println("update article love num failed , err is ", err)
		rsp.Code = model.ERR_ARTICLE_LOVENUM
		rsp.Msg = model.MSG_UPDATE_ARTICLE_LOVENUME
		return
	}

}

//点击评论
func Comment(c *gin.Context) {
	log.Println("收到评论文章请求")
	comentR := model.CommentRsp{}
	comentR.Res = model.MSG_SUCCESS
	comentR.Replys = []*model.ReplyR{}
	defer func() {
		c.HTML(http.StatusOK, "home/comment.html", comentR)
	}()
	req := &model.CommentReq{}
	err := c.BindJSON(req)
	if err != nil {
		log.Println("json unpack failed")
		comentR.Res = model.MSG_JSON_UNPACK
		return
	}

	comentR.Res = model.MSG_SUCCESS
	comentR.Content = template.HTML(req.Content)
	comentR.HeadIcon = req.HeadIcon
	comentR.Id = ksuid.New().String()
	comentR.LoveNum = 0
	comentR.Parent = req.Parent
	comentR.ReplyNum = 0
	tstamp := time.Now().Local().Unix()
	comentR.Time = time.Unix(tstamp, 0).Format("2006-01-02 15:04:05")
	comentR.UserName = req.UserName

	comentdb := &model.Comment{}
	comentdb.Id = comentR.Id
	comentdb.Content = req.Content
	comentdb.HeadIcon = req.HeadIcon
	comentdb.LoveNum = 0
	comentdb.Parent = req.Parent
	comentdb.Time = int(tstamp)
	comentdb.UserName = req.UserName
	comentdb.ArtId = req.ArtId

	if err := mongocli.AddComment(comentdb); err != nil {
		log.Println("insert comment db failed")
		comentR.Res = "insert comment db failed"
		return
	}
}

//点击评论喜欢数
func ComLove(c *gin.Context) {
	log.Println("点击评论喜欢")
	loveRsp := model.ComLoveRsp{}
	loveRsp.Code = model.SUCCESS_NO
	loveRsp.Msg = model.MSG_SUCCESS
	loveReq := &model.ComLoveReq{}
	defer func() {
		c.JSON(http.StatusOK, loveRsp)
	}()
	err := c.BindJSON(loveReq)
	if err != nil {
		loveRsp.Code = model.ERR_JSON_UNPACK
		loveRsp.Msg = model.MSG_JSON_UNPACK
		return
	}

	err = mongocli.AddComLove(loveReq)
	if err != nil {
		loveRsp.Code = model.ERR_COM_LOVE
		loveRsp.Msg = model.MSG_COM_LOVE
		return
	}
}

//点赞回复喜欢数
func ReplyLove(c *gin.Context) {
	log.Println("点击回复区喜欢")
	loveRsp := model.RplLoveRsp{}
	loveRsp.Code = model.SUCCESS_NO
	loveRsp.Msg = model.MSG_SUCCESS
	loveReq := &model.RplLoveReq{}
	defer func() {
		c.JSON(http.StatusOK, loveRsp)
	}()
	err := c.BindJSON(loveReq)
	if err != nil {
		loveRsp.Code = model.ERR_JSON_UNPACK
		loveRsp.Msg = model.MSG_JSON_UNPACK
		return
	}

	err = mongocli.AddRplLove(loveReq)
	if err != nil {
		loveRsp.Code = model.ERR_COM_LOVE
		loveRsp.Msg = model.MSG_COM_LOVE
		return
	}
}

//获取子分类下文章信息列表

func SubCatArtInfos(c *gin.Context) {
	log.Println("获取子分类下文章信息")
	infoR := model.SubCatArtInfoR{}
	infoR.Msg = model.MSG_SUCCESS
	infoR.SubCatArtInfos = []*model.ArtInfoR{}
	infoReq := &model.SubCatArtInfoReq{}
	defer func() {
		c.HTML(http.StatusOK, "home/leftsubcat.html", infoR)
	}()

	err := c.BindJSON(infoReq)
	if err != nil {
		infoR.Msg = model.MSG_JSON_UNPACK
		log.Println("json unparse failed, err is ", err)
		return
	}

	infos, err := mongocli.SubCatArtInfos(infoReq.Cat, infoReq.SubCat)
	if err != nil {
		infoR.Msg = model.MSG_COM_LOVE
		log.Println("get articles by cat & subcat failed , err is ", err)
		return
	}

	for _, info := range infos {
		artinfo := &model.ArtInfoR{}
		artinfo.ArtId = info.Id
		artinfo.ArtSubTitle = info.Subtitle
		infoR.SubCatArtInfos = append(infoR.SubCatArtInfos, artinfo)
	}
}

//ArtDetail
func ArtDetail(c *gin.Context) {
	log.Println("分类页面获取文章详情")
	detailR := model.ArticleDetailR{}
	detailR.Msg = model.MSG_SUCCESS

	infoReq := &model.ArtdetailReq{}
	defer func() {
		c.HTML(http.StatusOK, "home/articledetail.html", detailR)
	}()

	err := c.BindJSON(infoReq)
	if err != nil {
		detailR.Msg = model.MSG_JSON_UNPACK
		log.Println("json unparse failed, err is ", err)
		return
	}

	article, err := mongocli.GetArticleId(infoReq.Id)
	if err != nil {
		detailR.Msg = model.MSG_ARTICLE_ID
		log.Println("get article by id failed , err is ", err)
		return
	}

	//获取评论信息
	comments, err := mongocli.GetCommentByParent(article.ArticleInfo.Id)
	if err != nil {
		detailR.Msg = model.MSG_COMMENT_BYPARENT
		log.Println("get article by id failed , err is ", err)
		return
	}

	num, bres := c.Get("visitnum")
	if bres == false {
		log.Println("get visitnum failed !")
		detailR.Msg = model.MSG_VISITNUM_FAILED
		return
	}

	//内容区文章信息和内容
	err = mongocli.AddArticleScan(infoReq.Id)
	if err != nil {
		log.Println("add article scan num failed , error is ", err)
		detailR.Msg = model.MSG_ADD_VISITNUM
		return
	}

	detailR.Author = article.Author
	detailR.Cat = article.Cat
	detailR.CommentNum = len(comments)
	detailR.Content = template.HTML(article.Content)
	createtm := time.Unix(article.CreateAt, 0)
	detailR.CreateAt = createtm.Format("2006-01-02 15:04:05")

	lasttm := time.Unix(article.LastEdit, 0)
	detailR.LastEdit = lasttm.Format("2006-01-02 15:04:05")
	detailR.Id = article.ArticleInfo.Id
	detailR.Index = article.Index
	detailR.LoveNum = article.LoveNum
	detailR.ScanNum = article.ScanNum
	detailR.Subcat = article.Subcat
	detailR.Subtitle = article.Subtitle
	detailR.Title = article.Title
	detailR.VisitNum = num.(int64)

	for _, comment := range comments {
		tm := time.Unix(int64(comment.Time), 0)
		timestr := tm.Format("2006-01-02 15:04:05")

		commentR := &model.CommentR{}
		commentR.Time = timestr
		commentR.Content = template.HTML(comment.Content)
		commentR.Id = comment.Id
		commentR.LoveNum = comment.LoveNum
		commentR.Parent = comment.Parent
		commentR.Replys = []*model.ReplyR{}
		commentR.UserName = comment.UserName
		commentR.HeadIcon = comment.HeadIcon

		replys, err := mongocli.GetCommentByParent(comment.Id)
		if err != nil {
			log.Println("get reply by comment id ", comment.Id, "failed, error is ", err)
			continue
		}

		commentR.ReplyNum = len(replys)

		for _, reply := range replys {
			replyR := &model.ReplyR{}
			replyR.Content = template.HTML(reply.Content)
			replyR.Id = reply.Id
			replyR.LoveNum = reply.LoveNum
			replyR.Parent = reply.Parent
			rtm := time.Unix(int64(reply.Time), 0)
			timestr := rtm.Format("2006-01-02 15:04:05")
			replyR.Time = timestr
			replyR.UserName = reply.UserName
			replyR.HeadIcon = reply.HeadIcon
			commentR.Replys = append(commentR.Replys, replyR)
		}

		detailR.Comments = append(detailR.Comments, commentR)
	}
}

//评论回复数
func ComReply(c *gin.Context) {
	log.Println("点击评论回复")
	req := &model.ComReplyReq{}
	replyR := model.CommentReplyR{}
	replyR.Res = model.MSG_SUCCESS
	defer func() {
		c.HTML(http.StatusOK, "home/reply.html", replyR)
	}()

	err := c.BindJSON(req)
	if err != nil {
		replyR.Res = model.MSG_JSON_UNPACK
		return
	}

	dbreply := &model.Comment{}
	dbreply.Content = req.Content
	dbreply.HeadIcon = req.HeadIcon
	dbreply.Id = ksuid.New().String()
	dbreply.LoveNum = 0
	dbreply.Parent = req.Parent
	tmstamp := int(time.Now().Local().Unix())
	dbreply.Time = tmstamp
	dbreply.UserName = req.UserName
	dbreply.ArtId = req.ArtId

	err = mongocli.AddComment(dbreply)
	if err != nil {
		log.Println("insert reply into db failed, err is ", err)
		replyR.Res = "insert reply into db failed"
		return
	}

	replyR.Content = template.HTML(req.Content)
	replyR.HeadIcon = req.HeadIcon
	replyR.Id = dbreply.Id
	replyR.LoveNum = dbreply.LoveNum
	replyR.Parent = req.Parent
	replyR.Time = time.Unix(int64(tmstamp), 0).Format("2006-01-02 15:04:05")
	replyR.UserName = req.UserName
	replyR.Res = model.MSG_SUCCESS
}
