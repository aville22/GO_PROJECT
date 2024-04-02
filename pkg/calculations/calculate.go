package calculations

import "encoding/json"

func CalculateCalories(weight, height, age float64, gender, goal string, activityLevel float64) ([]byte, error) {
	var bmr float64
	type Result struct {
		TotalCalories float64 `json:"total_calories"`
		ProteinGrams  float64 `json:"protein_grams"`
		FatGrams      float64 `json:"fat_grams"`
		CarbGrams     float64 `json:"carb_grams"`
	}
	// Рассчитываем базовый метаболический коэффициент (BMR)
	if gender == "male" {
		bmr = 88.362 + (13.397 * weight) + (4.799 * height) - (5.677 * age)
	} else if gender == "female" {
		bmr = 447.593 + (9.247 * weight) + (3.098 * height) - (4.330 * age)
	}

	// Умножаем BMR на коэффициент активности
	totalCalories := bmr * activityLevel

	// Распределяем калории в зависимости от цели
	switch goal {
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
