package models


import (
  "github.com/jinzhu/gorm"
)



type Product struct {
  gorm.Model
  Name string `json:"name" form:"name" gorm:"size:255"`
  Brand string `json:"brand" form:"brand"`
  Category string `json:"category" form:"category"`
  Isactivated int `json:"isactivated" form:"isactivated" gorm:"default:1"`
  Isused int `json:"isused" form:"isused" gorm:"default:1"`
}


type MetaUser struct {
    Status bool `json:"status"`
    Code int `json:"code"`
    Message string `json:"message"`
}

type ResponseUser struct {
    Meta MetaUser `json:"meta"`
    Data []Product `json:"data"`
}