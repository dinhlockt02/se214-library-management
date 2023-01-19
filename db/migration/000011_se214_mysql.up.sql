CREATE TABLE PhieuThuTienPhat (
    MaPhieuThuTienPhat  VARCHAR(255),
    MaDocGia VARCHAR(255),
    TongNo INT,
    SoTienThu INT UNSIGNED,
    ConLai INT,
    NgayThu DATE,
    CONSTRAINT PhieuThuTienPhat_FK_DocGia FOREIGN KEY (MaDocGia) REFERENCES DocGia(MaDocGia)
);