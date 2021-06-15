package main

import (
	"fmt"
	"os/exec"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

func RunDb() (db *sqlx.DB, err error) {
	log, _ = zap.NewProduction()

	defer log.Sync()

	host := "127.0.0.1"
	port := 3306
	user := "root"
	pass := "root"
	dbname := "feedy"

	//// DB CONNECTION ////
	pathSQL := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, pass, host, port, dbname)
	db, err = sqlx.Connect("mysql", pathSQL)

	if err != nil {
		return db, err

	} else {
		log.Info("Connexion etablished ", zap.String("database", dbname),
			zap.Int("attempt", 3), zap.Duration("backoff", time.Second))
	}
	return db, err
}

func insertDb(mail string, name string, message string, direction string) (err error) {

	db, err := RunDb()

	attempt := 1
	for attempt <= 3 && err != nil {
		//printerr
		log.Error("failed to connect database", zap.String("database", "feedy"),
			zap.Int("attempt", attempt), zap.Duration("backoff", time.Second))

		//restart mysql
		exec.Command("sudo", "service", "mysqld", "restart").Output()

		//wait
		time.Sleep(4 * time.Second)

		//retry
		db, err = RunDb()
		attempt++
	}

	if direction == "IN" {
		_, err = db.Exec("INSERT INTO messages(mail, name, message) VALUES(?, ?, ?)", mail, name, message)
	}
	_, err = db.Exec("INSERT INTO mailingdb(mail) SELECT * FROM ( SELECT ? AS mail) AS ifexists WHERE NOT EXISTS (SELECT mail FROM mailingdb WHERE mail = ?) LIMIT 1", mail, mail)

	if err != nil {
		fmt.Println(err)
	}
	return err
}
