# pep2gene

Match peptide search results from TPP or MSPLIT to genes.

* [Installation](#installation)
* [Usage](#usage)

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

#### Pull image (and rename - optional)

```
docker pull knightjdr/pep2gene:v1.0.0
docker tag knightjdr/pep2gene:v1.0.0 pep2gene
```

Check for [versions](https://cloud.docker.com/repository/registry-1.docker.io/knightjdr/pep2gene/tags).

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

## Usage

### GO executable

```
pep2gene -db="database.fasta" -file="sample.pepxml" -enzyme="trypsin" -missedcleavages="2"
```

### Docker

```
docker run -rm -v $(pwd):/files/ pep2gene -db="database.fasta" -file="sample.pepxml" -enzyme="trypsin" -missedcleavages="2"
```

The database and peptide file must be located in the working directory Docker is called from. Relative or nested paths will not work, i.e. `./some-directory/database.fasta` or `../database.fasta`. The output file will also be written to the working directory.

### Flags

| Name | Description | Required | Default |
|------|-------------|----------|---------|
| -db | FASTA database | true | |
| -enzyme | digestion enzyme | false | |
| -fdr | MSPLIT peptide FDR | false | 0.01 |
| -file | peptide file | true | |
| -missedcleavages | number of missed cleavages | false | 0 |
| -output | output file format | false | json |
| -pepprob | TPP peptide probability | false | 0.85 |
| -pipeline | search pipeline | false | TPP |

#### Notes

**_-db (database)_**

The search database is expected to be in FASTA format, with headers of the format
> \>xx|accession|gn|\<gene name>:\<Entrez gene id>

E.G:

> \>gi|22538794|gn|PDCD10:11235| programmed cell death protein 10 [Homo sapiens]

**_-enzyme_**

If an enzyme is specified, the sequence database will be digested before peptide matching begins. This significantly speeds up the matching process. If no enzyme is used, peptides are matched against the any protein subsequence.

The available enzymes are:
* arg-c
* asp-n
* asp-n_ambic
* chymotrypsin
* cnbr
* lys-c
* lys-c/p
* lys-n
* pepsina
* trypchymo
* trypsin
* trypsin/p
* v8-de
* v8-e

**_-fdr_**

The FDR is used for parsing high-quality peptides from MSPLIT results, both DDA and DIA. It is ignored when parsing TPP results.

**_-pepprob_**

The peptide probability for parsing high-quality peptides from TPP results. It is ignored when parsing MSPLIT results.

**_-pipeline_**

The analysis pipeline used for searching peptides. The options are:
* MSPLIT_DDA
* MSPLIT_DIA
* TPP