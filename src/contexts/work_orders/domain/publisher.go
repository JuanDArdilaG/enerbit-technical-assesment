package workorders

type Publisher interface {
	Publish(orderID ID, status Status) error
}
