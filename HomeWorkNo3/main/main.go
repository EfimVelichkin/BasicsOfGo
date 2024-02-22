package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/eiannone/keyboard"
)

type URL struct {
	Name        string
	URL         string
	Description string
	Tags        []string
	Date        string
}

var urlList []URL

func main() {
	defer func() {
		_ = keyboard.Close()
	}()

	fmt.Println("Программа для добавления url в список")
	fmt.Println("Для выхода из приложения нажмите Esc")

OuterLoop:
	for {
		if err := keyboard.Open(); err != nil {
			log.Fatal(err)
		}

		char, key, err := keyboard.GetKey()
		if err != nil {
			log.Fatal(err)
		}

		switch char {
		case 'a':
			if err := keyboard.Close(); err != nil {
				log.Fatal(err)
			}

			fmt.Println("Введите новую запись в формате <url описание теги>")

			reader := bufio.NewReader(os.Stdin)
			text, _ := reader.ReadString('\n')
			args := strings.Fields(text)
			if len(args) < 3 {
				fmt.Println("Введите правильные аргументы в формате url описание теги")
				continue OuterLoop
			}

			newURL := URL{
				Name:        args[0],
				URL:         args[1],
				Description: args[2],
				Tags:        args[3:],
				Date:        "сегодня",
			}

			urlList = append(urlList, newURL)
			fmt.Println("URL успешно добавлен")

		case 'l':
			fmt.Printf("Всего добавлено URL: %d\n", len(urlList))
			for _, url := range urlList {
				fmt.Printf("Имя: %s\nURL: %s\nОписание: %s\nТеги: %s\nДата: %s\n\n", url.Name, url.URL, url.Description, strings.Join(url.Tags, ", "), url.Date)
			}

		case 'r':
			if err := keyboard.Close(); err != nil {
				log.Fatal(err)
			}

			fmt.Println("Введите имя ссылки, которую нужно удалить")
			reader := bufio.NewReader(os.Stdin)
			text, _ := reader.ReadString('\n')
			name := strings.TrimSpace(text)

			for i, url := range urlList {
				if url.Name == name {
					urlList = append(urlList[:i], urlList[i+1:]...)
					fmt.Printf("URL с именем '%s' успешно удален\n", name)
					break
				}
			}

		default:
			if key == keyboard.KeyEsc {
				return
			}
		}
	}
}
