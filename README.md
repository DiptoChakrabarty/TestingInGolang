# TestingInGolang
Learning how to perform unit tests in golang

go test -cover .

 go test -coverprofile=coverage.out

go tool cover -html=coverage.out

 alias coverage="go test -coverprofile=coverage.out && go tool cover -html=coverage.out"