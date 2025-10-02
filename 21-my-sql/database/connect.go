package database

import (
	"database/sql"
	"fmt"

	"sql-in-go/config"
)

func ConnectDB(cfg *config.DatabaseConfig) (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", 
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)
	
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	
	if err = db.Ping(); err != nil {
		return nil, err
	}
	
	fmt.Println("Conexión exitosa a la base de datos")
	
	return db, nil
}