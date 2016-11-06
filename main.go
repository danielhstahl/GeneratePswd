package main
import (
    "golang.org/x/crypto/bcrypt"
    "fmt"
    "os"
)
func main(){
    //args:=os.Args[1:]
    pswd:=os.Args[1]
    bytePassword:=[]byte(pswd)
    hashedPassword, err:=bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
    if err!=nil{
        panic(err)

    }
    fmt.Println(string(hashedPassword))
}