package task

const (
	BestBuy Task = 1
	Target  Task = 2
)

type Task uint32

func (tk Task) String() string {
	switch tk {
	case BestBuy:
		return "bestbuy.com"
	case Target:
		return "target.com"
	default:
		return ""
	}
}

type Result struct {
	// 黑色
	Name1      string
	Price1     string
	Available1 string
	Link1      string

	// 红蓝
	Name2      string
	Price2     string
	Available2 string
	Link2      string

	IsAlarm     bool
	UpdatedTime string
}
