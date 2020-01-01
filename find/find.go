package find

import (
	"log"
	"regexp"
	"strings"
)

func GetTableData(body []byte) string {
	start := regexp.MustCompile(`<tbody>`)
	end := regexp.MustCompile(`<(\/tbody)`)
	locationStart := start.FindIndex(body)
	locationEnd := end.FindIndex(body)
	log.Println(locationStart, locationEnd)
	data := string(body[(locationStart[1] + 1):(locationEnd[0] - 1)])
	return data
}

func BooksDetail(table string) [][]string {
	Regexp := regexp.MustCompile(`d>.*<`)
	var bookDetail []string
	var booksDetail [][]string
	for index, value := range Regexp.FindAllStringSubmatch(table, -1) {
		s := strings.TrimPrefix(value[0], "d>")
		s = strings.TrimSuffix(s, "<")
		if index%5 > 1 {
			bookDetail = append(bookDetail, s)
		}
		if index%5 == 4 {
			booksDetail = append(booksDetail, bookDetail)
			bookDetail = nil
		}
	}
	return booksDetail
}
