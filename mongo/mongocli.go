package mongocli

import (
	"context"
	"log"
	"time"

	"bstgo-blog/config"
	"bstgo-blog/logger"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var MongoClient *mongo.Client = nil
var MongoDb *mongo.Database = nil

func MongoInit() (e error) {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	// 连接uri
	uri := "mongodb://" + config.TotalCfgData.Mongo.User + ":" + config.TotalCfgData.Mongo.Passwd +
		"@" + config.TotalCfgData.Mongo.Host + "/?authSource=admin"
	log.Println("uri is ", uri)
	// 构建mongo连接可选属性配置
	opt := new(options.ClientOptions)
	// 设置最大连接的数量
	opt = opt.SetMaxPoolSize(uint64(config.TotalCfgData.Mongo.MaxPoolSize))
	// 设置连接超时时间 5000 毫秒
	du, _ := time.ParseDuration(config.TotalCfgData.Mongo.ConTimeOut)
	opt = opt.SetConnectTimeout(du)
	// 设置连接的空闲时间 毫秒
	mt, _ := time.ParseDuration(config.TotalCfgData.Mongo.MaxConIdle)
	opt = opt.SetMaxConnIdleTime(mt)
	// 开启驱动
	MongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(uri), opt)
	if err != nil {
		e = err
		log.Println("err is ", err)
		return
	}
	// 注意，在这一步才开始正式连接mongo
	e = MongoClient.Ping(ctx, readpref.Primary())
	if e != nil {
		log.Println("err is ", e)
	}

	log.Println("mongo init success!!!")
	logger.Sugar.Info("mongo init success!!!")
	//连接数据库
	MongoDb = MongoClient.Database(config.TotalCfgData.Mongo.Database)
	initAdmin()
	return

}

func MongoRelease() {
	MongoClient.Disconnect(context.TODO())
}
