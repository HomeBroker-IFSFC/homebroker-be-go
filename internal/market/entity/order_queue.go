package entity


type OrderQueue struct {
	Orders []*Order

}


func (oq * OrderQueue) Swap( k int, l int) {
	oq.Orders[k], oq.Orders[l] = oq.Orders[l], oq.Orders[k]
}

//Less
func (oq * OrderQueue) Less(i int,j int) bool {
	return oq.Orders[i].Price < oq.Orders[j].Price
}

//Len
func (oq * OrderQueue) Len() int {
	return len(oq.Orders)
}

//Push
func (oq * OrderQueue) Push (x any) {
	oq.Orders = append(oq.Orders, x.(*Order))
}

//Pop
func (oq * OrderQueue ) Pop () interface {} {
	old:= oq.Orders
	n:=len(old);
	item:= old[n-1]
	oq.Orders = old[0 : n-1]
	return item
}

func NewOrderQueue () * OrderQueue {
	return &OrderQueue{}
}


