package main

import (
    "fmt"
    "strconv"
    "strings"
)

func main() {
    input := ParseCommands(PuzzleInput)
    var part1BootCode = make([]Command, len(input))
    copy(part1BootCode, input)
    traverser := Traverser{BootCode: part1BootCode}
    traverser.Traverse()

    fmt.Printf("Part 1: Accumulator has value %v\n", traverser.Accumulator)

    var part2BootCode = make([]Command, len(input))
    copy(part2BootCode, input)
    b := GetTraverserForTraversibleRoute(part2BootCode)
    fmt.Printf("Part 2: Accumulator has value %v", b.Accumulator)
}

func ParseCommands(input string) []Command {
    split := strings.Split(input, "\n")

    commands := make([]Command, 0)
    for _, entry := range split {
        cmd := strings.Split(entry, " ")
        count, _ := strconv.Atoi(cmd[1])
        commands = append(commands, Command{Step: cmd[0], Count: count})
    }
    return commands
}

type Command struct {
    Step        string
    Count       int
    wasExecuted bool
}

type Traverser struct {
    BootCode           []Command
    Index, Accumulator int
}

func (t *Traverser) Traverse() {
    cmd := t.BootCode[t.Index]

    if cmd.wasExecuted {
        return
    }

    t.executeCommand(cmd)

    t.Traverse()
}

func (t *Traverser) executeCommand(cmd Command) {
    t.BootCode[t.Index].wasExecuted = true
    switch cmd.Step {
    case "nop":
        t.Index++
        break
    case "acc":
        t.Accumulator += cmd.Count
        t.Index++
        break
    case "jmp":
        t.Index += cmd.Count
        break
    }
}

func GetTraverserForTraversibleRoute(commands []Command) Traverser {
    var traverser Traverser

    for i := range commands {
        var bootCode = make([]Command, len(commands))
        copy(bootCode, commands)

        traverser = Traverser{BootCode: bootCode}
        cmdFix := traverser.BootCode[i]
        if cmdFix.Step != "jmp" {
            continue
        }

        traverser.BootCode[i].Step = "nop"

        if traverser.IsTraversable() {
            break
        }
    }

    return traverser
}

func (t *Traverser) IsTraversable() bool {
    if t.Index >= len(t.BootCode)  {
        return true
    }

    cmd := t.BootCode[t.Index]

    if cmd.wasExecuted {
        return false
    }

    t.executeCommand(cmd)

    return t.IsTraversable()
}
