package main
import "fmt"

func main () {
	messages := make(chan []string);

	go func(){
		messages <- []string{"hello", "ping"}
	}()
	
	msg := <-messages
	fmt.Println(msg);
}