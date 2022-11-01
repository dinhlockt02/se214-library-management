package entity

type TacGia struct {
	MaTacGia  ID
	TenTacGia string
}

func NewTacGia(tenTacGia string) *TacGia {
	return &TacGia{
		TenTacGia: tenTacGia,
	}
}

func (tacGia *TacGia) IsValid() bool {
	return true
}
