  # wezario

**Simple command-line tool to get current weather** 

[![asciicast](https://asciinema.org/a/7ZRUc8GnvQFU1QG4JxDMHAtbj.png)](https://asciinema.org/a/7ZRUc8GnvQFU1QG4JxDMHAtbj?autoplay=1)

## Usage

create your own `.env` file from `.env.example` with `SERVICE_KEY` value set to your openweathermap API key

run included docker-compose with redis `docker-compose ru -d`

If you run simple `./main` you'll get the weather info for Moscow region
But with `--city` flag you can specify which city you want to get, e.g:

```
./main --city London
```

Application at first time would get info from openweathermap and store it to redis with city name as a key
TTL for that cache is 15 minutes and within those 15 minutes next requests would be addressed to redis

  ![Image of Gopher with umbrella](https://phillipsoft.com/images/posts/alexa-golang/umbrella.svg)
