package mongocli

import (
	"bstgo-blog/model"
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {

}

func GetSessionById(sessionId string) (*model.Session, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	//指定连接集合
	col := MongoDb.Collection("sessions")
	//根据sessionId去数据库查找对应session
	session := model.Session{}
	err := col.FindOne(ctx, bson.M{"sid": sessionId}).Decode(&session)
	if err != nil {
		return nil, err
	}
	return &session, nil
}

func GetAdminByEmail(email string) (*model.Admin, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	//指定连接集合
	col := MongoDb.Collection("administrator")
	//根据email 查找
	administrator := model.Admin{}
	err := col.FindOne(ctx, bson.M{"email": email}).Decode(&administrator)
	if err != nil {
		return nil, err
	}

	return &administrator, nil
}

func InitAdmin(email string, pwd string) (primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	//指定连接集合
	col := MongoDb.Collection("administrator")
	//插入email
	administrator := &model.Admin{
		Email: email,
		Pwd:   pwd,
	}
	insertRes, err := col.InsertOne(ctx, administrator)
	if err != nil {
		return primitive.ObjectID{}, err
	}
	return insertRes.InsertedID.(primitive.ObjectID), nil
}

func initAdmin() {
	//初始化admin账户
	inseres, err := InitAdmin("secondtonone1@163.com", "123456")
	if err != nil {
		log.Println("init admin failed , err is ", err)
		return
	}

	log.Println("init admin success, insert id is ", inseres)
}

func SaveSession(session *model.Session) (primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	//指定连接集合
	col := MongoDb.Collection("sessions")
	insertRes, err := col.InsertOne(ctx, session)
	if err != nil {
		return primitive.ObjectID{}, err
	}

	return insertRes.InsertedID.(primitive.ObjectID), nil
}

func SaveLoginFailed(loginfailed *model.LoginFailed) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	//指定连接集合
	col := MongoDb.Collection("loginfaileds")
	_, err := col.InsertOne(ctx, loginfailed)
	if err != nil {
		return err
	}

	return nil
}

func GetLoginFailed(email string) (*model.LoginFailed, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	//指定连接集合
	col := MongoDb.Collection("loginfaileds")
	loginfailed := &model.LoginFailed{}
	//根据email 查找
	err := col.FindOne(ctx, bson.M{"email": email}).Decode(loginfailed)
	if err != nil {
		return nil, err
	}
	return loginfailed, nil
}

func UpdateLoginFailed(loginfailed *model.LoginFailed) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	//指定连接集合
	col := MongoDb.Collection("loginfaileds")
	//设置更新条件
	filter := bson.D{{"email", loginfailed.Email}}
	update := bson.D{{"$set",
		bson.D{
			{"count", loginfailed.Count},
		},
	}}

	_, err := col.UpdateOne(ctx, filter, update)
	return err
}

func GetMenuByCat(cat string) (*model.CatMenu, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	//指定连接集合
	col := MongoDb.Collection("menus")
	//设置查询条件
	filter := bson.M{"name": cat}

	catmenu := &model.CatMenu{}
	err := col.FindOne(ctx, filter).Decode(catmenu)

	if err != nil {
		return nil, err
	}

	return catmenu, nil
}

func GetMenuListByParent(parent string) ([]*model.CatMenu, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	//指定连接集合
	col := MongoDb.Collection("menus")
	//设置查询条件
	filter := bson.M{"parent": parent}

	SORT := bson.D{{"index", 1}}
	findOptions := options.Find().SetSort(SORT)

	cursor, err := col.Find(ctx, filter, findOptions)

	if err != nil {
		return nil, err
	}

	defer cursor.Close(context.TODO())
	menulist := []*model.CatMenu{}

	for cursor.Next(context.TODO()) {
		menu := &model.CatMenu{}
		err = cursor.Decode(menu)
		if err != nil {
			return nil, err
		}
		menulist = append(menulist, menu)
	}

	return menulist, nil
}

func GetMenuById(id string) (*model.CatMenu, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	//指定连接集合
	col := MongoDb.Collection("menus")
	//设置查询条件
	filter := bson.M{"catid": id}
	catmenu := &model.CatMenu{}
	err := col.FindOne(ctx, filter).Decode(catmenu)
	if err != nil {
		return nil, err
	}

	return catmenu, nil
}

