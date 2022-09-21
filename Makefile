all: build
build:
	rm -rf app/*
	cp -r static app
	go build -o exe .
	mv exe app/gym