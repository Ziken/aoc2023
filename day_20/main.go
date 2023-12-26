package main

import (
	"bytes"
	"fmt"
	"github.com/Ziken/advent_of_code_2023/utils"
	"strings"
)

const (
	IGNORE_PULSE = -1
	LOW_PULSE    = 0
	HIGH_PULSE   = 1
)

type Module interface {
	Pulse(module string, pulse int) int
	GetNext() []string
	GetName() string
	Copy() Module
}

type FlipFlopModule struct {
	Name  string
	State bool
	Next  []string
}

func (f *FlipFlopModule) Pulse(module string, p int) int {
	if p == HIGH_PULSE {
		return IGNORE_PULSE
	}
	if p == LOW_PULSE && f.State {
		f.State = false
		return LOW_PULSE
	}
	f.State = true
	return HIGH_PULSE
}
func (f *FlipFlopModule) GetNext() []string {
	return f.Next
}
func (f *FlipFlopModule) GetName() string {
	return f.Name
}
func (f *FlipFlopModule) Copy() Module {
	return &FlipFlopModule{Name: f.Name, State: f.State, Next: strings.Split(strings.Join(f.Next, ","), ",")}
}

type ConjunctionModule struct {
	Name   string
	Pulses map[string]int
	Next   []string
}

func (f *ConjunctionModule) Pulse(module string, pulse int) int {
	f.Pulses[module] = pulse
	var allHighPulse = true
	for _, p := range f.Pulses {
		if p != HIGH_PULSE {
			allHighPulse = false
			break
		}
	}
	if allHighPulse {
		return LOW_PULSE
	}
	return HIGH_PULSE
}
func (f *ConjunctionModule) GetNext() []string {
	return f.Next
}
func (f *ConjunctionModule) GetName() string {
	return f.Name
}
func (f *ConjunctionModule) Copy() Module {
	var pulses = make(map[string]int)
	for key, value := range f.Pulses {
		pulses[key] = value
	}
	return &ConjunctionModule{Name: f.Name, Pulses: pulses, Next: strings.Split(strings.Join(f.Next, ","), ",")}
}

type TestModule struct {
	Name string
}

func (f *TestModule) GetName() string {
	return f.Name
}
func (f *TestModule) Copy() Module {
	return &TestModule{Name: f.Name}
}
func (f *TestModule) Pulse(module string, pulse int) int {
	return IGNORE_PULSE
}
func (f *TestModule) GetNext() []string {
	return []string{}
}

type Pulse struct {
	Module     string
	PrevModule string
	Pulse      int
}

type Node struct {
	Module   string
	Children []*Node
}

func parseData(data [][]byte) ([]string, map[string]Module) {
	var cable = map[string]Module{}
	var broadcaster = []string{}
	var conjunctionConnections = make(map[string][]string)

	for _, line := range data {
		var splitted = bytes.Split(line, []byte(" -> "))
		var module = string(splitted[0])
		if module[0] == '%' {
			var name = module[1:]
			var m = &FlipFlopModule{Name: name, Next: strings.Split(string(splitted[1]), ", ")}
			cable[name] = m
		} else if module[0] == '&' {
			var name = module[1:]
			var m = &ConjunctionModule{Name: name, Pulses: make(map[string]int), Next: strings.Split(string(splitted[1]), ", ")}
			conjunctionConnections[name] = make([]string, 0)
			cable[name] = m
		} else if module == "broadcaster" {
			broadcaster = strings.Split(string(splitted[1]), ", ")
		} else {
			var name = module
			var m = &TestModule{Name: name}
			cable[name] = m
		}
	}

	// set connected modules to conjunction modules
	for key, module := range cable {
		for _, next := range module.GetNext() {
			if _, ok := conjunctionConnections[next]; ok {
				conjunctionConnections[next] = append(conjunctionConnections[next], key)
			}
		}
	}
	// Set default pulses to conjunction modules
	for key, modules := range conjunctionConnections {
		var module = cable[key].(*ConjunctionModule)
		for _, m := range modules {
			module.Pulses[m] = LOW_PULSE
		}
		cable[key] = module
	}
	// Set not registered modules
	for _, module := range cable {
		for _, next := range module.GetNext() {
			if _, ok := cable[next]; !ok {
				cable[next] = &TestModule{Name: next}
			}
		}
	}

	return broadcaster, cable
}

