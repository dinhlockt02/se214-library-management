CREATE TABLE IF NOT EXISTS DocGia (
    MaDocGia VARCHAR(255) PRIMARY KEY,
    HoTen VARCHAR(255),
    MaLoaiDocGia VARCHAR(255),
    NgaySinh DATE,
    DiaChi VARCHAR(255),
    Email VARCHAR(255),
    NgayLapThe DATE,
    NgayHetHan DATE,
    TongNo INT,
    CONSTRAINT DocGia_FK_MaLoaiDocGia FOREIGN KEY (MaLoaiDocGia) REFERENCES LoaiDocGia(MaLoaiDocGia),
    CONSTRAINT NgayLapThe_CK_NgayHetHan CHECK ( DocGia.NgayLapThe < DocGia.NgayHetHan )
);

