package require

import (
	testify_stats "github.com/elliotchance/testify-stats"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/url"
	"testing"
	"time"
)

func Condition(t *testing.T, comp assert.Comparison, msgAndArgs ...interface{}) {
	require.Condition(t, comp, msgAndArgs...)
	testify_stats.AddAssert(t, true)
}

func Conditionf(t *testing.T, comp assert.Comparison, msg string, args ...interface{}) {
	require.Conditionf(t, comp, msg, args)
	testify_stats.AddAssert(t, true)
}

func Contains(t *testing.T, s interface{}, contains interface{}, msgAndArgs ...interface{}) {
	require.Contains(t, s, contains, msgAndArgs...)
	testify_stats.AddAssert(t, true)
}

func Containsf(t *testing.T, s interface{}, contains interface{}, msg string, args ...interface{}) {
	require.Containsf(t, s, contains, msg, args...)
	testify_stats.AddAssert(t, true)
}

func DirExists(t *testing.T, path string, msgAndArgs ...interface{}) {
	require.DirExists(t, path, msgAndArgs...)
	testify_stats.AddAssert(t, true)
}

func DirExistsf(t *testing.T, path string, msg string, args ...interface{}) {
	require.DirExistsf(t, path, msg, args...)
	testify_stats.AddAssert(t, true)
}

func ElementsMatch(t *testing.T, listA interface{}, listB interface{}, msgAndArgs ...interface{}) {
	require.ElementsMatch(t, listA, listB, msgAndArgs...)
	testify_stats.AddAssert(t, true)
}

func ElementsMatchf(t *testing.T, listA interface{}, listB interface{}, msg string, args ...interface{}) {
	require.ElementsMatchf(t, listA, listB, msg, args...)
	testify_stats.AddAssert(t, true)
}

func Empty(t *testing.T, object interface{}, msgAndArgs ...interface{}) {
	require.Empty(t, object, msgAndArgs...)
	testify_stats.AddAssert(t, true)
}

func Emptyf(t *testing.T, object interface{}, msg string, args ...interface{}) {
	require.Emptyf(t, object, msg, args...)
	testify_stats.AddAssert(t, true)
}

func Equal(t *testing.T, expected interface{}, actual interface{}, msgAndArgs ...interface{}) {
	require.Equal(t, expected, actual, msgAndArgs...)
	testify_stats.AddAssert(t, true)
}

func EqualError(t *testing.T, theError error, errString string, msgAndArgs ...interface{}) {
	require.EqualError(t, theError, errString, msgAndArgs...)
	testify_stats.AddAssert(t, true)
}

func EqualErrorf(t *testing.T, theError error, errString string, msg string, args ...interface{}) {
	require.EqualErrorf(t, theError, errString, msg, args...)
	testify_stats.AddAssert(t, true)
}

func EqualValues(t *testing.T, expected interface{}, actual interface{}, msgAndArgs ...interface{}) {
	require.EqualValues(t, expected, actual, msgAndArgs...)
	testify_stats.AddAssert(t, true)
}

func EqualValuesf(t *testing.T, expected interface{}, actual interface{}, msg string, args ...interface{}) {
	require.EqualValuesf(t, expected, actual, msg, args...)
	testify_stats.AddAssert(t, true)
}

func Equalf(t *testing.T, expected interface{}, actual interface{}, msg string, args ...interface{}) {
	require.Equalf(t, expected, actual, msg, args...)
	testify_stats.AddAssert(t, true)
}

func Error(t *testing.T, err error, msgAndArgs ...interface{}) {
	require.Error(t, err, msgAndArgs...)
	testify_stats.AddAssert(t, true)
}

func Errorf(t *testing.T, err error, msg string, args ...interface{}) {
	require.Errorf(t, err, msg, args...)
	testify_stats.AddAssert(t, true)
}

func Exactly(t *testing.T, expected interface{}, actual interface{}, msgAndArgs ...interface{}) {
	require.Exactly(t, expected, actual, msgAndArgs...)
	testify_stats.AddAssert(t, true)
}