func partOne(broadcaster []string, cable map[string]Module) int {
	var cableCopy = make(map[string]Module)
	for key, module := range cable {
		cableCopy[key] = module.Copy()
	}
	var highPulses = 0
	var lowPulses = 0
	for i := 0; i < 1000; i++ {
		lowPulses++
		var nextPulses = make([]Pulse, 0)
		for _, module := range broadcaster {
			var pulse = LOW_PULSE
			nextPulses = append(nextPulses, Pulse{Module: module, Pulse: pulse, PrevModule: ""})
		}
		for len(nextPulses) > 0 {
			var pulse = nextPulses[0]
			nextPulses = nextPulses[1:]
			if pulse.Pulse == IGNORE_PULSE {
				continue
			}
			if pulse.Pulse == LOW_PULSE {
				lowPulses++
			} else {
				highPulses++
			}
			var nextPulse = cableCopy[pulse.Module].Pulse(pulse.PrevModule, pulse.Pulse)
			for _, nextModule := range cableCopy[pulse.Module].GetNext() {
				nextPulses = append(nextPulses, Pulse{Module: nextModule, Pulse: nextPulse, PrevModule: pulse.Module})
			}
		}
	}
	return lowPulses * highPulses
}

func partTwo(broadcaster []string, cable map[string]Module) int {
	var cableCopy = make(map[string]Module)
	for key, module := range cable {
		cableCopy[key] = module.Copy()
	}
	var root = createTree(cableCopy)

	var moduleBeforeRoot = root.Children[0].Module

	var out = 1
	for _, key := range broadcaster {
		out *= findCycle(cable, key, moduleBeforeRoot)
	}
	return out
}

func createTree(cable map[string]Module) Node {
	var root = Node{Module: "rx"}
	var visitedChildren = make(map[string]bool)
	var checkChildren = []*Node{&root}

	for len(checkChildren) > 0 {
		var node = checkChildren[0]
		checkChildren = checkChildren[1:]
		for key, module := range cable {
			for _, next := range module.GetNext() {
				if next == node.Module && !visitedChildren[next+module.GetName()] {
					var newNode = &Node{Module: key}
					node.Children = append(node.Children, newNode)
					visitedChildren[next+module.GetName()] = true
				}
			}
		}
		checkChildren = append(checkChildren, node.Children...)
	}
	return root
}

func findCycle(cable map[string]Module, broadcastModule string, rootModule string) int {
	for i := 1; ; i++ {
		var nextPulses = make([]Pulse, 0)
		nextPulses = append(nextPulses, Pulse{Module: broadcastModule, Pulse: LOW_PULSE, PrevModule: ""})
		for len(nextPulses) > 0 {
			var pulse = nextPulses[0]
			nextPulses = nextPulses[1:]
			if pulse.Pulse == IGNORE_PULSE {
				continue
			}
			if pulse.Pulse == HIGH_PULSE && pulse.Module == rootModule {
				//fmt.Println("abc")
				return i
			}
			var nextPulse = cable[pulse.Module].Pulse(pulse.PrevModule, pulse.Pulse)

			for _, nextModule := range cable[pulse.Module].GetNext() {
				nextPulses = append(nextPulses, Pulse{Module: nextModule, Pulse: nextPulse, PrevModule: pulse.Module})
			}
		}
	}

}

func main() {
	var data = utils.GetInput("day_20/input.txt")
	var broadcaster, cable = parseData(data)
	fmt.Println("Part one:", partOne(broadcaster, cable))
	fmt.Println("Part two:", partTwo(broadcaster, cable))
}
