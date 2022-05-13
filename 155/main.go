package main

type Minstack struct {
	elements, min []int
	l  int
}

func Constructor() Minstack {
	return Minstack{make([]int, 0), make([]int, 0), 0}
}

func (this *Minstack)Push(x int)  {
	this.elements = append(this.elements, x)

	if this.l == 0 {
		this.min = append(this.min, x)
	} else {
		min := this.GetMin()
		if x < min {
			this.min = append(this.min, x)
		} else {
			this.min = append(this.min, min)
		}
	}

	this.l++
}

func (this *Minstack)Pop() {
	this.l--
	this.min = this.min[:this.l]
	this.elements = this.elements[:this.l]
}

func (this *Minstack) Top() int {
	return this.elements[this.l-1]
}

func (this *Minstack)GetMin() int {
	return this.min[this.l - 1]
}

