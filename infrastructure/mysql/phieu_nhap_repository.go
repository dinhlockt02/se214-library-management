package mysql

import (
	"daijoubuteam.xyz/se214-library-management/core/entity"
	coreerror "daijoubuteam.xyz/se214-library-management/core/error"
	"daijoubuteam.xyz/se214-library-management/utils"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type PhieuNhapRepository struct {
	db          *sqlx.DB
	dauSachRepo *DauSachRepository
}

func NewPhieuNhapRepository(db *sqlx.DB, dauSachRepo *DauSachRepository) *PhieuNhapRepository {
	return &PhieuNhapRepository{
		db:          db,
		dauSachRepo: dauSachRepo,
	}
}

func (repo *PhieuNhapRepository) GetDanhSachPhieuNhap() (_ []*entity.PhieuNhap, err error) {
	tx := repo.db.MustBegin()
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			_ = tx.Commit()
		}
	}()
	var danhSachPhieuNhap []*entity.PhieuNhap
	rows, err := tx.Queryx(`SELECT MaPhieuNhap,  NgayNhap AS NgayLap, TongTien FROM PhieuNhap`)
	defer func() {
		_ = rows.Close()
	}()
	if err != nil {
		return nil, coreerror.NewInternalServerError("database error: can't not query phieu nhap", err)
	}
	for rows.Next() {
		var phieuNhap = entity.PhieuNhap{}
		var maPhieuNhap string
		err = rows.Scan(&maPhieuNhap, &(phieuNhap.NgayLap), &(phieuNhap.TongTien))
		if err != nil {
			return nil, coreerror.NewInternalServerError("database error: can't not query phieu nhap", err)
		}
		phieuNhap.MaPhieuNhap, err = entity.StringToID(maPhieuNhap)
		if err != nil {
			return nil, coreerror.NewInternalServerError("database error: can't not query phieu nhap", err)
		}
		danhSachPhieuNhap = append(danhSachPhieuNhap, &phieuNhap)
	}
	for i, _ := range danhSachPhieuNhap {
		var ctPhieuNhapRow *sqlx.Rows
		ctPhieuNhapRow, err = tx.Queryx(
			`SELECT CT.MaSach, DonGia, MaDauSach, NhaXuatBan, TriGia, NamXuatBan, TinhTrang, GhiChu
	FROM Ct_PhieuNhap CT JOIN Sach S on S.MaSach = CT.MaSach
	WHERE MaPhieuNhap = ?`,
			danhSachPhieuNhap[i].MaPhieuNhap,
		)
		if err != nil {
			fmt.Println(err)
			return nil, coreerror.NewInternalServerError("database error: can't not query chi tiet phieu nhap", err)
		}
		var ctPhieuNhap []*entity.CtPhieuNhap
		for ctPhieuNhapRow.Next() {
			var sach = entity.Sach{}
			var ct = entity.CtPhieuNhap{
				Sach: &sach,
			}
			var maSach string
			var maDauSach string

			err = ctPhieuNhapRow.Scan(&maSach, &(ct.DonGia), &maDauSach, &(ct.NhaXuatBan), &(ct.TriGia), &(ct.NamXuatBan), &(ct.TinhTrang), &(ct.GhiChu))
			if err != nil {
				return nil, coreerror.NewInternalServerError("database error: can't not query chi tiet phieu nhap", err)
			}
			ct.MaSach, err = entity.StringToID(maSach)
			if err != nil {
				fmt.Println(ct.Sach)
				return nil, coreerror.NewInternalServerError("database error: can't not execute query", err)
			}
			var mds *entity.ID
			mds, err = entity.StringToID(maDauSach)
			if err != nil {
				return nil, coreerror.NewInternalServerError("database error: can't not query phieu nhap sach", err)
			}
			ct.DauSach = utils.Ptr(entity.DauSach{MaDauSach: mds})
			ctPhieuNhap = append(ctPhieuNhap, &ct)
		}
		for i, _ := range ctPhieuNhap {
			ctPhieuNhap[i].DauSach, err = repo.dauSachRepo.getDauSachWithTx(ctPhieuNhap[i].DauSach.MaDauSach, tx)
			if err != nil {
				return nil, err
			}
		}
		danhSachPhieuNhap[i].CtPhieuNhap = ctPhieuNhap
	}
	return danhSachPhieuNhap, nil
}

