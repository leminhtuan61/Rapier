package data_structure

type KeySpaceStat struct {
	Key    int64
	Expire int64
}

var HashKeySpaceStat = KeySpaceStat{Key: 0, Expire: 0}
