# Scanpath

A simple path scanning utility written in Go. Detects the OS platform to determine the correct [scan function](./internal/scan) to use.

## Usage

*See full usage with `scanpath --help`*

By default, `scanpath` scans your current working directory (i.e. `.` or `$PWD`). You can change the path to scan with `-p path/to/scan` or `--scan-path path/to/scan`. Below are a list of additional/optional flags:

| Flag               | Description                                                                                                        | Default                 |
| ------------------ | ------------------------------------------------------------------------------------------------------------------ | ----------------------- |
| `-p`/`--scan-path` | Tell the app which path to scan.                                                                                   | `.`                     |
| `-l`/`--limit`     | Limit the number of results outputted to the terminal.                                                             | `0` (print all results) |
| `-s`/`--sort-name` | Column name to sort. Run the script once without defining a sort name to see all columns.                          | `name`                  |
| `-o`/`--order`     | `asc` or `desc` for ascending/descending sort.                                                                     | `asc`                   |
| `-f`/`--filter`    | A string defining your sort, i.e. `-s 'created >2022-01-01`, `--filter 'name ~ namePart` or `--filter 'name *part` | `""`/`nil`              |
