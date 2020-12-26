from pathlib import Path
import shutil

dest_path = Path('../frontend/public')
problem_path = Path('../problems')
python_path = Path('../python')
go_path = Path('../go/pkg')
ts_path = Path('../ts/solutions')

problems = []

for p in problem_path.glob('**/*.md'):
    shutil.copy(p, dest_path.joinpath(p.name + '.txt'))
    problems.append(p.stem)

for p in python_path.glob('**/*.py'):
    shutil.copy(p, dest_path.joinpath(p.name + '.txt'))

for p in go_path.glob('**/*.go'):
    shutil.copy(p, dest_path.joinpath(p.name + '.txt'))

for p in ts_path.glob('**/*.ts'):
    shutil.copy(p, dest_path.joinpath(p.name + '.txt'))


with open('../frontend/public/problem-list.txt', 'w') as f:
    for p in sorted(problems):
        f.write(p + '\n')
