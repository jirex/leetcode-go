type NestedIterator struct {
    stack [][]*NestedInteger
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
    return &NestedIterator{[][]*NestedInteger{nestedList}}
}

func (this *NestedIterator) Next() int {
	queue := this.stack[len(this.stack)-1]
	val := queue[0].GetInteger()
	this.stack[len(this.stack)-1] = queue[1:]
	return val
}

func (this *NestedIterator) HasNext() bool {
    for len(this.stack) > 0{
		queue := this.stack[len(this.stack)-1]
		if len(queue) == 0 {
			this.stack = this.stack[:len(this.stack)-1]
			continue
		}
		nest := queue[0]
		if nest.IsInteger() {
			return true
		}
		this.stack[len(this.stack)-1] = queue[1:]
		this.stack = append(this.stack, nest.GetList())
	}
	return false
}
