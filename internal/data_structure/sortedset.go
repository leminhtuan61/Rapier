package data_structure

type SortedSet struct {
	Tree         *BPlusTree
	MemberScores map[string]float64
}

func NewSortedSet(degree int) *SortedSet {
	return &SortedSet{
		Tree:         NewBPlusTree(degree),
		MemberScores: make(map[string]float64),
	}
}

func (ss *SortedSet) Add(score float64, member string) int {
	return ss.Tree.Add(score, member)
}

func (ss *SortedSet) GetScore(member string) (float64, bool) {
	return ss.Tree.Score(member)
}

func (ss *SortedSet) GetRank(member string) int {
	return ss.Tree.GetRank(member)
}
