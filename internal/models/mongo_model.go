package models

import (
	//"github.com/9Neechan/Startup/configs"
	"github.com/LSD-Learn-Strive-Develop/logarithm/tree/main/rest-api/configs"
)

var coll_channels = configs.GetCollection(configs.DB, "channels")
var coll_plans = configs.GetCollection(configs.DB, "plans")
var coll_task = configs.GetCollection(configs.DB, "tasks")
var coll_users = configs.GetCollection(configs.DB, "users")
var coll_vars = configs.GetCollection(configs.DB, "vars")
