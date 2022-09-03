package singly_test

import (
	"os"
	"testing"

	"github.com/mowazzem/ds/lists/singly"
	"github.com/stretchr/testify/assert"
)

var list *singly.List[int]

func TestMain(m *testing.M) {
	list = singly.NewList[int]()
	code := m.Run()
	os.Exit(code)
}

func TestPrepend(t *testing.T) {
	list.Append(0)
	assert.NotNil(t, list)
	assert.Equal(t, list.Len(), 1)
}

func TestAppend(t *testing.T) {
	list.Append(1)
	list.Append(2)
	list.Append(3)
	assert.NotNil(t, list)
	assert.Equal(t, list.Len(), 4)
}

func TestGetAll(t *testing.T) {
	vals := list.GetAll()
	assert.Equal(t, []int{0, 1, 2, 3}, vals)
}

func TestGetIndexOf(t *testing.T) {
	index := list.IndexOf(2)
	assert.Equal(t, 2, index)
	index = list.IndexOf(3)
	assert.Equal(t, 3, index)
	index = list.IndexOf(0)
	assert.Equal(t, 0, index)
}
