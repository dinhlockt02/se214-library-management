package entity

import "time"

type PhieuTra struct {
	MaPhieuTraSach *ID
	DocGia         *DocGia
	NgayTra        *time.Time
	TienPhatKyNay  uint
}
