package core

import "Rapier/internal/core/ds"

var dict *ds.Dictionary

func init() {
	dict = ds.NewDictionary()
}
