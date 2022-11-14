package entity

type TacGia struct {
	MaTacGia  *ID
	TenTacGia string
}

func NewTacGia(tenTacGia string) *TacGia {
	newId := NewID()
	return &TacGia{
		MaTacGia:  &newId,
		TenTacGia: tenTacGia,
	}
}

func (tacGia *TacGia) IsValid() bool {
	return true
}
