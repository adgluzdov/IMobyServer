package model

type IMobyAuthRequest struct {
	Accsses_token string
}

type IMobyAuthResponse struct {
	Uid string
}

func (this *IMobyAuthRequest)Init(accsses_token string)  {
	this.Accsses_token = accsses_token
}
