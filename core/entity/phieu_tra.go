package entity

import "time"

type PhieuTra struct {
	MaPhieuTraSach ID
	MaDocGia       ID
	NgayTra        time.Time
	TienPhatKyNay  uint
}
