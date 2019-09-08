package main;

import (
	"fmt"
	"os"
	"database/sql"
)
import _"github.com/go-sql-driver/mysql"

func sql(n){
	db,err:=sql.Open("mysql","user:pw@tcp(ip:port)/dbname")
	if err!=nil{
		fmt.Println("error: could not connect to sql database")
	}
	defer db.Close()
	stmt:= "SELECT * FROM FILE WHERE name=(?)"
	row,err=db.Query(stmt,1,n)
	if err!=nil{
		fmt.Println("error: cannot search name")
	}
	defer row.Close()
	var( name string )
	for row.Next(){
		err:=row.Scan(&name)
		if err!=nil{
			fmt.Println("error: empty result?")
		}
		fmt.Println(name)
	}
}
func main(){
	args:=os.Args[1:]
	if (len(args)!=1){
		fmt.Println("usage: ssearch <file or dir name>")
	}
	sql(args[0])
}
