CREATE TABLE IF NOT EXISTS ThamSo (
  Id ENUM('1') NOT NULL,
  ThoiHanThe INT,
  TuoiToiThieu INT,
  TuoiToiDa INT,
  DefaultPassword VARCHAR(256)
);
INSERT INTO ThamSo (Id, ThoiHanThe, TuoiToiThieu, TuoiToiDa, DefaultPassword)
VALUES (
    '1',
    6,
    18,
    55,
    '12345678'
  );