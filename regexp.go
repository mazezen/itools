package itools

import (
	"net"
	"regexp"
	"strconv"
	"strings"
)

type GoRegs struct{}

func NewGoRegs() *GoRegs {
	return &GoRegs{}
}

// MatchIntOrFloat 整数或者小数
func (g *GoRegs) MatchIntOrFloat(str string) bool {
	compile := regexp.MustCompile(regIntOrFloat)
	return compile.MatchString(str)
}

// MatchNumber 纯数字
func (g *GoRegs) MatchNumber(str string) bool {
	compile := regexp.MustCompile(regNumber)
	return compile.MatchString(str)
}

// MatchLenNNumber 长度为n的纯数字
func (g *GoRegs) MatchLenNNumber(str string, n int) bool {
	nu := strconv.Itoa(n)
	reg := strings.Replace(regLenNNumber, "n", nu, 1)
	compile := regexp.MustCompile(reg)
	return compile.MatchString(str)
}

// MatchGeNNumber 长度不小于n位的纯数字
func (g *GoRegs) MatchGeNNumber(str string, n int) bool {
	nu := strconv.Itoa(n)
	reg := strings.Replace(regGeNNumber, "n", nu, 1)
	compile := regexp.MustCompile(reg)
	return compile.MatchString(str)
}

// MatchMNIntervalNumber 长度m~n位的纯数字
func (g *GoRegs) MatchMNIntervalNumber(str string, m, n int) bool {
	mu := strconv.Itoa(m)
	nu := strconv.Itoa(n)
	reg := strings.Replace(regMNIntervalNumber, "m", mu, 1)
	reg = strings.Replace(reg, "n", nu, 1)
	compile := regexp.MustCompile(reg)
	return compile.MatchString(str)
}

// MatchStartingWithNonZero 非零开头的纯数字
func (g *GoRegs) MatchStartingWithNonZero(str string) bool {
	compile := regexp.MustCompile(regStartingWithNonZero)
	return compile.MatchString(str)
}

// MatchNNovelsOfRealNumber 有n位小数的正实数
func (g *GoRegs) MatchNNovelsOfRealNumber(str string, n int) bool {
	nu := strconv.Itoa(n)
	reg := strings.Replace(regNNovelsOfRealNumber, "n", nu, 1)
	compile := regexp.MustCompile(reg)
	return compile.MatchString(str)
}

// MatchMNNovelsOfRealNumber m~n位小数的正实数
func (g *GoRegs) MatchMNNovelsOfRealNumber(str string, m, n int) bool {
	mu := strconv.Itoa(m)
	nu := strconv.Itoa(n)
	reg := strings.Replace(regMNNovelsOfRealNumber, "m", mu, 1)
	reg = strings.Replace(reg, "n", nu, 1)
	compile := regexp.MustCompile(reg)
	return compile.MatchString(str)
}

// MatchNanZeroNumber 非零的正整数
func (g *GoRegs) MatchNanZeroNumber(str string) bool {
	compile := regexp.MustCompile(regNanZeroNumber)
	return compile.MatchString(str)
}

// MatchNanZeroNegNumber 非零的负整数
func (g *GoRegs) MatchNanZeroNegNumber(str string) bool {
	compile := regexp.MustCompile(regNanZeroNegNumber)
	return compile.MatchString(str)
}

// MatchNLeCharacter 长度为n的字符，特殊字符除外
func (g *GoRegs) MatchNLeCharacter(str string, n int) bool {
	nu := strconv.Itoa(n)
	reg := strings.Replace(regNLeCharacter, "n", nu, 1)
	compile := regexp.MustCompile(reg)
	return compile.MatchString(str)
}

// MatchMKNoContainSpecialCharacter 长度m-k的字符串，包含数字、字母、汉字，不包含特殊字符 一般多用于校验用户名或昵称等
func (g *GoRegs) MatchMKNoContainSpecialCharacter(str string, m, k int) bool {
	mu := strconv.Itoa(m)
	ku := strconv.Itoa(k)
	reg := strings.Replace(regMKNoContainSpecialCharacter, "m", mu, 1)
	reg = strings.Replace(reg, "k", ku, -1)
	compile := regexp.MustCompile(reg)
	return compile.MatchString(str)
}

// MatchMNLeCharacter 长度m-z的字符串，包含特殊字符、数字、字母、汉字， 一般多用于校验用户名或昵称等
func (g *GoRegs) MatchMNLeCharacter(str string, m, z int) bool {
	mu := strconv.Itoa(m)
	zu := strconv.Itoa(z)
	reg := strings.Replace(regMNLeCharacter, "m", mu, 1)
	reg = strings.Replace(reg, "z", zu, -1)
	compile := regexp.MustCompile(reg)
	return compile.MatchString(str)
}

// MatchEnCharacter 纯英文字符串,大小写不敏感
func (g *GoRegs) MatchEnCharacter(str string) bool {
	compile := regexp.MustCompile(regEnCharacter)
	return compile.MatchString(str)
}

// MatchUpEnCharacter 纯大写英文字符串
func (g *GoRegs) MatchUpEnCharacter(str string) bool {
	compile := regexp.MustCompile(regUpEnCharacter)
	return compile.MatchString(str)
}

