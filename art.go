package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

// Функция для выполнения git-коммита
func gitCommit(commitTime time.Time, message string) error {
	// Устанавливаем дату коммита через переменные окружения
	gitCommitTime := commitTime.Format(time.RFC3339)
	os.Setenv("GIT_COMMITTER_DATE", gitCommitTime)
	os.Setenv("GIT_AUTHOR_DATE", gitCommitTime)

	// Выполняем пустой git-коммит
	cmd := exec.Command("git", "commit", "--allow-empty", "-m", message)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("error committing: %v", err)
	}

	return nil
}

// Функция для пуша изменений
func gitPush() error {
	cmd := exec.Command("git", "push")
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("error pushing changes: %v", err)
	}
	return nil
}

// Функция для генерации случайного числа коммитов
func randomCommitsPerDay() int {
	return rand.Intn(5) + 1 // Количество коммитов от 1 до 5
}

func main() {
	rand.Seed(time.Now().UnixNano()) // Инициализация генератора случайных чисел

	startDate := time.Date(2022, 1, 1, 12, 0, 0, 0, time.UTC) // Начало года
	endDate := time.Date(2022, 12, 31, 12, 0, 0, 0, time.UTC) // Конец года

	// Перебираем дни в году
	for currentDate := startDate; currentDate.Before(endDate) || currentDate.Equal(endDate); currentDate = currentDate.AddDate(0, 0, 1) {
		weekday := currentDate.Weekday()

		// Пропускаем субботу и воскресенье
		if weekday == time.Saturday || weekday == time.Sunday {
			continue
		}

		// Рандомно выбираем количество коммитов для этого дня
		commits := randomCommitsPerDay()

		for i := 0; i < commits; i++ {
			// Делаем коммит
			err := gitCommit(currentDate, fmt.Sprintf("Random commit #%d on %s", i+1, currentDate.Format("2006-01-02")))
			if err != nil {
				fmt.Printf("Failed to commit: %v\n", err)
				return
			}
		}
	}

	// Пушим все изменения
	err := gitPush()
	if err != nil {
		fmt.Printf("Failed to push changes: %v\n", err)
		return
	}

	fmt.Println("Random commits for 2022 created successfully!")
}

