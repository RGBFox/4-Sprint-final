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

func parseTraining(data string) (int, string, time.Duration, error) {
	s := strings.Split(data, ",")
	if len(s) != 3 {
		return 0, "", 0, fmt.Errorf("неправильный формат ввода данных")
	}
	step, err := strconv.Atoi(s[0])
	if err != nil {
		return 0, "", 0, err
	}
	if step < 0 {
		return 0, "", 0, fmt.Errorf("отрицательное количество шагов")
	}
	traintype := s[1]
	time, err := time.ParseDuration(s[3])
	if err != nil {
		return 0, "", 0, err
	}
	return step, traintype, time, nil
}

func distance(steps int, height float64) float64 {
	/*	if steps < 0 || height < 0 {
		return 0
	}*/
	longStep := height * stepLengthCoefficient
	return float64(steps) * longStep
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	/*	if duration <= 0 || steps <= 0 || height <= 0 {
		return 0
	}*/
	longDistance := distance(steps, height)
	return longDistance / duration.Hours()
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, fmt.Errorf("Отрицательные или нулевые числа")
	}
	speed := meanSpeed(steps, height, duration)
	mins := duration.Minutes()
	kkal := (weight * speed * mins) / minInH
	return kkal, nil

}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, fmt.Errorf("Отрицательные или нулевые числа")
	}
	speed := meanSpeed(steps, height, duration)
	mins := duration.Minutes()
	kkal := ((weight * speed * mins) / minInH) * walkingCaloriesCoefficient
	return kkal, nil
}

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
		dist = distance(step, height)
		awrSpeed = meanSpeed(step, height, timeDur)
	case "Ходьба":
		kkal, err = WalkingSpentCalories(step, weight, height, timeDur)
		if err != nil {
			return "", err
		}
		dist = distance(step, height)
		awrSpeed = meanSpeed(step, height, timeDur)
	default:
		return "", fmt.Errorf("неизвестный тип тренировки")
	}

	announcement := fmt.Sprintf(
		`Тип тренировки: %s
Длительность: %.2f ч.
Дистанция: %.2f км.
Скорость: %.2f км/ч
Сожгли калорий: %.2f`,
		traintype, timeDur.Hours(), dist, awrSpeed, kkal,
	)
	return announcement, nil
}