func SaveMenu(menu *model.CatMenu) (primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	//指定连接集合
	col := MongoDb.Collection("menus")
	res, err := col.InsertOne(ctx, menu)
	if err != nil {
		return primitive.ObjectID{}, err
	}

	return res.InsertedID.(primitive.ObjectID), nil
}

func SaveMenuList(menulist []*model.CatMenu) error {
	if len(menulist) == 0 {
		log.Println("menulist is empty")
		return nil
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	models := []mongo.WriteModel{}
	for _, catmenu := range menulist {
		filter := bson.D{{"catid", catmenu.CatId}}
		updatecmd := bson.D{{"$set", bson.D{{"index", catmenu.Index}}}}
		model := mongo.NewUpdateOneModel().SetFilter(filter).SetUpdate(updatecmd).SetUpsert(false)
		models = append(models, model)
	}

	//log.Println("models are ", models)
	opts := options.BulkWrite().SetOrdered(false)
	_, err := MongoDb.Collection("menus").BulkWrite(ctx, models, opts)
	return err
}

func SaveArticle(article *model.Article) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := MongoDb.Collection("articles").InsertOne(ctx, &article)
	return err
}

//获取文章列表
func GetArticlesByPage(page int) ([]*model.Article, error) {
	articles := []*model.Article{}
	if page < 1 {
		return articles, nil
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	sort := bson.D{{"lastedit", -1}}
	findOptions := options.Find().SetSort(sort)

	//从第1页获取，每次获取5条
	skipTmp := int64((page - 1) * 5)
	limitTmp := int64(5)
	findOptions.Skip = &skipTmp
	findOptions.Limit = &limitTmp

	filter := bson.D{}
	cursor, err := MongoDb.Collection("articles").Find(ctx, filter, findOptions)

	if err != nil {
		return articles, err
	}

	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		article := &model.Article{}
		if err := cursor.Decode(article); err != nil {
			continue
		}

		articles = append(articles, article)
	}

	return articles, nil
}

//获取文章总数
func ArticleTotalCount() (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.D{}
	count, err := MongoDb.Collection("articles").CountDocuments(ctx, filter)

	if err != nil {
		return 0, err
	}

	return int(count), nil
}

//搜索文章
func SearchArticle(condition *model.SearchArticleReq) ([]*model.Article, int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	filter := bson.M{}
	if condition.Year != "" && condition.Year != "不限" {
		var stamp int64 = 0
		valStr := condition.Year
		if condition.Month != "" && condition.Month != "不限" {
			tempStr := "2006-01月"
			valStr = valStr + "-" + condition.Month
			localTime, err := time.ParseInLocation(tempStr, valStr, time.Local)
			if err != nil {
				log.Println("time parse failed, err is ", err)
			} else {
				stamp = localTime.Unix()
				log.Println("query time stamp is ", stamp)
			}
		} else {
			localTime, err := time.ParseInLocation("2006", valStr, time.Local)
			if err != nil {
				log.Println("time parse failed, err is ", err)
			} else {
				stamp = localTime.Unix()
				log.Println("query time stamp is ", stamp)
			}
		}

		filter["createdAt"] = bson.M{"$gte": stamp}
	}

	if condition.Cat != "" && condition.Cat != "不限" {
		filter["cat"] = condition.Cat
	}

	if condition.Keywords != "" && condition.Keywords != "不限" {

		filter["$or"] = []bson.M{
			bson.M{
				"content": bson.M{"$regex": condition.Keywords, "$options": "$i"},
			},

			bson.M{
				"title": bson.M{"$regex": condition.Keywords, "$options": "$i"},
			},
		}
	}

	log.Println("filter is ", filter)
	sort := bson.D{{"lastedit", -1}}
	findOptions := options.Find().SetSort(sort)

	//从第1页获取，每次获取5条
	skipTmp := int64((condition.Page - 1) * 5)
	limitTmp := int64(5)
	findOptions.Skip = &skipTmp
	findOptions.Limit = &limitTmp

	articles := []*model.Article{}
	cursor, err := MongoDb.Collection("articles").Find(ctx, filter, findOptions)
	if err != nil {
		return articles, 0, err
	}

	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		article := &model.Article{}
		if err := cursor.Decode(article); err != nil {
			continue
		}

		articles = append(articles, article)
	}

	count, err := MongoDb.Collection("articles").CountDocuments(ctx, filter)

	if err != nil {
		return articles, 0, err
	}

	return articles, int(count), nil
}

