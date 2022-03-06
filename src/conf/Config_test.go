package conf

import (
	"fmt"
	"testing"
)

func TestGetConfig(t *testing.T) {
	config := GetConfig()
	fmt.Println(config)
}
