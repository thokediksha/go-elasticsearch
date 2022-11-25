package main

import (
	"context"
	"encoding/json"
	// "strings"
	// "encoding/json"
	"log"

	// "encoding/json"
	// "encoding/json"
	// "bytes"
	"fmt"

	// "time"
	// "os"

	"github.com/olivere/elastic/v7"

	// "strings"
	// "log"
	// "github.com/gin-gonic/gin"
	// "github.com/go-redis/redis"
	// "go-elasticsearch/models"
	"github.com/gomodule/redigo/redis"
	// "strconv"
	// "github.com/nitishm/go-rejson/v4"
)

var elasticClient1 *elastic.Client
// var Rdb *redis.Client
var Rdb = newPool()


func newPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle: 80,
		MaxActive: 12000,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", ":6379")
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}
}

func main() {
	client := Rdb.Get()
    defer client.Close()
	// var c *gin.Context
	 if newClient, err := elastic.NewClient(
				elastic.SetURL("http://127.0.0.1:9200"),
				elastic.SetHealthcheck(false)); err != nil {
				panic(err)
			} else {
				elasticClient1 = newClient
			}
			ctx := context.TODO()		

    //  _, err := client.Do("SET", "mykey", "Hello from redigo!")
    //  if err != nil {
    //      panic(err)
    //  }

    //  value, err := client.Do("GET", "mykey")
    //  if err != nil {
    //      panic(err)
    //  }

    //  fmt.Printf("%s \n", value)
	type Name struct {
		First string `json:"first,omitempty"`
		Last  string `json:"last,omitempty"`
	}

	// Student - student object
	type Student struct {
		Name  Name    `json:"name,omitempty"`
		Age   int     `json:"rank,omitempty"`
		Score float64 `json:"score,omitempty"`
	}

	student := Student{
					Name: Name{
						"Mark",
						"Pronto",
					},
					Age: 23,
					Score: 14.5,
}

    //  _, err := client.Do("ZADD", "vehicles", 4, "car")
    //  if err != nil {
    //      panic(err)
    //  }
    //  _, err = client.Do("ZADD", "vehicles", 2, "bike")
    //  if err != nil {
    //      panic(err)
    //  }

    //  vehicles, err := client.Do("ZRANGE", "vehicles", 0, -1, "WITHSCORES")
    //  if err != nil {
    //      panic(err)
    //  }
    //  fmt.Printf("%s \n", vehicles)
	s, _ := json.Marshal(student)

	 _, err  := client.Do("RPUSH", "stud", s)
	 if err != nil {
		panic (err)
	 }

	 students, err := redis.String(client.Do("RPOP", "stud"))
	 if err != nil {
		panic (err)
	 }

	 fmt.Printf("%s \n", students)
	//  var task Student

	//  pd := students.(map[string]interface{})
	//  fmt.Println(pd)
	
    // v,_ := json.Marshal(students)
	// json.Unmarshal([]byte(students), &task)
	// fmt.Println(task.Name)
	// fmt.Println(students)
	// tempName := pd["name"]
	// fmt.Println(tempName)

	// v := tempName.(map[string]interface{})
	

	// fmt.Println(Name)

    //  task := Student{
    //   Name: student.Name,
	//   Age: student.Age,
	//   Score: student.Score,
	//  }


	 
// p, err := redis.Bytes(students, err)
// if err != nil {
// 	panic(err)
// }
// dataJSON, err := json.Marshal(p)
// js := string(dataJSON)
// fmt.Println(js)

	//  fmt.Println(task)

	//  Insert Document
	_, err = elasticClient1.Index().
		Index("student2").
		// Id(studentID).
		BodyJson(students).
		Do(ctx)
	log.Println(err)
	flushESDB("student2")

}


func flushESDB(indexname string) error {
	_, err := elasticClient1.Flush().Index(indexname).Do(context.TODO())
	return err
}

// func panicIfError(err error) {
// 	if err != nil {
// 		panic(err)
// 	}
// }