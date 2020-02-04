package GoStore

import (
	"testing"

	"github.com/youricorocks/shop_competition"
)

func TestAddWrongNameProducts(t *testing.T) {
	// var shop shop_competition.Shop
	shop := NewStore()

	product := shop_competition.Product{
		Name:  "",
		Price: 100,
		Type:  shop_competition.ProductNormal,
	}

	if err := shop.AddProduct(product); err == nil /*&& err.(StoreError).errStatus == I*/ {
		t.Fatalf(" могу добавлять товар с пустым названием %v", product)
	}
	t.Fatalf("Finish %v", shop.ProductList)
}

func TestAddWrongBundle(t *testing.T) {
	shop := NewStore()

	product := shop_competition.Product{
		Name:  "Brimbom",
		Price: 100,
		Type:  shop_competition.ProductNormal,
	}

	product1 := shop_competition.Product{
		Name:  "XXXXX",
		Price: 100,
		Type:  shop_competition.ProductSample,
	}

	product2 := shop_competition.Product{
		Name:  "YYYYY",
		Price: 100,
		Type:  shop_competition.ProductSample,
	}

	product3 := shop_competition.Product{
		Name:  "ZZZZZZ",
		Price: 100,
		Type:  shop_competition.ProductSample,
	}

	err := shop.AddBundle("", product, 0.00001, product1, product2, product3)

	if err == nil {
		t.Fatalf("")
	}

}
