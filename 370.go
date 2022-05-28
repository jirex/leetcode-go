type Difference struct {
	diff []int
}

func Constructor(len int) Difference {
	diff := make([]int, len)
	return Difference{diff: diff}
}

func (d *Difference) increment(i, j, val int) {
	d.diff[i] += val

	if j+1 < len(diff) {
		d.diff[j+1] -= val
	}
}

func (d *Difference) result() []int {
	res := make([]int, len(d.diff))

	res[0] = d.diff[0]
	for i, v := range d.diff {
		if i > 0 {
			res[i] = res[i-1] + v
		}
	}
	return res
}

