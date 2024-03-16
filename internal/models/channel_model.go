package models

import (
	"context"
	"sync"
	//"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	//"go.mongodb.org/mongo-driver/mongo"
)

type Channel struct {
	sync.Mutex

	Title              string 
	Tg_user_id         int 
	Channel_id         int64 
	Username           string
	Created_tasks int
	Done_tasks int
}

func (ch *Channel) GetChannelsByTgUserId(tg_user_id int) ([]Channel, error) {
	ch.Lock()
	defer ch.Unlock()

	opts := options.Find().SetSort(bson.D{{"created_tasks", -1}})
	filter := bson.D{{"tg_user_id", tg_user_id}}
	channels, err := ch.FindFewChannelsInMomgo(filter, opts)

	return channels, err
}

func (ch *Channel) FindFewChannelsInMomgo(filter primitive.D, opts *options.FindOptions) ([]Channel, error) {
	returnTasks := make([]Channel, 0)

	cur, err := coll_channels.Find(context.TODO(), filter, opts)
	if err != nil {
		return make([]Channel, 0), err
	}

	// проходим по всем items и добавляем в массив
	defer cur.Close(context.Background())
	for cur.Next(context.Background()) {
		var result Channel

		err := cur.Decode(&result)
		if err != nil {
			return make([]Channel, 0), err
		}

		returnTasks = append(returnTasks, result)
	}
	if err := cur.Err(); err != nil {
		return make([]Channel, 0), err
	}

	return returnTasks, nil
}