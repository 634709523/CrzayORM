package CrzayORM

import (
	"CrzayORM/dialect"
	"filestore-server/db"
	"database/sql"
	"crzayORM/log"
	"crzayORM/session"
)

type Engine struct {
	db *sql.DB
	dialect dialect.Dialect
}

func NewEngine(driver,source string)(e *Engine,err error){
	// Connect Database
	db,err := sql.Open(driver, source)
	if err != nil{
		log.Error((err)
		return 
	}
	// Send a Ping to make sure the database connection is alive
	if err = db.Ping(); err != nil{
		log.Error(err)
		return 
	}
	// make sure the spicific dialect exsits
	dial,ok := dialect.GetDialect(driver)
	if !ok {
		log.Errorf(("dialect %s Not Found"), driver)
		return 
	}
	e = &Engine{db: db}
	log.Info(“Connect Database success”)
	return 
}

func (engine *Engine)Close(){
	if err := engine.db.Close(); err != nil{
		log.Error(("Failed to close database"))
		return 
	}
	log.Info(("Close database success"))
}

// create session by Engine
func (engine *Engine)NewSession()*session.Session{
	return session.New(engine.db,engine.dialect)
}
