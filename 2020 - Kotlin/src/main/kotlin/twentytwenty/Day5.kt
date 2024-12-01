package twentytwenty

import kotlin.math.abs

private val ROW_RANGE = 0..127
private val COL_RANGE = 0..7

fun main() {
    val day5 = Day5()
    day5.partOne()
    day5.partTwo()
}

class Day5 {

    private val inputText = ResourceLoader.loadText("day5.txt")

    fun partOne(): Int? {
        val maxOrNull = inputText.lines().map {
            computeSeat(it)
        }.maxOrNull()
        println("Max seat: $maxOrNull")
        return maxOrNull
    }

    fun partTwo() {
        val seats = inputText.lines().map {
            computeSeat(it)
        }.sorted()
        val sumOfSeats = seats.sum()
        val mySeat = abs(sumOfSeats - (seats[1].rangeTo(seats.last())).sum())
        println("My seat: $mySeat")
        return
    }
}

fun IntRange.lower(): IntRange = this.first..(this.first + this.last) / 2

fun IntRange.upper(): IntRange = (this.first + this.last + 1) / 2..this.last

fun computeSeat(boardingPass: String): Int {
    val rowRange = search(boardingPass.take(7), ROW_RANGE)
    val colRange = search(boardingPass.takeLast(3), COL_RANGE)
    return rowRange.first * 8 + colRange.first
}

fun search(slice: String, startingRange: IntRange): IntRange {
    var range = startingRange
    for (char in slice) {
        when (char) {
            'F' -> range.lower()
            'B' -> range.upper()
            'L' -> range.lower()
            'R' -> range.upper()
            else -> 0..0
        }.also { range = it as IntRange }
    }
    return range
}