package core

// DataService 数据服务集成
type DataService interface {
	// 钱包服务
	WalletService
	// 用户服务
	UserManageService
}
