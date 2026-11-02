package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type User struct {
	Name          string
	Phone         string
	Books_on_hand []Book
	Penalties     []Penalty
}

type Book struct {
	Name                string
	Author              string
	Year_of_manufacture int
}

type Penalty struct {
	Name  string
	Price int
}

func main() {
	var users []User
	var books []Book
	new_user := User{}
	new_penalty := Penalty{}
	reader := bufio.NewReader(os.Stdin)
	for {
		var choice int

		fmt.Println("Меню программы библиотеки.")
		fmt.Println("1. Добавить человека.")
		fmt.Println("2. Убрать человека.")
		fmt.Println("3. Добавить штраф человеку.")
		fmt.Println("4. Добавить книгу.")
		fmt.Println("5. Узнать всё о человеке по читательскому билету.")

		fmt.Print("Выберите действие: ")
		fmt.Scan(&choice)
		fmt.Println("")

		if choice == 1 {

			fmt.Print("Укажите имя человека: ")
			fmt.Scan(&new_user.Name)

			fmt.Println("")

			fmt.Print("Укажите номер телефона: ")
			fmt.Scan(&new_user.Phone)

			users = append(users, new_user)
			fmt.Println("Человек успешно добавлен!")
		} else if choice == 2 {
			var name string

			fmt.Print("Введите имя человека: ")
			fmt.Scan(&name)
			fmt.Println("")

			found := false

			for i, user := range users {
				if user.Name == name {
					users = append(users[:i], users[i+1:]...)
					fmt.Println("Успешное удаление!")
					fmt.Println("")
					found = true
					break
				}
			}

			if !found {
				fmt.Println("Не удалось найти человека!")
			}
		} else if choice == 3 {
			var name string

			fmt.Print("Введите имя человека: ")
			fmt.Scan(&name)
			fmt.Println("")
			fmt.Scanln()

			found := false

			for i, user := range users {
				if user.Name == name {

					fmt.Print("Укажите название штрафа: ")
					new_penalty.Name, _ = reader.ReadString('\n')
					new_penalty.Name = strings.TrimSpace(new_penalty.Name)
					fmt.Println("")

					fmt.Print("Укажите цену штрафа: ")
					fmt.Scan(&new_penalty.Price)
					fmt.Println("")

					users[i].Penalties = append(users[i].Penalties, new_penalty)
					fmt.Println("Штраф успешно добавлен!")
					fmt.Println("")
					found = true
					break
				}
			}

			if !found {
				fmt.Println("Не удалось найти человека!")
			}
		} else if choice == 5 {
			var name string

			fmt.Print("Введите имя человека: ")
			fmt.Scan(&name)
			fmt.Println("")

			found := false

			if new_user.Name == name {
				found = true
				for userIndex := range users {
					fmt.Println("Имя:", users[userIndex].Name)
					fmt.Println("Номер телефона:", users[userIndex].Phone)
					fmt.Println("Штрафы:", len(users[userIndex].Penalties))
					if len(users[userIndex].Penalties) > 0 {
						var choice string

						fmt.Println("У человека обнаружены штрафы!")
						fmt.Print("Вывести на экран? (Y/n):")
						fmt.Scan(&choice)
						fmt.Println("")

						if strings.ToLower(choice) == "y" {
							for i, penalty := range users[userIndex].Penalties {
								var choice string
								fmt.Println("Штраф", i+1, "-", penalty.Name)
								fmt.Println("Цена:", penalty.Price)

								fmt.Print("Вы хотите снять штраф? (Y/n): ")
								fmt.Scan(&choice)
								fmt.Scanln()
								fmt.Println("")

								if strings.ToLower(choice) == "y" {
									var penalty_name string

									fmt.Print("Введите имя штрафа: ")
									penalty_name = readText(reader)
									fmt.Println("")

									penaltyFound := false

									for i, penalty := range users[userIndex].Penalties {
										if penalty.Name == penalty_name {
											users[userIndex].Penalties = append(
												users[userIndex].Penalties[:i],
												users[userIndex].Penalties[i+1:]...,
											)
											penaltyFound = true
											fmt.Println("Штраф успешно удалён!")
											fmt.Println("")
											break
										}
									}
									if !penaltyFound {
										fmt.Println("Такой штраф не найден!")
									}
								} else {
									fmt.Println("Штрафы не были сняты!")
									break
								}
								break
							}
						}
					}
					if !found {
						fmt.Println("Человек не найден!")
					}
				}
			}
		} else if choice == 4 {
			var name string

			fmt.Print("Введите имя человека: ")
			new_user.Name = readText(reader)
			fmt.Println("")

			found := false

			for i := range users {
				if users[i].Name == name {

					fmt.Print("Введите название книги: ")
					books[i].Name = readText(reader)
					fmt.Println("")

					fmt.Print("Введите автора книги: ")
					books[i].Author = readText(reader)
					fmt.Println("")

					fmt.Print("Введите год издания книги: ")
					fmt.Scan(&books[i].Year_of_manufacture)
					fmt.Println("")

					users[i].Books_on_hand = append(
						users[i].Books_on_hand,
						books...,
					)

					fmt.Println("Книга успешно добавлена!")
					break
				}
			}
			if !found {
				fmt.Println("Книга не найдена!")
			}
		} else {
			fmt.Println("Неизвестное действие!")
		}
	}
}

func readText(reader *bufio.Reader) string {
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}
