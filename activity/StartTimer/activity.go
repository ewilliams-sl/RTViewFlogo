package StartTimer

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"time"
	"fmt"
)

// Constants
const (
	result  = "result"
)


// MyActivity is a stub for your Activity implementation
type MyActivity struct {
	metadata *activity.Metadata
}


func makeTimestamp() int64 {
    return time.Now().UnixNano() / int64(time.Millisecond)
}





// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &MyActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *MyActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *MyActivity) Eval(context activity.Context) (done bool, err error) {
	result := "starttime"
	iStartTime := makeTimestamp()
	context.SetOutput(result, iStartTime)

	// Signal to the Flogo engine that the activity is completed
	return true, nil
	
}