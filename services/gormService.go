package services

import (
	"github.com/ITBread/lotteryweb/datamodels"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Product struct {
	gorm.Model
	// Code string `json:"code"`
	// Date string `json:"Date"`
	// Red  string `json:"Red"`
	// Blue string `json:"Blue"`
	Ssq datamodels.Ssq
}

func LoadProducts(products []Product) {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&Product{})
	for v := range products {
		db.Create(v)
	}
	// 创建
	//db.Create(&Product{Code: "L1212", Price: 1000})

	// // 读取
	// var product Product
	// db.First(&product, 1)                   // 查询id为1的product
	// db.First(&product, "code = ?", "L1212") // 查询code为l1212的product

	// // 更新 - 更新product的price为2000
	// db.Model(&product).Update("Price", 2000)

	// // 删除 - 删除product
	// db.Delete(&product)
}
