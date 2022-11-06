package service

import (
	"awesomeProject/model/web"
	"context"
)

type Requestotpservice_request interface {
	RequestOTP(ctx context.Context, request web.RequestOTP_request) web.RequestOTP_response
}
