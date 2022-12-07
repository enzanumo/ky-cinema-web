package dao

import (
	"github.com/enzanumo/ky-theater-web/internal/model"
	"github.com/enzanumo/ky-theater-web/pkg/types"
	"gorm.io/gorm"
)

var AttachmentIncomeRate = 0.8

type walletServant = dataServant

func (d *walletServant) GetRechargeByID(id int64) (*model.WalletRecharge, error) {
	recharge := &model.WalletRecharge{
		Model: &model.Model{
			ID: id,
		},
	}

	return recharge.Get(d.db)
}
func (d *walletServant) CreateRecharge(userId, amount int64) (*model.WalletRecharge, error) {
	recharge := &model.WalletRecharge{
		UserID: userId,
		Amount: amount,
	}

	return recharge.Create(d.db)
}

func (d *walletServant) GetUserWalletBills(userID int64, offset, limit int) ([]*model.WalletStatement, error) {
	statement := &model.WalletStatement{
		UserID: userID,
	}

	return statement.List(d.db, &model.ConditionsT{
		"ORDER": "id DESC",
	}, offset, limit)
}

func (d *walletServant) GetUserWalletBillCount(userID int64) (int64, error) {
	statement := &model.WalletStatement{
		UserID: userID,
	}
	return statement.Count(d.db, &model.ConditionsT{})
}

func (d *walletServant) HandleRechargeSuccess(recharge *model.WalletRecharge, tradeNo string) error {
	user, _ := (&model.User{
		Model: &model.Model{
			ID: recharge.UserID,
		},
	}).Get(d.db)

	return d.db.Transaction(func(tx *gorm.DB) error {
		// 扣除金额
		if err := tx.Model(user).Update("balance", gorm.Expr("balance + ?", recharge.Amount)).Error; err != nil {
			// 返回任何错误都会回滚事务
			return err
		}

		// 新增账单
		if err := tx.Create(&model.WalletStatement{
			UserID:          user.ID,
			ChangeAmount:    recharge.Amount,
			BalanceSnapshot: user.Balance + recharge.Amount,
			Reason:          "用户充值",
		}).Error; err != nil {
			return err
		}

		// 标记为已付款
		if err := tx.Model(recharge).Updates(map[string]types.Any{
			"trade_no":     tradeNo,
			"trade_status": "TRADE_SUCCESS",
		}).Error; err != nil {
			return err
		}

		// 返回 nil 提交事务
		return nil
	})
}

func (d *walletServant) HandlePostAttachmentBought(post *model.Post, user *model.User) error {
	return d.db.Transaction(func(tx *gorm.DB) error {
		// 扣除金额
		if err := tx.Model(user).Update("balance", gorm.Expr("balance - ?", post.AttachmentPrice)).Error; err != nil {
			// 返回任何错误都会回滚事务
			return err
		}

		// 新增账单
		if err := tx.Create(&model.WalletStatement{
			PostID:          post.ID,
			UserID:          user.ID,
			ChangeAmount:    -post.AttachmentPrice,
			BalanceSnapshot: user.Balance - post.AttachmentPrice,
			Reason:          "支出",
		}).Error; err != nil {
			return err
		}

		// 新增购买记录
		if err := tx.Create(&model.PostAttachmentBill{
			PostID:     post.ID,
			UserID:     user.ID,
			PaidAmount: post.AttachmentPrice,
		}).Error; err != nil {
			return err
		}

		// 返回 nil 提交事务
		return nil
	})
}
