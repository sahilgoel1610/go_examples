package main

import (
"fmt"
)

func main() {
	btn := MakeButton()

	handlerOne := make(chan string)
	handlerTwo := make(chan string)

	btn.AddEventListener("click",handlerOne)
	btn.AddEventListener("click", handlerTwo)
	go func () {
		for {
			msg := <- handlerOne
			fmt.Println("handler One : ", msg)
		}
	}()

	go func () {
		for {
			msg := <- handlerTwo
			fmt.Println("handler Two : ", msg)
		}
	}()

	btn.TriggerEvent("click","Buttong Clicked")
	btn.RemoveListener("click", handlerTwo)
	btn.TriggerEvent("click", "Buttong clicked again")
	fmt.Scanln()
}


type Button struct{
	eventListeners map[string][]chan string 
}


func MakeButton() *Button{
	result := new(Button)
	result.eventListeners = make(map[string][]chan string)
	return result
	
}


func (b Button) AddEventListener(event string, responseChannel chan string) {
	 _,present := b.eventListeners[event]
	 if present {
	 	b.eventListeners[event] = append(b.eventListeners[event], responseChannel)
	 } else{
	 	b.eventListeners[event] = []chan string{responseChannel}
	 }
}


func (b Button) RemoveListener(event string, removeChannel chan string) {
	_, present := b.eventListeners[event]

	if present {
		for idx,_ := range b.eventListeners[event]{
			if b.eventListeners[event][idx] == removeChannel {
				b.eventListeners[event] = append(b.eventListeners[event][:idx],b.eventListeners[event][idx+1:]...)
				break
			}
		}
	}
}

func (b Button) TriggerEvent(event string, response string ) {
	_,present := b.eventListeners[event]

	if present{
		for _, handler := range b.eventListeners[event]{
			go func (handler chan<- string) {
				handler <- response
			}(handler)

		}

	}
}

