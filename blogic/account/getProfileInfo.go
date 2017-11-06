package account

import "github.com/adgluzdov/IMobyServer/model/account"

type GetProfileInfo interface {
	Go(request *account.GetProfileInfoRequest)(response account.GetProfileInfoResponse,err error)
}
