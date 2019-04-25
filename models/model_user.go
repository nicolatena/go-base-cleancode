package models


import (
  "github.com/jinzhu/gorm"
)



type User struct {
  gorm.Model
  Email string `json:"email" form:"email" gorm:"size:255" gorm:"unique"`
  Password string `json:"password" form:"password" gorm:"size:255"`
  Role string `json:"role" form:"role" gorm:"size:20"`
  Isactivated int `json:"isactivated" form:"isactivated" gorm:"default:1`
  Isused int `json:"isused" form:"isused gorm:"default:1"`
}


type MetaUser struct {
    Status bool `json:"status"`
    Code int `json:"code"`
    Message string `json:"message"`
}

type ResponseUser struct {
    Meta MetaUser `json:"meta"`
    Data []User `json:"data"`
}