# must
[![GoDoc](https://godoc.org/github.com/xdefrag/must?status.svg)](https://godoc.org/github.com/xdefrag/must)
[![Build Status](https://travis-ci.org/xdefrag/must.svg?branch=master)](https://travis-ci.org/xdefrag/must)
[![codecov](https://codecov.io/gh/xdefrag/must/branch/master/graph/badge.svg)](https://codecov.io/gh/xdefrag/must)
[![Go Report Card](https://goreportcard.com/badge/github.com/xdefrag/must)](https://goreportcard.com/report/github.com/xdefrag/must)

Your lovely package has no Must functions but you really need them? FEAR NOT! Now you can just chain any functions to Must functions from this package!

```go
must.Must(json.Unmarshal(b, &v) // If something going wrong -- it's just panic!
must.Must2(db.Exec(`SELECT 1`)) // Yeah, jmoiron/sqlx has must function, BUT what
                                // if you want to stick with default?

r := must.Must2R(db.Exec(`INSERT INTO names (value) VALUES ('jim')`))
r.(sql.Result) // Go has no generics, this is true. But who needs them, if you
               // know what returning from your functions!

must.TMust2(t)(db.Exec(`SELECT 1`)) // Really useful with tests: "t" must have Error(args ...interface{})
                                    // func like in testing.TB interface, and whole thing call t.Error(err)
                                    // if err not nil.

// And so on.
```
