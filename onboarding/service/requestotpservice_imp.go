package service

import (
	"awesomeProject/helper"
	"awesomeProject/model/web"
	"awesomeProject/repository"
	"context"
	"database/sql"
)

type RequestOTPservice_imp struct {
	RequestRepo repository.Requestotprepository
	DB          sql.DB
}

func (service RequestOTPservice_imp) RequestOTP(ctx context.Context, request web.RequestOTP_request) web.RequestOTP_response {
	//TODO implement me
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	helper.CommitorRollback(tx)

	panic("implement me")
}