func Exactlyf(t *testing.T, expected interface{}, actual interface{}, msg string, args ...interface{}) {
	require.Exactlyf(t, expected, actual, msg, args...)
	testify_stats.AddAssert(t, true)
}

func Fail(t *testing.T, failureMessage string, msgAndArgs ...interface{}) {
	require.Fail(t, failureMessage, msgAndArgs...)
	testify_stats.AddAssert(t, true)
}

func FailNow(t *testing.T, failureMessage string, msgAndArgs ...interface{}) {
	require.FailNow(t, failureMessage, msgAndArgs...)
	testify_stats.AddAssert(t, true)
}

func FailNowf(t *testing.T, failureMessage string, msg string, args ...interface{}) {
	require.FailNowf(t, failureMessage, msg, args...)
	testify_stats.AddAssert(t, true)
}

func Failf(t *testing.T, failureMessage string, msg string, args ...interface{}) {
	require.Failf(t, failureMessage, msg, args...)
	testify_stats.AddAssert(t, true)
}

func False(t *testing.T, value bool, msgAndArgs ...interface{}) {
	require.False(t, value, msgAndArgs...)
	testify_stats.AddAssert(t, true)
}

func Falsef(t *testing.T, value bool, msg string, args ...interface{}) {
	require.Falsef(t, value, msg, args...)
	testify_stats.AddAssert(t, true)
}

func FileExists(t *testing.T, path string, msgAndArgs ...interface{}) {
	require.FileExists(t, path, msgAndArgs...)
	testify_stats.AddAssert(t, true)
}

func FileExistsf(t *testing.T, path string, msg string, args ...interface{}) {
	require.FileExistsf(t, path, msg, args...)
	testify_stats.AddAssert(t, true)
}

func HTTPBodyContains(t *testing.T, handler http.HandlerFunc, method string, url string, values url.Values, str interface{}, msgAndArgs ...interface{}) {
	require.HTTPBodyContains(t, handler, method, url, values, str, msgAndArgs...)
	testify_stats.AddAssert(t, true)
}

func HTTPBodyContainsf(t *testing.T, handler http.HandlerFunc, method string, url string, values url.Values, str interface{}, msg string, args ...interface{}) {
	require.HTTPBodyContainsf(t, handler, method, url, values, str, msg, args...)
	testify_stats.AddAssert(t, true)
}

func HTTPBodyNotContains(t *testing.T, handler http.HandlerFunc, method string, url string, values url.Values, str interface{}, msgAndArgs ...interface{}) {
	require.HTTPBodyNotContains(t, handler, method, url, values, str, msgAndArgs...)
	testify_stats.AddAssert(t, true)
}

func HTTPBodyNotContainsf(t *testing.T, handler http.HandlerFunc, method string, url string, values url.Values, str interface{}, msg string, args ...interface{}) {
	require.HTTPBodyNotContainsf(t, handler, method, url, values, str, msg, args...)
	testify_stats.AddAssert(t, true)
}

func HTTPError(t *testing.T, handler http.HandlerFunc, method string, url string, values url.Values, msgAndArgs ...interface{}) {
	require.HTTPError(t, handler, method, url, values, msgAndArgs...)
	testify_stats.AddAssert(t, true)
}

func HTTPErrorf(t *testing.T, handler http.HandlerFunc, method string, url string, values url.Values, msg string, args ...interface{}) {
	require.HTTPErrorf(t, handler, method, url, values, msg, args...)
	testify_stats.AddAssert(t, true)
}

func HTTPRedirect(t *testing.T, handler http.HandlerFunc, method string, url string, values url.Values, msgAndArgs ...interface{}) {
	require.HTTPRedirect(t, handler, method, url, values, msgAndArgs...)
	testify_stats.AddAssert(t, true)
}

func HTTPRedirectf(t *testing.T, handler http.HandlerFunc, method string, url string, values url.Values, msg string, args ...interface{}) {
	require.HTTPRedirectf(t, handler, method, url, values, msg, args...)
	testify_stats.AddAssert(t, true)
}

