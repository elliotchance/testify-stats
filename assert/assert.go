package assert

import (
	testify_stats "github.com/elliotchance/testify-stats"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
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

func Conditionf(t *testing.T, comp assert.Comparison, msg string, args ...interface{}) {
	testify_stats.AddAssert(t, assert.Conditionf(t, comp, msg, args))
}

func Containsf(t *testing.T, s interface{}, contains interface{}, msg string, args ...interface{}) {
	testify_stats.AddAssert(t, assert.Containsf(t, s, contains, msg, args...))
}

func DirExistsf(t *testing.T, path string, msg string, args ...interface{}) {
	testify_stats.AddAssert(t, assert.DirExistsf(t, path, msg, args...))
}

func ElementsMatchf(t *testing.T, listA interface{}, listB interface{}, msg string, args ...interface{}) {
	testify_stats.AddAssert(t, assert.ElementsMatchf(t, listA, listB, msg, args...))
}

func Emptyf(t *testing.T, object interface{}, msg string, args ...interface{}) {
	testify_stats.AddAssert(t, assert.Emptyf(t, object, msg, args...))
}

func EqualErrorf(t *testing.T, theError error, errString string, msg string, args ...interface{}) {
	testify_stats.AddAssert(t, assert.EqualErrorf(t, theError, errString, msg, args...))
}

func EqualValuesf(t *testing.T, expected interface{}, actual interface{}, msg string, args ...interface{}) {
	testify_stats.AddAssert(t, assert.EqualValuesf(t, expected, actual, msg, args...))
}

func Equalf(t *testing.T, expected interface{}, actual interface{}, msg string, args ...interface{}) {
	testify_stats.AddAssert(t, assert.Equalf(t, expected, actual, msg, args...))
}

func Errorf(t *testing.T, err error, msg string, args ...interface{}) {
	testify_stats.AddAssert(t, assert.Errorf(t, err, msg, args...))
}

func Exactlyf(t *testing.T, expected interface{}, actual interface{}, msg string, args ...interface{}) {
	testify_stats.AddAssert(t, assert.Exactlyf(t, expected, actual, msg, args...))
}

func FailNowf(t *testing.T, failureMessage string, msg string, args ...interface{}) {
	testify_stats.AddAssert(t, assert.FailNowf(t, failureMessage, msg, args...))
}

func Failf(t *testing.T, failureMessage string, msg string, args ...interface{}) {
	testify_stats.AddAssert(t, assert.Failf(t, failureMessage, msg, args...))
}

func Falsef(t *testing.T, value bool, msg string, args ...interface{}) {
	testify_stats.AddAssert(t, assert.Falsef(t, value, msg, args...))
}

func FileExistsf(t *testing.T, path string, msg string, args ...interface{}) {
	testify_stats.AddAssert(t, assert.FileExistsf(t, path, msg, args...))
}

func HTTPBodyContainsf(t *testing.T, handler http.HandlerFunc, method string, url string, values url.Values, str interface{}, msg string, args ...interface{}) {
	testify_stats.AddAssert(t, assert.HTTPBodyContainsf(t, handler, method, url, values, str, msg, args...))
}

func HTTPBodyNotContainsf(t *testing.T, handler http.HandlerFunc, method string, url string, values url.Values, str interface{}, msg string, args ...interface{}) {
	testify_stats.AddAssert(t, assert.HTTPBodyNotContainsf(t, handler, method, url, values, str, msg, args...))
}

func HTTPErrorf(t *testing.T, handler http.HandlerFunc, method string, url string, values url.Values, msg string, args ...interface{}) {
	testify_stats.AddAssert(t, assert.HTTPErrorf(t, handler, method, url, values, msg, args...))
}

func HTTPRedirectf(t *testing.T, handler http.HandlerFunc, method string, url string, values url.Values, msg string, args ...interface{}) {
	testify_stats.AddAssert(t, assert.HTTPRedirectf(t, handler, method, url, values, msg, args...))
}

func HTTPSuccessf(t *testing.T, handler http.HandlerFunc, method string, url string, values url.Values, msg string, args ...interface{}) {
	testify_stats.AddAssert(t, assert.HTTPSuccessf(t, handler, method, url, values, msg, args...))
}

func Implementsf(t *testing.T, interfaceObject interface{}, object interface{}, msg string, args ...interface{}) {
	testify_stats.AddAssert(t, assert.Implementsf(t, interfaceObject, object, msg, args...))
}

func InDeltaMapValuesf(t *testing.T, expected interface{}, actual interface{}, delta float64, msg string, args ...interface{}) {
	testify_stats.AddAssert(t, assert.InDeltaMapValuesf(t, expected, actual, delta, msg, args...))
}

func InDeltaSlicef(t *testing.T, expected interface{}, actual interface{}, delta float64, msg string, args ...interface{}) {
	testify_stats.AddAssert(t, assert.InDeltaSlicef(t, expected, actual, delta, msg, args...))
}

func InDeltaf(t *testing.T, expected interface{}, actual interface{}, delta float64, msg string, args ...interface{}) {
	testify_stats.AddAssert(t, assert.InDeltaf(t, expected, actual, delta, msg, args...))
}

func InEpsilonSlicef(t *testing.T, expected interface{}, actual interface{}, epsilon float64, msg string, args ...interface{}) {
	testify_stats.AddAssert(t, assert.InEpsilonSlicef(t, expected, actual, epsilon, msg, args...))
}

func InEpsilonf(t *testing.T, expected interface{}, actual interface{}, epsilon float64, msg string, args ...interface{}) {
	testify_stats.AddAssert(t, assert.InEpsilonf(t, expected, actual, epsilon, msg, args...))
}

func IsTypef(t *testing.T, expectedType interface{}, object interface{}, msg string, args ...interface{}) {
	testify_stats.AddAssert(t, assert.IsTypef(t, expectedType, object, msg, args...))
}

func JSONEqf(t *testing.T, expected string, actual string, msg string, args ...interface{}) {
	testify_stats.AddAssert(t, assert.JSONEqf(t, expected, actual, msg, args...))
}

func Lenf(t *testing.T, object interface{}, length int, msg string, args ...interface{}) {
	testify_stats.AddAssert(t, assert.Lenf(t, object, length, msg, args...))
}

func Nilf(t *testing.T, object interface{}, msg string, args ...interface{}) {
	testify_stats.AddAssert(t, assert.Nilf(t, object, msg, args...))
}

func NoErrorf(t *testing.T, err error, msg string, args ...interface{}) {
	testify_stats.AddAssert(t, assert.NoErrorf(t, err, msg, args...))
}

func NotContainsf(t *testing.T, s interface{}, contains interface{}, msg string, args ...interface{}) {
	testify_stats.AddAssert(t, assert.NotContainsf(t, s, contains, msg, args...))
}

func NotEmptyf(t *testing.T, object interface{}, msg string, args ...interface{}) {
	testify_stats.AddAssert(t, assert.NotEmptyf(t, object, msg, args...))
}

func NotEqualf(t *testing.T, expected interface{}, actual interface{}, msg string, args ...interface{}) {
	testify_stats.AddAssert(t, assert.NotEqualf(t, expected, actual, msg, args...))
}

func NotNilf(t *testing.T, object interface{}, msg string, args ...interface{}) {

	testify_stats.AddAssert(t, assert.NotNilf(t, object, msg, args...))
}

func NotPanicsf(t *testing.T, f assert.PanicTestFunc, msg string, args ...interface{}) {
	testify_stats.AddAssert(t, assert.NotPanicsf(t, f, msg, args...))
}

func NotRegexpf(t *testing.T, rx interface{}, str interface{}, msg string, args ...interface{}) {
	testify_stats.AddAssert(t, assert.NotRegexpf(t, rx, str, msg, args...))
}

func NotSubsetf(t *testing.T, list interface{}, subset interface{}, msg string, args ...interface{}) {
	testify_stats.AddAssert(t, assert.NotSubsetf(t, list, subset, msg, args...))
}

func NotZerof(t *testing.T, i interface{}, msg string, args ...interface{}) {
	testify_stats.AddAssert(t, assert.NotZerof(t, i, msg, args...))
}

func PanicsWithValuef(t *testing.T, expected interface{}, f assert.PanicTestFunc, msg string, args ...interface{}) {
	testify_stats.AddAssert(t, assert.PanicsWithValuef(t, expected, f, msg, args...))
}

func Panicsf(t *testing.T, f assert.PanicTestFunc, msg string, args ...interface{}) {
	testify_stats.AddAssert(t, assert.Panicsf(t, f, msg, args...))
}

func Regexpf(t *testing.T, rx interface{}, str interface{}, msg string, args ...interface{}) {
	testify_stats.AddAssert(t, assert.Regexpf(t, rx, str, msg, args...))
}

func Subsetf(t *testing.T, list interface{}, subset interface{}, msg string, args ...interface{}) {
	testify_stats.AddAssert(t, assert.Subsetf(t, list, subset, msg, args...))
}

func Truef(t *testing.T, value bool, msg string, args ...interface{}) {
	testify_stats.AddAssert(t, assert.Truef(t, value, msg, args...))
}

func WithinDurationf(t *testing.T, expected time.Time, actual time.Time, delta time.Duration, msg string, args ...interface{}) {
	testify_stats.AddAssert(t, assert.WithinDurationf(t, expected, actual, delta, msg, args...))
}

func Zerof(t *testing.T, i interface{}, msg string, args ...interface{}) {
	testify_stats.AddAssert(t, assert.Zerof(t, i, msg, args...))
}
