package models

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func Connection() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg := mysql.Config{
		User:      os.Getenv("DB_USER"),
		Passwd:    os.Getenv("DB_PASSWORD"),
		Net:       "tcp",
		Addr:      os.Getenv("DB_HOST"),
		DBName:    os.Getenv("DB_DATABASE"),
		ParseTime: true,
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func FetchAttendances(from string, to string, userId string) []Attendance {
	db := Connection()
	defer db.Close()

	sql := `
		SELECT
			*
		FROM
			Attendance
		WHERE
			date BETWEEN ? AND ?
		AND
			userId = ?;
		`
	rows, err := db.Query(sql, from, to, userId)
	if err != nil {
		return nil
	}

	var attendances []Attendance
	for rows.Next() {
		var atd Attendance
		rows.Scan(
			&atd.ID,
			&atd.Date,
			&atd.InTime,
			&atd.OutTime,
			&atd.BreakHour,
			&atd.UserID,
		)
		attendances = append(attendances, atd)
	}

	return attendances
}

func RegistrateWorkIn(data *Attendance) {
	db := Connection()
	defer db.Close()

	sql := `
		INSERT INTO
			Attendance
		values(
			?
			, ?
			, ?
			, null
			, null
			, ?
		);
	`
	dbh, err := db.Prepare(sql)
	if err != nil {
		// TODO Error.
		fmt.Println(err.Error())
		return
	}
	_, err = dbh.Exec(
		data.ID,
		data.Date,
		data.InTime,
		data.UserID,
	)
	if err != nil {
		// TODO Error.
		fmt.Println(err.Error())
		return
	}
}

func RegistrateWorkOut(id string) {
	db := Connection()
	defer db.Close()

	sql := `
		UPDATE 
			Attendance
		SET
			outTime = ?
		WHERE
			id = ?;
	`
	dbh, err := db.Prepare(sql)
	if err != nil {
		// TODO Error.
		fmt.Println(err.Error())
		return
	}
	_, err = dbh.Exec(
		time.Now(),
		id,
	)

	if err != nil {
		// TODO Error.
		fmt.Println(err.Error())
		return
	}
}
