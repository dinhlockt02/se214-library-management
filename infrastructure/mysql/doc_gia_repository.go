package mysql

import (
	"daijoubuteam.xyz/se214-library-management/core/entity"
	coreerror "daijoubuteam.xyz/se214-library-management/core/error"
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type DocGiaRepository struct {
	db *sqlx.DB
}

func NewDocGiaRepository(db *sqlx.DB) *DocGiaRepository {
	return &DocGiaRepository{
		db: db,
	}
}

func (repo *DocGiaRepository) CreateDocGia(docGia *entity.DocGia) (_ *entity.DocGia, err error) {
	tx := repo.db.MustBegin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	exec := `
	INSERT INTO DocGia (MaDocGia, HoTen, MaLoaiDocGia, NgaySinh, DiaChi, Email, NgayLapThe, NgayHetHan, TongNo) 
	VALUES (?, ?, ?, ? , ?, ?, ?, ?, ?)
	`
	_, err = tx.Exec(exec, docGia.MaDocGia.String(), docGia.HoTen, docGia.LoaiDocGia.MaLoaiDocGia.String(), docGia.NgaySinh, docGia.DiaChi, docGia.Email, docGia.NgayLapThe, docGia.NgayHetHan, docGia.TongNo)

	if err != nil {
		return nil, coreerror.NewInternalServerError("database error: insert new doc gia failed", err)
	}
	return docGia, nil
}

func (repo *DocGiaRepository) GetDocGia(maDocGia *entity.ID) (_ *entity.DocGia, err error) {
	tx := repo.db.MustBegin()
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			_ = tx.Commit()
		}
	}()
	return repo.getDocGiaWithTx(tx, maDocGia)
}

func (repo *DocGiaRepository) getDocGiaWithTx(tx *sqlx.Tx, maDocGia *entity.ID) (*entity.DocGia, error) {
	stmt, err := tx.Prepare(`
	SELECT HoTen, NgaySinh, DiaChi, Email, NgayLapThe, NgayHetHan, TongNo, LoaiDocGia.MaLoaiDocGia, TenLoaiDocGia 
	FROM DocGia 
	INNER JOIN LoaiDocGia ON DocGia.MaLoaiDocGia = LoaiDocGia.MaLoaiDocGia
	WHERE MaDocGia = ?;
	`)

	if err != nil {
		return nil, coreerror.NewInternalServerError("database error: can't not prepare query", err)
	}
	row := stmt.QueryRow(maDocGia.String())
	docGia := &entity.DocGia{LoaiDocGia: &entity.LoaiDocGia{}}
	var maLoaiDocGiaDB string = ""
	err = row.Scan(&(docGia.HoTen), &(docGia.NgaySinh),
		&(docGia.DiaChi), &(docGia.Email), &(docGia.NgayLapThe),
		&(docGia.NgayHetHan), &(docGia.TongNo), &maLoaiDocGiaDB,
		&(docGia.LoaiDocGia.TenLoaiDocGia))
	if err == sql.ErrNoRows {
		return nil, coreerror.NewNotFoundError("doc gia not found", err)
	} else if err != nil {
		return nil, coreerror.NewInternalServerError("scan query failed", err)
	}
	maLoaiDocGia, err := entity.StringToID(maLoaiDocGiaDB)
	if err != nil {
		return nil, coreerror.NewInternalServerError("database error: can't not convert from string to id", err)
	}
	docGia.LoaiDocGia.MaLoaiDocGia = maLoaiDocGia
	docGia.MaDocGia = maDocGia
	return docGia, nil
}

func (repo *DocGiaRepository) GetDanhSachDocGia() (_ []*entity.DocGia, err error) {
	tx := repo.db.MustBegin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	stmt, err := tx.Prepare(`
		SELECT MaDocGia, HoTen, NgaySinh, DiaChi, Email, NgayLapThe, NgayHetHan, TongNo, LoaiDocGia.MaLoaiDocGia, TenLoaiDocGia 
	FROM DocGia 
	INNER JOIN LoaiDocGia ON DocGia.MaLoaiDocGia = LoaiDocGia.MaLoaiDocGia
`)
	if err != nil {
		return nil, coreerror.NewInternalServerError("database error: can't not prepare query", err)
	}
	rows, err := stmt.Query()
	defer func() {
		rows.Close()
	}()
	if err != nil {
		return nil, coreerror.NewInternalServerError("database error: can't not execute query", err)
	}
	danhSachDocGia := make([]*entity.DocGia, 0)
	for rows.Next() {
		docGia := &entity.DocGia{LoaiDocGia: &entity.LoaiDocGia{}}
		var maLoaiDocGiaDB string = ""
		var maDocGiaDB string = ""
		err = rows.Scan(&maDocGiaDB, &(docGia.HoTen), &(docGia.NgaySinh),
			&(docGia.DiaChi), &(docGia.Email), &(docGia.NgayLapThe),
			&(docGia.NgayHetHan), &(docGia.TongNo), &maLoaiDocGiaDB,
			&(docGia.LoaiDocGia.TenLoaiDocGia))
		if err == sql.ErrNoRows {
			return nil, coreerror.NewNotFoundError("doc gia not found", err)
		} else if err != nil {
			fmt.Println(err)
			return nil, coreerror.NewInternalServerError("scan query failed", err)
		}
		maLoaiDocGia, err := entity.StringToID(maLoaiDocGiaDB)
		maDocGia, err := entity.StringToID(maDocGiaDB)
		if err != nil {
			return nil, coreerror.NewInternalServerError("database error: can't not convert from string to id", err)
		}
		docGia.LoaiDocGia.MaLoaiDocGia = maLoaiDocGia
		docGia.MaDocGia = maDocGia
		danhSachDocGia = append(danhSachDocGia, docGia)
	}
	return danhSachDocGia, nil
}

func (repo *DocGiaRepository) UpdateDocGia(docGia *entity.DocGia) (_ *entity.DocGia, err error) {
	tx := repo.db.MustBegin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	exec := `
	UPDATE DocGia
	SET HoTen = ?, MaLoaiDocGia = ?, NgaySinh = ?, DiaChi = ?, Email = ?, NgayLapThe = ?, NgayHetHan = ?, TongNo = ?
	WHERE MaDocGia = ?
	`
	_, err = tx.Exec(exec, docGia.HoTen, docGia.LoaiDocGia.MaLoaiDocGia.String(), docGia.NgaySinh, docGia.DiaChi, docGia.Email, docGia.NgayLapThe, docGia.NgayHetHan, docGia.TongNo, docGia.MaDocGia.String())
	if err != nil {
		return nil, coreerror.NewInternalServerError("database error: update doc gia failed", err)
	}
	return docGia, nil
}

func (repo *DocGiaRepository) RemoveDocGia(maDocGia *entity.ID) (err error) {
	tx := repo.db.MustBegin()
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			_ = tx.Commit()
		}
	}()
	exec := `
	DELETE 
	FROM DocGia
	WHERE MaDocGia = ?
	`
	_, err = tx.Exec(exec, maDocGia.String())
	if err != nil {
		if driverError, ok := err.(*mysql.MySQLError); ok {
			return DriverErrorHandling(driverError)
		}
		return coreerror.NewInternalServerError("database error: insert new doc gia failed", err)
	}

	return nil
}
