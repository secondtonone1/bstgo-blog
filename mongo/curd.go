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
	menulist := []*model.CatMenu{}
	defer cursor.Close(context.TODO())
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