// MatchLowerEnCharacter 纯小写英文字符串
func (g *GoRegs) MatchLowerEnCharacter(str string) bool {
	compile := regexp.MustCompile(regLowerEnCharacter)
	return compile.MatchString(str)
}

// MatchNumberEnCharacter 数字和26个英文字母组成的字符串,大小写不敏感
func (g *GoRegs) MatchNumberEnCharacter(str string) bool {
	compile := regexp.MustCompile(regNumberEnCharacter)
	return compile.MatchString(str)
}

// MatchNumberEnUnderscores 数字和26个英文字母组成的字符串,大小写不敏感
func (g *GoRegs) MatchNumberEnUnderscores(str string) bool {
	compile := regexp.MustCompile(regNumberEnUnderscores)
	return compile.MatchString(str)
}

// MatchPass1 密码1 由数字、26个英文字母或者下划线组成的英文开头的字符串, 长度m~n位
func (g *GoRegs) MatchPass1(str string, m, n int) bool {
	mu := strconv.Itoa(m)
	nu := strconv.Itoa(n)
	reg := strings.Replace(regPass1, "m", mu, 1)
	reg = strings.Replace(reg, "n", nu, 1)
	compile := regexp.MustCompile(reg)
	return compile.MatchString(str)
}

// MatchPass2 密码2
// 密码长度至少为8个字符。
// 包含至少一个小写字母。
// 包含至少一个大写字母。
// 包含至少一个数字。
// 包含至少一个特殊字符（例如 !@#$%^&*() 等
func (g *GoRegs) MatchPass2(str string) bool {
	if len(str) < 8 {
		return false
	}

	hasLower := regexp.MustCompile(`[a-z]`).MatchString(str)
	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(str)
	hasDigit := regexp.MustCompile(`[0-9]`).MatchString(str)
	hasSpecial := regexp.MustCompile(`[!@#~$%^&*(),.?":{}|<>]`).MatchString(str)

	return hasLower && hasUpper && hasDigit && hasSpecial
}

// MatchIsContainSpecialCharacter 验证是否包含特殊字符串
func (g *GoRegs) MatchIsContainSpecialCharacter(str string) bool {
	compile := regexp.MustCompile(regIsContainSpecialCharacter)
	return compile.MatchString(str)
}

// MatchChineseCharacter 纯汉字
func (g *GoRegs) MatchChineseCharacter(str string) bool {
	compile := regexp.MustCompile(regChineseCharacter)
	return compile.MatchString(str)
}

// MatchEmail email
func (g *GoRegs) MatchEmail(str string) bool {
	compile := regexp.MustCompile(regEmail)
	return compile.MatchString(str)
}

// MatchChinesePhoneNumber 大陆手机号
func (g *GoRegs) MatchChinesePhoneNumber(str string) bool {
	compile := regexp.MustCompile(regChinesePhoneNumber)
	return compile.MatchString(str)
}

// MatchChineseIDCardNumber 验证大陆身份证号
func (g *GoRegs) MatchChineseIDCardNumber(id string) bool {
	compile := regexp.MustCompile(regChineseIDCardNumber)
	if !compile.MatchString(id) {
		return false
	}
	if len(id) == 15 {
		// Convert 15-digit ID card to 18-digit
		id = id[:6] + "19" + id[6:]
		id = id + calculateChecksum(id)
	} else if len(id) == 18 {
		// Validate the checksum of 18-digit ID card
		return calculateChecksum(id[:17]) == string(id[17])
	}

	return false
}

// calculateChecksum calculates the checksum for the given 17-digit ID card number.
func calculateChecksum(id string) string {
	weights := []int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
	checkMap := []byte{'1', '0', 'X', '9', '8', '7', '6', '5', '4', '3', '2'}

	sum := 0
	for i, char := range id {
		num, _ := strconv.Atoi(string(char))
		sum += num * weights[i]
	}

	return string(checkMap[sum%11])
}

// MatchContainChineseCharacter 大陆手机号
func (g *GoRegs) MatchContainChineseCharacter(str string) bool {
	compile := regexp.MustCompile(regContainChineseCharacter)
	return compile.MatchString(str)
}

// MatchDoubleByte 匹配双字节字符(包括汉字在内)
func (g *GoRegs) MatchDoubleByte(input string) bool {
	re := regexp.MustCompile(regDoubleByte)
	return re.MatchString(input)
}

// MatchEmptyLine 匹配零个或多个空白字符（包括空格、制表符、换页符等）
func (g *GoRegs) MatchEmptyLine(input string) bool {
	re := regexp.MustCompile(regEmptyLine)
	return re.MatchString(input)
}

// MatchIPv4 ipv4
func (g *GoRegs) MatchIPv4(input string) bool {
	re := regexp.MustCompile(regIPv4)
	return re.MatchString(input)
}

// MatchIPv6 ipv6
func (g *GoRegs) MatchIPv6(input string) bool {
	return net.ParseIP(input) != nil && net.ParseIP(input).To4() == nil
}
