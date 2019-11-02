# CHANGELOG
All notable changes to this project will be documented in this file. This project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).
<a name="unreleased"></a>
## [Unreleased]


<a name="v0.7.3"></a>
## [v0.7.3] - 2019-10-29
### Bug Fixes
- **cmd:** enforce lowercase hash string (fixes [#67](https://github.com/vchain-us/vcn/issues/67))
- **cmd/inspect:** enforce lowercase hash string
- **cmd/serve:** enforce lowercase hash string
- **cmd/sign:** try again when notarization password is wrong

### Changes
- **cmd:** no echo when asking for confirmation before quitting


<a name="v0.7.2"></a>
## [v0.7.2] - 2019-10-24
### Bug Fixes
- **cmd/serve:** add missing event tracking

### Changes
- **ci:** add verify-commit action
- **ci:** github action for docs
- **cmd/sign:** return full output when notarizing

### Features
- store TX hash
- **cmd/inspect:** add ability to print only locally extracted info


<a name="v0.7.1"></a>
## [v0.7.1] - 2019-10-08
### Bug Fixes
- added ability to do not write on read only dir (fixes [#45](https://github.com/vchain-us/vcn/issues/45))
- **cmd:** typo and language improvement
- **cmd/serve:** CORS preflight (fixes [#60](https://github.com/vchain-us/vcn/issues/60))


<a name="v0.7.0"></a>
## [v0.7.0] - 2019-10-04
### Bug Fixes
- simplied error message when wrong notarization password is provided
- correct cross compile deps
- trust level was not shown when equals to DISABLED
- contentType was not shown when `ArtifactResponse` was not present
- make fields name consistent across output formats (json, yaml)
- **api:** correct tracing log label for `VerifyMatchingSignerIDWithFallback`
- **cmd/sign:** command must accept exactly one arg

### Changes
- update messages
- update dependencies
- platform managed secret is not stored locally anymore
- remove onboarding message (closes [#52](https://github.com/vchain-us/vcn/issues/52))
- use go v1.13 for the building system
- promote artifact fields to root
- move tracking call outside api pkg
- **api:** removed `BlockchainVerification.LevelName()` method
- **api:** add `BlockchainVerification.UnmarshalYAML`
- **api:** omit empy field in `ArtifactResponse`
- **cmd:** gracefully print error message if artifact is nil (fixes [#57](https://github.com/vchain-us/vcn/issues/57))
- **cmd:** added ability to use empty notarization password
- **cmd/info:** improved message for expired tokens
- **cmd/serve:** notarization API will allow only registered schemes for the kind field (closes [#51](https://github.com/vchain-us/vcn/issues/51))
- **extractor:** do not store version in metadata when empty
- **extractor/dir:** do not create `.vcnignore` on directory authentication (refs [#45](https://github.com/vchain-us/vcn/issues/45))
- **meta:** use Stringer interface for stage environment
- **meta:** use Stringer interface for status, level, and visibility

### Code Refactoring
- functional options for signing method
- improve secret handling API
- **api:** always check for user existence when login
- **cmd:** simplify cobra commands naming

### Features
- **api:** add `SignerID()` method
- **api:** added ability to upload the secret to the platform
- **api:** add functional options for signing
- **cmd:** improved help messages
- **cmd:** added `vcn info` command
- **cmd:** add `vcn serve` command
- **cmd:** added command for setting a custom notarization password (closes [#53](https://github.com/vchain-us/vcn/issues/53))
- **cmd/serve:** TLS support (closes [#48](https://github.com/vchain-us/vcn/issues/48))
- **cmd/serve:** print environment info at startup
- **cmd/serve:** allow to pass credentials via HTTP
- **extractor:** add functional options

### BREAKING CHANGE

`api.Sign` method signature has been changed.

`meta.StageName` has been removed; please use the `.String()` on value instead.

All command's factory methods have been renamed.

`BlockchainVerification.LevelName()` has been removed, please use `BlockchainVerification.Level.String()` instead.

`meta.StatusName`, `meta.LevelName`, `meta.VisibilityName` have been removed; please use the `.String()` on values instead.

`artifact` field is not present anymore in json/yaml results.

Extractors method signature have changed to accomodate functional options.

Secret storage is not used anymore for platform managed secrets.
`store.PublicAddress()` has been renamed to `store.SignerIDFromSecret()`, and will return a value only when a local secret is available


<a name="v0.6.3"></a>
## [v0.6.3] - 2019-09-25
### Bug Fixes
- **extractor/dir:** runtime error when passing a regular file instead of a directory (fixes [#56](https://github.com/vchain-us/vcn/issues/56))


<a name="v0.6.2"></a>
## [v0.6.2] - 2019-09-17
### Changes
- new AssetsRelay smart contract

### Code Refactoring
- use new smart contract functions


<a name="v0.6.1"></a>
## [v0.6.1] - 2019-09-10
### Bug Fixes
- correct error msg when entering empty email (fixes [#43](https://github.com/vchain-us/vcn/issues/43))

### Features
- enable git repo notarization and authentication
- **extractor:** add scheme for git repository


<a name="v0.6.0"></a>
## [v0.6.0] - 2019-09-04
### Bug Fixes
- correct `Signer` field semantic
- **api:** correct secret download content format
- **verify:** switch to single key for current user

### Changes
- update Smart Contract addresses to latest version
- removed profile migration from v0.3.x
- deprecate `KEYSTORE_PASSWORD` env var in favour of `VCN_NOTARIZATION_PASSWORD`
- add "your assets will not be uploaded" message
- deprecate `--key` in favor of `--signerID`
- improve API and user messages to reflect the new terminology
- correct typos
- **api:** refactor to single address and code cleanup
- **help:** add info about assets and env vars
- **meta:** removed event tracking for keystore creation
- **sign:** dropping support for multiple keys
- **store:** switch from multi-key to single secret config
- **terminology:** switch to `notarize` and `authenticate`

### Code Refactoring
- gofmt and golint
- reduce artifact API surface
- **meta:** rationalize config functions

### Features
- add trial expiration message
- automatically get secret from platform at login
- add automatic checking for newer versions

### BREAKING CHANGE

Env variable `KEYSTORE_PASSWORD` has been renamed to `VCN_NOTARIZATION_PASSWORD`.

All `api.BlockChainVerify*()` funcs have been renamed to `Verify*()`

`api.BlockchainVerification.Key()` has been renamed to `SignerID()`
`api.BlockChainOrganisation.MembersKeys()` has been renamed to `MembersIDs()`
`api.BlockChainGetOrganisation()` has been renamed to `GetBlockChainOrganisation()`

`meta`'s endpoint funcs has been removed in favour of single `meta.APIEndpoint()`. `meta.MainNetEndpoint()` has been renamed to `meta.MainNet()`.

Artifact field `Signer` will not contain the user pub key anymore.

`api.LoadArtifactForHash` has been renamed to `api.LoadArtifact`.
`api.ArtifactRequest` has been privatized. Artifact shall be created only thru the notarization process.

`--key` has been removed from `vcn sign`

Dropping `LoadAllArtifacts` and `LoadArtifacts` and wallets/keys/keystore related functions from `api.User`.
`api.Sign` won't accept anymore a pub key as parameter.

support for multiple keystores in config file and related APIs within the `store` have been removed.


<a name="v0.5.4"></a>
## [v0.5.4] - 2019-07-30
### Features
- **extractor:** add support for podman


<a name="v0.5.3"></a>
## [v0.5.3] - 2019-07-19
### Bug Fixes
- **build:** correct NSIS installation directory


<a name="v0.5.2"></a>
## [v0.5.2] - 2019-07-17
### Bug Fixes
- update go.sum for xgo
- **bundle:** use uint64 for size
- **dir:** OS agnostic paths
- **list:** show all assets with pagination (fixes [#28](https://github.com/vchain-us/vcn/issues/28))

### Changes
- **api:** add Artifact.Copy()
- **bundle:** return error if distinct sizes are found for the same digest
- **bundle:** descriptor test
- **bundle:** rework diff
- **cmd:** always show local extracted metadata
- **dir/extractor:** ignore irregular files

### Features
- directory signing and verify with manifest
- **api:** list assets grouped by hash with pagination
- **bundle:** enforce manifest specs
- **bundle:** arrange multiple items in a Merkle Directed Acyclic Graph
- **bundle:** diff
- **extractor:** scheme for directories
- **extractor/dir:** default ignore file
- **extractor/dir:** .vcnignore file support
- **verify/dir:** automatically check manifest integrity


<a name="v0.5.1"></a>
## [v0.5.1] - 2019-07-02
### Bug Fixes
- `BlockchainVerification` json unmarshalling

### Changes
- **cmd:** update cobra to 0.5.0
- **verify:** add explanatory output messages
- **verify:** minor cmd usage improvements

### Features
- **inspect:** add new `vcn inspect` command
- **sign:** `VCN_KEY` env variable for signing


<a name="v0.5.0"></a>
## [v0.5.0] - 2019-06-25

<a name="0.5.0"></a>
## [0.5.0] - 2019-06-25
### Bug Fixes
- ca-certificates for Dockerfile
- clean other context when user logs in
- **cmd:** show config file path if not default one
- **docker:** accept images but not other objects
- **extractor:** correct empty files handling

### Changes
- reduce public api surface
- code cleanup and fix comments
- bump version
- **api:** move type Error to its own source file
- **api:** deprecate public publisher APIs
- **api:** deprecate publisher fields
- **api:** allow empty size
- **cmd:** internal types
- **env:** remove VERIFY_ prefix
- **internal:** move cli package
- **output:** silence printing when formatted output
- **verify:** improve error labels

### Features
- autogenerated changelog
- vcn user agent
- promote --output to global flag and return formatted errors accordingly
- json and yaml output for sign
- **api:** add blockchain organisation support
- **list:** support for --output json/yaml
- **sign:** signing by --hash
- **verify:** add --org for verify

### BREAKING CHANGE

some methods are now private, tracking functions are now deprecated.

Some publisher and auth related methods and structs have been privatized and will removed in future

ArtifactResponse fields (related to the publisher) as been renamed as following:
```
Publisher -> Signer
PublisherCompany -> Company
PublisherWebsiteUrl -> Website
PublisherCount -> SignerCount
```
ArtifactRequest's `Url` has been fixed to `URL` too.

meta.VcnClientName() has been removed in favor of meta.UserAgent()

`VCN_VERIFY_KEYS` has been removed in favour of `VCN_KEY` with the same functionality.


<a name="0.4.3"></a>
## [0.4.3] - 2019-05-21
### Bug Fixes
- **api:** nil pointer dereference
- **extractor:** nil pointer dereference
- **sign:** flag accessed but not defined: hash
- **sign:** cross platform loading spinner

### Changes
- minor refactoring of login/logout
- generate markdown docs for commands
- **api:** load user's artifact by hash only
- **api:** size must be valid, name can be empty
- **sign:** prompt login if needed when started by explorer

### Features
- profile dir per stage
- add test env
- **sign:** untrust/unsupport by --hash flag


<a name="0.4.2"></a>
## [0.4.2] - 2019-05-14
### Bug Fixes
- print correct config file name
- do not read in env var for config
- **sign:** remove ownership disclaimer
- **verify:** cross platform coloured printing
- **verify:** only print size when available

### Changes
- disable config file message
- reorg publisher fields
- minor printing and marshaling improvements
- **api:** fine-tune blockchain verify funcs
- **build:** make makefile more resilient
- **log:** improve metahash tracing
- **output:** correct WriteTo interface

### Code Refactoring
- **sign:** printing

### Features
- improve printing and add yaml output format
- static build and docker cli support in dockerfile
- **verify:** with multiple key by VCN_VERIFY_KEYS env var
- **verify:** allow multiple --key flags
- **verify:** show asset visibility


<a name="0.4.1"></a>
## [0.4.1] - 2019-05-08
### Bug Fixes
- correct err msg when token has expired
- do not print verification error message if --output=json
- verify fallback
- clear context when logging out
- if logged in always show user own signature
- remove test-resources from dockerignore
- return proper error by checking quota before

### Changes
- executable file sniffer
- go mod tidy
- change metadata naming
- add Platform and Architecture fields
- **build:** minor improvements
- **ci:** initial circleci setup
- **cmd:** use extractors
- **meta:** increse signing timeout
- **meta:** increase tx verification rounds

### Code Refactoring
- asset and metadata fields naming

### Features
- user defined attributes for assets
- json output format (--output=json)
- infer asset version
- get metadata from docker image
- add --hash for verify
- improve mime and metadata extractor for executables
- new kind and mimeType attributes for assets
- modularize metadata extractors
- URI parser for assets
- subscription limit enforcement
- **build:** improved build system


<a name="0.4.0"></a>
## [0.4.0] - 2019-04-30
### Bug Fixes
- dashboard cmd has no args
- prefix for docker assets
- find home dir in the right way
- do not exit early when login is required
- do not quit when executed by Win context menu
- ensure default keystore before looking for keys
- remove log.Fatal in favour of returing errors
- gitignore
- ask for confirmation before quitting
- token deletion when already logged out
- improve err msgs when auth is required
- add missing header comments
- temporary dirs creation at startup
- **login:** silence usage message when errored
- **migrate:** do not create a new key if it already exists
- **store:** correct key funcs and tests
- **tracking:** send events at right place in time

### Changes
- use two-stage build
- make createArtifact private
- printing stuff
- no race test for make install
- add dev stage
- move cli package to internal
- improve naming
- code cleanup and minor improvements
- add store context and key utils
- introduce store package
- move logs to internal
- improve wallet not synced message
- correct printing func
- update gitignore
- switch main to cobra
- partial api logging refactor
- code clean up
- introducing Cobra
- move file hashing func to internal
- use make install for Dockerfile
- **build:** improve makefile
- **cli:** spinner changed and code cleaned up
- **dist:** omit symbol table and debug info
- **internal:** correct errors funcs and tests

### Code Refactoring
- login cmd
- list cmd
- untrust and unsupport cmds
- sign, dashboard cmds and APIs
- docker compose
- package oriented design

### Features
- key flag for sign and verify
- profile migration from v0.3.x
- config file and multi-keys support
- logout command
- upgrade to latest assetsrelay sc
- provide Dockerfile
- **build:** build system with Makefile
- **cli:** improved column printing

### BREAKING CHANGE

this commit changes the usage of verify and sign methods

this commit introduce the config file with multi-keys support, and a huge refactoring of vcn code.


<a name="0.3.6"></a>
## [0.3.6] - 2019-04-08

<a name="0.3.5"></a>
## [0.3.5] - 2019-03-28

<a name="0.3.4"></a>
## [0.3.4] - 2019-03-20

<a name="0.3.3"></a>
## [0.3.3] - 2019-03-13

<a name="0.3.2"></a>
## [0.3.2] - 2019-03-11

<a name="0.3.1"></a>
## [0.3.1] - 2019-03-11

<a name="0.3.0"></a>
## [0.3.0] - 2019-03-08

<a name="0.2.2"></a>
## [0.2.2] - 2019-03-07

<a name="v.0.2-beta.0"></a>
## [v.0.2-beta.0] - 2019-02-25

<a name="v.0.1-beta.2"></a>
## v.0.1-beta.2 - 2019-02-19

[Unreleased]: https://github.com/vchain-us/vcn/compare/v0.7.3...HEAD
[v0.7.3]: https://github.com/vchain-us/vcn/compare/v0.7.2...v0.7.3
[v0.7.2]: https://github.com/vchain-us/vcn/compare/v0.7.1...v0.7.2
[v0.7.1]: https://github.com/vchain-us/vcn/compare/v0.7.0...v0.7.1
[v0.7.0]: https://github.com/vchain-us/vcn/compare/v0.6.3...v0.7.0
[v0.6.3]: https://github.com/vchain-us/vcn/compare/v0.6.2...v0.6.3
[v0.6.2]: https://github.com/vchain-us/vcn/compare/v0.6.1...v0.6.2
[v0.6.1]: https://github.com/vchain-us/vcn/compare/v0.6.0...v0.6.1
[v0.6.0]: https://github.com/vchain-us/vcn/compare/v0.5.4...v0.6.0
[v0.5.4]: https://github.com/vchain-us/vcn/compare/v0.5.3...v0.5.4
[v0.5.3]: https://github.com/vchain-us/vcn/compare/v0.5.2...v0.5.3
[v0.5.2]: https://github.com/vchain-us/vcn/compare/v0.5.1...v0.5.2
[v0.5.1]: https://github.com/vchain-us/vcn/compare/v0.5.0...v0.5.1
[v0.5.0]: https://github.com/vchain-us/vcn/compare/0.5.0...v0.5.0
[0.5.0]: https://github.com/vchain-us/vcn/compare/0.4.3...0.5.0
[0.4.3]: https://github.com/vchain-us/vcn/compare/0.4.2...0.4.3
[0.4.2]: https://github.com/vchain-us/vcn/compare/0.4.1...0.4.2
[0.4.1]: https://github.com/vchain-us/vcn/compare/0.4.0...0.4.1
[0.4.0]: https://github.com/vchain-us/vcn/compare/0.3.6...0.4.0
[0.3.6]: https://github.com/vchain-us/vcn/compare/0.3.5...0.3.6
[0.3.5]: https://github.com/vchain-us/vcn/compare/0.3.4...0.3.5
[0.3.4]: https://github.com/vchain-us/vcn/compare/0.3.3...0.3.4
[0.3.3]: https://github.com/vchain-us/vcn/compare/0.3.2...0.3.3
[0.3.2]: https://github.com/vchain-us/vcn/compare/0.3.1...0.3.2
[0.3.1]: https://github.com/vchain-us/vcn/compare/0.3.0...0.3.1
[0.3.0]: https://github.com/vchain-us/vcn/compare/0.2.2...0.3.0
[0.2.2]: https://github.com/vchain-us/vcn/compare/v.0.2-beta.0...0.2.2
[v.0.2-beta.0]: https://github.com/vchain-us/vcn/compare/v.0.1-beta.2...v.0.2-beta.0
