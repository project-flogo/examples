package main

import (
	"fmt"

	"github.com/project-flogo/core/api"
	"github.com/project-flogo/core/engine"
	"github.com/project-flogo/flow"

	_ "github.com/project-flogo/contrib/activity/log" //our flow contains a log activity, so we need to include this
	"github.com/project-flogo/contrib/trigger/rest"

)

func main() {

	app := myApp()

	e, err := api.NewEngine(app)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	engine.RunEngine(e)
}

func myApp() *api.App {
	app := api.NewApp()

	trg := app.NewTrigger(&rest.Trigger{}, &rest.Settings{Port: 8080})
	h, _ := trg.NewHandler(&rest.HandlerSettings{Method: "GET", Path: "/blah/:val"})

	settings :=  &flow.Settings{FlowURI:"file://myflow.json"}
	a, _ := h.NewAction(&flow.FlowAction{}, settings)
	a.SetInputMappings("in=$.pathParams.val")

	return app
}
