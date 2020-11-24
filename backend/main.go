package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/websocket"
	"github.com/mrWinston/knuffon/backend/manager"

	log "github.com/sirupsen/logrus"
)

func readStdinLine(question string) (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(question)
	read, err := reader.ReadString('\n')
	return strings.TrimSpace(read), err
}

func checkOrigin(r *http.Request) bool {
	return true
}

var upgrader = websocket.Upgrader{
	EnableCompression: true,
	CheckOrigin:       checkOrigin,
}
var gm = manager.NewGameManager()

func HandleWS(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("Got an Error during Connection Upgrade")
		return
	}

	for {
		_, msg, err := conn.ReadMessage()

		if err != nil {
			log.WithFields(log.Fields{
				"error": err,
			}).Error("Got an Error during Connection Upgrade")
			return
		}
		log.WithFields(log.Fields{"msg": msg}).Debug("Got a new Message")
		gm.AddMessage(&manager.RawMessage{Raw: msg, Conn: conn})
	}
}

func main() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
	log.SetReportCaller(true)
	defer gm.Stop()
	http.HandleFunc("/ws", HandleWS)

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Error(err)
	}
}

//func main_old() {
//
//	g := game.CreateGame([]*game.Player{
//		game.NewPlayer("Alf", "1"),
//		game.NewPlayer("Rudi", "2"),
//		game.NewPlayer("Christin", "3"),
//		game.NewPlayer("Rachel", "4"),
//	})
//	fmt.Println("Commands: roll, reset, exit")
//	var result *game.RollResult
//	for {
//		currPlayer := g.GetCurrentPlayer().Name
//		input, err := readStdinLine(fmt.Sprintf("What to Do?(%s)", currPlayer))
//		if err != nil {
//			log.Fatal(err)
//		}
//
//		input = strings.TrimSpace(input)
//
//		if input == "roll" {
//			diceToRoll := queryDiceToRoll()
//			result, err = g.Roll(diceToRoll)
//			if err != nil {
//				fmt.Println(err)
//			} else {
//				printResult(result)
//			}
//
//		} else if input == "done" {
//			err := g.TurnDone(queryResultToChoose(result))
//			if err != nil {
//				log.Fatal(err)
//			}
//		} else if input == "exit" {
//			log.Fatal("Let's get outta here")
//		} else {
//			fmt.Println("Whut?")
//		}
//	}
//}
//
//func queryDiceToRoll() []bool {
//	diceToRollRaw, _ := readStdinLine("Which ones? 1 - 5, comma separated")
//	return []bool{
//		strings.Contains(diceToRollRaw, "1"),
//		strings.Contains(diceToRollRaw, "2"),
//		strings.Contains(diceToRollRaw, "3"),
//		strings.Contains(diceToRollRaw, "4"),
//		strings.Contains(diceToRollRaw, "5"),
//	}
//}
//
//func printDice(dice []int) {
//	res := ""
//	dieToString := []string{
//		"⚀", "⚁", "⚂", "⚃", "⚄", "⚅",
//	}
//
//	for _, die := range dice {
//		res += " " + dieToString[die-1]
//	}
//
//	fmt.Printf("Dice Are: %s\n", res)
//}
//
//func printResult(res *game.RollResult) {
//	printDice(res.Dice)
//	fmt.Printf("Results Are: \n")
//	for rt, score := range res.Result {
//		fmt.Printf("%s: %d\n", rt, score)
//	}
//}
//
//func queryResultToChoose(res *game.RollResult) game.ResultType {
//	var intToResultType map[int]game.ResultType = map[int]game.ResultType{}
//	i := 0
//	for rt, score := range res.Result {
//		fmt.Printf("%v - %s: %d\n", i, rt, score)
//		intToResultType[i] = rt
//		i++
//	}
//	numRaw, _ := readStdinLine("Select a Number")
//	num, err := strconv.Atoi(numRaw)
//	for err != nil {
//		fmt.Printf("'%s' is not a number\n", numRaw)
//		numRaw, _ = readStdinLine("Select a Number")
//		num, err = strconv.Atoi(numRaw)
//	}
//	return intToResultType[num]
//
//}
