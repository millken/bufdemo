```
buf registry login # rate limit 10/hour, if you have not logged in
buf config init
buf generate
go run .

curl \
    --header "Content-Type: application/json" \
    --data '{"pet_id": "Jane"}' http://127.0.0.1:8080/pet.v1.PetStoreService/PutPet

grpcurl \           
-protoset <(buf build -o -) -plaintext \
-d '{"pet_id": "Jane"}' \                   
localhost:8080 pet.v1.PetStoreService/PutPet #need in current directory
```

