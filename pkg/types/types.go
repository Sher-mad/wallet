package types

// Money представляет собой в минимальных единицах (центы, копейки, дирам и т.д. )
type Money int64

// Category представляет собой категориб, в котором был совершён платёж (авто, аптеки, рестораы и т.д.).
type PaymentCategory string

// Status представляет собой статус платежа.
type PaymentStatus string

// Status представляет статусы платежей.
const (
	PaymentStatusOk         PaymentStatus = "OK"
	PaymentStatusFail       PaymentStatus = "FAIL"
	PaymentStatusInProgress PaymentStatus = "INPROGRESS"
)

type Phone string

//Payment представляет информацию о платеже
type Payment struct {
	ID        string
	AccountID int64
	Amount    Money
	Category  PaymentCategory
	Status    PaymentStatus
}

// Accaunt представляет информацию о  счёте пользователя.
type Account struct {
	ID      int64
	Phone   Phone
	Balance Money
}