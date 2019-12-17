package model

import (
	"github.com/pelletier/go-toml"
	"gopkg.in/mgo.v2"
	"io/ioutil"
	"log"
	"strconv"
	"time"
	//"strconv"
)

//const (
//	url = "edgex-mongo:27017"
//	dbName = "metadata"
//	configCol = "deviceDataAccessConfig"
//)

var (
	session          *mgo.Session
	database         *mgo.Database
	configCollection *mgo.Collection
)

type DBConfig struct {
	Host       string
	Port       int
	DB         string
	Collection string
}

type Cfg struct {
	DBConfig DBConfig
}

func dbConfig() *DBConfig {
	configFile := "configuration.toml"

	contents, _ := ioutil.ReadFile(configFile)
	c := Cfg{}

	toml.Unmarshal(contents, &c)
	return &c.DBConfig
}

type mongoLog struct {
}

func (mongoLog) Output(calldepth int, s string) error {
	log.SetFlags(log.Lshortfile)
	return log.Output(calldepth, s)
}

func ConfigCollection() *mgo.Collection {
	if configCollection != nil {
		return configCollection
	}
	db := dbConfig()
	url := db.Host + ":" + strconv.Itoa(db.Port)
	log.Printf("mongo connect info(url:%s, dbName:%s, collection:%s)\n", url, db.DB, db.Collection)

	session, err := mgo.Dial(url)
	if err != nil {
		log.Fatalf("Connect mongo failed: %v \n", err)
	}

	mgo.SetDebug(true)
	//	mgo.SetLogger(new(mongoLog))
	//	session.SetMode(mgo.Monotonic, true)
	session.SetMode(mgo.Eventual, true)
	session.SetSocketTimeout(2 * time.Second)
	session.SetSyncTimeout(2 * time.Second)
	configCollection = session.DB(db.DB).C(db.Collection)
	return configCollection
}
