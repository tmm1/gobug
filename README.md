```
$ go test -v ./
=== RUN   TestTypeInfo
--- PASS: TestTypeInfo (0.00s)
PASS
ok      github.com/tmm1/gobug   0.090s
```

```
$ docker buildx build --platform linux/arm/v7 .
=== RUN   TestTypeInfo
    types_test.go:27:
              Error Trace:    types_test.go:27
              Error:          Not equal:
                              expected: []int{1, 0}
                              actual  : []int{1, 1}

                              Diff:
                              --- Expected
                              +++ Actual
                              @@ -2,3 +2,3 @@
                                (int) 1,
                              - (int) 0
                              + (int) 1
                               }
              Test:           TestTypeInfo
--- FAIL: TestTypeInfo (0.01s)
FAIL
FAIL  github.com/tmm1/gobug   0.105s
```
