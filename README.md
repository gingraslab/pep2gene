# pep2gene

Match peptide search results from TPP or MSPLIT to genes.

* [Motivation](#motivation)
* [Rules](#rules)
* [Installation](#installation)
* [Usage](#usage)

## Motivation

Matching peptides to the correct protein and protein isoform can be a challenging task requiring complicated rulesets that can create a confusing picture of the actual sample composition. If the principal interest is actually at the gene level, one way to simplify the interpretation of results is to match peptides to genes, since much of the complexity at the protein level is due to different transcripts arising from a single gene sequence. pep2gene was created to perform this task of matching peptides to genes.

## Rules

Peptides are matched via protein sequence to the corresponding gene identifier. A peptide that matches to multiple proteins that arise from a single gene will count as a single unique peptide match to the gene and any spectral counts for that peptide will be assigned to the gene. If a peptide matches to more than a single gene, whether the peptide gets assigned to the gene, and what portion of its spectral counts get assigned depend on the following rules.

1. If a peptide matches to multiple genes (shared peptide), and there is evidence that each of those genes is present in the sample, i.e. each gene has at least one unique peptide, then the shared peptide will be assigned to all of these matching genes and each gene will get a portion of the shared peptide's spectral counts relative to the evidence for each gene's existence. For example, if one gene has two unique peptides and another has one unique peptide, the first will get twice as many spectral counts from the shared peptide.

2. If a peptide matches to multiple genes but only a subset of those genes have unique peptides, the shared peptide will only be assigned to those genes for which there is definite evidence they are in the sample, i.e. have at least one unique peptide. The spectral counts for the shared peptide will be apportioned as in rule 1.

3. If a gene (A) matches to a peptide or peptides, but those same peptides also match to another gene (B) and that other gene has additional evidence for its existence, gene A is considered to be `subsumed` by B and will be listed as such in the output summary for gene B.

4. If two (or more) genes match to the exact same set of peptides and their is no evidence favouring the presence of one gene over the other, both genes are considered to be present and will evenly split the spectral counts of the shared peptides.

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

For Docker:

```
docker build -t pep2gene -f docker/standard/Dockerfile .
```

For Singularity:

```
docker build -t pep2gene -f docker/singularity/Dockerfile .
```

The Docker image is based from a minimal `scratch` image. While the image is small, it does not work with Singularity, and so Singularity has a slightly different build file.

## Usage

### GO executable

```
pep2gene -db="database.fasta" -file="sample.pepxml" -enzyme="trypsin"
```

### Docker

```
docker run -v $(pwd):/files/ pep2gene -db="database.fasta" -file="sample.pepxml" -enzyme="trypsin"
```

### Singularity

```
singularity run -B ./:/files/ docker://knightjdr/pep2gene:v1.0.0 -db="database.fasta" -file="sample.pepxml" -enzyme="trypsin"
```

The database and peptide file must be located in the working directory Docker/Singularity is called from. Relative or nested paths will not work, i.e. `./some-directory/database.fasta` or `../database.fasta`. The output file will also be written to the working directory.

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

The search database is expected to be in FASTA format, with headers following this convention:
> \>xx|accession|gn|\<gene symbol>:\<Entrez gene ID>

E.G:

> \>gi|22538794|gn|PDCD10:11235| programmed cell death protein 10 [Homo sapiens]

While pep2gene does try to parse the accession and gene name, currently only the gene symbol and gene ID are used.

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