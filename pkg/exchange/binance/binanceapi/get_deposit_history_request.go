package binanceapi

import (
	"time"

	"github.com/c9s/requestgen"

	"github.com/puper/bbgo/pkg/fixedpoint"
	"github.com/puper/bbgo/pkg/types"
)

type DepositStatus int

const (
	DepositStatusPending            DepositStatus = 0
	DepositStatusSuccess            DepositStatus = 1
	DepositStatusCredited           DepositStatus = 6
	DepositStatusWrong              DepositStatus = 7
	DepositStatusWaitingUserConfirm DepositStatus = 8
)

type DepositHistory struct {
	Amount        fixedpoint.Value           `json:"amount"`
	Coin          string                     `json:"coin"`
	Network       string                     `json:"network"`
	Status        DepositStatus              `json:"status"`
	Address       string                     `json:"address"`
	AddressTag    string                     `json:"addressTag"`
	TxId          string                     `json:"txId"`
	InsertTime    types.MillisecondTimestamp `json:"insertTime"`
	TransferType  int                        `json:"transferType"`
	UnlockConfirm int                        `json:"unlockConfirm"`

	// ConfirmTimes format = "current/required", for example: "7/16"
	ConfirmTimes string `json:"confirmTimes"`
	WalletType   int    `json:"walletType"`
}

//go:generate requestgen -method GET -url "/sapi/v1/capital/deposit/hisrec" -type GetDepositHistoryRequest -responseType []DepositHistory
type GetDepositHistoryRequest struct {
	client requestgen.AuthenticatedAPIClient

	coin *string `param:"coin"`

	startTime *time.Time `param:"startTime,milliseconds"`
	endTime   *time.Time `param:"endTime,milliseconds"`
}

func (c *RestClient) NewGetDepositHistoryRequest() *GetDepositHistoryRequest {
	return &GetDepositHistoryRequest{client: c}
}
