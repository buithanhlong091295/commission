THIS_BRANCH=$(git rev-parse --abbrev-ref HEAD)

docker build -t xtekexchange/commission-service:$THIS_BRANCH .

docker push xtekexchange/commission-service:$THIS_BRANCH