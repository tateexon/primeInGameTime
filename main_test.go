package main_test

import (
	"fmt"
	"testing"
	"io/ioutil"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tateexon/primeInGameTime/lib"
)

func TestFindOffsetFile1(t *testing.T) {
	filename := "test.gci"
	expectedTime := fmt.Sprintf("%f",float64(124.90000651404262))

	data, err := ioutil.ReadFile(filename)
	require.NoError(t, err)

	t1 := lib.FindFromOffsets(data, 44617)
	found := fmt.Sprintf("%f", t1)
	assert.Equal(t, expectedTime, found, "Did not find expected time")
}

func TestFindOffsetFile2(t *testing.T) {
	filename := "test.gci"
	expectedTime := fmt.Sprintf("%f",float64(2409.366792))

	data, err := ioutil.ReadFile(filename)
	require.NoError(t, err)

	t1 := lib.FindFromOffsets(data, 52137)
	found := fmt.Sprintf("%f", t1)
	assert.Equal(t, expectedTime, found, "Did not find expected time")
}

func TestFindOffsetFile3(t *testing.T) {
	filename := "test.gci"
	expectedTime := fmt.Sprintf("%f",float64(696.650036))

	data, err := ioutil.ReadFile(filename)
	require.NoError(t, err)

	t1 := lib.FindFromOffsets(data, 59657)
	found := fmt.Sprintf("%f", t1)
	assert.Equal(t, expectedTime, found, "Did not find expected time")
}
