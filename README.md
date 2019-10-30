[![Build Status](https://travis-ci.org/fpapadopou/poi.svg?branch=master)](https://travis-ci.org/fpapadopou/poi)[![codecov](https://codecov.io/gh/fpapadopou/poi/branch/master/graph/badge.svg)](https://codecov.io/gh/fpapadopou/poi)

# poi
A Go app for submitting &amp; searching places of interest

### Feature list
- [ ] POI type CRUD
- [ ] POI CRUD
- [ ] POI search with several search options like lat/lon combination, area extent (circle and bounded-box), ratings, etc. 

### Environment configuration
- The application can load configuration from `.env` files. A sample with the required
values can be found in `.env.dist`.
- If no `.env` file, system-wide environment variables are used instead.
- If none of the above is present, default (hardcoded) values will be loaded.
