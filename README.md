# Wrecker

Wrecker is a request builder for JSON based APIs, written in Go.

[![Build Status](https://travis-ci.org/BrandonRomano/wrecker.svg?branch=master)](https://travis-ci.org/BrandonRomano/wrecker)

## Fluent Builder

Without digging too far into anything, here is what your requests will look like with Wrecker:

```go
wreckerClient.Get("/users").
    WithParam("id", "1").
    Into(&user).            // Loads the response into the user struct
    Execute()
```

## Usage

### Wrecker Instance

To get started with Wrecker, you first have to prepare an instance for yourself:

```go
wreckerClient = &wrecker.Wrecker{
    BaseURL: "http://localhost:5000",
    HttpClient: &http.Client{
        Timeout: 10 * time.Second,
    },
    DefaultContentType: "application/x-www-form-urlencoded",
}
```

There are three values you'll have to set

- `BaseURL`: This is the base URL of the API you will be using
- `HttpClient`: This is an instance of an http.Client.  Feel free to tweak this as much as you would like.
- `DefaultContentType`: This is the value of the `Content-Type` header that will be added to every request.  This can be overridden at request level, but for convenience you can set a default.

You're only going to have to do this once, but if you still find it's too verbose or you don't need the flexibility you can do this:

```go
wreckerClient = wrecker.New("http://localhost:" + os.Getenv("PORT"))
```

This way will default the `HttpClient` + `DefaultContentType` to the values set in the first example.

### Creating Requests

Let's say that you wanted to hit an endpoint with a response that looked something like this:

```json
{
    "id": 1,
    "user_name": "BrandonRomano",
    "location": "Brooklyn, NY"
}
```

The first thing you would do is create a model that would represent the response:

```go
type User struct {
    Id       int    `json:"id"`
    UserName string `json:"user_name"`
    Location string `json:"location"`
}
```

Now that we have our struct and our instance of `Wrecker`, we're ready to actually create the response.

```go
user := User{}

err := wreckerClient.Get("/users").
    WithParam("id", "1").
    Into(&user).
    Execute()

fmt.Println(strconv.Itoa(user.Id)) // Prints 1
fmt.Println(user.UserName) // Prints "BrandonRomano"
fmt.Println(user.Location) // Prints "Brooklyn, NY"
```

## License

Wrecker is licensed under [MIT](LICENSE.md)
