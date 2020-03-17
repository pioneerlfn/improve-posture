In other words, a program needs to fulfill only two criteria to be considered a web app:
- The program must return HTML to a calling client that renders HTML and dis- plays to a user.
- The data must be transported to the client through HTTP.

As an extension of this definition, if a program doesn’t render HTML to a user but instead returns data in any other format to another program, it is a **web service** (that is, it provides a service to other programs). 

## HTTP动词

| Method | Meaning | Safe | Idempotent |
|--- | --- | --- | --- |
| GET | Tells the server to return the specified resource. | YES | YES |
| HEAD | The same as GET except that the server must not return a message body. This method is often used to get the response headers without carrying the weight of the rest of the message body over the network. | YES | YES |
| OPTIONS | Tells the server to return a list of HTTP methods that the server sup-ports. | YES | YES |
| TRACE | Tells the server to return the request. This way, the client can see what the intermediate servers did to the request. | YES | YES |
| POST | Tells the server that the data in the message body should be passed to the resource identified by the URI. What the server does with the message body is up to the server. | **NO** | **NO** |
| PUT | Tells the server that the data in the message body should be the resource at the given URI. If data already exists at the resource identified by the URI, that data is replaced. Otherwise, a new resource is created at the place where the URI is. | NO | YES |
| DELETE | Tells the server to remove the resource identified by the URI. | NO | YES |
| CONNECT | Tells the server to set up a network connection with the client. This method is used mostly for setting up SSL tunneling (to enable HTTPS). |  |  |
| PATCH | Tells the server that the data in the message body modifies the resource identified by the URI. | | | 

