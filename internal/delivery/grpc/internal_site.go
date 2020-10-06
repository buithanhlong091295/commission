package grpc

// import (
// 	"context"
// 	balance "xtek/exchange/commission/internal/domain/balance"
// 	coin "xtek/exchange/commission/internal/domain/coin"
// 	transaction "xtek/exchange/commission/internal/domain/transaction"
// 	useraddress "xtek/exchange/commission/internal/domain/useraddress"
// 	"xtek/exchange/commission/internal/models"
// 	model "xtek/exchange/commission/internal/models"
// 	pb "xtek/exchange/commission/pb/wallet"
// 	pbDTO "xtek/exchange/commission/pb/wallet/dto"

// 	"github.com/golang/protobuf/ptypes/empty"
// 	"github.com/richard-xtek/go-grpc-micro-kit/utils/math"

// 	"google.golang.org/grpc"
// 	"google.golang.org/grpc/codes"
// )

// // NewInternalSiteDelivery ...
// func NewInternalSiteDelivery(balanceDomain *balance.BalanceDomain,
// 	coinDomain *coin.CoinDomain,
// 	userAddressDomain *useraddress.UserAddressesDomain,
// 	transactionDomain *transaction.TransactionDomain,
// 	adminTransactionDomain *transaction.AdminTransactionDomain) (pb.InternalSiteServiceServer, error) {
// 	return &internalSiteDelivery{balanceDomain, coinDomain, userAddressDomain, transactionDomain, adminTransactionDomain}, nil
// }

// type internalSiteDelivery struct {
// 	balanceDomain          *balance.BalanceDomain
// 	coinDomain             *coin.CoinDomain
// 	userAddressDomain      *useraddress.UserAddressesDomain
// 	transactionDomain      *transaction.TransactionDomain
// 	adminTransactionDomain *transaction.AdminTransactionDomain
// }

// func (s *internalSiteDelivery) CreateBalanceDefault(ctx context.Context, req *pbDTO.CreateBalanceDefaultRequest) (*pbDTO.CreateBalanceDefaultResponse, error) {
// 	if err := req.Validate(); err != nil {
// 		return nil, err
// 	}
// 	err := s.balanceDomain.CreateBalanceDefault(ctx, req.GetUserID())
// 	if err != nil {
// 		return nil, err
// 	}
// 	var rep = new(pbDTO.CreateBalanceDefaultResponse)
// 	rep.Result = true
// 	return rep, nil
// }

// func (s *internalSiteDelivery) CreateCoin(ctx context.Context, req *pbDTO.CreateCoinRequest) (*pbDTO.CreateCoinResponse, error) {
// 	if err := req.Validate(); err != nil {
// 		return nil, err
// 	}
// 	coin := &model.Coin{
// 		Coin:    req.GetCoin(),
// 		Name:    req.GetName(),
// 		LogoURL: req.GetLogoUrl(),
// 	}
// 	coin, err := s.coinDomain.CreateCoin(ctx, coin)
// 	if err != nil {
// 		return nil, err
// 	}
// 	var rep = new(pbDTO.CreateCoinResponse)
// 	rep.Result = coin.ID
// 	return rep, nil
// }

// func (s *internalSiteDelivery) HandleDepositWallet(ctx context.Context, req *pbDTO.HandleDepositnWithWalletRequest) (*pbDTO.HandleDepositnWithWalletResponse, error) {
// 	if err := req.Validate(); err != nil {
// 		return nil, err
// 	}
// 	err := s.transactionDomain.HandleDepositWallet(
// 		ctx,
// 		req,
// 	)
// 	if err != nil {
// 		return nil, err
// 	}
// 	var resp = &pbDTO.HandleDepositnWithWalletResponse{
// 		Result: true,
// 	}
// 	return resp, nil
// }

// func (s *internalSiteDelivery) HandleWithdrawalWallet(ctx context.Context, req *pbDTO.HandleDepositnWithdrawalWalletRequest) (*pbDTO.HandleDepositnWithdrawalWalletResponse, error) {
// 	if err := req.Validate(); err != nil {
// 		return nil, err
// 	}
// 	err := s.transactionDomain.HandleWithdrawalWallet(
// 		ctx,
// 		req,
// 	)
// 	if err != nil {
// 		return nil, err
// 	}
// 	var resp = &pbDTO.HandleDepositnWithdrawalWalletResponse{
// 		Result: true,
// 	}
// 	return resp, nil
// }

