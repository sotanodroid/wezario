  # wezario

**Telegram chat bot to get current weather info** 

Application at first time would get info from openweathermap and store it to redis with city name as a key. 
TTL for that cache is 15 minutes and within those 15 minutes next requests would be addressed to redis

  ![Image of Gopher with umbrella](https://phillipsoft.com/images/posts/alexa-golang/umbrella.svg)
