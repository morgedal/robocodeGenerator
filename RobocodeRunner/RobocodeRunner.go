package RobocodeRunner

import (
	"bufio"
	"../Config"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)



func CompileRobot(name string) {
	robotStr := Config.ROBOT_DIR + name + ".java"
	cmd := exec.Command("javac", robotStr, "-cp", Config.ROBOCODE_LIBS_DIR)
	err := cmd.Run()
	if err != nil {
		log.Panicf("\nERROR during compilation of robot class %s '%s'\n",name,err)
	}
}

func RunRobocode() {
	//java -Xmx512M -cp libs/robocode.jar robocode.Robocode -battle generator\ConsoleBattle.battle -results generator\results.txt
	args := []string { "-Xmx512M", "-cp", Config.ROBOCODE_JAR, "robocode.Robocode", "-battle", Config.BATTLE_FILE, "-results", Config.RESULTS_FILE }
	if !Config.DISPLAY_BATTLE {
		args = append(args, "-nodisplay")
	}
	cmd := exec.Command("java", args...)
	err := cmd.Run()
	if err != nil {
		fmt.Printf("\nERROR during running battle\n")
	}
}

func GetResults() map[string]int {
	file, err := os.Open(Config.RESULTS_FILE)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	return parseResults(scanner)
}

func parseResults(scanner *bufio.Scanner) map[string]int {
	scores := make(map[string]int)
	skipLines := 2
	for scanner.Scan() {
		line := scanner.Text()
		if skipLines > 0 { skipLines--; continue }
		splitted := strings.Split(line, "\t")
		if len(splitted) > 1 {
			scoreStr := (strings.Split(splitted[1], " "))[0]
			robotName := (strings.Split(splitted[0], " "))[1]
			score,_ := strconv.Atoi(scoreStr)
			robotName = strings.TrimPrefix(robotName, "sample.")
			scores[robotName] = score
		}
	}
	return scores
}