package calculations

import (
	"encoding/json"
	"github.com/aville22/greeneats/internal/models"
)

func CalculateCalories(profile models.ProfileForm) ([]byte, error) {
	var bmr float64
	type Result struct {
		TotalCalories float64 `json:"total_calories"`
		ProteinGrams  float64 `json:"protein_grams"`
		FatGrams      float64 `json:"fat_grams"`
		CarbGrams     float64 `json:"carb_grams"`
	}
	// Рассчитываем базовый метаболический коэффициент (BMR)
	if profile.Gender == "male" {
		bmr = 88.362 + (13.397 * profile.Weight) + (4.799 * profile.Height) - (5.677 * profile.Age)
	} else if profile.Gender == "female" {
		bmr = 447.593 + (9.247 * profile.Weight) + (3.098 * profile.Height) - (4.330 * profile.Age)
	}

	// Умножаем BMR на коэффициент активности
	totalCalories := bmr * profile.Activity

	// Распределяем калории в зависимости от цели
	switch profile.Goal {
	case "weight_gain":
		// Увеличиваем общее количество калорий для набора массы
		totalCalories *= 1.1
	case "weight_loss":
		// Уменьшаем общее количество калорий для похудения
		totalCalories *= 0.9
		// Дополнительные кейсы для других целей, если необходимо
	}

	// Процентное соотношение белков, жиров и углеводов от общего количества калорий
	proteinPercent := 0.25 // Примерное соотношение белков: 25%
	fatPercent := 0.35     // Примерное соотношение жиров: 35%
	carbPercent := 0.40    // Примерное соотношение углеводов: 40%

	// Рассчитываем количество граммов для каждого макронутриента (1 г белков/углеводов = 4 калории, 1 г жиров = 9 калорий)
	proteinCalories := totalCalories * proteinPercent
	fatCalories := totalCalories * fatPercent
	carbCalories := totalCalories * carbPercent

	proteinGrams := proteinCalories / 4
	fatGrams := fatCalories / 9
	carbGrams := carbCalories / 4

	result := Result{
		TotalCalories: totalCalories,
		ProteinGrams:  proteinGrams,
		FatGrams:      fatGrams,
		CarbGrams:     carbGrams,
	}

	// Преобразуем результаты в JSON
	jsonData, err := json.Marshal(result)
	if err != nil {
		return nil, err
	}

	return jsonData, nil
}
