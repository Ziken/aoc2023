package main

import (
	"bytes"
	"fmt"
	"github.com/Ziken/advent_of_code_2023/utils"
	"regexp"
	"strconv"
)

type Condition struct {
	Left       string
	Right      int
	Operator   string
	Evaluate   func(n int) bool
	IfResult   string
	ElseResult string
}

type VariableRange struct {
	Min, Max int
}
type Workflow struct {
	Name       string
	Conditions []Condition
}

func parseData(data [][]byte) (map[string]Workflow, []map[string]int) {
	var workflows = make(map[string]Workflow)
	var variables []map[string]int
	var nameRegexp = regexp.MustCompile(`(\w+){(.+)}`)
	var conditionRegexp = regexp.MustCompile(`(\w+)(<|>)(\d+):(\w+)`)
	var variableRegexp = regexp.MustCompile(`(\w)=(\d+)`)

	var i = 0
	// parse workflows
	for ; i < len(data); i++ {
		var line = data[i]
		if len(line) == 0 {
			break
		}
		var parsedLine = nameRegexp.FindSubmatch(line)
		var workflowName, rawConditions = parsedLine[1], parsedLine[2]
		var workflow = Workflow{Name: string(workflowName)}
		var conditions = bytes.Split(rawConditions, []byte(","))

		for _, condition := range conditions {
			var parsedCondition = conditionRegexp.FindSubmatch(condition)
			// last condition
			if len(parsedCondition) < 5 {
				workflow.Conditions[len(workflow.Conditions)-1].ElseResult = string(condition)
				continue
			}
			var left, operator, right, ifResult = parsedCondition[1], parsedCondition[2], parsedCondition[3], parsedCondition[4]
			var rightInt, _ = strconv.Atoi(string(right))
			var evaluate func(int int) bool
			if string(operator) == "<" {
				evaluate = func(n int) bool {
					return n < rightInt
				}
			} else {
				evaluate = func(int int) bool {
					return int > rightInt
				}
			}
			var c = Condition{
				Left:       string(left),
				Right:      rightInt,
				Evaluate:   evaluate,
				IfResult:   string(ifResult),
				Operator:   string(operator),
				ElseResult: "",
			}
			workflow.Conditions = append(workflow.Conditions, c)
		}
		workflows[workflow.Name] = workflow
	}
	// omit empty line
	i++
	for ; i < len(data); i++ {
		var line = data[i]
		var parsedLine = variableRegexp.FindAllSubmatch(line, -1)
		var parts = map[string]int{}
		for _, matched := range parsedLine {
			var variableName, variableValue = matched[1], matched[2]
			var value, _ = strconv.Atoi(string(variableValue))
			parts[string(variableName)] = value
		}
		variables = append(variables, parts)
	}
	return workflows, variables
}

func partOne(workflows map[string]Workflow, variables []map[string]int) (out int) {
	var variableNames = []string{"x", "m", "a", "s"}
	var foundAnswer = false
	for _, part := range variables {
		var workflow = workflows["in"]
		var usedWorkflows = map[string]bool{"in": true}
		foundAnswer = false
		for !foundAnswer {
			for _, condition := range workflow.Conditions {
				variableName := condition.Left
				var result = ""
				if condition.Evaluate(part[variableName]) {
					result = condition.IfResult
				} else {
					result = condition.ElseResult
				}
				if result == "A" {
					for _, vn := range variableNames {
						out += part[vn]
					}
					foundAnswer = true
					break
				} else if result == "R" {
					foundAnswer = true
					break
				} else if result != "" {
					if ok, _ := usedWorkflows[result]; !ok {
						usedWorkflows[result] = true
						workflow = workflows[result]
						break
					}
				}

			}
		}
	}
	return
}

func traverseWorkflow(workflows map[string]Workflow, workflow Workflow, currentPart map[string]VariableRange, visitedWorkflows map[string]bool, gatheredResult *[]map[string]VariableRange) {
	if _, ok := visitedWorkflows[workflow.Name]; ok {
		return
	}
	visitedWorkflows[workflow.Name] = true
	var elseConditionPart = map[string]VariableRange{
		"x": currentPart["x"],
		"m": currentPart["m"],
		"a": currentPart["a"],
		"s": currentPart["s"],
	}
	for _, condition := range workflow.Conditions {
		var ifConditionPart = map[string]VariableRange{
			"x": elseConditionPart["x"],
			"m": elseConditionPart["m"],
			"a": elseConditionPart["a"],
			"s": elseConditionPart["s"],
		}
		variableName := condition.Left
		if condition.Operator == "<" {
			if ifConditionPart[variableName].Max > condition.Right {
				ifConditionPart[variableName] = VariableRange{Min: ifConditionPart[variableName].Min, Max: condition.Right - 1}
			}
			if elseConditionPart[variableName].Min < condition.Right {
				elseConditionPart[variableName] = VariableRange{Min: condition.Right, Max: elseConditionPart[variableName].Max}
			}
		} else {
			if ifConditionPart[variableName].Min < condition.Right {
				ifConditionPart[variableName] = VariableRange{Min: condition.Right + 1, Max: ifConditionPart[variableName].Max}
			}
			if elseConditionPart[variableName].Max > condition.Right {
				elseConditionPart[variableName] = VariableRange{Min: elseConditionPart[variableName].Min, Max: condition.Right}
			}
		}
		if condition.IfResult == "A" {
			*gatheredResult = append(*gatheredResult, ifConditionPart)
		} else if condition.IfResult == "R" {
			// ignore
		} else {
			var newVisitedWorkflows = map[string]bool{}
			for k, v := range visitedWorkflows {
				newVisitedWorkflows[k] = v
			}
			traverseWorkflow(workflows, workflows[condition.IfResult], ifConditionPart, newVisitedWorkflows, gatheredResult)
		}

		if condition.ElseResult == "A" {
			*gatheredResult = append(*gatheredResult, elseConditionPart)
		} else if condition.ElseResult == "R" {
			// ignore
		} else if condition.ElseResult != "" {
			var newVisitedWorkflows = map[string]bool{}
			for k, v := range visitedWorkflows {
				newVisitedWorkflows[k] = v
			}
			traverseWorkflow(workflows, workflows[condition.ElseResult], elseConditionPart, newVisitedWorkflows, gatheredResult)
		}
	}
}
func partTwo(workflows map[string]Workflow) (out int) {
	var currentPart = map[string]VariableRange{
		"x": {Min: 1, Max: 4000},
		"m": {Min: 1, Max: 4000},
		"a": {Min: 1, Max: 4000},
		"s": {Min: 1, Max: 4000},
	}
	var visitedWorkflows = map[string]bool{}
	var gatheredResult []map[string]VariableRange
	traverseWorkflow(workflows, workflows["in"], currentPart, visitedWorkflows, &gatheredResult)

	var result = 0
	for _, part := range gatheredResult {
		result += (part["x"].Max - part["x"].Min + 1) * (part["m"].Max - part["m"].Min + 1) * (part["a"].Max - part["a"].Min + 1) * (part["s"].Max - part["s"].Min + 1)
	}
	return result

}

func main() {
	var data = utils.GetInput("day_19/input.txt")
	var workflows, variables = parseData(data)
	fmt.Println("Part one:", partOne(workflows, variables))
	fmt.Println("Part Two:", partTwo(workflows))
}
