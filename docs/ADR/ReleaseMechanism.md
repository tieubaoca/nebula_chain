# Release Mechanism Interface

As mentioned in [launchpad](/x/launchpad/README.md), release mechanism is an interface required by a launchpad project.

For maximum flexibility, release mechanism interface should strive to describe as broad of a release mechanism as possible.

In broad term, release mechanism is just a way to release assets of varied attributes and varied quantities to user under a specific mechanism.

Because, each of these release mechanism will have different ways for user to interact with it. Therefore, each release mechanism must include user interaction tx as well.

# Structure overview
1. asset: there can be many asset types in the future. For now, it will be token. NFT will be supported later.
    * In case of NFT, I need project address to store NFTs. I should probably expose project address as well.
1. total_amount: total amount of assets for release.
1. asset_listing_price: an asset needs to have listing price on launchpad in UST.
1. asset_accepted_payment: an asset needs to declare what types of asset it accepts as means of payment. 
    * To ensure that project creator does not accept highly fluctuating token that can hurt user, there will only be a subset of tokens accepted.
    * To update this subset of tokens, governance proposal is required.
    * The only accepted token at beginning will be UST only.
1. amount_released: amount of tokens already released. Project ends when __total_amount__ - __amount_released__ = 0.

# Common Interface (I think that I don't need this) (If each release mechanism maps to a project, then a project doesn't need to store information on what type of mechanism it offers) (If follow this path, there has to be a way for project to query all mechanisms linked to it)
1. QueryAssetAmount(): query the total amount of such asset
1. QueryProgress(): query the progress of the distribution project
1. QueryAssetListingPrice(): query the listing price of an asset
1. QueryAssetAcceptedPayment(): query the coin accepted as means of payment
1. Release(): release logic of an asset

# Implementation
1. [ICO](/x/ico/README.md) 