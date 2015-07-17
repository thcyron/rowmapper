**You probably shouldn’t use rowmapper. Consider using [sqlbuilder](https://github.com/thcyron/sqlbuilder) instead.**

---

rowmapper
=========

`rowmapper` is a Go library for mapping database columns to variables.

Are you tired of calling the `Scan` function with countless of arguments?
Then `rowmapper` is for you. Just define the mapping from column to variable
beforehand and let `rowmapper` call `Scan` with the right arguments for you.

So instead of doing this:

```go
rows, _ := db.Query("SELECT id, name, age FROM users")
rows.Scan(&id, &name, &age)
```

You do this:

```go
mapper := rowmapper.New()
mapper.Map("id", &id)
mapper.Map("name", &name)
mapper.Map("age", &age)

rows, _ := db.Query("SELECT id, name, age FROM users")
mapper.Do(rows)
```

If this short example doesn’t show you the benefits, imagine querying a table
with 20 columns and a corresponding `Scan` call with 20 arguments where the
order is significant.

It also works well with `SELECT * FROM` queries where the column order isn’t obvious.

Use `DoOne` if you want to query at most one row. It works just like `QueryRow`
from database/sql.
