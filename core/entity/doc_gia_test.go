package entity

import (
	"fmt"
	"testing"
	"time"

	"daijoubuteam.xyz/se214-library-management/utils"
	"github.com/stretchr/testify/assert"
)

func TestNewDocGia(t *testing.T) {

	tests := []struct {
		hoTen      string
		loaiDocGia *LoaiDocGia
		ngaySinh   time.Time
		diaChi     string
		email      string
		ngayLapThe time.Time
		ngayHetHan time.Time
	}{
		{
			hoTen:      "mock ho ten",
			loaiDocGia: NewLoaiDocGia("mock loai doc gia"),
			ngaySinh:   time.Date(2002, 2, 2, 0, 0, 0, 0, time.UTC),
			diaChi:     "mock dia chi",
			email:      "test@email.com",
			ngayLapThe: time.Now(),
			ngayHetHan: time.Now().AddDate(0, 6, 0),
		},
		{
			hoTen:      "mock ho ten",
			loaiDocGia: nil,
			ngaySinh:   time.Date(2002, 2, 2, 0, 0, 0, 0, time.UTC),
			diaChi:     "mock dia chi",
			email:      "test@email.com",
			ngayLapThe: time.Date(2022, 2, 2, 0, 0, 0, 0, time.UTC),
			ngayHetHan: time.Date(2002, 8, 2, 0, 0, 0, 0, time.UTC),
		},
	}

	for _, value := range tests {
		docGia := NewDocGia(value.hoTen, value.loaiDocGia, &value.ngaySinh, value.diaChi, value.email, &value.ngayLapThe, &value.ngayHetHan)

		if assert.NotNil(t, docGia) {
			assert.Equal(t, value.hoTen, docGia.HoTen)
			assert.Equal(t, value.loaiDocGia, docGia.LoaiDocGia)
			assert.Equal(t, &value.ngaySinh, docGia.NgaySinh)
			assert.Equal(t, value.diaChi, docGia.DiaChi)
			assert.Equal(t, value.email, docGia.Email)
			assert.Equal(t, &value.ngayLapThe, docGia.NgayLapThe)
			assert.Equal(t, &value.ngayHetHan, docGia.NgayHetHan)
			assert.NotNil(t, docGia.MaDocGia)
		}
	}

}

