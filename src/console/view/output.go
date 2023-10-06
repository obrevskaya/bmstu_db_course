package view

import (
	openapi "console/internal/client"
	"fmt"

	"github.com/fatih/color"
)

func PrintRunMenu() {
	fmt.Printf("Выберите как зайти\n" +
		"0) пользователь\n1) админ\nВыбор:\n")
}

func PrintGuestMenu() {
	fmt.Printf("Гостевое меню пользователя\n")
	fmt.Printf("0) Выход\n")
	fmt.Printf("1) Регистрация\n")
	fmt.Printf("2) Авторизация\n")
	fmt.Printf("3) Получить каталог вин\n")
	fmt.Printf("Выбор:\n")
}

func PrintUserMenu() {
	fmt.Printf("Пользовательское меню\n")
	fmt.Printf("0) Выход\n")
	fmt.Printf("1) Разлогиниться\n")
	fmt.Printf("2) Получить каталог вин\n")
	fmt.Printf("3) Добавить вино в заказ\n")
	fmt.Printf("4) Увеличить количество вина в заказе\n")
	fmt.Printf("5) Уменьшить количество вина в заказе\n")
	fmt.Printf("6) Удалить вино из заказа\n")
	fmt.Printf("7) Получить текущий заказ\n")
	fmt.Printf("8) Оформить заказ\n")
	fmt.Printf("9) Добавить вино в любимое\n")
	fmt.Printf("10) Получить список любимых вин\n")
	fmt.Printf("11) Удалить вино из любимого\n")
	fmt.Printf("Выбор:\n")
}

func PrintAdminMenu() {
	fmt.Printf("Меню администратора\n")
	fmt.Printf("0) Выход\n")
	fmt.Printf("1) Регистрация\n")
	fmt.Printf("2) Получить каталог вин\n")
	fmt.Printf("3) Подтвердить оплату счета\n")
	fmt.Printf("4) Добавить вино\n")
	fmt.Printf("5) Удалить вино\n")
	fmt.Printf("6) Изменить вино\n")
	fmt.Printf("Выбор:\n")
}

func PrintWine(wine openapi.Wine) {
	fmt.Printf("\tID: %s\n", wine.Id)
	fmt.Printf("\tНазвание: %s\n", wine.Name)
	fmt.Printf("\tКоличество: %s\n", wine.Count)
	fmt.Printf("\tГод: %d\n", wine.Year)
	fmt.Printf("\tКрепость: %d\n", wine.Strength)
	fmt.Printf("\tЦена: %s\n", wine.Price)
	fmt.Printf("\tСорт: %s\n", wine.Type)
}

func PrintWines(wines []openapi.Wine) {
	c := color.New(color.FgCyan)
	for i, wine := range wines {
		c.Printf("Вино №%d:\n", i+1)
		PrintWine(wine)
	}
}

func PrintUserWines(wines []openapi.UserWine) {
	c := color.New(color.FgCyan)
	for i, wine := range wines {
		c.Printf("Вино №%d:\n", i+1)
		fmt.Println(wine.IdWine)
	}
}

func PrintOrder(o *openapi.Order) {
	c := color.New(color.FgCyan)
	c.Printf("Заказ:\n")
	fmt.Printf("\tID: %s\n", o.Id)
	fmt.Printf("\tID пользователя: %s\n", o.IdUser)
	fmt.Printf("\tОбщая стоимость: %s\n", o.TotalPrice)
	fmt.Printf("\tОплачено ли баллами: %s\n", o.IsPoints)
	fmt.Printf("\tСтатус: %s\n", o.Status)
}

func PrintWineInOrder(w *openapi.Wine, el *openapi.Elem) {
	c := color.New(color.FgCyan)
	c.Printf("Элемент:\n")
	fmt.Printf("\tID элемента: %s\n", el.Id)
	fmt.Printf("\tID вина: %s\n", w.Id)
	fmt.Printf("\tНазвание: %s\n", w.Name)
	fmt.Printf("\tСтоимость вина: %s\n", w.Price)
	fmt.Printf("\tКоличество: %d\n\n", el.Count)
}
