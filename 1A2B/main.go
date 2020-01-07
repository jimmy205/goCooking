package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/google/uuid"
)

var r = rand.New(rand.NewSource(time.Now().Unix()))
var memberAns = map[string]string{}

func main() {
	RunEngine()
}

type errorRes struct {
	ErrorCode string `json:"error_code"`
	ErrorText string `json:"error_text"`
}

type checkAnsReq struct {
	Sid   string `json:"sid"`
	Guess string `json:"guess"`
}

func checkAns(c *gin.Context) {

	req := checkAnsReq{}
	c.ShouldBindJSON(&req)
	answer, ok := memberAns[req.Sid]
	if !ok {
		c.JSON(http.StatusOK, errorRes{
			ErrorCode: "1111",
			ErrorText: "使用者還沒有設定答案喔！",
		})
		return
	}

	guess := req.Guess

	var A, B int
	ansMap := map[string]int{}
	for i := 0; i < len(answer); i++ {
		ansMap[string(answer[i])] = i
	}

	for i := 0; i < len(guess); i++ {
		ansIndex, ok := ansMap[string(guess[i])]
		if ok {
			if ansIndex == i {
				A++
				continue
			}
			B++
		}
	}

	// for i, v := range s {
	// 	ansMap[string(v)] = string(i)
	// }

	// for i, v := range g {
	// 	ans, ok := ansMap[string(v)]
	// 	if ok {
	// 		if string(ans) == string(i) {
	// 			A++
	// 			continue
	// 		}
	// 		B++
	// 	}
	// }

	fmt.Println(A, "A", B, "B")
}

type loginRes struct {
	Sid string `json:"sid"`
}

func fakeLogin(c *gin.Context) {
	u := uuid.New()

	id := u.String()
	ans := generateAns()
	memberAns[id] = ans

	fmt.Println("m :", memberAns)
	c.JSON(http.StatusOK, loginRes{
		Sid: u.String(),
	})
}

func generateAns() (ans string) {

	for {
		k := rand.Intn(10)
		isUse, _ := BasicNumber[k]
		if isUse {
			continue
		}
		if len(ans) == 4 {
			break
		}
		ans += strconv.Itoa(k)
		BasicNumber[k] = true
	}

	return ans
}
