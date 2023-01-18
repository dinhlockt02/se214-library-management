package mysql

import (
	"daijoubuteam.xyz/se214-library-management/core/entity"
	coreerror "daijoubuteam.xyz/se214-library-management/core/error"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type SachRepository struct {
	*sqlx.DB
}

func (r SachRepository) GetDanhSachSach() (_ []*entity.Sach, err error) {
	tx := r.DB.MustBegin()
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			_ = tx.Commit()
		}
	}()
	sachRows, err := tx.Queryx(
		`SELECT S.MaSach, 
					   S.MaDauSach, 
					   NhaXuatBan, 
					   TriGia, 
					   NamXuatBan, 
					   TinhTrang, 
					   GhiChu 
	FROM Sach S INNER JOIN DauSach DS on S.MaDauSach = DS.MaDauSach`)
	if err != nil {
		return nil, coreerror.NewInternalServerError("database error: can't not query sach", err)
	}
	var danhSachSach []*entity.Sach
	for sachRows.Next() {
		var (
			maSach     string
			maDauSach  string
			nhaXuatBan string
			triGia     uint
			namXuatban uint
			tinhTrang  bool
			ghiChu     string
		)
		err = sachRows.Scan(&maSach, &maDauSach, &nhaXuatBan, &triGia, &namXuatban, &tinhTrang, &ghiChu)
		if err != nil {
			fmt.Println(err)
			return nil, coreerror.NewInternalServerError("database error: can't not query sach", err)
		}
		var ms *entity.ID
		ms, err = entity.StringToID(maSach)
		if err != nil {
			return nil, coreerror.NewInternalServerError("database error: can't not query sach", err)
		}
		var mds *entity.ID
		mds, err = entity.StringToID(maDauSach)
		if err != nil {
			return nil, coreerror.NewInternalServerError("database error: can't not query sach", err)
		}
		var dauSach = &entity.DauSach{MaDauSach: mds}

		if err != nil {
			return nil, coreerror.NewInternalServerError("database error: can't not query dau sach", err)
		}
		sach := entity.Sach{
			MaSach:     ms,
			DauSach:    dauSach,
			NhaXuatBan: nhaXuatBan,
			TriGia:     triGia,
			NamXuatBan: namXuatban,
			TinhTrang:  tinhTrang,
			GhiChu:     ghiChu,
		}
		danhSachSach = append(danhSachSach, &sach)
	}
	for i, _ := range danhSachSach {
		danhSachSach[i].DauSach, err = NewDauSachRepository(r.DB).getDauSachWithTx(danhSachSach[i].DauSach.MaDauSach, tx)
	}
	return danhSachSach, nil
}

func (r SachRepository) GetSach(maSach *entity.ID) (_ *entity.Sach, err error) {
	tx := r.DB.MustBegin()
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			_ = tx.Commit()
		}
	}()

	var sach *entity.Sach
	if sach, err = r.getSachWithTx(maSach, tx); err != nil {
		return nil, err
	}

	return sach, nil
}

func (r SachRepository) getSachWithTx(maSach *entity.ID, tx *sqlx.Tx) (_ *entity.Sach, err error) {
	var sach *entity.Sach
	row := tx.QueryRowx(
		`SELECT S.MaDauSach, 
					   NhaXuatBan, 
					   TriGia, 
					   NamXuatBan, 
					   TinhTrang, 
					   GhiChu 
	FROM Sach S WHERE S.MaSach = ?`, maSach)
	if row.Err() == sql.ErrNoRows {
		return nil, coreerror.NewNotFoundError("sach not found", err)
	}
	if row.Err() != nil {
		return nil, coreerror.NewInternalServerError("database error: can't not query sach", err)
	}
	var (
		maDauSach  string
		nhaXuatBan string
		triGia     uint
		namXuatban uint
		tinhTrang  bool
		ghiChu     string
	)
	err = row.Scan(&maDauSach, &nhaXuatBan, &triGia, &namXuatban, &tinhTrang, &ghiChu)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, coreerror.NewNotFoundError("sach not found", err)
		}
		return nil, coreerror.NewInternalServerError("database error: can't not query sach", err)
	}
	var mds *entity.ID
	mds, err = entity.StringToID(maDauSach)
	if err != nil {
		return nil, coreerror.NewInternalServerError("database error: can't not query sach", err)
	}
	var dauSach = &entity.DauSach{MaDauSach: mds}

	if err != nil {
		return nil, coreerror.NewInternalServerError("database error: can't not query dau sach", err)
	}
	sach = &entity.Sach{
		MaSach:     maSach,
		DauSach:    dauSach,
		NhaXuatBan: nhaXuatBan,
		TriGia:     triGia,
		NamXuatBan: namXuatban,
		TinhTrang:  tinhTrang,
		GhiChu:     ghiChu,
	}
	sach.DauSach, err = NewDauSachRepository(r.DB).getDauSachWithTx(sach.DauSach.MaDauSach, tx)
	return sach, err
}

func (r SachRepository) UpdateSach(sach *entity.Sach) (_ *entity.Sach, err error) {
	tx := r.DB.MustBegin()
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			_ = tx.Commit()
		}
	}()
	_, err = tx.Exec(
		`UPDATE Sach SET 
                NhaXuatBan = ? , 
                TriGia = ? , 
                NamXuatBan = ? , 
                TinhTrang = ?, 
                GhiChu = ? 
            WHERE MaSach = ?`,
		sach.NhaXuatBan, sach.TriGia, sach.NamXuatBan, sach.TinhTrang, sach.GhiChu, sach.MaSach)
	if err != nil {
		return nil, coreerror.NewInternalServerError("database error: can't not update database", err)
	}
	return sach, nil
}

func NewSachRepository(db *sqlx.DB) SachRepository {
	return SachRepository{
		db,
	}
}
