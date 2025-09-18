package ds

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestZSet_Skiplist_GetRank(t *testing.T) {
	ss := CreateZSet()
	ss.Add(20.0, "k2")
	ss.Add(40.0, "k4")
	ss.Add(10.0, "k1")
	ss.Add(30.0, "k3")
	ss.Add(50.0, "k5")
	ss.Add(60.0, "k6")
	ss.Add(80.0, "k8")
	ss.Add(70.0, "k7")

	rank, score := ss.GetRank("k1", false)
	assert.EqualValues(t, 0, rank)
	assert.EqualValues(t, 10.0, score)

	rank, score = ss.GetRank("k2", false)
	assert.EqualValues(t, 1, rank)
	assert.EqualValues(t, 20.0, score)

	rank, score = ss.GetRank("k3", false)
	assert.EqualValues(t, 2, rank)
	assert.EqualValues(t, 30.0, score)

	rank, score = ss.GetRank("k4", false)
	assert.EqualValues(t, 3, rank)
	assert.EqualValues(t, 40.0, score)

	rank, score = ss.GetRank("k5", false)
	assert.EqualValues(t, 4, rank)
	assert.EqualValues(t, 50.0, score)

	rank, score = ss.GetRank("k6", false)
	assert.EqualValues(t, 5, rank)
	assert.EqualValues(t, 60.0, score)

	rank, score = ss.GetRank("k7", false)
	assert.EqualValues(t, 6, rank)
	assert.EqualValues(t, 70.0, score)

	rank, score = ss.GetRank("k8", false)
	assert.EqualValues(t, 7, rank)
	assert.EqualValues(t, 80.0, score)
}
