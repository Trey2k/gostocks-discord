# gostocks-discord
goStocks-discord is a discord client that is made to read trading info from a subscribed to channel and make automated trades based off of messages recived.

**NEED TO HAVES:**

- **Ability to cap trading for the day**

Ex. There are some rules about trading that make no sense. In this case, the bot is using a cash account. This is the best account to use unless you have $25,000 or more, which means you can make day-trades to your hearts content.

With a cash account, you can trade through your entire account&#39;s settled cash value you STARTED THE DAY with, and NO MORE. Unsettled cash is cash that hasn&#39;t settled because, with option trades, they do not settle until the following business day. If you trade over the settled cash amount, the account is limited from placing trades for 90 days after the 3rd occurrence within a 12-month calendar period.

Say an account has $2000, and the risk tolerance for one trade is 10%. **This means, IDEALLY, the bot can buy and sell a total of 5 times in one day because a buy/sell each count toward fund usage.** If the bot detects that it cannot make anymore trades without breaking this rule, it will stop for the day.

- **Variable percentage for how much of the total portfolio value the bot can use in one trade.**

Ex. Portfolio value of $1000:

**You only want the bot to risk 10% of your total portfolio in one trade. This means in any one alert the bot can spend a maximum of $100 total.** If the bot cannot afford a minimum of one contract with the total, it will reject making the trade. If the bot can make the trade, it should round up to the closest number of contracts it can get for the total. In this case, a contract may cost $50 each. This means the bot will make a max purchase of 2 contracts.

- **Ability to not take risky/lotto trades.**

Ex. A trade is considered a **lotto or risky** trade when it is formatted in a few different ways:

BTO AMZN 2550C 6/5 @ 1.09 **lotto**

BTO COST 310c 4/9 @ 0.39 **very risky**

in addition to the keyword lotto or risky being present:
Something else that can make the trade risky is the expiration date being the next day. Ex. Today's date is 6/13, but the contract is set to expire 6/14.

Some trades are risky, and risk isn&#39;t for everyone.

- **Ability to set own stop losses.**

Ex. Not all alerts will include a stop loss. This can be dangerous because if the analyst forgets to close the trade or is just doing a risky hold overnight because it&#39;s down, the contract could plummet in value.

Stop losses should be able to be variable depending on your risk tolerance. Say my risk tolerance is relatively high; in this case I might want the stop loss to be 50% of the contracts total value. **This means,**** if for some reason my contract does lose 50% of its value without any analyst setting their own stop loss beforehand, it will market sell the contract once it hits a 50% loss.**

**NICE TO HAVES:**

- **Independent variable percentage of how much total portfolio value can be used in a risky/lotto trade**

Ex. For non-risky/lotto trades, my total portfolio risk in one trade can be 10%. However, with trades that are deemed risky/lottos my total portfolio risk in one trade is only 5%

Risky/lotto trades are inherently more dangerous than those that aren&#39;t, so being able to set a risk tolerance independently for those trades will decreases losses if the trade doesn&#39;t work out.

- **Ability to only make trades that are alerted by a list of specified users**

Ex. Every analyst has their own trading style in the channel. Different risks and different methods are all good, but they aren&#39;t all necessarily favorable. Being able to create a list so the bot only takes trades alerted by specified users would possibly help prevent bad trades.

Let&#39;s say I only wanted to take Blee&#39;s trades. His tag is **Blee#0001.** In this case, only trades announced by him would be acknowledged. All others will be ignored.
