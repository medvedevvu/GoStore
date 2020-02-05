package GoStore

import (
	"testing"

	sc "github.com/youricorocks/shop_competition"
)

func InitProductList() map[string]*sc.Product {
	/* Получилось название товара в 2 местах */
	productList := map[string]*sc.Product{}
	productList["Колбаса"] = &sc.Product{Name: "Колбаса", Price: 102.45, Type: sc.ProductNormal}
	productList["Вареники"] = &sc.Product{Name: "Вареники", Price: 148.25, Type: sc.ProductPremium}
	productList["Вермишель"] = &sc.Product{Name: "Вермишель", Price: 249.31, Type: sc.ProductPremium}
	return productList
}

func TestAddWrongProduct(t *testing.T) {
	var testProductList ProductList = InitProductList()
	product := sc.Product{Name: "", Price: 995.31, Type: sc.ProductNormal}
	err := testProductList.AddProduct(product)
	if err == nil {
		t.Fatalf("Могу добавлять товары с пустым названием %v ", testProductList)
	}
}

func TestModifyProductOnWrongValue(t *testing.T) {
	var testProductList ProductList = InitProductList()
	err := testProductList.ModifyProduct(sc.Product{Name: "вареники:",
		Price: 995.31, Type: /*sc.ProductNormal*/ 's'})
	if err == nil {
		t.Fatalf("Могу использовать с типом символа s = %v ", testProductList)
	}
}

func TestAddSamegProduct(t *testing.T) {
	var testProductList ProductList = InitProductList()
	product := sc.Product{Name: "вареники", Price: 995.31, Type: sc.ProductNormal}
	err := testProductList.AddProduct(product)
	product = sc.Product{Name: "вАрЕники", Price: 995.31, Type: sc.ProductNormal}
	err = testProductList.AddProduct(product)
	product = sc.Product{Name: "вАрЕнИкИ", Price: 995.31, Type: sc.ProductNormal}
	err = testProductList.AddProduct(product)
	if err == nil {
		t.Fatalf("Могу добавлять товары с одинаковыми\n атрибутами но с похожими именами\n %v\n ", testProductList)
	}
}
