package pq

import "github.com/kelbwah/pq-374/heap"

type PriorityQueue struct {
    Queue *heap.MaxHeap
}

func (pq *PriorityQueue) Enqueue(value int) {
    pq.Queue.Insert(value)
}

func (pq *PriorityQueue) Peek() (int, error) {
    return pq.Queue.GetMax()
}

func (pq *PriorityQueue) Dequeue() (int, error) {
    return pq.Queue.ExtractMax()
}

func (pq *PriorityQueue) IsEmpty() bool {
    return len(pq.Queue.List) == 0
}

func (pq *PriorityQueue) ChangePriorityByIndex(idx, newVal int) {
    pq.Queue.UpdateByIndex(idx, newVal)
}

func (pq *PriorityQueue) ChangePriority(oldVal, newVal int) {
    pq.Queue.Update(oldVal, newVal)
}

func CreatePriorityQueue() *PriorityQueue {
    return &PriorityQueue{
        Queue: heap.CreateMaxHeap([]int{}),
    }
}
