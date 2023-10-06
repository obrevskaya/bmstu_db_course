package view

import (
	openapi "console/internal/client"
	"console/internal/consts"
	"fmt"
	"strconv"
)

func InputUser() (*openapi.RegisterRequest, error) {
	user := &openapi.RegisterRequest{}

	fmt.Print("Введите логин: ")
	if _, err := fmt.Scan(&user.Login); err != nil {
		return nil, fmt.Errorf("input login: %w", err)
	}

	fmt.Print("Введите пароль: ")
	if _, err := fmt.Scan(&user.Password); err != nil {
		return nil, fmt.Errorf("input password: %w", err)
	}

	fmt.Print("Введите имя: ")
	if _, err := fmt.Scan(&user.Fio); err != nil {
		return nil, fmt.Errorf("input name: %w", err)
	}

	fmt.Print("Введите электронную почту: ")
	if _, err := fmt.Scan(&user.Email); err != nil {
		return nil, fmt.Errorf("input email: %w", err)
	}

	user.Points = nil
	user.Status = strconv.Itoa(consts.Customer)

	return user, nil
}

func InputAuth() (*openapi.AuthRequest, error) {
	auth := &openapi.AuthRequest{}

	fmt.Print("Введите логин: ")
	if _, err := fmt.Scan(&auth.Login); err != nil {
		return nil, fmt.Errorf("input login: %w", err)
	}

	fmt.Print("Введите пароль: ")
	if _, err := fmt.Scan(&auth.Password); err != nil {
		return nil, fmt.Errorf("input password: %w", err)
	}

	return auth, nil
}

func InputID() (string, error) {
	var idString string
	fmt.Print("Введите ID: ")
	if _, err := fmt.Scan(&idString); err != nil {
		return "", fmt.Errorf("input id order: %w", err)
	}
	return idString, nil
}

func InputIsPoints() (string, error) {
	var s string
	fmt.Print("Использовать баллы? Введите yes или no: ")
	if _, err := fmt.Scan(&s); err != nil {
		return "", fmt.Errorf("input id order: %w", err)
	}
	if s == "yes" || s == "y" {
		return "true", nil
	}

	return "false", nil
}

func InputLimitSkip() (*openapi.GetWinesRequest, error) {
	getWines := &openapi.GetWinesRequest{}

	fmt.Print("Сколько вывести: ")
	if _, err := fmt.Scan(&getWines.Limit); err != nil {
		return nil, fmt.Errorf("input limit: %w", err)
	}

	fmt.Print("Сколько пропустить: ")
	if _, err := fmt.Scan(&getWines.Skip); err != nil {
		return nil, fmt.Errorf("input skip: %w", err)
	}

	return getWines, nil
}

func InputElem() (*openapi.CreateElemRequest, error) {
	elem := &openapi.CreateElemRequest{}

	fmt.Print("Введите id вина: ")
	if _, err := fmt.Scan(&elem.IdWine); err != nil {
		return nil, fmt.Errorf("input id wine: %w", err)
	}

	fmt.Print("Введите количество: ")
	if _, err := fmt.Scan(&elem.Count); err != nil {
		return nil, fmt.Errorf("input count: %w", err)
	}

	return elem, nil
}

func InputWine() (*openapi.AddWineRequest, error) {
	w := &openapi.AddWineRequest{}

	fmt.Print("Введите название вина: ")
	if _, err := fmt.Scan(&w.Name); err != nil {
		return nil, fmt.Errorf("input name wine: %w", err)
	}

	fmt.Print("Введите количество: ")
	if _, err := fmt.Scan(&w.Count); err != nil {
		return nil, fmt.Errorf("input count: %w", err)
	}

	fmt.Print("Введите год: ")
	if _, err := fmt.Scan(&w.Year); err != nil {
		return nil, fmt.Errorf("input year: %w", err)
	}

	fmt.Print("Введите крепость: ")
	if _, err := fmt.Scan(&w.Strength); err != nil {
		return nil, fmt.Errorf("input strength: %w", err)
	}

	fmt.Print("Введите цену: ")
	if _, err := fmt.Scan(&w.Price); err != nil {
		return nil, fmt.Errorf("input price: %w", err)
	}

	fmt.Print("Введите сорт: ")
	if _, err := fmt.Scan(&w.Type); err != nil {
		return nil, fmt.Errorf("input type: %w", err)
	}

	return w, nil
}

func UpdateWine(wine *openapi.Wine) (*openapi.Wine, error) {
	w := &openapi.Wine{
		Id:       wine.Id,
		Name:     wine.Name,
		Count:    wine.Count,
		Year:     wine.Year,
		Strength: wine.Strength,
		Price:    wine.Price,
		Type:     wine.Type,
	}
	var flag string
	fmt.Printf("Изменить название вина (Старое: %s) (yes или no): ", wine.Name)

	if _, err := fmt.Scan(&flag); err != nil {
		return nil, fmt.Errorf("input flag name: %w", err)
	}
	if flag == "yes" || flag == "y" {
		fmt.Printf("Введите названия вина: ")
		if _, err := fmt.Scan(&w.Name); err != nil {
			return nil, fmt.Errorf("input name wine: %w", err)
		}
	}

	fmt.Printf("Изменить количество вина (Старое: %s) (yes или no): ", wine.Count)
	if _, err := fmt.Scan(&flag); err != nil {
		return nil, fmt.Errorf("input flag count: %w", err)
	}
	if flag == "yes" || flag == "y" {
		fmt.Printf("Введите количество вина: ")
		if _, err := fmt.Scan(&w.Count); err != nil {
			return nil, fmt.Errorf("input count wine: %w", err)
		}
	}

	fmt.Printf("Изменить год вина (Старый: %d) (yes или no): ", wine.Year)
	if _, err := fmt.Scan(&flag); err != nil {
		return nil, fmt.Errorf("input flag year: %w", err)
	}
	if flag == "yes" || flag == "y" {
		fmt.Printf("Введите год вина: ")
		if _, err := fmt.Scan(&w.Year); err != nil {
			return nil, fmt.Errorf("input year wine: %w", err)
		}
	}

	fmt.Printf("Изменить крепость вина (Старое: %d) (yes или no):", wine.Strength)
	if _, err := fmt.Scan(&flag); err != nil {
		return nil, fmt.Errorf("input flag strength: %w", err)
	}
	if flag == "yes" || flag == "y" {
		fmt.Printf("Введите крепость вина: ")
		if _, err := fmt.Scan(&w.Strength); err != nil {
			return nil, fmt.Errorf("input strength wine: %w", err)
		}
	}

	fmt.Printf("Изменить цену вина (Старая: %s) (yes или no): ", wine.Price)
	if _, err := fmt.Scan(&flag); err != nil {
		return nil, fmt.Errorf("input flag price: %w", err)
	}
	if flag == "yes" || flag == "y" {
		fmt.Printf("Введите цену вина: ")
		if _, err := fmt.Scan(&w.Price); err != nil {
			return nil, fmt.Errorf("input price wine: %w", err)
		}
	}

	fmt.Printf("Изменить сорт вина (Старый: %s) (yes или no):: ", wine.Type)
	if _, err := fmt.Scan(&flag); err != nil {
		return nil, fmt.Errorf("input flag type: %w", err)
	}
	if flag == "yes" || flag == "y" {
		fmt.Printf("Введите сорт вина: ")
		if _, err := fmt.Scan(&w.Type); err != nil {
			return nil, fmt.Errorf("input type wine: %w", err)
		}
	}
	return w, nil
}
