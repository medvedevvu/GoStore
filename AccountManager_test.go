package GoStore

import (
	"testing"
)

func InitAcc() Accounts {
	acc := Accounts{}
	acc.Register("Petya")
	acc.Register("Kola")
	acc.Register("Vitya")
	return acc
}

func TestWrongUserNameRegistert(t *testing.T) {
	accounts := InitAcc()
	err := accounts.Register("")
	if err == nil {
		t.Fatal("Можно использовать в качестве Username пустое значение")
	}
}

func TestWrongAccountType(t *testing.T) {
	accounts := InitAcc()
	err := accounts.EditType("Vitya", 's')
	if err == nil {
		t.Fatal("Можно использовать символ в качестве типа пользователя")
	}
}
