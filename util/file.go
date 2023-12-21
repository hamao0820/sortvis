package util

import (
	"bufio"
	"os"
	"strconv"
)

func ReadFileAndConvertToIntSlice(filePath string) ([]int, error) {
	// ファイルを開く
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// スライスを初期化
	var numbers []int

	// ファイルを行ごとに読み込む
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// 文字列を整数に変換
		number, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}

		// 整数をスライスに追加
		numbers = append(numbers, number)
	}

	// エラーチェック
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return numbers, nil
}
