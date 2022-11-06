package repository

import (
	"awesomeProject/helper"
	"awesomeProject/model/domain"
	"context"
	"database/sql"
)

type Requestotprepository_imp struct {
}

func (repository *Requestotprepository_imp) Save(ctx context.Context, tx sql.Tx, requstotp domain.Requestotp) int {
	//TODO implement me
	//panic("implement me")

	res := 0
	SQL := "insert into initregistration (phonenum ) values (?) "
	result, err := tx.ExecContext(ctx, SQL, requstotp.Phonenum)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	if err != nil {
		res = int(id)
	}

	return res
}
