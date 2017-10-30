# CiscoComProductCrawler

When tested this package will go through cisco.com's product catalogue and download all device models it can find.  
At the moment these are `3881` models that get stored under `data/categories` and are later on parsed.  
The output of the test script is `data/parsed.json`.  

To run the tests one must execute the following:
```
time julia -e 'Pkg.test("CiscoComProductCrawler", coverage=true)'
```

This command will scrape the data and parse the `model.html` files, gathering product names, release date, end of sale date and end of support date.  
