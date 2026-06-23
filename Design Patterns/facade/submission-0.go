type Order struct {
    contents string
    takeOut  bool
}

func NewOrder(contents string, takeOut bool) *Order {
    return &Order{contents: contents, takeOut: takeOut}
}

func (o *Order) GetOrder() string {
    return o.contents
}

func (o *Order) IsTakeOut() bool {
    return o.takeOut
}

type Cashier struct{}

func NewCashier() *Cashier {
    return &Cashier{}
}

func (c *Cashier) TakeOrder(contents string, takeOut bool) *Order {
    return NewOrder(contents, takeOut)
}

type Food struct {
    contents string
}

func NewFood(order string) *Food {
    return &Food{contents: order}
}

func (f *Food) GetFood() string {
    return f.contents
}

type Chef struct{}

func NewChef() *Chef {
    return &Chef{}
}

func (c *Chef) PrepareFood(order *Order) *Food {
    return NewFood(order.GetOrder())
}

type PackagedFood struct {
    *Food
}

func NewPackagedFood(food *Food) *PackagedFood {
    return &PackagedFood{Food: NewFood(food.GetFood() + " in a bag")}
}

type KitchenStaff struct{}

func NewKitchenStaff() *KitchenStaff {
    return &KitchenStaff{}
}

func (k *KitchenStaff) PackageOrder(food *Food) *PackagedFood {
    return NewPackagedFood(food)
}

type DriveThruFacade struct {
    cashier      *Cashier
    chef         *Chef
    kitchenStaff *KitchenStaff
}

func NewDriveThruFacade() *DriveThruFacade {
    return &DriveThruFacade{
        cashier:      NewCashier(),
        chef:         NewChef(),
        kitchenStaff: NewKitchenStaff(),
    }
}

func (d *DriveThruFacade) TakeOrder(orderContents string, takeOut bool) *Food {
    order := d.cashier.TakeOrder(orderContents, takeOut)
    food:= d.chef.PrepareFood(order)
    if order.IsTakeOut(){
        return d.kitchenStaff.PackageOrder(food).Food
    }
    return food
}
