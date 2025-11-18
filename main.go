package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Matrix [][]float64

func printMatrix(matrix Matrix) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("%8.2f", matrix[i][j])
		}
		fmt.Println()
	}
}

func inputMatrix(size int) Matrix {
	scanner := bufio.NewScanner(os.Stdin)
	matrix := make(Matrix, size)
	
	for i := 0; i < size; i++ {
		matrix[i] = make([]float64, size)
		
		for {
			fmt.Printf("Строка %d: ", i+1)
			scanner.Scan()
			line := scanner.Text()
			numbers := strings.Fields(line)
			
			if len(numbers) != size {
				fmt.Printf("Ошибка! Введите %d числа через пробел: ", size)
				continue
			}
			
			valid := true
			for j := 0; j < size; j++ {
				num, err := strconv.ParseFloat(numbers[j], 64)
				if err != nil {
					fmt.Printf("Ошибка! Введите числа: ")
					valid = false
					break
				}
				matrix[i][j] = num
			}
			
			if valid {
				break
			}
		}
	}
	return matrix
}

func addMatrix(a, b Matrix) Matrix {
	size := len(a)
	result := make(Matrix, size)
	for i := 0; i < size; i++ {
		result[i] = make([]float64, size)
		for j := 0; j < size; j++ {
			result[i][j] = a[i][j] + b[i][j]
		}
	}
	return result
}

func multiplyMatrixByScalar(matrix Matrix, scalar float64) Matrix {
	size := len(matrix)
	result := make(Matrix, size)
	for i := 0; i < size; i++ {
		result[i] = make([]float64, size)
		for j := 0; j < size; j++ {
			result[i][j] = matrix[i][j] * scalar
		}
	}
	return result
}

func multiplyMatrices(a, b Matrix) Matrix {
	size := len(a)
	result := make(Matrix, size)
	for i := 0; i < size; i++ {
		result[i] = make([]float64, size)
		for j := 0; j < size; j++ {
			sum := 0.0
			for k := 0; k < size; k++ {
				sum += a[i][k] * b[k][j]
			}
			result[i][j] = sum
		}
	}
	return result
}

func readInt() int {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		line := scanner.Text()
		num, err := strconv.Atoi(strings.TrimSpace(line))
		if err == nil {
			return num
		}
		fmt.Print("Ошибка! Введите число: ")
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	
	for {
		var choice int
		for {
			fmt.Println("1. Сложение")
			fmt.Println("2. Умножение на число")
			fmt.Println("3. Умножение матриц")
			fmt.Println("4. Выход")
			fmt.Print("Выбор: ")
			choice = readInt()
			if choice >= 1 && choice <= 4 {
				break
			}
			fmt.Println("Ошибка! Введите число от 1 до 4")
		}

		if choice == 4 {
			break
		}

		var size int
		for {
			fmt.Print("Размер (2 или 3): ")
			scanner.Scan()
			line := scanner.Text()
			num, err := strconv.Atoi(strings.TrimSpace(line))
			if err == nil && (num == 2 || num == 3) {
				size = num
				break
			}
			fmt.Println("Ошибка! Введите 2 или 3")
		}

		switch choice {
		case 1:
			fmt.Println("Матрица 1:")
			matrix1 := inputMatrix(size)
			fmt.Println("Матрица 2:")
			matrix2 := inputMatrix(size)
			result := addMatrix(matrix1, matrix2)
			fmt.Println("Результат:")
			printMatrix(result)
		case 2:
			fmt.Println("Матрица:")
			matrix := inputMatrix(size)
			fmt.Print("Число: ")
			scalar := float64(readInt())
			result := multiplyMatrixByScalar(matrix, scalar)
			fmt.Println("Результат:")
			printMatrix(result)
		case 3:
			fmt.Println("Матрица 1:")
			matrix1 := inputMatrix(size)
			fmt.Println("Матрица 2:")
			matrix2 := inputMatrix(size)
			result := multiplyMatrices(matrix1, matrix2)
			fmt.Println("Результат:")
			printMatrix(result)
		}
		fmt.Println()
	}
}