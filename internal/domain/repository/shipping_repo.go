package repository

import (
	"errors"
	"fmt"
	"log"

	"github.com/GratiaManullang03/ekspedisi-go/internal/domain/model"
	"gorm.io/gorm"
)

type ShippingRepository struct {
	db *gorm.DB
}

func NewShippingRepository(db *gorm.DB) *ShippingRepository {
	return &ShippingRepository{db: db}
}

func (r *ShippingRepository) SelectShipping(role string, nik string, costCenter string) ([]model.ShippingResponse, error) {
	var shippings []model.ShippingResponse
	var err error

	// Membuat base query
	query := `
		SELECT
			tr_id,
			tr_tgl_request_kirim,
			tr_nama_mitra,
			tr_up,
			tr_no_telp,
			tr_alamat,
			tr_kelurahan,
			tr_kecamatan,
			tr_kota,
			tr_provinsi,
			tr_kode_pos,
			tr_jenis_barang,
			tr_ml_id,
			tr_asuransi,
			tr_packing_kayu,
			tr_ms_id,
			tr_created_by,
			tr_no_referensi,
			tr_catatan,
			tr_mcc_code,
			tr_no_resi,
			tr_me_id,
			tr_tgl_kirim
		FROM ekspedisi.trx_request
		WHERE 1=1
	`

	// Menambahkan kondisi sesuai role
	var args []interface{}
	
	if role == "USER" {
		query += " AND tr_created_by = ?"
		args = append(args, nik)
	} else if role == "MANAGER" {
		query += " AND tr_ms_id IN (1, 2, 5) AND tr_mcc_code = ?"
		args = append(args, costCenter)
	} else if role == "SUPER_ADMIN" {
		query += " AND tr_ms_id IN (2, 3, 4)"
	}

	log.Printf("Executing query: %s | Params: %v", query, args)
	
	// Eksekusi query dan scan hasilnya
	rows, err := r.db.Raw(query, args...).Rows()
	if err != nil {
		log.Printf("Database query failed: %v", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var shipping model.ShippingResponse
		err = r.db.ScanRows(rows, &shipping)
		if err != nil {
			log.Printf("Error scanning rows: %v", err)
			return nil, err
		}
		shippings = append(shippings, shipping)
	}

	log.Printf("Query Result: Retrieved %d records", len(shippings))
	
	if len(shippings) == 0 {
		return []model.ShippingResponse{}, nil
	}
	
	return shippings, nil
}

func (r *ShippingRepository) SelectByID(trID int) (model.ShippingResponse, error) {
	var shipping model.ShippingResponse

	query := `
		SELECT
			tr_id,
			tr_tgl_request_kirim,
			tr_nama_mitra,
			tr_up,
			tr_no_telp,
			tr_alamat,
			tr_kelurahan,
			tr_kecamatan,
			tr_kota,
			tr_provinsi,
			tr_kode_pos,
			tr_jenis_barang,
			tr_ml_id,
			tr_asuransi,
			tr_packing_kayu,
			tr_ms_id,
			tr_created_by,
			tr_no_referensi,
			tr_catatan,
			tr_mcc_code,
			tr_no_resi,
			tr_me_id,
			tr_tgl_kirim
		FROM ekspedisi.trx_request
		WHERE tr_id = ?
	`

	err := r.db.Raw(query, trID).Scan(&shipping).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.ShippingResponse{}, fmt.Errorf("shipping with ID %d not found", trID)
		}
		return model.ShippingResponse{}, err
	}

	return shipping, nil
}