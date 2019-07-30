
## validator/new

> Verify fields in a structure or structure slice

## define struct use `validate` tag

```golang
type Fruit struct {
	Name string `json:"name" validate:"lte=5"`
}
```

## more

> https://godoc.org/gopkg.in/go-playground/validator.v9
