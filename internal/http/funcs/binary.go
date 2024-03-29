package funcs

import (
	"strconv"
	"strings"
)

func BinToByte(binStr string) []byte {
	// Вычисляем длину массива байтов
	byteLen := (len(binStr) + 7) / 8
	bytes := make([]byte, byteLen)

	// Проходим по каждому биту в строке и формируем байты
	for i := 0; i < len(binStr); i += 8 {
		byteEnd := i + 8
		if byteEnd > len(binStr) {
			byteEnd = len(binStr)
		}
		byteStr := binStr[i:byteEnd]

		byteVal := byte(0)
		for j, bit := range byteStr {
			byteVal <<= 1
			if bit == '1' {
				byteVal |= 1
			} else if bit != '0' {
				panic("Недопустимый символ в бинарной строке")
			}
			if j == 7 {
				break
			}
		}
		bytes[i/8] = byteVal
	}

	return bytes
}

func DecimalArrayToBinary(decArr []int, last_len int) (result string) {
	for i := range decArr {
		if i == len(decArr)-1 {
			result += strings.Repeat("0", (last_len - len(DecimalToBinary(decArr[i]))))
			result += DecimalToBinary(decArr[i])
		} else {
			result += strings.Repeat("0", (11 - len(DecimalToBinary(decArr[i]))))
			result += DecimalToBinary(decArr[i])
		}
	}
	return result
}

func BinaryArrayToDecimalArray(binaryArray [][]bool) []int {
	var decimalArray []int
	for _, row := range binaryArray {
		var decimal int
		for _, val := range row {
			if val {
				decimal = (decimal << 1) | 1
			} else {
				decimal = decimal << 1
			}
		}
		decimalArray = append(decimalArray, decimal)
	}
	return decimalArray
}

// Переводим текст в биты
func ByteToBin(bytes []byte) []bool {
	bits := make([]bool, len(bytes)*8)
	for i, b := range bytes {
		for j := 0; j < 8; j++ {
			bits[i*8+j] = b&(1<<(7-j)) != 0
		}
	}
	return bits
}

func GetRemainder(x, d string) string {
	var r string
	var n, m, z, i int
	var arr []byte
	var zeroes string
	if len(x) < 5 {
		zeroes += strings.Repeat("0", (5 - len(x)))
		x = zeroes + x
	}

	z = 0
	n = len(x)
	m = len(d)
	// Инициализация
	for i = 0; i < n; i++ {
		if x[i] != '0' && x[i] != '1' {
			return ""
		}
		r += string(x[i])
	}
	arr = []byte(r)

	for {
		if x == "0" {
			arr = []byte{48, 48, 48, 48}
			break
		} else if x == "1" {
			arr = []byte{48, 48, 48, 49}
			break
		}
		for i = 0; i < m; i++ {
			// Прибавляем делитель в текущей позиции указателя
			arr[z+i] = Plus(arr[z+i], d[i])
		}
		for arr[z] == '0' {
			z++ // Сдвиг указателя влево, до первой единицы
			if z >= len(arr) {
				break
			}
		}
		if z > n-m {
			break // Конец деления
		}
	}

	return string(arr)[(len(string(arr)) - 4):]
}

func Plus(a, b byte) byte {
	if a == '0' && b == '0' {
		return '0'
	}
	if a == '1' && b == '1' {
		return '0'
	}
	return '1'
}

func DecimalToBinary(dec int) string {
	return strconv.FormatInt(int64(dec), 2)
}
func BinaryToDecimal(bin string) int {
	dec, err := strconv.ParseInt(bin, 2, 64)
	if err != nil {
		return 0
	}
	return int(dec)
}
func Polynom_vector(err string) (result int) {
	lastFour := err[len(err)-4:]
	if lastFour == "0001" {
		result = 1
	} else if lastFour == "0010" {
		result = 2
	} else if lastFour == "0100" {
		result = 4
	} else if lastFour == "1000" {
		result = 8
	} else if lastFour == "0011" {
		result = 16
	} else if lastFour == "0110" {
		result = 32
	} else if lastFour == "1100" {
		result = 64
	} else if lastFour == "1011" {
		result = 128
	} else if lastFour == "0101" {
		result = 256
	} else if lastFour == "1010" {
		result = 512
	} else if lastFour == "0111" {
		result = 1024
	} else if lastFour == "1110" {
		result = 2048
	} else if lastFour == "1111" {
		result = 4096
	} else if lastFour == "1101" {
		result = 8192
	} else if lastFour == "1001" {
		result = 16384
	}
	return result
}
