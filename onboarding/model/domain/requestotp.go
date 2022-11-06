package domain

import (
	"time"
)

type Requestotp struct {
	Initregid      int
	Phonenum       string
	Otp            string
	Otpexpireddate time.Time
	Otpsendvia     string
	Isotpmatched   string
	Otpmatcheddate time.Time
	Secretquestion string
	Secretanswer   string
}
