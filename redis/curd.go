package redis

import (
	"bstgo-blog/model"
	"encoding/json"
	"log"
	"sort"
	"strconv"
	"time"

	"github.com/go-redis/redis"
)

//keys
const (
	STRING_ToTAL_VISIT_NUM_KEY = "total_visit"
	HSET_LV1_MENU_KEY          = "lv1_menu"
	HSET_HOT_ARTICLES_KEY      = "hot_arts"
	HSET_NEW_COMMENTS_KEY      = "new_comments"
	HSET_HOME_ARTICLES_KEY     = "home_arts"
	STRING_TOTAL_ARTICLE_NUM   = "total_art_num"
	HSET_MENUS_KEY             = "menus"
	HSET_LV2_MENU_PREFIX       = "lv2menu_"
	HSET_ARTINFO_PREFIX        = "artinfo_"
	HSET_FIRSTART              = "first_art"
	HSET_COM_PREFIX            = "com_"
	HSET_ARTICLES_KEY          = "articles"
	HSET_RELCOM_PREFIX         = "relc_"
	HSET_ARTSCAT_PREFIX        = "art_"
)

func AddVisitNum() (int64, error) {
	visit, err := rediscli.Incr(STRING_ToTAL_VISIT_NUM_KEY).Result()
	if err != nil {
		return 0, err
	}

	return visit, nil
}

func SetVisitNum(num int64) (string, error) {
	val, err := rediscli.Set(STRING_ToTAL_VISIT_NUM_KEY, num, 0).Result()
	if err != nil {
		return "", err
	}

	return val, nil
}

func GetVisitNum() (string, error) {
	val, err := rediscli.Get(STRING_ToTAL_VISIT_NUM_KEY).Result()

	if err == redis.Nil {
		return "key not exists in redis", err
	}

	if err != nil {
		return "", err
	}

	return val, nil
}

func GetLv1Menus() ([]*model.CatMenu, error) {
	menulist := []*model.CatMenu{}
	menus, err := rediscli.HGetAll(HSET_LV1_MENU_KEY).Result()
	if err != nil {
		return menulist, err
	}

	for _, val := range menus {
		menu := &model.CatMenu{}
		err := json.Unmarshal([]byte(val), menu)
		if err != nil {
			log.Println("json unmarshal failed, err is ", err)
			continue
		}
		menulist = append(menulist, menu)
	}
	return menulist, nil
}

func SetLv1Menus(menus []*model.CatMenu) error {
	for _, menu := range menus {
		menujs, err := json.Marshal(menu)
		if err != nil {
			log.Println("json marshal failed, err is ", err)
			continue
		}

		rediscli.HSet(HSET_LV1_MENU_KEY, menu.CatId, menujs)
	}

	rediscli.Expire(HSET_LV1_MENU_KEY, time.Minute*33)

	return nil
}

func GetHotArticles() ([]*model.ArticleInfo, error) {
	artinfos := []*model.ArticleInfo{}
	infomap, err := rediscli.HGetAll(HSET_HOT_ARTICLES_KEY).Result()
	if err != nil {
		return artinfos, err
	}

	for _, info := range infomap {
		artinfo := &model.ArticleInfo{}
		err := json.Unmarshal([]byte(info), artinfo)
		if err != nil {
			log.Println("json unmarshal failed, err is ", err)
			continue
		}

		artinfos = append(artinfos, artinfo)
	}

	return artinfos, nil
}

func SetHotArts(hotarts []*model.ArticleInfo) error {
	for _, art := range hotarts {
		artjs, err := json.Marshal(art)
		if err != nil {
			log.Println("json marshal failed, err is ", err)
			continue
		}

		rediscli.HSet(HSET_HOT_ARTICLES_KEY, art.Id, artjs)
	}

	rediscli.Expire(HSET_HOT_ARTICLES_KEY, time.Minute*18)
	return nil
}

func GetNewComments() ([]*model.Comment, error) {
	comments := []*model.Comment{}
	comInfos, err := rediscli.HGetAll(HSET_NEW_COMMENTS_KEY).Result()
	if err != nil {
		return comments, err
	}

	for _, info := range comInfos {
		comment := &model.Comment{}
		err := json.Unmarshal([]byte(info), comment)
		if err != nil {
			log.Println("json unmarshal failed, err is ", err)
			continue
		}

		comments = append(comments, comment)
	}

	return comments, nil
}

func SetNewComments(comments []*model.Comment) error {
	for _, comment := range comments {
		comjs, err := json.Marshal(comment)
		if err != nil {
			log.Println("json marshal failed, err is ", err)
			continue
		}

		rediscli.HSet(HSET_NEW_COMMENTS_KEY, comment.Id, comjs)
	}

	rediscli.Expire(HSET_HOT_ARTICLES_KEY, time.Minute*10)
	return nil
}

