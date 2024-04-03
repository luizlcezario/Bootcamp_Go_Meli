package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"reflect"
)

// representation of the a db with basic functions

type BD interface {
	SalvarAllStruct(objcs []Model) error
	ListAll() error
	Close()
}

// the struct of the DB if its use with a CSV

type CsvBd struct {
	fd *os.File
}

func newBd(filename string) BD {
	file, err := os.OpenFile(filename, os.O_RDWR, os.ModeAppend)
	if err != nil {
		panic("file not open")
	}
	return &CsvBd{
		fd: file,
	}
}

func (c *CsvBd) Close() {
	defer c.fd.Close()
}

func (c *CsvBd) SalvarAllStruct(objcs []Model) error {
	//config writer
	writer := csv.NewWriter(c.fd)
	writer.Comma = ';'

	// get fields and number of fields
	record := make([][]string, len(objcs)+1)
	var rowL int
	record[0], rowL = objcs[0].fieldsList()

	// create each line with each product using the number and the name of fields
	for i, obj := range objcs {
		record[i+1] = make([]string, rowL)
		for a := 0; a < rowL; a++ {
			record[i+1][a] = fmt.Sprintf("%v", reflect.ValueOf(obj).Elem().FieldByName(record[0][a]))
		}
	}

	// write all the table in csv
	err := writer.WriteAll(record)
	c.fd.Seek(0, 0)
	return err
}

func (c *CsvBd) ListAll() error {
	reader := csv.NewReader(c.fd)
	reader.Comma = ';'

	data, err := reader.ReadAll()
	if err != nil || err == io.EOF {
		return errors.New("errro on read file")
	}
	for _, lst := range data {
		line := ""
		for i, op := range lst {
			if i == 0 {
				line += fmt.Sprintf("%-20s", op)
			} else {
				line += fmt.Sprintf(" %10s", op)
			}
		}
		fmt.Println(line)
	}
	return nil
}
