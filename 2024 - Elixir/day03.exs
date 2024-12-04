defmodule Day03 do
  def solve(filename) do
    input =
      filename
      |> File.read!()

    part1_result =
      Regex.scan(~r/mul\((\d{1,3}),(\d{1,3})\)/, input, capture: :all_but_first)
      |> Enum.reduce(0, fn [a, b], acc ->
        acc + String.to_integer(a) * String.to_integer(b)
      end)

    part2_result =
      Regex.scan(
        ~r/mul\((\d{1,3}),(\d{1,3})\) | don\'t\(\) | do\(\)/x,
        input
      )
      |> filter_instructions()
      |> Enum.reduce(0, fn [_, b, c], acc ->
        acc + String.to_integer(b) * String.to_integer(c)
      end)

    {part1_result, part2_result}
  end

  defp filter_instructions(instructions) do
    {filtered, _} =
      Enum.reduce(instructions, {[], :enabled}, fn
        ["do()"], {acc, _} ->
          {acc, :enabled}

        ["don't()"], {acc, _} ->
          {acc, :disabled}

        instruction, {acc, state} ->
          case state do
            :enabled -> {acc ++ [instruction], state}
            :disabled -> {acc, state}
          end
      end)

    filtered
  end
end

# Main script execution
{part1, part2} = Day03.solve("input/day03.txt")
IO.puts("Day 3, Part 1: #{part1}")
IO.puts("Day 3, Part 2: #{part2}")
