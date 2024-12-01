defmodule Day01 do
  def solve(filename) do
    values =
      filename
      |> File.read!()
      |> String.split("\n", trim: true)
      |> Enum.map(&parse_line/1)

    # Part 1: Restore the original logic of sorted absolute differences
    part1_result =
      values
      |> Enum.unzip()
      |> then(fn {left, right} ->
        left
        |> Enum.sort()
        |> Enum.zip(Enum.sort(right))
        |> Enum.reduce(0, fn {x, y}, acc -> abs(x - y) + acc end)
      end)

    # Part 2: Frequency-based calculation
    frequencies =
      values
      |> Enum.map(fn {_, y} -> y end)
      |> Enum.frequencies()

    part2_result =
      values
      |> Enum.map(fn {x, _} -> x end)
      |> Enum.reduce(0, fn x, acc ->
        acc + x * Map.get(frequencies, x, 0)
      end)

    {part1_result, part2_result}
  end

  defp parse_line(line) do
    line
    |> String.split()
    |> Enum.map(&Integer.parse/1)
    |> Enum.filter(fn
      {_, ""} -> true
      _ -> false
    end)
    |> Enum.map(fn {num, _} -> num end)
    |> List.to_tuple()
  end
end

# Main script execution
{part1, part2} = Day01.solve("input/day01.txt")
IO.puts("Day 1, Part 1: #{part1}")
IO.puts("Day 1, Part 2: #{part2}")
