package prototype

// Cloneable interface or should it be Cloner or something like that? Go idioms help!
type Cloneable[T any] interface {
	Clone() T
}

// I don't know how data people do things, so the struct might be a bit off.
//But it is just an example to understand Prototype Design pattern...

// Data for our example
type Data struct {
	Map map[string]any
}

func NewData(m map[string]any) *Data {
	return &Data{
		Map: m,
	}
}

// DataFrame takes pointer to the Data struct, to make things a bit harder by nesting stuff...
type DataFrame struct {
	Name string
	Data *Data
}

// NewDataFrame constructs new DataFrame
func NewDataFrame(name string, data *Data) *DataFrame {
	return &DataFrame{
		Name: name,
		Data: data,
	}
}

// Clone method for DataFrame, to implement Cloneable interface
func (d DataFrame) Clone() DataFrame {
	// We need to clone map, not just copy the reference of it!
	deepCloneMap := make(map[string]any)
	for k, v := range d.Data.Map {
		deepCloneMap[k] = v
	}

	// Primitives are just copied, no need to clone them
	return DataFrame{
		Name: d.Name,
		Data: &Data{
			Map: deepCloneMap,
		},
	}
}
