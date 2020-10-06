package grpc

// import (
// 	"context"
// 	"xtek/exchange/commission/internal/domain/withdraw"

// 	pbPaginate "xtek/exchange/commission/pb/paginate"
// 	pb "xtek/exchange/commission/pb/wallet"
// 	pbDTO "xtek/exchange/commission/pb/wallet/dto"

// 	balance "xtek/exchange/commission/internal/domain/balance"
// 	coin "xtek/exchange/commission/internal/domain/coin"
// 	transaction "xtek/exchange/commission/internal/domain/transaction"
// 	useraddress "xtek/exchange/commission/internal/domain/useraddress"

// 	"github.com/golang/protobuf/ptypes/empty"
// 	"google.golang.org/grpc"
// 	"google.golang.org/grpc/codes"
// )

// // NewAdminSiteDelivery ...
// func NewAdminSiteDelivery(balanceDomain *balance.BalanceDomain,
// 	coinDomain *coin.CoinDomain,
// 	userAddressDomain *useraddress.UserAddressesDomain,
// 	transactionDomain *transaction.TransactionDomain,
// 	adminTransactionDomain *transaction.AdminTransactionDomain,
// 	withdrawDomain *withdraw.Withdraw,
// ) (pb.AdminSiteServiceServer, error) {
// 	return &adminSiteDelivery{balanceDomain, coinDomain, userAddressDomain, transactionDomain, adminTransactionDomain, withdrawDomain}, nil
// }

// type adminSiteDelivery struct {
// 	balanceDomain          *balance.BalanceDomain
// 	coinDomain             *coin.CoinDomain
// 	userAddressDomain      *useraddress.UserAddressesDomain
// 	transactionDomain      *transaction.TransactionDomain
// 	adminTransactionDomain *transaction.AdminTransactionDomain
// 	withdrawDomain         *withdraw.Withdraw
// }

// func (s *adminSiteDelivery) GetListWithdrawal(ctx context.Context, req *pbDTO.GetListWithdrawalRequest) (*pbDTO.GetListWithdrawalResponse, error) {
// 	if err := req.Validate(); err != nil {
// 		return nil, err
// 	}
// 	if req.GetPagination().GetPage() <= 0 {
// 		return nil, grpc.Errorf(codes.InvalidArgument, "Page cannot empty")
// 	}
// 	if req.GetPagination().GetLimit() <= 0 || req.GetPagination().GetLimit() > 50 {
// 		return nil, grpc.Errorf(codes.InvalidArgument, "Limit cannot empty")
// 	}

// 	lists, totalWithdrawal, totalPages, err := s.adminTransactionDomain.GetListWithdrawal(ctx, req)
// 	if err != nil {
// 		return nil, err
// 	}
// 	var pagination = &pbPaginate.PaginateResponse{
// 		Limit:       req.GetPagination().GetLimit(),
// 		CurrentPage: req.GetPagination().GetPage(),
// 		TotalPages:  int32(totalPages),
// 	}

// 	var rep = &pbDTO.GetListWithdrawalResponse{
// 		Lists:           lists,
// 		Pagination:      pagination,
// 		TotalWithdrawal: int32(totalWithdrawal),
// 	}
// 	return rep, nil
// }

// // GetUserTransactionHistories ...
// func (s *adminSiteDelivery) GetUserTransactionHistories(ctx context.Context, req *pbDTO.GetUserTransactionHistoriesRequest) (*pbDTO.GetUserTransactionHistoriesResponse, error) {
// 	if err := req.Validate(); err != nil {
// 		return nil, err
// 	}
// 	if req.GetPagination().GetPage() <= 0 {
// 		return nil, grpc.Errorf(codes.InvalidArgument, "Invalid page")
// 	}
// 	if req.GetPagination().GetLimit() <= 0 && req.GetPagination().GetLimit() > 20 {
// 		return nil, grpc.Errorf(codes.InvalidArgument, "Invalid limit")
// 	}
// 	lists, totalPages, err := s.transactionDomain.GetTransactionHistoriesByUserIDNCoin(ctx, req.GetUserID(), req.GetCoin(), int(req.GetType()), int(req.GetPagination().GetPage()), int(req.GetPagination().GetLimit()))
// 	if err != nil {
// 		return nil, err
// 	}
// 	var pagination = &pbPaginate.PaginateResponse{
// 		Limit:       req.GetPagination().GetLimit(),
// 		CurrentPage: req.GetPagination().GetPage(),
// 		TotalPages:  int32(totalPages),
// 	}
// 	var rep = &pbDTO.GetUserTransactionHistoriesResponse{
// 		Histories:  lists,
// 		Pagination: pagination,
// 	}

// 	return rep, nil
// }

// func (s *adminSiteDelivery) UpdateTransactionStatus(ctx context.Context, req *pbDTO.UpdateTransactionStatusRequest) (*pbDTO.UpdateTransactionStatusResponse, error) {
// 	if err := req.Validate(); err != nil {
// 		return nil, err
// 	}
// 	err := s.adminTransactionDomain.UpdateTransactionStatus(req.GetId(), req.GetStatus())
// 	if err != nil {
// 		return nil, err
// 	}

// 	var rep = &pbDTO.UpdateTransactionStatusResponse{
// 		IsSuccess: true,
// 	}
// 	return rep, nil
// }

// func (s *adminSiteDelivery) GetBalancesByUserID(ctx context.Context, req *pbDTO.GetBalancesByUserIDRequest) (*pbDTO.GetBalancesByUserIDResponse, error) {
// 	if err := req.Validate(); err != nil {
// 		return nil, err
// 	}
// 	balances, err := s.balanceDomain.GetBalancesByUserID(ctx, req.GetUserID())
// 	if err != nil {
// 		return nil, err
// 	}

// 	var rep = &pbDTO.GetBalancesByUserIDResponse{
// 		Balances: balances,
// 	}
// 	return rep, nil
// }

// // ConfirmUserWithdraw ...
// func (s *adminSiteDelivery) ConfirmUserWithdraw(ctx context.Context, req *pbDTO.ConfirmUserWithdrawRequest) (*empty.Empty, error) {
// 	if err := req.Validate(); err != nil {
// 		return nil, err
// 	}

// 	if err := s.withdrawDomain.AdminConfirmUserWithdraw(ctx, req); err != nil {
// 		return nil, err
// 	}

// 	return &empty.Empty{}, nil
// }
