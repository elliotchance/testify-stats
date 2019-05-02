# 🔢 github.com/elliotchance/testify-stats

`testify-stats` is a drop in replacement for
[testify](https://github.com/stretchr/testify) that will print stats at the end
of the test suite:

```txt
$ go test
PASS
Tests: 460, Assertions: 1319, Time: 18.600047ms
ok      github.com/elliotchance/testify-stats   0.014s
```

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
