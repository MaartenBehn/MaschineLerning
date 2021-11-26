# Doku

## [matrix.go](matrix.go)

```go
type matrix struct {
	row    int
	collum int
	data   []float64
}
```

`row` ist die Menge and Zeilen der Matrix.  
`collum` ist die Menge and Spalten der Matrix.  
`data` ist der Inhalt der Matrix.  
Oft werden im neuralen Netz Vektoren genutzt. Diese werden auch als Matrizen dargestellt. Bei diesen ist `collum = 1` 

Die restlichen Funktionen von der Matrix Datei sind selbst erklärend.

## [math.go](math.go)

```go
func sigmoidFunc(x float64) float64 {
	return 1 / (1 + math.Pow(math.E, -x))
}
```
Die Signumd-Funktion.

```go
func sigmoidDerivationFunc(x float64) float64 {
	return sigmoidFunc(x) * (1 - sigmoidFunc(x))
}
```
Die Ableitung der Signumd-Funktion.

## [net.go](net.go)
### Net Datenstruktur

```go
type NeuralNet struct {
    hiddenLayersAmmount int
    inputNodes          int
    outputNodes         int
    learnparam          float64
    
    input       *matrix // Vector
    layers      []Layer
    outputLayer *Layer
}

type Layer struct {
    nodesAmount int
    input       *matrix  // Vector
    weights     *matrix 
    netInput    *matrix  // Vector
    output      *matrix  // Vector
    errSig      *matrix  // Vector
    expected    *matrix  // Vector
}
```

Die Datenstruktur des neuralen Netzwerkes besteht aus einem Struct mit einem eingefassten Structs für die Layers.  
`inputNodes` gibt die Menge an input nodes an. Also wie viele input Parameter man hat.  
`hiddenLayersAmmount` gibt die Menge an hidden Layers an.
Je mehr hidden Layers genutzt werden, umso komplexere Aufgaben kann das Netz loesen.  
Aber die rechen Zeit erhoet sich auch signifikant.  
`input` ist ein Vector mit den input Parametern.  
`layers` ist ein slice von allen Layers. Der Outputlayer ist in dem slice enthalten.  
`outputLayer` ist ein pointer auf den letzten Layer in `layers` (Der Outputlayer).  

In dem `Layer` Struct ist `nodesAmount` die Menge an Node von diesem Layer.  
`input` ist eine Matrix der input Werte des Layers.
`output` sind die errechneten Ergebnisse des Layers.
`input` ist der gleiche Pointer wie `output` des vorigem Layers (index -1).  
Im ersten Hiddenlayer ist `input` == `NeuralNet.input`.  
`weights` repräsentiert die Gewichtungen der Verbindungen zum vorigem Layers.
`weights` hat die Groeße `nodesAmount` * `input.row`.  
`netInput` und `errSig` sind zwischen Variablen, die für die backward Propergation benötigt werden.  
`expected` wird nur im Outputlayer genutzt. Dort enthält es die Lösungswerte für den Datensatz.

Alle floats in allen Matrizen sind immer zwischen 0 und 1.  
In einigen Fällen ist bei den Vektorzugriffen index + 1 noetig. 
Dies hängt damit zusammen, das `input`/`output` das bias an index 0 haben und dann erst die Werte ab 1 kommen.

### Net Funktionen

`func NewNet(...) *NeuralNet` initialisiert ein neues Net.

`func setRandomWeigts()` setzt alle `weights` auf eine zufällige Zahl zwischen 0 und 1.

`func print()` und `printOutput()` sind nur zum debuggen und loggen da.

### Forward Path

```go
func (net *NeuralNet) forwardPath() {
    for i := 0; i < net.hiddenLayersAmmount+1; i++ {
        layer := net.layers[i]
        result := layer.weights.Mul(layer.input)
    
            for i := 0; i < result.row; i++ {
                val := result.Get(i, 0)
                layer.netInput.Set(i, 0, val)
                layer.output.Set(i+1, 0, sigmoidFunc(val))
            }
        }
}
```
Rechnet den Forward Path für das Net durch.
Für jeden Layer wird erst `weights` mal `input` genommen.
Dann wird das Ergebnis in `netInput` gespeichert.
`output` enthält die sigmund Funktion des Ergebnisses.

### Backward Path

```go
layer := *net.outputLayer
for i := 0; i < layer.errSig.row; i++ {
    sig := sigmoidDerivationFunc(layer.netInput.Get(i, 0)) *
    (layer.expected.Get(i, 0) - layer.output.Get(i+1, 0))
    
    layer.errSig.Set(i, 0, sig)
}
```

