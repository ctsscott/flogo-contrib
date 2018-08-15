package jsonparse

import (
	"strconv"

	"github.com/Jeffail/gabs"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

// log is the default package logger
var activityLog = logger.GetLogger("activity-flogo-jsonparse")

const (
	jsonPath1   = "jsonPath1"
	jsonString1 = "jsonString1"
	jsonType1   = "jsonType1"
	ovOutput1   = "output1"

	jsonPath2   = "jsonPath2"
	jsonString2 = "jsonString2"
	jsonType2   = "jsonType2"
	ovOutput2   = "output2"

	jsonPath3   = "jsonPath3"
	jsonString3 = "jsonString3"
	jsonType3   = "jsonType3"
	ovOutput3   = "output3"
)

func init() {
	activityLog.SetLogLevel(logger.DebugLevel)
}

// MYJSONParseActivity is a stub for your Activity implementation
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

	// Set 1
	jsonP1 := context.GetInput(jsonPath1).(string)
	jsonS1 := context.GetInput(jsonString1).(string)
	jsonType1 := context.GetInput(jsonType1).(string)

	activityLog.Debugf("Parsing -- %s --for value %s", jsonS1, jsonP1)
	jsonParsed, err := gabs.ParseJSON([]byte(jsonS1))
	var value1 string
	var ok bool
	// var ok bool
	switch jsonType1 {

	case "boolean":
		var boolVal bool
		boolVal, ok = jsonParsed.Path(jsonP1).Data().(bool)
		value1 = strconv.FormatBool(boolVal)
	default:
		value1, ok = jsonParsed.Path(jsonP1).Data().(string)

	}

	if !ok {
		activityLog.Error("Parse failed")
		context.SetOutput(ovOutput1, "Parse failed for -- "+jsonP1+" -- in string -- "+jsonS1)
	}

	activityLog.Debugf("Parsed value from json: %s", value1)
	if value1 != "" {
		context.SetOutput(ovOutput1, value1)
	}

	// Set 2
	jsonP2 := context.GetInput(jsonPath2).(string)
	jsonS2 := context.GetInput(jsonString2).(string)
	jsonType2 := context.GetInput(jsonType2).(string)

	activityLog.Debugf("Parsing -- %s --for value %s", jsonS2, jsonP2)
	jsonParsed, err = gabs.ParseJSON([]byte(jsonS2))
	var value2 string
	// var ok bool
	switch jsonType2 {
	case "boolean":
		var boolVal bool
		boolVal, ok = jsonParsed.Path(jsonP2).Data().(bool)
		value2 = strconv.FormatBool(boolVal)
	default:
		value2, ok = jsonParsed.Path(jsonP2).Data().(string)

	}

	if !ok {
		activityLog.Error("Parse failed")
		context.SetOutput(ovOutput2, "Parse failed for -- "+jsonP2+" -- in string -- "+jsonS2)
	}

	activityLog.Debugf("Parsed value from json: %s", value2)
	if value2 != "" {
		context.SetOutput(ovOutput2, value2)
	}

	// Set 3
	jsonP3 := context.GetInput(jsonPath3).(string)
	jsonS3 := context.GetInput(jsonString3).(string)
	jsonType3 := context.GetInput(jsonType3).(string)

	activityLog.Debugf("Parsing -- %s --for value %s", jsonS3, jsonP3)
	jsonParsed, err = gabs.ParseJSON([]byte(jsonS3))
	var value3 string
	// var ok bool
	switch jsonType3 {
	case "boolean":
		var boolVal bool
		boolVal, ok = jsonParsed.Path(jsonP3).Data().(bool)
		value3 = strconv.FormatBool(boolVal)
	default:
		value3, ok = jsonParsed.Path(jsonP3).Data().(string)

	}

	if !ok {
		activityLog.Error("Parse failed")
		context.SetOutput(ovOutput3, "Parse failed for -- "+jsonP3+" -- in string -- "+jsonS3)
	}

	activityLog.Debugf("Parsed value from json: %s\n", value3)
	if value3 != "" {
		context.SetOutput(ovOutput3, value3)
	}

	return true, nil
}
