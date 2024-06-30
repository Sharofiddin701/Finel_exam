package models

type ScanBarCode struct {
	BarCode string `json:"barcode"`
	SaleId  string `json:"sale_id"`
}
