#!/bin/bash

echo "Copying all source files to public..."
cp python/**/*.py frontend/public
cp go/pkg/**/*.go frontend/public
cp ts/solutions/**/*.ts frontend/public

# echo "Building and deploying the Vue application..."
# cd frontend
# yarn build
# rsync --delete -azv -e "ssh -i ~/.ssh/mcculleydj.dev.pem" dist/ ubuntu@54.203.119.188:/home/ubuntu/rosetta-code
# ssh -i ~/.ssh/mcculleydj.dev.pem ubuntu@54.203.119.188 sudo systemctl restart nginx
# cd ..
