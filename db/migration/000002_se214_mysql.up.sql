CREATE TABLE IF NOT EXISTS ThamSo (
  Id ENUM('1') NOT NULL,
  ThoiHanThe INT,
  TuoiToiThieu INT,
  TuoiToiDa INT,
  ThoiGianLuuHanh INT,
  SoNgayMuonMax INT,
  SoSachMuonMax INT,
  MucThuTienPhat INT,
  DefaultPassword VARCHAR(256)
);
INSERT INTO ThamSo
VALUES (
    '1',
    6,
    18,
    55,
    8,
    4,
    5,
    1000,
    '12345678'
  );