package response

// 我的装备列表
type MyEquipmentList struct {
	//主武器
	MainWeapon MyEquipment `json:"mainWeapon"`
	//副武器
	SecondaryWeapon MyEquipment `json:"secondaryWeapon"`
	//眼镜
	Glasses MyEquipment `json:"glasses"`
	//服装
	Clothing MyEquipment `json:"clothing"`
	//鞋子
	Shoes MyEquipment `json:"shoes"`
}

// 我的装备
type MyEquipment struct {
	//装备名称
	Name string `json:"name"`
	//装备图标
	Icon string `json:"icon"`
}
