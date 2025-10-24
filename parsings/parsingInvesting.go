package parsings

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"time"
)

func AllParsingInvesting(url, selector string) (string, error) {
	// Создаем HTTP-клиент
	client := &http.Client{
		Timeout: 10 * time.Second, // Прервет запрос через 10 секунд
	}
	// Создаем HTTP-запрос
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("Ошибка при создании запроса: %v", err)
	}

	// Добавляем заголовки
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 "+
		"(KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	req.Header.Set("Accept-Language", "en-US,en;q=0.5")
	req.Header.Set("Referer", "https://ru.investing.com/")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Cache-Control", "max-age=0")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-User", "?1")
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Host", "ru.investing.com")
	req.Header.Set("DNT", "1")
	req.Header.Set("TE", "Trailers")

	// Выполняем запрос
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Ошибка при выполнении запроса: %v", err)
	}
	defer resp.Body.Close()

	// Проверяем статус ответа
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Ошибка: статус код %d", resp.StatusCode)
	}

	// Парсим HTML-страницу
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatalf("Ошибка при парсинге HTML: %v", err)
	}

	// Ищем элемент с указанным классом
	element := doc.Find(selector).First()

	// Проверяем, найден ли элемент
	if element.Length() == 0 {
		log.Fatal("Элемент не найден")
	}

	// Извлекаем текст из элемента
	//text := element.Text()
	return element.Text(), nil
}
