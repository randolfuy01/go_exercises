package bank

import (
	"fmt"
)

func main() {
	fmt.Print("Starting banking service\n")
	fmt.Print("What would you like to do?\n [1] create account, [0] exit\n")

  var user_choice int
  fmt.Scan(&user_choice)

  if user_choice == 1 {
    fmt.Print("Please enter an account name")
    var user_name string
    fmt.Scanln(&user_name)

    fmt.Print("Please enter amount you would for an initial deposit")
    var initial_deposit float32
    fmt.Scan(&initial_deposit)

    fmt.Print("Please enter an address in the form of 'number, street, city, state'"")
    fmt.Scanln()
  }
}
