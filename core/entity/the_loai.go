package entity

type TheLoai struct {
	MaTheLoai  ID
	TenTheLoai string
}

func NewTheLoai(tenTheLoai string) *TheLoai {
	return &TheLoai{
		TenTheLoai: tenTheLoai,
	}
}

func (theLoai *TheLoai) IsValid() bool {
	return true
}
