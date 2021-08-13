package mongocli

import (
	"bstgo-blog/model"
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func init() {

}

func GetSessionById(sessionId string) (*model.Session, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
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
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
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
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
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
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	//指定连接集合
	col := MongoDb.Collection("sessions")
	insertRes, err := col.InsertOne(ctx, session)
	if err != nil {
		return primitive.ObjectID{}, err
	}

	return insertRes.InsertedID.(primitive.ObjectID), nil
}

func SaveLoginFailed(loginfailed *model.LoginFailed) error {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	//指定连接集合
	col := MongoDb.Collection("loginfaileds")
	_, err := col.InsertOne(ctx, loginfailed)
	if err != nil {
		return err
	}

	return nil
}

func GetLoginFailed(email string) (*model.LoginFailed, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
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
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
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

func GetMenuList() (*model.Menu, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	//指定连接集合
	col := MongoDb.Collection("menu")
	//设置查询条件
	filter := bson.M{}
	menu := &model.Menu{}
	err := col.FindOne(ctx, filter).Decode(menu)
	if err != nil {
		return nil, err
	}

	return menu, nil
}

func SaveMenuList(menu *model.Menu) error {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	//指定连接集合
	col := MongoDb.Collection("menu")
	_, err := col.InsertOne(ctx, menu)
	if err != nil {
		return err
	}

	return nil
}

func UpdateMenuList(menu *model.Menu) error {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	//指定连接集合
	col := MongoDb.Collection("menu")
	//设定更新filter
	filter := bson.D{{}}
	update := bson.D{{"$set",
		bson.D{
			{"catmenus", menu.CatMenus_},
		}}}

	_, err := col.UpdateOne(ctx, filter, update)
	return err
}

func UpdateSortMenu(submenu *model.SortMenuReq) error {

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	//指定连接集合
	col := MongoDb.Collection("menu")
	//设定更新filter
	//filter := bson.D{{"catmenus.catid", submenu.ParentId}}
	filter := bson.D{{"catmenus",
		bson.D{{"$elemMatch",
			bson.D{{"catid",
				submenu.ParentId}},
		}},
	}}
	update := bson.D{{"$set",
		bson.D{
			{"catmenus.$.subcatmenus", submenu.Menu},
		}}}

	_, err := col.UpdateOne(ctx, filter, update)
	return err
}

func GetSubCatSelect(catId string) (*model.Menu, error) {
	log.Println("cat id is ", catId)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	//指定连接集合
	col := MongoDb.Collection("menu")
	//设定更新filter
	filter := bson.D{{},
		{"catmenus",
			bson.D{{"$elemMatch",
				bson.D{{"catid",
					catId}}}},
		},
	}
	menu := model.Menu{}
	err := col.FindOne(ctx, filter).Decode(&menu)
	if err != nil {
		return nil, err
	}
	return &menu, nil
}
