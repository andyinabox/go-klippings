GIT_UP_TO_DATE="Already up to date."

cd ./web/
  PULL=$(git pull origin master)
  if [ "$PULL" = "$GIT_UP_TO_DATE" ]
  then
    echo "UI repo is up-to-date"
  else
    cp .env.example .env
    yarn install
    # go test # need to make this safe
    yarn build
  fi
cd ..

go test ./...
go run ./cmd/klippings/ -gui