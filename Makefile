proto:
	protoc --go_out=. --go_opt=paths=source_relative v1/person.proto
tag:
	protoc --go_out=. --go_opt=paths=source_relative --go_opt=tags="bson,json" v1/person.proto


clean:
	rm -f v1/patient.pb.go
