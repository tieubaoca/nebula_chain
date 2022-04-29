# Release Mechanism Interface

As mentioned in [launchpad](/x/launchpad/README.md), release mechanism is an interface required by a launchpad project.

For maximum flexibility, release mechanism interface should strive to describe as broad of a release mechanism as possible.

# Structure overview
1. asset_type: there can be many asset types in the future. For now, it will be token. NFT will be supported later.
2. total_amount: total amount of assets for release.
3. amount_released: amount of tokens already released. Project ends when __total_amount__ - __amount_released__ = 0.

# Common Interface
# Implementation
1. [ICO](/x/ico/README.md) 