package twentytwenty

import java.io.File

//TODO: Split using "\n\n" silly.
fun main() {
    val day4 = Day4()
    day4.partOne()
    day4.partTwo()
}

class Day4 {
    private val mandatoryFields = setOf("byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid")
    private val validEyecolors = setOf("amb", "blu", "brn", "gry", "grn", "hzl", "oth")

    private val validators = mapOf<String, (String) -> Boolean>(
        "byr" to { i -> i.toInt() in 1920..2002 },
        "iyr" to { i -> i.toInt() in 2010..2020 },
        "eyr" to { i -> i.toInt() in 2020..2030 },
        "hgt" to { i ->
            val result = """(?<hgt>\d+)(?<unit>.*?)""".toRegex().matchEntire(i)!!
            val height = result.groups["hgt"]!!.value
            val unit = result.groups["unit"]!!.value
            when (unit) {
                "in" -> height.toInt() in 59..76
                "cm" -> height.toInt() in 150..193
                else -> false
            }
        },
        "hcl" to { i -> i.matches("""#[0-9a-f]{6}""".toRegex()) },
        "ecl" to { i -> i in validEyecolors },
        "pid" to { i -> i.matches("""[0-9]{9}""".toRegex()) },
        "cid" to { i -> true }
    )

    fun partOne() {
        var valid = 0
        val text =
            File("/Users/colemackenzie/github/pragmaticdev-io/adventofcode/src/main/resources/twentytwenty/day4.txt").readText()
        chunkText(text).forEach {
            val passport = mutableSetOf<String>()
            val a = it.flatMap { it.split(" ").map { it.split(":") } }
            passport.addAll(a.map { it.first() })
            valid += passport.containsAll(mandatoryFields).toInt()
        }
        println("Part One Valid Passports: ${valid}")
        return
    }

    fun partTwo() {
        var valid = 0
        val text =
            File("/Users/colemackenzie/github/pragmaticdev-io/adventofcode/src/main/resources/twentytwenty/day4.txt").readText()
        chunkText(text).forEach {
            // Each string in the List<String> belongs to a single passport
            val passport = mutableMapOf<String, String>()
            val pairs = it.flatMap { it.split(" ").map { it.split(":") } } // Put it into a list of pairs

            val passportFieldSet = mutableSetOf<String>()
            passportFieldSet.addAll(pairs.map { it.first() })

            var validPassport = passportFieldSet.containsAll(mandatoryFields)
            if (validPassport) { // Don't bother if its not valid to begin with
                pairs.forEach {
                    passport[it.first()] = it.last()
                }
                validPassport = validPassport.and(passport.map {
                    validators[it.key]?.invoke(it.value)!!
                }.reduce { acc, b -> b.and(acc) })
                valid += validPassport.toInt()
            }
        }
        println("Part Two Valid Passports: ${valid}")
        return
    }
}

fun chunkText(text: String): MutableList<List<String>> {
    val chunks = mutableListOf<List<String>>()
    var chunk = mutableListOf<String>()
    for (line in text.lines()) {
        if (line.isNotBlank()) {
            chunk.add(line)
        } else {
            chunks.add(chunk)
            chunk = mutableListOf()
        }
    }
    return chunks
}

fun Boolean.toInt() = if (this) 1 else 0
