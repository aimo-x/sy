package app

import (
	"time"
)

// XlServiceOrder 服务订单
type XlServiceOrder struct {
	ID               uint       `gorm:"primary_key" json:"id"`
	UserID           uint       `json:"user_id"`                         // 创建ID 的用户
	OrderNO          string     `json:"order_no" gorm:"unique;not null"` // 订单号码
	TotalPrice       int        `json:"total_price"`                     // 订单金额
	PayType          string     `json:"pay_type"`                        // 支付方式
	ServiceUserID    uint       `json:"service_user_id"`                 // 服务者用户ID
	ServiceGoodsID   uint       `json:"service_goods_id"`                // 购买商品ID
	ServicePrice     int        `json:"service_price"`                   // 服务单价
	ServicePriceUnit string     `json:"service_price_unit"`              // 服务单价的单位
	ServiceFee       int        `json:"service_fee"`                     // 非服务费【手续费】
	HeadImg          string     `json:"head_img"`                        // 订单缩略图
	Body             string     `json:"body"`                            // 订单名称
	Desc             string     `json:"desc"`                            // 订单描述
	Status           string     `json:"status"`                          // 待支付，已支付【待接单】，已确认【已接单】，已服务【服务者修改】，已完成【需求者确认】，退款中【需求者申请退款】, 已退款，交易关闭
	PayAt            time.Time  `json:"pay_time"`                        // 支付时间
	RefundAt         time.Time  `json:"refun_at"`                        // 退款时间
	CloseAt          time.Time  `json:"close_at"`                        // 交易关闭时间
	CompletedAt      time.Time  `json:"completed_at"`                    // 交易完成时间
	CreatedAt        time.Time  `json:"CreatedAt"`                       // 订单创建时间
	UpdatedAt        time.Time  `json:"UpdatedAt"`                       // 订单更新时间
	DeletedAt        *time.Time `sql:"index" json:"DeletedAt"`           // 软删除
}
