# 🔢 github.com/elliotchance/testify-stats

`testify-stats` is a drop in replacement for
[testify](https://github.com/stretchr/testify) that will print stats at the end
of the test suite:

```txt
$ go test -v
PASS
Tests: 460, Assertions: 1319, Time: 18.600047ms
ok      github.com/elliotchance/testify-stats   0.014s
```

**Note:** The statistics will only print with the `-v` option. This is because the testing package will capture all stdout and only print it under verbose more, or if there was a failure. I couldn't find a way around this. If you know of one, please let me know.

# How to Use It

1. Swap all imports of `github.com/stretchr/testify/assert` with
`github.com/elliotchance/testify-stats/assert`.

2. Either update or create a `TestMain`:

```go
import "github.com/elliotchance/testify-stats"

func TestMain(m *testing.M) {
	os.Exit(testify_stats.Run(m))
}
```

3. Run `go test` as normal.

# Build/CI Check

To make sure all assertions are recorded you can add the following step to your
build/CI:

```bash
$(exit $(grep -r '"github.com/stretchr/testify"' . --include \*.go | wc -l))
```
