package model

import (
	"time"
)

// ShippingResponse menggantikan struct entities.ShippingResponse dari Python
type ShippingResponse struct {
	TrID              int       `json:"tr_id"`
	TrTglRequestKirim time.Time `json:"tr_tgl_request_kirim"`
	TrNamaMitra       string    `json:"tr_nama_mitra"`
	TrUp              string    `json:"tr_up"`
	TrNoTelp          int       `json:"tr_no_telp"`
	TrAlamat          string    `json:"tr_alamat"`
	TrKelurahan       string    `json:"tr_kelurahan"`
	TrKecamatan       string    `json:"tr_kecamatan"`
	TrKota            string    `json:"tr_kota"`
	TrProvinsi        string    `json:"tr_provinsi"`
	TrKodePos         int       `json:"tr_kode_pos"`
	TrJenisBarang     string    `json:"tr_jenis_barang"`
	TrMlID            int       `json:"tr_ml_id"`
	TrAsuransi        bool      `json:"tr_asuransi"`
	TrPackingKayu     bool      `json:"tr_packing_kayu"`
	TrMsID            int       `json:"tr_ms_id"`
	TrCreatedBy       string    `json:"tr_created_by"`
	TrNoReferensi     *string   `json:"tr_no_referensi,omitempty"`
	TrCatatan         *string   `json:"tr_catatan,omitempty"`
	TrMccCode         *string   `json:"tr_mcc_code,omitempty"`
	TrNoResi          *string   `json:"tr_no_resi,omitempty"`
	TrMeID            *int      `json:"tr_me_id,omitempty"`
	TrTglKirim        *time.Time `json:"tr_tgl_kirim,omitempty"`
}

// BaseResponse adalah wrapper untuk response API
type BaseResponse struct {
	ResponseCode string      `json:"response_code"`
	Message      string      `json:"message"`
	Data         interface{} `json:"data"`
}