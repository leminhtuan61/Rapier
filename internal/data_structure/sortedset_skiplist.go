package data_structure

type ZSet struct {
	zskiplist *Skiplist
	// map from ele to score
	dict map[string]float64
}

func CreateZSet() *ZSet {
	zs := ZSet{
		zskiplist: CreateSkiplist(),
		dict:      map[string]float64{},
	}
	return &zs
}

func (zs *ZSet) Add(score float64, ele string) int {
	if len(ele) == 0 {
		return 0
	}
	if curScore, exist := zs.dict[ele]; exist {
		if curScore != score {
			znode := zs.zskiplist.UpdateScore(curScore, ele, score)
			zs.dict[ele] = znode.score
		}
		return 1
	}

	znode := zs.zskiplist.Insert(score, ele)
	zs.dict[ele] = znode.score
	return 1
}

/*
Returns the 0-based rank of the object or -1 if the object does not exist.
If reverse is false, rank is computed considering as first element the one
with the lowest score. If reverse is true, rank is computed considering as element with rank 0 the
one with the highest score.
*/
func (zs *ZSet) GetRank(ele string, reverse bool) (rank int64, score float64) {
	setSize := zs.zskiplist.length
	score, exist := zs.dict[ele]
	if !exist {
		return -1, 0
	}
	rank = int64(zs.zskiplist.GetRank(score, ele))
	if reverse {
		rank = int64(setSize) - rank
	} else {
		rank--
	}
	return rank, score
}

func (zs *ZSet) GetScore(ele string) (int, float64) {
	score, exist := zs.dict[ele]
	if !exist {
		return -1, 0
	}
	return 0, score
}

func (zs *ZSet) Len() int {
	return len(zs.dict)
}
