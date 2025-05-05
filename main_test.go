package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateRandomElements(t *testing.T) {
	negativeTestValues := []int{-1, 0}

	for _, val := range negativeTestValues {
		_, err := generateRandomElements(val)
		assert.Error(t, err)
	}

	positiveTestValues := []int{1, 10, 100}

	for _, val := range positiveTestValues {
		slice, err := generateRandomElements(val)
		assert.NoError(t, err)
		assert.Equal(t, val, len(slice))
	}
}

func TestMaximum(t *testing.T) {
	emptySlice := []int{}
	oneValuePositive := []int{1}
	oneValueNegative := []int{-1}
	negativeValues := []int{-1, 0, 1}
	positiveValues := []int{1, 2, 3, 4, 5, 6}

	_, err := maximum(emptySlice)
	assert.Error(t, err)

	max, err := maximum(oneValuePositive)
	assert.NoError(t, err)
	assert.Equal(t, 1, max)

	_, err = maximum(oneValueNegative)
	assert.Error(t, err)

	_, err = maximum(negativeValues)
	assert.Error(t, err)

	max, err = maximum(positiveValues)
	assert.NoError(t, err)
	assert.Equal(t, 6, max)
}

func TestMaxChunks(t *testing.T) {
	data := []int{1, 5, 10, 3, 4, 2, 9, 8}
	max, err := maxChunks(data)
	assert.NoError(t, err)
	assert.Equal(t, 10, max)

	_, err = maxChunks([]int{})
	assert.Error(t, err)
}

func TestEqualMax(t *testing.T) {
	data := []int{1, 5, 10, 3, 4, 2, 9, 8}
	maxOfMaximum, err := maximum(data)
	assert.NoError(t, err)

	maxOfMaxChunks, err := maxChunks(data)
	assert.NoError(t, err)

	assert.Equal(t, maxOfMaximum, maxOfMaxChunks)
}