Als Erstes wird `errSig` für den Outputlayer ausgerechnet.  
Dabei wird `netInput` durch die Ableitung der Signumd-Funktion gerechnet.  
Und danach das Ergebnis der Ableitung der Signumd-Funktion mit der Differenz von `output` zu `expected` multipliziert.
Dies wird dann in `errSig` gespeichert.

```go
for k := net.hiddenLayersAmmount - 1; k >= 0; k-- {
    layer := net.layers[k]
    for i := 0; i < layer.errSig.row; i++ {
        sig := sigmoidDerivationFunc(layer.netInput.Get(i, 0))

        a := 0.0
        backLayer := net.layers[k+1]
        for l := 0; l < backLayer.errSig.row; l++ {
            a += backLayer.errSig.Get(l, 0) * backLayer.weights.Get(l, i+1)
        }
        sig *= a

        layer.errSig.Set(i, 0, sig)
    }
}
```

Dann wird `errSig` für alle anderen Layer ausgerechnet.
Dabei wird auch erst `netInput` durch die Ableitung der Signumd-Funktion gerechnet.  
Dann werden alle `weights` des nachkommenden Layers (index +1), die mit der Note verbunden sind, 
mit deren `errSig` multipliziert.
Die Summe all dieser Produkte wird mit dem Ergebnis der Ableitung der Signumd-Funktion multipliziert.
Dies wird dann in `errSig` gespeichert.

```go
for _, layer := range net.layers {
    for i := 0; i < layer.errSig.row; i++ {
        for j := 0; j < layer.input.row; j++ {
            weigthDelta := net.learnparam * layer.errSig.Get(i, 0) * layer.input.Get(j, 0)
            weigth := layer.weights.Get(i, j) + weigthDelta
            layer.weights.Set(i, j, weigth)
        }
    }
}
```

Nun wird noch das `weigthDelta` ausgrechenet, indem der Lernparameter, `errSig` und `input` multipliziert werden.
Die wird auf `weigth` addiert.


## [learn.go](learn.go)

### Learn Settings Datenstruktur
```go
type LearnSettings struct {
    numOfHiddenLayers        int
    numOfNodesInHiddenLayers int
    numOfInputs              int
    datasets                 []*DataSet
    names                    []string
}

type DataSet struct {
    inputs []float64
    name   string
    id     int
}
```

`LearnSettings` ist ein struct der Anweisungen enthält, wie ein neurales Netz gebaut werden soll.  
`numOfHiddenLayers` ist die Menge an Hiddenlayers. Sie muss >= 1 sein.  
`numOfNodesInHiddenLayers` ist die Menge an Notes in jedem Hiddenlayer. Sie muss >= 1 sein.  
`numOfInputs` ist die Menge an input Parameter.
Sie muss mit der Menge an input Werten der Testdaten übereinstimmen.  
`datasets` repräsentiert den Testdatensatz. `name` ist dabei das Ergebnis der Testdaten. 
Somit müssen alle Testdaten mit gleichem Ergebnis auch den gleichen Namen haben.
`names` Ist ein slice aller Namen die in den Testdaten vorkommen.

### Learn Funktionen
`func NewLearnSettings(...) *LearnSettings` initialisiert neue Learn Settings.

`func CreateNetFromLearnSettings(settings *LearnSettings, learparam float64) *NeuralNet` 
erzeugt ein neues Net von `LearnSettings`. 
Dabei wird der `names` slice gefüllt und `DataSet.id` auf den zugehörigen Index gesetzt.

```go
func (net *NeuralNet) loadDataSet(dataset *DataSet) {
    for i, input := range dataset.inputs {
        net.input.Set(i+1, 0, input)
    }
    
    for i := 0; i < net.outputLayer.expected.row; i++ {
        net.outputLayer.expected.Set(i, 0, 0)
    }
    net.outputLayer.expected.Set(dataset.id, 0, 1)
}
```
`loadDataSet` setzt die Inputs des Nets auf die Werte das Datenset. 
Außerdem wird `expected` überall auf 0 gesetzt. Außer bei der Note mit demselben Index wie `id` dort auf 1. 
 
## [file.go](file.go)

`func loadFile(path string) (file *os.File)` gibt die Datei an dem Pfad zurück.

`func loadBmp(path string) (img image.Image)` gibt ein Bildstruct von der .bmp Datei an dem Pfad zurück.

