package docgia

import (
	coreerror "daijoubuteam.xyz/se214-library-management/core/error"
	"time"

	"daijoubuteam.xyz/se214-library-management/core/entity"
	"daijoubuteam.xyz/se214-library-management/core/repository"
	loaidocgia "daijoubuteam.xyz/se214-library-management/usecase/loai_doc_gia"
)

type DocGiaService struct {
	loaiDocGiaUsecase loaidocgia.LoaiDocGiaUsecase
	docGiaRepo        repository.DocGiaRepository
	thamSoRepo        repository.ThamSoRepository
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
		return nil, coreerror.NewNotFoundError("doc gia not found", nil)
	}

	return docGia, err
}

func (service *DocGiaService) CreateDocGia(hoTen string, maLoaiDocGia *entity.ID, ngaySinh *time.Time, diaChi string, email string, ngayLapThe *time.Time) (*entity.DocGia, error) {

	loaiDocGia, err := service.loaiDocGiaUsecase.GetLoaiDocGia(maLoaiDocGia)
	if err != nil {
		return nil, err
	}

	// Get tham so

	thoiHanThe, err := service.thamSoRepo.GetThoiHanThe()
	if err != nil {
		return nil, err
	}
	tuoiToiThieu, err := service.thamSoRepo.GetTuoiToiThieu()
	if err != nil {
		return nil, err
	}
	tuoiToiDa, err := service.thamSoRepo.GetTuoiToiDa()
	if err != nil {
		return nil, err
	}

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

func (service *DocGiaService) UpdateDocGia(maDocGia *entity.ID, hoTen string, maLoaiDocGia *entity.ID, ngaySinh *time.Time, diaChi string, email string) (*entity.DocGia, error) {
	docGia, err := service.docGiaRepo.GetDocGia(maDocGia)
	if err != nil {
		return nil, err
	}

	if docGia == nil {
		return nil, coreerror.NewNotFoundError("doc gia not found", nil)
	}

	// Update hoten

	docGia.HoTen = hoTen

	// Update loai doc gia

	loaiDocGia, err := service.loaiDocGiaUsecase.GetLoaiDocGia(maLoaiDocGia)
	if err != nil {
		return nil, err
	}

	docGia.LoaiDocGia = loaiDocGia

	// Update ngay sinh

	if ngaySinh != nil {
		docGia.NgaySinh = ngaySinh
	}

	// Update dia chi

	docGia.DiaChi = diaChi

	// Update email

	docGia.Email = email
	// Get tham so

	thoiHanThe, err := service.thamSoRepo.GetThoiHanThe()
	if err != nil {
		return nil, err
	}
	tuoiToiThieu, err := service.thamSoRepo.GetTuoiToiThieu()
	if err != nil {
		return nil, err
	}
	tuoiToiDa, err := service.thamSoRepo.GetTuoiToiDa()
	if err != nil {
		return nil, err
	}

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
