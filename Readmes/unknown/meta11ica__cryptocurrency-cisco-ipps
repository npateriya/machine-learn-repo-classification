# What is it ?
It's an Cisco IP Phone Service (IPPS) to retrieve Cryptocurrencies Exchange Rate, there are so far 4 crytocurrencies listed.
# What are those 4 cryptos ?
They are : Bitcoin (BTC), Ethereum (ETH), Litecoin (LTC) and Dogecoin (DOGE).
# What about other crypto-currencies, you don't like 'em huh ?
No I really love everything about this brilliant technology. However, I wrote this IPPS for 4 cryptocurrencies but customizing the IPPS to retrieve other currencies is extremely easy. (I will explain later how to do it)
# What does it need to work ?
You need to place all files in a Web Server supporting PHP, like EasyPHP. The Web Server needs to access https://www.cryptonator.com/ website API (api.cryptonator.com) in order to download the latest exchange rates.
Of course you also need a CUCM deployment in order to implement this service to IP phones. Please refer to the Cisco Unified Communications Manager XML Developers Guide for further information about how to implement Cisco IP phone XML services as well as the compatibility matrix for IP phones. The cryptocurrency-cisco-ipps will basically work on all Cisco phones supporting XML applications. Your IP phones may need a DNS server if the web server is called by its FQDN and not by IP address.

Otherwise, if the service is hosted publicly (like my demo below), you just need the IP phones to reach this public server where you hosted it.

# How does the service work ?
The service retrieves the Exchange Rate in US Dollar of the chosen cryptocurrency. Bitcoin is the default currency, you may select another one while navigating the application through your IP phone.
# How to implement it from PHP server side ?
1- in index.php change the value of the $base_url variable to reflect the folder where the IPPS is hosted, for example:
$base_url = "http://faucet.bitcoin.com.tn/ipps";
This means I created a root folder in my PHP server and put all files in that folder
2- If you would like to an additional cryptocurrency, edit the select.php file and add a section like as part of the appropriate tag (there are 4 "MenuItem" tags each one representing one of the 4 cryptocurrencies.
<MenuItem>
    <Name>Dash</Name>
    <URL>http://faucet.bitcoin.com.tn/ipps/index.php?coin=DASH</URL>
</MenuItem>
Everything is case-sensitive, at least I didn't test but I faced an error when I wrote it downcase.
3- If you would like to change the default currency, change the section 
if (isset($_GET["coin"])) {
   $coin = $_GET["coin"];
}
else {
   $coin = "BTC";
}
to reflect the other default currency for example. For this, in the "else" tag, put rather affect "ETH" to the $coin variable.
That's it!
4- The IP phone service has a built-in cache feature in order to not overwhelm the cryptonator.com API. So the Exchange Rate is refreshed every 60 seconds. The variable $ctime in index.php hold the cache time in seconds e.g. = 60.

# How to implement it from CUCM side ?
Add an IP phone service to CUCM by going to CUCM Administration, meny Device > Device Settings > Phones services
Click add new, give the service a name and put the URL of your web page in the "URL Service" field. If you use HTTPS, put the HTTPS URL in the "Secure URL Service" field. Tick "Enable"
If you would like to make it available for all IP phones in the organization, tick the "Enterprise Subscription" option. and click Save.
Otherwise, you will want to subscribe each specific IP phone to your new IP phone service.

# Is there's a demo ?
Yes! Here's it : https://faucet.bitcoin.com.tn/ipps/
You can use it at will.
