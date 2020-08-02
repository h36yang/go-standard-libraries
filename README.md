# go-standard-libraries
Practice materials for "The Go Standard Library" course

### CLI Applications
* Leveraging Command Line Arguments with `os.Args` slice
* Adding Flags with `flag` package
* Reading Keyboard Inputs with `fmt.Scanf()` function
* Scanning Keyboard Inputs or Files with `bufio` package

### Using the Fmt Package
* Scanning Input with `fmt.Sscanf()` and `fmt.Scanln()` functions - generally not as useful as `bufio` scanner
* Formatting Output with `fmt.Printf()` function
* Manipulating Strings with `fmt.Sprintf()` function and special color formats
* Formatting other Data Types with different format specifiers - refer to https://golang.org/pkg/fmt for more details

### Using the Log Package
* Understanding Common Error Levels:
  * **Information**
    * If you want to confirm something
    * If you are logging transactions
    * If you are keeping track of something
    * If you want to show runtime information
  * **Warning**
    * If you need to get user's attention
    * If something is misconfigured
  * **Error**
    * If operation is compromised
    * If something very unexpected happened
  * **Fatal**
    * If the program must stop running
  * **Debug**
    * If you want to write out data or examine variables
* Formatting Log Output with `log` package
  * Log fatal messages and stop the program with `log.Fatal()`
  * Log other types of messages with `log.Println()`
  * Format messages with different flags - refer to https://golang.org/pkg/log for more details
* Utilizing the Trace Logger with `runtime/trace` package
  * Use `go tool trace <TRACE_FILE>.out` command to open and view the trace file

### Using the Time Package
* Understanding Wall Clock vs Monotonic Clock
  * **Wall Clock**
    * Used to keep track of the current time of day
    * Subject to variation
    * Great for humans
    * Not great for measuring time
  * **Monotonic Clock**
    * Used for measuring time
    * Not subject to variation
    * Only meaningful for the process calling it
    * Can be simulated by wall clock time
* Formatting Time Output with `time` package
  * Get current time with `time.Now()` function
  * Format date with `t.Format()` method
  * Refer to https://golang.org/pkg/time for more details
* Calculating Time Spans
  * Measure timespan using `time.Since()` and `time.Until()` functions
  * Obtain dates by adding/subtracting timespan using `t.AddDate()` and `t.Add()` methods
* Calculating Elapsed Time for Applications using `time.Since()` function
