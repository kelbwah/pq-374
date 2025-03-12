package heap

import (
    "errors"
    "slices"
)

type MaxHeap struct {
    List []int
}

func (h *MaxHeap) Shiftup(idx int) {
    parent := (idx - 1) / 2
    for idx != 0 && h.List[idx] > h.List[parent] {
        h.List[idx], h.List[parent] = h.List[parent], h.List[idx]
        idx = parent
        parent = (idx - 1) / 2
    }
}

func (h *MaxHeap) Shiftdown(idx int) {
    left := (2 * idx) + 1
    right := (2 * idx) + 2
    for (left < len(h.List) && h.List[idx] < h.List[left]) || (right < len(h.List) && h.List[idx] < h.List[right]) {
        var biggest int
        if right >= len(h.List) || h.List[left] > h.List[right] {
            biggest = left
        } else {
            biggest = right  
        }
        h.List[idx], h.List[biggest] = h.List[biggest], h.List[idx]
        idx = biggest
        left = (2 * idx) + 1
        right = (2 * idx) + 2
    }
}

func (h *MaxHeap) Insert(val int) {
    h.List = append(h.List, val)
    h.Shiftup(len(h.List) - 1)
}

func (h *MaxHeap) GetMax() (int, error) {
    if len(h.List) > 0 {
        return h.List[0], nil
    } else {
        return 0, errors.New("Heap is empty!")
    }
}

func (h *MaxHeap) ExtractMax() (int, error) {
    if len(h.List) == 0 {
        return 0, errors.New("Heap is empty!")
    }
    maxVal := h.List[0]
    h.List[0], h.List[len(h.List) - 1] = h.List[len(h.List) - 1], h.List[0]
    h.List = h.List[:len(h.List) - 1]
    h.Shiftdown(0)

    return maxVal, nil
}

func (h *MaxHeap) UpdateByIndex(idx, newVal int) {
    oldVal := h.List[idx]
    h.List[idx] = newVal
    if newVal > oldVal {
        h.Shiftup(idx)
    } else {
        h.Shiftdown(idx)
    }
}

func (h *MaxHeap) Update(oldVal, newVal int) {
    if slices.Contains(h.List, oldVal) {
        h.UpdateByIndex(slices.Index(h.List, oldVal), newVal)
    }
}

func CreateMaxHeap(arr []int) *MaxHeap {
    heapedArr := make([]int, len(arr))
    copy(heapedArr, arr)
    newMaxHeap := &MaxHeap{
        List: heapedArr, 
    }
    
    for i := len(heapedArr) - 1; i >= 0; i-- {
      newMaxHeap.Shiftdown(i)
    }
    
    return newMaxHeap
}
