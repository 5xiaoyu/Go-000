package Week02

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type Person struct {
	Id int `db:"id"`
	UserName string `db:"user_name"`
	Sex int `db:"sex"`
	Phone string `db:"phone"`

}
//var Db *sqlx.DB

func MysqlInstance (dsn string) (Db *sqlx.DB, err error) {
	Db,err = sqlx.Open("mysql", dsn)
	return
}

func GetPersonById (id int ) (person *Person,err error){
	//var person *Person
	db,err := MysqlInstance("root:root@tcp(127.0.0.1:3306)/test")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}

	err = db.Select(&person, "select id, user_name, sex, phone from person where id=?", id)
	if err != nil {
		fmt.Println("exec failed, ", err)
		//return
	}
	return
}