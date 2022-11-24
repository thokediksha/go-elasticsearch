package main

import (
	"context"
	"encoding/json"
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

	 _, err  := client.Do("RPUSH", "students", student.Name)
	 if err != nil {
		panic (err)
	 }


	 students, err := client.Do("RPOP", "students")
	 if err != nil {
		panic (err)
	 }

	 fmt.Printf("%s \n", students)
    
     task := Student{
      Name: student.Name,
	  Age: student.Age,
	  Score: student.Score,
	 }
	 
p, err := redis.Bytes(students, err)
if err != nil {
	panic(err)
}
dataJSON, err := json.Marshal(p)
js := string(dataJSON)
fmt.Println(js)

	 fmt.Println(task)

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


// func Example_JSONSet(rh *rejson.Handler) {
	

// 		student := Student{
// 			Name: Name{
// 				"Mark",
// 				"Pronto",
// 			},
// 			Age: 23,
// 			Score: 14.5,
// 		}
// 		res, err := rh.JSONSet("student", ".", student)
// 		if err != nil {
// 			fmt.Println(err)
// 			// log.Fatalf("Failed to JSONSet")
// 			return
// 		}
	
// 		if res.(string) == "OK" {
// 			fmt.Printf("Success: %s\n", res)
// 		} else {
// 			fmt.Println("Failed to Set: ")
// 		}
	
// 		studentJSON, err := redis.Bytes(rh.JSONGet("student", "."))
// 		if err != nil {
// 			log.Fatalf("Failed to JSONGet")
// 			return
// 		}
	
// 		readStudent := Student{}
// 		err = json.Unmarshal(studentJSON, &readStudent)
// 		if err != nil {
// 			log.Fatalf("Failed to JSON Unmarshal")
// 			return
// 		}
	
// 		fmt.Printf("Student read from redis : %#v\n", readStudent)
// 		fmt.Println(readStudent)
// 	}	


// NewRedisCache(c, student)
// func main() {
// 	// var addr = os.Getenv("REDIS_PORT")
// 	// var pass = os.Getenv("REDIS_PASSWORD")
// 	// db, _ := strconv.Atoi(os.Getenv("REDIS_DB"))
// 	// // fmt.Println(addr, pass, db)
// 	// Rdb = redis.NewClient(&redis.Options{
// 	// 	Addr:     addr,
// 	// 	Password: pass, // no password set
// 	// 	DB:       db,   // use default DB
// 	// })
// 	// NewRedisCache()

// 	rh := rejson.NewReJSONHandler()
// 	Example_JSONSet(rh)
// 	// var c *gin.Context
// 	if newClient, err := elastic.NewClient(
// 		elastic.SetURL("http://127.0.0.1:9200"),
// 		elastic.SetHealthcheck(false)); err != nil {
// 		panic(err)
// 	} else {
// 		elasticClient1 = newClient
// 	}

// 	ctx := context.TODO()

// 	//    student_id := 1

// 	//    if student_id == 1 {
// 	// 	NewRedisCache(c, student)
// 	//    }

// 	// var id,name,study string
// 	// student := map[string]interface{}{
// 	// 	"id":        Rdb.HGet("id", id),
// 	// 	"name":      Rdb.HGet("name" , name),
// 	// 	"study":     Rdb.HGet("study" , study),
// 	// }
// 	// studentID := "1"

// 	//    ID := 1

// 	// if ID == 1 {
// 	// 	NewRedisCache()
// 	// }

// 	student2 := rh


// 	// Insert Document
// 	_, err := elasticClient1.Index().
// 		Index("student2").
// 		// Id(studentID).
// 		BodyJson(student2).
// 		Do(ctx)
// 	panicIfError(err)
// 	flushESDB("student2")

// 	// Update Document by Id
// 	elasticClient1.Update().
// 		Index("student2").
// 		// Id(studentID).
// 		Doc(student2).
// 		Do(ctx)
// 	panicIfError(err)
// 	flushESDB("student2")

// 	elasticClient1.Refresh("student2")
// }

// func flushESDB(indexname string) error {
// 	_, err := elasticClient1.Flush().Index(indexname).Do(context.TODO())
// 	return err
// }

// func panicIfError(err error) {
// 	if err != nil {
// 		panic(err)
// 	}
// }
