package cmsStruct

type AddAnimationReq struct {
	Label string `schema:"label,required"`
	Desc  string `schema:"desc"`
}
