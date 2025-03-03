package server

import (
	"context"
	"f1betting/betting_system"
	"f1betting/proto"

	"github.com/jackc/pgx/v5"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type BettingSystemServer struct {
	proto.UnimplementedBettingServiceServer
	conn *pgx.Conn
}

func NewBettingSystemServer(conn *pgx.Conn) *BettingSystemServer {
	return &BettingSystemServer{conn: conn}
}

func (s *BettingSystemServer) CreateFastestLapBet(ctx context.Context, req *proto.FastestLapBetRequest) (*proto.BetResponse, error) {
	var bet betting_system.FastestLapBet
	bet.SetFastestLapBet(req.UserId, int(req.SessionId), req.DriverId, int64(req.BettingPool))

	betID, err := betting_system.CreateFastestLapBet(ctx, s.conn, bet)
	if err != nil {
		return nil, err
	}

	return &proto.BetResponse{
		BetId:  betID,
		Status: "PENDING",
		Amount: 0.95 * float64(req.BettingPool), // 5% commission
	}, nil
}

func (s *BettingSystemServer) GetPendingBetsForSession(ctx context.Context, req *proto.SessionRequest) (*proto.SessionBetsResponse, error) {
	bets, err := betting_system.GetFastestLapBetsByRace(ctx, s.conn, int64(req.SessionId), "PENDING")
	if err != nil {
		return nil, err
	}

	var response proto.SessionBetsResponse
	for _, bet := range *bets {
		response.Bets = append(response.Bets, &proto.Bet{
			BetType:     "FASTEST_LAP",
			Id:          bet.ID,
			UserId:      bet.UserID,
			SessionId:   int32(bet.SessionID),
			BetDetails:  &proto.Bet_Driver{Driver: &proto.DriverBetDetails{DriverId: bet.DriverID}},
			Amount:      bet.Amount,
			Status:      bet.Status,
			BettingPool: int32(bet.BettingPool),
			CreatedAt:   timestamppb.New(bet.CreateAt),
		})
	}

	return &response, nil
}

func (s *BettingSystemServer) GetUserActiveBets(ctx context.Context, req *proto.UserRequest) (*proto.UserBetsResponse, error) {
	bets, err := betting_system.GetFastestLapBetsByRace(ctx, s.conn, 0, "PENDING") // 0 to get all sessions
	if err != nil {
		return nil, err
	}

	var response proto.UserBetsResponse
	for _, bet := range *bets {
		if bet.UserID == req.UserId {
			response.Bets = append(response.Bets, &proto.Bet{
				BetType:     "FASTEST_LAP",
				Id:          bet.ID,
				UserId:      bet.UserID,
				SessionId:   int32(bet.SessionID),
				BetDetails:  &proto.Bet_Driver{Driver: &proto.DriverBetDetails{DriverId: bet.DriverID}},
				Amount:      bet.Amount,
				Status:      bet.Status,
				BettingPool: int32(bet.BettingPool),
				CreatedAt:   timestamppb.New(bet.CreateAt),
			})
		}
	}

	return &response, nil
}

func (s *BettingSystemServer) GetUserBalance(ctx context.Context, req *proto.UserRequest) (*proto.BalanceResponse, error) {
	// TODO: Implement balance tracking
	return &proto.BalanceResponse{Balance: 0}, nil
}

// Implement other required methods with placeholder returns for now
func (s *BettingSystemServer) CreatePodiumBet(ctx context.Context, req *proto.PodiumBetRequest) (*proto.BetResponse, error) {
	return &proto.BetResponse{}, nil
}

func (s *BettingSystemServer) CreatePolePositionBet(ctx context.Context, req *proto.PolePositionBetRequest) (*proto.BetResponse, error) {
	return &proto.BetResponse{}, nil
}

func (s *BettingSystemServer) CreateRainBet(ctx context.Context, req *proto.RainBetRequest) (*proto.BetResponse, error) {
	return &proto.BetResponse{}, nil
}

func (s *BettingSystemServer) CreateRetirementBet(ctx context.Context, req *proto.RetirementBetRequest) (*proto.BetResponse, error) {
	return &proto.BetResponse{}, nil
}

func (s *BettingSystemServer) CreateLapTimingBet(ctx context.Context, req *proto.LapTimingBetRequest) (*proto.BetResponse, error) {
	return &proto.BetResponse{}, nil
}

func (s *BettingSystemServer) GetUserBetHistory(ctx context.Context, req *proto.UserRequest) (*proto.UserBetsResponse, error) {
	return &proto.UserBetsResponse{}, nil
}
