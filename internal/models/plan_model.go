package models

import (
	//"fmt"
	//"context"
	//"fmt"
	//"log"
	"sync"
	"time"

	//"go.mongodb.org/mongo-driver/bson"
	//"go.mongodb.org/mongo-driver/mongo/options"
	//"go.mongodb.org/mongo-driver/mongo"
	//"github.com/9Neechan/Startup/configs"
	//"github.com/9Neechan/Startup/internal/models"

	//"github.com/9Neechan/Startup/internal/mongoclient"
)

type Plan struct {
	sync.Mutex

	Create_time time.Time
	tasks       map[int] Task
	//tasks *mongo.Collection
}
/*
func NewPost() *Post {
	p := &Post{}
	p.Create_time = time.Now()
	p.tasks = make(map[int]taskstore.Task)
	return p
}*/

// CreateTask creates a new post in mongo.

/*
func (p *Plan) CreatePost(id int, title string, tg_user_id int) int {
	p.Lock()
	defer p.Unlock()

	post := Plan{
		Create_time:      title,
		Tg_user_id: tg_user_id,
	}

	result, err := coll_task.InsertOne(context.Background(), task)
	if err != nil {
		//log.Fatal(err)
		fmt.Printf("Documents inserted: %v\n", result)
	}

	return task.Id
}
*/
