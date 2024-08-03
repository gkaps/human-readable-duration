import (
  "fmt" 
  "strings"
)

func FormatDuration(seconds int64) string {
  
  // Create a list of the words to be shown on the output
  words := []string{"year", "day", "hour", "minute", "second"}
  
  if seconds == 0 {
    return "now"
  }
  
  // Calculate years, days, hours, minutes and seconds values
  seconds = seconds % (100 * 365 * 24 * 3600)
  years := seconds / (365 * 24 * 3600)
  seconds = seconds % (365 * 24 * 3600)
  days := seconds / (24 * 3600)
  seconds = seconds % (24 * 3600)
  hours := seconds / 3600
  seconds = seconds % 3600
  minutes := seconds / 60
  seconds = seconds % 60
  
  // put all the values on an integer array
  time := []int64{years, days, hours, minutes, seconds}
  
  // create an empty duration string array
  duration := make([]string, 0)
  
  // Loop over time array and add the respective strings to the duration array
  for i:=0; i<len(time); i++ {
    if time[i] == 1 {
      duration = append(duration, fmt.Sprint(time[i]) + " " + words[i])
    } else if time[i] > 1 {
      duration = append(duration, fmt.Sprint(time[i]) + " " + words[i] + "s")
    }
  }
  
  // Check the scenarios and return the correct output
  if len(duration) == 1 {
    return duration[0]
  } else if len(duration) == 2 {
    return duration[0] + " and " + duration[1]
  } else {
    return strings.Join(duration[:len(duration)-1],", ") + " and " + duration[len(duration)-1]
  }
  
}
