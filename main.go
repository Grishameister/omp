package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/Grishameister/omp/csvParser"
	"github.com/Grishameister/omp/domain"
	"github.com/Grishameister/omp/jsonParser"
	"log"
	"os"
	"strings"
)

func open(filename *string) (domain.Result, error) {
	if filename == nil {
		return domain.Result{}, errors.New("invalid ptr")
	}

	idx := strings.LastIndex(*filename, ".")
	if idx == -1 {
		return domain.Result{}, errors.New("unknown format")
	}

	format := string([]byte(*filename)[idx + 1:])
	if format != "json" && format != "csv" {
		return domain.Result{}, errors.New("unknown format")
	}

	file, err := os.Open(*filename)
	if err != nil {
		return domain.Result{}, errors.New("can't open file")
	}
	defer func() {
		if err := file.Close(); err != nil {
			return
		}
	}()

	res := domain.Result{}
	if format == "json" {
		parser := jsonParser.JsonParser{}
		err := parser.Parse(file, &res)
		return res, err
	} else {
		parser := csvParser.CsvParser{}
		err := parser.Parse(file, &res)
		return res, err
	}
}

func main() {
	flag.Parse()
	if len(flag.Args()) != 1 {
		log.Fatal("first arg is filename")
	}

	filename := flag.Arg(0)
	res, err := open(&filename)

	if err != nil {
		log.Println(err.Error())
		return
	}
	fmt.Printf("Максимальная цена %d у продукта %s, Максимальный рейтинг %d у Продукта %s",
		res.MaxPrice, res.MaxPriceProduct, res.MaxRating, res.MaxRatingProduct)
}
