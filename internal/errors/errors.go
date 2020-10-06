package errors

import (
	"math/big"

	"github.com/richard-xtek/go-grpc-micro-kit/utils/math"

	"google.golang.org/grpc"
)

// Define errors
var (
	ErrExceededYourWithdrawalLimit = grpc.Errorf(20001, "You have exceeded your withdrawal limit")
	ErrAuthenticationFail          = grpc.Errorf(20002, "Please login to access")
	ErrAmountInvaild               = grpc.Errorf(20003, "You entered an invalid number, please enter again")
	ErrWithdrawDisable             = grpc.Errorf(20004, "Withdrawal is temporarily suspended, Please try again later")
	ErrInvaildAddress              = grpc.Errorf(20005, "Your wallet address is not valid, please enter it again")
	ErrMinWithdrawMustIs           = func(coin string, minAmount *big.Int) error {
		return grpc.Errorf(20006, "The withdrawal amount must be greater %s %s", math.ToDecimalString(minAmount), coin)
	}
	ErrInsufficientBalance     = grpc.Errorf(20006, "Insufficient balance")
	ErrGoogleVerifyCodeInvaild = grpc.Errorf(20007, "That 2FA verification code was invalid. Please try again.")
	ErrEmailVerifyCodeInvaild  = grpc.Errorf(20007, "That email code was invalid. Please try again.")
	ErrAdminConfirmNotMatch    = func(coin string, maxAmount *big.Int) error {
		return grpc.Errorf(20008, "The withdrawal amount less than %s %s are automatic", math.ToDecimalString(maxAmount), coin)
	}
	ErrInvaildTransaction = grpc.Errorf(20009, "Invaild transaction")
	ErrWithdrawHitLimit   = func(maxAmount *big.Int, availableCoinAmount *big.Int, coin string) error {
		if math.IsStrictlySmallerThan(availableCoinAmount, big.NewInt(0)) {
			return grpc.Errorf(20011, "You only withdraw %s USDT within 24 hours", math.ToDecimalString(maxAmount))
		}
		return grpc.Errorf(20010, "You only withdraw %s USDT within 24 hours and now you can only withdraw up to %s %s", math.ToDecimalString(maxAmount), math.ToDecimalString(availableCoinAmount), coin)
	}
	ErrNotFoundPriceByCoin       = grpc.Errorf(20012, "Not found price by coin")
	ErrInvaildTransferWalletType = grpc.Errorf(20013, "Invaild transfer wallet type")
	ErrCoinInvaild               = grpc.Errorf(20014, "Coin invaild")
	ErrMinTransferMustIs         = func(coin string, minAmount *big.Int) error {
		return grpc.Errorf(20005, "The transfer amount must be greater %s %s", math.ToDecimalString(minAmount), coin)
	}
)