//home??????????????????
func GetHomeArticleDetails() ([]*model.Article_, error) {
	articles := []*model.Article_{}
	artmap, err := rediscli.HGetAll(HSET_HOME_ARTICLES_KEY).Result()
	if err != nil {
		return articles, err
	}

	for _, info := range artmap {
		article := &model.Article_{}
		err := json.Unmarshal([]byte(info), article)
		if err != nil {
			log.Println("json unmarshal failed, err is ", err)
			continue
		}

		articles = append(articles, article)
	}

	return articles, nil
}

//home????????????????????????redis
func SetHomeArticleDetails(arts []*model.Article_) error {
	for _, art := range arts {
		artjs, err := json.Marshal(art)
		if err != nil {
			log.Println("json marshal failed, err is ", err)
			continue
		}

		rediscli.HSet(HSET_HOME_ARTICLES_KEY, art.ArticleInfo.Id, artjs)
	}

	rediscli.Expire(HSET_HOME_ARTICLES_KEY, time.Hour*2)
	return nil
}

//??????????????????
func GetTotalArtNum() (int, error) {
	val, err := rediscli.Get(STRING_TOTAL_ARTICLE_NUM).Result()

	if err == redis.Nil {
		return 0, err
	}

	if err != nil {
		return 0, err
	}

	valint, err := strconv.Atoi(val)
	if err != nil {
		return 0, err
	}

	return valint, nil
}

//??????????????????
func SetTotalArtNum(total int) error {
	_, err := rediscli.Set(STRING_TOTAL_ARTICLE_NUM, total, 15*time.Hour).Result()
	return err
}

//??????id????????????
func GetMenuById(catid string) (*model.CatMenu, error) {
	catmenu := &model.CatMenu{}
	menujs, err := rediscli.HGet(HSET_MENUS_KEY, catid).Result()
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(menujs), catmenu)
	if err != nil {
		return nil, err
	}
	return catmenu, nil
}

//?????????????????????redis ????????????
func SetMenuToSet(menu *model.CatMenu) error {
	menujs, err := json.Marshal(menu)
	if err != nil {
		return err
	}

	rediscli.HSet(HSET_MENUS_KEY, menu.CatId, menujs)
	rediscli.Expire(HSET_MENUS_KEY, time.Minute*40)
	return nil
}

//??????????????????????????????
func GetLv2MenusByCatId(cat string) ([]*model.CatMenu, error) {
	catkey := HSET_LV2_MENU_PREFIX + cat
	menus := []*model.CatMenu{}
	menumap, err := rediscli.HGetAll(catkey).Result()
	if err != nil {
		return menus, err
	}

	for _, info := range menumap {
		menu := &model.CatMenu{}
		err := json.Unmarshal([]byte(info), menu)
		if err != nil {
			log.Println("json unmarshal failed, err is ", err)
			continue
		}

		menus = append(menus, menu)
	}

	sort.Sort(model.MenuSlice(menus))

	return menus, nil
}

//home????????????????????????redis
func SetLv2MenusByCatId(cat string, menus []*model.CatMenu) error {
	catkey := HSET_LV2_MENU_PREFIX + cat
	for _, menu := range menus {
		menujs, err := json.Marshal(menu)
		if err != nil {
			log.Println("json marshal failed, err is ", err)
			continue
		}

		rediscli.HSet(catkey, menu.CatId, menujs)
	}

	rediscli.Expire(catkey, time.Hour*5)
	return nil
}

//??????????????????????????????
func GetArtInfoByCat(cat string) ([]*model.ArticleInfo, error) {
	artinfos := []*model.ArticleInfo{}
	catkey := HSET_ARTINFO_PREFIX + cat
	infomap, err := rediscli.HGetAll(catkey).Result()
	if err != nil {
		return artinfos, err
	}

	for _, info := range infomap {
		artinfo := &model.ArticleInfo{}
		err := json.Unmarshal([]byte(info), artinfo)
		if err != nil {
			log.Println("json unmarshal failed, err is ", err)
			continue
		}

		artinfos = append(artinfos, artinfo)
	}

	return artinfos, nil
}

//??????????????????
func SetArtInfoByCat(cat string, infos []*model.ArticleInfo) error {
	catkey := HSET_ARTINFO_PREFIX + cat
	for _, info := range infos {
		infojs, err := json.Marshal(info)
		if err != nil {
			log.Println("json marshal failed, err is ", err)
			continue
		}

		rediscli.HSet(catkey, info.Id, infojs)
	}

	rediscli.Expire(catkey, time.Minute*55)
	return nil
}

//??????????????????
func GetFirstArtByCat(cat string, subcat string) (*model.Article_, error) {
	article := &model.Article_{}
	artjs, err := rediscli.HGet(HSET_FIRSTART, cat+"_"+subcat).Result()
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(artjs), article)
	if err != nil {
		return nil, err
	}

	return article, nil
}