func HTTPSuccess(t *testing.T, handler http.HandlerFunc, method string, url string, values url.Values, msgAndArgs ...interface{}) {
	require.HTTPSuccess(t, handler, method, url, values, msgAndArgs...)
	testify_stats.AddAssert(t, true)
}

func HTTPSuccessf(t *testing.T, handler http.HandlerFunc, method string, url string, values url.Values, msg string, args ...interface{}) {
	require.HTTPSuccessf(t, handler, method, url, values, msg, args...)
	testify_stats.AddAssert(t, true)
}

func Implements(t *testing.T, interfaceObject interface{}, object interface{}, msgAndArgs ...interface{}) {
	require.Implements(t, interfaceObject, object, msgAndArgs...)
	testify_stats.AddAssert(t, true)
}

func Implementsf(t *testing.T, interfaceObject interface{}, object interface{}, msg string, args ...interface{}) {
	require.Implementsf(t, interfaceObject, object, msg, args...)
	testify_stats.AddAssert(t, true)
}

func InDelta(t *testing.T, expected interface{}, actual interface{}, delta float64, msgAndArgs ...interface{}) {
	require.InDelta(t, expected, actual, delta, msgAndArgs...)
	testify_stats.AddAssert(t, true)
}

func InDeltaMapValues(t *testing.T, expected interface{}, actual interface{}, delta float64, msgAndArgs ...interface{}) {
	require.InDeltaMapValues(t, expected, actual, delta, msgAndArgs...)
	testify_stats.AddAssert(t, true)
}

func InDeltaMapValuesf(t *testing.T, expected interface{}, actual interface{}, delta float64, msg string, args ...interface{}) {
	require.InDeltaMapValuesf(t, expected, actual, delta, msg, args...)
	testify_stats.AddAssert(t, true)
}

func InDeltaSlice(t *testing.T, expected interface{}, actual interface{}, delta float64, msgAndArgs ...interface{}) {
	require.InDeltaSlice(t, expected, actual, delta, msgAndArgs...)
	testify_stats.AddAssert(t, true)
}

func InDeltaSlicef(t *testing.T, expected interface{}, actual interface{}, delta float64, msg string, args ...interface{}) {
	require.InDeltaSlicef(t, expected, actual, delta, msg, args...)
	testify_stats.AddAssert(t, true)
}

func InDeltaf(t *testing.T, expected interface{}, actual interface{}, delta float64, msg string, args ...interface{}) {
	require.InDeltaf(t, expected, actual, delta, msg, args...)
	testify_stats.AddAssert(t, true)
}

func InEpsilon(t *testing.T, expected interface{}, actual interface{}, epsilon float64, msgAndArgs ...interface{}) {
	require.InEpsilon(t, expected, actual, epsilon, msgAndArgs...)
	testify_stats.AddAssert(t, true)
}

func InEpsilonSlice(t *testing.T, expected interface{}, actual interface{}, epsilon float64, msgAndArgs ...interface{}) {
	require.InEpsilonSlice(t, expected, actual, epsilon, msgAndArgs...)
	testify_stats.AddAssert(t, true)
}

func InEpsilonSlicef(t *testing.T, expected interface{}, actual interface{}, epsilon float64, msg string, args ...interface{}) {
	require.InEpsilonSlicef(t, expected, actual, epsilon, msg, args...)
	testify_stats.AddAssert(t, true)
}

func InEpsilonf(t *testing.T, expected interface{}, actual interface{}, epsilon float64, msg string, args ...interface{}) {
	require.InEpsilonf(t, expected, actual, epsilon, msg, args...)
	testify_stats.AddAssert(t, true)
}

func IsType(t *testing.T, expectedType interface{}, object interface{}, msgAndArgs ...interface{}) {
	require.IsType(t, expectedType, object, msgAndArgs...)
	testify_stats.AddAssert(t, true)
}

func IsTypef(t *testing.T, expectedType interface{}, object interface{}, msg string, args ...interface{}) {
	require.IsTypef(t, expectedType, object, msg, args...)
	testify_stats.AddAssert(t, true)
}

