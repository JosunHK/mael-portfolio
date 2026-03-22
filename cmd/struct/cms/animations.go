package cmsStruct

import "database/sql"

type AddAnimationReq struct {
	Label string `schema:"label,required"`
}

type AddSubAnimationReq struct {
	Label string `schema:"label,required"`
}

type ModifyAnimationReq struct {
	Label string        `schema:"label,required"`
	Fps   sql.NullInt32 `schema:"fps"`
	Desc  string        `schema:"desc"`
}

type ModifySubAnimationReq struct {
	Label string        `schema:"label,required"`
	Fps   sql.NullInt32 `schema:"fps"`
	Desc  string        `schema:"desc"`
}

type ModifyThumbModeReq struct {
	MobileId  int `schema:"mobileId"`
	DesktopId int `schema:"desktopId"`
}

