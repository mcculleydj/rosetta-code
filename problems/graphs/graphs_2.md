# Build Order

Given a list of projects and a list of dependencies find a valid build order. If there is no valid build order return an error.

Example input / output where (a, d) means a depends on d:

- projects: a, b, c, d, e, f
- dependencies: (a, d), (f, b), (b, d), (f, a), (d, c)
- output: f, e, a, b, d, c
