package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Bank struct {
	Name   string
	Prefix string
}

func DetectBank(cardNumber string, banks []Bank) *Bank {
	if len(cardNumber) == 0 {
		return nil
	}

	for i := range banks {
		if strings.HasPrefix(cardNumber, banks[i].Prefix) {
			return &banks[i]
		}
	}

	return nil
}

func LuhnCheck(cardNumber string) bool {
	if len(cardNumber) == 0 {
		return false
	}

	total := 0
	isSecond := false
	for i := len(cardNumber) - 1; i >= 0; i-- {
		digit := int(cardNumber[i] - '0')
		if digit < 0 || digit > 9 {
			return false
		}

		if isSecond {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}

		total += digit
		isSecond = !isSecond
	}

	return total%10 == 0
}

func loadBankData(path string) ([]Bank, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var banks []Bank
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}

		parts := strings.Split(line, ",")
		if len(parts) != 2 {
			return nil, fmt.Errorf("неверный формат строки: %q", line)
		}

		banks = append(banks, Bank{Name: strings.TrimSpace(parts[0]), Prefix: strings.TrimSpace(parts[1])})
	}

	if scanner.Err() != nil {
		return nil, err
	}

	return banks, nil
}

func getUserInput() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите номер карты (или Enter для выхода): ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	input = strings.ReplaceAll(input, " ", "")
	input = strings.ReplaceAll(input, "-", "")
	return input
}

func validateInput(cardNumber string) error {
	for _, char := range cardNumber {
		digit := char - '0'
		if digit < 0 || digit > 9 {
			return fmt.Errorf("номер должен содержать только цифры")
		}
	}

	minLen, maxLen := 13, 19
	if len(cardNumber) < minLen || len(cardNumber) > maxLen {
		return fmt.Errorf("номер должен содержать от %d до %d цифр", minLen, maxLen)
	}

	return nil
}

func main() {
	banks, err := loadBankData("./banks.txt")
	if err != nil {
		fmt.Println("Не удалось загрузить банки:", err)
		return
	}

	fmt.Println("Загружено банков:", len(banks))

	for {
		cardNumber := getUserInput()
		if len(cardNumber) == 0 {
			fmt.Println("До свидания!")
			break
		}

		err = validateInput(cardNumber)
		if err != nil {
			fmt.Println("Ошибка:", err)
			continue
		}

		if !LuhnCheck(cardNumber) {
			fmt.Println("Ошибка: не прошел проверку по Луна")
			continue
		}

		detectedBank := DetectBank(cardNumber, banks)
		bankName := "не определен"
		if detectedBank != nil {
			bankName = detectedBank.Name
		}
		fmt.Println("Банк:", bankName)
	}
}
