package main;

import(
	"os"
	"log"
	"io/ioutil"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)
var (tag string)
func fcheck(string check,Conn db) bool{
	stm,err:=db.Query("SELECT * FROM tree WHERE tag =(?)")
	err!=nil{
		log.Fatal(err)
	}
	defer stm.Close()
	for stm.Next(){
		err:=stm.Scan(&tag)
		if err!=nil{
			log.Fatal(err)
		}
		if check==tag {
			return true
		}
	}
	return false
}
func sql(tag string, Conn db) bool{
	if err!=nil {
		log.Fatal(err)
		return false
	}
	defer db.Close()
	err=db.Query("INSERT INTO FILE (name) VALUES(?)",1,tag)
	if err!=nil{
		log.Fatal(err)
		return false
	}
	return true
}
func main(){
	arg:=os.Args[1:]
	if (len(arg)==0) || (len(arg)>1) {
		fmt.Println("usage: mst <filename>")
		return;
	}
	tag:=arg[0]
	if ( (tag[len(tag)-1]=='\\') || (tag[len(tag)-1]=='/') ){
		fmt.Println("error: cannot use / or \\ in filename")
		return
	}
	if fcheck(tag){
		fmt.Println("error: tag exists")
		return
	}
	if err !=nil{
		fmt.Println("error: file cannot be created check permissions")
	}
	os.Chmod(fn,0777)
	db,err := sql.Open("mysql","user:pw@tcp(ip:port)/dbname")
	if err!=nil{
		fmt.Println("error could not connnect to sql check address")
		return
	}
	sql(fn,db)
}
