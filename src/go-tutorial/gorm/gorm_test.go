package gorm

import (
	"fmt"
	"testing"

	mygorm "go-tutorial/gorm"
)

func TestFindProduct1(t *testing.T) {
	product, err := mygorm.FindProduct("L1212")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(product)
}

func TestFindProducts(t *testing.T) {
	products, err := mygorm.FindProducts()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(products)
}

func TestUpdateProduct(t *testing.T) {
	var price uint = 2000
	product, err := mygorm.UpdateProductPrice("L1212", price)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(product)
}

func TestDeleteProduct(t *testing.T) {
	err := mygorm.DeleteProduct("L1212")
	if err != nil {
		t.Error(err)
	}
}