func JSONEq(t *testing.T, expected string, actual string, msgAndArgs ...interface{}) {
	require.JSONEq(t, expected, actual, msgAndArgs...)
	testify_stats.AddAssert(t, true)
}

func JSONEqf(t *testing.T, expected string, actual string, msg string, args ...interface{}) {
	require.JSONEqf(t, expected, actual, msg, args...)
	testify_stats.AddAssert(t, true)
}

func Len(t *testing.T, object interface{}, length int, msgAndArgs ...interface{}) {
	require.Len(t, object, length, msgAndArgs...)
	testify_stats.AddAssert(t, true)
}

func Lenf(t *testing.T, object interface{}, length int, msg string, args ...interface{}) {
	require.Lenf(t, object, length, msg, args...)
	testify_stats.AddAssert(t, true)
}

func Nil(t *testing.T, object interface{}, msgAndArgs ...interface{}) {
	require.Nil(t, object, msgAndArgs...)
	testify_stats.AddAssert(t, true)
}

func Nilf(t *testing.T, object interface{}, msg string, args ...interface{}) {
	require.Nilf(t, object, msg, args...)
	testify_stats.AddAssert(t, true)
}

func NoError(t *testing.T, err error, msgAndArgs ...interface{}) {
	require.NoError(t, err, msgAndArgs...)
	testify_stats.AddAssert(t, true)
}

func NoErrorf(t *testing.T, err error, msg string, args ...interface{}) {
	require.NoErrorf(t, err, msg, args...)
	testify_stats.AddAssert(t, true)
}

func NotContains(t *testing.T, s interface{}, contains interface{}, msgAndArgs ...interface{}) {
	require.NotContains(t, s, contains, msgAndArgs...)
	testify_stats.AddAssert(t, true)
}

func NotContainsf(t *testing.T, s interface{}, contains interface{}, msg string, args ...interface{}) {
	require.NotContainsf(t, s, contains, msg, args...)
	testify_stats.AddAssert(t, true)
}

func NotEmpty(t *testing.T, object interface{}, msgAndArgs ...interface{}) {
	require.NotEmpty(t, object, msgAndArgs...)
	testify_stats.AddAssert(t, true)
}

func NotEmptyf(t *testing.T, object interface{}, msg string, args ...interface{}) {
	require.NotEmptyf(t, object, msg, args...)
	testify_stats.AddAssert(t, true)
}

func NotEqual(t *testing.T, expected interface{}, actual interface{}, msgAndArgs ...interface{}) {
	require.NotEqual(t, expected, actual, msgAndArgs...)
	testify_stats.AddAssert(t, true)
}

func NotEqualf(t *testing.T, expected interface{}, actual interface{}, msg string, args ...interface{}) {
	require.NotEqualf(t, expected, actual, msg, args...)
	testify_stats.AddAssert(t, true)
}

func NotNil(t *testing.T, object interface{}, msgAndArgs ...interface{}) {
	require.NotNil(t, object, msgAndArgs...)
	testify_stats.AddAssert(t, true)
}

func NotNilf(t *testing.T, object interface{}, msg string, args ...interface{}) {
	require.NotNilf(t, object, msg, args...)
	testify_stats.AddAssert(t, true)
}

func NotPanics(t *testing.T, f assert.PanicTestFunc, msgAndArgs ...interface{}) {
	require.NotPanics(t, f, msgAndArgs...)
	testify_stats.AddAssert(t, true)
}

func NotPanicsf(t *testing.T, f assert.PanicTestFunc, msg string, args ...interface{}) {
	require.NotPanicsf(t, f, msg, args...)
	testify_stats.AddAssert(t, true)
}

func NotRegexp(t *testing.T, rx interface{}, str interface{}, msgAndArgs ...interface{}) {
	require.NotRegexp(t, rx, str, msgAndArgs...)
	testify_stats.AddAssert(t, true)
}

