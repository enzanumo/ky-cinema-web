package service

import (
	"time"

	"github.com/enzanumo/ky-theater-web/internal/conf"
	"github.com/enzanumo/ky-theater-web/internal/model"
	"github.com/enzanumo/ky-theater-web/pkg/errcode"
	"github.com/gin-gonic/gin"
)

type RechargeReq struct {
	Amount int64 `json:"amount" form:"amount" binding:"required"`
}

func GetRechargeByID(id int64) (*model.WalletRecharge, error) {
	return DS.GetRechargeByID(id)
}

func CreateRecharge(userID, amount int64) (*model.WalletRecharge, error) {
	return DS.CreateRecharge(userID, amount)
}

func FinishRecharge(ctx *gin.Context, id int64, tradeNo string) error {
	if ok, _ := conf.Redis.SetNX(ctx, "PaoPaoRecharge:"+tradeNo, 1, time.Second*5).Result(); ok {
		recharge, err := DS.GetRechargeByID(id)
		if err != nil {
			return err
		}

		if recharge.TradeStatus != "TRADE_SUCCESS" {

			// 标记为已付款
			err := DS.HandleRechargeSuccess(recharge, tradeNo)
			defer conf.Redis.Del(ctx, "PaoPaoRecharge:"+tradeNo)

			if err != nil {
				return err
			}
		}

	}

	return nil
}

func BuyPostAttachment(post *model.Post, user *model.User) error {
	if user.Balance < post.AttachmentPrice {
		return errcode.InsuffientDownloadMoney
	}

	// 执行购买
	return DS.HandlePostAttachmentBought(post, user)
}
