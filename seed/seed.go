package seed

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/jackc/pgx/v4/pgxpool"
)

type SeedData struct {
	TableName string
	Columns   []string
	Values    []interface{}
}

type Seed struct {
	Db *pgxpool.Pool
}

func seed(ctx context.Context, s Seed, Method string) {

	m := reflect.ValueOf(&s).MethodByName(Method)

	if !m.IsValid() {

		log.Fatal("no seed method with name :", Method)

	}

	log.Println("seeding ...")

	// TODO fix call
	m.Call(nil)

	log.Println("seeding completed")

}

func Excute(pool *pgxpool.Pool, Methods ...string) {

	s := Seed{Db: pool}

	seedType := reflect.TypeOf(s)

	if len(Methods) == 0 {

		// excute all seeders if no method is given

		for i := 0; i < seedType.NumMethod(); i++ {

			m := seedType.Method(i)

			seed(context.Background(), s, m.Name)

		}

	}

	// excute only specified methods
	for _, item := range Methods {
		seed(context.Background(), s, item)
	}

}

func (s *Seed) SeedTable(tableName string, columns []string, values []interface{}) error {

	if len(tableName) == 0 || len(columns) == 0 {

		return fmt.Errorf("columns and values must not be  empty")

	}

	placeholders := make([]string, len(values))
	for i := range values {
		placeholders[i] = fmt.Sprintf("$%d", i+1)
	}

	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES ( %s)", tableName, strings.Join(columns, ","),
		strings.Join(placeholders, ","))

	_, err := s.Db.Exec(context.Background(), query, values...)

	if err != nil {

		return err
	}

	fmt.Printf("inserted into %s value %v", tableName, values)
	return nil

}
