# go-isocountries
A useful list of ISO 3166-1 alpha-3 country code constants for Go.

The package contains a named string type and a whole bunch of constants representing three letter ISO 3166-1 alpha-3 country codes (https://en.wikipedia.org/wiki/ISO_3166-1_alpha-3).

Currently it does not contain special codes like unofficial ones, NATO codes, and disputed zones, but these will be added soon. Another thing you may notice is that the constants are ordered by their code, not by their long name. As some codes are derived from the native-language versions of the country name it may not always be obvious where it appears in the list.
