package spentcalories

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

// parseTraining — Распаковывает полученные шаги, тип тренировки и время из строки.
func parseTraining(data string) (int, string, time.Duration, error) {
	s := strings.Split(data, ",")
	if len(s) != 3 {
		return 0, "", 0, fmt.Errorf("неверное количество элементов: ожидалось 3, получено %d", len(s))
	}
	step, err := strconv.Atoi(s[0])
	if err != nil {
		return 0, "", 0, err
	}
	traintype := s[1]
	dur, err := time.ParseDuration(s[2])
	if err != nil {
		return 0, "", 0, err
	}
	if step <= 0 || dur <= 0 {
		return 0, "", 0, fmt.Errorf("шаги и длительность должны быть положительными")
	}
	return step, traintype, dur, nil
}

// distance — Рассчитывает дистанцию в километрах.
func distance(steps int, height float64) float64 {
	longStep := height * stepLengthCoefficient / mInKm
	return float64(steps) * longStep
}

// meanSpeed — Рассчитывает среднюю скорость.
func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 || steps <= 0 || height <= 0 {
		return 0
	}
	longDistance := distance(steps, height)
	return longDistance / duration.Hours()
}

// RunningSpentCalories — Рассчитывает сожженые калории при беге.
func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, fmt.Errorf("все параметры должны быть положительными")
	}
	speed := meanSpeed(steps, height, duration)
	mins := duration.Minutes()
	kkal := (weight * speed * mins) / minInH
	return kkal, nil

}

// WalkingSpentCalories — Рассчитывает сожженые калории при ходьбе.
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, fmt.Errorf("все параметры должны быть положительными")
	}
	speed := meanSpeed(steps, height, duration)
	mins := duration.Minutes()
	kkal := ((weight * speed * mins) / minInH) * walkingCaloriesCoefficient
	return kkal, nil
}

// TrainingInfo — Общая функция, показывает результат.
func TrainingInfo(data string, weight, height float64) (string, error) {
	step, traintype, timeDur, err := parseTraining(data)
	if err != nil {
		log.Println(err)
		return "", err
	}
	var dist, awrSpeed, kkal float64
	switch traintype {
	case "Бег":
		kkal, err = RunningSpentCalories(step, weight, height, timeDur)
		if err != nil {
			return "", err
		}
	case "Ходьба":
		kkal, err = WalkingSpentCalories(step, weight, height, timeDur)
		if err != nil {
			return "", err
		}
	default:
		return "", fmt.Errorf("неизвестный тип тренировки")
	}
	dist = distance(step, height)
	awrSpeed = meanSpeed(step, height, timeDur)
	// Переменная для вывода
	announcement := fmt.Sprintf(
		`Тип тренировки: %s
Длительность: %.2f ч.
Дистанция: %.2f км.
Скорость: %.2f км/ч
Сожгли калорий: %.2f
`,
		traintype, timeDur.Hours(), dist, awrSpeed, kkal,
	)
	return announcement, nil
}
