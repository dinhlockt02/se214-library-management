package entity

type TheLoai struct {
	MaTheLoai  *ID
	TenTheLoai string
}

func NewTheLoai(tenTheLoai string) *TheLoai {
	newId := NewID()
	return &TheLoai{
		MaTheLoai:  &newId,
		TenTheLoai: tenTheLoai,
	}
}

func (theLoai *TheLoai) IsValid() bool {
	return true
}
