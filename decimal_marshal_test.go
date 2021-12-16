package decimal

import (
	"gopkg.in/yaml.v3"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecimal_MarshalJSON(t *testing.T) {
	x, ok := NewDecimalFromString("1")
	assert.True(t, ok)
	xJson, err := json.Marshal(x)
	assert.NoError(t, err)
	assert.Equal(t, []byte(`"1"`), xJson)
	assert.NoError(t, json.Unmarshal(xJson, &x))
	assert.Equal(t, x, NewDecimalFromStringOrDefault("1", NewDecimalZero()))

	x, ok = NewDecimalFromString("1.5555555")
	assert.True(t, ok)
	xJson, err = json.Marshal(x)
	assert.NoError(t, err)
	assert.Equal(t, []byte(`"1.5555555"`), xJson)
	assert.NoError(t, json.Unmarshal(xJson, &x))
	assert.Equal(t, x, NewDecimalFromStringOrDefault("1.5555555", NewDecimalZero()))
}

func TestDecimal_MarshalYAML(t *testing.T) {
	x, ok := NewDecimalFromString("1")
	assert.True(t, ok)
	xYaml, err := yaml.Marshal(x)
	assert.NoError(t, err)
	actualXYaml := append([]byte(`"1"`), 0xa)
	assert.Equal(t, actualXYaml, xYaml)
	assert.NoError(t, yaml.Unmarshal(actualXYaml, &x))
	assert.Equal(t, x, NewDecimalFromStringOrDefault("1", NewDecimalZero()))

	x, ok = NewDecimalFromString("1.5555555")
	assert.True(t, ok)
	xYaml, err = yaml.Marshal(x)
	assert.NoError(t, err)
	actualXYaml = append([]byte(`"1.5555555"`), 0xa)
	assert.Equal(t, actualXYaml, xYaml)
	assert.NoError(t, yaml.Unmarshal(actualXYaml, &x))
	assert.Equal(t, x, NewDecimalFromStringOrDefault("1.5555555", NewDecimalZero()))
}
