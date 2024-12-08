package utils

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
)

func Scan(fp string, ch chan string) error {
	defer close(ch)

	file, err := os.Open(fp)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ch <- scanner.Text()
	}

	if err = scanner.Err(); err != nil {
		return err
	}

	return nil
}

func Fetch(day string) (string, error) {
	fp := "input/" + day + ".txt"
	if _, err := os.Stat(fp); err == nil {
		return fp, nil
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://adventofcode.com/2024/day/"+day+"/input", nil)
	if err != nil {
		return "", err
	}

	sess, ok := os.LookupEnv("AOC24_SESSION")
	if !ok || sess == "" {
		return "", fmt.Errorf("AOC24_SESSION environment variable not set")
	}

	req.AddCookie(&http.Cookie{
		Name:  "session",
		Value: sess,
	})

	fmt.Println("Fetching input for day ", day)
	resp, err := client.Do(req)
	if err != nil {
		return "", nil
	}
	defer resp.Body.Close()

	file, err := os.Create(fp)
	if err != nil {
		return "", err
	}
	defer file.Close()

	file.ReadFrom(resp.Body)

	return fp, nil
}
