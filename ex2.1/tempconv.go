// tempconv 包负责转换摄氏温度与华氏温度和开尔文温度间的转换
package tempconv

import (
	"fmt"
)

type Celsius float64
type Fahrenheit float64
type Kelvins float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string {
	return fmt.Sprintf("%gC", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%gF", f)
}

func (k Kelvins) String() string {
	return fmt.Sprintf("%gK", k)
}
