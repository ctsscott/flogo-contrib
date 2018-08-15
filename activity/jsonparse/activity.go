package jsonparse

import (
	"github.com/Jeffail/gabs"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

// log is the default package logger
var activityLog = logger.GetLogger("activity-flogo-log")

const (
	jsonPath   = "jsonPath"
	jsonString = "jsonString"
	ovOutput   = "output"
)

// MyJsonParseActivity is a stub for your Activity implementation
type MYJSONParseActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &MYJSONParseActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *MYJSONParseActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *MYJSONParseActivity) Eval(context activity.Context) (done bool, err error) {

	jsonP := context.GetInput(jsonPath).(string)
	jsonS := context.GetInput(jsonString).(string)
	activityLog.Infof("Parsing -- %s --for value %s\n", jsonS, jsonP)
	jsonParsed, err := gabs.ParseJSON([]byte(jsonS))
	var value string
	// var ok bool
	value, _ = jsonParsed.Path(jsonP).Data().(string)
	activityLog.Infof("Parsed value from json: %s\n", value)
	context.SetOutput(ovOutput, value)

	return true, nil
}
