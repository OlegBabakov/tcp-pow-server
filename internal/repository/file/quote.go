package file

import (
	"bufio"
	"bytes"
	_ "embed"
	"math/rand"
	"strings"
)

//go:embed quotes.txt
var quotes []byte

type QuoteRepo struct {
	quotes []string
}

func NewQuote() *QuoteRepo {
	qr := new(QuoteRepo)

	reader := bytes.NewReader(quotes)
	s := bufio.NewScanner(reader)

	for s.Scan() {
		if q := strings.TrimSpace(s.Text()); q != "" {
			qr.quotes = append(qr.quotes, q)
		}
	}

	return qr
}

func (q *QuoteRepo) GetQuote() (string, error) {
	return q.quotes[rand.Intn(len(q.quotes))], nil
}
