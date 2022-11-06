package repository

import (
	"awesomeProject/model/domain"
	"context"
	"database/sql"
)

type Requestotprepository interface {
	Save(ctx context.Context, tx sql.Tx, requstotp domain.Requestotp) int
}
