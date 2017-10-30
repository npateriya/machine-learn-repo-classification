# data_analysis_note
Cisco Skillzone : data analysis track study notes

## Types of Data Science Questions 
### Descriptive
* The first kind of data analysis performed
* Commonly applied to census data
* The description and interpretation are different steps
* Descriptions can usually not be generalized without additional statistical modeling
Example: http://books.google.com/ngrams

### Exploratory
Goal: Find relationships you didn't know about
* Exploratory models are good for discovering new connections 
* They are also useful for defining future studies
* Exploratory analyses are usually not the final say
* Exploratory analyses alone should not be used for generalizing/predicting
* Correlation does not imply causation
Example: http://www.sdss.org (Sloan Digital Sky Survey)

### Inferential
Goal: Use a relatively small sample of data to say something about a bigger population
* Inference is commonly the goal of statistical models
* Inference involves estimating both the quantity you care about and you uncertainly about your estimate
* Inference depends heavily on both the population and the sampling scheme
Example: Effect of air pollution control on life expectancy in the US: an analysis of 545 US counties for the period from 2000 to 2007

### Predictive
Goal: To use the data on some objects to predict values for another object
* If X predicts Y it does not mean that X causes Y
* Accurate prediction depends heavily on measuring the right variables
* Although there are better and worse prediction models, more data and a simple model works really well
* Prediction is very hard, especially about the future references
Example: 
1. http://fivethirtyeight.blogs.nytimes.com (predict US selections)
2. Target figured out a teen girl was pregnant before her father did

### Causal
Goal: To find out what happens to one variable when you make another variable change.
* Usually randomized studies are required to identify causation
* There are approaches to inferring causation in non-randomized studies, but they are complicated and sensitive to assumptions
* Causal relationships are usually identified as average effects, but may not apply to every individual
* Causal models are usually the "gold standard" for data analysis

### Mechanistic
Goal: Understand the exact changes in variables that lead to changes in other variables for individual obejcts.
* Incredibly hard to infer, except in simple situations
* Usually modeled by a deterministic set of equations (physical/engineering science)
* Generally the random component of the data is measurement error
* If the equations are known but the parameters are not, they may be inferred with data analysis


## In general look at the documentation of APIs
- httr allows GET, POST, PUT, DELETE requests if you are authorized
- You can authenticate with a user name or a password
- Most modern APIs use something like oauth
- httr works well with Facebook, Google, Twitter, Github, etc.

## Reading From Other Sources

### Interacting more directly with files
* file - open a connection to a text file
* url - open a connection to a url 
* gzfile - open a connection to a .gz file
* bzfile - open a connection to a .bz2 file
* ?connections for more information
* Remember to close connections

### foreign package
* Loads data from Minitab, S, SAS, SPSS, Stata, Systat
* Basic functions read.foo
  - read.arff (Weka)
  - read.dta (Stata)
  - read.mtp (Minitab)
  - read.octave (Octave)
  - read.spss (SPSS)
  - read.xport (SAS)
