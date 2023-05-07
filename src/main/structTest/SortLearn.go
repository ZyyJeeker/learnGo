package structTest

import (
	"fmt"
	"sort"
)

type Data struct {
	nodes []Node "一个Node切片"
}

func (d *Data) index(i int) Node {
	return d.nodes[i]
}

type Node struct {
	num1 int
	num2 int
}

type Sorter interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

func (d *Data) Len() int {
	return len(d.nodes)
}

func (d *Data) Less(i, j int) bool {
	switch {
	case d.index(i).num1 < d.index(j).num1:
		return true
	case d.index(i).num1 > d.index(j).num1:
		return false
	default:
		return d.index(i).num2 < d.index(j).num2
	}
}

func (d *Data) Swap(i, j int) {
	// go的特色
	d.nodes[i], d.nodes[j] = d.nodes[j], d.nodes[i]
}

// Sort 简单的冒泡排序
func Sort(data Sorter) {
	for pass := 1; pass < data.Len(); pass++ {
		for i := 0; i < data.Len()-pass; i++ {
			if data.Less(i+1, i) {
				data.Swap(i, i+1)
			}
		}
	}
}

func SortTest() {
	data := &Data{[]Node{{1, 2}, {1, 3}, {3, 4}, {3, 5}}}
	sort.Sort(data)
	fmt.Println(data)
	data1 := &Data{[]Node{{4, 2}, {1, 3}, {3, 1}, {3, 5}}}
	Sort(data1)
	fmt.Println(data1)

	days()
}

type day struct {
	num       int
	shortName string
	longName  string
}
type dayArray struct {
	data []*day
}

func (p *dayArray) Len() int           { return len(p.data) }
func (p *dayArray) Less(i, j int) bool { return p.data[i].num < p.data[j].num }
func (p *dayArray) Swap(i, j int)      { p.data[i], p.data[j] = p.data[j], p.data[i] }

func days() {
	Sunday := day{0, "SUN", "Sunday"}
	Monday := day{1, "MON", "Monday"}
	Tuesday := day{2, "TUE", "Tuesday"}
	Wednesday := day{3, "WED", "Wednesday"}
	Thursday := day{4, "THU", "Thursday"}
	Friday := day{5, "FRI", "Friday"}
	Saturday := day{6, "SAT", "Saturday"}
	data := []*day{&Tuesday, &Thursday, &Wednesday, &Sunday, &Monday, &Friday, &Saturday}
	a := dayArray{data}
	sort.Sort(&a)
	if !sort.IsSorted(&a) {
		panic("fail")
	}
	for _, d := range data {
		fmt.Printf("%s ", d.longName)
	}
	fmt.Printf("\n")
}
