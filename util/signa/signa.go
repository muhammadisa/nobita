package signa

import (
	"crypto/sha512"
	"fmt"
)

type CharityData struct {
	SC string // Super Secret
	WI string // Wallet ID
	PD string // Pool Date
	BA int64  // Before Amount
	AA int64  // After Amount
}

type WalletData struct {
	SC string // Super Secret
	AI int64  // Account ID
	LA int64  // Last Amount
	AA int64  // After Amount
}

func NewCharitySigna(data CharityData) string {
	return genCharitySignature(data.SC, data.WI, data.PD, data.BA, data.AA)
}

func NewCharitySignaVerify(data CharityData, signature string) bool {
	current := genCharitySignature(data.SC, data.WI, data.PD, data.BA, data.AA)
	return current == signature
}

func NewWalletSigna(data WalletData) string {
	return genWalletSignature(data.SC, data.AI, data.LA, data.AA)
}

func NewWalletSignaVerify(data WalletData, signature string) bool {
	current := genWalletSignature(data.SC, data.AI, data.LA, data.AA)
	return current == signature
}

func genCharitySignature(sc string, wi string, pd string, bp int64, ap int64) string {
	newSha512 := sha512.New()
	newSha512.Write([]byte(fmt.Sprintf("%s%s%s%d%d", sc, wi, pd, bp, ap)))
	return fmt.Sprintf("%x", newSha512.Sum(nil))
}

func genWalletSignature(sc string, ai int64, la int64, aa int64) string {
	newSha512 := sha512.New()
	newSha512.Write([]byte(fmt.Sprintf("%s%d%d%d", sc, ai, la, aa)))
	return fmt.Sprintf("%x", newSha512.Sum(nil))
}
