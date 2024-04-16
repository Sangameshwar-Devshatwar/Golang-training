package account

import "time"

type Account struct {
	AcNum         int
	Name          string
	Date_of_Birth time.Time
	Acc_Type      string
	Pin           int
	Balance       float64
}

func (a *Account) SetName(nm string) {
	a.Name = nm
}
func (a *Account) SetBalance(balance float64) {
	a.Balance = balance
}
func (a *Account) GetACNum(ac int) int {
	a.AcNum = ac
	return a.AcNum

}

func (a *Account) SetAccType(ac_type string) {
	a.Acc_Type = ac_type
}
func (a *Account) GetPin(pincode int) int {
	a.Pin = pincode
	return a.Pin
}
func SetPin(input int, acc_num int, bank *map[int]Account) {
	holder := (*bank)[acc_num]
	holder.Pin = input
	(*bank)[acc_num] = holder
}
func (a *Account) SetDob(date time.Time) {
	date.Format("2006-01-02")
	a.Date_of_Birth = date
}

func DepositAmount(amount float64, acc int, Bank *map[int]Account) {

	holder := (*Bank)[acc]
	holder.Balance = holder.Balance + amount
	(*Bank)[acc] = holder

}

func WithdrawAmount(amount float64, Bank *map[int]Account, acc int) {
	holder := (*Bank)[acc]
	holder.Balance = holder.Balance - amount
	(*Bank)[acc] = holder
}
func CheckBalance(acc int, Bank *map[int]Account) float64 {
	holder := (*Bank)[acc]
	return holder.Balance
}
