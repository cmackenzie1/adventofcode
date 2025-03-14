defmodule Day04 do
  def solve(filename) do
    input =
      filename
      |> File.read!()
      |> String.split("\n", trim: true)

    matrix = input |> Enum.map(&String.graphemes/1)

    # Part 1: Find XMAS
    word = "XMAS"
    part1_result = count_word_occurrences(matrix, word)

    # Part 2: Find X-MAS pattern
    part2_result = count_x_mas_patterns(matrix)

    {part1_result, part2_result}
  end

  defp count_word_occurrences(matrix, word) do
    directions = [
      # Horizontal
      fn r, c -> horizontal_search(matrix, word, r, c) end,
      # Vertical
      fn r, c -> vertical_search(matrix, word, r, c) end,
      # Diagonal top-left to bottom-right
      fn r, c -> diagonal_search(matrix, word, r, c, 1, 1) end,
      # Diagonal top-right to bottom-left
      fn r, c -> diagonal_search(matrix, word, r, c, 1, -1) end,
      # Diagonal bottom-left to top-right
      fn r, c -> diagonal_search(matrix, word, r, c, -1, 1) end,
      # Diagonal bottom-right to top-left
      fn r, c -> diagonal_search(matrix, word, r, c, -1, -1) end
    ]

    rows = length(matrix)
    cols = length(Enum.at(matrix, 0) || [])

    Enum.reduce(0..(rows - 1), 0, fn r, acc1 ->
      Enum.reduce(0..(cols - 1), acc1, fn c, acc2 ->
        Enum.reduce(directions, acc2, fn direction, acc3 ->
          acc3 + direction.(r, c)
        end)
      end)
    end)
  end

  defp count_x_mas_patterns(matrix) do
    rows = length(matrix)
    cols = length(Enum.at(matrix, 0) || [])

    Enum.reduce(0..(rows - 1), 0, fn r, acc1 ->
      Enum.reduce(0..(cols - 1), acc1, fn c, acc2 ->
        if check_x_mas_pattern(matrix, r, c), do: acc2 + 1, else: acc2
      end)
    end)
  end

  defp check_x_mas_pattern(matrix, row, col) do
    # Possible directions for MAS
    mas_directions = [
      # Down
      {1, 0},
      # Up
      {-1, 0},
      # Right
      {0, 1},
      # Left
      {0, -1},
      # Down-Right
      {1, 1},
      # Down-Left
      {1, -1},
      # Up-Right
      {-1, 1},
      # Up-Left
      {-1, -1}
    ]

    # Try all combinations of two different directions
    Enum.any?(mas_directions, fn {row_step1, col_step1} ->
      Enum.any?(mas_directions -- [{row_step1, col_step1}], fn {row_step2, col_step2} ->
        check_mas_in_x(matrix, row, col, row_step1, col_step1, row_step2, col_step2)
      end)
    end)
  end

  defp check_mas_in_x(matrix, row, col, row_step1, col_step1, row_step2, col_step2) do
    # Check if we can place 'MAS' in both directions
    mas_words = ["MAS", "SAM"]

    Enum.any?(mas_words, fn mas ->
      # First MAS direction
      first_match = check_word_match(matrix, mas, row, col, row_step1, col_step1) == 1

      # Second MAS direction
      second_match = check_word_match(matrix, mas, row, col, row_step2, col_step2) == 1

      first_match and second_match
    end)
  end

  defp horizontal_search(matrix, word, row, col) do
    cols = length(Enum.at(matrix, 0) || [])
    word_length = String.length(word)

    ltr =
      if col + word_length <= cols do
        check_word_match(matrix, word, row, col, 0, 1)
      else
        0
      end

    # Right to left
    rtl =
      if col - word_length + 1 >= 0 do
        check_word_match(matrix, word, row, col, 0, -1)
      else
        0
      end

    ltr + rtl
  end

  defp vertical_search(matrix, word, row, col) do
    rows = length(matrix)
    word_length = String.length(word)

    ttb =
      if row + word_length <= rows do
        check_word_match(matrix, word, row, col, 1, 0)
      else
        0
      end

    # Bottom to top
    btt =
      if row - word_length + 1 >= 0 do
        check_word_match(matrix, word, row, col, -1, 0)
      else
        0
      end

    ttb + btt
  end

  defp diagonal_search(matrix, word, row, col, row_step, col_step) do
    rows = length(matrix)
    cols = length(Enum.at(matrix, 0) || [])
    word_length = String.length(word)

    last_row = row + (word_length - 1) * row_step
    last_col = col + (word_length - 1) * col_step

    if last_row >= 0 and last_row < rows and last_col >= 0 and last_col < cols do
      check_word_match(matrix, word, row, col, row_step, col_step)
    else
      0
    end
  end

  defp check_word_match(matrix, word, start_row, start_col, row_step, col_step) do
    # Validate input matrix
    if is_nil(matrix) or matrix == [] do
      0
    else
      word_chars = String.graphemes(word)

      match =
        Enum.with_index(word_chars)
        |> Enum.all?(fn {char, index} ->
          curr_row = start_row + index * row_step
          curr_col = start_col + index * col_step

          # Additional boundary checks
          cond do
            curr_row < 0 or curr_row >= length(matrix) ->
              false

            curr_col < 0 or curr_col >= length(Enum.at(matrix, 0)) ->
              false

            true ->
              matrix_char = matrix |> Enum.at(curr_row) |> Enum.at(curr_col)
              char == matrix_char
          end
        end)

      if match, do: 1, else: 0
    end
  end
end

# Main script execution
{part1, part2} = Day04.solve("input/day04.txt")
IO.puts("Day 4, Part 1: #{part1}")
IO.puts("Day 4, Part 2: #{part2}")
