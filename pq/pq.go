package pq

import "github.com/kelbwah/pq-374/heap"

type PriorityQueue struct {
    Queue *heap.MaxHeap
}

// O(logn)
func (pq *PriorityQueue) Enqueue(value int) {
    pq.Queue.Insert(value)
}

// O(1)
func (pq *PriorityQueue) Peek() (int, error) {
    return pq.Queue.GetMax()
}

// O(logn)
func (pq *PriorityQueue) Dequeue() (int, error) {
    return pq.Queue.ExtractMax()
}

// O(1)
func (pq *PriorityQueue) IsEmpty() bool {
    return len(pq.Queue.List) == 0
}

// O(logn)
func (pq *PriorityQueue) ChangePriorityByIndex(idx, newVal int) {
    pq.Queue.UpdateByIndex(idx, newVal)
}

// O(n)
func (pq *PriorityQueue) ChangePriority(oldVal, newVal int) {
    pq.Queue.Update(oldVal, newVal)
}

// O(1)
func CreatePriorityQueue() *PriorityQueue {
    return &PriorityQueue{
        Queue: heap.CreateMaxHeap([]int{}),
    }
}
