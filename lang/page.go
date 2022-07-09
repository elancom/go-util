package lang

import "encoding/json"

// 默认页面
const defaultPage = 1

// 默认页大小
const defaultRows = 10

func NewPage(pageAndRows ...int) *Page {
	t := new(Page)
	pLen := len(pageAndRows)
	if pLen > 0 {
		t.page = pageAndRows[0]
		if pLen > 1 {
			t.rows = pageAndRows[1]
		}
	}
	return t
}

type Page struct {
	page    int
	rows    int
	maxRows int // 0代表未设置
}

// SetPage 设置页码
func (p *Page) SetPage(page int) {
	if page > 0 {
		p.page = page
	}
}

// GetPage 页码 基于1索引开始的页
func (p *Page) GetPage() int {
	if p.page <= 0 {
		return defaultPage
	}
	return p.page
}

// GetPageIndex 页码 基于0索引开始的页
func (p *Page) GetPageIndex() int {
	return p.GetPage() - 1
}

// TrySetRows 如果页大小未指定则指定
func (p *Page) TrySetRows(rows int) {
	// 页未设置
	if p.rows <= 0 {
		p.SetRows(rows)
	}
}

// SetRows 设置页大小 最大不超过最大页面(如果指定最大页面)
func (p *Page) SetRows(rows int) {
	if rows >= 0 {
		p.rows = rows
	}
	p.checkAndSetRows()
}

// 检查是否超过最大页面 如果超过则设置页大小为最大页大小
func (p *Page) checkAndSetRows() {
	// 未指定
	if p.maxRows <= 0 {
		return
	}
	// 防止超过最大页面
	if p.rows > p.maxRows {
		p.rows = p.maxRows
	}
}

// SetMaxRows 设置最大页大小 自动修正页大小
func (p *Page) SetMaxRows(maxRows int) {
	if maxRows < 0 {
		return
	}
	p.maxRows = maxRows
	p.checkAndSetRows()
}

// GetMaxRows 取最大页大小
func (p *Page) GetMaxRows() int {
	return p.maxRows
}

// GetRows 取页大小
func (p *Page) GetRows() int {
	r := p.rows

	// 未设置页大小
	if r <= 0 {
		r = defaultRows
	}

	// 指定最大页
	if p.maxRows > 0 && r > p.maxRows {
		r = p.maxRows
	}

	return r
}

// GetFirst 取第1条记录索引
func (p *Page) GetFirst() int {
	return p.GetPageIndex() * p.GetRows()
}

func (p *Page) Clone() *Page {
	p2 := new(Page)
	p2.page = p.page
	p2.rows = p.rows
	p2.maxRows = p.maxRows
	return p2
}

func (p *Page) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"page":  p.GetPage(),
		"rows":  p.GetRows(),
		"first": p.GetFirst(),
	})
}

func NewFlagList() *Flag {
	f := new(Flag)
	f.IsList = true
	return f
}

func NewFlagCount() *Flag {
	f := new(Flag)
	f.IsCount = true
	return f
}

func NewFlagSummary() *Flag {
	f := new(Flag)
	f.IsSummary = true
	return f
}

type Flag struct {
	IsList    bool
	IsCount   bool
	IsSummary bool
	// IsV1      bool
	// IsV2      bool
}

func NewQueryRs(list any, total any, summary ...any) *QueryRs {
	q := new(QueryRs)
	q.List = list
	q.Total = total
	if len(summary) > 0 {
		q.Summary = summary[0]
	}
	return q
}

type QueryRs struct {
	List    any `json:"list,omitempty"`
	Total   any `json:"total,omitempty"`
	Summary any `json:"summary,omitempty"`
}
