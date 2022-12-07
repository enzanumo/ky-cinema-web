package core

// DataService 数据服务集成
type DataService interface {
	MovieInfoService
	TicketService
	UserManageService
	WalletService
	RoomService
}
