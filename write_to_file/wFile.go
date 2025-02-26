package write_to_file

import (
	"fmt"
	"os"
)

func WriteToFile(data []byte) (string, error) {

	file, err := os.OpenFile("exchangeRate.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		return "", err
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		fmt.Println("Ошибка при записи файла:", err)
		return "", err
	}

	return "Данные успешно добавлены", nil

}
