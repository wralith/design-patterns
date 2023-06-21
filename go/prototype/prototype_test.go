package prototype

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDataFrame_Clone(t *testing.T) {
	// It's neither data nor frame I know, but it is not the point here, right?
	prototype := NewDataFrame("Test DataFrame",
		NewData(map[string]any{
			"name":     "Test",
			"category": "Creational",
			"pattern":  "Prototype",
		}))

	clone := prototype.Clone()
	// assert.Equal runs reflect.DeepEqual that recursively check values
	assert.Equal(t, prototype.Data.Map, clone.Data.Map)

	clone.Data.Map = map[string]any{
		"title":   "Prototype, Creational Design Pattern",
		"version": 2,
	}

	assert.NotEqual(t, prototype, clone)
	assert.NotSame(t, prototype, clone)
	assert.NotSame(t, prototype.Data, clone.Data)
	assert.NotSame(t, prototype.Data.Map, clone.Data.Map)

	assert.Equal(t, "Prototype", prototype.Data.Map["pattern"])
	assert.Equal(t, "Prototype, Creational Design Pattern", clone.Data.Map["title"])

	assert.Nil(t, clone.Data.Map["name"])
	assert.Nil(t, prototype.Data.Map["version"])

	assert.Equal(t, 2, clone.Data.Map["version"])
}
