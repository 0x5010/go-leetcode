package leetcode0165

func compareVersion(version1 string, version2 string) int {
	bToI := map[byte]int{
		'0': 0,
		'1': 1,
		'2': 2,
		'3': 3,
		'4': 4,
		'5': 5,
		'6': 6,
		'7': 7,
		'8': 8,
		'9': 9,
	}
	i, j := -1, -1
	for i < len(version1) || j < len(version2) {
		num1, num2 := 0, 0
		for i++; i < len(version1) && version1[i] != '.'; i++ {
			num1 = num1*10 + bToI[version1[i]]
		}
		for j++; j < len(version2) && version2[j] != '.'; j++ {
			num2 = num2*10 + bToI[version2[j]]
		}
		if num1 > num2 {
			return 1
		} else if num1 < num2 {
			return -1
		}
	}
	return 0
}
