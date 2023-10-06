package main

import (
	"context"
	"fmt"
	myContext "git.iu7.bmstu.ru/ovv20u676/ppo/src/backend/internal/context"
	controllers2 "git.iu7.bmstu.ru/ovv20u676/ppo/src/backend/internal/logic/controllers"
	"git.iu7.bmstu.ru/ovv20u676/ppo/src/backend/internal/logic/models"
	postgres2 "git.iu7.bmstu.ru/ovv20u676/ppo/src/backend/internal/repository/postgres"
	"github.com/google/uuid"
	"github.com/testcontainers/testcontainers-go"
	"gorm.io/gorm"
	"os"
	"strconv"
)

const N = 500
const M = 10000

func main() {
	step := 0

	file, err := os.Create("result.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	for i := 10; i <= M; i += step {
		fmt.Println("Records count: ", i)

		resultTimeTr, errorCountTr, err := researchCreateRecordWithTrigger(i)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Research with trigger end ok!")
		}

		resultTime, errorCount, err := researchCreateRecordWithoutTrigger(i)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Research without trigger end ok!")
		}

		_, err = file.WriteString(strconv.Itoa(i) + " " + strconv.Itoa(resultTimeTr) + " " + strconv.Itoa(errorCountTr) + " ")
		if err != nil {
			fmt.Println(err)
			return
		}

		_, err = file.WriteString(strconv.Itoa(resultTime) + " " + strconv.Itoa(errorCount) + "\n")
		if err != nil {
			fmt.Println(err)
			return
		}

		if i < 100 {
			i += 10
		} else if i == 100 {
			step = 100
		} else if i == 1000 {
			step = 1000
		}
	}
}

func setupData(count int) (error, *gorm.DB, testcontainers.Container) {
	dbContainer, db, err := SetupTestDatabase()
	if err != nil {
		return err, nil, nil
	}

	text, err := os.ReadFile("./deployments/res/" + strconv.Itoa(count) + ".sql")
	if err != nil {
		return fmt.Errorf("read file: %w", err), nil, nil
	}

	if err := db.Exec(string(text)).Error; err != nil {
		return fmt.Errorf("exec: %w", err), nil, nil
	}

	return nil, db, dbContainer
}

func setupDataTr(count int) (error, *gorm.DB, testcontainers.Container) {
	dbContainer, db, err := SetupTestDatabase()
	if err != nil {
		return err, nil, nil
	}

	//text, err := os.ReadFile("./deployments/res/tr.sql")
	//if err != nil {
	//	return fmt.Errorf("read file: %w", err), nil, nil
	//}
	//
	//if err := db.Exec(string(text)).Error; err != nil {
	//	return fmt.Errorf("exec: %w", err), nil, nil
	//}

	text, err := os.ReadFile("./deployments/res/" + strconv.Itoa(count) + ".sql")
	if err != nil {
		return fmt.Errorf("read file: %w", err), nil, nil
	}

	if err := db.Exec(string(text)).Error; err != nil {
		return fmt.Errorf("exec: %w", err), nil, nil
	}

	return nil, db, dbContainer
}

func researchCreateRecordWithoutTrigger(count int) (int, int, error) {

	err, db, dbContainer := setupData(count)
	if err != nil {
		return 0, 0, err
	}
	defer dbContainer.Terminate(context.Background())

	billRep := postgres2.NewBR(db)
	orderRep := postgres2.NewOR(db)
	elemRep := postgres2.NewOElR(db)
	userRep := postgres2.NewUR(db)
	wineRep := postgres2.NewWR(db)

	orderController := controllers2.NewOrderController(billRep, elemRep, userRep, orderRep, wineRep)
	elemController := controllers2.NewElemController(elemRep, orderRep, wineRep, orderController)

	var result int64
	var errorCount int64
	var successCount int64

	for successCount != N {

		fmt.Println(successCount)
		ctx := myContext.UserToContext(context.Background(), &models.User{
			ID:       uuid.MustParse("85684bce-b987-46d3-8d9a-9c83d64b8e65"),
			Login:    "obrevskaya",
			Password: "qwerty1234",
			Fio:      "Obrevskaya Veronika",
			Email:    "obrevskaya.vera@mail.ru",
			Points:   0,
			Status:   models.Customer,
		})
		elem := &models.OrderElement{
			IDWine:  uuid.MustParse("8168077b-f285-4ca2-87b4-735c95ccae5b"),
			IDOrder: uuid.MustParse("5b376dba-89a4-4a83-ab80-433c49b37db6"),
			Count:   22,
		}
		err, duration := elemController.ResearchCreate(ctx, elem)
		if err != nil {
			errorCount += 1
			fmt.Println(err)
		} else {
			successCount += 1
			result += duration.Nanoseconds()
		}

		err = elemController.Delete(ctx, elem.ID)
		if err != nil {
			return 0, 0, err
		}

	}

	fmt.Println("итог время!!!!! ", result/N)
	fmt.Println("итого ошибок!!!!", errorCount)
	return int(result), int(errorCount), err
}

func researchCreateRecordWithTrigger(count int) (int, int, error) {

	err, db, dbContainer := setupDataTr(count)
	if err != nil {
		return 0, 0, err
	}
	defer dbContainer.Terminate(context.Background())

	billRep := postgres2.NewBR(db)
	orderRep := postgres2.NewOR(db)
	elemRep := postgres2.NewOElR(db)
	userRep := postgres2.NewUR(db)
	wineRep := postgres2.NewWR(db)

	orderController := controllers2.NewOrderController(billRep, elemRep, userRep, orderRep, wineRep)
	elemController := controllers2.NewElemController(elemRep, orderRep, wineRep, orderController)

	var result int64
	var errorCount int64
	var successCount int64

	for successCount != N {

		fmt.Println(successCount)
		ctx := myContext.UserToContext(context.Background(), &models.User{
			ID:       uuid.MustParse("85684bce-b987-46d3-8d9a-9c83d64b8e65"),
			Login:    "obrevskaya",
			Password: "qwerty1234",
			Fio:      "Obrevskaya Veronika",
			Email:    "obrevskaya.vera@mail.ru",
			Points:   0,
			Status:   models.Customer,
		})
		elem := &models.OrderElement{
			IDWine:  uuid.MustParse("8168077b-f285-4ca2-87b4-735c95ccae5b"),
			IDOrder: uuid.MustParse("5b376dba-89a4-4a83-ab80-433c49b37db6"),
			Count:   22,
		}
		err, duration := elemController.ResearchCreateTr(ctx, elem)
		if err != nil {
			errorCount += 1
			fmt.Println(err)
		} else {
			successCount += 1
			result += duration.Nanoseconds()
		}

		err = elemController.Delete(ctx, elem.ID)
		if err != nil {
			return 0, 0, err
		}

	}

	fmt.Println("итог время!!!!! ", result/N)
	fmt.Println("итого ошибок!!!!", errorCount)
	return int(result), int(errorCount), err
}
