package data

import (
	"context"
)

type Pic interface {
	RespPicMessage(context.Context, int32) *PicRepo
}

type PicRepo struct {
	Id int32
	Name string
	Url string
}

func NewPicRepo() *PicRepo {
	return &PicRepo{}
}

func (p *PicRepo) RespPicMessage(ctx context.Context, id int32) *PicRepo {
	return &PicRepo{Id: id, Name: "PicName", Url: "PicUrl"}
}