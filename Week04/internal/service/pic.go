package service

import (
	v1 "Week04/api/pic/v1"
	"Week04/internal/biz"
	"Week04/internal/data"
	"context"
)

type PicService struct {}

func NewPicService() *PicService {
	return &PicService{}
}

func (s *PicService) GetPicInfoById(ctx context.Context, r *v1.GetPicInfoByIdRequest) (*v1.GetPicInfoByIdResponse, error) {
	pr := data.NewPicRepo()
	rpm := pr.RespPicMessage(context.TODO(), r.Id)
	pl := biz.NewPicLogic()
	pl.DealPicInfo(rpm)
	return &v1.GetPicInfoByIdResponse{Id: rpm.Id, Name: rpm.Name, Url: rpm.Url}, nil
}