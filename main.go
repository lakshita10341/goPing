package main
import ("goPing/internal/ping"
		"log"
	)
func main(){
	err:=ping.Ping("8.8.8.8")
	if err!=nil{
		log.Fatal(err)
	}
}