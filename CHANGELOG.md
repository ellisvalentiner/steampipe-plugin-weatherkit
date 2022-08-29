<a name="unreleased"></a>
## [Unreleased]


<a name="v0.0.1"></a>
## v0.0.1 - 2022-08-29
### Chore
- update CHANGELOG and tag for release
- update CHANGELOG

### Docs
- how to configure credentials using env vars
- add CHANGELOG.md

### Feat
- add support for setting config using env vars [#14](https://github.com/ellisvalentiner/steampipe-plugin-weatherkit/issues/14)
- add support for using a pre-generated JWT for authorization

### Fix
- interfaceToColumnValue failed for column 'moonset'
- missing config panics too soon [#15](https://github.com/ellisvalentiner/steampipe-plugin-weatherkit/issues/15)
- more helpful error handling for invalid credentials
- empty control flow branch
- **docs:** reformat queries [#16](https://github.com/ellisvalentiner/steampipe-plugin-weatherkit/issues/16)
- **docs:** token missing docs/index.md [#13](https://github.com/ellisvalentiner/steampipe-plugin-weatherkit/issues/13)
- **docs:** update slack channel link [#12](https://github.com/ellisvalentiner/steampipe-plugin-weatherkit/issues/12)

### Refactor
- use a common method for all Weather data set requests

### Pull Requests
- Merge pull request [#21](https://github.com/ellisvalentiner/steampipe-plugin-weatherkit/issues/21) from ellisvalentiner/docs/configuration-using-env-vars
- Merge pull request [#19](https://github.com/ellisvalentiner/steampipe-plugin-weatherkit/issues/19) from ellisvalentiner/refactor/common-method-for-weather-requests
- Merge pull request [#18](https://github.com/ellisvalentiner/steampipe-plugin-weatherkit/issues/18) from ellisvalentiner/fix/interfaceToColumnValue-failed
- Merge pull request [#17](https://github.com/ellisvalentiner/steampipe-plugin-weatherkit/issues/17) from ellisvalentiner/release/v0.1.0
- Merge pull request [#9](https://github.com/ellisvalentiner/steampipe-plugin-weatherkit/issues/9) from ellisvalentiner/fix/improve-bad-creds-messaging
- Merge pull request [#8](https://github.com/ellisvalentiner/steampipe-plugin-weatherkit/issues/8) from ellisvalentiner/feat/add-pregenerated-jwt-support
- Merge pull request [#7](https://github.com/ellisvalentiner/steampipe-plugin-weatherkit/issues/7) from ellisvalentiner/dependabot/go_modules/github.com/turbot/steampipe-plugin-sdk/v3-3.3.2


[Unreleased]: https://github.com/ellisvalentiner/steampipe-plugin-weatherkit/compare/v0.0.1...HEAD
