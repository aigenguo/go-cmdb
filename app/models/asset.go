package models

import (
	"github.com/jinzhu/gorm"
)

// Asset 资产表
type Asset struct {
	gorm.Model
	Hostname  string `json:"hostname" gorm:"comment '主机名'"`
	IP        string `json:"ip" gorm:";comment '内网ip'"`
	Host      string `json:"host" gorm:"comment '宿主机'"`
	OS        string `json:"os" gorm:"comment '系统版本'"`
	OIP       string `json:"oip" gorm:"comment '外网ip'"`
	CPU       int8   `json:"cores" gorm:"comment 'cpu核数/个'"`
	Mem       int16  `json:"mem" gorm:"comment '内存/G'"`
	Disk      string `json:"disk" gorm:"comment '硬盘/G'"`
	Bandwidth int16  `json:"bandwidth" gorm:"comment '带宽/MB'"`
	UseOf     string `json:"use_of" gorm:"comment '用途'"`
	Principal string `json:"principal" gorm:"comment '负责人'"`
}

var Assets []Asset

// 添加资产
func (asset Asset) Insert() (id int64, err error) {
	// 添加数据
	result := DB.Create(&asset)
	id = int64(asset.ID)
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

// 列表
func (asset *Asset) Assets(pageNum, pageSize int, condition string) (assets []Asset, count int, err error) {
	var result = DB
	// 判断是否有查询条件
	if condition == "" {
		// 分页判断
		if pageNum > 0 {
			result = DB.Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&assets)
		} else {
			result = DB.Offset(pageNum).Limit(pageSize).Find(&assets)
		}
	} else {
		if pageNum > 0 && pageSize > 0 {
			// 支持模糊查询
			result = DB.Offset((pageNum-1)*pageSize).Limit(pageSize).Where("hostname like ?", "%"+condition+"%").Or(
				"ip like ?", "%"+condition+"%").Or("host like ?", "%"+condition+"%").Or(
				"os like ?", ""+condition+"%").Or("o_ip like ?", "%"+condition+"%").Or(
				"cpu like ?", "%"+condition+"%").Or("mem like ?", "%"+condition+"%").Or(
				"disk like ?", "%"+condition+"%").Or("bandwidth like ?", "%"+condition+"%").Or(
				"use_of like ?", "%"+condition+"%").Or("principal like ?", "%"+condition+"%").Find(&assets)
		} else {
			result = DB.Offset(pageNum).Limit(pageSize).Where("hostname like ?", "%"+condition+"%").Or(
				"ip like ?", "%"+condition+"%").Or("host like ?", "%"+condition+"%").Or(
				"os like ?", ""+condition+"%").Or("o_ip like ?", "%"+condition+"%").Or(
				"cpu like ?", "%"+condition+"%").Or("mem like ?", "%"+condition+"%").Or(
				"disk like ?", "%"+condition+"%").Or("bandwidth like ?", "%"+condition+"%").Or(
				"use_of like ?", "%"+condition+"%").Or("principal like ?", "%"+condition+"%").Find(&assets)
		}
	}
	DB.Table("tb_assets").Count(&count)
	err = result.Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return
	}
	return
}

// 删除资产
func (asset *Asset) Destroy(id int64) (result Asset, err error) {
	if err = DB.Select([]string{"id"}).First(&asset, id).Error; err != nil {
		return
	}
	if err = DB.Delete(&asset).Error; err != nil {
		return
	}
	result = *asset
	return
}

// 修改资产
func (asset *Asset) Update(id int64) (updateAsset Asset, err error) {
	if err = DB.Select([]string{"id"}).First(&updateAsset, id).Error; err != nil {
		return
	}
	if err = DB.Model(&updateAsset).Updates(&asset).Error; err != nil {
		return
	}
	return
}
