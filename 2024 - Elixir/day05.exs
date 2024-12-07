defmodule Day05 do
  def solve(filename) do
    {graph, updates} =
      filename
      |> File.read!()
      |> parse()

    correct_updates =
      Enum.filter(updates, fn update ->
        update
        |> Enum.chunk_every(2, 1, :discard)
        |> Enum.reduce_while(true, fn [parent, child], _ ->
          if child in graph[parent] do
            {:cont, true}
          else
            {:halt, false}
          end
        end)
      end)

    part1_result = sum_of_center_elements(correct_updates)

    incorrect_updates =
      Enum.map(updates -- correct_updates, fn incorrect_update ->
        Enum.sort(incorrect_update, &(&1 not in graph[&2]))
      end)

    part2_result = sum_of_center_elements(incorrect_updates)

    {part1_result, part2_result}
  end

  defp sum_of_center_elements(lists) do
    lists
    |> Enum.map(fn list ->
      center = div(length(list), 2)
      list |> Enum.at(center) |> String.to_integer()
    end)
    |> Enum.sum()
  end

  defp parse(input) do
    # Split into two sections based on the double newline
    [rules, updates] = String.split(input, "\n\n", trim: true)

    rules =
      rules
      |> String.split("\n", trim: true)
      |> Enum.map(&String.split(&1, "|", trim: true))

    graph =
      rules
      |> List.flatten()
      |> MapSet.new()
      |> Enum.reduce(%{}, fn node, graph -> Map.put(graph, node, []) end)

    graph =
      Enum.reduce(rules, graph, fn [parent, child], graph ->
        Map.put(graph, parent, Map.fetch!(graph, parent) ++ [child])
      end)

    updates =
      updates
      |> String.split("\n", trim: true)
      |> Enum.map(&String.split(&1, ",", trim: true))

    {graph, updates}
  end
end

# Main script execution
{part1, part2} = Day05.solve("input/day05.txt")
IO.puts("Day 5, Part 1: #{part1}")
IO.puts("Day 5, Part 2: #{part2}")
