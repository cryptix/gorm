package gorm_test

import (
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/shopspring/decimal"
)

type TestDecimalPtrStruct struct {
	gorm.Model
	Number *decimal.Decimal
}

func TestDecimalPtrDataType(t *testing.T) {
	var v TestDecimalPtrStruct
	scope := gorm.Scope{Value: &v}
	field, ok := scope.FieldByName("Number")
	if !ok {
		t.Fatal("no such field")
	}
	s := DB.Dialect().DataTypeOf(field.StructField)
	t.Log("data type for decimal:", s)
	if s != "string" {
		t.Errorf("data type for decimal is wrong")
	}

	n, _ := decimal.NewFromString("0.01")
	v.Number = &n
	if err := DB.Save(&v).Error; err != nil {
		t.Errorf("decimal save failed: %s", err)
	}

}

type TestDecimalStruct struct {
	gorm.Model
	Number decimal.Decimal
}

func TestDecimalDataType(t *testing.T) {
	var v TestDecimalStruct
	scope := gorm.Scope{Value: &v}
	field, ok := scope.FieldByName("Number")
	if !ok {
		t.Fatal("no such field")
	}
	s := DB.Dialect().DataTypeOf(field.StructField)
	t.Log("data type for decimal:", s)
	if s != "string" {
		t.Errorf("data type for decimal is wrong")
	}

	v.Number, _ = decimal.NewFromString("0.01")
	if err := DB.Save(&v).Error; err != nil {
		t.Errorf("decimal save failed: %s", err)
	}
}
