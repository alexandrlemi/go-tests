package main

import (
	"fmt"
)

type Saver interface {
	Save(v string)

}

type Getter interface {
	Get() []string
}

func main() {
	//конфиг читаем
	typeBD := "MS"
	var MSDB *MSSQL
	var PGDB *PG
	if typeBD == "MS" {
		MSDB = &MSSQL{addr: "localhost"}
	} else {
		PGDB = &PG{addr: "localhost"}
	}
	var s Saver
	//выбираем бд
	if typeBD == "MS" {
		s =MSDB
	} else {
		s = PGDB
	}

	var g Getter
	//выбираем бд
	if typeBD == "MS" {
		g  =&Cash{DB: MSDB}
	} else {
		g =&Cash{DB: PGDB}
	}

	
	s.Save("test")
	//куча бизнес логики
	fmt.Println(g.Get())

	//сохраняем в базу значение

}
 

type PG struct {
	addr          string
	maxConnection int
	stor          []string
}

func (pg *PG) Save(value string) {
	pg.stor = append(pg.stor, value)
}
func (pg *PG) Get() []string{
	return pg.stor
}
type MSSQL struct {
	addr          string
	maxMemorySize int64
	MSstor        []string
}

func (ms *MSSQL) Save(MSvalue string) {
	ms.MSstor = append(ms.MSstor, MSvalue)
	ms.maxMemorySize++
	//
}
func (ms *MSSQL) TestConsistansy() {
	fmt.Println("ok")
	//
}
func (ms *MSSQL) Get() []string{
	return ms.MSstor
}

type Cash struct{
	DB Getter
	tmp []string	
	flag bool
}

func (ch *Cash) Get() []string{
	if !ch.flag{
		ch.tmp=ch.DB.Get()
		ch.flag=true
	}
	return ch.tmp
} 