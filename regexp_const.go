package itools

const (
	// 整数或者小数
	regIntOrFloat = `^[0-9]+\.{0,1}[0-9]{0,2}$`

	// 纯数字
	regNumber = `^[0-9]*$`

	// 长度为n的纯数字
	regLenNNumber = `^\d{n}$`

	// 长度不小于n位的纯数字
	regGeNNumber = `^\d{n,}$`

	// 长度m~n位的纯数字
	regMNIntervalNumber = `^\d{m,n}$`

	// 非零开头的纯数字
	regStartingWithNonZero = `^(0|[1-9][0-9]*)$`

	// n位小数的正实数
	regNNovelsOfRealNumber = `^[0-9]+(.[0-9]{n})?$`

	// m~n位小数的正实数
	regMNNovelsOfRealNumber = `^[0-9]+(.[0-9]{m,n})?$`

	// 非零的正整数
	regNanZeroNumber = `^\+?[1-9][0-9]*$`

	// 非零的负整数
	regNanZeroNegNumber = `^\-[1-9][0-9]*$`

	// 长度为3的字符
	regNLeCharacter = `^.{n}$`

	// 长度m-k的字符串，包含数字、字母、汉字，不包含特殊字符 一般多用于校验用户名或昵称等
	regMKNoContainSpecialCharacter = `^[\p{Han}a-zA-Z0-9]{m,k}$`

	// 长度m-n的字符串，包含特殊字符、数字、字母、汉字， 一般多用于校验用户名或昵称等
	regMNLeCharacter = `^[\p{Han}\p{L}\p{N}\p{P}\p{S}]{m,z}$`

	// 纯英文字符串,大小写不敏感
	regEnCharacter = `^[A-Za-z]+$`

	// 纯大写英文字符串
	regUpEnCharacter = `^[A-Z]+$`

	// 纯小写英文字符串
	regLowerEnCharacter = `^[a-z]+$`

	// 数字和26个英文字母组成的字符串,大小写不敏感
	regNumberEnCharacter = `^[A-Za-z0-9]+$`

	// 数字、26个英文字母或者下划线组成的字符串
	regNumberEnUnderscores = `^\w+$`

	// 密码1 由数字、26个英文字母或者下划线组成的英文开头的字符串, 长度m~n位
	regPass1 = `^[a-zA-Z]\w{m,n}$`

	// 验证是否包含特殊字符串
	regIsContainSpecialCharacter = `[!@#\$%\^&\*\(\)_\+\[\]{}|;':",./<>?]`

	// 纯汉字
	regChineseCharacter = `^[\p{Han}]+$`

	// email
	regEmail = `^\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$`

	// 大陆手机号
	regChinesePhoneNumber = `^1[3-9]\d{9}$`

	// 验证大陆身份证号
	regChineseIDCardNumber = `^\d{15}$|^\d{17}(\d|X|x)$`

	// 匹配中文
	regContainChineseCharacter = `[\p{Han}]`

	// 匹配双字节字符(包括汉字在内)
	regDoubleByte = `[^\x00-\xff]`

	// 匹配零个或多个空白字符（包括空格、制表符、换页符等）
	regEmptyLine = `^\s*$`

	// ipv4
	regIPv4 = `^(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.` +
		`(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.` +
		`(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.` +
		`(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$`
)
