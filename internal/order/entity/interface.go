package entity

type OrderRepositoryInterface interface {
	Save(*Order) error
	GetTotal() (int, error)
}
