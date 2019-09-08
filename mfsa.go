package main;

import(
	"os"
	"fmt"
	"io/ioutil"
	"database/sql"
)
import _"github.com/go-sql-driver/mysql"

func fcheck(filename string) bool{
	ret,err:=os.Stat(filename)
	if err!=nil{
		return false;
	}
	return !ret.IsDir();
}
func sql(fn) bool{
	db,err:=sql.Open("mysql","user:pw@tcp(ip:port)/dbname")
	if err!=nil{
		fmt.Println("error: could not connect to sql database")
		return false
	}
	defer db.Close()
	stmt:= "INSERT INTO FILE (name) VALUES(?)"
	err=db.Query(stmt,1,fn)
	if err!={
		fmt.Println("error: cannot insert file into sql")
		return false
	}
}
func main(){
	arg:=os.Args[1:]
	if (len(arg)==0) || (len(arg)>1) {
		fmt.Println("usage: mfsa <filename>")
		return;
	}
	fn:=arg[0]
	if ( (fn[len(fn)-1]=='\\') || (fn[len(fn)-1]=='/') ){
		fmt.Println("error: cannot use / or \\ in filename")
		return
	}
	if fcheck(fn){
		fmt.Println("error: file exists")
		return
	}
	err := ioutil.WriteFile(fn,[]byte(""),os.ModePerm)
	if err !=nil{
		fmt.Println("error: file cannot be created check permissions")
	}
	os.Chmod(fn,0777)
	sql(fn)
}
