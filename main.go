package main

import (
	"fmt"

	"business/auth/common/mq"
	"business/user/impls"

	"github.com/mz-eco/mz/app"
	"github.com/mz-eco/mz/http"
)

type Application struct {
}

func (m *Application) Run(args []string) {

	service := http.NewService()
	service.AddAccessControlHandlers(impls.AccessControlHandlers)
	service.AddHandlers(impls.Handlers)

	//订阅
	subscriber, err := mq.NewSubscriber()
	if err != nil {
		panic(err)
		return
	}
	subscriber.Run()

	service.Run()

}

func (m *Application) Flags(flags *app.Flags) {
	//TODO: application flags
}

func (m *Application) GetName() string {
	return "user"
}

func main() {
	err := app.Main(&Application{})

	if err != nil {
		fmt.Println(err)
	}
}
