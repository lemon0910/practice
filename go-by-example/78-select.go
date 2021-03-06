package main
import "time"
import "fmt"

func main(){
	c1 := make(chan string)
	c2 := make(chan string)

	go func(){
		time.Sleep(time.Second * 1)
		c1<- "hello"
	}()
	go func(){
		time.Sleep(time.Second * 1)
		c1<- "world"
	}()

	for i := 0; i < 2; i++ {
		select{
		case msg1 := <- c1:
			fmt.Println("received", msg1)
		case msg2 := <- c2:
			fmt.Println("received", msg2)
		}
	}
}
