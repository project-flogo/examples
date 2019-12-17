# Run Actions using Go-API.

This short tutorial walks through how to run Flogo Actions using the Go-API provided by the flogo core.

First Import the action you need to run .

```go 
import (
    ...
    "github.com/project-flogo/core/api"
    "github.com/project-flogo/flow"
    ...
)
```

Then we need to initalize the Flogo app using Go-API

```go
app := api.NewApp()
```

Setup the trigger and handler that invokes your action.   In this case will setup a REST trigger on port 8080.  We'll use the method GET /blah/:num to kick off the action.
```go

trg := app.NewTrigger(&rest.Trigger{}, &rest.Settings{Port: 8080})
h, _ := trg.NewHandler(&rest.HandlerSettings{Method: "GET", Path: "/blah/:num"})

```

Lets setup Flow Action that runs the flow from our file myflow.json.  We'll then associate that action with our handler.
```go

settings :=  &flow.Settings{FlowURI:"file://myflow.json"}
a, _ h.NewAction(&flow.FlowAction{}, settings)

```

Now lets map the path parameter from the handler to input of our flow.
```go

a.SetInputMappings("in=$.pathParams.val")

```

####Full Example
```go
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
```
