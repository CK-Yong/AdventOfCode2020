package main

import (
    "testing"
)

var validPassports = []string{
    "ecl:gry pid:860033327 eyr:2020 hcl:#fffffd\nbyr:1937 iyr:2017 cid:147 hgt:183cm",
    "hcl:#ae17e1 iyr:2013\neyr:2024\necl:brn pid:760753108 byr:1931\nhgt:179cm",
}

func Test_should_be_valid_passport(test *testing.T) {
    for _, input := range validPassports {
        passport := new(Passport)
        passport.Init(input)

        result := passport.IsValid()
        if result == false {
            test.Errorf("Expected passport %v to be valid. Got %v", input, result)
        }
    }
}

var invalidPassports = []string{
    "iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884\nhcl:#cfa07d byr:1929",
    "hcl:#cfa07d eyr:2025 pid:166559648\niyr:2011 ecl:brn hgt:59in",
}

func Test_should_be_invalid_passport(test *testing.T) {
    for _, input := range invalidPassports {
        passport := new(Passport)
        passport.Init(input)

        result := passport.IsValid()
        if result == true {
            test.Errorf("Expected passport %v to be invalid. Got %v", input, result)
        }
    }
}

var validPassportsV2 = []string{
    "pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980\nhcl:#623a2f",
    "eyr:2029 ecl:blu cid:129 byr:1989\niyr:2014 pid:896056539 hcl:#a97842 hgt:165cm",
    "hcl:#888785\nhgt:164cm byr:2001 iyr:2015 cid:88\npid:545766238 ecl:hzl\neyr:2022",
    "iyr:2010 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719",
}

func Test_should_be_valid_V2_passports(test *testing.T){
    for _, input := range validPassportsV2 {
        passport := new(Passport)
        passport.Init(input)

        result := passport.IsValidV2()
        if result == false {
            test.Errorf("Expected passport %v to be valid. Got %v", input, result)
        }
    }
}

var invalidPassportsV2 = []string{
    "eyr:1972 cid:100\nhcl:#18171d ecl:amb hgt:170 pid:186cm iyr:2018 byr:1926",
    "iyr:2019\nhcl:#60297 eyr:1967 hgt:170cm\necl:grn pid:012533040 byr:1946",
    "hcl:dab227 iyr:2012\necl:brn hgt:182cm pid:021572410 eyr:2020 byr:1992 cid:277",
    "hgt:59cm ecl:zzz\neyr:2038 hcl:74454a iyr:2023\npid:3556412378 byr:2007",
}

func Test_should_be_invalid_V2_passports(test *testing.T){
    for _, input := range invalidPassportsV2 {
        passport := new(Passport)
        passport.Init(input)

        result := passport.IsValidV2()
        if result == true {
            test.Errorf("Expected passport %v to be invalid. Got %v", input, result)
        }
    }
}

func Test_should_be_invalid_height(test *testing.T){
    passport := new(Passport)
    passport.Init("hgt:170")

    result := passport.hasValidHeight()

    if result == true {
        test.Errorf("Expected invalid height")
    }
}

func Test_should_be_invalid_hair_color(test *testing.T){
    passport := new(Passport)
    passport.Init("hcl:#123abz")

    result := passport.hasValidHairColor()

    if result == true {
        test.Errorf("Expected invalid hair color")
    }
}

func Test_should_be_invalid_eye_color(test *testing.T){
    passport := new(Passport)
    passport.Init("ecl:zzz")

    result := passport.hasValidEyeColor()

    if result == true {
        test.Errorf("Expected invalid eye color")
    }
}

func Test_should_be_valid_pid(test *testing.T){
    passport := new(Passport)
    passport.Init("pid:0123456789")

    result := passport.hasValidPassportId()

    if result == true {
        test.Errorf("Expected invalid passport ID")
    }
}