func (repo *PhieuNhapRepository) GetPhieuNhap(maPhieuNhap *entity.ID) (_ *entity.PhieuNhap, err error) {
	tx := repo.db.MustBegin()
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			_ = tx.Commit()
		}
	}()
	var phieuNhap = entity.PhieuNhap{}
	err = tx.QueryRowx(`SELECT NgayNhap AS NgayLap, TongTien FROM PhieuNhap WHERE MaPhieuNhap = ?`, maPhieuNhap).Scan(&(phieuNhap.NgayLap), &(phieuNhap.TongTien))
	phieuNhap.MaPhieuNhap = maPhieuNhap
	ctPhieuNhapRow, err := tx.Queryx(
		`SELECT CT.MaSach, DonGia, MaDauSach, NhaXuatBan, TriGia, NamXuatBan, TinhTrang, GhiChu
	FROM Ct_PhieuNhap CT JOIN Sach S on S.MaSach = CT.MaSach
	WHERE MaPhieuNhap = ?`, maPhieuNhap)
	if err != nil {
		fmt.Println(err)
		return nil, coreerror.NewInternalServerError("database error: can't not query chi tiet phieu nhap", err)
	}
	var ctPhieuNhap []*entity.CtPhieuNhap
	for ctPhieuNhapRow.Next() {
		var sach = entity.Sach{}
		var ct = entity.CtPhieuNhap{
			Sach: &sach,
		}
		var maSach string
		var maDauSach string

		err = ctPhieuNhapRow.Scan(&maSach, &(ct.DonGia), &maDauSach, &(ct.NhaXuatBan), &(ct.TriGia), &(ct.NamXuatBan), &(ct.TinhTrang), &(ct.GhiChu))
		if err != nil {
			return nil, coreerror.NewInternalServerError("database error: can't not query chi tiet phieu nhap", err)
		}
		ct.MaSach, err = entity.StringToID(maSach)
		if err != nil {
			fmt.Println(ct.Sach)
			return nil, coreerror.NewInternalServerError("database error: can't not execute query", err)
		}
		mds, _ := entity.StringToID(maDauSach)
		ct.DauSach = utils.Ptr(entity.DauSach{MaDauSach: mds})
		ctPhieuNhap = append(ctPhieuNhap, &ct)
	}
	for i, _ := range ctPhieuNhap {
		ctPhieuNhap[i].DauSach, err = repo.dauSachRepo.getDauSachWithTx(ctPhieuNhap[i].DauSach.MaDauSach, tx)
		if err != nil {
			return nil, err
		}
	}
	phieuNhap.CtPhieuNhap = ctPhieuNhap
	return &phieuNhap, nil
}

func (repo *PhieuNhapRepository) CreatePhieuNhap(phieuNhap *entity.PhieuNhap) (_ *entity.PhieuNhap, err error) {
	tx := repo.db.MustBegin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	phieuNhapExec := `INSERT INTO PhieuNhap(MaPhieuNhap, NgayNhap, TongTien) VALUES (?, ?, ?)`
	_, err = tx.Exec(phieuNhapExec, phieuNhap.MaPhieuNhap.String(), phieuNhap.NgayLap, phieuNhap.TongTien)
	if err != nil {
		return nil, coreerror.NewInternalServerError("database error: can't not create phieu nhap", err)
	}

	for _, ctPhieuNhap := range phieuNhap.CtPhieuNhap {
		var sach *entity.Sach
		sach, err = repo.createSach(tx, ctPhieuNhap.Sach)
		if err != nil {
			return nil, err
		}
		err = repo.createCtPhieuNhap(tx, sach.MaSach, phieuNhap.MaPhieuNhap, ctPhieuNhap.DonGia)
		if err != nil {
			return nil, err
		}
	}

	return phieuNhap, nil
}

