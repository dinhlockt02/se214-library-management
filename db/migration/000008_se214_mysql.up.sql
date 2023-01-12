CREATE TABLE PhieuNhap (
    MaPhieuNhap VARCHAR(255) PRIMARY KEY,
    NgayNhap DATE,
    TongTien INT UNSIGNED
);

CREATE TABLE Sach (
    MaSach VARCHAR(255) PRIMARY KEY ,
    MaDauSach VARCHAR(255),
    NhaXuatBan VARCHAR(255),
    TriGia INT UNSIGNED,
    NamXuatBan SMALLINT UNSIGNED,
    TinhTrang BOOLEAN,
    GhiChu TEXT,
    CONSTRAINT Sach_FK_DauSach FOREIGN KEY (MaDauSach) REFERENCES DauSach(MaDauSach)
);

CREATE TABLE Ct_PhieuNhap (
    MaPhieuNhap VARCHAR(255),
    MaSach VARCHAR(255) PRIMARY KEY,
    DonGia INT UNSIGNED,
    CONSTRAINT Ct_PhieuNhap_FK_PhieuNhap FOREIGN KEY (MaPhieuNhap) REFERENCES PhieuNhap(MaPhieuNhap),
    CONSTRAINT Ct_PhieuNhap_FK_Sach FOREIGN KEY (MaSach) REFERENCES Sach(MaSach)
);