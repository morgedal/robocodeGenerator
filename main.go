package main

import (
	"./Config"
	"./Generator"
	"./RobocodeRunner"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	StartDuelingGenerator()
}

func StartDuelingGenerator() {
	template := Generator.GetTemplate()
	robotNames := []string{"FirstRobot", "SecondRobot"}
	histories := InitDuelingRobots(robotNames[0], robotNames[1], template, Generator.MainGrammar)
	results := RobocodeRunner.GetResults()

	for i := 0; i < Config.GENERATION_STEPS; i++ {
		var worseIndex, betterIndex int
		if results[robotNames[0]] >= results[robotNames[1]] {
			worseIndex = 1
			betterIndex = 0
		} else if results[robotNames[0]] < results[robotNames[1]] {
			worseIndex = 0
			betterIndex = 1
		} else {
			panic("Cannot read results after battle")
		}

		ReGenerateRobot(robotNames[worseIndex], template, results[robotNames[betterIndex]], Generator.MainGrammar,
			&(histories[betterIndex]), &(histories[worseIndex]))

		RobocodeRunner.RunRobocode()
		results = RobocodeRunner.GetResults()
	}
}

func InitDuelingRobots(firstName, secondName, template string, grammar Generator.Grammar) [][]Generator.Production {
	var history1 []Generator.Production
	var history2 []Generator.Production

	generatedRobot_1 := Generator.GenerateFromGrammar(template, grammar, &history1)
	generatedRobot_2 := Generator.GenerateFromGrammar(template, grammar, &history2)
	Generator.SaveGeneratedRobot(firstName, generatedRobot_1)
	Generator.SaveGeneratedRobot(secondName, generatedRobot_2)

	RobocodeRunner.CompileRobot(firstName)
	RobocodeRunner.CompileRobot(secondName)
	RobocodeRunner.RunRobocode()

	return [][]Generator.Production{history1, history2}
}

func ReGenerateRobot(robotName, template string, winnerResult int, grammar Generator.Grammar, betterHist, worseHist *[]Generator.Production) {
	var generatedRobot string
	if winnerResult > Config.MINIMAL_RESULT_THRESHOLD {
		generatedRobot = ReGenerateFromHistory(template, grammar, betterHist, worseHist)
	} else {
		generatedRobot = ReGenerateFromScratch(template, grammar, worseHist)
	}
	Generator.SaveGeneratedRobot(robotName, generatedRobot)
	RobocodeRunner.CompileRobot(robotName)
}

func ReGenerateFromHistory(template string, grammar Generator.Grammar, betterHist, worseHist *[]Generator.Production) string {
	*worseHist = make([]Generator.Production, len(*betterHist))
	copy(*worseHist, *betterHist)
	return Generator.GenerateFromHistory(template, grammar, worseHist)
}

func ReGenerateFromScratch(template string, grammar Generator.Grammar, history *[]Generator.Production) string {
	*history = nil
	return Generator.GenerateFromGrammar(template, grammar, history)
}