//??????????????????
func SetFristArtByCat(cat string, subcat string, art *model.Article_) error {
	artjs, err := json.Marshal(art)
	if err != nil {
		return err
	}
	_, err = rediscli.HSet(HSET_FIRSTART, cat+"_"+subcat, artjs).Result()
	if err != nil {
		return err
	}

	rediscli.Expire(HSET_FIRSTART, time.Minute*35)
	return nil
}

//?????????????????????
func GetCommentsByParent(pid string) ([]*model.Comment, error) {
	comments := []*model.Comment{}
	valmap, err := rediscli.HGetAll(HSET_COM_PREFIX + pid).Result()
	if err != nil {
		return comments, err
	}
	for _, val := range valmap {
		comment := &model.Comment{}
		err := json.Unmarshal([]byte(val), comment)
		if err != nil {
			continue
		}

		comments = append(comments, comment)
	}
	return comments, nil
}

//??????????????????
func SetCommentsByParent(pid string, comments []*model.Comment) error {
	for _, comment := range comments {
		comjs, err := json.Marshal(comment)
		if err != nil {
			continue
		}
		rediscli.HSet(HSET_COM_PREFIX+pid, comment.Id, comjs)
	}

	rediscli.Expire(HSET_COM_PREFIX+pid, time.Minute*45)
	return nil
}

//????????????id????????????
func GetCommentByParent(pid string, id string) (*model.Comment, error) {
	comment := &model.Comment{}
	comjs, err := rediscli.HGet(HSET_COM_PREFIX+pid, id).Result()
	if err != nil {
		return comment, err
	}
	err = json.Unmarshal([]byte(comjs), comment)
	if err != nil {
		return comment, err
	}
	return comment, nil
}

//??????????????????
func SetCommentByParent(pid string, id string, comment *model.Comment) error {
	comjs, err := json.Marshal(comment)
	if err != nil {
		return err
	}

	_, err = rediscli.HSet(HSET_COM_PREFIX+pid, id, comjs).Result()
	if err != nil {
		return err
	}

	return nil
}

//??????id????????????
func GetArticleById(id string) (*model.Article_, error) {
	article := &model.Article_{}
	artjs, err := rediscli.HGet(HSET_ARTICLES_KEY, id).Result()
	if err != nil {
		return article, err
	}

	err = json.Unmarshal([]byte(artjs), article)
	if err != nil {
		return article, err
	}

	return article, nil
}

//????????????id????????????
func SetArticleById(art *model.Article_) error {

	artjs, err := json.Marshal(art)
	if err != nil {
		return err
	}
	_, err = rediscli.HSet(HSET_ARTICLES_KEY, art.ArticleInfo.Id, artjs).Result()
	if err != nil {
		return err
	}

	rediscli.Expire(HSET_ARTICLES_KEY, time.Minute*30)
	return nil
}

//??????????????????
func GetRelcByCat(catId string) ([]*model.ArticleInfo, error) {
	infos := []*model.ArticleInfo{}
	valmap, err := rediscli.HGetAll(HSET_RELCOM_PREFIX + catId).Result()
	if err != nil {
		return infos, err
	}
	for _, val := range valmap {
		info := &model.ArticleInfo{}
		err := json.Unmarshal([]byte(val), info)
		if err != nil {
			continue
		}

		infos = append(infos, info)
	}
	return infos, nil
}

//redis??????????????????
func SetRelcByCat(artId string, arts []*model.ArticleInfo) error {
	for _, art := range arts {
		artjs, err := json.Marshal(art)
		if err != nil {
			continue
		}
		rediscli.HSet(HSET_RELCOM_PREFIX+artId, art.Id, artjs)
	}

	rediscli.Expire(HSET_RELCOM_PREFIX+artId, time.Minute*15)
	return nil
}

//??????????????????
func ClearAll() error {
	go rediscli.FlushDBAsync().Result()
	return nil
}

//??????????????????????????????????????????
func GetArtsByCatSubCat(cat string, subcat string) ([]*model.ArticleInfo, error) {
	infos := []*model.ArticleInfo{}
	mapval, err := rediscli.HGetAll(HSET_ARTSCAT_PREFIX + cat + "_" + subcat).Result()
	if err != nil {
		return infos, err
	}

	for _, val := range mapval {
		art := &model.ArticleInfo{}
		err = json.Unmarshal([]byte(val), art)
		if err != nil {
			continue
		}

		infos = append(infos, art)
	}

	return infos, nil
}

//???????????????????????????????????????
func SetArtByCatSubCat(cat string, subcat string, infos []*model.ArticleInfo) error {

	for _, val := range infos {
		valjs, err := json.Marshal(val)
		if err != nil {
			continue
		}

		_, err = rediscli.HSet(HSET_ARTSCAT_PREFIX+cat+"_"+subcat, val.Id, valjs).Result()
		if err != nil {
			continue
		}
	}

	return nil
}
