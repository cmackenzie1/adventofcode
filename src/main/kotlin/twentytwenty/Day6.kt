package twentytwenty


fun main() {
    val day6 = Day6()
    day6.partOne()
    day6.partTwo()
}

class Day6 {

    private val inputText = ResourceLoader.loadText("day6.txt")

    fun partOne() {
        inputText.split("\n\n").flatMap { group ->
            group.toCharArray().filter { it.isLetter() }.toSet()
        }.count().let(::println)
        return
    }

    fun partTwo() {
        inputText.split("\n\n").map { group ->
            group.lines().map {
                it.toCharArray().filter { it.isLetter() }.toSet()
            }.reduceRight { set, acc -> set.intersect(acc) }.count()
        }.sum().let(::println)
    }
}