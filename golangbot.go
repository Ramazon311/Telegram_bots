package main

import (
	//"example.com/m/projects/go/pkg/mod/golang.org/x/vuln@v0.0.0-20220503210553-a5481fb0c8be/client"
	//"example.com/m/projects/go/pkg/mod/golang.org/x/text@v0.3.7/message"
	"github.com/Squirrel-Network/gobotapi"
	"github.com/Squirrel-Network/gobotapi/methods"
	"github.com/Squirrel-Network/gobotapi/types"
)
func main(){
	client:=gobotapi.NewClient("5579790457:AAHJ1eQtHoYFCHgRhuYBbpPmvhDi27hKVtg")
	client.OnMessage(func(message types.Message){
		client.Invoke(&methods.SendMessage{
			ChatID: message.Chat.ID,
			Text: "Hello Ramazon",
		})
	})
	client.Run()

}