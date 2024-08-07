package MaekSmol

import (
	"container/heap"
)

type ByteFreq struct {
    val byte
    freq uint
}
 
type PrioQue []*ByteFreq

func (pq PrioQue) Len() int { 
    return len(pq) 
}

func (pq PrioQue) Less(i, j int) bool {
	return pq[i].freq < pq[j].freq
}

func (pq PrioQue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PrioQue) Push(bf any) {
	n := len(*pq)
	*pq = append(*pq, bf.(*ByteFreq))
}

func (pq *PrioQue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil 
	*pq = old[0 : n-1]
	return item
}
