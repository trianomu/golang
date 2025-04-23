
package domain

import "time"

type Customer struct {
    ID          string    `gorm:"primaryKey" json:"id"`
    NIK         string    `json:"nik"`
    FullName    string    `json:"full_name"`
    LegalName   string    `json:"legal_name"`
    BirthPlace  string    `json:"birth_place"`
    BirthDate   time.Time `json:"birth_date"`
    Salary      int64     `json:"salary"`
    PhotoKTP    string    `json:"photo_ktp"`
    PhotoSelfie string    `json:"photo_selfie"`
    CreatedAt   time.Time `json:"created_at"`
}

type Limit struct {
    ID         string `gorm:"primaryKey"`
    CustomerID string
    Tenor      int
    Amount     int64
}

type Transaction struct {
    ID             string    `gorm:"primaryKey"`
    CustomerID     string
    ContractNumber string
    OTR            int64
    AdminFee       int64
    Installment    int64
    Interest       int64
    AssetName      string
    CreatedAt      time.Time
}
