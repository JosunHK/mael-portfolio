package cmsStruct

import "database/sql"

type AddAnimationReq struct {
	Label string `schema:"label,required"`
}

type ModifyAnimationReq struct {
	Label       string        `schema:"label,required"`
	FramesCount sql.NullInt32 `schema:"framesCount"`
	Fps         sql.NullInt32 `schema:"fps"`
	Desc        string        `schema:"desc"`
}
