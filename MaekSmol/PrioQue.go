package MaekSmol

type PrioQue []ByteNode

func (pq PrioQue) Len() int { 
    return len(pq) 
}

func (pq PrioQue) Less(i, j int) bool {
    if (pq[i].freq == pq[j].freq) {
        return pq[i].val < pq[j].val
    }
	return pq[i].freq < pq[j].freq
}

func (pq PrioQue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PrioQue) Push(bf any) {
	*pq = append(*pq, bf.(ByteNode))
}

func (pq *PrioQue) Pop() any {
	old := *pq
	n := len(old)
	bf := old[n-1]
	*pq = old[0:n-1]
	return bf
}
