package main

import "time"

type Bin struct {
	id        string
	private   bool
	createdAt time.Time
	name      string
}

type BinList struct {
	Bins []Bin
}

func newBin(id string, private bool, createdAt time.Time, name string) *Bin {
	return &Bin{id, private, createdAt, name}
}

func newBinList() *BinList {
	return &BinList{
		Bins: make([]Bin, 0), // инициализируем пустой срез
	}
}

func main() {

}
