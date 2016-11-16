package gorm

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Diary struct {
	ID    int
	Title string
	Ymd   time.Time `sql:"not null;type:date"`
}

type Product struct {
	ID    int
	Code  string
	Price uint
}

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open("mysql", "root:@/go_tutorial?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	db.LogMode(true)
	db.AutoMigrate(&Diary{})
	db.AutoMigrate(&Product{})
	db.Delete(&Diary{})
	db.Delete(&Product{})
	ymd := time.Date(2014, time.December, 31, 12, 13, 24, 0, time.UTC)
	db.Create(&Diary{Title: "titile", Ymd: ymd})
	db.Create(&Product{Code: "L1212", Price: 1000})
}

func FindDiary(day time.Time) (*Diary, error) {
	var diary Diary
	err := db.Where("ymd = ?", day.Format("2006-01-02")).First(&diary).Error
	if err != nil {
		return nil, err
	}
	return &diary, nil
}

func FindProduct(code string) (*Product, error) {
	var product Product
	err := db.Where("code = ?", code).First(&product).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func FindProducts() ([]Product, error) {
	var products []Product
	err := db.Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

func UpdateProductPrice(code string, price uint) (*Product, error) {
	tx := db.Begin()
	product, err := FindProduct(code)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	product.Price = price
	db.Save(product)
	tx.Commit()
	return product, nil
}

func DeleteProduct(code string) error {
	tx := db.Begin()
	product, err := FindProduct(code)
	if err != nil {
		tx.Rollback()
		return err
	}
	db.Delete(product)
	tx.Commit()
	return nil
}
