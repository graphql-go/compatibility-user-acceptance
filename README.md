# Compatibility User Acceptance

Golang CLI library for **validating compatibility** of any GraphQL implementation DX metrics against the GraphQL reference implementation: [graphql-js](https://github.com/graphql/graphql-js).

Current implementation supports the following GraphQL implementations:
- [https://github.com/graphql-go/graphql](https://github.com/graphql-go/graphql)

## Use Cases

- Cross validation of compatibility between implementations leveraging the following DX metrics of the GitHub repository:
  - Stars.
  - Number of issues opened.
  - Number of issues closed.
  - Number of issues opened (Comments related to the GraphQL Compatibility, keywords: specification, standard and parity).
  - Number of issues closed(Comments related to the GraphQL Compatibility, keywords: specification, standard and parity).
  - Number of pull requests opened.
  - Number of pull requests closed.
  - Number of forks.
  - License.
  - Last commit date.
  - Number of contributors.
  - GraphQL specification version.
 
The difference ratio is calculated by dividing the (implementation - spec/implementation) and showed in percentage, and the result per item information is display as succeeded or failure whether or not it does not exceed the max difference ratio allowed.

Eg.

```
Example 1:
Specification License: MIT
Reference Implementation License: MIT
Difference Ratio: 0%
Max Difference Ratio Allowed: 0%
```

```
Example 2:
Number Of Stars: 100k - 100%
Reference Implementation Stars: 10k - 10%
Difference Ratio: 90%
Max Difference Ratio Allowed: 90%
```

It is useful to have the max difference ratio allowed because it helps to check per item level what is going on at that metric side and gives clear idea on what actions to take in order to make the reference implementation more user accepted.

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

Running the library in debug mode:

```
DEBUG=true ./bin/start.sh
```
