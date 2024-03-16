package models

import (
	"context"
	//"fmt"
	//"fmt"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	//"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	//"go.mongodb.org/mongo-driver/mongo"
)

type Task struct {
	sync.Mutex

	Task_id            int
	Title              string
	Tg_user_id         int
	Status             bool
	Created_time       int64
	Change_status_time int64
	Channel_id         int64
	Channel_username   string
	Plan_id            int
}

/*
func NewTask() *Task {
	t := &Task{}
	t.Id = 0
	t.Title = ""
	t.Tg_user_id = 0
	return t
}*/

// creates a new task in mongo
func (t *Task) CreateTask() (*Task, error) {
	t.Lock()
	defer t.Unlock()

	timestamp := time.Now().Unix()
	t.Created_time = timestamp

	_, err := coll_task.InsertOne(context.Background(), &t) // task
	if err != nil {
		return &Task{}, err
	}

	return t, nil
}

// retrieves a task from the store, by id
func (t *Task) GetTaskById(task_id int) (*Task, error) {
	t.Lock()
	defer t.Unlock()

	results, err := t.FindTaskInMongo(task_id)

	return results, err
}

// deletes the task with the given id
func (t *Task) DeleteTaskById(task_id int) error {
	t.Lock()
	defer t.Unlock()

	if _, err := t.FindTaskInMongo(task_id); err != nil {
		return err
	}

	_, err := coll_task.DeleteOne(context.TODO(), bson.D{{"task_id", task_id}})
	if err != nil {
		return err
	}

	return nil
}

// Меняет статус таски и время изменения
func (t *Task) ChangeTaskStatus(task_id int) error {
	t.Lock()
	defer t.Unlock()

	task, err := t.FindTaskInMongo(task_id)
	if err != nil {
		return err
	}

	timestamp := time.Now().Unix()

	filter := bson.D{{"task_id", task_id}}
	update := bson.D{{"$set", bson.D{{"status", !task.Status}, {"change_status_time", timestamp}}}}
	_, err = coll_task.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}

	return nil
}

// returns all the tasks in the store
func (ts *Task) GetAllTasks() ([]Task, error) {
	ts.Lock()
	defer ts.Unlock()

	allTasks, err := FindFewTasksInMomgo(bson.D{}, nil, coll_task)

	return allTasks, err
}

// Возвращает список таксок по tg_user_id
func (ts *Task) GetTasksByTgUserId(tg_user_id int) ([]Task, error) {
	ts.Lock()
	defer ts.Unlock()

	//opts := options.Find().SetSort(bson.D{{"tg_user_id", tg_user_id}})
	filter := bson.D{{"tg_user_id", tg_user_id}}
	userTasks, err := FindFewTasksInMomgo(filter, nil, coll_task)

	return userTasks, err
}

// Среди тасок, сортированных в порядке убывания по времени создания,
// пропускает shift первых и после них возвращает с следующих
func (ts *Task) GetShiftTasksByTimeStamp(shift, c int) ([]Task, error) {
	ts.Lock()
	defer ts.Unlock()

	opts := options.Find().SetSort(bson.D{{"change_status_time", -1}, {"task_id", -1}}).SetLimit(int64(c)).SetSkip(int64(shift))
	shiftTasks, err := FindFewTasksInMomgo(bson.D{}, opts, coll_task)

	return shiftTasks, err
}

// Возвращает список тасок последнего плана по tg_user_id, т.е. max(plan_id)
func (ts *Task) GetLastPlanByTgUserId(tg_user_id int) ([]Task, error) {
	ts.Lock()
	defer ts.Unlock()

	filter1 := bson.D{{"tg_user_id", tg_user_id}}
	opts := options.Find().SetSort(bson.D{{"plan_id", -1}}).SetLimit(1)
	task, err := FindFewTasksInMomgo(filter1, opts, coll_task)
	if err != nil {
		return task, err
	}

	filter2 := bson.D{{"plan_id", task[0].Plan_id}} // {"tg_user_id", tg_user_id},
	planTasks, err := FindFewTasksInMomgo(filter2, nil, coll_task)

	return planTasks, err
}

func (t *Task) FindTaskInMongo(task_id int) (*Task, error) {
	filter := bson.D{{"task_id", task_id}}
	err := coll_task.FindOne(context.TODO(), filter).Decode(t)
	if err != nil {
		return &Task{}, err
	}

	return t, nil
}

/*
// GetTasksByDueDate returns all the tasks that have the given due date, in
// arbitrary order.
func (ts *TaskStore) GetTasksByDueDate(year int, month time.Month, day int) []Task {
	ts.Lock()
	defer ts.Unlock()

	var tasks []Task

	for _, task := range ts.tasks {
		y, m, d := task.Due.Date()
		if y == year && m == month && d == day {
			tasks = append(tasks, task)
		}
	}
	return tasks
}

*/

// num items in collection
/*
	opts := options.Count().SetHint("_id_")
	count, err := coll_task.CountDocuments(context.TODO(), bson.D{}, opts)
	if err != nil {
		//panic(err)
		return make([]Task, 0, 0), err
	}*/
