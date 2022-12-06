package twentytwenty

import java.io.File


fun main() {
    val day3 = Day3()
    day3.partOne()
    day3.partTwo()
}

class Day3 {
    fun partOne() {
        val lines =
            File("/Users/colemackenzie/github/pragmaticdev-io/adventofcode/src/main/resources/twentytwenty/day3.txt")
                .readLines().map { it.trim() }
        val answer = toboggan(3, 1, lines)
        println("Part One Trees Hit: $answer")
    }

    fun partTwo() {
        val lines =
            File("/Users/colemackenzie/github/pragmaticdev-io/adventofcode/src/main/resources/twentytwenty/day3.txt")
                .readLines().map { it.trim() }
        val answer = arrayListOf(
            toboggan(1, 1, lines),
            toboggan(3, 1, lines),
            toboggan(5, 1, lines),
            toboggan(7, 1, lines),
            toboggan(1, 2, lines),
        ).reduce { acc, i -> acc * i }
        println("Part Two Trees Hit: $answer")
    }
}

fun isTree(char: Char): Boolean = char == '#'

fun toboggan(right: Int, down: Int, lines: List<String>): Int {
    var x = 0
    var y = 0
    var bonks = 0
    while (y < lines.size) {
        if (isTree(lines[y][x % lines.first().length])) bonks++
        x += right
        y += down
    }
    return bonks
}