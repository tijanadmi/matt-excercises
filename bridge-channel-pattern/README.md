#Bridge-Channel Pattern

The bridge-channel pattern refers to connect or combine multiple channels into a single channel. It allows you to merge the streams of data from different channels into a unified stream, making it easier to consume and process the combined data.In this case, the bridge channel acts as a connector between multiple channels, enabling the flow of data between them.
Here's how it would work:
To implement a bridge channel in Go, you can create a function that takes multiple input channels and returns a single output channel.
The bridge function is called with sensorData1 and sensorData2 as arguments, which bridges these two input channels into a single output channel.

The code sets up two channels to send sensor data, bridges these channels into a single output channel, and consumes the values from the output channel concurrently using goroutines. The use of channels and goroutines allows for concurrent communication and processing of sensor data.

