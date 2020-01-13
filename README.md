### Metric calculator
This is a CLI application with commands to calculate following metrics
- uptime
- allowed-downtime [TODO]


#### Installation
```
go get github.com/asheet-bhaskar/metric-calculator
```

#### Uptime calculation
##### Command Syntax
```
uptime <software name> <monitoring duration> <downtime duration>
```

##### Example
```
metric-calculator uptime prometheus 86400 300

Calculating the uptime
prometheus uptime is = 99.653
Nice, Improve it ❤
```
