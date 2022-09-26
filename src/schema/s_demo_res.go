package schema

import "maxblog-me-main/src/pb"

type DemoRes struct {
	Id        uint32 `json:"id"`
	Title     string `json:"title"`
	Desc      string `json:"desc"`
	CreatedAt string `json:"createdAt"`
	Content   string `json:"content"`
	View      uint32 `json:"view"`
	Ref       string `json:"ref"`
}

type DemosRes []DemoRes

func Pb2Res(dataRes *pb.DemoRes) DemoRes {
	return DemoRes{
		Id:        dataRes.Id,
		Title:     dataRes.Title,
		Desc:      dataRes.Desc,
		CreatedAt: dataRes.CreatedAt,
		Content:   dataRes.Content,
		View:      dataRes.View,
		Ref:       dataRes.Ref,
	}
}
