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
	new_user := User{}
	new_penalty := Penalty{}
	reader := bufio.NewReader(os.Stdin)
	all_books := []Book{
		{"Преступление и наказание", "Ф.М. Достоевский", 1866},
		{"Война и мир", "Л.Н. Толстой", 1869},
		{"Евгений Онегин", "А.С. Пушкин", 1833},
		{"Мёртвые души", "Н.В. Гоголь", 1842},
		{"Отцы и дети", "И.С. Тургенев", 1862},
		{"Мастер и Маргарита", "М.А. Булгаков", 1967},
		{"Анна Каренина", "Л.Н. Толстой", 1877},
		{"Идиот", "Ф.М. Достоевский", 1869},
		{"Герой нашего времени", "М.Ю. Лермонтов", 1840},
	}
	for {
		var choice int

		fmt.Println("Меню программы библиотеки.")
		fmt.Println("1. Добавить человека.")
		fmt.Println("2. Убрать человека.")
		fmt.Println("3. Добавить штраф человеку.")
		fmt.Println("4. Добавить книгу у человека.")
		fmt.Println("5. Узнать всё о человеке по читательскому билету.")
		fmt.Println("6. Посмотреть реестр книг.")
		fmt.Println("0. Выход")

		fmt.Print("Выберите действие: ")
		fmt.Scan(&choice)
		fmt.Scanln()
		fmt.Println("")

		if choice == 1 {

			fmt.Print("Укажите имя человека: ")
			new_user.Name = readText(reader)

			if new_user.Name == "" {
				fmt.Println("Имя не может быть пустой строкой!")
				continue
			}

			fmt.Println("")

			fmt.Print("Укажите номер телефона: ")
			new_user.Phone = readText(reader)

			if new_user.Phone == "" {
				fmt.Println("Телефон не может быть пустой строкой!")
				continue
			}
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
		} else if choice == 4 {
			var name string

			fmt.Print("Введите имя человека: ")
			name = readText(reader)
			fmt.Println("")

			found := false

			for i := range users {
				if users[i].Name == name {
					var newBook Book

					fmt.Print("Введите название книги: ")
					newBook.Name = readText(reader)
					fmt.Println("")

					users[i].Books_on_hand = append(
						users[i].Books_on_hand,
						newBook,
					)
					found = true
					fmt.Println("Книга успешно добавлена!")
					break
				}
			}
			if !found {
				fmt.Println("Книга не найдена!")
			}
		} else if choice == 5 {
			var name string
			var num_or_name string

			fmt.Println("1. Найти человека по имени.")
			fmt.Println("2. Найти человека по номеру телефона.")
			fmt.Println("")

			fmt.Print("Выберите действие: ")
			num_or_name = readText(reader)

			found := false

			if num_or_name == "1" {
				fmt.Print("Введите имя человека: ")
				name = readText(reader)
				fmt.Println("")

				for userIndex := range users {
					if users[userIndex].Name == name {
						found = true

						fmt.Println("Имя:", users[userIndex].Name)
						fmt.Println("Номер телефона:", users[userIndex].Phone)
						fmt.Println("Штрафы:", len(users[userIndex].Penalties))
						fmt.Println("Количество книг на руках:", len(users[userIndex].Books_on_hand))
						fmt.Println("")

						manageUser(reader, &users[userIndex])
						break
					}
				}

				if !found {
					fmt.Println("Человек не найден!")
				}

			} else if num_or_name == "2" {
				var phone string

				fmt.Print("Введите номер телефона: ")
				phone = readText(reader)
				fmt.Println("")

				for userIndex := range users {
					if users[userIndex].Phone == phone {
						found = true

						fmt.Println("Имя:", users[userIndex].Name)
						fmt.Println("Телефон:", users[userIndex].Phone)
						fmt.Println("Штрафы:", len(users[userIndex].Penalties))
						fmt.Println("Количество книг на руках:", len(users[userIndex].Books_on_hand))
						fmt.Println("")

						manageUser(reader, &users[userIndex])
						break
					}
				}

				if !found {
					fmt.Println("Такого номера телефона не найдено!")
				}
			}

		} else if choice == 6 {
			fmt.Println("Список доступных книг:")
			fmt.Println("")

			for i := 0; i < len(all_books); i++ {
				fmt.Println("Название:", all_books[i].Name)
				fmt.Println("Автор:", all_books[i].Author)
				fmt.Println("Год выпуска:", all_books[i].Year_of_manufacture)
				fmt.Println("")
			}
			var choice123 string

			fmt.Println("1. Добавить книгу в реестр.")
			fmt.Println("2. Убрать книгу с реестра.")
			fmt.Println("0. Выход")
			fmt.Println("")

			fmt.Print("Выберите действие: ")
			choice123 = readText(reader)

			if choice123 == "1" {
				var name string
				var author string
				var year int

				fmt.Print("Введите название книги: ")
				name = readText(reader)
				fmt.Println("")

				if name == "" {
					fmt.Println("Название не может быть пустым!")
					continue
				}

				fmt.Print("Введите автора книги: ")
				author = readText(reader)
				fmt.Println("")

				if author == "" {
					fmt.Println("Имя автора не может быть пустым!")
					continue
				}

				fmt.Print("Введите год выпуска: ")
				fmt.Scan(&year)
				fmt.Scanln()
				fmt.Println("")

				all_books = append(all_books, Book{
					Name:                name,
					Author:              author,
					Year_of_manufacture: year,
				})
				fmt.Println("Книга успешно добавлена!")
			} else if choice123 == "2" {
				var book_name string

				fmt.Print("Введите имя книги: ")
				book_name = readText(reader)
				fmt.Println("")

				bookFound := false

				for i := range all_books {
					if all_books[i].Name == book_name {
						all_books = append(
							all_books[:i],
							all_books[i+1:]...,
						)

						bookFound = true
						fmt.Println("Книга успешно удалена с реестра!")
						break
					}
				}
				if !bookFound {
					fmt.Println("Книга не найдена!")
					continue
				}

			} else if choice123 == "0" {
				fmt.Println("Ничего не было изменено!")
			}
		} else if choice == 0 {
			fmt.Println("До свидания!")
			break
		}
	}
}

