package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// to be exported, name need to start with Capitalize the name of the function
// type Configuration map[string]string
// var Conf Configuration = map[string]string{}
// var Conf map[string]interface{}
type ConfMap map[string]interface{}

var Conf ConfMap = ConfMap{}

func LoadEnv() error {
	l := Logger()
	err := godotenv.Load(".env")
	if err != nil {
		l.Error("Config: Error loading .env file, error:"+ err.Error())
		return err
	}
	l.Info("Config: .env loaded")
	database := os.Getenv("db_pg_name")
	Conf["jwt_expiration_time"], err = strconv.Atoi(os.Getenv("jwt_expiration_time"))
	// Conf["jwt_secret"] = os.Getenv("jwt_secret")
	Conf["jwt_secret"] = []byte(os.Getenv("jwt_secret"))
	l.Info("database: " + database)
	// abc := os.Getenv("abc") // no existge. pero no da error...
	// fmt.Println("abb: " + abc)
	// Conf["db_host"] = os.Getenv("db_pg_host")
	// Conf["db_port"] = os.Getenv("db_pg_port")
	Conf["db_name"] = os.Getenv("db_pg_name")
	// Conf["db_username"] = os.Getenv("db_pg_username")
	// Conf["db_passwd"] = os.Getenv("db_pg_password")
	Conf["db_url"] = "postgres://" + os.Getenv("db_pg_username") + ":" + os.Getenv("db_pg_password") + "@" + os.Getenv("db_pg_host") + ":" + os.Getenv("db_pg_port") + "/" + os.Getenv("db_pg_name")
	l.Info("conf db: " + Conf["db_name"].(string))
	// fmt.Println("conf secret: " + Conf["jwt_secret"].(string))
	return err
}