//通过子分类搜索文章
func SearchArticleBySubCat(category string) ([]*model.Article, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	filter := bson.M{}
	filter["subcat"] = category

	log.Println("filter is ", filter)
	sort := bson.D{{"index", 1}}
	findOptions := options.Find().SetSort(sort)
	articles := []*model.Article{}
	cursor, err := MongoDb.Collection("articles").Find(ctx, filter, findOptions)
	if err != nil {
		return articles, err
	}

	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		article := &model.Article{}
		if err := cursor.Decode(article); err != nil {
			continue
		}

		articles = append(articles, article)
	}

	return articles, nil
}

//批量更新文章列表序列
func UpdateArticleSort(sortArt *model.ArticleSortReq) error {
	if len(sortArt.ArticleList) == 0 {
		log.Println("sort article list is empty")
		return nil
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	models := []mongo.WriteModel{}
	for _, article := range sortArt.ArticleList {
		filter := bson.M{"id": article.Id, "title": article.Title}
		updatecmd := bson.D{{"$set", bson.D{{"index", article.Index}}}}
		model := mongo.NewUpdateOneModel().SetFilter(filter).SetUpdate(updatecmd).SetUpsert(false)
		models = append(models, model)
	}

	//log.Println("models are ", models)
	opts := options.BulkWrite().SetOrdered(false)
	_, err := MongoDb.Collection("articles").BulkWrite(ctx, models, opts)
	return err
}

//获取子分类下最大index
func GetSubCatMaxIndex(subcat string) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	pipeline := bson.A{
		bson.M{
			"$match": bson.M{"subcat": subcat},
		},

		bson.M{
			"$group": bson.M{
				"_id":      bson.M{"subcat_": "$subcat"},
				"maxIndex": bson.M{"$max": "$index"}},
		},
	}
	cursor, err := MongoDb.Collection("articles").Aggregate(ctx, pipeline)
	if err != nil {
		log.Println("aggrete failed, error is ", err)
		return 0, err
	}

	maxIndex := 0
	for cursor.Next(context.Background()) {
		doc := cursor.Current
		maxindex_, err := doc.LookupErr("maxIndex")
		if err != nil {
			log.Println("LookupErr failed, error is ", err)
			return maxIndex, err
		}

		maxIndex = int(maxindex_.Int32())
		log.Println("maxindex is ", maxIndex)
	}
	log.Println("get max index is ", maxIndex)
	return maxIndex, nil
}

//删除文章
func DelArticle(title string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	filter := bson.M{"title": title}
	_, err := MongoDb.Collection("articles").DeleteOne(ctx, filter)
	if err != nil {
		log.Println("del article ", title, " failed, err is ", err)
		return err
	}
	return nil
}

//通过id获取文章
func GetArticleId(id string) (*model.Article, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	filter := bson.M{"id": id}
	article := &model.Article{}
	err := MongoDb.Collection("articles").FindOne(ctx, filter).Decode(article)
	if err != nil {
		return nil, err
	}

	return article, nil
}

//更新文章
func UpdateArticle(req *model.UpdateArticleReq) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	filter := bson.M{"id": req.Id}
	value := bson.M{}
	value["title"] = req.Title
	value["subtitle"] = req.SubTitle
	value["cat"] = req.Cat
	value["subcat"] = req.SubCat
	value["lastedit"] = req.LastEdit
	value["author"] = req.Author
	value["content"] = req.Content

	upvalue := bson.M{"$set": value}
	_, err := MongoDb.Collection("articles").UpdateOne(ctx, filter, upvalue)

	return err
}

//获取草稿总数
func DraftTotalCount() (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.D{}
	count, err := MongoDb.Collection("drafts").CountDocuments(ctx, filter)

	if err != nil {
		return 0, err
	}

	return int(count), nil
}

//获取草稿列表
func GetDraftsByPage(page int) ([]*model.Article, error) {
	articles := []*model.Article{}
	if page < 1 {
		return articles, nil
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	sort := bson.D{{"lastedit", -1}}
	findOptions := options.Find().SetSort(sort)

	//从第1页获取，每次获取5条
	skipTmp := int64((page - 1) * 5)
	limitTmp := int64(5)
	findOptions.Skip = &skipTmp
	findOptions.Limit = &limitTmp

	filter := bson.D{}
	cursor, err := MongoDb.Collection("drafts").Find(ctx, filter, findOptions)

	if err != nil {
		return articles, err
	}

	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		article := &model.Article{}
		if err := cursor.Decode(article); err != nil {
			continue
		}

		articles = append(articles, article)
	}

	return articles, nil
}
