# Stream in the lines
lines =
  File.stream!("day01.txt")

# Split lines into left and right, then parse them into integers
values =
  lines
  |> Enum.map(fn string ->
    String.split(string)
    |> Enum.flat_map(fn string ->
      case Integer.parse(string) do
        # transform to integer
        {int, _rest} -> [int]
        # skip the value
        :error -> []
      end
    end)
    |> List.to_tuple()
  end)

# split them into their left and right halves
{left, right} = Enum.unzip(values)

left = left |> Enum.sort()
right = right |> Enum.sort()

# Sum the absolute differences across all entries
result = Enum.zip_reduce(left, right, 0, fn x, y, acc -> abs(x - y) + acc end)
IO.puts("Day 1, Part 1: #{result}")

# Do frequency count on right list
frequencies = Enum.frequencies(right) |> IO.inspect()
# Sum the left * right (if present), otherwise 0
result = Enum.reduce(left, 0, fn x, acc -> acc + x * Map.get(frequencies, x, 0) end)

IO.puts("Day 1, Part 2: #{result}")
