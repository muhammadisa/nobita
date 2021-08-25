package mdt

import (
	"crypto/sha512"
	"errors"
	"fmt"
	"github.com/veritrans/go-midtrans"
	"time"
)

type Env string

const (
	Sandbox    Env = `SANDBOX`
	Production Env = `PRODUCTION`
	TimeUnit       = `MINUTE`
	TimeLayout     = `2006-01-02 15:04:05 +0700`
)

func NewMidtransClient(sk string, ck string, env Env) *MtCore {
	mclient := midtrans.NewClient()
	mclient.ServerKey = sk
	mclient.ClientKey = ck
	switch env {
	case Sandbox:
		mclient.APIEnvType = midtrans.Sandbox
	case Production:
		mclient.APIEnvType = midtrans.Production
	default:
		panic(errors.New("invalid environment type"))
	}
	return &MtCore{
		SK: sk,
		CK: ck,
		Core: midtrans.CoreGateway{
			Client: mclient,
		},
	}
}

type Trx struct {
	Status     string
	Gross      string
	ExpireAt   int64
	TrxCode    string
	TrxType    int
	Message    string
	Bank       string
	VaNumber   string
	BillerCode string
	BillKey    string
}

type MtChargingResponse struct {
	MtResponse *midtrans.Response
}

func (mcr MtChargingResponse) Trx(expireTime time.Time) (*Trx, error) {
	res := mcr.MtResponse

	var va, bankNames string
	var err error
	if res.PermataVaNumber != "" {
		va = res.PermataVaNumber
		bankNames = "permata"
	} else {
		if res.BillKey != "" && res.BillerCode != "" {
			va = ""
			bankNames = "mandiri"
		} else {
			if len(res.VANumbers) > 0 {
				va = res.VANumbers[0].VANumber
				bankNames = res.VANumbers[0].Bank
			} else {
				err = errors.New("error while doing payment")
			}
		}
	}
	if err != nil {
		return nil, err
	}
	return &Trx{
		ExpireAt:   expireTime.Unix(),
		TrxType:    1,
		TrxCode:    res.OrderID,
		Status:     res.StatusCode,
		Gross:      res.GrossAmount,
		Message:    res.StatusMessage,
		Bank:       bankNames,
		VaNumber:   va,
		BillerCode: res.BillerCode,
		BillKey:    res.BillKey,
	}, nil
}

type MtCore struct {
	SK, CK   string
	Core     midtrans.CoreGateway
	TimeZone string
}

func (mc MtCore) RetrieveStatus(orderID string) (midtrans.Response, error) {
	return mc.Core.Status(orderID)
}

func (mc MtCore) VerifySignature(signature string, orderID, statusCode, grossAmt string) bool {
	formulas := orderID + statusCode + grossAmt + mc.SK
	newSha512 := sha512.New()
	newSha512.Write([]byte(formulas))
	formatted := fmt.Sprintf("%x", newSha512.Sum(nil))
	return signature == formatted
}

func (mc MtCore) RequestCharge(chargeReq *midtrans.ChargeReq) (*midtrans.Response, error) {
	response, err := mc.Core.Charge(chargeReq)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

type ItemDetail struct {
	ID           string
	Name         string
	Price        int64
	Qty          int32
	Brand        string
	Category     string
	MerchantName string
}

type DetailTX struct {
	Phone, Email, FName, LName, TrxID string
	TrxTime                           time.Time
	Items                             []ItemDetail
}

func (dt DetailTX) grossAmtCount() int64 {
	var grossAmt int64
	for _, item := range dt.Items {
		grossAmt += item.Price
	}
	return grossAmt
}

func (dt DetailTX) toMidtransItemDetail() *[]midtrans.ItemDetail {
	midtransItem := make([]midtrans.ItemDetail, len(dt.Items))
	for i, item := range dt.Items {
		midtransItem[i] = midtrans.ItemDetail{
			ID:           item.ID,
			Name:         item.Name,
			Price:        item.Price,
			Qty:          item.Qty,
			Brand:        item.Brand,
			Category:     item.Category,
			MerchantName: item.MerchantName,
		}
	}
	return &midtransItem
}

func (dt DetailTX) BCAVirtualAccountCharger() *midtrans.ChargeReq {
	return &midtrans.ChargeReq{
		PaymentType: midtrans.SourceBankTransfer,
		BankTransfer: &midtrans.BankTransferDetail{
			Bank: midtrans.BankBca,
		},
		CustomExpiry: &midtrans.CustomExpiry{
			OrderTime:      dt.TrxTime.Format(TimeLayout),
			ExpiryDuration: 60,
			Unit:           TimeUnit,
		},
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  dt.TrxID,
			GrossAmt: dt.grossAmtCount(),
		},
		CustomerDetail: &midtrans.CustDetail{
			Email: dt.Email,
			Phone: dt.Phone,
			FName: dt.FName,
			LName: dt.LName,
		},
		Items: dt.toMidtransItemDetail(),
	}
}

