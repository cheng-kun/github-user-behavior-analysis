package db

//import (
//	"context"
//	"github.com/github-user-behavior-analysis/backend-go/logs"
//	"github.com/mongodb/mongo-go-driver/mongo"
//	"github.com/mongodb/mongo-go-driver/x/bsonx"
//	"log"
//)
//
//func SetupMongoDB() (*mongo.Client, error)  {
//	client, err := mongo.NewClient("mongodb://nebula=ai-chengkun:Cck8079567@localhost:27017")
//	if err != nil {
//		logs.PrintLogger().Error(err)
//		return nil, err
//	}
//
//	err = client.Connect(context.TODO())
//	if err != nil {
//		logs.PrintLogger().Error()
//		return nil, err
//	}
//
//	return client, err
//}
//
//func InsertMongoDB()  {
//	client, err := SetupMongoDB()
//	if err != nil {
//		logs.PrintLogger().Error(err)
//	}
//
//	collection := client.Database("demoDB").Collection("demoDB")
//
//	logs.PrintLogger().Info(collection)
//
//	//res, err := collection.InsertOne(context.Background(), Doc{{"x", Int32(1)}})
//	//if err != nil {
//	//	logs.PrintLogger().Error(err)
//	//}
//	//id := res.InsertedID
//	//
//	//fmt.Println(id)
//	cur, err := collection.Find(context.Background(), nil)
//	if err != nil { log.Fatal(err) }
//	defer cur.Close(context.Background())
//	for cur.Next(context.Background()) {
//		elem := bsonx.Doc{}
//		err := cur.Decode(elem)
//		if err != nil { log.Fatal(err) }
//		// do something with elem....
//	}
//	if err := cur.Err(); err != nil {
//		log.Fatal(err)
//	}
//
//
//	//rows, _ := ConnDB.Query("select * from top_ten ")
//	//
//	//for rows.Next() {
//	//
//	//
//	//
//	//}
//
//}


