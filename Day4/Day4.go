package main

import (
    "fmt"
    "strconv"
    "strings"
)

type Passport struct {
    input    string
    innerMap map[string]string
}

func toKeyValuePair(value string) (string, string) {
    array := strings.Split(value, ":")
    return array[0], array[1]
}

func (p *Passport) Init(input string) *Passport {
    p.innerMap = make(map[string]string)
    sanitized := strings.Replace(input, "\n", " ", -1)
    array := strings.Split(sanitized, " ")
    for _, line := range array {
        key, val := toKeyValuePair(line)
        p.innerMap[key] = val
    }
    return p
}

func (p Passport) IsValid() bool {
    _, cidExists := p.innerMap["cid"]
    if cidExists {
        return len(p.innerMap) == 8
    } else {
        return len(p.innerMap) == 7
    }
}

/* Logic for Part 2 */

func (p Passport) IsValidV2() bool {
    if !p.IsValid() {
        return false
    }

    data := p.innerMap

    birthYear, _ := strconv.Atoi(data["byr"])
    if birthYear < 1920 || birthYear > 2002 {
        return false
    }

    issueYear, _ := strconv.Atoi(data["iyr"])
    if issueYear < 2010 || issueYear > 2020 {
        return false
    }

    expirationYear, _ := strconv.Atoi(data["eyr"])
    if expirationYear < 2020 || expirationYear > 2030 {
        return false
    }

    heightIsValid := p.hasValidHeight()
    if !heightIsValid {
        return false
    }

    hairColorIsValid := p.hasValidHairColor()
    if !hairColorIsValid {
        return false
    }

    eyeColorIsValid := p.hasValidEyeColor()
    if !eyeColorIsValid {
        return false
    }

    passportIdIsValid := p.hasValidPassportId()
    if !passportIdIsValid {
        return false
    }

    return true
}

func (p Passport) hasValidHeight() bool {
    height, exists := p.innerMap["hgt"]
    if !exists {
        return false
    }

    if strings.Contains(height, "in") {
        val := strings.Replace(height, "in", "", -1)
        intVal, _ := strconv.Atoi(val)
        if intVal >= 59 && intVal <= 76 {
            return true
        }
    }

    if strings.Contains(height, "cm") {
        val := strings.Replace(height, "cm", "", -1)
        intVal, _ := strconv.Atoi(val)
        if intVal >= 150 && intVal <= 193 {
            return true
        }
    }

    return false
}

func (p Passport) hasValidHairColor() bool {
    hcl, exists := p.innerMap["hcl"]
    if !exists {
        return false
    }

    if !strings.HasPrefix(hcl, "#") || len(hcl) != 7 {
        return false
    }

    alphaNumerics := "0123456789abcdef"
    for i := 1; i < len(hcl); i++ {
        if !strings.ContainsRune(alphaNumerics, rune(hcl[i])) {
            return false
        }
    }

    return true
}

func (p Passport) hasValidEyeColor() bool {
    validEyeColors := strings.Split("amb blu brn gry grn hzl oth", " ")
    eyeColor, exists := p.innerMap["ecl"]
    if !exists {
        return false
    }
    for _, validColor := range validEyeColors {
        if eyeColor == validColor {
            return true
        }
    }
    return false
}

func (p Passport) hasValidPassportId() bool {
    passportId, exists := p.innerMap["pid"]
    if !exists {
        return false
    }
    if len(passportId) != 9 {
        return false
    }

    for _, character := range passportId {
        if character < '0' || character > '9' {
            return false
        }
    }
    return true
}

func main() {
    inputArray := strings.Split(PuzzleInput, "\n\n")
    var count int
    for _, passportInput := range inputArray {
        passport := new(Passport)
        passport.Init(passportInput)
        if passport.IsValid() {
            count++
        }
    }

    fmt.Printf("Part 1: %v valid passports\n", count)

    count = 0
    for _, passportInput := range inputArray {
        passport := new(Passport)
        passport.Init(passportInput)
        if passport.IsValidV2() {
            count++
        }
    }

    fmt.Printf("Part 2: %v valid passports", count)
}
