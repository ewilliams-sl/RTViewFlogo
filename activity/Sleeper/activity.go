package Sleeper

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"time"
	"math/rand"

)

// Constants
const (
	spMinMillis = "Min"
	spRndMillis = "Rnd"
)



// MyActivity is a stub for your Activity implementation
type MyActivity struct {
	metadata *activity.Metadata
}


func makeTimestamp() int64 {
    return time.Now().UnixNano() / int64(time.Millisecond)
}


func sleepFor(iLow int, iRand int) int{
		iSleep := (iLow + rand.Intn(iRand))
		time.Sleep(time.Duration(iSleep) * time.Millisecond)
		return iSleep
}

func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &MyActivity{metadata: metadata}
}

func (a *MyActivity) Metadata() *activity.Metadata {
	return a.metadata
}

func (a *MyActivity) Eval(context activity.Context) (done bool, err error) {

	fMin := context.GetInput(spMinMillis).(float64)
	fRnd := context.GetInput(spRndMillis).(float64)
	
	iMin := int(fMin)
	iRnd := int(fRnd)
	
	sleepFor(iMin, iRnd)

	return true, nil
	
}
