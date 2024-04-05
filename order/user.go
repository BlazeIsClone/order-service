package order

import "strconv"

type Order struct {
	ID          uint     `json:"id"`
	Products    []string `json:"products"`
	Destination string   `json:"destination"`
}

func orders() []Order {
	return []Order{
		{
			ID:          1,
			Products:    []string{"shampoo", "soap"},
			Destination: "No.12, york street, new burnden",
		},
		{
			ID:          2,
			Products:    []string{"phone", "endurance"},
			Destination: "No.302, frankenstein, new castle",
		},
		{
			ID:          3,
			Products:    []string{"candy bar", "t-shirt"},
			Destination: "No.92, old town, new castle",
		},
	}
}

func loadOrders() map[string]Order {
	orders := orders()
	res := make(map[string]Order, len(orders))

	for _, x := range orders {
		res[strconv.Itoa(int(x.ID))] = x
	}

	return res
}
