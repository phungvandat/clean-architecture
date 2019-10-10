package engine

type Condition string

const (
	Equal              Condition = "equal"
	LessThan           Condition = "less_than"
	LessThanOrEqual    Condition = "less_than_or_equal"
	GreaterThan        Condition = "greater_than"
	GreaterThanOrEqual Condition = "greater_than_or_equal"
	And                Condition = "and"
	Or                 Condition = "or"
)

type Direction int

const (
	Ascending  Direction = 1
	Descending Direction = -1
)

type Query struct {
	Name    string
	Skip    int
	Limit   int
	Orders  []*Order
	Filters []*Filter
}

type Filter struct {
	Property  string
	Condition Condition
	Value     interface{}
}

type Order struct {
	Property  string
	Direction Direction
}

func (q *Query) AddFilter(property string, condition Condition, value interface{}) {
	q.Filters = append(q.Filters, &Filter{
		Property:  property,
		Condition: condition,
		Value:     value,
	})
}

func (q *Query) AddOrder(property string, direction Direction) {
	q.Orders = append(q.Orders, &Order{
		Property:  property,
		Direction: direction,
	})
}

func (q *Query) AddSlice(skip, limit int) {
	if skip > 0 {
		q.Skip = skip
	}
	if limit > 0 {
		q.Limit = limit
	}
}
