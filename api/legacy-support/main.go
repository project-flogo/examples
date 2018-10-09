package main

import (
	"context"

	"github.com/TIBCOSoftware/flogo-contrib/activity/log"
	"github.com/TIBCOSoftware/flogo-contrib/trigger/rest"
	"github.com/project-flogo/core/api"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/engine"
	"github.com/project-flogo/core/support/logger"
	"github.com/project-flogo/legacybridge"
)

//go:generate go run $GOPATH/src/github.com/project-flogo/legacybridge/gen $GOPATH

func main() {

	app := myApp()

	e, err := api.NewEngine(app)

	if err != nil {
		logger.Error(err)
		return
	}

	engine.RunEngine(e)
}

func myApp() *api.App {
	app := api.NewApp()

	restTrg := legacybridge.GetTrigger(&rest.RestTrigger{})
	trg := app.NewTrigger(restTrg, map[string]interface{}{"port": 8080})
	trg.NewHandler(map[string]interface{}{"method": "GET", "path": "/blah/:num"}, RunActivities)

	return app
}

func RunActivities(ctx context.Context, input map[string]interface{}) (map[string]interface{}, error) {

	msg, _ := coerce.ToString(input["pathParams"])
	in := map[string]interface{}{"message": msg}

	logAct := legacybridge.GetActivity(&log.LogActivity{})
	_, err := api.EvalActivity(logAct, in)

	if err != nil {
		return nil, err
	}

	response := make(map[string]interface{})

	response["id"] = "123"
	response["amount"] = "1"
	response["balance"] = "500"
	response["currency"] = "USD"

	ret := make(map[string]interface{})
	ret["code"] = 200
	ret["data"] = response

	return ret, nil
}
