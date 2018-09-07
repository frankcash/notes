# Going Severless

[Kelsey Hightower](https://twitter.com/kelseyhightower)

August 29, 2018

### Functions

Some peolpe believe serverless is just writing functions, i.e. like CGI

Invoke and handle - move this into provider

[xinetd](https://en.wikipedia.org/wiki/Xinetd) made writing a daemon easier.

### Events

Needs to be invoked.

REST API: function per method/resource. API Gateway on the edge.

Data Pipelines:  Event occurs, process, and then return it

## In Go

AWS is the first to support it

Google Cloud is about to be there w/ go

https://stackoverflow.com/questions/52075778/what-does-a-production-ready-google-cloud-function-look-like

Use go vendor

Make sure to configure proper pooling when using a db connection.  It can be complicated to manage this in a serverless environment.

Can make use of multiple forms of triggers to make life easier.

Need to import tracing libraries and better logging tools.

Only one db connection per function call, so not to exhaust db connections.

Prototyping serverless becomes easy, and prototypes mirror prod a lot.

