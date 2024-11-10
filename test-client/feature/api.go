package feature

type IF[TReq any, TRsp any] interface {
	MakeReq() (TReq, error)
	HandleRsp(TRsp) error
}
