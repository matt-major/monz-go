# monzgo

## Usage
Create a new instance of the `Monzgo` client:

```go
m := monzgo.Monzgo{
        APIKey:  os.Getenv("MONZO_API_KEY"),
        BaseURL: "https://api.monzo.com/",
    }
```
