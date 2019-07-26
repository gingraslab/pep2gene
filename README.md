# Gene-Peptide

Match peptide search results from TPP or MSPLIT to genes.

## Installation

This was built as a GO module using go1.12.7. If you have GO installed, you can build/install the binary.
Otherwise it can be run as a container, using Docker for example.

### GO executable

1. Ensure [GO](https://golang.org/doc/install) is installed.

2. Clone repo
```
git clone https://github.com/knightjdr/pep2gene.git
cd pep2gene
```

3. Build executable
```
go build
```

The executable will be called `pep2gene`.

### Docker

#### Build

1. Clone repo
```
git clone https://github.com/knightjdr/pep2gene.git
cd pep2gene
```

2. Build the image
```
docker build -t pep2gene .
```

## Run

### GO executable

```
pep2gene -db="database.fasta" -file="sample.pepxml" -enzyme="trypsin" -missedcleavages="2"
```

### Docker

```
docker run -rm -v $(pwd):/files/ pep2gene -db="database.fasta" -file="sample.pepxml" -enzyme="trypsin" -missedcleavages="2"
```

The database and peptide file must be located in the working directory Docker is called from. Relative or nested paths will not work, i.e. `./some-directory/database.fasta` or `../database.fasta`. The output file will also be written to the working directory.