package main

import (
	"fmt"
	"os"
)

type Trade struct {
	Symbol string  // stock symbol
	Price  float64 // stock price
	Volume int     // Number of Shares
	Buy    bool    // True is buying and false if sell trade
}

// Value function returns the trade value (volume * price)
func (t *Trade) Value() float64 {
	value := float64(t.Volume) * t.Price

	// if t.Buy {
	// 	value = -value
	// }

	return value
}

// NewTrade will validate the input and create a new trade
// It will return a pointer to the new trade object, or any errors
func NewTrade(symbol string, volume int, price float64, buy bool) (*Trade, error) {
	if symbol == "" {
		return nil, fmt.Errorf("symbol can't be empty")
	}

	if volume <= 0 {
		return nil, fmt.Errorf("volume must be >= 0 (was %d)", volume)
	}

	if price <= 0.0 {
		return nil, fmt.Errorf("price must be >= 0 (was %f)", price)
	}

	trade := &Trade{
		Symbol: symbol,
		Volume: volume,
		Price:  price,
		Buy:    buy,
	}
	return trade, nil
}

func main() {
	// trade1 := Trade{
	// 	Symbol: "APPL",
	// 	Price:  149.50,
	// 	Volume: 100,
	// 	Buy:    true,
	// }

	// fmt.Println(trade1.Value())

	t, err := NewTrade("MSFT", 10, 99.98, true)

	if err != nil {
		fmt.Printf("error: can't create trade - %s\n", err)
		os.Exit(1)
	}

	fmt.Println(t.Value())
}
