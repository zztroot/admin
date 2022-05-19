package util

import (
	"fmt"

	"strconv"

	"github.com/zztroot/zztlog"
)

func Float32ToString(n int, v float32) string {
	return fmt.Sprintf("%s%%", strconv.FormatFloat(float64(v), 'f', n, 64))
}

func Float32ToFloat32(n int, v float32) float32 {
	value, err := strconv.ParseFloat(strconv.FormatFloat(float64(v), 'f', n, 64), 64)
	if err != nil {
		zztlog.Error(err)
	}
	return float32(value)
}
