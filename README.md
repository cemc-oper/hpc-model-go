# hpc-model-go

A key-value model for HPC resources in NWPC. Including models for:

* Slurm

## Installation

Use `go get` to install the latest version.

```bash
go get http://www.github.com/perillaroc/hpc-model-go
```

## Getting Started

The following example uses `hpc-model-go` to extract job id and job owner from a `squeue -o %all` query.
 
Create category list:

```go
categoryList := slurm.QueryCategoryList{
    QueryCategoryList: hpcmodel.QueryCategoryList{
        CategoryList: []*hpcmodel.QueryCategory{
            {
                ID:                      "sinfo.partition",
                DisplayName:             "Partition",
                Label:                   "PARTITION",
                PropertyClass:           "StringProperty",
                PropertyCreateArguments: []string{},
                RecordParserClass:       "TokenRecordParser",
            },
            {
                ID:                      "sinfo.avail",
                DisplayName:             "Avail",
                Label:                   "AVAIL",
                PropertyClass:           "StringProperty",
                PropertyCreateArguments: []string{},
                RecordParserClass:       "TokenRecordParser",
            },
            {
                ID:                      "sinfo.nodes",
                DisplayName:             "Nodes(A/I/O/T)",
                Label:                   "NODES(A/I/O/T)",
                PropertyClass:           "StringProperty",
                PropertyCreateArguments: []string{},
                RecordParserClass:       "TokenRecordParser",
            },
            {
                ID:                      "sinfo.cpus",
                DisplayName:             "CPUs(A/I/O/T)",
                Label:                   "CPUS(A/I/O/T)",
                PropertyClass:           "StringProperty",
                PropertyCreateArguments: []string{},
                RecordParserClass:       "TokenRecordParser",
            },
        },
    },
}
```

Get `squeue -o %all` output.

```go
cmd := exec.Command("sinfo", params...)
//fmt.Println(cmd.Args)
var out bytes.Buffer
cmd.Stdout = &out
err := cmd.Run()
if err != nil {
    return nil, fmt.Errorf("command ran error: %v", err)
}
s := out.String()
lines := strings.Split(s, "\n")
```

Build model from category list.

```go
model, err := slurm.BuildModel(lines, categoryList, " ")
```

`model` contains data of all categories.

## Test

Use `go test` to run all tests.

## License

Copyright &copy; 2019, Perilla Roc.

`hpc-model-go` is licensed under [The MIT License](https://opensource.org/licenses/MIT).
