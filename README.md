# ketchup
random golang utilities


```go
package simpay // import "."

var InvalidNumber = errors.New("invalid number")
func Check(auth Auth, service_id, number, code string) (price int, err error)
func IsCodeNotFoundErr(err error) bool
func IsCodeUsedErr(err error) bool
func NumberToPrice(number string) (price int, err error)
type Auth struct{ ... }
type Error struct{ ... }
```

```golang
package dbutils // import "."

//CheckAffected checks if result.RowsAffected() are equal to the expected numeber
func CheckAffected(result sql.Result, sqlError error, expected int64) error

//IsMySQLDuplicate checks if mysql error is ER_DUP_ENTRY mysql error
func IsMySQLDuplicate(err error) bool
```

```go
package jsonutils // import "."

type Time struct{ ... }
    func New(t time.Time) *Time
```