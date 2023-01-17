CREATE TABLE PhieuMuon (
    MaPhieuMuon VARCHAR(255) PRIMARY KEY,
    MaDocGia VARCHAR(255),
    MaSach VARCHAR(255),
    NgayMuon DATE,
    CONSTRAINT PhieuMuon_FK_DocGia FOREIGN KEY (MaDocGia) REFERENCES DocGia(MaDocGia)
);

CREATE TABLE PhieuTra (
    MaPhieuMuon VARCHAR(255) PRIMARY KEY ,
    TienPhat INT UNSIGNED ,
    NgayTra DATE,
    GhiChu TEXT,
    CONSTRAINT PhieuTra_FK_PhieuMuon FOREIGN KEY (MaPhieuMuon) REFERENCES PhieuMuon(MaPhieuMuon)
);