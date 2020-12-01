package twentytwenty

import java.io.File

fun main() {
    day1part1()
    day1part2()
}

fun day1part1() {
    print("Day 1 - ")
    val seen = mutableSetOf<Int>()
    File("/Users/colemackenzie/github/pragmaticdev-io/adventofcode/src/main/resources/twentytwenty/day1.txt").forEachLine {
        val x = it.toInt()
        val target = 2020 - x
        if (target in seen) {
            val answer = target * x
            println("Total: $answer")
        } else {
            seen.add(x)
        }
    }
}

fun day1part2() {
    val lines =
        File("/Users/colemackenzie/github/pragmaticdev-io/adventofcode/src/main/resources/twentytwenty/day1.txt").readLines()
            .map {
                it.toInt()
            }

    for (x in lines) for (y in lines) for (z in lines) {
        if (x + y + z == 2020) {
            val total = x * y * z
            println("Day 2 - Total: $total")
            return
        }
    }
}
