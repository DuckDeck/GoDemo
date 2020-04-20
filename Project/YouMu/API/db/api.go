package db

import (
	"database/sql"
	"log"
	"socket/Project/YouMu/API/defs"

)

func addUserCredential(name string, password string) error {
	stmt, err := db.Prepare("insert into user(login_name,password) values (?,?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(name, password)
	if err != nil {
		return err
	}
	defer stmt.Close()
	return nil
}

func getUserCredential(name string) (string, error) {
	stmt, err := db.Prepare("select password from user where login_name = ?")
	if err != nil {
		log.Printf("%s", err)
		return "", err
	}
	var password string
	err = stmt.QueryRow(name).Scan(&password)
	if err != nil && err != sql.ErrNoRows {
		return "", err
	}
	defer stmt.Close()
	return password, nil
}

func deleteUser(name string, password string) error {
	stmt, err := db.Prepare("delete from user where login_name = ? and password = ?")
	if err != nil {
		log.Printf("%s", err)
		return err
	}
	_, err = stmt.Exec(name, password)
	if err != nil {
		return err
	}
	defer stmt.Close()
	return nil
}

func addNewVideo(aid int, name string) (*defs.VideoInfo error)) {
	var uuid =  uuid.Must(uuid.NewV4()).String()
	t := time.now()
}
