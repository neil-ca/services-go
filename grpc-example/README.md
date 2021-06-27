When building distributed services, you’re communicating between the services
over a network. To send data (such as your structs) over a network, you need
to encode the data in a format to transmit, and lots of programmers choose
JSON. When you’re building public APIs or you’re creating a project where
you don’t control the clients, JSON makes sense because it’s accessible—both
for humans to read and computers to parse. But when you’re building private
APIs or building projects where you do control the clients, you can make use
of a mechanism for structuring and transmitting data that—compared to
JSON—makes you more productive and helps you create services that are
faster, have more features, and have fewer bugs.
So what is this mechanism? Protocol buffers (also known as protobuf), which
is Google’s language and platform-neutral extensible mechanism for structur-
ing and serializing data. The advantages of using protobuf are that it:

- Guarantees type-safety
