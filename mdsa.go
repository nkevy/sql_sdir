package main

import (
	"fmt"
	"os"
	"runtime"
	"database/sql"
)
import _"github.com/go-sql-driver/mysql"
func pathSep() string{
	if runtime.GOOS == "windows" {
		return string('\\')
	}else{
		return string('/')
	}
}
func sql(fn) bool{
	db,err:=sql.Open("mysql","user:pw@tcp(ip:port)/dbname")
	if err!=nil{
		fmt.Println("error: could not connect to sql database")
		return false
	}
	defer db.Close()
	stmt:= "INSERT INTO TREE (name) VALUES(?)"
	err=db.Query(stmt,1,fn)
	if err!={
		fmt.Println("error: cannot insert file into sql")
		return false
	}
}
func main(){
	if (len(os.Args)>1)&&os.Args[1:][0][0]!='/'&&os.Args[1:][0][0]!='\\'{
		wd,err:=os.Getwd()
		if err!=nil{
			fmt.Println("cannot get working dir, check permissions")
			return;
		}
		fn:=wd+pathSep()+os.Args[1:][0]
		if '/'==fn[len(fn)-1]{
			fmt.Println("usage: mdsa <dir cannot end with />")
			return;
		}
		if _,err:=os.Stat(fn); os.IsNotExist(err){
			err:=os.Mkdir(fn,0000)
			if err!=nil{
				fmt.Println("mkdir error run as elivated user")
				fmt.Println(err)
				return;
			}
		}else{
			fmt.Println("dir name exists")
		}
		os.Chmod(fn,0777)
		sql(fn)
	}else{
		fmt.Println("usage: mdsa <dir name to be created>")
	}
}