// func (s *internalSiteDelivery) ValidateAvailableBalance(ctx context.Context, req *pbDTO.ValidateAvailableBalanceRequest) (*pbDTO.ValidateAvailableBalanceResponse, error) {
// 	if err := req.Validate(); err != nil {
// 		return nil, err
// 	}
// 	if math.IsZero(math.ToBigInt(req.GetAmount())) {
// 		return nil, grpc.Errorf(codes.InvalidArgument, "Amount must greater than zero")
// 	}

// 	isAvailable, err := s.balanceDomain.IsAvailableBalance(ctx, req)
// 	if err != nil {
// 		return nil, err
// 	}

// 	rs := &pbDTO.ValidateAvailableBalanceResponse{
// 		IsVaild: isAvailable,
// 	}
// 	return rs, nil
// }

// func (s *internalSiteDelivery) ValidateBalance(ctx context.Context, req *pbDTO.ValidateBalanceRequest) (*pbDTO.ValidateBalanceResponse, error) {
// 	if err := req.Validate(); err != nil {
// 		return nil, err
// 	}
// 	if math.IsZero(math.ToBigInt(req.GetAmount())) {
// 		return nil, grpc.Errorf(codes.InvalidArgument, "Amount must greater than zero")
// 	}
// 	isVaild, err := s.balanceDomain.ValidateBalance(ctx, req)
// 	if err != nil {
// 		return nil, err
// 	}

// 	rs := &pbDTO.ValidateBalanceResponse{
// 		IsVaild: isVaild,
// 	}
// 	return rs, nil
// }

// func (s *internalSiteDelivery) GetUser(ctx context.Context, req *pbDTO.ValidateBalanceRequest) (*pbDTO.ValidateBalanceResponse, error) {
// 	if err := req.Validate(); err != nil {
// 		return nil, err
// 	}
// 	if math.IsZero(math.ToBigInt(req.GetAmount())) {
// 		return nil, grpc.Errorf(codes.InvalidArgument, "Amount must greater than zero")
// 	}
// 	isVaild, err := s.balanceDomain.ValidateBalance(ctx, req)
// 	if err != nil {
// 		return nil, err
// 	}

// 	rs := &pbDTO.ValidateBalanceResponse{
// 		IsVaild: isVaild,
// 	}
// 	return rs, nil
// }

// // GetBalanceByUserIDnCoin ...
// func (s *internalSiteDelivery) GetBalanceByUserIDnCoin(ctx context.Context, req *pbDTO.GetBalanceByUserIDnCoinRequest) (*pbDTO.GetBalanceByUserIDnCoinResponse, error) {
// 	if err := req.Validate(); err != nil {
// 		return nil, err
// 	}

// 	balance, err := s.balanceDomain.GetBalanceByUserIDnCoin(ctx, req.GetUserID(), req.GetCoin())
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &pbDTO.GetBalanceByUserIDnCoinResponse{Amount: balance.Amount.String()}, nil
// }

// // GetBalancesByUserIDNType ...
// func (s *internalSiteDelivery) GetBalancesByUserIDNType(ctx context.Context, req *pbDTO.GetBalancesByUserIDNTypeRequest) (*pbDTO.GetBalancesByUserIDNTypeResponse, error) {
// 	if err := req.Validate(); err != nil {
// 		return nil, err
// 	}

// 	balances, err := s.balanceDomain.GetBalancesByUserIDNType(ctx, req.GetUserID(), models.BalanceType(req.GetType()))
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &pbDTO.GetBalancesByUserIDNTypeResponse{Listings: balances}, nil
// }

// // GenerateUserAddress ...
// func (s *internalSiteDelivery) GenerateUserAddress(ctx context.Context, req *pbDTO.GenerateUserAddressRequest) (*empty.Empty, error) {
// 	if err := s.userAddressDomain.GenerateUserAddress(ctx, req); err != nil {
// 		return nil, err
// 	}
// 	return &empty.Empty{}, nil
// }

// func (s *internalSiteDelivery) GetCoins(ctx context.Context, req *empty.Empty) (*pbDTO.GetCoinsResponse, error) {
// 	coins, err := s.coinDomain.GetCoins(ctx)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &pbDTO.GetCoinsResponse{Listings: coins}, nil
// }

// func (s *internalSiteDelivery) GetBalanceByUserIDnCoinnType(ctx context.Context, req *pbDTO.GetBalanceByUserIDnCoinnTypeRequest) (*pbDTO.GetBalanceByUserIDnCoinnTypeResponse, error) {
// 	balance, err := s.balanceDomain.GetBalanceByUserIDnCoinnType(ctx, req.GetUserID(), req.GetCoin(), models.BalanceType(req.GetType()))
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &pbDTO.GetBalanceByUserIDnCoinnTypeResponse{
// 		Amount: balance.String(),
// 	}, nil
// }
