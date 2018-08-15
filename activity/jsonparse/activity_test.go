package jsonparse

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/TIBCOSoftware/flogo-contrib/action/flow/test"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
)

var activityMetadata *activity.Metadata

func getActivityMetadata() *activity.Metadata {

	if activityMetadata == nil {
		jsonMetadataBytes, err := ioutil.ReadFile("activity.json")
		if err != nil {
			panic("No Json Metadata found for activity.json path")
		}

		activityMetadata = activity.NewMetadata(string(jsonMetadataBytes))
	}

	return activityMetadata
}

func TestCreate(t *testing.T) {

	act := NewActivity(getActivityMetadata())

	if act == nil {
		t.Error("Activity Not Created")
		t.Fail()
		return
	}
}

// Test success
func TestEval(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	tc.SetInput("jsonPath1", "Publisher.IsDocker")
	tc.SetInput("jsonString1", "{\"Publisher\":{\"Timestamp\":\"15/08/2018 03:14:58.958 +0000\",\"IsDocker\":true,\"Domain\":\"CTSL\",\"RvService\":\"27496\",\"RvDaemon\":\"cdcxpd0698.con-way.com:7474\",\"RvNetwork\":\";\",\"Host\":\"48336f75b07a\"},\"Event\":{\"MicroagentID\":\"COM.TIBCO.ADAPTER.bwengine.CTSL.ConwayRamNotification.ConwayRamNotification\",\"Agent\":\"cdcxpd0578\",\"NewPID\":28916,\"Alive\":true}}")
	tc.SetInput("jsonType1", "boolean")
	tc.SetInput("jsonPath2", "Publisher.Domain")
	tc.SetInput("jsonString2", "{\"Publisher\":{\"Timestamp\":\"15/08/2018 03:14:58.958 +0000\",\"IsDocker\":true,\"Domain\":\"CTSL\",\"RvService\":\"27496\",\"RvDaemon\":\"cdcxpd0698.con-way.com:7474\",\"RvNetwork\":\";\",\"Host\":\"48336f75b07a\"},\"Event\":{\"MicroagentID\":\"COM.TIBCO.ADAPTER.bwengine.CTSL.ConwayRamNotification.ConwayRamNotification\",\"Agent\":\"cdcxpd0578\",\"NewPID\":28916,\"Alive\":true}}")
	tc.SetInput("jsonType2", "string")

	tc.SetInput("jsonPath3", "Event.MicroagentID")
	tc.SetInput("jsonString3", "{\"Publisher\":{\"Timestamp\":\"15/08/2018 03:14:58.958 +0000\",\"IsDocker\":true,\"Domain\":\"CTSL\",\"RvService\":\"27496\",\"RvDaemon\":\"cdcxpd0698.con-way.com:7474\",\"RvNetwork\":\";\",\"Host\":\"48336f75b07a\"},\"Event\":{\"MicroagentID\":\"COM.TIBCO.ADAPTER.bwengine.CTSL.ConwayRamNotification.ConwayRamNotification\",\"Agent\":\"cdcxpd0578\",\"NewPID\":28916,\"Alive\":true}}")
	tc.SetInput("jsonType3", "string")

	done, err := act.Eval(tc)
	if !done {
		fmt.Println(err)
	}

	//check result attr
	a, _ := json.Marshal(tc.GetOutput("output1"))
	fmt.Println(string(a))

	b, _ := json.Marshal(tc.GetOutput("output2"))
	fmt.Println(string(b))

	c, _ := json.Marshal(tc.GetOutput("output3"))
	fmt.Println(string(c))

	//if strings.Compare(string(b), "\"value\"") != 0 {
	//	t.FailNow()
	//}

}

/* Test success
func TestFail(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	tc.SetInput("jsonPath", "test1111")
	tc.SetInput("jsonString", "{\"test\":\"value\"}")

	done, err := act.Eval(tc)
	if !done {
		fmt.Println(err)
	}

	//check result attr
	b, _ := json.Marshal(tc.GetOutput("output"))
	fmt.Println(string(b))

	if strings.Compare(string(b), "\"value\"") != 0 {
		t.FailNow()
	}

} */
