package domain

import "github.com/wazwki/WearStore/cart-service/api/proto/cartpb"

type CartItem struct {
	ProductID string  `json:"product_id" redis:"product_id"`
	Name      string  `json:"name" redis:"name"`
	Price     float64 `json:"price" redis:"price"`
	Quantity  int     `json:"quantity" redis:"quantity"`
}

type Cart struct {
	UserID    string      `json:"user_id" redis:"user_id"`
	Items     []*CartItem `json:"items" redis:"items"`
	TotalCost float64     `json:"total_cost" redis:"total_cost"`
}

func CartEntityToDTO(cart *Cart) *cartpb.Cart {
	items := make([]*cartpb.CartItem, len(cart.Items))
	for _, item := range cart.Items {
		items = append(items, CartItemEntityToDTO(item))
	}
	return &cartpb.Cart{
		UserId:    cart.UserID,
		Items:     items,
		TotalCost: cart.TotalCost,
	}
}

func CartItemEntityToDTO(cartItem *CartItem) *cartpb.CartItem {
	return &cartpb.CartItem{
		ProductId: cartItem.ProductID,
		Name:      cartItem.Name,
		Price:     cartItem.Price,
		Quantity:  int64(cartItem.Quantity),
	}
}
