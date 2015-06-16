# Wrecker

Wrecker is a request builder for JSON based APIs, written in Go.

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
}
```

There are two values you'll have to set

- `BaseURL`: This is the base URL of the API you will be using
- `HttpClient`: This is an instance of an http.Client.  Feel free to tweak this as much as you would like.

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
