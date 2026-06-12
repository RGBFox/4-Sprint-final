package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	"../spentcalories"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

// распаковка полученных шагов и времени (вспомогательная функция)
func parsePackage(data string) (int, time.Duration, error) {
	s := strings.Split(data, ",")
	if len(s) != 2 {
		return 0, 0, fmt.Errorf("ошибка длины")
	}
	step, err := strconv.Atoi(s[0])
	if err != nil {
		return 0, 0, err
	}
	time, err := time.ParseDuration(s[1])
	if err != nil {
		return 0, 0, err
	}
	if step < 0 {
		return 0, 0, fmt.Errorf("ошибка длины")
	}
	return step, time, nil
}

// DayActionInfo - функция вычисляющая количество шагов, дистанции и калорий
func DayActionInfo(data string, weight, height float64) string {
	walk, dur, err := parsePackage(data)
	if err != nil {
		return ""
	}
	km := float64(walk) * stepLength / float64(mInKm)
	// далее будет определение калорий, не забыть дописать
	kkal := 
	return fmt.Printf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", walk, km, kkal)
}
