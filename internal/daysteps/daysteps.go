package daysteps

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/RGBFox/4-Sprint-final/internal/spentcalories"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

// parsePackage — распаковывает полученные шаги и время из строки.
func parsePackage(data string) (int, time.Duration, error) {
	s := strings.Split(data, ",")
	if len(s) != 2 {
		return 0, 0, fmt.Errorf("неверное количество элементов: ожидалось 2, получено %d", len(s))
	}
	step, err := strconv.Atoi(s[0])
	if err != nil {
		return 0, 0, err
	}
	dur, err := time.ParseDuration(s[1])
	if err != nil {
		return 0, 0, err
	}
	if step <= 0 || dur <= 0 {
		return 0, 0, fmt.Errorf("шаги и длительность должны быть положительными")
	}
	return step, dur, nil
}

// DayActionInfo — вычисляет количество шагов, дистанцию и калории.
func DayActionInfo(data string, weight, height float64) string {
	step, dur, err := parsePackage(data)
	if err != nil {
		log.Println(err)
		return ""
	}
	km := float64(step) * stepLength / float64(mInKm)
	kkal, err := spentcalories.WalkingSpentCalories(step, weight, height, dur)
	if err != nil {
		log.Println(err)
		return ""
	}
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", step, km, kkal)
}
