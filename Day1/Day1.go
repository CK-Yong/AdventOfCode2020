package main

type Pair struct{ first, second int }

func Find2020Sum(array []int) Pair {
    for _, first := range array {
        for _, second := range array {
            if first == second {
                continue
            }

            if first+second == 2020 {
                return Pair{first, second}
            }
        }
    }

    return Pair{0, 0}
}

type Triplet struct{ first, second, third int }

func Find2020SumTriplet(array []int) Triplet {
    for _, first := range array {
        for _, second := range array {
            for _, third := range array {
                if first == second && second == third {
                    continue
                }

                if first+second+third == 2020 {
                    return Triplet{first, second, third}
                }
            }
        }
    }
    return Triplet{0, 0, 0}
}

func main() {

    result := Find2020Sum(PuzzleInput)
    println("Results: ", result.first, result.second)
    // Result was 1073 * 947 = 1016131

    resultTriplet := Find2020SumTriplet(PuzzleInput)
    println("Results: ", resultTriplet.first, resultTriplet.second, resultTriplet.third)
    // Result was 911 * 618 * 491 = 276432018
}
