package schema

import "maxblog-me-main/src/pb"

type DataRes struct {
	Mobile string `json:"mobile"`
}

func Pb2Res(dataRes *pb.DataRes) DataRes {
	return DataRes{
		Mobile: dataRes.Mobile,
	}
}
