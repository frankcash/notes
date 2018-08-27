# Introduction to ML and AI in Golang

Gophercon 2018


##  ML and the ML Workflow
[ML and the ML Workflow](https://github.com/ardanlabs/training-ai/tree/master/machine-learning-with-go/ml_intro)

- Start doing experimentation w/ gophernotes
- Eventually move to full fledged program

#### Important Terms

- Data in (images, numbers). "featues".
- Model (modeling some relationships) - Can use a combination of equations and parameters.  "model" takes in "features".
- Data out (label, number) to represent reality. "labels" or "response".

The Model will have a definition: these are the set of equations and the parameters. 

#### Data in => Model => Data Out
#### Data in => f(...) => Data Out (INFERENCE)

#### How to find which parameters to use?

-  Create a training data set
-  Throw training data at models to find which parameters to use. (AKA TRAINING)
-  Try to reduce bias in training data

#### Finding a Use Case

Define a case where there is value to be created by the use of ML/AI. (Identify)
-> Highly repetitive with lots of examples.
-> Might need an expert to help create training or help with the identify step

#### Explore, Select Data
-  Gather Data
-  Profile Data (Good use for Jupyter)
-  Clean / Manipulate Data

#### Train and Test Model(s)
-  Prepare Training and Test Data
-  Define Model(s)
-  Train Model(s): Using training data
-  Test Model(s): Run testing data set through this

Training and Test data is general a split from the training data, so we have a set to see model performance.  This is called "evaluation".

Be wary of over-fitting.  Which is when you make a model super specific to one set of data.

Validation set should be similar percentage of labels as you're expecting in reality.

#### Deploy and Utilize Model(s)
-  Export / Compile Model(s)
-  Deploy Model Inference
-  Deploy Online/Prod Model Training
-  Automation, Monitoring, and Updating

## ML with Go

[Gophernote](https://github.com/gopherdata/gophernotes)

## Building full ML workflow (Lab)

## Concluding Remarks

### Extra stuff to look into

[Go / Fuzzy String Matching](https://github.com/schollz/closestmatch)