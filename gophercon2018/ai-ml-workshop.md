# Introduction to ML and AI in Golang

Gophercon 2018

August 27, 2018

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
[Gonum](https://github.com/gonum/gonum)
[Gota](https://github.com/kniren/gota)

Reading a CSV with Go using stdlib:

```go
import (
    "os"
    "fmt"
    "encoding/csv"
    "io/ioutil"
)

// Open the csv file at ../../data/5kings_battles_v1.csv.
file, err := os.Open("../data/5kings_battles_v1.csv")
if err != nil {
    fmt.Println(err)
}
// Create a new CSV reader.
reader := csv.NewReader(file)

// Read in all the records via the CSV reader method ReadAll.
records, err := reader.ReadAll()
if err != nil {
    fmt.Println(err)
}

// Close the file.
file.Close()

// Let's get a sense of what the records look like
// by printing a few of them.
for idx, record := range records {
    // Examine the header row.
    if idx == 0 {
        fmt.Println("Header: ", record)
    }else{
        // Print a few of the actual records.
        fmt.Printf("\nname: %s\nyear: %s\nattacker_king: %s\ndefender_king: %s\nattacker_1: %s\n", record[0], record[1], record[3], record[4], record[5])
        if idx > 5 {
            break
        }
    }
}
```

It is important to profile the data before getting started on other things.

Regression: using variable x to model the changes in variable y.

Linear regression assumes data is in a bell curve. https://github.com/gonum/gonum/blob/73ea1e732937f96d723d31dc5263d214a275d204/stat/stat.go#L747

Classification models will output labels to define something.

[Go Learn](https://github.com/sjwhitworth/golearn)

#### Iris KNN example
```go
rawData, err := base.ParseCSVToInstances("../data/iris.csv", true)
if err != nil {
	panic(err)
}

//Initialises a new KNN classifier
cls := knn.NewKnnClassifier("euclidean", "linear", 3)


//Do a training-test split
trainData, testData := base.InstancesTrainTestSplit(rawData, 0.50)
cls.Fit(trainData)

//Calculates the Euclidean distance and returns the most popular label
predictions, err := cls.Predict(testData)
if err != nil {
	panic(err)
}
fmt.Println(predictions)

// Prints precision/recall metrics
confusionMat, err := evaluation.GetConfusionMatrix(testData, predictions)
if err != nil {
	panic(fmt.Sprintf("Unable to get confusion matrix: %s", err.Error()))
}
fmt.Println(evaluation.GetSummary(confusionMat))

```

#### Exercise 3 Solution

```go

irisData, err := base.ParseCSVToInstances("../data/iris.csv", true)
if err != nil {
    fmt.Println(err)
}

for i:= 2; i <10; i++{
    fmt.Println(i)
    knn := knn.NewKnnClassifier("euclidean", "linear", i)
    // Use cross-fold validation to evaluate the kNN model
    // on 5 folds of the data set.
    cv, err := evaluation.GenerateCrossFoldValidationConfusionMatrices(irisData, knn, 5)
    if err != nil {
        fmt.Println(err)
    }

    // Get the mean, variance and standard deviation of the accuracy for the
    // cross validation.
    mean, variance := evaluation.GetCrossValidatedMetric(cv, evaluation.GetAccuracy)
    stdev := math.Sqrt(variance)

    // Output the cross metrics to standard out.
    fmt.Printf("\n\nkNN Accuracy:\n%.2f (+/- %.2f)\n\n", mean, stdev*2)

}
```

## Building full ML workflow (Lab)

## Concluding Remarks

### Extra stuff to look into

[Go / Fuzzy String Matching](https://github.com/schollz/closestmatch)

[Gopherdata](http://gopherdata.io/)