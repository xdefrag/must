package must

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAll(t *testing.T) {
	assert := assert.New(t)

	err := errors.New("")

	assert.Panics(func() { Must(err) })
	assert.NotPanics(func() { Must(nil) })

	assert.Panics(func() { Must2("Hello", err) })
	assert.NotPanics(func() { Must2("Hello", nil) })

	assert.Panics(func() { Must3("Hello", "World", err) })
	assert.NotPanics(func() { Must3("Hello", "World", nil) })

	assert.Panics(func() { Must2R("Hello", err) })
	assert.NotPanics(func() { Must2R("Hello", nil) })
	m2rval := Must2R("Hello", nil)
	assert.Equal("Hello", m2rval)

	assert.Panics(func() { Must3R("Hello", "World", err) })
	assert.NotPanics(func() { Must3R("Hello", "World", nil) })
	m3rval1, m3rval2 := Must3R("Hello", "World", nil)
	assert.Equal("Hello", m3rval1)
	assert.Equal("World", m3rval2)

	m := new(testingmock)

	TMust(m)(nil)
	m.AssertExpectations(t)

	m.On("Error", err).Return()
	TMust(m)(err)
	m.AssertExpectations(t)

	TMust2(m)("Hello", nil)
	m.AssertExpectations(t)

	m.On("Error", err).Return()
	TMust2(m)("Hello", err)
	m.AssertExpectations(t)

	TMust3(m)("Hello", "World", nil)
	m.AssertExpectations(t)

	m.On("Error", err).Return()
	TMust3(m)("Hello", "World", err)
	m.AssertExpectations(t)
}

type testingmock struct {
	mock.Mock
}

func (t *testingmock) Error(args ...interface{}) {
	t.Called(args...)
}
