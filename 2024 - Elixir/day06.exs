defmodule GridWalker do
  def walk(grid, {start_x, start_y}, {dx, dy}, target) do
    max_x = length(Enum.at(grid, 0)) - 1
    max_y = length(grid) - 1

    do_walk(grid, {start_x, start_y}, {dx, dy}, target, max_x, max_y)
  end

  defp do_walk(grid, {x, y}, {dx, dy}, target, max_x, max_y) do
    cond do
      x < 0 or x > max_x or y < 0 or y > max_y ->
        :not_found

      grid
      |> Enum.at(y, [])
      |> Enum.at(x, :out_of_bounds) == target ->
        {:found, {x, y}}

      true ->
        do_walk(grid, {x + dx, y + dy}, {dx, dy}, target, max_x, max_y)
    end
  end
end

defmodule Day06 do
  def solve(filename) do
    grid =
      filename
      |> File.read!()
      |> parse()
      |> IO.inspect(lists: :as_lists)

    curr = current_pos(grid) |> IO.inspect()

    GridWalker.walk(grid, curr, {0, -1}, "#") |> IO.inspect()

    part1_result = nil
    part2_result = nil

    {part1_result, part2_result}
  end

  defp traverse(grid, seen) do
    start = current_pos(grid)
    dir = direction(get_element(grid, start))
  end

  defp current_pos(grid) do
    rows = length(grid)
    cols = length(Enum.at(grid, 0))

    Enum.with_index(grid)
    |> Enum.find_value(fn {row, y} ->
      case Enum.find_index(row, fn elem -> String.match?(elem, ~r/\^|>|v|</) end) do
        nil -> nil
        x -> {x, y}
      end
    end)
  end

  defp direction(c) do
    case c do
      "^" -> {0, -1}
      ">" -> {1, 0}
      "v" -> {0, 1}
      "<" -> {-1, 0}
    end
  end

  defp rotate_90(c) do
    case c do
      "^" -> ">"
      ">" -> "v"
      "v" -> "<"
      "<" -> "^"
    end
  end

  defp get_element(grid, pos) do
    {x, y} = pos
    row = Enum.at(grid, y)
    Enum.at(row, x)
  end

  defp parse(input) do
    String.split(input, "\n", trim: true)
    |> Enum.map(fn row -> String.graphemes(row) end)
  end
end

# Main script execution
{part1, part2} = Day06.solve("input/day06.txt")
IO.puts("Day 5, Part 1: #{part1}")
IO.puts("Day 5, Part 2: #{part2}")
