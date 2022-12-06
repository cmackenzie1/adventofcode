package twentytwenty

class PasswordPolicy(
    val min: Int,
    val max: Int,
    val char: Char
) {
    companion object Factory {
        fun fromString(input: String): PasswordPolicy {
            val matchResult = """(?<min>\d+)\-(?<max>\d+)\s(?<char>\w)""".toRegex().matchEntire(input)!!
            val min = matchResult.groups["min"]!!.value.toInt()
            val max = matchResult.groups["max"]!!.value.toInt()
            val char = matchResult.groups["char"]!!.value
            return PasswordPolicy(min = min, max = max, char = char.toCharArray().first())
        }
    }

    fun isOldPolicyValid(password: String): Boolean {
        val count = password.asIterable().groupingBy { it }.eachCount()
        if (char !in count) return false
        return count.getOrElse(char) { return false } in min..max
    }

    fun isNewPolicyValid(password: String): Boolean {
        val first = min - 1
        val last = max - 1
        if (last > password.length) return false
        return (password[first] == char).xor(password[last] == char)
    }

    override fun toString(): String {
        return "PasswordPolicy(min=$min, max=$max, char=$char)"
    }

}

fun main() {
    val day2 = Day2()
    day2.partOne()
    day2.partTwo()
}

class Day2 {
    private val inputText = ResourceLoader.loadText("day2.txt")
    fun partOne() {
        var correct = 0
        inputText.reader().forEachLine {
            val split = it.split(":")
            val password = split.last().trim()
            val passwordPolicy = PasswordPolicy.fromString(split.first())
            when (passwordPolicy.isOldPolicyValid(password)) {
                true -> correct += 1
            }
        }
        println("Part One Correct Passwords: $correct")
    }

    fun partTwo() {
        var correct = 0
        inputText.reader().forEachLine {
            val split = it.split(":")
            val password = split.last().trim()
            val passwordPolicy = PasswordPolicy.fromString(split.first())
            when (passwordPolicy.isNewPolicyValid(password)) {
                true -> correct += 1
            }
        }
        println("Part Two Correct Passwords: $correct")
    }
}