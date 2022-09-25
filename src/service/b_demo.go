package service

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	logger "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"maxblog-me-main/internal/core"
	"maxblog-me-main/src/pb"
	"maxblog-me-main/src/schema"
)

var DemoSet = wire.NewSet(wire.Struct(new(BDemo), "*"))

type BDemo struct{}

func (b *BDemo) GetDemoById(c *gin.Context, id uint32) (*schema.DemoRes, error) {
	addr := core.GetDownstreamMaxblogBETemplateAddr()
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		logger.WithFields(logger.Fields{
			"失败方法": core.GetFuncName(),
		}).Fatal(core.FormatError(300, err).Error())
		return nil, err
	}
	client := pb.NewDemoServiceClient(conn)
	pbRes, err := client.GetDemoById(context.Background(), &pb.IdRequest{Id: id})
	if err != nil {
		return nil, err
	}
	dataRes := schema.Pb2Res(pbRes)
	return &dataRes, nil
}
