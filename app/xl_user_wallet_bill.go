package app

// XlUserWalletBill xl_user_wallet_bill
type XlUserWalletBill struct {
	// gorm.Model
	ID                     uint    `gorm:"primary_key"`                           // id
	UserID                 uint    `json:"user_id"`                               // user_id
	Type                   uint    `json:"type"`                                  // 1 增加 or 0 减少
	Scenes                 uint    `json:"scenes"`                                // 0 订单-消费/退款, 1 红包收/发, 2 提现/充值,
	Amount                 float64 `json:"amount" sql:"type:decimal(10,2);"`      // amount 交易金额 【+/-】
	ServiceFee             float64 `json:"service_fee" sql:"type:decimal(10,2);"` // service_fee 交易手续费
	PartnerTradeNo         string  `json:"partner_trade_no"`                      // 交易号
	Body                   string  `json:"body"`                                  // 交易内容 | scenes = 0 : 购买的商品名【amount -(消费)/+(退款) | scenes = 1 红包 amount +(来着某某的红包)/-(发给某某的红包) | scenes = 2 amount +(充值)/-(提现)
	Goods                  string  `json:"goods"`                                 // 商品名 scenes = 0 商品详细名称
	PaymentMethod          string  `json:"payment_method"`                        // 支付方式 scenes = 0 订单支付方式/退款路径
	RedPacketFlow          string  `json:"red_packet_flow"`                       // 红包流 scenes = 1 amount +(发送方)/-(接收方)
	RechargeWithdrawMethod string  `json:"recharge_withdrawal_method"`            // 充值/提现方式 scenes = 2 amount +(充值方式)/-(提现方式) ['支付宝'，'微信']
	RechargeWithdrawFlow   string  `json:"recharge_withdrawal_flow"`              // 充值/提现账户 scenes = 2 amount +(微信-bank_type 付款银行，支付宝/fund_bill_list 支付渠道)/-(提现到openid)
	Status                 string  `json:"status"`                                // 交易状态 scenes = 0 [消费-待支付/已支付，退款-待退款/已退款] | scenes = 1 [发-待支付未发出/已支付并发送, 收-已收到] | scenes = 3 提现-发起提现/提现到账，充值-发起充值/充值成功
	ErrCode                string  `json:"err_code"`                              // 错误代码
	ErrCodeDesc            string  `json:"err_code_desc"`                         // 错误代码说明
	CreateTime             uint    // 创建时间 scenes = 0 发起消费/退款时间 | scenes = 1  | scenes = 2 发起充值/提现时间
	UpdateTime             uint    // 更新时间
}
