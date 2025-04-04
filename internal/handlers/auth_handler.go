package handlers

import (
	"context"

	"github.com/PharmaKart/authentication-svc/internal/proto"
	"github.com/PharmaKart/authentication-svc/internal/repositories"
	"github.com/PharmaKart/authentication-svc/internal/services"
	"github.com/PharmaKart/authentication-svc/pkg/errors"
	"github.com/PharmaKart/authentication-svc/pkg/utils"
)

type AuthHandler interface {
	Register(ctx context.Context, req *proto.RegisterRequest) (*proto.RegisterResponse, error)
	Login(ctx context.Context, req *proto.LoginRequest) (*proto.LoginResponse, error)
	VerifyToken(ctx context.Context, req *proto.VerifyTokenRequest) (*proto.VerifyTokenResponse, error)
}

type authHandler struct {
	proto.UnimplementedAuthServiceServer
	authService services.AuthService
}

func NewAuthHandler(userRepo repositories.UserRepository, customerRepo repositories.CustomerRepository, jwtSecret string) *authHandler {
	return &authHandler{
		authService: services.NewAuthService(userRepo, customerRepo, jwtSecret),
	}
}

func (h *authHandler) Register(ctx context.Context, req *proto.RegisterRequest) (*proto.RegisterResponse, error) {
	err := h.authService.Register(
		req.Username,
		req.Email,
		req.Password,
		req.FirstName,
		req.LastName,
		req.Phone,
		req.DateOfBirth,
		req.StreetLine1,
		req.StreetLine2,
		req.City,
		req.Province,
		req.PostalCode,
		req.Country,
	)

	if err != nil {
		// Convert the app error to proto response
		if appErr, ok := errors.IsAppError(err); ok {
			return &proto.RegisterResponse{
				Success: false,
				Message: appErr.Message,
				Error: &proto.Error{
					Type:    string(appErr.Type),
					Message: appErr.Message,
					Details: utils.ConvertMapToKeyValuePairs(appErr.Details),
				},
			}, nil
		}
		return &proto.RegisterResponse{
			Success: false,
			Message: err.Error(),
			Error: &proto.Error{
				Type:    string(errors.InternalError),
				Message: "An unexpected error occurred",
			},
		}, nil
	}

	return &proto.RegisterResponse{Success: true, Message: "Registered Successfully"}, nil
}

func (h *authHandler) Login(ctx context.Context, req *proto.LoginRequest) (*proto.LoginResponse, error) {
	token, userid, username, role, err := h.authService.Login(req.Email, req.Username, req.Password)

	if err != nil {
		// Convert the app error to proto response
		if appErr, ok := errors.IsAppError(err); ok {
			return &proto.LoginResponse{
				Success: false,
				Message: appErr.Message,
				Error: &proto.Error{
					Type:    string(appErr.Type),
					Message: appErr.Message,
					Details: utils.ConvertMapToKeyValuePairs(appErr.Details),
				},
			}, nil
		}
		return &proto.LoginResponse{
			Success: false,
			Message: err.Error(),
			Error: &proto.Error{
				Type:    string(errors.InternalError),
				Message: "An unexpected error occurred",
			},
		}, nil
	}

	return &proto.LoginResponse{Success: true, Message: "Logged in Successfully", Token: token, UserId: userid, Username: username, Role: role}, nil
}

func (h *authHandler) VerifyToken(ctx context.Context, req *proto.VerifyTokenRequest) (*proto.VerifyTokenResponse, error) {
	userid, role, err := h.authService.VerifyToken(req.Token)

	if err != nil {
		if appErr, ok := errors.IsAppError(err); ok {
			return &proto.VerifyTokenResponse{
				Success: false,
				Message: appErr.Message,
				Error: &proto.Error{
					Type:    string(appErr.Type),
					Message: appErr.Message,
					Details: utils.ConvertMapToKeyValuePairs(appErr.Details),
				},
			}, nil
		}
		return &proto.VerifyTokenResponse{
			Success: false,
			Message: err.Error(),
			Error: &proto.Error{
				Type:    string(errors.InternalError),
				Message: "An unexpected error occurred",
			},
		}, nil
	}

	return &proto.VerifyTokenResponse{Success: true, Message: "Token validated", Role: role, UserId: userid}, nil
}
