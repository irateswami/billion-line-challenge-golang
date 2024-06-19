# billion-line-challenge-golang
My attempt(s) at the billion line challenge using golang 

# Assumptions
There are some assumptions I have to make about the data, for the sake of my 
sanity.
1. Temp data is realistic, this means it will fall in the typical range covered 
by an signed 8 bit integer + a decimal. 
2. The cities are normally distributed, as in all cities appear in the data a 
roughly equal number of times
3. The data is formatted correctly and won't require checks along the way. 

# My plan
## Main runner thread
Thread that will read in line by line and then place each line onto a 
byLineChannel

## Broker thread
Thread that will consume the byLineChannel and then broker on which thread the 
line should be sent. 

Each line put over the channel will have the appearance of:

```
Hamburg;12.0
Bulawayo;8.9
Palembang;38.8
St. John's;15.2
Cracow;12.6
Bridgetown;26.9
Istanbul;6.2
Roseau;34.4
Conakry;31.2
Istanbul;23.0
```

Once the broker receives the line, it will placed into the broker worker pool. 

Each worker will:
1. Parse the line into two objects, a string and an 32 bit integer
2. Determine which city to place the data onto
3. Place the data onto the appropriate channel for the city
