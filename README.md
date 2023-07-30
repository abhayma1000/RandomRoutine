# Workout Routine

This repo uses [MuscleWiki](https://musclewiki.com/) and returns the link to a randomized workout to do.

This takes the links specified in ```config.json``` and cyles through them. A ```SlidingWindow``` is specified which is how many different workouts until a certain workout can be repeated.

### Dependencies:

- Go

### Steps

- Download repo
- Run ```go mod tidy && go mod vendor```
- Run ```go run main.go```

