package dao

import (
	"database/sql"
	"errors"
	"exam/model"
	"log"
)

func OpenDb() {
	Db, err := sql.Open("mysql", "root:admin@tcp(127.0.0.1:3306)/user")
	if err != nil {
		log.Println(err)
		return
	}
	err = Db.Ping()
	if err != nil {
		log.Println(err)
		return
	}
}
func NameQuery(name string) (*sql.Rows, error) {
	Db, err := sql.Open("mysql", "root:admin@tcp(127.0.0.1:3306)/user")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	stmt, err := Db.Prepare("select * from user where name= ?")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	row, err := stmt.Query(name)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return row, nil
}
func PasswordQuery(name string) (*sql.Rows, error) {
	Db, err := sql.Open("mysql", "root:admin@tcp(127.0.0.1:3306)/user")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	stmt, err := Db.Prepare("select * from user where name= ?")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	row, err := stmt.Query(name)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	err = row.Scan()
	return row, nil
}
func ExecName(name string) error {
	Db, err := sql.Open("mysql", "root:admin@tcp(127.0.0.1:3306)/user")
	if err != nil {
		log.Println(err)
		return err
	}
	_, err = Db.Exec("insert into user name value ?", name)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
func ExecPassword(password string) error {
	Db, err := sql.Open("mysql", "root:admin@tcp(127.0.0.1:3306)/user")
	if err != nil {
		log.Println(err)
		return err
	}
	_, err = Db.Exec("insert into user password value ?", password)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func Recharge(money int, name string) error {
	Db, err := sql.Open("mysql", "root:admin@tcp(127.0.0.1:3306)/user")
	if err != nil {
		log.Println(err)
		return err
	}
	rows, err := NameQuery(name)
	if err != nil {
		log.Println(err)
	}
	var user model.User
	err = rows.Scan(&user.Id, &user.Name, &user.Password, &user.Balance)
	if err != nil {
		log.Println(err)
		return err
	}

	_, err = Db.Exec("update user set balance =? where name=?", user.Balance+money)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil

}

func ReduceMoney(money int, name string) error {
	Db, err := sql.Open("mysql", "root:admin@tcp(127.0.0.1:3306)/user")
	if err != nil {
		log.Println(err)
		return err
	}
	rows, err := NameQuery(name)
	if err != nil {
		log.Println(err)
	}
	var user model.User
	err = rows.Scan(&user.Id, &user.Name, &user.Password, &user.Balance)
	if err != nil {
		log.Println(err)
		return err
	}
	if user.Balance-money < 0 {
		return err
	}
	_, err = Db.Exec("update user set balance =? where name=?", user.Balance-money)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil

}
func Checkmoney(name string) error {

	rows, err := NameQuery(name)
	if err != nil {
		log.Println(err)
	}
	var user model.User
	err = rows.Scan(&user.Id, &user.Name, &user.Password, &user.Balance)
	if user.Balance < 0 {
		return errors.New("钱低于0")
	}
	return nil

}