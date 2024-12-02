defmodule Day02 do
  def solve(filename) do
    reports =
      filename
      |> File.read!()
      |> String.split("\n", trim: true)
      |> Enum.map(&parse_line/1)

    part1_result =
      filter_stable(reports)
      |> Enum.count()

    part2_result =
      filter_stable_with_tolerance(reports)
      |> Enum.count()

    {part1_result, part2_result}
  end

  defp filter_stable(list) do
    Enum.filter(list, fn list ->
      valid_range? =
        list
        |> Enum.chunk_every(2, 1, :discard)
        |> Enum.all?(fn [a, b] ->
          diff = abs(b - a)
          diff >= 1 and diff <= 3
        end)

      trend_check =
        list
        |> Enum.chunk_every(2, 1, :discard)
        |> Enum.map(fn [a, b] ->
          cond do
            b > a -> :increase
            b < a -> :decrease
            true -> :same
          end
        end)
        |> Enum.uniq()

      valid_range? and (trend_check == [:increase] or trend_check == [:decrease])
    end)
  end

  defp filter_stable_with_tolerance(list) do
    Enum.filter(list, fn report ->
      # Try removing each index once
      Enum.with_index(report)
      |> Enum.any?(fn {_, remove_index} ->
        modified_report = List.delete_at(report, remove_index)

        valid_range? =
          modified_report
          |> Enum.chunk_every(2, 1, :discard)
          |> Enum.all?(fn [a, b] ->
            diff = abs(b - a)
            diff >= 1 and diff <= 3
          end)

        trend_check =
          modified_report
          |> Enum.chunk_every(2, 1, :discard)
          |> Enum.map(fn [a, b] ->
            cond do
              b > a -> :increase
              b < a -> :decrease
              true -> :same
            end
          end)
          |> Enum.uniq()

        # Only count as safe if modified report is consistently increasing or decreasing
        valid_range? and (trend_check == [:increase] or trend_check == [:decrease])
      end)
    end)
  end

  defp parse_line(line) do
    line
    |> String.split()
    |> Enum.map(&String.trim/1)
    |> Enum.map(&String.to_integer/1)
  end
end

{part1, part2} = Day02.solve("input/day02.txt")
IO.puts("Day 2, Part 1: #{part1}")
IO.puts("Day 2, Part 2: #{part2}")
