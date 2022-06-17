package model

const (
	//用户封禁状态
	StatusNormal = 1
	StatusBan    = 0

	//地址
	RegionOther = 4 //4表示地址的其他

	//用户在线状态
	UserOnline    = 1 //在线
	UserNotOnline = 0 //离线

	//用户登录类型
	AdminLogin = 1 //管理员登录
	UserLogin  = 0 //普通用户登录

	//分类等级
	CategoryFirstLevel  = 1 //一级
	CategorySecondLevel = 2 //二级

	//销售状态
	IdleSaled   = 1
	IdleNotSale = 0

	//上架状态
	Published    = 1
	NotPublished = 0

	//求购物品状态
	PurchaseBuy    = 1
	PurchaseNotBuy = 0

	//订单状态
	OrderCompleted = 1 //订单已确认
	OrderOnGoing   = 2 //进行中，待确认
	OrderCancel    = 3 //订单已关闭（取消）

	//订单类型
	OrderTypeIdle     = 1
	OrderTypePurchase = 2

	//不存在的id
	NULLID = 0

	//物品类型
	TypeIdle     = 1
	TypePurchase = 2

	//自增类型
	IncLoginCount       = 1
	IncFinishOrderCount = 2

	//系统id
	SystemID = 1
)