func (dt DetailTX) BNIVirtualAccountCharger() *midtrans.ChargeReq {
	return &midtrans.ChargeReq{
		PaymentType: midtrans.SourceBankTransfer,
		BankTransfer: &midtrans.BankTransferDetail{
			Bank: midtrans.BankBni,
		},
		CustomExpiry: &midtrans.CustomExpiry{
			OrderTime:      dt.TrxTime.Format(TimeLayout),
			ExpiryDuration: 60,
			Unit:           TimeUnit,
		},
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  dt.TrxID,
			GrossAmt: dt.grossAmtCount(),
		},
		CustomerDetail: &midtrans.CustDetail{
			Email: dt.Email,
			Phone: dt.Phone,
			FName: dt.FName,
			LName: dt.LName,
		},
		Items: dt.toMidtransItemDetail(),
	}
}

func (dt DetailTX) BRIVirtualAccountCharger() *midtrans.ChargeReq {
	return &midtrans.ChargeReq{
		PaymentType: midtrans.SourceBankTransfer,
		BankTransfer: &midtrans.BankTransferDetail{
			Bank: midtrans.BankBri,
		},
		CustomExpiry: &midtrans.CustomExpiry{
			OrderTime:      dt.TrxTime.Format(TimeLayout),
			ExpiryDuration: 60,
			Unit:           TimeUnit,
		},
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  dt.TrxID,
			GrossAmt: dt.grossAmtCount(),
		},
		CustomerDetail: &midtrans.CustDetail{
			Email: dt.Email,
			Phone: dt.Phone,
			FName: dt.FName,
			LName: dt.LName,
		},
		Items: dt.toMidtransItemDetail(),
	}
}

func (dt DetailTX) PermataVirtualAccountCharger() *midtrans.ChargeReq {
	return &midtrans.ChargeReq{
		PaymentType: midtrans.SourceBankTransfer,
		BankTransfer: &midtrans.BankTransferDetail{
			Bank: midtrans.BankPermata,
		},
		CustomExpiry: &midtrans.CustomExpiry{
			OrderTime:      dt.TrxTime.Format(TimeLayout),
			ExpiryDuration: 60,
			Unit:           TimeUnit,
		},
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  dt.TrxID,
			GrossAmt: dt.grossAmtCount(),
		},
		CustomerDetail: &midtrans.CustDetail{
			Email: dt.Email,
			Phone: dt.Phone,
			FName: dt.FName,
			LName: dt.LName,
		},
		Items: dt.toMidtransItemDetail(),
	}
}

func (dt DetailTX) MandiriBillVirtualAccountCharger() *midtrans.ChargeReq {
	return &midtrans.ChargeReq{
		PaymentType: midtrans.SourceEchannel,
		MandiriBillBankTransferDetail: &midtrans.MandiriBillBankTransferDetail{
			BillInfo1: "please complete payment",
			BillInfo2: "dept",
		},
		CustomExpiry: &midtrans.CustomExpiry{
			OrderTime:      dt.TrxTime.Format(TimeLayout),
			ExpiryDuration: 60,
			Unit:           TimeUnit,
		},
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  dt.TrxID,
			GrossAmt: dt.grossAmtCount(),
		},
		CustomerDetail: &midtrans.CustDetail{
			Email: dt.Email,
			Phone: dt.Phone,
			FName: dt.FName,
			LName: dt.LName,
		},
		Items: dt.toMidtransItemDetail(),
	}
}
