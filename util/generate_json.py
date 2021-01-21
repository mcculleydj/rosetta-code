from pathlib import Path
import os
import shutil
import json

dest_path = Path('../frontend/public')
problem_path = Path('../problems')
python_path = Path('../python')
go_path = Path('../go/pkg')
ts_path = Path('../ts/solutions')

name_to_category = {
    'strings': 'Strings',
    'linked': 'Linked Lists',
    'stacks': 'Stacks',
    'queues': 'Queues',
    'matrices': 'Matrices',
    'trees': 'Trees',
    'graphs': 'Graphs',
    'files': 'Files',
    'bits': 'Bits',
}

problems_dict = {}

for p in sorted(problem_path.glob('**/*.md')):
    with open(p, 'r') as f:
        lines = list(f.readlines())

    name = p.stem.split('_')[0]
    number = p.stem.split('_')[1]
    title = lines[0].replace('# ', '').strip()
    description = lines[2].strip()

    if problems_dict.get(name_to_category[name]):
        problems_dict[name_to_category[name]].append({
            'stem': p.stem,
            'number': number,
            'title': title,
            'description': description,
        })
    else:
        problems_dict[name_to_category[name]] = [{
            'stem': p.stem,
            'number': number,
            'title': title,
            'description': description,
        }]

for category, problems in problems_dict.items():
    for p in problems:
        dir_name = p['stem'].split('_')[0]
        with open(python_path.joinpath(dir_name, p['stem'] + '.py'), 'r') as f:
            p['python'] = f.read()
        with open(go_path.joinpath(dir_name, p['stem'] + '.go'), 'r') as f:
            p['go'] = f.read()
        with open(ts_path.joinpath(dir_name, p['stem'] + '.ts'), 'r') as f:
            p['ts'] = f.read()

with open(dest_path.joinpath('problems.json'), 'w') as f:
    json.dump(problems_dict, f, sort_keys=True, indent=4)
