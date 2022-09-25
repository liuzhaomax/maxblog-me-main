package schema

import "maxblog-me-main/src/pb"

type DemoRes struct {
	Mobile string `json:"mobile"`
}

func Pb2Res(dataRes *pb.DemoRes) DemoRes {
	return DemoRes{
		Mobile: dataRes.Mobile,
	}
}
