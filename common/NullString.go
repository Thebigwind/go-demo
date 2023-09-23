package common

import (
	"database/sql"
	"fmt"
	"time"
)

//sql.NullString

func GetTarget(src string) sql.NullString {

	valid := false
	if src != "" {
		valid = true
	}

	return sql.NullString{
		String: src,
		Valid:  valid,
	}
}

func GetResult(nullStr sql.NullString) string {
	str := ""
	if nullStr.Valid {
		str = nullStr.String
	}
	return str
}

func GetValid(src string) bool {
	if src != "" {
		return true
	}
	return false
}

type Certificate struct {
	Id            int64          `gorm:"column:id;primary_key" json:"id"`
	ParentSn      string         `gorm:"column:parent_sn" json:"parent_sn"`
	GmKeyId       sql.NullString `gorm:"column:gm_key_id" json:"gm_key_id"`   // id
	PqcKeyId      sql.NullString `gorm:"column:pqc_key_id" json:"pqc_key_id"` // pcaid
	Cert          string         `gorm:"column:cert" json:"cert"`
	Sn            string         `gorm:"column:sn" json:"sn"` // sn
	Name          string         `gorm:"column:name" json:"name"`
	Status        int64          `gorm:"column:status" json:"status"` // 12;3:
	ValidDate     int64          `gorm:"column:valid_date" json:"valid_date"`
	Type          int64          `gorm:"column:type" json:"type"` // 1-2-3-4
	IssuedBy      string         `gorm:"column:issued_by" json:"issued_by"`
	IssuedAt      time.Time      `gorm:"column:issued_at" json:"issued_at"`
	ExpiredAt     time.Time      `gorm:"column:expired_at" json:"expired_at"`
	NotBefore     time.Time      `gorm:"column:not_before" json:"not_before"`
	TenantId      int64          `gorm:"column:tenant_id" json:"tenant_id"` // id
	CertSubject   string         `gorm:"column:cert_subject" json:"cert_subject"`
	OrgName       string         `gorm:"column:org_name" json:"org_name"`
	SignAlgorithm string         `gorm:"column:sign_algorithm" json:"sign_algorithm"`
	KeyAlgorithm  string         `gorm:"column:key_algorithm" json:"key_algorithm"`
	KeyLen        int64          `gorm:"column:key_len" json:"key_len"`
	CreatedAt     time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt     time.Time      `gorm:"column:updated_at" json:"updated_at"`
}

func GetTarget2(src string) *string {
	if src == "" {
		return nil
	}
	return &src
}

func GetResult2(src interface{}) string {
	fmt.Printf("src:%v\n ", src)
	if src == nil {
		return ""
	}
	s, ok := src.(*string)
	if ok {
		return *s
	}
	return "xx"
}

func GetResult3(nullStr sql.NullString) string {
	str := ""
	if nullStr.Valid {
		str = nullStr.String
	}
	return str
}

func test(gmkeyId, pqcKeyId string) {
	//insert
	var certificate = Certificate{
		//Id:            0,

		GmKeyId:  sql.NullString{String: gmkeyId, Valid: GetValid(gmkeyId)},   // common.GetTarget(gmKeyId),  //&gmKeyId
		PqcKeyId: sql.NullString{String: pqcKeyId, Valid: GetValid(pqcKeyId)}, //common.GetTarget(pqcKeyId),                                       //签名证书如果携带keyId,只能是pqc_key_id &pqcKeyId
	}
	fmt.Println(certificate)
}
