// tempconv负责摄氏温度与华氏温度和开尔文温度间的转换
package tempconv

// CToF将摄氏度转换为华氏度
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

// FToC将华氏度转换为摄氏度
func FToc(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

// CToK将摄氏度转换为开尔文
func CToK(c Celsius) Kelvins {
	return Kelvins(c + 273.15)
}

// KToC将开尔文转换为摄氏度
func KToC(k Kelvins) Celsius {
	return Celsius(k - 273.15)
}
