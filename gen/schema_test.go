package gen

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidDataType(t *testing.T) {
	assert.NoError(t, CheckSchemaType("string"))
	assert.NoError(t, CheckSchemaType("number"))
	assert.NoError(t, CheckSchemaType("integer"))
	assert.NoError(t, CheckSchemaType("boolean"))
	assert.NoError(t, CheckSchemaType("array"))
	assert.NoError(t, CheckSchemaType("object"))

	assert.Error(t, CheckSchemaType("oops"))
}

func TestTransToValidSchemeType(t *testing.T) {
	assert.Equal(t, TransToValidSchemeType("uint"), "integer")
	assert.Equal(t, TransToValidSchemeType("uint32"), "integer")
	assert.Equal(t, TransToValidSchemeType("uint64"), "integer")
	assert.Equal(t, TransToValidSchemeType("float32"), "number")
	assert.Equal(t, TransToValidSchemeType("bool"), "boolean")
	assert.Equal(t, TransToValidSchemeType("string"), "string")

	// should accept any type, due to user defined types
	TransToValidSchemeType("oops")
}

func TestIsGolangPrimitiveType(t *testing.T) {
	t.Run("fase", func(tt *testing.T) {
		assert.False(tt, IsGolangPrimitiveType("sdafdsa"), "sdafds 不是基本类型")
	})
	t.Run("true", func(tt *testing.T) {
		assert.True(tt, IsGolangPrimitiveType("int"), "int 是基本类型")
		assert.True(tt, IsGolangPrimitiveType("string"), "string 不是基本类型")
	})
}

func TestIsNumericType(t *testing.T) {
	assert.True(t, IsNumericType("integer"))
	assert.True(t, IsNumericType("number"))
	assert.False(t, IsNumericType("numbers"))
}
