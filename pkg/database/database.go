package database

import (
	"database/sql"
	"fmt"
	"github.com/JonasEnglund/tgBotTest/pkg/consts"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func ConnectDb() (*sql.DB, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", consts.Db_user, consts.Db_password, consts.Db_host, consts.Db_name))
	//defer db.Close()

	if err != nil {
		return db, err
	}

	log.Println("Successfully connected to Database")
	return db, nil
}

func InsertData(db *sql.DB, person *consts.Person) error {
	insert, err := db.Query(fmt.Sprintf("INSERT INTO test VALUES('%s', '%s', '%s', '%s', '%s')", person.Id, person.FirstName, person.SecondName, person.Role, person.Points))
	if err != nil {
		return err
	}
	defer insert.Close()
	return nil
}

func ShowData(db *sql.DB) error {
	var persons []consts.Person

	show, err := db.Query(fmt.Sprintf("SELECT * FROM %s", consts.Db_table))
	defer show.Close()
	if err != nil {
		return err
	}

	for show.Next() {
		p := consts.Person{}
		err := show.Scan(&p.Id, &p.FirstName, &p.SecondName, &p.Role, &p.Points)
		if err != nil {
			return err
		}
		persons = append(persons, p)
	}

	for i := range persons {
		fmt.Println(persons[i])
	}

	return nil
}

func CheckUserExistance(Id int, db *sql.DB) (bool, error) {
	show, err := db.Query(fmt.Sprintf("SELECT id FROM %s", consts.Db_table))
	defer show.Close()
	if err != nil {
		return false, err
	}
	for show.Next() {
		p := consts.Person{}
		err := show.Scan(&p.Id)
		if err != nil {
			return false, err
		}
		if p.Id == Id {
			return true, nil
		}
	}

	return false, nil
}
