package workorders

type Status string

var (
	StatusNew       Status = "new"
	StatusDone      Status = "done"
	StatusCancelled Status = "cancelled"
)
