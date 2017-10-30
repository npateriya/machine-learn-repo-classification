####Bot Commands####

######IP Calculations######

***subnetting***

    <user> 192.168.10.5/24
    <bot> subnet: 192.168.10.0, mask: 255.255.255.0, broacast: 192.168.10.255, first: 192.168.10.1, last: 192.168.10.254
    
    <user> next 192.168.10.5/24
    <bot> subnet: 192.168.11.0, mask: 255.255.255.0, broacast: 192.168.11.255, first: 192.168.11.1, last: 192.168.11.254

***binary***

    <user> 1010
    <bot> {"dec":10,"hex":"a"}

***hex***

    <user> 0xFF
    <bot> {"dec":255,"bin":"11111111"}
   

######Quiz######

***View available certs***

    <user> certs
    <bot> icnd1, icnd2, ...
   

***Count database entries***

Syntax: count &lt;*cert*&gt; &lt;[topic]&gt;

Example:

    <user> count icnd1
    <bot> 354
    
    <user> count icnd1 routing
    <bot> 125
   
***Get cert topics***

Syntax: &lt;*cert*&gt; topics

Example:

    <user> topics icnd1
    <bot> ios, security, subnetting, routing

***Start quiz with a random topic***

Syntax: quiz &lt;*cert*&gt;

Example:

    <user> quiz icnd1
    <bot> How do you enable RIP routing?

***Start quiz with a specific topic***

Syntax: quiz &lt;*cert*&gt;

Example:

    <user> quiz icnd1 security
    <bot> This command tells the IOS to prompt for a password

***Answer the quiz***

Be sure to include a ***?*** at the end of your answer

Example:

    <user> quiz icnd1
    <bot> How do you enable RIP routing?
    <user> router rip?
    <bot> good job

***Don't know the answer?***

    <bot> How do you enable RIP routing?
    <user> i give up
    <bot> router rip


