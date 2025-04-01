# Compatibility User Acceptance

Golang CLI library for **validating compatibility** of any GraphQL implementation UX metrics against the GraphQL reference implementation: [graphql-js](https://github.com/graphql/graphql-js).

Current implementation supports the following GraphQL implementations:
- [https://github.com/graphql-go/graphql](https://github.com/graphql-go/graphql)

## Use Cases

- Cross validation of compatibility between implementations of the following UX metrics:
  - GitHub Stars.
  - GitHub Issues (Open/Closed).
  - GitHub Issues (Open/Closed, comments that are related to the specification).
  - GitHub Pull Requests (open/closed).
  - GitHub Forks.
- The end result compares the implementation and the reference in terms of percentage consolidating into a single number that is wired to a % that is the base line for considering that it passes or not the check. 


### Contribute Back

Friendly reminder links are available in case you would like to contribute back into our commitment with Go and open-source.

| Author        |  PayPal Link  |
|:-------------:|:-------------:|
| [Chris Ramón](https://github.com/chris-ramon) | https://www.paypal.com/donate/?hosted_button_id=WHUQQYEMTRQBJ |

## Quick Start

Running the library:

```
./bin/start.sh
```
