# Random Routine

This repo is for people looking to generate a random action without repeating for a while and saving it all in a file

Inspiration comes from wanting to generate a workout based off of [MuscleWiki](https://musclewiki.com/).
The program returns a link to a random workout that I haven't done in a while.

This takes the links (options) specified in ```<project>.json``` and cyles through them. Then, logs them in ```<project>.csv```.

A ```SlidingWindow``` is specified which is how many different options until a certain option can be repeated.

### Dependencies:

- Go

### Steps

- Download repo
- Run ```go mod tidy && go mod vendor```
- Run ```go run main.go init <project>``` --> Initializes the project CSV and JSON files
- Populate ```<project>.json``` file with info (see example JSON config file below)
  - BaseUrl is a base url for all of the options. Option generated will append to this BaseUrl when outputted
  - Sliding window is number of different options before a repeat
- Run ```go run main.go exec <project>``` --> Generates random instance depending on JSON file and outputs to server. Then writes to CSV to not be repeated for SlidingWindow options

### Example
- ```go run main.go init workouts```
- Example JSON Config file:
```json
{"AllOptions":["chest","shoulders","traps","traps-middle","biceps","triceps","forearms","lats","obliques","lowerback","abdominals","glutes","quads","hamstrings","calves"],"BaseUrl":"https://musclewiki.com/exercises/male","SlidingWindow":10}
```
- ```go run main.go exec workouts```
- Output: [https://musclewiki.com/exercises/male/shoulders](https://musclewiki.com/exercises/male/shoulders)

