

# Error Wrapping Utility

This utility provides two functions, `Wrap` and `WrapLog`, designed to enhance error handling and logging in Go applications.

## Function `Wrap`

The `Wrap` function is used to wrap an existing error with additional messages. It facilitates creating more informative error messages by combining custom messages with the original error message. The resulting error provides improved context and helps with error tracing.

### Usage

```go
newErr := Wrap(&originalErr, "Additional message 1", "Additional message 2")
```

If `originalErr` is not nil, `Wrap` will create a new error that includes the provided additional messages along with the original error message.

## Function `WrapLog`

The `WrapLog` function combines error wrapping with logging, making it useful for debugging and monitoring purposes. It logs the wrapped error along with request-related information and custom messages.

### Usage

```go
WrapLog(&originalErr, requestInfo, "Additional log message")
```

If `originalErr` is not nil, `WrapLog` will log a message that combines the additional log message, request information, and the original error message.

## Benefits

- **Enhanced Error Information**: By using `Wrap` and `WrapLog`, you can easily add contextual information to error messages, making it easier to identify the source of errors.
- **Simplified Debugging**: The functions simplify debugging by providing detailed error messages and logging information. This can significantly reduce the time it takes to identify and fix issues.
- **Effective Monitoring**: With `WrapLog`, you can create informative log entries for error situations, aiding in real-time monitoring and issue detection.
- **Readable and Maintainable Code**: By encapsulating error wrapping and logging logic, your codebase remains cleaner and easier to maintain.

## Caution

- Be mindful when using these functions. Overuse of excessive messages might clutter logs and make debugging more difficult.
- Always use relevant, concise, and helpful messages to ensure that the combined error messages provide meaningful context.

## Example

Here's an example demonstrating the use of the `Wrap` and `WrapLog` functions:

```go
func PerformOperation() error {
    err := SomeFunction()
    if err != nil {
        WrapLog(&err, requestDetails, "Error during operation")
        return Wrap(&err, "Failed to perform operation")
    }
    return nil
}
```

In this example, `WrapLog` logs the error with additional information, while `Wrap` combines the error with a custom error message for better error handling.

---

Feel free to adapt and expand this README.md to match your project's specific needs and context.

```

You can copy and paste this content into your project's README.md file and adjust it according to your actual project details and preferences.
```
