# Problems and selected solutions

## Why?
Because I want to document what problems I had, what solutions I considered, which one I chose, and why. It will help later, when the project grows bigger.

## Glossary
- **Clicks:** the number of times the user clicked the button.
- **Coins:** the number of virtual tokens the user earned from clicks. This includes power-ups and passive skills, so 1 click can be worth anything from 1 to 100+ coins.

## Sending every click immediately
### Scenario
100 users click for 10 minutes at about 10 clicks per second. (Speed measured with [this site](https://www.rapidtables.com/tools/click-test.html?t=1).)

### Problem
100 users × 10 clicks/s = 1000 requests/s. This is an unnecessarily high load.

### Solutions
**1:** Store the number of clicks on the client side and synchronize with the server every 3-4 seconds. This brings the load down to ~30 requests/s.

**2:** Ask the user to move the earned coins to a stash manually.

I chose **solution 1**: it's better UX, since there are no manual steps.

## Cheating
### Scenario
A user finds the endpoint and sends a huge number of coins.

### Problem
The in-game economy collapses.

### Solution
Cap the coins a batch can earn at `(max user speed + upgrades + passives) × elapsed time in seconds`. This requires storing the last synchronization time for each user.

## Connection lost during synchronization
### Scenario
The user loses the connection right in the middle of synchronization.

### Problem
We can lose the coins the user earned in that moment.

### Solution
A batch that has been sent is **frozen**. New clicks go into a new batch. The old batch is retried with the same UUID until it succeeds.

## Retries cause duplicated batches
### Scenario
The server applied a batch, but the response was lost, so the client retries the same batch twice or more.

### Problem
The user is credited with the same coins several times.

### Solutions
**1:** Every batch gets an increasing ID, and the server checks only whether the ID is bigger than the last one in the database. The problem: two devices can send batches with the same ID, or a device can send a smaller ID, and that batch gets discarded.

**2:** Every batch gets a UUID, which the server stores with a TTL of 24 hours (enough for a clicker). The server deduplicates received batches. This also handles multiple devices.

I chose **solution 2**. The trade-off: it costs storage and has a 24-hour window (a retry after the TTL would be counted twice), but it works across devices.