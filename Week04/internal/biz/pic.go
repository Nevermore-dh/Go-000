package biz

import (
	"Week04/internal/data"
	"github.com/spf13/viper"
)

type PicBiz interface {
	DealPicInfo(*data.PicRepo)
}

type PicLogic struct {}

func NewPicLogic() *PicLogic {
	return &PicLogic{}
}

func (pl *PicLogic) DealPicInfo(pr *data.PicRepo) {
	pr.Url = viper.GetString("PicHost") + pr.Url
}