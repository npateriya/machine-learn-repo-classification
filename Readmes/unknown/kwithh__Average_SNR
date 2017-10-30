# Average SNR
<h5>[Initially written for py 2.7.2]</h5>

<strong><p>How it works:</p></strong>
<p>This script will thread an SSH connection to a cisco UBR (in case of multiple UBR's) and compute average SNR for all cable modems.
<ul><li>Output of a command to pull cable upstream labels is parsed into a dictionary.</li>
<li>Each thread will then create a second SSH session to the same UBR to pull upstream SNR by cable modem for every cable interface on the UBR into lists.</li>
<li>It is used to dynamically create indicies in a dictionary by cable interface, and add each modem's SNR to the value of said index.</li>
<li>To compute average, the number of occurrences of each dictionary index inthe list of cable interfaces (pulled with modem upstream SNR value into separate list) is counted, then used as a divisor of each SNR SUM value in the dictionary and print if below a certain value (in this case 30).</li></p>