func NotRegexpf(t *testing.T, rx interface{}, str interface{}, msg string, args ...interface{}) {
	require.NotRegexpf(t, rx, str, msg, args...)
	testify_stats.AddAssert(t, true)
}

func NotSubset(t *testing.T, list interface{}, subset interface{}, msgAndArgs ...interface{}) {
	require.NotSubset(t, list, subset, msgAndArgs...)
	testify_stats.AddAssert(t, true)
}

func NotSubsetf(t *testing.T, list interface{}, subset interface{}, msg string, args ...interface{}) {
	require.NotSubsetf(t, list, subset, msg, args...)
	testify_stats.AddAssert(t, true)
}

func NotZero(t *testing.T, i interface{}, msgAndArgs ...interface{}) {
	require.NotZero(t, i, msgAndArgs...)
	testify_stats.AddAssert(t, true)
}

func NotZerof(t *testing.T, i interface{}, msg string, args ...interface{}) {
	require.NotZerof(t, i, msg, args...)
	testify_stats.AddAssert(t, true)
}

func Panics(t *testing.T, f assert.PanicTestFunc, msgAndArgs ...interface{}) {
	require.Panics(t, f, msgAndArgs...)
	testify_stats.AddAssert(t, true)
}

func PanicsWithValue(t *testing.T, expected interface{}, f assert.PanicTestFunc, msgAndArgs ...interface{}) {
	require.PanicsWithValue(t, expected, f, msgAndArgs...)
	testify_stats.AddAssert(t, true)
}

func PanicsWithValuef(t *testing.T, expected interface{}, f assert.PanicTestFunc, msg string, args ...interface{}) {
	require.PanicsWithValuef(t, expected, f, msg, args...)
	testify_stats.AddAssert(t, true)
}

func Panicsf(t *testing.T, f assert.PanicTestFunc, msg string, args ...interface{}) {
	require.Panicsf(t, f, msg, args...)
	testify_stats.AddAssert(t, true)
}

func Regexp(t *testing.T, rx interface{}, str interface{}, msgAndArgs ...interface{}) {
	require.Regexp(t, rx, str, msgAndArgs...)
	testify_stats.AddAssert(t, true)
}

func Regexpf(t *testing.T, rx interface{}, str interface{}, msg string, args ...interface{}) {
	require.Regexpf(t, rx, str, msg, args...)
	testify_stats.AddAssert(t, true)
}

func Subset(t *testing.T, list interface{}, subset interface{}, msgAndArgs ...interface{}) {
	require.Subset(t, list, subset, msgAndArgs...)
	testify_stats.AddAssert(t, true)
}

func Subsetf(t *testing.T, list interface{}, subset interface{}, msg string, args ...interface{}) {
	require.Subsetf(t, list, subset, msg, args...)
	testify_stats.AddAssert(t, true)
}

func True(t *testing.T, value bool, msgAndArgs ...interface{}) {
	require.True(t, value, msgAndArgs...)
	testify_stats.AddAssert(t, true)
}

func Truef(t *testing.T, value bool, msg string, args ...interface{}) {
	require.Truef(t, value, msg, args...)
	testify_stats.AddAssert(t, true)
}

func WithinDuration(t *testing.T, expected time.Time, actual time.Time, delta time.Duration, msgAndArgs ...interface{}) {
	require.WithinDuration(t, expected, actual, delta, msgAndArgs...)
	testify_stats.AddAssert(t, true)
}

func WithinDurationf(t *testing.T, expected time.Time, actual time.Time, delta time.Duration, msg string, args ...interface{}) {
	require.WithinDurationf(t, expected, actual, delta, msg, args...)
	testify_stats.AddAssert(t, true)
}

func Zero(t *testing.T, i interface{}, msgAndArgs ...interface{}) {
	require.Zero(t, i, msgAndArgs...)
	testify_stats.AddAssert(t, true)
}

func Zerof(t *testing.T, i interface{}, msg string, args ...interface{}) {
	require.Zerof(t, i, msg, args...)
	testify_stats.AddAssert(t, true)
}
