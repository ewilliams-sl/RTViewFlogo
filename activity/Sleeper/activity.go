package Sleeper

import (
	"encoding/json"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"strings"
	"time"
	"fmt"
)

// Constants
const (
	spMinMillis = "Min"
	spRndMillis = "Rnd"
	params  = "params"
	result  = "result"
)

// log is the default package logger which we'll use to log
var log = logger.GetLogger("activity-setQoS")

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

	iMin := context.GetInput(spMinMillis).(int64)
	iRnd := context.GetInput(spRndMillis).(int64)
	
	sleepFor(iMn, iRnd)

	// Signal to the Flogo engine that the activity is completed
	return true, nil
	
}

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	m := make(map[string]int)
	for _, word := range words {
		m[word] += 1
	}
	return m
}

func split(tosplit string, sep rune) []string {
	var fields []string

	last := 0
	for i, c := range tosplit {
		if c == sep {
			// Found the separator, append a slice
			fields = append(fields, string(tosplit[last:i]))
			last = i + 1
		}
	}

	// Don't forget the last field
	fields = append(fields, string(tosplit[last:]))

	return fields
}
