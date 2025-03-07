package main

import (
	"awesomeProject1/parsings"
	"awesomeProject1/write_to_file"
	_ "awesomeProject1/write_to_file"
	"fmt"
	"time"
)

type Currency struct {
	Name     string
	URL      string
	Selector string
}

func main() {
	now := time.Now()
	timeFormat := now.Format("2006-01-02 15:04:05")
	timeWr := fmt.Sprintf("\n[----------%s----------] \n", timeFormat)

	htmlSearch := "div.text-5xl\\/9.font-bold.text-\\[\\#232526\\].md\\:text-\\[42px\\].md\\:leading-\\[60px\\]"

	currency := []Currency{
		{"USD к рублю", "https://ru.investing.com/currencies/usd-rub", htmlSearch},
		{"EUR к рублю", "https://ru.investing.com/currencies/eur-rub", htmlSearch},
		{"CNY к рублю", "https://ru.investing.com/currencies/cny-rub", htmlSearch},
		{"Bitcoin к USD", "https://ru.investing.com/crypto/bitcoin", htmlSearch},
		{"Bitcoin к рублю", "https://ru.investing.com/crypto/bitcoin/btc-rub", htmlSearch},
	}

	msg, err := write_to_file.WriteToFile([]byte(timeWr))
	if err != nil {
		fmt.Println("Ошибка при записи файла", err)
	} else {
		fmt.Println(msg)
	}

	for _, c := range currency {
		money, err := parsings.AllParsingInvesting(c.URL, c.Selector)
		if err != nil {
			fmt.Printf("Ошибка для $s: %v\n", c.Name, err)
			continue
		}
		fmt.Println(c.Name, money)

		msg, err := write_to_file.WriteToFile([]byte(fmt.Sprintf("%s: %s\n", c.Name, money)))
		if err != nil {
			fmt.Println("Ошибка при записи файла", err)
		} else {
			fmt.Println(msg)
		}

	}
	fmt.Println("Нажмите Enter чтобы закрыть...")
	_, _ = fmt.Scanln() // Ждет нажатия Enter

}
