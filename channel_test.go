package golang_goroutines

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel :=make(chan string)

	defer close(channel)

	go func(){
		time.Sleep(2 *time.Second)
		// Mengirimkan ke channel
		channel <- "Muhammad Irfan Zidni"
		fmt.Println("Selesai mengirim data ke channel")

	}()
	data:= <-channel
	fmt.Println(data)

	time.Sleep(5*time.Second)

}

// Channel as parameter

func GiveMeResponse(channel chan string){
	time.Sleep(2*time.Second)
	channel <- "Muhammad Irfan Zidni"

}

func TestChannelAsParameter(t *testing.T){
	channel := make(chan string)
	defer close(channel)
	
    go GiveMeResponse(channel)
    data := <-channel
    fmt.Println(data)

    time.Sleep(5*time.Second)


}