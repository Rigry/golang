package main
 
import (
    "fmt"
    "log"
    "time"
 
	"github.com/tbrandon/mbserver"
)
 
type Exception uint8

func main() {
 
    serv := mbserver.NewServer()
    err := serv.ListenTCP("127.0.0.1:502")
    if err != nil {
        log.Printf("%v\n", err)
    }
    defer serv.Close()
 
    
    for {
 
        fmt.Println(serv.Coils[1])
        fmt.Println(serv.HoldingRegisters[0])
		time.Sleep(1 * time.Second)
	}

}