package bins

import "time"

type Bin struct {
	Id        string    `json:"id"`
	Private   bool      `json:"private"`
	CreatedAt time.Time `json:"createdAt"`
	Name      string    `json:"name"`
}

type BinList struct {
	Bins []Bin `json:"bins"`
}

func NewBin(id string, private bool, createdAt time.Time, name string) *Bin {
	return &Bin{id, private, createdAt, name}
}

func NewBinList() *BinList {
	return &BinList{
		Bins: make([]Bin, 0), // инициализируем пустой срез
	}
}
