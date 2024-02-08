package utils

// import "sort"

// type PriorityQueue struct {
// 	items  []*Node
// 	length int
// }

// func NewPriorityQueue() *PriorityQueue {
// 	return &PriorityQueue{
// 		items:  make([]*Node, 0),
// 		length: 0,
// 	}
// }

// func (q *PriorityQueue) IsEmpty() bool {
// 	return q.length == 0
// }

// func (q *PriorityQueue) Enqueue(item *Node) {
// 	if len(q.items) == 0 {
// 		q.items = append(q.items, item)
// 		q.length++
// 		return
// 	}

// 	insertionInd := 0
// 	for _, qitem := range q.items {
// 		if qitem.distance < item.distance {
// 			insertionInd++
// 			continue
// 		}
// 	}

// 	q.items = append(q.items, &Node{})

// 	copy(q.items[insertionInd+1:], q.items[insertionInd:])
// 	q.items[insertionInd] = item
// 	q.length++
// }

// func (q *PriorityQueue) DecreaseDist(item *Node, dist int) {
// 	item.distance = dist
// 	sort.Slice(q.items, func(i, j int) bool {
// 		return q.items[i].distance < q.items[j].distance
// 	})
// 	// nodeInd := 0
// 	// for ind, qitem := range q.items {
// 	// 	if qitem == item {
// 	// 		nodeInd = ind
// 	// 		break
// 	// 	}
// 	// }

// 	// newNodeInd := 0
// 	// for _, qitem := range q.items {
// 	// 	if qitem.distance < dist {
// 	// 		newNodeInd++
// 	// 	}
// 	// }

// 	// fmt.Println(nodeInd, newNodeInd)

// 	// q.items[newNodeInd], q.items[nodeInd] = q.items[nodeInd], q.items[newNodeInd]
// }

// func (q *PriorityQueue) Dequeue() (*Node, bool) {
// 	if q.length > 0 {
// 		item := q.items[0]
// 		q.items = q.items[1:]
// 		q.length--
// 		return item, true
// 	}
// 	var zero *Node
// 	return zero, false
// }

// func (q *PriorityQueue) Items() []*Node {
// 	return q.items
// }

// func (q *PriorityQueue) Length() int {
// 	return q.length
// }
