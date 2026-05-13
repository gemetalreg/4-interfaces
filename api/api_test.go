package api

import (
	"encoding/json"
	"io"
	"os"
	"strings"
	"testing"
)

// Вспомогательная функция для создания временного файла с json-содержимым
func createTempJSONFile(t *testing.T, content string) *os.File {
	t.Helper()
	tmpfile, err := os.CreateTemp("", "*.json")
	if err != nil {
		t.Fatal(err)
	}
	if _, err := tmpfile.Write([]byte(content)); err != nil {
		t.Fatal(err)
	}
	tmpfile.Close()
	return tmpfile
}

// Вспомогательная функция для извлечения id bin из ответа сервера
func extractIDFromCreateResponse(resp string) string {
	type response struct {
		Record struct {
			ID string `json:"id"`
		} `json:"record"`
	}
	var r response
	_ = json.Unmarshal([]byte(resp), &r)
	return r.Record.ID
}

// Тест создания bin
func TestCreateBin(t *testing.T) {
	file := createTempJSONFile(t, `{"field": "test-create"}`)
	defer os.Remove(file.Name())

	// Подготовка параметров (имя не влияет на API)

	fileFlag := file.Name()
	nameFlag := new(string)
	*nameFlag = "test-bin"

	output := captureOutput(func() {
		Create(&fileFlag, nameFlag)
	})

	id := extractIDFromCreateResponse(output)
	if id == "" {
		t.Fatalf("Не удалось получить id из ответа: %s", output)
	}

	// Удалить после теста
	defer Delete(&id)
}

// Тест обновления bin
func TestUpdateBin(t *testing.T) {
	// Сначала создаём bin
	file := createTempJSONFile(t, `{"field": "test-update-init"}`)
	defer os.Remove(file.Name())
	fileFlag := file.Name()
	nameFlag := new(string)
	*nameFlag = "test-bin"
	createOutput := captureOutput(func() {
		Create(&fileFlag, nameFlag)
	})
	id := extractIDFromCreateResponse(createOutput)
	defer Delete(&id)

	// Теперь обновляем содержимое
	updateFile := createTempJSONFile(t, `{"field": "test-updated"}`)
	defer os.Remove(updateFile.Name())

	updateFileName := updateFile.Name()
	updateFileFlag := &updateFileName

	idFlag := &id

	updateOutput := captureOutput(func() {
		Update(updateFileFlag, idFlag)
	})

	if !strings.Contains(updateOutput, "record") {
		t.Fatalf("Некорректный ответ на update: %s", updateOutput)
	}
}

// Тест получения bin
func TestGetBin(t *testing.T) {
	file := createTempJSONFile(t, `{"field": "test-get"}`)
	defer os.Remove(file.Name())
	fileFlag := file.Name()
	nameFlag := new(string)
	*nameFlag = "test-bin"
	createOutput := captureOutput(func() {
		Create(&fileFlag, nameFlag)
	})
	id := extractIDFromCreateResponse(createOutput)
	defer Delete(&id)

	idFlag := &id
	getOutput := captureOutput(func() {
		Get(idFlag)
	})
	if !strings.Contains(getOutput, `"field":"test-get"`) && !strings.Contains(getOutput, `"field": "test-get"`) {
		t.Fatalf("Ошибка получения bin: %s", getOutput)
	}
}

// Тест удаления bin
func TestDeleteBin(t *testing.T) {
	file := createTempJSONFile(t, `{"field": "test-delete"}`)
	defer os.Remove(file.Name())
	fileFlag := file.Name()
	nameFlag := new(string)
	*nameFlag = "test-bin"
	createOutput := captureOutput(func() {
		Create(&fileFlag, nameFlag)
	})
	id := extractIDFromCreateResponse(createOutput)
	idFlag := &id

	captureOutput(func() {
		Delete(idFlag)
	})
	// Тут можно проверить, что bin действительно удалён,
	// например попробовав получить его — и ожидать ошибку 404
}

// Вспомогательная функция для захвата вывода stdout
func captureOutput(f func()) string {
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	_ = w.Close()
	out, _ := io.ReadAll(r)
	os.Stdout = rescueStdout
	return string(out)
}
