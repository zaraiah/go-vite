package dex

import (
	"errors"
)

var (
	InvalidOrderIdErr       = errors.New("invalid order id")
	InvalidOrderTypeErr     = errors.New("invalid order type")
	InvalidOrderPriceErr    = errors.New("invalid order price format")
	InvalidOrderQuantityErr = errors.New("invalid order quantity")
	OrderNotExistsErr       = errors.New("order not exists")
	OrderAmountTooSmallErr  = errors.New("order amount too small")

	TradeMarketExistsErr            = errors.New("trade market already exists")
	TradeMarketNotExistsErr         = errors.New("trade market not exists")
	TradeMarketStoppedErr           = errors.New("trade market stopped")
	ComposeOrderIdFailErr           = errors.New("compose order id fail")
	DeComposeOrderIdFailErr         = errors.New("decompose order id fail")
	TradeMarketInvalidQuoteTokenErr = errors.New("invalid quote token")
	TradeMarketInvalidTokenPairErr  = errors.New("invalid token pair")
	TradeMarketAllowMineErr         = errors.New("token pair already allow mine")
	TradeMarketNotAllowMineErr      = errors.New("token pair already not allow mine")

	CancelOrderOwnerInvalidErr  = errors.New("order to cancel not own to initiator")
	CancelOrderInvalidStatusErr = errors.New("order status is invalid to cancel")

	OnlyOwnerAllowErr = errors.New("only owner allow")

	ExceedFundAvailableErr     = errors.New("exceed fund available")
	ExceedFundLockedErr        = errors.New("try release locked amount exceed locked")
	InvalidPledgeAmountErr     = errors.New("invalid pledge amount")
	InvalidPledgeActionTypeErr = errors.New("invalid pledge action type")
	ExceedPledgeAvailableErr   = errors.New("exceed pledge available")
	PledgeForVipExistsErr      = errors.New("pledge for vip exists")
	PledgeForVipNotExistsErr   = errors.New("pledge for vip not exists")
	PledgeForVipNotExpireErr   = errors.New("pledge for vip not expire")

	InvalidSourceAddressErr           = errors.New("invalid source address")
	InvalidAmountForPledgeCallbackErr = errors.New("invalid amount for pledge callback")

	InvalidTokenErr                      = errors.New("invalid token")
	PendingDonateAmountSubExceedErr      = errors.New("pending donate amount sub exceed")
	PendingNewMarketInnerConflictErr     = errors.New("pending new market inner conflict")
	GetTokenInfoCallbackInnerConflictErr = errors.New("get token info callback inner conflict")
	InvalidTimestampFromTimerErr         = errors.New("invalid timestamp from timer")
	InvalidBrokerFeeRateErr              = errors.New("invalid broker fee rate")

	NoFeeSumFoundForValidPeriodErr = errors.New("no fee sum found for valid period")
	NotSetTimestampErr             = errors.New("not set timestamp")
	IterateVmDbFailedErr           = errors.New("iterate vm db failed")

	InvalidInviterFeeAmountErr = errors.New("invalid inviter fee amount")
	AlreadyIsInviterErr        = errors.New("already is inviter")

	InvalidInviteCodeErr  = errors.New("invalid invite code")
	NotBindInviterErr     = errors.New("not bind inviter")
	AlreadyBindInviterErr = errors.New("already bind inviter")
	NewInviteCodeFailErr = errors.New("new invite code fail")
	InternalErr = errors.New("internal error")
)
