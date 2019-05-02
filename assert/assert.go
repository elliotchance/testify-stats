package assert

import (
	testify_stats "github.com/elliotchance/testify-stats"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func FailNow(t *testing.T, failureMessage string, msgAndArgs ...interface{}) bool {
	return testify_stats.AddAssert(t, assert.FailNow(t, failureMessage, msgAndArgs...))
}

func Fail(t *testing.T, failureMessage string, msgAndArgs ...interface{}) bool {
	return testify_stats.AddAssert(t, assert.Fail(t, failureMessage, msgAndArgs...))
}

func Implements(t *testing.T, interfaceObject interface{}, object interface{}, msgAndArgs ...interface{}) bool {
	return testify_stats.AddAssert(t, assert.Implements(t, interfaceObject, object, msgAndArgs...))
}

func IsType(t *testing.T, expectedType interface{}, object interface{}, msgAndArgs ...interface{}) bool {
	return testify_stats.AddAssert(t, assert.IsType(t, expectedType, object, msgAndArgs...))
}

func Equal(t *testing.T, expected, actual interface{}, msgAndArgs ...interface{}) bool {
	return testify_stats.AddAssert(t, assert.Equal(t, expected, actual, msgAndArgs...))
}

func EqualValues(t *testing.T, expected, actual interface{}, msgAndArgs ...interface{}) bool {
	return testify_stats.AddAssert(t, assert.EqualValues(t, expected, actual, msgAndArgs...))
}

func Exactly(t *testing.T, expected, actual interface{}, msgAndArgs ...interface{}) bool {
	return testify_stats.AddAssert(t, assert.Exactly(t, expected, actual, msgAndArgs...))
}

func NotNil(t *testing.T, object interface{}, msgAndArgs ...interface{}) bool {
	return testify_stats.AddAssert(t, assert.NotNil(t, object, msgAndArgs...))
}

func Nil(t *testing.T, object interface{}, msgAndArgs ...interface{}) bool {
	return testify_stats.AddAssert(t, assert.Nil(t, object, msgAndArgs...))
}

func Empty(t *testing.T, object interface{}, msgAndArgs ...interface{}) bool {
	return testify_stats.AddAssert(t, assert.Empty(t, object, msgAndArgs...))
}

func NotEmpty(t *testing.T, object interface{}, msgAndArgs ...interface{}) bool {
	return testify_stats.AddAssert(t, assert.NotEmpty(t, object, msgAndArgs...))
}

func Len(t *testing.T, object interface{}, length int, msgAndArgs ...interface{}) bool {
	return testify_stats.AddAssert(t, assert.Len(t, object, length, msgAndArgs...))
}

func True(t *testing.T, value bool, msgAndArgs ...interface{}) bool {
	return testify_stats.AddAssert(t, assert.True(t, value, msgAndArgs...))
}

func False(t *testing.T, value bool, msgAndArgs ...interface{}) bool {
	return testify_stats.AddAssert(t, assert.False(t, value, msgAndArgs...))
}

func NotEqual(t *testing.T, expected, actual interface{}, msgAndArgs ...interface{}) bool {
	return testify_stats.AddAssert(t, assert.NotEqual(t, expected, actual, msgAndArgs...))
}

func Contains(t *testing.T, s, contains interface{}, msgAndArgs ...interface{}) bool {
	return testify_stats.AddAssert(t, assert.Contains(t, s, contains, msgAndArgs...))
}

func NotContains(t *testing.T, s, contains interface{}, msgAndArgs ...interface{}) bool {
	return testify_stats.AddAssert(t, assert.NotContains(t, s, contains, msgAndArgs...))
}

func Subset(t *testing.T, list, subset interface{}, msgAndArgs ...interface{}) (ok bool) {
	return testify_stats.AddAssert(t, assert.Subset(t, list, subset, msgAndArgs...))
}

func NotSubset(t *testing.T, list, subset interface{}, msgAndArgs ...interface{}) (ok bool) {
	return testify_stats.AddAssert(t, assert.NotSubset(t, list, subset, msgAndArgs...))
}

func ElementsMatch(t *testing.T, listA, listB interface{}, msgAndArgs ...interface{}) (ok bool) {
	return testify_stats.AddAssert(t, assert.ElementsMatch(t, listA, listB, msgAndArgs...))
}

func Condition(t *testing.T, comp assert.Comparison, msgAndArgs ...interface{}) bool {
	return testify_stats.AddAssert(t, assert.Condition(t, comp, msgAndArgs...))
}

func Panics(t *testing.T, f assert.PanicTestFunc, msgAndArgs ...interface{}) bool {
	return testify_stats.AddAssert(t, assert.Panics(t, f, msgAndArgs...))
}

func PanicsWithValue(t *testing.T, expected interface{}, f assert.PanicTestFunc, msgAndArgs ...interface{}) bool {
	return testify_stats.AddAssert(t, assert.PanicsWithValue(t, expected, f, msgAndArgs...))
}

func NotPanics(t *testing.T, f assert.PanicTestFunc, msgAndArgs ...interface{}) bool {
	return testify_stats.AddAssert(t, assert.NotPanics(t, f, msgAndArgs...))
}

func WithinDuration(t *testing.T, expected, actual time.Time, delta time.Duration, msgAndArgs ...interface{}) bool {
	return testify_stats.AddAssert(t, assert.WithinDuration(t, expected, actual, delta, msgAndArgs...))
}

func InDelta(t *testing.T, expected, actual interface{}, delta float64, msgAndArgs ...interface{}) bool {
	return testify_stats.AddAssert(t, assert.InDelta(t, expected, actual, delta, msgAndArgs...))
}

func InDeltaSlice(t *testing.T, expected, actual interface{}, delta float64, msgAndArgs ...interface{}) bool {
	return testify_stats.AddAssert(t, assert.InDeltaSlice(t, expected, actual, delta, msgAndArgs...))
}

func InDeltaMapValues(t *testing.T, expected, actual interface{}, delta float64, msgAndArgs ...interface{}) bool {
	return testify_stats.AddAssert(t, assert.InDeltaMapValues(t, expected, actual, delta, msgAndArgs...))
}

func InEpsilon(t *testing.T, expected, actual interface{}, epsilon float64, msgAndArgs ...interface{}) bool {
	return testify_stats.AddAssert(t, assert.InEpsilon(t, expected, actual, epsilon, msgAndArgs...))
}

// InEpsilonSlice is the same as InEpsilon, except it compares each value from two slices.
func InEpsilonSlice(t *testing.T, expected, actual interface{}, epsilon float64, msgAndArgs ...interface{}) bool {
	return testify_stats.AddAssert(t, assert.InEpsilonSlice(t, expected, actual, epsilon, msgAndArgs...))
}

func NoError(t *testing.T, err error, msgAndArgs ...interface{}) bool {
	return testify_stats.AddAssert(t, assert.NoError(t, err, msgAndArgs...))
}

func Error(t *testing.T, err error, msgAndArgs ...interface{}) bool {
	return testify_stats.AddAssert(t, assert.Error(t, err, msgAndArgs...))
}

func EqualError(t *testing.T, theError error, errString string, msgAndArgs ...interface{}) bool {
	return testify_stats.AddAssert(t, assert.EqualError(t, theError, errString, msgAndArgs...))
}

func Regexp(t *testing.T, rx interface{}, str interface{}, msgAndArgs ...interface{}) bool {
	return testify_stats.AddAssert(t, assert.Regexp(t, rx, str, msgAndArgs...))
}

func NotRegexp(t *testing.T, rx interface{}, str interface{}, msgAndArgs ...interface{}) bool {
	return testify_stats.AddAssert(t, assert.NotRegexp(t, rx, str, msgAndArgs...))
}

func Zero(t *testing.T, i interface{}, msgAndArgs ...interface{}) bool {
	return testify_stats.AddAssert(t, assert.Zero(t, i, msgAndArgs...))
}

func NotZero(t *testing.T, i interface{}, msgAndArgs ...interface{}) bool {
	return testify_stats.AddAssert(t, assert.NotZero(t, i, msgAndArgs...))
}

func FileExists(t *testing.T, path string, msgAndArgs ...interface{}) bool {
	return testify_stats.AddAssert(t, assert.FileExists(t, path, msgAndArgs...))
}

func DirExists(t *testing.T, path string, msgAndArgs ...interface{}) bool {
	return testify_stats.AddAssert(t, assert.DirExists(t, path, msgAndArgs...))
}

func JSONEq(t *testing.T, expected string, actual string, msgAndArgs ...interface{}) bool {
	return testify_stats.AddAssert(t, assert.JSONEq(t, expected, actual, msgAndArgs...))
}
