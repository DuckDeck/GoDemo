package db

import "log"

func addUserCredential(name string, password string) error {
	stmt, err := db.Prepare("insert into user(login_name,password) values (?,?)")
	if err != nil {
		return err
	}
	stmt.Exec(name, password)
	stmt.Close()
	return nil
}

func getUserCredential(name string) (string, error) {
	stmt, err := db.Prepare("select password from user where login_name = ?")
	if err != nil {
		log.Printf("%s", err)
		return "", err
	}
	var password string
	stmt.QueryRow(name).Scan(&password)
	stmt.Close()
	return password, nil
}

func deleteUser(name string, password string) error {
	stmt, err := db.Prepare("delete from user where login_name = ? and password = ?")
	if err != nil {
		log.Printf("%s", err)
		return err
	}
	stmt.Exec(name, password)
	stmt.Close()
	return nil
}
