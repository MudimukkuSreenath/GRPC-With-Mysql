package mysql

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Food struct {
	Id         int32
	Name       string
	Price      float32
	TypeId     int32
	CreateTime int64 `gorm:"column:createtime"`
}

func connMysql() *gorm.DB {
	//username := "root"
	//password := "sreenath15"
	//address := localhost(":3306")
	//dbname := "sreenath"
	//timeout = "10s"

	//dsn := fmt.Sprint("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", username, password, dbname)

	var db, err = gorm.Open("mysql", "root:sreenath15@tcp(127.0.0.1:3306)/sreenath?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect to the database: " + err.Error())
	}
	return db

}

func InsDelUpd(op string, id int32, name string, price float32, typeId int32, createTime int64) {
	db := connMysql()
	defer db.Close()

	switch op {
	case "insert":
		food := &Food{
			id,
			name,
			price,
			typeId,
			createTime,
		}
		db.Create(food)
	case "delete":
		food := &Food{
			id,
			name,
			price,
			typeId,
			createTime,
		}
		db.Delete(&food)
	case "update":
		food := &Food{Id: id}
		db.Model(&food).Update(Food{Name: name, Price: price, TypeId: typeId, CreateTime: createTime})

	}

}

func Select(table string, columns string, condition string) string {
	db := connMysql()
	defer db.Close()
	var foods []Food
	db.Where(condition).Select(columns).Find(&foods)
	response := fmt.Sprintf("%v", foods)
	return response
}

func ExecSql(sql string) {
	db := connMysql()
	defer db.Close()
	db.Exec(sql)
}
