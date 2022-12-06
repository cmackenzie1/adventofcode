package twentytwenty

class Day1 {
    val inputText = ResourceLoader.loadText("day1.txt")
    fun partOne() {
        val seen = mutableSetOf<Int>()
        inputText.lines().forEach {
            val x = it.toInt()
            val target = 2020 - x
            if (target in seen) {
                val answer = target * x
                println("Day 2 - Total: $answer")

            } else {
                seen.add(x)
            }
        }
    }

    fun partTwo() {
        val lines =
            inputText.reader().readLines()
                .map {
                    it.toInt()
                }

        for (x in lines) for (y in lines) for (z in lines) {
            if (x + y + z == 2020) {
                val total = x * y * z
                println("Day 2 - Total: $total")
            }
        }
    }
}

fun main() {
    val day1 = Day1()
    day1.partOne()
    day1.partTwo()
}