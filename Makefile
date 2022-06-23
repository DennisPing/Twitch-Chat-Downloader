all:
	@go build -o main && echo Built: main

clean:
	rm ./main