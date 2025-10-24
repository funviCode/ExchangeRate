//package main   ЧЕРЕЗ ГОРУТИНЫ
//
//import (
//	"awesomeProject1/parsings"
//	"awesomeProject1/write_to_file"
//	_ "awesomeProject1/write_to_file"
//	"fmt"
//	"log"
//	"sync"
//	"time"
//)
//
//type Currency struct {
//	Name     string
//	URL      string
//	Selector string
//}
//type ParseResult struct {
//	Product Currency
//	Price   string
//	Error   error
//}
//
//func main() {
//	htmlSearch := "div.text-5xl\\/9.font-bold.text-\\[\\#232526\\].md\\:text-\\[42px\\].md\\:leading-\\[60px\\]"
//	currency := []Currency{
//		{"USD к рублю", "https://ru.investing.com/currencies/usd-rub", htmlSearch},
//		{"EUR к рублю", "https://ru.investing.com/currencies/eur-rub", htmlSearch},
//		{"CNY к рублю", "https://ru.investing.com/currencies/cny-rub", htmlSearch},
//		{"Bitcoin к USD", "https://ru.investing.com/crypto/bitcoin", htmlSearch},
//		{"Bitcoin к рублю", "https://ru.investing.com/crypto/bitcoin/btc-rub", htmlSearch},
//	}
//
//	//настраиваем время
//	now := time.Now()
//	timeFormat := now.Format("2006-01-02 15:04:05")
//	timeWr := fmt.Sprintf("\n[----------%s----------] \n", timeFormat)
//
//	//пишем время в файл
//	msg, err := write_to_file.WriteToFile([]byte(timeWr))
//	if err != nil {
//		fmt.Println("Ошибка при записи файла", err)
//	} else {
//		fmt.Println(msg)
//	}
//
//	results := make(chan ParseResult, len(currency))
//
//	var wg sync.WaitGroup
//	// Запуск горутин для каждого продукта
//	for _, currencys := range currency {
//		wg.Add(1)
//
//		go func(cur Currency) {
//			defer wg.Done()
//
//			price, err := parsings.AllParsingInvesting(cur.URL, cur.Selector)
//			results <- ParseResult{Product: cur, Price: price, Error: err}
//		}(currencys)
//	}
//
//	// Закрываем канал, когда все горутины завершатся
//	go func() {
//		wg.Wait()
//		close(results)
//	}()
//
//	// Обработка результатов
//	for result := range results {
//		if result.Error != nil {
//			log.Printf("Ошибка для %s: %v\n", result.Product.Name, result.Error)
//			continue
//		}
//		fmt.Println(result.Product.Name, result.Price)
//
//		msg, err := write_to_file.WriteToFile([]byte(fmt.Sprintf("%s: %s\n", result.Product.Name, result.Price)))
//		if err != nil {
//			fmt.Println("Ошибка при записи файла", err)
//		} else {
//			fmt.Println(msg)
//		}
//
//	}
//
//	fmt.Println("Нажмите Enter чтобы закрыть...")
//	_, _ = fmt.Scanln() // Ждет нажатия Enter
//}

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
	_, _ = fmt.Scanln() // Ждет нажатия Enterv

}

//make build
