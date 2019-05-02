package testify_stats

import (
	"fmt"
	"testing"
	"time"
)

var assertions = 0
var failures = 0
var tests = map[string]bool{}

func AddAssert(t *testing.T, result bool) bool {
	assertions++
	tests[t.Name()] = true
	if !result {
		failures++
	}

	return result
}

func Run(m *testing.M) (status int) {
	startTime := time.Now()
	status = m.Run()
	runDuration := time.Now().Sub(startTime)

	if failures == 0 {
		fmt.Printf("Tests: %d, Assertions: %d, Time: %v\n",
			len(tests), assertions, runDuration)
	} else {
		fmt.Printf("Tests: %d, Assertions: %d, Failures: %d, Time: %v\n",
			len(tests), assertions, failures, runDuration)
	}

	return status
}
