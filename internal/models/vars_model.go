package models

import (
	"context"
	"sync"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson/primitive"
	//"go.mongodb.org/mongo-driver/mongo"
)

type Vars struct {
	sync.Mutex

	Last_task_id int
	Last_plan_id int
}

func (v *Vars) GetVars() ([]Vars, error) {
	v.Lock()
	defer v.Unlock()

	vars, err := v.FindVarsInMomgo(bson.D{}, nil)

	return vars, err
}

func (v *Vars) FindVarsInMomgo(filter primitive.D, opts *options.FindOptions) ([]Vars, error) {
		returnVars := make([]Vars, 0)

		cur, err := coll_vars.Find(context.TODO(), filter, opts)
		if err != nil {
			return make([]Vars, 0), err
		}

		// проходим по всем items и добавляем в массив
		defer cur.Close(context.Background())
		for cur.Next(context.Background()) {
			var result Vars

			err := cur.Decode(&result)
			if err != nil {
				return make([]Vars, 0), err
			}

			returnVars = append(returnVars, result)
		}
		if err := cur.Err(); err != nil {
			return make([]Vars, 0), err
		}

		return returnVars, nil
	
}
