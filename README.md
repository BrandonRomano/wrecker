# Wrecker

Wrecker is a request builder for JSON based APIs, written in Go.

> Please note that the API is being actively developed, so please expect breaking changes until we post a 1.0 release.  If you would like to use wrecker now, we suggest creating a fork and using that.

[![Build Status](https://travis-ci.org/BrandonRomano/wrecker.svg?branch=master)](https://travis-ci.org/BrandonRomano/wrecker)

## Fluent Builder

Without digging too far into anything, here is what your requests will look like with Wrecker:

```go
wreckerClient.Get("/users").
    URLParam("id", "1").
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
    RequestInterceptor: func(req *wrecker.Request) error {
        req.URLParam("id", "1")
        return nil
    },
}
```

There are three values you'll have to set

- `BaseURL`: This is the base URL of the API you will be using
- `HttpClient`: This is an instance of an http.Client.  Feel free to tweak this as much as you would like.
- `DefaultContentType`: This is the value of the `Content-Type` header that will be added to every request.  This can be overridden at request level, but for convenience you can set a default.
- `RequestInterceptor`: This is an interceptor that can update every request before it gets sent out.  This can safely be set to nil if you do not have a need for a request interceptor.

You're only going to have to do this once, but if you still find it's too verbose or you don't need the flexibility you can do this:

```go
wreckerClient = wrecker.New("http://localhost:" + os.Getenv("PORT"))
```

This way will default the `HttpClient` + `DefaultContentType` to the values set in the first example, and `RequestInterceptor` will be set to nil.

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

response, err := wreckerClient.Get("/users").
    URLParam("id", "1").
    Into(&user).
    Execute()

fmt.Println(strconv.Itoa(user.Id)) // Prints 1
fmt.Println(user.UserName) // Prints "BrandonRomano"
fmt.Println(user.Location) // Prints "Brooklyn, NY"
```

## License

Wrecker is licensed under [MIT](LICENSE.md)
