# testing-service
A Micro service useful for testing


## Endpoints

### /
root path, indicates it is running

```
curl localhost:8000
if you see this page, at least something is working
```

### /v0/hc

#### GET
returns hc status code, default `200`

```
curl localhost:8000/v0/hc
this is the hc response with status code: 200
```

#### POST
set hc status code

```
curl -X POST 'localhost:8000/v0/hc?status=500'
hc status code set to: 500
```

### /v0/fib
calculate fib sequence the slow way

```
time curl localhost:8000/v0/fib/45
1134903170
curl localhost:8000/v0/fib/45  0.00s user 0.00s system 0% cpu 6.506 total
```

### /v0/control
control what to do

```
time curl 'localhost:8000/v0/control?response_code=233&response_body=helloworld&response_delay=2s'
helloworld
curl   0.00s user 0.00s system 0% cpu 2.009 total
```

### /v0/reflect
prints what its given

```
curl localhost:8000/v0/reflect/hello
Method:
GET

Headers:
User-Agent: []string{"curl/7.68.0"}
Accept: []string{"*/*"}

Path:
/v0/reflect/hello

Query Params:

Body:
```
