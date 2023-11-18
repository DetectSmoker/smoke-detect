package utils

import (
	"fmt"
	"time"
)

func GenerateFileName() string {
	t := time.Now()
	s := t.Unix()
	return fmt.Sprint(s)
}