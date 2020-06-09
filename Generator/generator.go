package Generator

import (
	"../Config"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
)

func getRandom(strList []string) string {
	r := rand.Intn(len(strList))
	return strList[r]
}

func findNTSymbol(text string, g Grammar) (bool, string) {
	for ntSymbol := range g {
		if strings.Contains(text, ntSymbol) {
			return true, ntSymbol
		}
	}
	return false,""
}

func GenerateFromGrammar(text string, g Grammar, history* []Production) string {
	found, ntSymbol := findNTSymbol(text, g)
	for found {
		replacement := getRandom(g[ntSymbol])
		text = strings.Replace(text, ntSymbol, replacement, 1 )
		p := Production{ntSymbol, replacement}
		*history = append(*history, p)
		found, ntSymbol = findNTSymbol(text, g)
	}
	return text
}

func GenerateFromHistory(text string, g Grammar, history* []Production) string {
	mutationProbability := Config.MUTATION_PROB
	for i,prod := range *history {
		if rand.Float32() <= mutationProbability && prod.left != "BODY" {
			replacement := getRandom(g[prod.left])
			text = strings.Replace(text, prod.left, replacement, 1 )
			(*history)[i].right = replacement
		} else {
			text = strings.Replace(text, prod.left, prod.right, 1)
		}
	}
	return GenerateFromGrammar(text, g, history)
}

func GetTemplate() string {
	content, err := ioutil.ReadFile(Config.TEMPLATE_FILE)
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}

func SaveGeneratedRobot(name, robotSrc string) {
	robotPath := Config.ROBOT_DIR + name + ".java"
	robotSrc = strings.Replace(robotSrc, "ROBOT_NAME", name, 1)

	if _, err := os.Stat(robotPath); err == nil {
		err = os.Remove(robotPath)
		if err != nil {log.Println("Error removing old robot class: ",err)}
	}
	file, err := os.OpenFile(robotPath, os.O_CREATE, 0777)
	defer file.Close()
	if err != nil { fmt.Println(err) }
	_, err = file.WriteString(robotSrc)
	if err != nil { fmt.Printf("Error while writing to file: %s\n",err) }
}