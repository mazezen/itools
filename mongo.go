package itools

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDb struct {
	c       *mongo.Client
	options *options.ClientOptions
}

// NewMongoDb mongodb 连接
func NewMongoDb(c *mongo.Client, addr string) (*MongoDb, error) {
	var err error
	if c == nil {
		c2 := &MongoDb{
			c:       &mongo.Client{},
			options: options.Client().ApplyURI(addr),
		}

		c2.c, err = mongo.Connect(context.TODO(), c2.options)
		if err != nil {
			return nil, err
		}

		if err = c2.c.Ping(context.TODO(), nil); err != nil {
			return nil, err
		}

		return c2, nil
	}
	return &MongoDb{
		c:       c,
		options: options.Client().ApplyURI(addr),
	}, nil
}

// InsertCollection 写入 db 写入的数据库  collection 写入的文档（表） data 写入的数据
func (m *MongoDb) InsertCollection(c *mongo.Client, db string, collection string, data interface{}) (
	res *mongo.InsertOneResult, err error) {
	return c.Database(db).Collection(collection).InsertOne(context.TODO(), data)
}

// BatchInsertCollection 批量写入
func (m *MongoDb) BatchInsertCollection(c *mongo.Client, db string, collection string, data []interface{}) (
	res *mongo.InsertManyResult, err error) {
	return c.Database(db).Collection(collection).InsertMany(context.TODO(), data)
}

// UpdateOneRecord 修改单条记录
func (m *MongoDb) UpdateOneRecord(c *mongo.Client, db string, collection string, id string, data bson.D) (
	res *mongo.UpdateResult, err error) {
	col := c.Database(db).Collection(collection)
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.D{{"_id", objId}}
	d := bson.D{{"$set", data}}
	res, err = col.UpdateOne(context.TODO(), filter, d)
	return
}

// CountCollection 统计
func (m *MongoDb) CountCollection(c *mongo.Client, db string, collection string, filter map[string]interface{}) (int64, int64, error) {
	col := c.Database(db).Collection(collection)

	esCount, err := col.EstimatedDocumentCount(context.TODO())
	if err != nil {
		return -1, -1, err
	}

	if len(filter) == 0 {
		count, err := col.CountDocuments(context.TODO(), bson.D{{}})
		if err != nil {
			return -1, -1, err
		}
		return esCount, count, nil
	}

	if len(filter) > 0 {
		for k, v := range filter {
			count, err := col.CountDocuments(context.TODO(), bson.D{{k, v}})
			if err != nil {
				return -1, -1, err
			}
			return esCount, count, nil
		}
	}

	return -1, -1, errors.New("查询条件异常")
}

// DeleteOneRecord 删除单条记录
func (m *MongoDb) DeleteOneRecord(c *mongo.Client, db string, collection string, filter map[string]interface{}) (
	res *mongo.DeleteResult, err error) {
	col := c.Database(db).Collection(collection)
	for k, v := range filter {
		bd := bson.D{{k, v}}
		return col.DeleteOne(context.TODO(), bd)
	}
	return
}
