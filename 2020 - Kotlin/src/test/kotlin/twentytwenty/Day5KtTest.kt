package twentytwenty

import org.junit.jupiter.api.Test

internal class Day5KtTest {

    @Test
    fun lower() {
        assert((0..127).lower() == 0..63) { "failed to meet expected range of 0..63" }
        assert((32..63).lower() == 32..47) { "failed to meet expected range of 32..47" }
        assert((16..16).lower() == 16..16) { "failed to meet expected range of 16..16" }
        assert((44..45).lower() == 44..44) { "failed to meet expected range of 44..44" }
    }

    @Test
    fun upper() {
        assert((0..63).upper() == 32..63) { "failed to meet expected range of 32..63" }
        assert((32..47).upper() == 40..47) { "failed to meet expected range of 40..47" }
        assert((16..16).upper() == 16..16) { "failed to meet expected range of 16..16" }
        assert((4..5).upper() == 5..5) { "failed to meet expected range of 16..16" }
    }

    @Test
    fun computeSeat() {
        val seat = "FBFBBFFRLR"
        println(computeSeat(seat))
    }
}