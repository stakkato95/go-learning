package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	Id      string
	Name    string
	Address string
}

type Account struct {
	Money     int
	CompanyId string
	Person
}

type Building struct {
	Address string
	Name    string
	Owner   Person
}

func main() {
	defer fmt.Println("defer one")
	defer fmt.Println("defer two")
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic happened and recovered :)", err)
		}
	}()
	defer func() {
		deferredFun()
	}()

	//ar
	//ar := [...]int{34, 2, 90}

	//add multiple values to slice
	sl := make([]int, 3)
	sl = append(sl, 76, 91)
	printSlice(sl)

	//concat two clices
	sl = append(sl, sl...)
	fmt.Println("concatenated")
	printSlice(sl)
	fmt.Println(strconv.Itoa(len(sl)) + " " + strconv.Itoa(cap(sl)))

	//copy slice
	sl2 := make([]int, len(sl), len(sl))
	copy(sl2, sl)
	fmt.Println("copied slice")
	fmt.Println(sl2)

	if element := sl2[3]; element == 76 {
		fmt.Println("It's " + strconv.Itoa(element))
	}

	//control structs
	strVal := "str"
	switch strVal {
	case "name":
		fallthrough
	case "str", "str1":
		fmt.Println("It's either str or str1")
	default:
		fmt.Println("unknown option")
	}
	//breaking a loop
	fmt.Println("break example")

Loop:
	for _, v := range sl2 {
		switch {
		case v < 90:
			if v == 76 {
				//breaks case
				break
			}
			fmt.Println("first " + strconv.Itoa(v))
		case v == 91:
			fmt.Println("second")
			break Loop
		}
	}

	//panic example
	//panic("something bad")

	//struct pretty print example
	p := Person{
		Id:      "id-1",
		Name:    "Alex",
		Address: "Hello Addr",
	}
	fmt.Printf("%#v\n", p)

	ac := Account{
		Money:     100500,
		CompanyId: "id-company",
		Person:    p,
	}
	fmt.Printf("%#v\n", ac)
	//accessing implicitly calling person
	fmt.Println(ac.Address)

	b := Building{
		Address: "build addr",
		Name:    "high tower",
		Owner: Person{
			Id:      "id-2",
			Name:    "John",
			Address: "build addr",
		},
	}
	fmt.Printf("%#v\n", b)
	//accessing explicitly calling person
	fmt.Printf("%#v\n", b.Owner.Address)

	//inheritance
	ac.SetId("new-id")
	fmt.Printf("%#v\n", ac)
	fmt.Printf("%T\n", ac) //prints type

	//interfaces
	payer1 := &Wallet{Cash: 100}
	var payer Payer = &Wallet{Cash: 100}
	applePay := ApplePay{Bonus: 10, Money: 50}
	Buy(100, payer1)
	Buy(100, payer)
	Buy(100, &applePay)
	//type switch
	whoYouAre(payer1)

	//empty interface
	fmt.Printf("info about wallet: %s\n", payer)
	fmt.Printf("info about apple: %s\n", &applePay)
	whoYouAreForAll(payer1)
	whoYouAreForAll([]int{1, 2, 3})

	//composition of interfaces
	samsung := &Samsung{
		Number: "+4366566666666",
		payer:  ApplePay{Bonus: 12, Money: 45},
		PhoneBook: map[string]string{
			"Mike": "+2",
			"John": "+1",
		},
	}
	Buy(13, samsung)
	TakeCall(samsung, "+3")
}

func printSlice(slice []int) {
	for _, v := range slice {
		fmt.Println(v)
	}
}

func deferredFun() {
	fmt.Println("hello defer")
}

//struct methods
func (this *Person) SetId(id string) {
	this.Id = id
}

func (this *Account) SetId(id string) {
	this.CompanyId = id
	//access to all methods of person
	this.Person.SetId(id + "-person")
	//this.SetId <- this call invokes Account.SetId recursively
}

//interfaces
type Payer interface {
	Pay(int) error
}

type Wallet struct {
	Cash int
}

type ApplePay struct {
	Bonus int
	Money int
}

func (this *Wallet) Pay(amount int) error {
	if this.Cash < amount {
		return fmt.Errorf("not enough money in wallet")
	}
	this.Cash -= amount
	return nil
}

func (this *ApplePay) Pay(amount int) error {
	if this.Money+this.Bonus < amount {
		return fmt.Errorf("not enough money on apple pay")
	}
	this.Money = this.Money + this.Bonus - amount
	this.Bonus = 0

	return nil
}

func (this *Wallet) String() string {
	return "Wallet with " + strconv.Itoa(this.Cash) + " euros"
}

func (this *ApplePay) String() string {
	return "Apple Pay account with " + strconv.Itoa(this.Money+this.Bonus) + " euros"
}

func Buy(amount int, payer Payer) {
	if err := payer.Pay(amount); err != nil {
		fmt.Printf("Error while paying with %T: %#v\n", payer, err)
		return
	}
	fmt.Printf("Successfully paid with %T\n", payer)
}

func whoYouAre(payer Payer) {
	switch payer.(type) {
	case *Wallet:
		fmt.Println("It's a wallet")
		_, ok := payer.(*Wallet)
		if ok {
			fmt.Println("It's 100% wallet")
		}
	case *ApplePay:
		fmt.Println("It's ApplePay")
		_, ok := payer.(*ApplePay)
		if ok {
			fmt.Println("It's 100% apple")
		}
	}
}

func whoYouAreForAll(in interface{}) {
	if _, ok := in.(Payer); !ok {
		fmt.Printf("%T is not a payment method\n", in)
		return
	}

	fmt.Printf("%T is a payment method\n", in)
}

//interface composition
type Ringer interface {
	Call(string) error
	Answer(string) error
}

type NFCPhone interface {
	Payer
	Ringer
}

type Samsung struct {
	Number    string
	PhoneBook map[string]string
	payer     ApplePay
}

func (this *Samsung) Pay(amount int) error {
	return this.payer.Pay(amount)
}

func (this *Samsung) Call(contact string) error {
	number, isPresent := this.PhoneBook[contact]
	if !isPresent {
		return fmt.Errorf("contact not found\n")
	}
	fmt.Printf("Calling number %s\n", number)

	return nil
}

func (this *Samsung) Answer(callingNumber string) error {
	for _, number := range this.PhoneBook {
		if number == callingNumber {
			return nil
		}
	}
	return fmt.Errorf("an unknown number is calling\n")
}

func TakeCall(ringer Ringer, phoneNumber string) {
	if err := ringer.Answer(phoneNumber); err != nil {
		fmt.Printf("Error when taking a call: %s", err)
		return
	}
	fmt.Println("Successfully taken call")
}
