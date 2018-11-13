package protocerr

import "github.com/luckahx/go-lib/cerr"

func ToProto(e error) *CError {
	res := &CError{}
	if eFull, ok := e.(cerr.CError); ok {
		res.Code = eFull.GetCode()
		res.Msg = eFull.GetMsg()
		if eFull.GetCause() != nil {
			if causeErr, ok := eFull.GetCause().(cerr.CError); ok {
				res.Cause = ToProto(causeErr)
			}
		}
	} else {
		res.Msg = e.Error()
	}
	return res
}

func FromProto(e *CError) *cerr.CErrorMessage {
	res := cerr.CErrorMessage{
		Code: e.Code,
		Msg:  e.Msg,
	}
	if e.Cause != nil {
		res.Cause = FromProto(e.Cause)
	}
	return &res
}

func (e *CError) Error() string {
	return FromProto(e).Error()
}

func (e *CError) FullMessage() string {
	return FromProto(e).FullMessage()
}
func (e *CError) FullErrorStack() []cerr.CError {
	return FromProto(e).FullErrorStack()
}
func (e *CError) Has(code string) bool {
	return FromProto(e).Has(code)
}
func (e *CError) Get(code string) cerr.CError {
	return FromProto(e).Get(code)
}
func (e *CError) GetFirst(codes ...string) (string, cerr.CError) {
	return FromProto(e).GetFirst(codes...)
}
func (e *CError) GetTS() int64 {
	return FromProto(e).GetTS()
}
