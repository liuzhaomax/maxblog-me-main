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

func Pb2DemoRes(demoRes *pb.DemoRes) DemoRes {
	return DemoRes{
		Id:        demoRes.Id,
		Title:     demoRes.Title,
		Desc:      demoRes.Desc,
		CreatedAt: demoRes.CreatedAt,
		Content:   demoRes.Content,
		View:      demoRes.View,
		Ref:       demoRes.Ref,
	}
}

type DemosDataRes struct {
	Id        uint32 `json:"id"`
	Title     string `json:"title"`
	Desc      string `json:"desc"`
	CreatedAt string `json:"createdAt"`
	View      uint32 `json:"view"`
	Preview   string `json:"preview"`
}

type DemosRes []DemosDataRes

func Pb2DemosDataRes(demosDataRes *pb.DemoRes) DemosDataRes {
	return DemosDataRes{
		Id:        demosDataRes.Id,
		Title:     demosDataRes.Title,
		Desc:      demosDataRes.Desc,
		CreatedAt: demosDataRes.CreatedAt,
		Preview:   demosDataRes.Preview,
		View:      demosDataRes.View,
	}
}