func TestIsValid(t *testing.T) {

	tests := []struct {
		testcasename string
		docGia       *DocGia
		tuoiToiDa    uint
		tuoiToiThieu uint
		thoiHanThe   uint
		want         bool
	}{
		//region test case
		{
			testcasename: "Valid data",
			tuoiToiThieu: 20,
			tuoiToiDa:    60,
			thoiHanThe:   6,
			docGia: &DocGia{
				MaDocGia:   utils.Ptr(NewID()),
				HoTen:      "mock ho ten",
				LoaiDocGia: NewLoaiDocGia("mock loai doc gia"),
				NgaySinh:   utils.Ptr(time.Now().AddDate(-20, 0, 0)),
				DiaChi:     "mock dia chi",
				Email:      "mock@email.com",
				NgayLapThe: utils.Ptr(time.Now()),
				NgayHetHan: utils.Ptr(time.Now().AddDate(0, 6, 0)),
				TongNo:     0,
			},
			want: true,
		},
		{
			testcasename: "invalid ma doc gia: nil",
			tuoiToiThieu: 20,
			tuoiToiDa:    60,
			thoiHanThe:   6,
			docGia: &DocGia{
				MaDocGia:   nil,
				HoTen:      "mock ho ten",
				LoaiDocGia: NewLoaiDocGia("mock loai doc gia"),
				NgaySinh:   utils.Ptr(time.Now().AddDate(-20, 0, 0)),
				DiaChi:     "mock dia chi",
				Email:      "mock@email.com",
				NgayLapThe: utils.Ptr(time.Now()),
				NgayHetHan: utils.Ptr(time.Now().AddDate(0, 6, 0)),
				TongNo:     0,
			},
			want: false,
		},
		{
			testcasename: "invalid data: loai doc gia is nil",
			tuoiToiThieu: 20,
			tuoiToiDa:    60,
			thoiHanThe:   6,
			docGia: &DocGia{
				MaDocGia:   utils.Ptr(NewID()),
				HoTen:      "mock ho ten",
				LoaiDocGia: nil,
				NgaySinh:   utils.Ptr(time.Now().AddDate(-20, 0, 0)),
				DiaChi:     "mock dia chi",
				Email:      "mock@email.com",
				NgayLapThe: utils.Ptr(time.Now()),
				NgayHetHan: utils.Ptr(time.Now().AddDate(0, 6, 0)),
				TongNo:     0,
			},
			want: false,
		},
		{
			testcasename: "invalid ngay sinh: nil",
			tuoiToiThieu: 20,
			tuoiToiDa:    60,
			thoiHanThe:   6,
			docGia: &DocGia{
				MaDocGia:   utils.Ptr(NewID()),
				HoTen:      "mock ho ten",
				LoaiDocGia: NewLoaiDocGia("mock loai doc gia"),
				NgaySinh:   nil,
				DiaChi:     "mock dia chi",
				Email:      "mock@email.com",
				NgayLapThe: utils.Ptr(time.Now()),
				NgayHetHan: utils.Ptr(time.Now().AddDate(0, 6, 0)),
				TongNo:     0,
			},
			want: false,
		},
		{
			testcasename: "invalid age: age is negative",
			tuoiToiThieu: 20,
			tuoiToiDa:    60,
			thoiHanThe:   6,
			docGia: &DocGia{
				MaDocGia:   utils.Ptr(NewID()),
				HoTen:      "mock ho ten",
				LoaiDocGia: NewLoaiDocGia("mock loai doc gia"),
				NgaySinh:   utils.Ptr(time.Now().AddDate(20, 0, 0)),
				DiaChi:     "mock dia chi",
				Email:      "mock@email.com",
				NgayLapThe: utils.Ptr(time.Now()),
				NgayHetHan: utils.Ptr(time.Now().AddDate(0, 6, 0)),
				TongNo:     0,
			},
			want: false,
		},
		{
			testcasename: "invalid age: too young",
			tuoiToiThieu: 20,
			tuoiToiDa:    60,
			thoiHanThe:   6,
			docGia: &DocGia{
				MaDocGia:   utils.Ptr(NewID()),
				HoTen:      "mock ho ten",
				LoaiDocGia: NewLoaiDocGia("mock loai doc gia"),
				NgaySinh:   utils.Ptr(time.Now().AddDate(-19, 0, 0)),
				DiaChi:     "mock dia chi",
				Email:      "mock@email.com",
				NgayLapThe: utils.Ptr(time.Now()),
				NgayHetHan: utils.Ptr(time.Now().AddDate(0, 6, 0)),
				TongNo:     0,
			},
			want: false,
		},
		{
			testcasename: "invalid age: too old",
			tuoiToiThieu: 20,
			tuoiToiDa:    60,
			thoiHanThe:   6,
			docGia: &DocGia{
				MaDocGia:   utils.Ptr(NewID()),
				HoTen:      "mock ho ten",
				LoaiDocGia: NewLoaiDocGia("mock loai doc gia"),
				NgaySinh:   utils.Ptr(time.Now().AddDate(-61, 0, 0)),
				DiaChi:     "mock dia chi",
				Email:      "mock@email.com",
				NgayLapThe: utils.Ptr(time.Now()),
				NgayHetHan: utils.Ptr(time.Now().AddDate(0, 6, 0)),
				TongNo:     0,
			},
			want: false,
		},
		{
			testcasename: "invalid ngay het han: too early",
			tuoiToiThieu: 20,
			tuoiToiDa:    60,
			thoiHanThe:   6,
			docGia: &DocGia{
				MaDocGia:   utils.Ptr(NewID()),
				HoTen:      "mock ho ten",
				LoaiDocGia: NewLoaiDocGia("mock loai doc gia"),
				NgaySinh:   utils.Ptr(time.Now().AddDate(-20, 0, 0)),
				DiaChi:     "mock dia chi",
				Email:      "mock@email.com",
				NgayLapThe: utils.Ptr(time.Now()),
				NgayHetHan: utils.Ptr(time.Now().AddDate(0, 4, 0)),
				TongNo:     0,
			},
			want: false,
		},
		{
			testcasename: "invalid ngay het han: too late",
			tuoiToiThieu: 20,
			tuoiToiDa:    60,
			thoiHanThe:   6,
			docGia: &DocGia{
				MaDocGia:   utils.Ptr(NewID()),
				HoTen:      "mock ho ten",
				LoaiDocGia: NewLoaiDocGia("mock loai doc gia"),
				NgaySinh:   utils.Ptr(time.Now().AddDate(-20, 0, 0)),
				DiaChi:     "mock dia chi",
				Email:      "mock@email.com",
				NgayLapThe: utils.Ptr(time.Now()),
				NgayHetHan: utils.Ptr(time.Now().AddDate(0, 8, 0)),
				TongNo:     0,
			},
			want: false,
		},
		//endregion
	}

	for _, value := range tests {
		result, _ := value.docGia.IsValid(value.tuoiToiDa, value.tuoiToiThieu, value.thoiHanThe)

		assert.Equal(t, result, value.want, fmt.Sprintf("IsValid() - %v", value.testcasename))
	}
}
