# Over Engineering with Macaroons

[Ms. Tess Rinearson](https://twitter.com/_tessr)

August 28, 2018

[Google Paper](https://static.googleusercontent.com/media/research.google.com/en//pubs/archive/41892.pdf)
[Go Implementation](https://github.com/go-macaroon/macaroon)

- Created an ACL
- Capability vs Identity
- Tokens can have too much power in just a capability world
- Tokens relying on identity can be dangerous
- ref ["ACLs Don't"](http://waterken.sourceforge.net/aclsdont/current.pdf)
- CSRF is caused by identity problems
- Can add safety by adding capability to an identity model
- But, there is a middle ground: keys that work with identity
- Macaroons are contextual bearer tokens

#### Problems w/ Basic Auth

- Username/passw on every req.

#### HMAC

Can be used instead of username / passw

https://security.stackexchange.com/questions/20129/how-and-when-do-i-use-hmac/20301

#### More on Macaroons 

Macaroons are cookies with layers
1. hmac(human emoji, key emokji)
2. hmac("perms:read-only", 1)

Macaroons eventually look like OAuth process w/ fewer requests

### Macaroons: The Bad Parts

1. New availability dependency on Macaroon discharger (Dashboard in this case)
2. Difficult to use locally - Adds depndency on Macaroon system to all be running
3. Confusing behavior for users - As permissions change
4. Revoking a credential took a decent amount of time (discharge macaroon is valid for 5 minutes).  They had to add a check w/ dashboard for disabled macaroon
5. Macaroon had a tricky format as they are base64 and it was hard for some formatting errors (lol?)
6. Macaroon was maybe misnamed

Eventually the team removed macaroons (took 5 months).  Had to phase out macaroons and convert them to their basic auth ++ .

Macaroons are exciting technology, but the tech is pretty sparse with others using them.  Bad if you want to use boring technology.