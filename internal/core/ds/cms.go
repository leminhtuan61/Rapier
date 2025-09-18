package ds

import (
	"github.com/spaolacci/murmur3"
	"math"
)

// Log10PointFive is a precomputed value for log10(0.5).
const Log10PointFive = -0.30102999566

// CMS is the Count-Min Sketch data structure.
// The counter field has been changed to a 2D slice for better clarity and indexing.
type CMS struct {
	width uint32
	depth uint32
	// counter is now a 2D slice of uint32. The outer slice represents the rows (depth),
	// and the inner slice represents the columns (width).
	counter [][]uint32
}

// CreateCMS initializes a new Count-Min Sketch with a given width and depth.
func CreateCMS(w uint32, d uint32) *CMS {
	cms := &CMS{
		width: w,
		depth: d,
	}
	// Initialize the 2D slice.
	// We create a slice of slices, where the outer slice has 'd' elements (for depth).
	cms.counter = make([][]uint32, d)
	// We then loop through each "row" and initialize a slice of 'w' elements for the width.
	for i := uint32(0); i < d; i++ {
		cms.counter[i] = make([]uint32, w)
	}
	return cms
}

// CalcCMSDim calculates the dimensions (width and depth) of the CMS
// based on the desired error rate and probability.
func CalcCMSDim(errRate float64, errProb float64) (uint32, uint32) {
	w := uint32(math.Ceil(2.0 / errRate))
	d := uint32(math.Ceil(math.Log10(errProb) / Log10PointFive))
	return w, d
}

// calcHash calculates a 32-bit hash for the given item and seed.
func (c *CMS) calcHash(item string, seed uint32) uint32 {
	hasher := murmur3.New32WithSeed(seed)
	hasher.Write([]byte(item))
	return hasher.Sum32()
}

// IncrBy increments the count for an item by a specific value.
// It returns the estimated count for the item after the increment.
func (c *CMS) IncrBy(item string, value uint32) uint32 {
	var minCount uint32 = math.MaxUint32

	// Loop through each row of the 2D array.
	for i := uint32(0); i < c.depth; i++ {
		// Calculate a new hash for each row using the row index as the seed.
		hash := c.calcHash(item, i)
		// Use the hash to get the column index within the row.
		j := hash % c.width

		// Safely add the value to prevent overflow.
		if math.MaxUint32-c.counter[i][j] < value {
			c.counter[i][j] = math.MaxUint32
		} else {
			c.counter[i][j] += value
		}

		// Keep track of the minimum count across all rows.
		if c.counter[i][j] < minCount {
			minCount = c.counter[i][j]
		}
	}
	return minCount
}

// Count returns the estimated count for an item.
// It retrieves the minimum count across all hash functions to provide the most accurate estimate.
func (c *CMS) Count(item string) uint32 {
	var minCount uint32 = math.MaxUint32

	// Loop through each row of the 2D array.
	for i := uint32(0); i < c.depth; i++ {
		// Calculate the hash for this row.
		hash := c.calcHash(item, i)
		// Determine the column index.
		j := hash % c.width

		// Find the minimum count across all rows.
		if c.counter[i][j] < minCount {
			minCount = c.counter[i][j]
		}
	}
	return minCount
}
