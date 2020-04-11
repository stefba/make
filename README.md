

# Make

Make initiates Gatsby builds via a webhook interface. It takes requests on `/rl/` to run a build. If one is already underway, a second one is queued. Additional requests are dropped until the queue empties.

Furthermore, finished builds are copied to a `pre` folder. The `live` build is then switched out via renaming folders. This way the downtime is reduced from 10-20 seconds to 35µs.