func (repo *PhieuNhapRepository) createSach(tx *sqlx.Tx, sach *entity.Sach) (_ *entity.Sach, err error) {
	sachExec := `INSERT INTO 
    Sach(MaSach, 
         MaDauSach, 
         NhaXuatBan, 
         TriGia, 
         NamXuatBan, 
         TinhTrang, 
         GhiChu) 
	VALUES (?,?,?,?,?,?,?)`
	_, err = tx.Exec(sachExec,
		sach.MaSach,
		sach.DauSach.MaDauSach,
		sach.NhaXuatBan,
		sach.TriGia,
		sach.NamXuatBan,
		sach.TinhTrang,
		sach.GhiChu)
	if err != nil {
		if driverError, ok := err.(*mysql.MySQLError); ok {
			return nil, DriverErrorHandling(driverError)
		}
		return nil, coreerror.NewInternalServerError("database error: can't not create chi tiet phieu nhap", err)
	}
	return sach, nil
}

func (repo *PhieuNhapRepository) createCtPhieuNhap(tx *sqlx.Tx, maSach *entity.ID, maPhieuNhap *entity.ID, donGia uint) error {
	ctPhieuNhapExec := `INSERT INTO 
    Ct_PhieuNhap(MaPhieuNhap, MaSach, DonGia) 
	VALUES (?,?,?)`
	_, err := tx.Exec(ctPhieuNhapExec,
		maPhieuNhap.String(), maSach.String(), donGia)
	if err != nil {
		if driverError, ok := err.(*mysql.MySQLError); ok {
			return DriverErrorHandling(driverError)
		}
		err = coreerror.NewInternalServerError("database error: can't not create chi tiet phieu nhap", err)
	}
	return err
}

func (repo *PhieuNhapRepository) UpdatePhieuNhap(phieuNhap *entity.PhieuNhap) (_ *entity.PhieuNhap, err error) {
	tx := repo.db.MustBegin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	phieuNhapExec := `UPDATE PhieuNhap SET NgayNhap = ? WHERE MaPhieuNhap = ?`
	_, err = tx.Exec(phieuNhapExec, phieuNhap.NgayLap, phieuNhap.MaPhieuNhap.String())
	if err != nil {
		return nil, coreerror.NewInternalServerError("database error: can't not create phieu nhap", err)
	}
	return phieuNhap, nil
}

func (repo *PhieuNhapRepository) RemovePhieuNhap(maPhieuNhap *entity.ID) (err error) {
	tx := repo.db.MustBegin()
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			_ = tx.Commit()
		}
	}()

	// Get all ma sach in phieu nhap
	maSachRow, err := tx.Queryx(
		`SELECT MaSach From Ct_PhieuNhap WHERE MaPhieuNhap = ?`,
		maPhieuNhap.String(),
	)
	defer func() {
		_ = maSachRow.Close()
	}()
	var maSach []interface{}
	for maSachRow.Next() {
		var ms string
		err = maSachRow.Scan(&ms)
		maSach = append(maSach, ms)
	}
	if err != nil {
		fmt.Println(err)
		return coreerror.NewInternalServerError("database error: can't not query ma sach", err)
	}
	// Delete sach if exists
	if len(maSach) > 0 {
		_, err = tx.Exec(
			`DELETE FROM Ct_PhieuNhap WHERE MaPhieuNhap = ?`,
			maPhieuNhap,
		)
		if err != nil {
			if driverError, ok := err.(*mysql.MySQLError); ok {
				return DriverErrorHandling(driverError)
			}
			return coreerror.NewInternalServerError("database error: can't not delete chi tiet phieu nhap", err)
		}
		var query string
		var args []interface{}
		query, args, err = sqlx.In(
			`DELETE FROM Sach WHERE Sach.MaSach IN (?);`,
			maSach,
		)
		_, err = tx.Exec(tx.Rebind(query), args...)
		if err != nil {
			return coreerror.NewInternalServerError("database error: can't not delete sach", err)
		}
	}

	// Delete phieu nhap
	_, err = tx.Exec(
		`DELETE FROM PhieuNhap WHERE MaPhieuNhap = ?`,
		maPhieuNhap.String(),
	)
	if err != nil {
		if driverError, ok := err.(*mysql.MySQLError); ok {
			return DriverErrorHandling(driverError)
		}
		return coreerror.NewInternalServerError("database error: can't not delete phieu nhap", err)
	}

	return nil
}
