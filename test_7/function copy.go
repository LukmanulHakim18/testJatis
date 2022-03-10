package test8

import (
	"errors"
	"time"
)

var TotalPrice = 1000000

type Voucher struct {
	VoucherId      string
	VoucherNominal int
	Expired        time.Time
	IsUsed         bool
}

var DBvochs = make(map[string]Voucher)

func init() {
}

func ClimeVoucher(invoiceNumber string) (*Voucher, error) {
	if DBvochs[invoiceNumber].VoucherId != "" {
		return nil, errors.New("invoice hase redeem to voucher")
	}
	if TotalPrice < 1 {
		return nil, errors.New("out of stock")
	}
	generateVoucher := Voucher{VoucherId: invoiceNumber, VoucherNominal: 10000, Expired: time.Now().AddDate(0, 2, 0), IsUsed: false}
	DBvochs[invoiceNumber] = generateVoucher
	TotalPrice -= 10000
	return &generateVoucher, nil
}

func RedeemVoucher(VoucherId string) string {
	voucherFound := DBvochs[VoucherId]
	if voucherFound.Expired.After(time.Now()) {
		return "foucher expired"
	}
	if voucherFound.IsUsed {
		return "voucher have ben reedem"
	}
	voucherFound.IsUsed = true
	DBvochs[VoucherId] = voucherFound
	return "success"
}
