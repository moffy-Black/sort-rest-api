# heapsort-server

go で作られた受け取った配列を heap sort して返すサーバ

## modulename

```
github.com/moffy-Black/sort-server
```

## server sample test snippet

```
curl http://localhost:8080/ \                                               ✔  16:19:42 
--header "Content-Type: application/json" \
--include \
--request "POST" \
--data '{"sortType": "quicksort", "array": [0,1,9,2,8,3,7,4,6,5]}'
```
