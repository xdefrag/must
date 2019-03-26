# must
Your lovely package has no Must functions but you really need them? FEAR NOT! Now you can just chain any functions to Must functions from this package!

```go
must.Must(json.Unmarshal(b, &v) // If something going wrong -- it's just panic!
must.Must2(db.Exec(`SELECT 1`)) // Yeah, jmoiron/sqlx has must function, BUT what if you don't have it?
r := must.Must2R(db.Exec(`INSERT INTO names (value) VALUES ('jim')`))
r.(db.Result) // Go has no generics, this is true. But who needs them, if you 
              // know what returning from your functions!

// And so on.
```
