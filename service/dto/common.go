package dto

type CommonIDDTO struct {
	Id uint `json:"id" form:"id" uri:"id" binding:"required" message:"id不能为空" required_err:"id不能为空"`
}

type Paginate struct {
	Page  int `json:"page,omitempty" form:"page"`
	Limit int `json:"limit,omitempty" form:"limit"`
}

func (m *Paginate) GetPage() int {
	if m.Page <= 0 {
		m.Page = 1
	}
	return m.Page
}

func (m *Paginate) GetLimit() int {
	if m.Limit <= 0 {
		m.Limit = 10
	}
	return m.Limit
}
