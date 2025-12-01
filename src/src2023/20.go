package src2023

import (
	"regexp"
	"strings"

	"github.com/vkalekis/advent-of-code/pkg/logger"
	"github.com/vkalekis/advent-of-code/pkg/utils"
)

type broadcaster struct {
	label      string
	dstModules []string
}

type flipflop struct {
	label      string
	state      string
	dstModules []string
}

type conjuction struct {
	label      string
	state      map[string]int
	dstModules []string
}

type output struct {
	label string
	state int
}

type pulse struct {
	src, dst string
	state    int
}

type component interface {
	fire(pulse) ([]pulse, bool)
}

var LOW = -1
var HIGH = +1

func parseLines20(reader utils.Reader) (broadcaster, map[string]component) {
	re := regexp.MustCompile(`(%?&?\w+) -> (.+)`)

	var b broadcaster
	components := make(map[string]component)
	cjs := make([]*conjuction, 0)

	for line := range reader.Stream() {
		line = utils.StandardizeSpaces(line)

		logger.Debugf("Line: %v", line)

		match := re.FindStringSubmatch(line)

		if len(match) > 2 {
			left := match[1]
			right := strings.ReplaceAll(match[2], " ", "")
			splRight := strings.Split(right, ",")

			logger.Debugf("L: %v R: %v", left, right)

			if left == "broadcaster" {
				b = broadcaster{
					label:      "broadcaster",
					dstModules: splRight,
				}
			} else if strings.HasPrefix(left, "%") {
				label := strings.TrimPrefix(left, "%")
				components[label] = &flipflop{
					label:      label,
					state:      "off",
					dstModules: splRight,
				}
			} else if strings.HasPrefix(left, "&") {
				label := strings.TrimPrefix(left, "&")
				cj := &conjuction{
					label:      label,
					state:      make(map[string]int),
					dstModules: splRight,
				}
				components[label] = cj
				cjs = append(cjs, cj)
			}
		}

	}

	components["output"] = &output{
		label: "output",
		state: -1,
	}
	components["rx"] = &output{
		label: "rx",
		state: 0,
	}

	for _, cj := range cjs {
		for _, c := range components {
			switch v := c.(type) {
			case *flipflop:
				for _, dst := range v.dstModules {
					if dst == cj.label {
						cj.addSrc(v.label)
					}
				}
			case *output:
			case *conjuction:
			default:
				logger.Errorf("Unknown type")
				panic("unknown type")
			}
		}
		for _, dst := range b.dstModules {
			if dst == cj.label {
				cj.addSrc(b.label)
			}
		}
	}
	return b, components
}

func (ff *flipflop) fire(p pulse) ([]pulse, bool) {
	pulses := make([]pulse, 0)

	if p.state == HIGH {
		return pulses, false
	}

	switch ff.state {
	case "off":
		ff.state = "on"
		for _, dst := range ff.dstModules {
			pulses = append(pulses, pulse{
				src:   ff.label,
				dst:   dst,
				state: HIGH,
			})
		}
	case "on":
		ff.state = "off"
		for _, dst := range ff.dstModules {
			pulses = append(pulses, pulse{
				src:   ff.label,
				dst:   dst,
				state: LOW,
			})
		}
	}

	return pulses, true
}

func (cj *conjuction) fire(p pulse) ([]pulse, bool) {
	pulses := make([]pulse, 0)

	cj.state[p.src] = p.state

	sumStates := 0
	for _, state := range cj.state {
		sumStates += state
	}
	pulsesState := HIGH
	if sumStates == len(cj.state) {
		pulsesState = LOW
	}

	for _, dst := range cj.dstModules {
		pulses = append(pulses, pulse{
			src:   cj.label,
			dst:   dst,
			state: pulsesState,
		})
	}

	return pulses, true
}

func (cj *conjuction) addSrc(src string) {
	if _, found := cj.state[src]; !found {
		cj.state[src] = LOW
	}
}

func (o *output) fire(p pulse) ([]pulse, bool) {
	o.state++
	logger.Debugf(">> Pew: %+v", p)
	return nil, true
}

func pushButton(components map[string]component, b broadcaster) (int, int) {
	lowPulses, highPulses := 0, 0

	q := utils.NewQueue[pulse]()

	// the original button press
	lowPulses++

	logger.Debugf("%v", q.Items())

	if output, ok := components["rx"].(*output); ok {
		output.state = 0
	}

	for _, dst := range b.dstModules {
		q.Enqueue(pulse{
			src:   "br",
			dst:   dst,
			state: LOW,
		})
		lowPulses++
	}

	for !q.IsEmpty() {
		pulse, ok := q.Dequeue()
		if !ok {
			logger.Errorf("Error on dequeue on empty queue")
			panic("Error on dequeue on empty queue")
		}

		logger.Debugf("Popped: %+v", pulse)

		comp, found := components[pulse.dst]
		if !found {
			continue
		}

		compPulses, fired := comp.fire(pulse)
		if fired {
			for _, compPulse := range compPulses {
				q.Enqueue(compPulse)
				switch compPulse.state {
				case LOW:
					lowPulses++
				case HIGH:
					highPulses++
				}
			}
		}
	}

	return lowPulses, highPulses
}

func (s Solver) Day_20(part int, reader utils.Reader) int {

	b, components := parseLines20(reader)

	logger.Infof("Broadcaster: %+v", b)
	for _, c := range components {
		logger.Infof("Component: %+v", c)
	}

	switch part {
	case 1:

		buttonPress, maxButtonPresses := 0, 1000
		var highPulses, lowPulses, iterHighPulses, iterLowPulses int
		for buttonPress < maxButtonPresses {

			iterLowPulses, iterHighPulses = pushButton(components, b)
			for _, c := range components {
				logger.Debugf("Component: %+v", c)
			}
			buttonPress++

			highPulses += iterHighPulses
			lowPulses += iterLowPulses
		}
		logger.Infof("High pulses: %d Low Pulses: %d", highPulses, lowPulses)
		return highPulses * lowPulses
	case 2:

		// buttonPress, maxButtonPresses := 0, 10000000

		// for buttonPress < maxButtonPresses {
		// 	if output, ok := components["rx"].(*output); ok {
		// 		output.state = 0
		// 	}

		// 	_, _ = pushButton(components, b)
		// 	buttonPress++

		// 	if output, ok := components["rx"].(*output); ok {
		// 		logger.Infof("Rx: %d", output.state)

		// 		if output.state == 1 {
		// 			return buttonPress
		// 		}
		// 	}
		// }
		return -1

	default:
		//shouldn't reach here
		return -1
	}
	return -1
}
