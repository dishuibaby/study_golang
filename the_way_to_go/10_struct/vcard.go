package main

import (
	"fmt"
	"time"
)

// Address 结构体，包含地址相关信息
type Address struct {
	Street  string
	City    string
	State   string
	Country string
	ZipCode string
}

// VCard 结构体，包含个人信息
type VCard struct {
	Name      string
	Addresses []*Address // 使用指针类型
	Birthday  time.Time
	ImageURL  string
}

func main() {
	// 创建 Address 实例
	homeAddress := &Address{
		Street:  "123 Main St",
		City:    "Anytown",
		State:   "CA",
		Country: "USA",
		ZipCode: "12345",
	}
	workAddress := &Address{
		Street:  "456 Elm St",
		City:    "Anytown",
		State:   "CA",
		Country: "USA",
		ZipCode: "12345",
	}

	// 创建 VCard 实例
	myVCard := VCard{
		Name:      "John Doe",
		Addresses: []*Address{homeAddress, workAddress},
		Birthday:  time.Date(1990, time.January, 1, 0, 0, 0, 0, time.UTC),
		ImageURL:  "http://example.com/image.jpg",
	}

	// 打印 VCard
	fmt.Printf("Name: %s\n", myVCard.Name)
	fmt.Println("Addresses:")
	for _, addr := range myVCard.Addresses {
		fmt.Printf(" - Street: %s, City: %s, State: %s, Country: %s, ZipCode: %s\n",
			addr.Street, addr.City, addr.State, addr.Country, addr.ZipCode)
	}
	fmt.Printf("Birthday: %s\n", myVCard.Birthday.Format("2006-01-02"))
	fmt.Printf("Image URL: %s\n", myVCard.ImageURL)
}
