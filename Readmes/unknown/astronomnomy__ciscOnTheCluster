# ciscOnTheCluster
Some example scripts for Sussex's Apollo cluster useage.


Two examples of qsub scripts that work on an example nifti (a standard MNI image from fsl) and fslmaths and fslstats. 

Both scripts call in the fsl module on the cluster and will submit a job to the cisc queue. One is set up to allow you to do other commands before submitting the job. The other will submit everything to the queue.


### commandlineqsubexample.sh
To run this code, type this in the command line: 

```bash
qsub -q cisc.q -sync n ./commandlineqsubexample.sh
```

qsub is the command to submit to the queue

-q  [name].q : is the queue name. CISC should have access to cisc.q. Others include e.g. mps.q for maths+ physical sciences school. 

-sync y/n : Either the script waits for the job to finish before moving on in your script (y) or it submits the job and frees up the command line, effectively running in the background (n).

./[].sh : The name of the script you're submitting to the queue. They usually have to be set out in a certain way.


Within the script there are a number of arguments that are read in by qsub, preceded by #$

```
#$ -j y
#$ -cwd
#$ -S /bin/bash
#$ -N script_qsub
#$ -pe mpich 2
#$ -o script_qsub.log
```


-j y : join together the error and output streams into one log file

-cwd : start running the code from the directory you're in when you submit the job (otherwise, home directory)

-S /bin/bash : default using the bash shell (maybe you want to use cshell for example) 

-N [name] : The name of the code you want to appear in the queue when you use 'qstat'

-pe mpich 2 : How many cores you want to use (in this case 2 cores of 4GB each). This line also tells the cluster to use mpich code to make these cores talk to each other. The larger this number, the harder it is to find a spot on the queue, so guesstimate wisely! 

-o [name].log : Name of the output and error files


Other arguments you might like:

-M name@email.com : emails you when the code has completed/died in error. Useful for long jobs so you don't have to keep checking it. 

### withinscriptqsubexample.sh
This one can be run simply with
```bash
./withinscriptqsubexample.sh
```

This is good if you want to run some simple checks before hand or small calculations before submitting. If you want to do large calculations before hand please log into the interactive nodes


### interactive node login

qlogin -q interactive.q -l m_mem_free=4G

This is a way to use command line work without running on the login node and getting in trouble! You can change m_mem_free=4G to any number of gigabytes (multiples of 4 are good on apollo) 
