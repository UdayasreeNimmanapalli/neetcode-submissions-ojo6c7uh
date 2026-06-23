type Meal struct {
    cost    float64
    takeOut bool
    main    string
    drink   string
}

func (m *Meal) GetCost() float64 {
    return m.cost
}

func (m *Meal) GetTakeOut() bool {
    return m.takeOut
}

func (m *Meal) GetMain() string {
    return m.main
}

func (m *Meal) GetDrink() string {
    return m.drink
}

func (m *Meal) SetCost(cost float64) {
    m.cost = cost
}

func (m *Meal) SetTakeOut(takeOut bool) {
    m.takeOut = takeOut
}

func (m *Meal) SetMain(main string) {
    m.main = main
}

func (m *Meal) SetDrink(drink string) {
    m.drink = drink
}

type MealBuilder struct {
    meal *Meal
}

func NewMealBuilder() *MealBuilder {
    return &MealBuilder{
        meal:&Meal{},
    }
}

func (mb *MealBuilder) AddCost(cost float64) *MealBuilder {
    mb.meal.SetCost(cost)
    return mb
}

func (mb *MealBuilder) AddTakeOut(takeOut bool) *MealBuilder {
    mb.meal.SetTakeOut(takeOut)
    return mb
}

func (mb *MealBuilder) AddMainCourse(main string) *MealBuilder {
    mb.meal.SetMain(main)
    return mb
}

func (mb *MealBuilder) AddDrink(drink string) *MealBuilder {
    mb.meal.SetDrink(drink)
    return mb
}

func (mb *MealBuilder) Build() *Meal {
   return mb.meal
}
