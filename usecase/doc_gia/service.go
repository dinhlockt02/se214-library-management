package docgia

import (
	"time"

	"daijoubuteam.xyz/se214-library-management/core/entity"
	businessError "daijoubuteam.xyz/se214-library-management/core/error"
	"daijoubuteam.xyz/se214-library-management/core/repository"
)

type DocGiaService struct {
	loaiDocGiaRepo repository.LoaiDocGiaRepository
	docGiaRepo     repository.DocGiaRepository
	thamSoRepo     repository.ThamSoRepository
}

func NewDocGiaService(docGiaRepo repository.DocGiaRepository) *DocGiaService {
	return &DocGiaService{
		docGiaRepo: docGiaRepo,
	}
}
func (service *DocGiaService) GetDanhSachDocGia() ([]*entity.DocGia, error) {
	danhSachDocGia, err := service.docGiaRepo.GetDanhSachDocGia()
	if err != nil {
		return nil, err
	}
	return danhSachDocGia, err
}

func (service *DocGiaService) GetDocGia(maDocGia *entity.ID) (*entity.DocGia, error) {
	docGia, err := service.docGiaRepo.GetDocGia(maDocGia)
	if err != nil {
		return nil, err
	}

	if docGia == nil {
		return nil, businessError.NewBusinessError("doc gia not found")
	}

	return docGia, err
}

func (service *DocGiaService) CreateDocGia(hoTen string, loaiDocGia *entity.ID, ngaySinh *time.Time, diaChi string, email string, ngayLapThe *time.Time) (*entity.DocGia, error) {

	loaiDocGiaRs, err := service.loaiDocGiaRepo.GetLoaiDocGia(loaiDocGia)
	if err != nil {
		return nil, err
	}

	if loaiDocGiaRs == nil {
		return nil, businessError.NewBusinessError("loai doc gia not found")
	}

	// Get tham so

	thoiHanThe := service.thamSoRepo.GetThoiHanThe()
	tuoiToiThieu := service.thamSoRepo.GetTuoiToiThieu()
	tuoiToiDa := service.thamSoRepo.GetTuoiToiDa()

	ngayHetHan := ngayLapThe.AddDate(0, 0, int(thoiHanThe))

	docGia := entity.NewDocGia(hoTen, loaiDocGia, ngaySinh, diaChi, email, ngayLapThe, &ngayHetHan)

	if isValid, err := docGia.IsValid(tuoiToiDa, tuoiToiThieu, thoiHanThe); isValid {
		return nil, err
	}

	docGiaRs, err := service.docGiaRepo.CreateDocGia(docGia)

	if err != nil {
		return nil, err
	}

	return docGiaRs, nil
}

func (service *DocGiaService) UpdateDocGia(maDocGia entity.ID, hoTen *string, loaiDocGia *entity.ID, ngaySinh *time.Time, diaChi *string, email *string) (*entity.DocGia, error) {
	docGia, err := service.docGiaRepo.GetDocGia(&maDocGia)
	if err != nil {
		return nil, err
	}

	if docGia == nil {
		return nil, businessError.NewBusinessError("doc gia not found")
	}

	// Update hoten

	if hoTen != nil {
		docGia.HoTen = *hoTen
	}

	// Update loai doc gia

	loaiDocGiaRs, err := service.loaiDocGiaRepo.GetLoaiDocGia(loaiDocGia)
	if err != nil {
		return nil, err
	}

	if loaiDocGiaRs == nil {
		return nil, businessError.NewBusinessError("loai doc gia not found")
	}

	docGia.MaLoaiDocGia = loaiDocGia

	// Update ngay sinh

	if ngaySinh != nil {
		docGia.NgaySinh = ngaySinh
	}

	// Update dia chi

	if diaChi != nil {
		docGia.DiaChi = *diaChi
	}

	// Update email

	if email != nil {
		docGia.Email = *email
	}
	// Get tham so

	thoiHanThe := service.thamSoRepo.GetThoiHanThe()
	tuoiToiThieu := service.thamSoRepo.GetTuoiToiThieu()
	tuoiToiDa := service.thamSoRepo.GetTuoiToiDa()

	// Validate

	if isValid, err := docGia.IsValid(tuoiToiDa, tuoiToiThieu, thoiHanThe); isValid {
		return nil, err
	}

	// Update doc gia

	updatedDocGia, err := service.docGiaRepo.UpdateDocGia(docGia)

	if err != nil {
		return nil, err
	}

	return updatedDocGia, nil
}

func (service *DocGiaService) RemoveDocGia(maDocGia *entity.ID) error {
	err := service.docGiaRepo.RemoveDocGia(maDocGia)
	return err
}