func readText(reader *bufio.Reader) string {
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func manageUser(reader *bufio.Reader, user *User) {
	for {
		fmt.Println("\n1. Действия с книгами")
		fmt.Println("2. Действия со штрафами")
		fmt.Println("0. Назад")

		fmt.Print("Выберите действие: ")
		choice := readText(reader)

		switch choice {
		case "1":
			manageBooks(reader, user)

		case "2":
			managePenalties(reader, user)

		case "0":
			return

		default:
			fmt.Println("Неверный выбор!")
		}
	}
}

func manageBooks(reader *bufio.Reader, user *User) {
	for {
		fmt.Println("\nКниги человека:")

		if len(user.Books_on_hand) == 0 {
			fmt.Println("У человека нет книг на руках!")
			return
		}

		for i, book := range user.Books_on_hand {
			fmt.Println("Книга #", i+1)
			fmt.Println("Название:", book.Name)
			fmt.Println("Автор:", book.Author)
			fmt.Println("Год выпуска:", book.Year_of_manufacture)
			fmt.Println("")
		}

		fmt.Println("1. Убрать книгу")
		fmt.Println("0. Назад")

		fmt.Print("Выберите действие: ")
		choice := readText(reader)

		if choice == "1" {
			var name_book string

			fmt.Print("Введите название книги: ")
			name_book = readText(reader)

			foundBook := false

			for i, book := range user.Books_on_hand {
				if book.Name == name_book {
					user.Books_on_hand = append(
						user.Books_on_hand[:i],
						user.Books_on_hand[i+1:]...,
					)

					fmt.Println("Книга успешно удалена!")
					foundBook = true
					break
				}
			}

			if !foundBook {
				fmt.Println("Книга не найдена!")
			}

		} else if choice == "0" {
			return
		} else {
			fmt.Println("Неверный выбор!")
		}
	}
}

func managePenalties(reader *bufio.Reader, user *User) {
	for {
		fmt.Println("\nШтрафы человека:")

		if len(user.Penalties) == 0 {
			fmt.Println("У человека нет штрафов!")
			return
		}

		for i, penalty := range user.Penalties {
			fmt.Println("Штраф #", i+1)
			fmt.Println("Название:", penalty.Name)
			fmt.Println("Цена:", penalty.Price, "рублей")
			fmt.Println("")
		}

		fmt.Println("1. Снять штраф")
		fmt.Println("0. Назад")

		fmt.Print("Выберите действие: ")
		choice := readText(reader)

		if choice == "1" {
			var penalty_name string

			fmt.Print("Введите название штрафа: ")
			penalty_name = readText(reader)

			penaltyFound := false

			for i, penalty := range user.Penalties {
				if penalty.Name == penalty_name {
					user.Penalties = append(
						user.Penalties[:i],
						user.Penalties[i+1:]...,
					)

					fmt.Println("Штраф успешно удалён!")
					penaltyFound = true
					break
				}
			}

			if !penaltyFound {
				fmt.Println("Такой штраф не найден!")
			}

		} else if choice == "0" {
			return
		} else {
			fmt.Println("Неверный выбор!")
		}
	}
}
