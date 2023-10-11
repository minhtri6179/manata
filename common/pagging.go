package common

type Pagging struct {
	Page   int32 `json:"page" form:"page"`
	Limit  int32 `json:"limit" form:"limit"`
	Offset int32 `json:"offset" `
}

func (p *Pagging) HandlePaging() {
	if p.Page < 1 {
		p.Page = 1
	}
	if p.Limit <= 1 {
		p.Limit = 1
	}
	if p.Limit > 100 {
		p.Limit = 100
	}
}
