package world

// Waiter 服務生
type Waiter struct {
	customer   *Customer
	restaurant *Restaurant
}

//DefaultWaiter 預設的服務生
var DefaultWaiter Waiter

// InitWaiter 初始化服務生
func InitWaiter() {
	DefaultWaiter = Waiter{}
}

// SetWaiter 設定服務生
func (w *Waiter) SetWaiter(r *Restaurant) {
	w.restaurant = r
}

// deliveryOrder 送做好的菜給客人
func (w *Waiter) deliveryMeal(s string) {

}

// OrderSomething 點餐
func (w *Waiter) OrderSomething(s string, cus *Customer) {

	w.customer = cus

	go func() {
		// 送餐點給廚師
		w.toChef(s)
	}()

}

func (w *Waiter) toChef(s string) {
	w.restaurant.orderChan <- s
}

func (w *Waiter) toCustomer(s string) {

	go w.customer.dining()
	w.customer.FinishChan <- s

}
